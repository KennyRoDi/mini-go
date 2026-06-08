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
		Errors:            []string{},
	}
}

func (v *ScopeVisitor) reportError(msg string, ctx antlr.ParserRuleContext) {
	line := ctx.GetStart().GetLine()
	column := ctx.GetStart().GetColumn()
	v.Errors = append(v.Errors, fmt.Sprintf("ERROR DE ALCANCE: %s [linea:%d - columna:%d]", msg, line, column))
}

func (v *ScopeVisitor) resolveType(ctx parser.IDeclTypeContext) Type {
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
		// Basic support for structs in scope pass
		return StructType{Fields: make(map[string]Type)}
	}
	if ctx.DeclType() != nil {
		return v.resolveType(ctx.DeclType())
	}
	return T_UNKNOWN
}

func (v *ScopeVisitor) VisitFuncArgDecls(ctx *parser.FuncArgDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ScopeVisitor) VisitSingleVarDecl(ctx *parser.SingleVarDeclContext) interface{} {
	var tipo Type = T_UNKNOWN
	if ctx.DeclType() != nil {
		tipo = v.resolveType(ctx.DeclType())
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
	var tipo Type = T_UNKNOWN
	if ctx.DeclType() != nil {
		tipo = v.resolveType(ctx.DeclType())
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
	
	var tipoRetorno Type = T_VOID
	if funcFront.DeclType() != nil {
		tipoRetorno = v.resolveType(funcFront.DeclType())
	}

	// For now, store it as its return type. In Phase 2 we will use FuncType.
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
