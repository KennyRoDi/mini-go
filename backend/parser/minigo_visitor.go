// Code generated from backend/minigo.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // minigo
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by minigoParser.
type minigoVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by minigoParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by minigoParser#topDeclarationList.
	VisitTopDeclarationList(ctx *TopDeclarationListContext) interface{}

	// Visit a parse tree produced by minigoParser#variableDecl.
	VisitVariableDecl(ctx *VariableDeclContext) interface{}

	// Visit a parse tree produced by minigoParser#innerVarDecls.
	VisitInnerVarDecls(ctx *InnerVarDeclsContext) interface{}

	// Visit a parse tree produced by minigoParser#singleVarDecl.
	VisitSingleVarDecl(ctx *SingleVarDeclContext) interface{}

	// Visit a parse tree produced by minigoParser#singleVarDeclNoExps.
	VisitSingleVarDeclNoExps(ctx *SingleVarDeclNoExpsContext) interface{}

	// Visit a parse tree produced by minigoParser#typeDecl.
	VisitTypeDecl(ctx *TypeDeclContext) interface{}

	// Visit a parse tree produced by minigoParser#innerTypeDecls.
	VisitInnerTypeDecls(ctx *InnerTypeDeclsContext) interface{}

	// Visit a parse tree produced by minigoParser#singleTypeDecl.
	VisitSingleTypeDecl(ctx *SingleTypeDeclContext) interface{}

	// Visit a parse tree produced by minigoParser#declType.
	VisitDeclType(ctx *DeclTypeContext) interface{}

	// Visit a parse tree produced by minigoParser#sliceDeclType.
	VisitSliceDeclType(ctx *SliceDeclTypeContext) interface{}

	// Visit a parse tree produced by minigoParser#arrayDeclType.
	VisitArrayDeclType(ctx *ArrayDeclTypeContext) interface{}

	// Visit a parse tree produced by minigoParser#structDeclType.
	VisitStructDeclType(ctx *StructDeclTypeContext) interface{}

	// Visit a parse tree produced by minigoParser#structMemDecls.
	VisitStructMemDecls(ctx *StructMemDeclsContext) interface{}

	// Visit a parse tree produced by minigoParser#funcDecl.
	VisitFuncDecl(ctx *FuncDeclContext) interface{}

	// Visit a parse tree produced by minigoParser#funcFrontDecl.
	VisitFuncFrontDecl(ctx *FuncFrontDeclContext) interface{}

	// Visit a parse tree produced by minigoParser#funcArgDecls.
	VisitFuncArgDecls(ctx *FuncArgDeclsContext) interface{}

	// Visit a parse tree produced by minigoParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by minigoParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by minigoParser#unaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by minigoParser#primaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by minigoParser#addExpr.
	VisitAddExpr(ctx *AddExprContext) interface{}

	// Visit a parse tree produced by minigoParser#mulExpr.
	VisitMulExpr(ctx *MulExprContext) interface{}

	// Visit a parse tree produced by minigoParser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by minigoParser#relExpr.
	VisitRelExpr(ctx *RelExprContext) interface{}

	// Visit a parse tree produced by minigoParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by minigoParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by minigoParser#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by minigoParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by minigoParser#index.
	VisitIndex(ctx *IndexContext) interface{}

	// Visit a parse tree produced by minigoParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by minigoParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by minigoParser#appendExpression.
	VisitAppendExpression(ctx *AppendExpressionContext) interface{}

	// Visit a parse tree produced by minigoParser#lengthExpression.
	VisitLengthExpression(ctx *LengthExpressionContext) interface{}

	// Visit a parse tree produced by minigoParser#capExpression.
	VisitCapExpression(ctx *CapExpressionContext) interface{}

	// Visit a parse tree produced by minigoParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by minigoParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by minigoParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by minigoParser#simpleStatement.
	VisitSimpleStatement(ctx *SimpleStatementContext) interface{}

	// Visit a parse tree produced by minigoParser#assignmentStatement.
	VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{}

	// Visit a parse tree produced by minigoParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by minigoParser#loop.
	VisitLoop(ctx *LoopContext) interface{}

	// Visit a parse tree produced by minigoParser#switch_stmt.
	VisitSwitch_stmt(ctx *Switch_stmtContext) interface{}

	// Visit a parse tree produced by minigoParser#expressionCaseClauseList.
	VisitExpressionCaseClauseList(ctx *ExpressionCaseClauseListContext) interface{}

	// Visit a parse tree produced by minigoParser#expressionCaseClause.
	VisitExpressionCaseClause(ctx *ExpressionCaseClauseContext) interface{}

	// Visit a parse tree produced by minigoParser#expressionSwitchCase.
	VisitExpressionSwitchCase(ctx *ExpressionSwitchCaseContext) interface{}
}
