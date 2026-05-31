// Code generated from AlphaCompiler.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // AlphaCompiler

import "github.com/antlr4-go/antlr/v4"

type BaseAlphaCompilerVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseAlphaCompilerVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitTopDeclarationList(ctx *TopDeclarationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitVariableDecl(ctx *VariableDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitInnerVarDecls(ctx *InnerVarDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitSingleVarDecl(ctx *SingleVarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitSingleVarDeclNoExps(ctx *SingleVarDeclNoExpsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitTypeDecl(ctx *TypeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitInnerTypeDecls(ctx *InnerTypeDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitSingleTypeDecl(ctx *SingleTypeDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitFuncDecl(ctx *FuncDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitFuncFrontDecl(ctx *FuncFrontDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitFuncArgDecls(ctx *FuncArgDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitDeclType(ctx *DeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitSliceDeclType(ctx *SliceDeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitArrayDeclType(ctx *ArrayDeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitStructDeclType(ctx *StructDeclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitStructMemDecls(ctx *StructMemDeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitAddExpr(ctx *AddExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitMulExpr(ctx *MulExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitRelExpr(ctx *RelExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitOperand(ctx *OperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitIndex(ctx *IndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitSelector(ctx *SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitAppendExpression(ctx *AppendExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitLengthExpression(ctx *LengthExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitCapExpression(ctx *CapExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitSimpleStatement(ctx *SimpleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitLoop(ctx *LoopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitSwitch_stmt(ctx *Switch_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitExpressionCaseClauseList(ctx *ExpressionCaseClauseListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitExpressionCaseClause(ctx *ExpressionCaseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseAlphaCompilerVisitor) VisitExpressionSwitchCase(ctx *ExpressionSwitchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}
