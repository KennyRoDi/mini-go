// Code generated from AlphaCompiler.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // AlphaCompiler

import "github.com/antlr4-go/antlr/v4"

// BaseAlphaCompilerListener is a complete listener for a parse tree produced by AlphaCompilerParser.
type BaseAlphaCompilerListener struct{}

var _ AlphaCompilerListener = &BaseAlphaCompilerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAlphaCompilerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAlphaCompilerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAlphaCompilerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAlphaCompilerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseAlphaCompilerListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseAlphaCompilerListener) ExitRoot(ctx *RootContext) {}

// EnterTopDeclarationList is called when production topDeclarationList is entered.
func (s *BaseAlphaCompilerListener) EnterTopDeclarationList(ctx *TopDeclarationListContext) {}

// ExitTopDeclarationList is called when production topDeclarationList is exited.
func (s *BaseAlphaCompilerListener) ExitTopDeclarationList(ctx *TopDeclarationListContext) {}

// EnterVariableDecl is called when production variableDecl is entered.
func (s *BaseAlphaCompilerListener) EnterVariableDecl(ctx *VariableDeclContext) {}

// ExitVariableDecl is called when production variableDecl is exited.
func (s *BaseAlphaCompilerListener) ExitVariableDecl(ctx *VariableDeclContext) {}

// EnterInnerVarDecls is called when production innerVarDecls is entered.
func (s *BaseAlphaCompilerListener) EnterInnerVarDecls(ctx *InnerVarDeclsContext) {}

// ExitInnerVarDecls is called when production innerVarDecls is exited.
func (s *BaseAlphaCompilerListener) ExitInnerVarDecls(ctx *InnerVarDeclsContext) {}

// EnterSingleVarDecl is called when production singleVarDecl is entered.
func (s *BaseAlphaCompilerListener) EnterSingleVarDecl(ctx *SingleVarDeclContext) {}

// ExitSingleVarDecl is called when production singleVarDecl is exited.
func (s *BaseAlphaCompilerListener) ExitSingleVarDecl(ctx *SingleVarDeclContext) {}

// EnterSingleVarDeclNoExps is called when production singleVarDeclNoExps is entered.
func (s *BaseAlphaCompilerListener) EnterSingleVarDeclNoExps(ctx *SingleVarDeclNoExpsContext) {}

// ExitSingleVarDeclNoExps is called when production singleVarDeclNoExps is exited.
func (s *BaseAlphaCompilerListener) ExitSingleVarDeclNoExps(ctx *SingleVarDeclNoExpsContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *BaseAlphaCompilerListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *BaseAlphaCompilerListener) ExitTypeDecl(ctx *TypeDeclContext) {}

// EnterInnerTypeDecls is called when production innerTypeDecls is entered.
func (s *BaseAlphaCompilerListener) EnterInnerTypeDecls(ctx *InnerTypeDeclsContext) {}

// ExitInnerTypeDecls is called when production innerTypeDecls is exited.
func (s *BaseAlphaCompilerListener) ExitInnerTypeDecls(ctx *InnerTypeDeclsContext) {}

// EnterSingleTypeDecl is called when production singleTypeDecl is entered.
func (s *BaseAlphaCompilerListener) EnterSingleTypeDecl(ctx *SingleTypeDeclContext) {}

// ExitSingleTypeDecl is called when production singleTypeDecl is exited.
func (s *BaseAlphaCompilerListener) ExitSingleTypeDecl(ctx *SingleTypeDeclContext) {}

// EnterFuncDecl is called when production funcDecl is entered.
func (s *BaseAlphaCompilerListener) EnterFuncDecl(ctx *FuncDeclContext) {}

// ExitFuncDecl is called when production funcDecl is exited.
func (s *BaseAlphaCompilerListener) ExitFuncDecl(ctx *FuncDeclContext) {}

// EnterFuncFrontDecl is called when production funcFrontDecl is entered.
func (s *BaseAlphaCompilerListener) EnterFuncFrontDecl(ctx *FuncFrontDeclContext) {}

// ExitFuncFrontDecl is called when production funcFrontDecl is exited.
func (s *BaseAlphaCompilerListener) ExitFuncFrontDecl(ctx *FuncFrontDeclContext) {}

// EnterFuncArgDecls is called when production funcArgDecls is entered.
func (s *BaseAlphaCompilerListener) EnterFuncArgDecls(ctx *FuncArgDeclsContext) {}

// ExitFuncArgDecls is called when production funcArgDecls is exited.
func (s *BaseAlphaCompilerListener) ExitFuncArgDecls(ctx *FuncArgDeclsContext) {}

// EnterDeclType is called when production declType is entered.
func (s *BaseAlphaCompilerListener) EnterDeclType(ctx *DeclTypeContext) {}

// ExitDeclType is called when production declType is exited.
func (s *BaseAlphaCompilerListener) ExitDeclType(ctx *DeclTypeContext) {}

// EnterSliceDeclType is called when production sliceDeclType is entered.
func (s *BaseAlphaCompilerListener) EnterSliceDeclType(ctx *SliceDeclTypeContext) {}

// ExitSliceDeclType is called when production sliceDeclType is exited.
func (s *BaseAlphaCompilerListener) ExitSliceDeclType(ctx *SliceDeclTypeContext) {}

// EnterArrayDeclType is called when production arrayDeclType is entered.
func (s *BaseAlphaCompilerListener) EnterArrayDeclType(ctx *ArrayDeclTypeContext) {}

// ExitArrayDeclType is called when production arrayDeclType is exited.
func (s *BaseAlphaCompilerListener) ExitArrayDeclType(ctx *ArrayDeclTypeContext) {}

// EnterStructDeclType is called when production structDeclType is entered.
func (s *BaseAlphaCompilerListener) EnterStructDeclType(ctx *StructDeclTypeContext) {}

// ExitStructDeclType is called when production structDeclType is exited.
func (s *BaseAlphaCompilerListener) ExitStructDeclType(ctx *StructDeclTypeContext) {}

// EnterStructMemDecls is called when production structMemDecls is entered.
func (s *BaseAlphaCompilerListener) EnterStructMemDecls(ctx *StructMemDeclsContext) {}

// ExitStructMemDecls is called when production structMemDecls is exited.
func (s *BaseAlphaCompilerListener) ExitStructMemDecls(ctx *StructMemDeclsContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseAlphaCompilerListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseAlphaCompilerListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *BaseAlphaCompilerListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *BaseAlphaCompilerListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BaseAlphaCompilerListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BaseAlphaCompilerListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterAddExpr is called when production addExpr is entered.
func (s *BaseAlphaCompilerListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production addExpr is exited.
func (s *BaseAlphaCompilerListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterMulExpr is called when production mulExpr is entered.
func (s *BaseAlphaCompilerListener) EnterMulExpr(ctx *MulExprContext) {}

// ExitMulExpr is called when production mulExpr is exited.
func (s *BaseAlphaCompilerListener) ExitMulExpr(ctx *MulExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseAlphaCompilerListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseAlphaCompilerListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterRelExpr is called when production relExpr is entered.
func (s *BaseAlphaCompilerListener) EnterRelExpr(ctx *RelExprContext) {}

// ExitRelExpr is called when production relExpr is exited.
func (s *BaseAlphaCompilerListener) ExitRelExpr(ctx *RelExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseAlphaCompilerListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseAlphaCompilerListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseAlphaCompilerListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseAlphaCompilerListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseAlphaCompilerListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseAlphaCompilerListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseAlphaCompilerListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseAlphaCompilerListener) ExitOperand(ctx *OperandContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseAlphaCompilerListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseAlphaCompilerListener) ExitLiteral(ctx *LiteralContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseAlphaCompilerListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseAlphaCompilerListener) ExitIndex(ctx *IndexContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseAlphaCompilerListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseAlphaCompilerListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseAlphaCompilerListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseAlphaCompilerListener) ExitSelector(ctx *SelectorContext) {}

// EnterAppendExpression is called when production appendExpression is entered.
func (s *BaseAlphaCompilerListener) EnterAppendExpression(ctx *AppendExpressionContext) {}

// ExitAppendExpression is called when production appendExpression is exited.
func (s *BaseAlphaCompilerListener) ExitAppendExpression(ctx *AppendExpressionContext) {}

// EnterLengthExpression is called when production lengthExpression is entered.
func (s *BaseAlphaCompilerListener) EnterLengthExpression(ctx *LengthExpressionContext) {}

// ExitLengthExpression is called when production lengthExpression is exited.
func (s *BaseAlphaCompilerListener) ExitLengthExpression(ctx *LengthExpressionContext) {}

// EnterCapExpression is called when production capExpression is entered.
func (s *BaseAlphaCompilerListener) EnterCapExpression(ctx *CapExpressionContext) {}

// ExitCapExpression is called when production capExpression is exited.
func (s *BaseAlphaCompilerListener) ExitCapExpression(ctx *CapExpressionContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseAlphaCompilerListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseAlphaCompilerListener) ExitStatementList(ctx *StatementListContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseAlphaCompilerListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseAlphaCompilerListener) ExitBlock(ctx *BlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseAlphaCompilerListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseAlphaCompilerListener) ExitStatement(ctx *StatementContext) {}

// EnterSimpleStatement is called when production simpleStatement is entered.
func (s *BaseAlphaCompilerListener) EnterSimpleStatement(ctx *SimpleStatementContext) {}

// ExitSimpleStatement is called when production simpleStatement is exited.
func (s *BaseAlphaCompilerListener) ExitSimpleStatement(ctx *SimpleStatementContext) {}

// EnterAssignmentStatement is called when production assignmentStatement is entered.
func (s *BaseAlphaCompilerListener) EnterAssignmentStatement(ctx *AssignmentStatementContext) {}

// ExitAssignmentStatement is called when production assignmentStatement is exited.
func (s *BaseAlphaCompilerListener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseAlphaCompilerListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseAlphaCompilerListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterLoop is called when production loop is entered.
func (s *BaseAlphaCompilerListener) EnterLoop(ctx *LoopContext) {}

// ExitLoop is called when production loop is exited.
func (s *BaseAlphaCompilerListener) ExitLoop(ctx *LoopContext) {}

// EnterSwitch_stmt is called when production switch_stmt is entered.
func (s *BaseAlphaCompilerListener) EnterSwitch_stmt(ctx *Switch_stmtContext) {}

// ExitSwitch_stmt is called when production switch_stmt is exited.
func (s *BaseAlphaCompilerListener) ExitSwitch_stmt(ctx *Switch_stmtContext) {}

// EnterExpressionCaseClauseList is called when production expressionCaseClauseList is entered.
func (s *BaseAlphaCompilerListener) EnterExpressionCaseClauseList(ctx *ExpressionCaseClauseListContext) {
}

// ExitExpressionCaseClauseList is called when production expressionCaseClauseList is exited.
func (s *BaseAlphaCompilerListener) ExitExpressionCaseClauseList(ctx *ExpressionCaseClauseListContext) {
}

// EnterExpressionCaseClause is called when production expressionCaseClause is entered.
func (s *BaseAlphaCompilerListener) EnterExpressionCaseClause(ctx *ExpressionCaseClauseContext) {}

// ExitExpressionCaseClause is called when production expressionCaseClause is exited.
func (s *BaseAlphaCompilerListener) ExitExpressionCaseClause(ctx *ExpressionCaseClauseContext) {}

// EnterExpressionSwitchCase is called when production expressionSwitchCase is entered.
func (s *BaseAlphaCompilerListener) EnterExpressionSwitchCase(ctx *ExpressionSwitchCaseContext) {}

// ExitExpressionSwitchCase is called when production expressionSwitchCase is exited.
func (s *BaseAlphaCompilerListener) ExitExpressionSwitchCase(ctx *ExpressionSwitchCaseContext) {}
