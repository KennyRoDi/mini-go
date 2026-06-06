package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"minigo-backend/parser"
)

type ScopeVisitor struct {
	*parser.BaseminigoVisitor
	Symbols *TablaSimbolos
	Errors  []string
}

func NewScopeVisitor() *ScopeVisitor {
	return &ScopeVisitor{
		BaseminigoVisitor: &parser.BaseminigoVisitor{},
		Symbols:           NewTablaSimbolos(),
		Errors:            []string{},
	}
}

func (v *ScopeVisitor) reportError(msg string, ctx antlr.ParserRuleContext) {
	line := ctx.GetStart().GetLine()
	column := ctx.GetStart().GetColumn()
	v.Errors = append(v.Errors, fmt.Sprintf("ERROR DE ALCANCE: %s [linea:%d - columna:%d]", msg, line, column))
}

func (v *ScopeVisitor) VisitSingleVarDecl(ctx *parser.SingleVarDeclContext) interface{} {
	tipo := T_UNKNOWN
	if ctx.DeclType() != nil {
		tipo = GetMagicType(ctx.DeclType().GetText())
	}
	if ctx.IdentifierList() != nil {
		ids := ctx.IdentifierList().AllIDENTIFIER()
		for _, id := range ids {
			name := id.GetText()
			if err := v.Symbols.Insert(name, tipo); err != nil {
				v.reportError(err.Error(), ctx)
			}
		}
	}
	return v.VisitChildren(ctx)
}

func (v *ScopeVisitor) VisitSingleVarDeclNoExps(ctx *parser.SingleVarDeclNoExpsContext) interface{} {
	tipo := T_UNKNOWN
	if ctx.DeclType() != nil {
		tipo = GetMagicType(ctx.DeclType().GetText())
	}
	if ctx.IdentifierList() != nil {
		ids := ctx.IdentifierList().AllIDENTIFIER()
		for _, id := range ids {
			name := id.GetText()
			if err := v.Symbols.Insert(name, tipo); err != nil {
				v.reportError(err.Error(), ctx)
			}
		}
	}
	return v.VisitChildren(ctx)
}

func (v *ScopeVisitor) VisitFuncDecl(ctx *parser.FuncDeclContext) interface{} {
	funcFront := ctx.FuncFrontDecl().(*parser.FuncFrontDeclContext)
	funcName := funcFront.IDENTIFIER().GetText()
	
	tipoRetorno := T_VOID
	if funcFront.DeclType() != nil {
		tipoRetorno = GetMagicType(funcFront.DeclType().GetText())
	}

	if err := v.Symbols.Insert(funcName, tipoRetorno); err != nil {
		v.reportError(err.Error(), ctx)
	}

	v.Symbols.OpenScope()
	v.VisitChildren(ctx)
	v.Symbols.CloseScope()
	return nil
}

func (v *ScopeVisitor) VisitOperand(ctx *parser.OperandContext) interface{} {
	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		if _, found := v.Symbols.Lookup(name); !found {
			v.reportError(fmt.Sprintf("el identificador '%s' no ha sido declarado", name), ctx)
		}
	}
	return v.VisitChildren(ctx)
}

func (v *ScopeVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	_, isFunc := ctx.GetParent().(*parser.FuncDeclContext)
	if !isFunc {
		v.Symbols.OpenScope()
		defer v.Symbols.CloseScope()
	}
	return v.VisitChildren(ctx)
}



// TODO: Implement Struct, Type, etc. later
