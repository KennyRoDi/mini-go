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
	fmt.Println(errorStr)
	v.Errors = append(v.Errors, errorStr)
	// Eliminamos el panic para permitir la recuperación de errores y evitar el desplome del backend.
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
		typeName := ctx.IDENTIFIER().GetText()
		basic := GetBasicType(typeName)
		if !basic.Equals(T_UNKNOWN) { return basic }
		ident, found := v.Symbols.Lookup(typeName)
		if found { return ident.Tipo }
		return T_UNKNOWN
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
		st := ctx.StructDeclType()
		fields := make(map[string]Type)
		var ordered []string
		if st.StructMemDecls() != nil {
			for _, mem := range st.StructMemDecls().(*parser.StructMemDeclsContext).AllSingleVarDeclNoExps() {
				t := v.resolveType(mem.DeclType())
				for _, id := range mem.IdentifierList().AllIDENTIFIER() {
					name := id.GetText()
					fields[name] = t
					ordered = append(ordered, name)
				}
			}
		}
		return StructType{Fields: fields, OrderedFields: ordered}
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
	var types []Type
	for _, expr := range ctx.AllExpression() {
		res := expr.Accept(v)
		if t, ok := res.(Type); ok {
			types = append(types, t)
		}
	}
	return types
}

func (v *TypeVisitor) VisitArguments(ctx *parser.ArgumentsContext) interface{} {
	if ctx.ExpressionList() != nil {
		return ctx.ExpressionList().Accept(v)
	}
	return []Type{}
}

func (v *TypeVisitor) VisitSingleVarDecl(ctx *parser.SingleVarDeclContext) interface{} {
	var targetType Type = T_UNKNOWN
	if ctx.DeclType() != nil {
		targetType = v.resolveType(ctx.DeclType())
	}
	var exprTypes []Type
	if ctx.ExpressionList() != nil {
		res := ctx.ExpressionList().Accept(v)
		if res != nil {
			exprTypes = res.([]Type)
		}
	}
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
	exprs := ctx.AllExpression()
	if len(exprs) >= 2 {
		left := exprs[0].Accept(v)
		right := exprs[1].Accept(v)
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
	// Also map the expression base
	v.NodeTypes[ctx.GetRuleContext()] = t
	return t
}

func (v *TypeVisitor) VisitMulExpr(ctx *parser.MulExprContext) interface{} {
	r1 := ctx.Expression(0).Accept(v)
	r2 := ctx.Expression(1).Accept(v)
	var t Type = T_UNKNOWN
	if r1 != nil && r2 != nil {
		t1 := r1.(Type)
		t2 := r2.(Type)
		if t1.Equals(T_INT) && t2.Equals(T_INT) {
			t = T_INT
		} else {
			op := ctx.GetOp().GetText()
			v.reportError("Operacion invalida", fmt.Sprintf("'%s' requiere enteros, se encontro %s y %s", op, t1.String(), t2.String()), ctx)
		}
	}
	v.NodeTypes[ctx] = t
	v.NodeTypes[ctx.GetRuleContext()] = t
	return t
}

func (v *TypeVisitor) VisitAndExpr(ctx *parser.AndExprContext) interface{} {
	r1 := ctx.Expression(0).Accept(v)
	r2 := ctx.Expression(1).Accept(v)
	var t Type = T_BOOL
	if r1 != nil && r2 != nil {
		t1 := r1.(Type)
		t2 := r2.(Type)
		if !t1.Equals(T_BOOL) || !t2.Equals(T_BOOL) {
			v.reportError("Operacion logica invalida", "&& requiere booleanos", ctx)
		}
	}
	v.NodeTypes[ctx] = t
	v.NodeTypes[ctx.GetRuleContext()] = t
	return t
}

func (v *TypeVisitor) VisitOrExpr(ctx *parser.OrExprContext) interface{} {
	r1 := ctx.Expression(0).Accept(v)
	r2 := ctx.Expression(1).Accept(v)
	var t Type = T_BOOL
	if r1 != nil && r2 != nil {
		t1 := r1.(Type)
		t2 := r2.(Type)
		if !t1.Equals(T_BOOL) || !t2.Equals(T_BOOL) {
			v.reportError("Operacion logica invalida", "|| requiere booleanos", ctx)
		}
	}
	v.NodeTypes[ctx] = t
	v.NodeTypes[ctx.GetRuleContext()] = t
	return t
}

func (v *TypeVisitor) VisitUnaryExpr(ctx *parser.UnaryExprContext) interface{} {
	res := ctx.Expression().Accept(v)
	var t Type = T_UNKNOWN
	if res != nil {
		t = res.(Type)
		op := ctx.GetChild(0).(antlr.TerminalNode).GetText()
		if op == "!" {
			if !t.Equals(T_BOOL) {
				v.reportError("Operacion unaria invalida", "! requiere booleano", ctx)
				t = T_UNKNOWN
			}
		} else if op == "+" || op == "-" {
			if !t.Equals(T_INT) {
				v.reportError("Operacion unaria invalida", op + " requiere entero", ctx)
				t = T_UNKNOWN
			}
		}
	}
	v.NodeTypes[ctx] = t
	v.NodeTypes[ctx.GetRuleContext()] = t
	return t
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
	v.NodeTypes[ctx] = T_BOOL
	v.NodeTypes[ctx.GetRuleContext()] = T_BOOL
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

func (v *TypeVisitor) VisitSwitch_stmt(ctx *parser.Switch_stmtContext) interface{} {
	if ctx.SimpleStatement() != nil { ctx.SimpleStatement().Accept(v) }
	if ctx.Expression() != nil { ctx.Expression().Accept(v) }
	if ctx.ExpressionCaseClauseList() != nil { ctx.ExpressionCaseClauseList().Accept(v) }
	return nil
}

func (v *TypeVisitor) VisitExpressionCaseClauseList(ctx *parser.ExpressionCaseClauseListContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			node.Accept(v)
		}
	}
	return nil
}

func (v *TypeVisitor) VisitExpressionCaseClause(ctx *parser.ExpressionCaseClauseContext) interface{} {
	if ctx.ExpressionSwitchCase() != nil { ctx.ExpressionSwitchCase().Accept(v) }
	if ctx.StatementList() != nil { ctx.StatementList().Accept(v) }
	return nil
}

func (v *TypeVisitor) VisitExpressionSwitchCase(ctx *parser.ExpressionSwitchCaseContext) interface{} {
	if ctx.ExpressionList() != nil { ctx.ExpressionList().Accept(v) }
	return nil
}

func (v *TypeVisitor) VisitLoop(ctx *parser.LoopContext) interface{} {
	v.Symbols.OpenScope()
	defer v.Symbols.CloseScope()
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok {
			res := node.Accept(v)
			if nodeCtx, ok := node.(parser.IExpressionContext); ok {
				if res != nil && !res.(Type).Equals(T_BOOL) {
					v.reportError("Condicion invalida", "el ciclo requiere una condicion booleana", nodeCtx)
				}
			}
		}
	}
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
	t := ctx.PrimaryExpression().Accept(v).(Type)
	v.NodeTypes[ctx] = t
	v.NodeTypes[ctx.GetRuleContext()] = t
	return t
}

func (v *TypeVisitor) VisitTypeDecl(ctx *parser.TypeDeclContext) interface{} {
	for _, child := range ctx.GetChildren() {
		if st, ok := child.(*parser.SingleTypeDeclContext); ok {
			name := st.IDENTIFIER().GetText()
			t := v.resolveType(st.DeclType())
			v.Symbols.Insert(name, t)
		}
	}
	return nil
}

func (v *TypeVisitor) VisitIndex(ctx *parser.IndexContext) interface{} {
	return ctx.Expression().Accept(v)
}

func (v *TypeVisitor) VisitSelector(ctx *parser.SelectorContext) interface{} {
	return nil
}

func (v *TypeVisitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	var t Type = T_UNKNOWN
	if ctx.Operand() != nil {
		t = ctx.Operand().Accept(v).(Type)
	} else if ctx.PrimaryExpression() != nil {
		res := ctx.PrimaryExpression().Accept(v)
		if res != nil { t = res.(Type) }
	}
	if ctx.Arguments() != nil {
		ctx.Arguments().Accept(v)
	}
	if ctx.Index() != nil {
		ctx.Index().Accept(v)
		if at, ok := t.(ArrayType); ok {
			t = at.Elem
		} else if st, ok := t.(SliceType); ok {
			t = st.Elem
		} else {
			if !t.Equals(T_UNKNOWN) {
				v.reportError("Indexacion invalida", fmt.Sprintf("no se puede indexar tipo %s", t.String()), ctx)
			}
		}
	}
	if ctx.Selector() != nil {
		fieldName := ctx.Selector().IDENTIFIER().GetText()
		if st, ok := t.(StructType); ok {
			if ft, found := st.Fields[fieldName]; found {
				t = ft
			} else {
				v.reportError("Campo no encontrado", fmt.Sprintf("struct no tiene campo '%s'", fieldName), ctx)
			}
		} else {
			if !t.Equals(T_UNKNOWN) {
				v.reportError("Acceso invalido", fmt.Sprintf("tipo %s no tiene campos", t.String()), ctx)
			}
		}
	}
	v.NodeTypes[ctx] = t
	return t
}

func (v *TypeVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	if terminal, ok := ctx.GetChild(0).(antlr.TerminalNode); ok {
		text := terminal.GetText()
		if text == "print" || text == "println" || text == "return" {
			if ctx.ExpressionList() != nil { ctx.ExpressionList().Accept(v) }
			if ctx.Expression() != nil { ctx.Expression().Accept(v) }
			return nil
		}
	}
	for _, child := range ctx.GetChildren() {
		if node, ok := child.(antlr.ParseTree); ok { node.Accept(v) }
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
