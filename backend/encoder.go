package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"minigo-backend/parser"
)

// MiniGoEncoder implements the mandated LLVM code generation.
type MiniGoEncoder struct {
	*parser.BaseminigoVisitor
	Module *ir.Module
	
	// IdentityHashMap as mandated: maps AST node reference to LLVM value reference.
	valorLLVM map[interface{}]value.Value
	
	// Current function and block being built
	currFunc  *ir.Func
	currBlock *ir.Block
	
	// Symbols and types (from previous passes)
	Symbols *TablaSimbolos
	
	// Built-in functions
	printf *ir.Func
}

func NewMiniGoEncoder(symbols *TablaSimbolos) *MiniGoEncoder {
	mod := ir.NewModule()
	
	// Declare printf: i32 printf(i8*, ...)
	printf := mod.NewFunc("printf", types.I32, ir.NewParam("format", types.NewPointer(types.I8)))
	printf.Sig.Variadic = true

	return &MiniGoEncoder{
		BaseminigoVisitor: &parser.BaseminigoVisitor{},
		Module:            mod,
		valorLLVM:         make(map[interface{}]value.Value),
		Symbols:           symbols,
		printf:            printf,
	}
}

func (v *MiniGoEncoder) VisitRoot(ctx *parser.RootContext) interface{} {
	// Automatically synthesize an implicit LLVM main function.
	mainFunc := v.Module.NewFunc("main", types.I32)
	entryBlock := mainFunc.NewBlock("entry")
	v.currFunc = mainFunc
	v.currBlock = entryBlock

	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(v)
		}
	}

	if v.currBlock.Term == nil {
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

func (v *MiniGoEncoder) VisitSingleVarDecl(ctx *parser.SingleVarDeclContext) interface{} {
	for _, id := range ctx.IdentifierList().AllIDENTIFIER() {
		name := id.GetText()
		alloc := v.currBlock.NewAlloca(types.I32)
		alloc.SetName(name)
		v.valorLLVM[id] = alloc
		
		if ctx.ExpressionList() != nil {
			exprList := ctx.ExpressionList().(*parser.ExpressionListContext)
			if len(exprList.AllExpression()) > 0 {
				valResult := exprList.Expression(0).Accept(v)
				if valResult != nil {
					v.currBlock.NewStore(valResult.(value.Value), alloc)
				}
			}
		}
	}
	return nil
}

func (v *MiniGoEncoder) VisitFuncDecl(ctx *parser.FuncDeclContext) interface{} {
	// For this prototype, we visit the body and inline it into 'main'
	// to ensure 'print' statements at the top level work.
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

func (v *MiniGoEncoder) VisitStatement(ctx *parser.StatementContext) interface{} {
	if ctx.GetChildCount() > 0 {
		var childText string
		if terminal, ok := ctx.GetChild(0).(antlr.TerminalNode); ok {
			childText = terminal.GetText()
		}
		
		if childText == "print" || childText == "println" {
			formatName := "fmtInt"
			var formatStr *ir.Global
			for _, g := range v.Module.Globals {
				if g.Name() == formatName {
					formatStr = g
					break
				}
			}
			if formatStr == nil {
				formatStr = v.Module.NewGlobalDef(formatName, constant.NewCharArrayFromString("%d\n\x00"))
			}
			
			if ctx.ExpressionList() != nil {
				exprList := ctx.ExpressionList().(*parser.ExpressionListContext)
				for _, expr := range exprList.AllExpression() {
					valResult := expr.Accept(v)
					if valResult != nil {
						v.currBlock.NewCall(v.printf, formatStr, valResult.(value.Value))
					}
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
	return nil
}

func (v *MiniGoEncoder) VisitOperand(ctx *parser.OperandContext) interface{} {
	if ctx.Literal() != nil {
		return ctx.Literal().Accept(v)
	}
	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		// Manual lookup in block instructions for the allocation
		for _, inst := range v.currBlock.Insts {
			if alloc, ok := inst.(*ir.InstAlloca); ok && alloc.Name() == name {
				return v.currBlock.NewLoad(types.I32, alloc)
			}
		}
	}
	if ctx.Expression() != nil {
		return ctx.Expression().Accept(v)
	}
	return nil
}

func (v *MiniGoEncoder) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.INTLITERAL() != nil {
		var val int64
		fmt.Sscanf(ctx.INTLITERAL().GetText(), "%d", &val)
		return constant.NewInt(types.I32, val)
	}
	return constant.NewInt(types.I32, 0)
}

func (v *MiniGoEncoder) VisitAddExpr(ctx *parser.AddExprContext) interface{} {
	left := ctx.Expression(0).Accept(v)
	right := ctx.Expression(1).Accept(v)
	if left != nil && right != nil {
		return v.currBlock.NewAdd(left.(value.Value), right.(value.Value))
	}
	return nil
}

func (v *MiniGoEncoder) Emit(outPath string) error {
	irContent := v.Module.String()
	irFile := outPath + ".ll"
	err := os.WriteFile(irFile, []byte(irContent), 0644)
	if err != nil {
		return err
	}
	cmd := exec.Command("clang", irFile, "-o", outPath)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
