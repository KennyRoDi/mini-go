package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"minigo-backend/parser"
)

// TypeErrorException is used to halt deep parsing of erroneous branches.
type TypeErrorException struct {
	Message string
}

type TypeVisitor struct {
	*parser.BaseminigoVisitor
	Symbols   *TablaSimbolos
	Errors    []string
	NodeTypes map[antlr.ParseTree]Type
	SymbolMap map[antlr.ParseTree]*Ident
}

func NewTypeVisitor(symbols *TablaSimbolos) *TypeVisitor {
	return &TypeVisitor{
		BaseminigoVisitor: &parser.BaseminigoVisitor{},
		Symbols:           symbols,
		Errors:            []string{},
		NodeTypes:         make(map[antlr.ParseTree]Type),
		SymbolMap:         make(map[antlr.ParseTree]*Ident),
	}
}

func (v *TypeVisitor) reportError(msg string, detail string, ctx antlr.ParserRuleContext) {
	line := ctx.GetStart().GetLine()
	column := ctx.GetStart().GetColumn()
	errorStr := fmt.Sprintf("ERROR DE TIPO: %s (%s) [linea:%d - columna:%d]", msg, detail, line, column)
	v.Errors = append(v.Errors, errorStr)
	panic(TypeErrorException{Message: errorStr})
}

func (v *TypeVisitor) visitIsolated(ctx antlr.ParseTree) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(TypeErrorException); ok {
			} else {
				panic(r)
			}
		}
	}()
	if ctx != nil {
		ctx.Accept(v)
	}
}

func (v *TypeVisitor) resolveType(ctx parser.IDeclTypeContext) Type {
	if ctx == nil { return T_UNKNOWN }
	if ctx.IDENTIFIER() != nil {
		return GetBasicType(ctx.IDENTIFIER().GetText())
	}
	if ctx.SliceDeclType() != nil {
		return SliceType{Elem: v.resolveType(ctx.SliceDeclType().DeclType())}
	}
	if ctx.ArrayDeclType() != nil {
		var size int
		fmt.Sscanf(ctx.ArrayDeclType().INTLITERAL().GetText(), "%d", &size)
		return ArrayType{Elem: v.resolveType(ctx.ArrayDeclType().DeclType()), Size: size}
	}
	if ctx.StructDeclType() != nil {
		return StructType{Fields: make(map[string]Type)}
	}
	if ctx.DeclType() != nil {
		return v.resolveType(ctx.DeclType())
	}
	return T_UNKNOWN
}

// --- VISITOR METHODS ---

func (v *TypeVisitor) VisitRoot(ctx *parser.RootContext) interface{} {
	if ctx.TopDeclarationList() != nil {
		return ctx.TopDeclarationList().Accept(v)
	}
	return nil
}

func (v *TypeVisitor) VisitTopDeclarationList(ctx *parser.TopDeclarationListContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(v)
		}
	}
	return nil
}

func (v *TypeVisitor) VisitVariableDecl(ctx *parser.VariableDeclContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(v)
		}
	}
	return nil
}

func (v *TypeVisitor) VisitStatementList(ctx *parser.StatementListContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(v)
		}
	}
	return nil
}

func (v *TypeVisitor) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(v)
		}
	}
	return nil
}

func (v *TypeVisitor) VisitArguments(ctx *parser.ArgumentsContext) interface{} {
	if ctx.ExpressionList() != nil {
		ctx.ExpressionList().Accept(v)
	}
	return nil
}

func (v *TypeVisitor) VisitSingleVarDecl(ctx *parser.SingleVarDeclContext) interface{} {
	var targetType Type = T_UNKNOWN
	if ctx.DeclType() != nil {
		targetType = v.resolveType(ctx.DeclType())
	}

	var exprTypes []Type
	if ctx.ExpressionList() != nil {
		list := ctx.ExpressionList().(*parser.ExpressionListContext)
		for _, expr := range list.AllExpression() {
			res := expr.Accept(v)
			if res != nil {
				exprTypes = append(exprTypes, res.(Type))
			}
		}
	}

	// Type Inference
	if targetType.Equals(T_UNKNOWN) && len(exprTypes) > 0 {
		targetType = exprTypes[0]
	}

	if ctx.IdentifierList() != nil {
		for _, id := range ctx.IdentifierList().AllIDENTIFIER() {
			v.Symbols.Insert(id.GetText(), targetType)
			v.NodeTypes[id] = targetType
			if ident, found := v.Symbols.Lookup(id.GetText()); found {
				v.SymbolMap[id] = ident
			}
		}
	}

	if !targetType.Equals(T_UNKNOWN) && len(exprTypes) > 0 {
		for _, et := range exprTypes {
			if !targetType.Equals(et) {
				v.reportError("Incompatibilidad de tipos", 
					fmt.Sprintf("no se puede asignar tipo %s a variable de tipo %s", et.String(), targetType.String()), ctx)
			}
		}
	}
	return nil
}

func (v *TypeVisitor) VisitSingleVarDeclNoExps(ctx *parser.SingleVarDeclNoExpsContext) interface{} {
	var targetType Type = T_UNKNOWN
	if ctx.DeclType() != nil {
		targetType = v.resolveType(ctx.DeclType())
	}
	if ctx.IdentifierList() != nil {
		for _, id := range ctx.IdentifierList().AllIDENTIFIER() {
			v.Symbols.Insert(id.GetText(), targetType)
			v.NodeTypes[id] = targetType
			if ident, found := v.Symbols.Lookup(id.GetText()); found {
				v.SymbolMap[id] = ident
			}
		}
	}
	return nil
}

func (v *TypeVisitor) VisitFuncDecl(ctx *parser.FuncDeclContext) interface{} {
	funcFront := ctx.FuncFrontDecl().(*parser.FuncFrontDeclContext)
	funcName := funcFront.IDENTIFIER().GetText()
	
	var tipoRetorno Type = T_VOID
	if funcFront.DeclType() != nil {
		tipoRetorno = v.resolveType(funcFront.DeclType())
	}
	
	v.Symbols.Insert(funcName, tipoRetorno)
	v.SymbolMap[funcFront.IDENTIFIER()] = &Ident{Nombre: funcName, Tipo: tipoRetorno, Nivel: v.Symbols.NivelActual}

	v.Symbols.OpenScope()
	defer v.Symbols.CloseScope()
	if ctx.FuncFrontDecl() != nil { ctx.FuncFrontDecl().Accept(v) }
	if ctx.Block() != nil { ctx.Block().Accept(v) }
	return nil
}

func (v *TypeVisitor) VisitFuncFrontDecl(ctx *parser.FuncFrontDeclContext) interface{} {
	if ctx.FuncArgDecls() != nil {
		ctx.FuncArgDecls().Accept(v)
	}
	return nil
}

func (v *TypeVisitor) VisitFuncArgDecls(ctx *parser.FuncArgDeclsContext) interface{} {
	for _, decl := range ctx.AllSingleVarDeclNoExps() {
		decl.Accept(v)
	}
	return nil
}

func (v *TypeVisitor) VisitAssignmentStatement(ctx *parser.AssignmentStatementContext) interface{} {
	if ctx.GetOp() != nil {
		left := ctx.Expression(0).Accept(v)
		right := ctx.Expression(1).Accept(v)
		if left != nil && right != nil {
			lt := left.(Type)
			rt := right.(Type)
			if !lt.Equals(rt) {
				v.reportError("Incompatibilidad de tipos en asignacion", 
					fmt.Sprintf("tipos %s y %s no coinciden", lt.String(), rt.String()), ctx)
			}
		}
	}
	return nil
}

func (v *TypeVisitor) VisitAddExpr(ctx *parser.AddExprContext) interface{} {
	r1 := ctx.Expression(0).Accept(v)
	r2 := ctx.Expression(1).Accept(v)
	var t Type = T_UNKNOWN
	if r1 != nil && r2 != nil {
		t1 := r1.(Type)
		t2 := r2.(Type)
		if t1.Equals(T_INT) && t2.Equals(T_INT) {
			t = T_INT
		} else if t1.Equals(T_STRING) && t2.Equals(T_STRING) && ctx.GetOp().GetText() == "+" {
			t = T_STRING
		} else {
			v.reportError("Operacion aritmetica invalida", 
				fmt.Sprintf("no se puede operar entre tipos %s y %s", t1.String(), t2.String()), ctx)
		}
	}
	v.NodeTypes[ctx] = t
	return t
}

func (v *TypeVisitor) VisitMulExpr(ctx *parser.MulExprContext) interface{} {
	r1 := ctx.Expression(0).Accept(v)
	r2 := ctx.Expression(1).Accept(v)
	if r1 == nil || r2 == nil { return T_UNKNOWN }
	t1 := r1.(Type)
	t2 := r2.(Type)
	if t1.Equals(T_INT) && t2.Equals(T_INT) {
		return T_INT
	}
	v.reportError("Operacion invalida", "multiplicacion requiere enteros", ctx)
	return T_UNKNOWN
}

func (v *TypeVisitor) VisitUnaryExpr(ctx *parser.UnaryExprContext) interface{} {
	res := ctx.Expression().Accept(v)
	if res == nil { return T_UNKNOWN }
	t := res.(Type)
	op := ctx.GetChild(0).(antlr.TerminalNode).GetText()
	if op == "!" {
		if t.Equals(T_BOOL) { return T_BOOL }
		v.reportError("Operacion unaria invalida", "! requiere booleano", ctx)
	}
	if op == "+" || op == "-" {
		if t.Equals(T_INT) { return T_INT }
		v.reportError("Operacion unaria invalida", op + " requiere entero", ctx)
	}
	return T_UNKNOWN
}

func (v *TypeVisitor) VisitRelExpr(ctx *parser.RelExprContext) interface{} {
	r1 := ctx.Expression(0).Accept(v)
	r2 := ctx.Expression(1).Accept(v)
	if r1 != nil && r2 != nil {
		t1 := r1.(Type)
		t2 := r2.(Type)
		if !t1.Equals(t2) {
			v.reportError("Comparacion invalida", 
				fmt.Sprintf("tipos %s y %s no son comparables", t1.String(), t2.String()), ctx)
		}
	}
	return T_BOOL
}

func (v *TypeVisitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	if ctx.SimpleStatement() != nil { v.visitIsolated(ctx.SimpleStatement()) }
	if ctx.Expression() != nil {
		res := ctx.Expression().Accept(v)
		if res != nil && !res.(Type).Equals(T_BOOL) {
			v.reportError("Condicion no booleana", "if requiere bool", ctx)
		}
	}
	for _, b := range ctx.AllBlock() { v.visitIsolated(b) }
	if ctx.IfStatement() != nil { v.visitIsolated(ctx.IfStatement()) }
	return nil
}

func (v *TypeVisitor) VisitOperand(ctx *parser.OperandContext) interface{} {
	var t Type = T_UNKNOWN
	if ctx.Literal() != nil {
		t = ctx.Literal().Accept(v).(Type)
	} else if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		ident, found := v.Symbols.Lookup(name)
		if found {
			t = ident.Tipo
			v.SymbolMap[ctx.IDENTIFIER()] = ident
		}
	} else if ctx.Expression() != nil {
		t = ctx.Expression().Accept(v).(Type)
	}
	v.NodeTypes[ctx] = t
	return t
}

func (v *TypeVisitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	var t Type = T_UNKNOWN
	if ctx.INTLITERAL() != nil { t = T_INT }
	if ctx.FLOATLITERAL() != nil { t = T_INT } 
	if ctx.RUNELITERAL() != nil { t = T_CHAR }
	if ctx.RAWSTRINGLITERAL() != nil { t = T_STRING }
	if ctx.INTERPRETEDSTRINGLITERAL() != nil { t = T_STRING }
	v.NodeTypes[ctx] = t
	return t
}

func (v *TypeVisitor) VisitPrimaryExpr(ctx *parser.PrimaryExprContext) interface{} {
	return ctx.PrimaryExpression().Accept(v)
}

func (v *TypeVisitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	if ctx.Operand() != nil { return ctx.Operand().Accept(v) }
	
	var t Type = T_UNKNOWN
	if ctx.PrimaryExpression() != nil {
		res := ctx.PrimaryExpression().Accept(v)
		if res != nil { t = res.(Type) }
	}
	
	if ctx.Arguments() != nil {
		ctx.Arguments().Accept(v)
	}
	if ctx.Index() != nil { ctx.Index().Accept(v) }
	if ctx.Selector() != nil { ctx.Selector().Accept(v) }
	
	v.NodeTypes[ctx] = t
	return t
}

func (v *TypeVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	if terminal, ok := ctx.GetChild(0).(antlr.TerminalNode); ok {
		text := terminal.GetText()
		if text == "print" || text == "println" || text == "return" {
			if ctx.ExpressionList() != nil {
				ctx.ExpressionList().Accept(v)
			}
			if ctx.Expression() != nil {
				ctx.Expression().Accept(v)
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

func (v *TypeVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	v.Symbols.OpenScope()
	defer v.Symbols.CloseScope()
	if ctx.StatementList() != nil {
		ctx.StatementList().Accept(v)
	}
	return nil
}
