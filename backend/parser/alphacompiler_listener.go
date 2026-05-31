// Code generated from AlphaCompiler.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // AlphaCompiler

import "github.com/antlr4-go/antlr/v4"

// AlphaCompilerListener is a complete listener for a parse tree produced by AlphaCompilerParser.
type AlphaCompilerListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterTopDeclarationList is called when entering the topDeclarationList production.
	EnterTopDeclarationList(c *TopDeclarationListContext)

	// EnterVariableDecl is called when entering the variableDecl production.
	EnterVariableDecl(c *VariableDeclContext)

	// EnterInnerVarDecls is called when entering the innerVarDecls production.
	EnterInnerVarDecls(c *InnerVarDeclsContext)

	// EnterSingleVarDecl is called when entering the singleVarDecl production.
	EnterSingleVarDecl(c *SingleVarDeclContext)

	// EnterSingleVarDeclNoExps is called when entering the singleVarDeclNoExps production.
	EnterSingleVarDeclNoExps(c *SingleVarDeclNoExpsContext)

	// EnterTypeDecl is called when entering the typeDecl production.
	EnterTypeDecl(c *TypeDeclContext)

	// EnterInnerTypeDecls is called when entering the innerTypeDecls production.
	EnterInnerTypeDecls(c *InnerTypeDeclsContext)

	// EnterSingleTypeDecl is called when entering the singleTypeDecl production.
	EnterSingleTypeDecl(c *SingleTypeDeclContext)

	// EnterFuncDecl is called when entering the funcDecl production.
	EnterFuncDecl(c *FuncDeclContext)

	// EnterFuncFrontDecl is called when entering the funcFrontDecl production.
	EnterFuncFrontDecl(c *FuncFrontDeclContext)

	// EnterFuncArgDecls is called when entering the funcArgDecls production.
	EnterFuncArgDecls(c *FuncArgDeclsContext)

	// EnterDeclType is called when entering the declType production.
	EnterDeclType(c *DeclTypeContext)

	// EnterSliceDeclType is called when entering the sliceDeclType production.
	EnterSliceDeclType(c *SliceDeclTypeContext)

	// EnterArrayDeclType is called when entering the arrayDeclType production.
	EnterArrayDeclType(c *ArrayDeclTypeContext)

	// EnterStructDeclType is called when entering the structDeclType production.
	EnterStructDeclType(c *StructDeclTypeContext)

	// EnterStructMemDecls is called when entering the structMemDecls production.
	EnterStructMemDecls(c *StructMemDeclsContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterUnaryExpr is called when entering the unaryExpr production.
	EnterUnaryExpr(c *UnaryExprContext)

	// EnterPrimaryExpr is called when entering the primaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterAddExpr is called when entering the addExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterMulExpr is called when entering the mulExpr production.
	EnterMulExpr(c *MulExprContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterRelExpr is called when entering the relExpr production.
	EnterRelExpr(c *RelExprContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterAppendExpression is called when entering the appendExpression production.
	EnterAppendExpression(c *AppendExpressionContext)

	// EnterLengthExpression is called when entering the lengthExpression production.
	EnterLengthExpression(c *LengthExpressionContext)

	// EnterCapExpression is called when entering the capExpression production.
	EnterCapExpression(c *CapExpressionContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterSimpleStatement is called when entering the simpleStatement production.
	EnterSimpleStatement(c *SimpleStatementContext)

	// EnterAssignmentStatement is called when entering the assignmentStatement production.
	EnterAssignmentStatement(c *AssignmentStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterLoop is called when entering the loop production.
	EnterLoop(c *LoopContext)

	// EnterSwitch_stmt is called when entering the switch_stmt production.
	EnterSwitch_stmt(c *Switch_stmtContext)

	// EnterExpressionCaseClauseList is called when entering the expressionCaseClauseList production.
	EnterExpressionCaseClauseList(c *ExpressionCaseClauseListContext)

	// EnterExpressionCaseClause is called when entering the expressionCaseClause production.
	EnterExpressionCaseClause(c *ExpressionCaseClauseContext)

	// EnterExpressionSwitchCase is called when entering the expressionSwitchCase production.
	EnterExpressionSwitchCase(c *ExpressionSwitchCaseContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitTopDeclarationList is called when exiting the topDeclarationList production.
	ExitTopDeclarationList(c *TopDeclarationListContext)

	// ExitVariableDecl is called when exiting the variableDecl production.
	ExitVariableDecl(c *VariableDeclContext)

	// ExitInnerVarDecls is called when exiting the innerVarDecls production.
	ExitInnerVarDecls(c *InnerVarDeclsContext)

	// ExitSingleVarDecl is called when exiting the singleVarDecl production.
	ExitSingleVarDecl(c *SingleVarDeclContext)

	// ExitSingleVarDeclNoExps is called when exiting the singleVarDeclNoExps production.
	ExitSingleVarDeclNoExps(c *SingleVarDeclNoExpsContext)

	// ExitTypeDecl is called when exiting the typeDecl production.
	ExitTypeDecl(c *TypeDeclContext)

	// ExitInnerTypeDecls is called when exiting the innerTypeDecls production.
	ExitInnerTypeDecls(c *InnerTypeDeclsContext)

	// ExitSingleTypeDecl is called when exiting the singleTypeDecl production.
	ExitSingleTypeDecl(c *SingleTypeDeclContext)

	// ExitFuncDecl is called when exiting the funcDecl production.
	ExitFuncDecl(c *FuncDeclContext)

	// ExitFuncFrontDecl is called when exiting the funcFrontDecl production.
	ExitFuncFrontDecl(c *FuncFrontDeclContext)

	// ExitFuncArgDecls is called when exiting the funcArgDecls production.
	ExitFuncArgDecls(c *FuncArgDeclsContext)

	// ExitDeclType is called when exiting the declType production.
	ExitDeclType(c *DeclTypeContext)

	// ExitSliceDeclType is called when exiting the sliceDeclType production.
	ExitSliceDeclType(c *SliceDeclTypeContext)

	// ExitArrayDeclType is called when exiting the arrayDeclType production.
	ExitArrayDeclType(c *ArrayDeclTypeContext)

	// ExitStructDeclType is called when exiting the structDeclType production.
	ExitStructDeclType(c *StructDeclTypeContext)

	// ExitStructMemDecls is called when exiting the structMemDecls production.
	ExitStructMemDecls(c *StructMemDeclsContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitUnaryExpr is called when exiting the unaryExpr production.
	ExitUnaryExpr(c *UnaryExprContext)

	// ExitPrimaryExpr is called when exiting the primaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitAddExpr is called when exiting the addExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitMulExpr is called when exiting the mulExpr production.
	ExitMulExpr(c *MulExprContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitRelExpr is called when exiting the relExpr production.
	ExitRelExpr(c *RelExprContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitAppendExpression is called when exiting the appendExpression production.
	ExitAppendExpression(c *AppendExpressionContext)

	// ExitLengthExpression is called when exiting the lengthExpression production.
	ExitLengthExpression(c *LengthExpressionContext)

	// ExitCapExpression is called when exiting the capExpression production.
	ExitCapExpression(c *CapExpressionContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitSimpleStatement is called when exiting the simpleStatement production.
	ExitSimpleStatement(c *SimpleStatementContext)

	// ExitAssignmentStatement is called when exiting the assignmentStatement production.
	ExitAssignmentStatement(c *AssignmentStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitLoop is called when exiting the loop production.
	ExitLoop(c *LoopContext)

	// ExitSwitch_stmt is called when exiting the switch_stmt production.
	ExitSwitch_stmt(c *Switch_stmtContext)

	// ExitExpressionCaseClauseList is called when exiting the expressionCaseClauseList production.
	ExitExpressionCaseClauseList(c *ExpressionCaseClauseListContext)

	// ExitExpressionCaseClause is called when exiting the expressionCaseClause production.
	ExitExpressionCaseClause(c *ExpressionCaseClauseContext)

	// ExitExpressionSwitchCase is called when exiting the expressionSwitchCase production.
	ExitExpressionSwitchCase(c *ExpressionSwitchCaseContext)
}
