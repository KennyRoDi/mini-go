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
	Symbols *TablaSimbolos
	Errors  []string
}

func NewTypeVisitor(symbols *TablaSimbolos) *TypeVisitor {
	return &TypeVisitor{
		BaseminigoVisitor: &parser.BaseminigoVisitor{},
		Symbols:           symbols,
		Errors:            []string{},
	}
}

func (v *TypeVisitor) reportError(msg string, detail string, ctx antlr.ParserRuleContext) {
	line := ctx.GetStart().GetLine()
	column := ctx.GetStart().GetColumn()
	// Format strictly according to SPEC.md:
	// ERROR DE TIPO: <mensaje> (<detalle>) [linea:<L> - columna:<C>]
	errorStr := fmt.Sprintf("ERROR DE TIPO: %s (%s) [linea:%d - columna:%d]", msg, detail, line, column)
	v.Errors = append(v.Errors, errorStr)
	panic(TypeErrorException{Message: errorStr})
}

func (v *TypeVisitor) visitIsolated(ctx antlr.ParseTree) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(TypeErrorException); ok {
				// Halt current branch, but continue other branches
			} else {
				panic(r)
			}
		}
	}()
	if ctx != nil {
		ctx.Accept(v)
	}
}

// --- VISITOR METHODS ---

func (v *TypeVisitor) VisitRoot(ctx *parser.RootContext) interface{} {
	if ctx.TopDeclarationList() != nil {
		ctx.TopDeclarationList().Accept(v)
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

func (v *TypeVisitor) VisitSingleVarDecl(ctx *parser.SingleVarDeclContext) interface{} {
	targetType := T_UNKNOWN
	if ctx.DeclType() != nil {
		targetType = GetMagicType(ctx.DeclType().GetText())
	}

	if ctx.IdentifierList() != nil {
		for _, id := range ctx.IdentifierList().AllIDENTIFIER() {
			v.Symbols.Insert(id.GetText(), targetType)
		}
	}

	if ctx.ExpressionList() != nil {
		exprs := ctx.ExpressionList().AllExpression()
		for _, expr := range exprs {
			result := expr.Accept(v)
			if result == nil { continue }
			exprType := result.(int)
			if targetType != T_UNKNOWN && exprType != targetType {
				v.reportError("Incompatibilidad de tipos", 
					fmt.Sprintf("no se puede asignar tipo %d a variable de tipo %d", exprType, targetType), ctx)
			}
		}
	}
	return nil
}

func (v *TypeVisitor) VisitSingleVarDeclNoExps(ctx *parser.SingleVarDeclNoExpsContext) interface{} {
	targetType := T_UNKNOWN
	if ctx.DeclType() != nil {
		targetType = GetMagicType(ctx.DeclType().GetText())
	}

	if ctx.IdentifierList() != nil {
		for _, id := range ctx.IdentifierList().AllIDENTIFIER() {
			v.Symbols.Insert(id.GetText(), targetType)
		}
	}
	return nil
}

func (v *TypeVisitor) VisitFuncDecl(ctx *parser.FuncDeclContext) interface{} {
	v.Symbols.OpenScope()
	if ctx.FuncFrontDecl() != nil {
		ctx.FuncFrontDecl().Accept(v)
	}
	if ctx.Block() != nil {
		ctx.Block().Accept(v)
	}
	v.Symbols.CloseScope()
	return nil
}

func (v *TypeVisitor) VisitFuncFrontDecl(ctx *parser.FuncFrontDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeVisitor) VisitFuncArgDecls(ctx *parser.FuncArgDeclsContext) interface{} {
	for _, decl := range ctx.AllSingleVarDeclNoExps() {
		tipo := GetMagicType(decl.DeclType().GetText())
		for _, id := range decl.IdentifierList().AllIDENTIFIER() {
			v.Symbols.Insert(id.GetText(), tipo)
		}
	}
	return nil
}

func (v *TypeVisitor) VisitAssignmentStatement(ctx *parser.AssignmentStatementContext) interface{} {
	if ctx.GetOp() != nil {
		left := ctx.Expression(0).Accept(v)
		right := ctx.Expression(1).Accept(v)
		if left != nil && right != nil {
			leftType := left.(int)
			rightType := right.(int)
			if leftType != rightType {
				v.reportError("Incompatibilidad de tipos en asignacion", 
					fmt.Sprintf("tipos %d y %d no coinciden", leftType, rightType), ctx)
			}
		}
	}
	return nil
}

func (v *TypeVisitor) VisitPrimaryExpr(ctx *parser.PrimaryExprContext) interface{} {
	return ctx.PrimaryExpression().Accept(v)
}

func (v *TypeVisitor) VisitAddExpr(ctx *parser.AddExprContext) interface{} {
	r1 := ctx.Expression(0).Accept(v)
	r2 := ctx.Expression(1).Accept(v)
	if r1 == nil || r2 == nil { return T_UNKNOWN }
	
	t1 := r1.(int)
	t2 := r2.(int)
	
	if t1 == T_INT && t2 == T_INT {
		return T_INT
	}
	if t1 == T_STRING && t2 == T_STRING && ctx.GetOp().GetText() == "+" {
		return T_STRING
	}
	
	v.reportError("Operacion aritmetica invalida", 
		fmt.Sprintf("no se puede operar entre tipos %d y %d", t1, t2), ctx)
	return T_UNKNOWN
}

func (v *TypeVisitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	if ctx.Operand() != nil {
		return ctx.Operand().Accept(v)
	}
	return T_UNKNOWN
}

func (v *TypeVisitor) VisitOperand(ctx *parser.OperandContext) interface{} {
	if ctx.Literal() != nil {
		return ctx.Literal().Accept(v)
	}
	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		ident, found := v.Symbols.Lookup(name)
		if !found {
			return T_UNKNOWN
		}
		return ident.Tipo
	}
	if ctx.Expression() != nil {
		return ctx.Expression().Accept(v)
	}
	return T_UNKNOWN
}

func (v *TypeVisitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.INTLITERAL() != nil { return T_INT }
	if ctx.FLOATLITERAL() != nil { return T_INT } 
	if ctx.RUNELITERAL() != nil { return T_CHAR }
	if ctx.RAWSTRINGLITERAL() != nil { return T_STRING }
	if ctx.INTERPRETEDSTRINGLITERAL() != nil { return T_STRING }
	return T_UNKNOWN
}

func (v *TypeVisitor) VisitRelExpr(ctx *parser.RelExprContext) interface{} {
	r1 := ctx.Expression(0).Accept(v)
	r2 := ctx.Expression(1).Accept(v)
	if r1 != nil && r2 != nil {
		t1 := r1.(int)
		t2 := r2.(int)
		if t1 != t2 {
			v.reportError("Comparacion invalida", 
				fmt.Sprintf("tipos %d y %d no son comparables", t1, t2), ctx)
		}
	}
	return T_BOOL
}

func (v *TypeVisitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	if ctx.SimpleStatement() != nil {
		ctx.SimpleStatement().Accept(v)
	}
	
	if ctx.Expression() != nil {
		res := ctx.Expression().Accept(v)
		if res != nil {
			condType := res.(int)
			if condType != T_BOOL {
				v.reportError("Condicion no booleana", 
					fmt.Sprintf("el if requiere bool, se encontro %d", condType), ctx)
			}
		}
	}
	
	// Visit the main block
	if ctx.Block(0) != nil {
		ctx.Block(0).Accept(v)
	}
	
	// Visit the 'else' block if present
	if len(ctx.AllBlock()) > 1 {
		ctx.Block(1).Accept(v)
	}
	
	// Visit the 'else if' statement if present
	if ctx.IfStatement() != nil {
		ctx.IfStatement().Accept(v)
	}
	
	return nil
}

func (v *TypeVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	if ctx.GetChildCount() > 0 {
		return ctx.GetChild(0).(antlr.ParseTree).Accept(v)
	}
	return nil
}

func (v *TypeVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	_, isFunc := ctx.GetParent().(*parser.FuncDeclContext)
	if !isFunc {
		v.Symbols.OpenScope()
		defer v.Symbols.CloseScope()
	}
	
	if ctx.StatementList() != nil {
		for _, stmt := range ctx.StatementList().AllStatement() {
			v.visitIsolated(stmt)
		}
	}
	return nil
}
