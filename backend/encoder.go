package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"minigo-backend/parser"
)

type MiniGoEncoder struct {
	*parser.BaseminigoVisitor
	Module *ir.Module
	
	valorLLVM map[interface{}]value.Value
	
	currFunc  *ir.Func
	currBlock *ir.Block
	
	Symbols *TablaSimbolos
	
	printf *ir.Func
	
	// IdentityHashMap as mandated: maps unique symbol instance to its LLVM pointer
	symbolPointers map[*Ident]value.Value
	
	NodeTypes map[antlr.ParseTree]Type
	SymbolMap map[antlr.ParseTree]*Ident
}

func NewMiniGoEncoder(symbols *TablaSimbolos) *MiniGoEncoder {
	mod := ir.NewModule()
	
	printf := mod.NewFunc("printf", types.I32, ir.NewParam("format", types.NewPointer(types.I8)))
	printf.Sig.Variadic = true

	return &MiniGoEncoder{
		BaseminigoVisitor: &parser.BaseminigoVisitor{},
		Module:            mod,
		valorLLVM:         make(map[interface{}]value.Value),
		Symbols:           symbols,
		printf:            printf,
		symbolPointers:    make(map[*Ident]value.Value),
	}
}

func (v *MiniGoEncoder) getLLVMType(t Type) types.Type {
	if t == nil { return types.I32 }
	if t.Equals(T_INT) { return types.I32 }
	if t.Equals(T_BOOL) { return types.I1 }
	if t.Equals(T_CHAR) { return types.I8 }
	if t.Equals(T_STRING) { return types.NewPointer(types.I8) }
	if t.Equals(T_VOID) { return types.Void }
	
	if at, ok := t.(ArrayType); ok {
		return types.NewArray(uint64(at.Size), v.getLLVMType(at.Elem))
	}
	
	if st, ok := t.(StructType); ok {
		var fields []types.Type
		for _, name := range st.OrderedFields {
			fields = append(fields, v.getLLVMType(st.Fields[name]))
		}
		return types.NewStruct(fields...)
	}
	
	return types.I32
}

func (v *MiniGoEncoder) VisitRoot(ctx *parser.RootContext) interface{} {
	if ctx.TopDeclarationList() != nil {
		ctx.TopDeclarationList().Accept(v)
	}
	
	if v.currFunc != nil && v.currFunc.Name() == "main" && v.currBlock.Term == nil {
		v.currBlock.NewRet(constant.NewInt(types.I32, 0))
	}
	
	return nil
}

func (v *MiniGoEncoder) VisitTopDeclarationList(ctx *parser.TopDeclarationListContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(v)
		}
	}
	return nil
}

func (v *MiniGoEncoder) VisitVariableDecl(ctx *parser.VariableDeclContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(v)
		}
	}
	return nil
}

func (v *MiniGoEncoder) VisitBlock(ctx *parser.BlockContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(v)
		}
	}
	return nil
}

func (v *MiniGoEncoder) VisitStatementList(ctx *parser.StatementListContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(v)
		}
	}
	return nil
}

func (v *MiniGoEncoder) VisitSingleVarDeclNoExps(ctx *parser.SingleVarDeclNoExpsContext) interface{} {
	if ctx.IdentifierList() == nil {
		return nil
	}
	for _, id := range ctx.IdentifierList().AllIDENTIFIER() {
		name := id.GetText()
		var t Type = T_UNKNOWN
		if nt, ok := v.NodeTypes[id]; ok {
			t = nt
		}
		llvmType := v.getLLVMType(t)

		ident, hasIdent := v.SymbolMap[id]

		isGlobal := v.currFunc == nil
		var ptr value.Value
		if isGlobal {
			var init constant.Constant
			switch llvmType {
			case types.I1:
				init = constant.NewInt(types.I1, 0)
			case types.I8:
				init = constant.NewInt(types.I8, 0)
			default:
				init = constant.NewInt(types.I32, 0)
			}
			ptr = v.Module.NewGlobalDef(name, init)
		} else {
			alloc := v.currBlock.NewAlloca(llvmType)
			alloc.SetName(name)
			ptr = alloc
		}
		if hasIdent {
			v.symbolPointers[ident] = ptr
		}
		v.valorLLVM[id] = ptr
	}
	return nil
}

func (v *MiniGoEncoder) VisitSingleVarDecl(ctx *parser.SingleVarDeclContext) interface{} {
	if ctx.IdentifierList() == nil {
		if ctx.SingleVarDeclNoExps() != nil {
			return ctx.SingleVarDeclNoExps().Accept(v)
		}
		return nil
	}
	for _, id := range ctx.IdentifierList().AllIDENTIFIER() {
		name := id.GetText()
		var t Type = T_UNKNOWN
		if nt, ok := v.NodeTypes[id]; ok {
			t = nt
		}
		llvmType := v.getLLVMType(t)
		
		var ptr value.Value
		ident, hasIdent := v.SymbolMap[id]
		
		isGlobal := v.currFunc == nil || (v.currFunc.Name() == "main" && hasIdent && ident.Nivel == 0)
		
		if isGlobal {
			var init constant.Constant = constant.NewInt(types.I32, 0)
			if ctx.ExpressionList() != nil {
				expr := ctx.ExpressionList().(*parser.ExpressionListContext).Expression(0)
				if literalCtx, ok := expr.Accept(v).(constant.Constant); ok {
					init = literalCtx
				}
			}
			ptr = v.Module.NewGlobalDef(name, init)
		} else {
			ptr = v.currBlock.NewAlloca(llvmType)
			ptr.(*ir.InstAlloca).SetName(name)
			
			if ctx.ExpressionList() != nil {
				exprList := ctx.ExpressionList().(*parser.ExpressionListContext)
				if len(exprList.AllExpression()) > 0 {
					valResult := exprList.Expression(0).Accept(v)
					if valResult != nil {
						v.currBlock.NewStore(valResult.(value.Value), ptr)
					}
				}
			}
		}
		
		if hasIdent {
			v.symbolPointers[ident] = ptr
		}
		v.valorLLVM[id] = ptr
	}
	return nil
}

func (v *MiniGoEncoder) VisitFuncDecl(ctx *parser.FuncDeclContext) interface{} {
	funcFront := ctx.FuncFrontDecl().(*parser.FuncFrontDeclContext)
	id := funcFront.IDENTIFIER()
	name := id.GetText()
	
	ident, hasIdent := v.SymbolMap[id]
	
	retType := v.getLLVMType(T_VOID)
	if hasIdent {
		retType = v.getLLVMType(ident.Tipo)
	}
	
	if name == "main" {
		retType = types.I32
	}
	
	var params []*ir.Param
	if funcFront.FuncArgDecls() != nil {
		args := funcFront.FuncArgDecls().(*parser.FuncArgDeclsContext)
		for _, decl := range args.AllSingleVarDeclNoExps() {
			t := v.getLLVMType(GetBasicType(decl.DeclType().GetText()))
			for _, id := range decl.IdentifierList().AllIDENTIFIER() {
				params = append(params, ir.NewParam(id.GetText(), t))
			}
		}
	}
	
	f := v.Module.NewFunc(name, retType, params...)
	oldFunc, oldBlock := v.currFunc, v.currBlock
	v.currFunc = f
	v.currBlock = f.NewBlock("entry")
	
	if funcFront.FuncArgDecls() != nil {
		args := funcFront.FuncArgDecls().(*parser.FuncArgDeclsContext)
		pIdx := 0
		for _, decl := range args.AllSingleVarDeclNoExps() {
			for _, id := range decl.IdentifierList().AllIDENTIFIER() {
				p := params[pIdx]
				alloc := v.currBlock.NewAlloca(p.Type())
				alloc.SetName(p.LocalName + "_addr")
				v.currBlock.NewStore(p, alloc)
				if ident, ok := v.SymbolMap[id]; ok {
					v.symbolPointers[ident] = alloc
				}
				pIdx++
			}
		}
	}
	
	ctx.Block().Accept(v)
	
	if v.currBlock.Term == nil {
		if name == "main" {
			v.currBlock.NewRet(constant.NewInt(types.I32, 0))
		} else if retType.Equal(types.Void) {
			v.currBlock.NewRet(nil)
		} else {
			v.currBlock.NewRet(constant.NewInt(types.I32, 0))
		}
	}
	
	v.currFunc, v.currBlock = oldFunc, oldBlock
	return nil
}

func (v *MiniGoEncoder) getPointer(ctx antlr.ParseTree) value.Value {
	if ctx == nil {
		return nil
	}

	switch node := ctx.(type) {
	case *parser.OperandContext:
		if node.IDENTIFIER() != nil {
			term := node.IDENTIFIER()
			if ident, ok := v.SymbolMap[term]; ok {
				return v.symbolPointers[ident]
			}
		}
	case *parser.PrimaryExpressionContext:
		if node.Operand() != nil {
			return v.getPointer(node.Operand())
		}
		if node.PrimaryExpression() != nil {
			basePtr := v.getPointer(node.PrimaryExpression())
			if basePtr == nil {
				return nil
			}
			baseType := basePtr.Type().(*types.PointerType).ElemType
			if node.Index() != nil {
				idxVal := node.Index().(*parser.IndexContext).Expression().Accept(v).(value.Value)
				zero := constant.NewInt(types.I32, 0)
				return v.currBlock.NewGetElementPtr(baseType, basePtr, zero, idxVal)
			}
			if node.Selector() != nil {
				fieldName := node.Selector().(*parser.SelectorContext).IDENTIFIER().GetText()
				t := v.NodeTypes[node.PrimaryExpression()]
				if st, ok := t.(StructType); ok {
					fieldIdx := -1
					for i, name := range st.OrderedFields {
						if name == fieldName {
							fieldIdx = i
							break
						}
					}
					if fieldIdx != -1 {
						zero := constant.NewInt(types.I32, 0)
						targetIdx := constant.NewInt(types.I32, int64(fieldIdx))
						return v.currBlock.NewGetElementPtr(baseType, basePtr, zero, targetIdx)
					}
				}
			}
		}
	case *parser.PrimaryExprContext:
		return v.getPointer(node.PrimaryExpression())
	case parser.IExpressionContext:
		if node.GetChildCount() > 0 {
			if child, ok := node.GetChild(0).(antlr.ParseTree); ok {
				return v.getPointer(child)
			}
		}
	}

	return nil
}

func (v *MiniGoEncoder) VisitAssignmentStatement(ctx *parser.AssignmentStatementContext) interface{} {
	var op string
	if ctx.GetOp() != nil {
		op = ctx.GetOp().GetText()
	} else {
		op = "="
	}

	var ptr value.Value
	var val value.Value

	if op == "=" {
		if len(ctx.AllExpressionList()) >= 2 {
			exprListLeft := ctx.ExpressionList(0).(*parser.ExpressionListContext)
			exprListRight := ctx.ExpressionList(1).(*parser.ExpressionListContext)
			if len(exprListLeft.AllExpression()) > 0 && len(exprListRight.AllExpression()) > 0 {
				ptr = v.getPointer(exprListLeft.Expression(0))
				res := exprListRight.Expression(0).Accept(v)
				if res != nil { val = res.(value.Value) }
			}
		}
	} else {
		ptr = v.getPointer(ctx.Expression(0))
		res := ctx.Expression(1).Accept(v)
		if res != nil { val = res.(value.Value) }
	}

	if ptr == nil || val == nil { return nil }
	
	if op == "=" {
		v.currBlock.NewStore(val, ptr)
	} else {
		oldVal := v.currBlock.NewLoad(ptr.Type().(*types.PointerType).ElemType, ptr)
		var newVal value.Value
		switch op {
		case "+=": newVal = v.currBlock.NewAdd(oldVal, val)
		case "-=": newVal = v.currBlock.NewSub(oldVal, val)
		case "*=": newVal = v.currBlock.NewMul(oldVal, val)
		case "/=": newVal = v.currBlock.NewSDiv(oldVal, val)
		case "<<=": newVal = v.currBlock.NewShl(oldVal, val)
		case ">>=": newVal = v.currBlock.NewAShr(oldVal, val)
		}
		if newVal != nil {
			v.currBlock.NewStore(newVal, ptr)
		}
	}
	return nil
}

func (v *MiniGoEncoder) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	if ctx.SimpleStatement() != nil { ctx.SimpleStatement().Accept(v) }
	
	res := ctx.Expression().Accept(v)
	if res == nil { return nil }
	cond := res.(value.Value)
	
	thenBlock := v.currFunc.NewBlock("if.then")
	elseBlock := v.currFunc.NewBlock("if.else")
	mergeBlock := v.currFunc.NewBlock("if.merge")
	
	v.currBlock.NewCondBr(cond, thenBlock, elseBlock)
	
	v.currBlock = thenBlock
	ctx.Block(0).Accept(v)
	if v.currBlock.Term == nil { v.currBlock.NewBr(mergeBlock) }
	
	v.currBlock = elseBlock
	if len(ctx.AllBlock()) > 1 {
		ctx.Block(1).Accept(v)
	} else if ctx.IfStatement() != nil {
		ctx.IfStatement().Accept(v)
	}
	if v.currBlock.Term == nil { v.currBlock.NewBr(mergeBlock) }
	
	v.currBlock = mergeBlock
	return nil
}

func (v *MiniGoEncoder) VisitLoop(ctx *parser.LoopContext) interface{} {
	if ctx.SimpleStatement(0) != nil {
		ctx.SimpleStatement(0).Accept(v)
	}
	
	condBlock := v.currFunc.NewBlock("loop.cond")
	bodyBlock := v.currFunc.NewBlock("loop.body")
	postBlock := v.currFunc.NewBlock("loop.post")
	exitBlock := v.currFunc.NewBlock("loop.exit")
	
	v.currBlock.NewBr(condBlock)
	v.currBlock = condBlock
	
	var cond value.Value = constant.NewInt(types.I1, 1)
	if ctx.Expression() != nil {
		res := ctx.Expression().Accept(v)
		if res == nil { return nil }
		cond = res.(value.Value)
	}
	v.currBlock.NewCondBr(cond, bodyBlock, exitBlock)
	
	v.currBlock = bodyBlock
	if ctx.Block() != nil {
		ctx.Block().Accept(v)
	}
	if v.currBlock.Term == nil { v.currBlock.NewBr(postBlock) }
	
	v.currBlock = postBlock
	if len(ctx.AllSimpleStatement()) > 1 {
		ctx.SimpleStatement(1).Accept(v)
	}
	v.currBlock.NewBr(condBlock)
	
	v.currBlock = exitBlock
	return nil
}

func (v *MiniGoEncoder) VisitSwitch_stmt(ctx *parser.Switch_stmtContext) interface{} {
	if ctx.SimpleStatement() != nil { ctx.SimpleStatement().Accept(v) }

	if ctx.Expression() == nil {
		return nil
	}
	res := ctx.Expression().Accept(v)
	if res == nil {
		return nil
	}
	switchVal := res.(value.Value)
	
	exitBlock := v.currFunc.NewBlock("switch.exit")
	
	clauses := ctx.ExpressionCaseClauseList().(*parser.ExpressionCaseClauseListContext).AllExpressionCaseClause()
	var defaultClause *parser.ExpressionCaseClauseContext
	
	for _, clause := range clauses {
		caseCtx := clause.(*parser.ExpressionCaseClauseContext)
		esc := caseCtx.ExpressionSwitchCase().(*parser.ExpressionSwitchCaseContext)
		if esc.GetText() == "default" {
			defaultClause = caseCtx
			continue
		}
		
		if esc.ExpressionList() != nil {
			for _, expr := range esc.ExpressionList().(*parser.ExpressionListContext).AllExpression() {
				cRes := expr.Accept(v)
				if cRes == nil { continue }
				caseVal := cRes.(value.Value)
				isMatch := v.currBlock.NewICmp(enum.IPredEQ, switchVal, caseVal)
				matchBlock := v.currFunc.NewBlock("switch.match")
				nextCaseBlock := v.currFunc.NewBlock("switch.next")
				v.currBlock.NewCondBr(isMatch, matchBlock, nextCaseBlock)
				v.currBlock = matchBlock
				if caseCtx.StatementList() != nil {
					caseCtx.StatementList().Accept(v)
				}
				if v.currBlock.Term == nil { v.currBlock.NewBr(exitBlock) }
				v.currBlock = nextCaseBlock
			}
		}
	}
	if defaultClause != nil {
		if defaultClause.StatementList() != nil {
			defaultClause.StatementList().Accept(v)
		}
	}
	if v.currBlock.Term == nil { v.currBlock.NewBr(exitBlock) }
	v.currBlock = exitBlock
	return nil
}

func (v *MiniGoEncoder) VisitSimpleStatement(ctx *parser.SimpleStatementContext) interface{} {
	if ctx.Expression() != nil {
		if ctx.GetChildCount() > 1 {
			if opNode, ok := ctx.GetChild(1).(antlr.TerminalNode); ok {
				op := opNode.GetText()
				if op == "++" || op == "--" {
					ptr := v.getPointer(ctx.Expression())
					if ptr != nil {
						val := v.currBlock.NewLoad(ptr.Type().(*types.PointerType).ElemType, ptr)
						var newVal value.Value
						if op == "++" {
							newVal = v.currBlock.NewAdd(val, constant.NewInt(types.I32, 1))
						} else {
							newVal = v.currBlock.NewSub(val, constant.NewInt(types.I32, 1))
						}
						v.currBlock.NewStore(newVal, ptr)
					}
					return nil
				}
			}
		}
		// Expresión standalone (sin ++ ni --)
		return ctx.Expression().Accept(v)
	}
	if ctx.AssignmentStatement() != nil {
		return ctx.AssignmentStatement().Accept(v)
	}
	// expressionList ':=' expressionList (short variable declaration)
	if ctx.GetChildCount() >= 3 {
		for _, child := range ctx.GetChildren() {
			if node, ok := child.(antlr.ParseTree); ok {
				node.Accept(v)
			}
		}
	}
	return nil
}

func (v *MiniGoEncoder) VisitStatement(ctx *parser.StatementContext) interface{} {
	if terminal, ok := ctx.GetChild(0).(antlr.TerminalNode); ok {
		text := terminal.GetText()
		if text == "print" || text == "println" {
			if ctx.ExpressionList() != nil {
				exprs := ctx.ExpressionList().(*parser.ExpressionListContext).AllExpression()
				for _, expr := range exprs {
					res := expr.Accept(v)
					if res == nil { continue }
					val := res.(value.Value)
					t := v.NodeTypes[expr]
					
					formatStr := "%d"
					if t != nil && t.Equals(T_STRING) { formatStr = "%s" }
					if text == "println" { formatStr += "\n" } else { formatStr += " " }
					
					fmtConst := constant.NewCharArrayFromString(formatStr + "\x00")
					fmtName := fmt.Sprintf("fmt.%d", len(v.Module.Globals))
					fmtGlobal := v.Module.NewGlobalDef(fmtName, fmtConst)
					
					v.currBlock.NewCall(v.printf, fmtGlobal, val)
				}
			}
			return nil
		}
		if text == "return" {
			if ctx.Expression() != nil {
				res := ctx.Expression().Accept(v)
				if res != nil {
					val := res.(value.Value)
					v.currBlock.NewRet(val)
				}
			} else {
				if v.currFunc.Name() == "main" {
					v.currBlock.NewRet(constant.NewInt(types.I32, 0))
				} else {
					v.currBlock.NewRet(nil)
				}
			}
			return nil
		}
	}
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(v)
		}
	}
	return nil
}

func (v *MiniGoEncoder) VisitPrimaryExpr(ctx *parser.PrimaryExprContext) interface{} {
	return ctx.PrimaryExpression().Accept(v)
}

func (v *MiniGoEncoder) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	if ctx.Operand() != nil {
		return ctx.Operand().Accept(v)
	}
	
	if ctx.PrimaryExpression() != nil {
		if ctx.Arguments() != nil {
			funcName := ctx.PrimaryExpression().GetText()
			var target *ir.Func
			for _, f := range v.Module.Funcs {
				if f.Name() == funcName { target = f; break }
			}
			if target != nil {
				var args []value.Value
				argCtx := ctx.Arguments().(*parser.ArgumentsContext)
				if argCtx.ExpressionList() != nil {
					for _, expr := range argCtx.ExpressionList().(*parser.ExpressionListContext).AllExpression() {
						res := expr.Accept(v)
						if res == nil {
							return nil
						}
						args = append(args, res.(value.Value))
					}
				}
				return v.currBlock.NewCall(target, args...)
			}
		}
		
		ptr := v.getPointer(ctx)
		if ptr != nil {
			return v.currBlock.NewLoad(ptr.Type().(*types.PointerType).ElemType, ptr)
		}
	}
	return nil
}

func (v *MiniGoEncoder) VisitUnaryExpr(ctx *parser.UnaryExprContext) interface{} {
	val := ctx.Expression().Accept(v).(value.Value)
	op := ctx.GetChild(0).(antlr.TerminalNode).GetText()
	if op == "-" { return v.currBlock.NewMul(constant.NewInt(types.I32, -1), val) }
	if op == "!" { return v.currBlock.NewXor(val, constant.NewInt(types.I1, 1)) }
	return val
}
func (v *MiniGoEncoder) VisitMulExpr(ctx *parser.MulExprContext) interface{} {
	lRes := ctx.Expression(0).Accept(v)
	rRes := ctx.Expression(1).Accept(v)
	if lRes == nil || rRes == nil {
		return nil
	}
	l := lRes.(value.Value)
	r := rRes.(value.Value)
	op := ctx.GetOp().GetText()
	switch op {
	case "*": return v.currBlock.NewMul(l, r)
	case "/": return v.currBlock.NewSDiv(l, r)
	case "%": return v.currBlock.NewSRem(l, r)
	case "<<": return v.currBlock.NewShl(l, r)
	case ">>": return v.currBlock.NewAShr(l, r)
	case "&": return v.currBlock.NewAnd(l, r)
	}
	return nil
}

func (v *MiniGoEncoder) VisitOperand(ctx *parser.OperandContext) interface{} {
	if ctx.Literal() != nil { return ctx.Literal().Accept(v) }
	if ctx.IDENTIFIER() != nil {
		term := ctx.IDENTIFIER()
		if ident, ok := v.SymbolMap[term]; ok {
			if ptr, ok := v.symbolPointers[ident]; ok {
				return v.currBlock.NewLoad(ptr.Type().(*types.PointerType).ElemType, ptr)
			}
		}
	}
	if ctx.Expression() != nil { return ctx.Expression().Accept(v) }
	return nil
}

func (v *MiniGoEncoder) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.INTLITERAL() != nil {
		var val int64
		fmt.Sscanf(ctx.INTLITERAL().GetText(), "%d", &val)
		return constant.NewInt(types.I32, val)
	}
	if ctx.INTERPRETEDSTRINGLITERAL() != nil {
		text := ctx.INTERPRETEDSTRINGLITERAL().GetText()
		text = text[1 : len(text)-1]
		strConst := constant.NewCharArrayFromString(text + "\x00")
		strGlobal := v.Module.NewGlobalDef(fmt.Sprintf("str.%d", len(v.Module.Globals)), strConst)
		zero := constant.NewInt(types.I32, 0)
		return constant.NewGetElementPtr(strConst.Typ, strGlobal, zero, zero)
	}
	if ctx.RAWSTRINGLITERAL() != nil {
		text := ctx.RAWSTRINGLITERAL().GetText()
		text = text[1 : len(text)-1]
		strConst := constant.NewCharArrayFromString(text + "\x00")
		strGlobal := v.Module.NewGlobalDef(fmt.Sprintf("str.%d", len(v.Module.Globals)), strConst)
		zero := constant.NewInt(types.I32, 0)
		return constant.NewGetElementPtr(strConst.Typ, strGlobal, zero, zero)
	}
	return constant.NewInt(types.I32, 0)
}

func (v *MiniGoEncoder) VisitAddExpr(ctx *parser.AddExprContext) interface{} {
	lRes := ctx.Expression(0).Accept(v)
	rRes := ctx.Expression(1).Accept(v)
	if lRes == nil || rRes == nil {
		return nil
	}
	l := lRes.(value.Value)
	r := rRes.(value.Value)
	op := ctx.GetOp().GetText()
	switch op {
	case "+": return v.currBlock.NewAdd(l, r)
	case "-": return v.currBlock.NewSub(l, r)
	case "|": return v.currBlock.NewOr(l, r)
	case "^": return v.currBlock.NewXor(l, r)
	}
	return nil
}

func (v *MiniGoEncoder) VisitRelExpr(ctx *parser.RelExprContext) interface{} {
	lRes := ctx.Expression(0).Accept(v)
	rRes := ctx.Expression(1).Accept(v)
	if lRes == nil || rRes == nil {
		return nil
	}
	l := lRes.(value.Value)
	r := rRes.(value.Value)
	op := ctx.GetOp().GetText()
	var pred enum.IPred
	switch op {
	case "==": pred = enum.IPredEQ
	case "!=": pred = enum.IPredNE
	case "<": pred = enum.IPredSLT
	case "<=": pred = enum.IPredSLE
	case ">": pred = enum.IPredSGT
	case ">=": pred = enum.IPredSGE
	}
	return v.currBlock.NewICmp(pred, l, r)
}

func (v *MiniGoEncoder) Emit(outPath string) error {
	irContent := v.Module.String()
	irFile := outPath + ".ll"
	os.WriteFile(irFile, []byte(irContent), 0644)
	return exec.Command("clang", irFile, "-o", outPath).Run()
}
