// Code generated from AlphaCompiler.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // AlphaCompiler

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by AlphaCompilerParser.
type AlphaCompilerVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by AlphaCompilerParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#topDeclarationList.
	VisitTopDeclarationList(ctx *TopDeclarationListContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#variableDecl.
	VisitVariableDecl(ctx *VariableDeclContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#innerVarDecls.
	VisitInnerVarDecls(ctx *InnerVarDeclsContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#singleVarDecl.
	VisitSingleVarDecl(ctx *SingleVarDeclContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#singleVarDeclNoExps.
	VisitSingleVarDeclNoExps(ctx *SingleVarDeclNoExpsContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#typeDecl.
	VisitTypeDecl(ctx *TypeDeclContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#innerTypeDecls.
	VisitInnerTypeDecls(ctx *InnerTypeDeclsContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#singleTypeDecl.
	VisitSingleTypeDecl(ctx *SingleTypeDeclContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#funcDecl.
	VisitFuncDecl(ctx *FuncDeclContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#funcFrontDecl.
	VisitFuncFrontDecl(ctx *FuncFrontDeclContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#funcArgDecls.
	VisitFuncArgDecls(ctx *FuncArgDeclsContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#declType.
	VisitDeclType(ctx *DeclTypeContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#sliceDeclType.
	VisitSliceDeclType(ctx *SliceDeclTypeContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#arrayDeclType.
	VisitArrayDeclType(ctx *ArrayDeclTypeContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#structDeclType.
	VisitStructDeclType(ctx *StructDeclTypeContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#structMemDecls.
	VisitStructMemDecls(ctx *StructMemDeclsContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#unaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#primaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#addExpr.
	VisitAddExpr(ctx *AddExprContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#mulExpr.
	VisitMulExpr(ctx *MulExprContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#relExpr.
	VisitRelExpr(ctx *RelExprContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#index.
	VisitIndex(ctx *IndexContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#appendExpression.
	VisitAppendExpression(ctx *AppendExpressionContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#lengthExpression.
	VisitLengthExpression(ctx *LengthExpressionContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#capExpression.
	VisitCapExpression(ctx *CapExpressionContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#simpleStatement.
	VisitSimpleStatement(ctx *SimpleStatementContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#assignmentStatement.
	VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#loop.
	VisitLoop(ctx *LoopContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#switch_stmt.
	VisitSwitch_stmt(ctx *Switch_stmtContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#expressionCaseClauseList.
	VisitExpressionCaseClauseList(ctx *ExpressionCaseClauseListContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#expressionCaseClause.
	VisitExpressionCaseClause(ctx *ExpressionCaseClauseContext) interface{}

	// Visit a parse tree produced by AlphaCompilerParser#expressionSwitchCase.
	VisitExpressionSwitchCase(ctx *ExpressionSwitchCaseContext) interface{}
}
