// Code generated from backend/minigo.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // minigo
import "github.com/antlr4-go/antlr/v4"

// BaseminigoListener is a complete listener for a parse tree produced by minigoParser.
type BaseminigoListener struct{}

var _ minigoListener = &BaseminigoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseminigoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseminigoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseminigoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseminigoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseminigoListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseminigoListener) ExitRoot(ctx *RootContext) {}

// EnterTopDeclarationList is called when production topDeclarationList is entered.
func (s *BaseminigoListener) EnterTopDeclarationList(ctx *TopDeclarationListContext) {}

// ExitTopDeclarationList is called when production topDeclarationList is exited.
func (s *BaseminigoListener) ExitTopDeclarationList(ctx *TopDeclarationListContext) {}

// EnterVariableDecl is called when production variableDecl is entered.
func (s *BaseminigoListener) EnterVariableDecl(ctx *VariableDeclContext) {}

// ExitVariableDecl is called when production variableDecl is exited.
func (s *BaseminigoListener) ExitVariableDecl(ctx *VariableDeclContext) {}

// EnterInnerVarDecls is called when production innerVarDecls is entered.
func (s *BaseminigoListener) EnterInnerVarDecls(ctx *InnerVarDeclsContext) {}

// ExitInnerVarDecls is called when production innerVarDecls is exited.
func (s *BaseminigoListener) ExitInnerVarDecls(ctx *InnerVarDeclsContext) {}

// EnterSingleVarDecl is called when production singleVarDecl is entered.
func (s *BaseminigoListener) EnterSingleVarDecl(ctx *SingleVarDeclContext) {}

// ExitSingleVarDecl is called when production singleVarDecl is exited.
func (s *BaseminigoListener) ExitSingleVarDecl(ctx *SingleVarDeclContext) {}

// EnterSingleVarDeclNoExps is called when production singleVarDeclNoExps is entered.
func (s *BaseminigoListener) EnterSingleVarDeclNoExps(ctx *SingleVarDeclNoExpsContext) {}

// ExitSingleVarDeclNoExps is called when production singleVarDeclNoExps is exited.
func (s *BaseminigoListener) ExitSingleVarDeclNoExps(ctx *SingleVarDeclNoExpsContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *BaseminigoListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *BaseminigoListener) ExitTypeDecl(ctx *TypeDeclContext) {}

// EnterInnerTypeDecls is called when production innerTypeDecls is entered.
func (s *BaseminigoListener) EnterInnerTypeDecls(ctx *InnerTypeDeclsContext) {}

// ExitInnerTypeDecls is called when production innerTypeDecls is exited.
func (s *BaseminigoListener) ExitInnerTypeDecls(ctx *InnerTypeDeclsContext) {}

// EnterSingleTypeDecl is called when production singleTypeDecl is entered.
func (s *BaseminigoListener) EnterSingleTypeDecl(ctx *SingleTypeDeclContext) {}

// ExitSingleTypeDecl is called when production singleTypeDecl is exited.
func (s *BaseminigoListener) ExitSingleTypeDecl(ctx *SingleTypeDeclContext) {}

// EnterDeclType is called when production declType is entered.
func (s *BaseminigoListener) EnterDeclType(ctx *DeclTypeContext) {}

// ExitDeclType is called when production declType is exited.
func (s *BaseminigoListener) ExitDeclType(ctx *DeclTypeContext) {}

// EnterSliceDeclType is called when production sliceDeclType is entered.
func (s *BaseminigoListener) EnterSliceDeclType(ctx *SliceDeclTypeContext) {}

// ExitSliceDeclType is called when production sliceDeclType is exited.
func (s *BaseminigoListener) ExitSliceDeclType(ctx *SliceDeclTypeContext) {}

// EnterArrayDeclType is called when production arrayDeclType is entered.
func (s *BaseminigoListener) EnterArrayDeclType(ctx *ArrayDeclTypeContext) {}

// ExitArrayDeclType is called when production arrayDeclType is exited.
func (s *BaseminigoListener) ExitArrayDeclType(ctx *ArrayDeclTypeContext) {}

// EnterStructDeclType is called when production structDeclType is entered.
func (s *BaseminigoListener) EnterStructDeclType(ctx *StructDeclTypeContext) {}

// ExitStructDeclType is called when production structDeclType is exited.
func (s *BaseminigoListener) ExitStructDeclType(ctx *StructDeclTypeContext) {}

// EnterStructMemDecls is called when production structMemDecls is entered.
func (s *BaseminigoListener) EnterStructMemDecls(ctx *StructMemDeclsContext) {}

// ExitStructMemDecls is called when production structMemDecls is exited.
func (s *BaseminigoListener) ExitStructMemDecls(ctx *StructMemDeclsContext) {}

// EnterFuncDecl is called when production funcDecl is entered.
func (s *BaseminigoListener) EnterFuncDecl(ctx *FuncDeclContext) {}

// ExitFuncDecl is called when production funcDecl is exited.
func (s *BaseminigoListener) ExitFuncDecl(ctx *FuncDeclContext) {}

// EnterFuncFrontDecl is called when production funcFrontDecl is entered.
func (s *BaseminigoListener) EnterFuncFrontDecl(ctx *FuncFrontDeclContext) {}

// ExitFuncFrontDecl is called when production funcFrontDecl is exited.
func (s *BaseminigoListener) ExitFuncFrontDecl(ctx *FuncFrontDeclContext) {}

// EnterFuncArgDecls is called when production funcArgDecls is entered.
func (s *BaseminigoListener) EnterFuncArgDecls(ctx *FuncArgDeclsContext) {}

// ExitFuncArgDecls is called when production funcArgDecls is exited.
func (s *BaseminigoListener) ExitFuncArgDecls(ctx *FuncArgDeclsContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseminigoListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseminigoListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseminigoListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseminigoListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *BaseminigoListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *BaseminigoListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BaseminigoListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BaseminigoListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterAddExpr is called when production addExpr is entered.
func (s *BaseminigoListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production addExpr is exited.
func (s *BaseminigoListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterMulExpr is called when production mulExpr is entered.
func (s *BaseminigoListener) EnterMulExpr(ctx *MulExprContext) {}

// ExitMulExpr is called when production mulExpr is exited.
func (s *BaseminigoListener) ExitMulExpr(ctx *MulExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseminigoListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseminigoListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterRelExpr is called when production relExpr is entered.
func (s *BaseminigoListener) EnterRelExpr(ctx *RelExprContext) {}

// ExitRelExpr is called when production relExpr is exited.
func (s *BaseminigoListener) ExitRelExpr(ctx *RelExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseminigoListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseminigoListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseminigoListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseminigoListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseminigoListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseminigoListener) ExitOperand(ctx *OperandContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseminigoListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseminigoListener) ExitLiteral(ctx *LiteralContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseminigoListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseminigoListener) ExitIndex(ctx *IndexContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseminigoListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseminigoListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseminigoListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseminigoListener) ExitSelector(ctx *SelectorContext) {}

// EnterAppendExpression is called when production appendExpression is entered.
func (s *BaseminigoListener) EnterAppendExpression(ctx *AppendExpressionContext) {}

// ExitAppendExpression is called when production appendExpression is exited.
func (s *BaseminigoListener) ExitAppendExpression(ctx *AppendExpressionContext) {}

// EnterLengthExpression is called when production lengthExpression is entered.
func (s *BaseminigoListener) EnterLengthExpression(ctx *LengthExpressionContext) {}

// ExitLengthExpression is called when production lengthExpression is exited.
func (s *BaseminigoListener) ExitLengthExpression(ctx *LengthExpressionContext) {}

// EnterCapExpression is called when production capExpression is entered.
func (s *BaseminigoListener) EnterCapExpression(ctx *CapExpressionContext) {}

// ExitCapExpression is called when production capExpression is exited.
func (s *BaseminigoListener) ExitCapExpression(ctx *CapExpressionContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseminigoListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseminigoListener) ExitBlock(ctx *BlockContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseminigoListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseminigoListener) ExitStatementList(ctx *StatementListContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseminigoListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseminigoListener) ExitStatement(ctx *StatementContext) {}

// EnterSimpleStatement is called when production simpleStatement is entered.
func (s *BaseminigoListener) EnterSimpleStatement(ctx *SimpleStatementContext) {}

// ExitSimpleStatement is called when production simpleStatement is exited.
func (s *BaseminigoListener) ExitSimpleStatement(ctx *SimpleStatementContext) {}

// EnterAssignmentStatement is called when production assignmentStatement is entered.
func (s *BaseminigoListener) EnterAssignmentStatement(ctx *AssignmentStatementContext) {}

// ExitAssignmentStatement is called when production assignmentStatement is exited.
func (s *BaseminigoListener) ExitAssignmentStatement(ctx *AssignmentStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseminigoListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseminigoListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterLoop is called when production loop is entered.
func (s *BaseminigoListener) EnterLoop(ctx *LoopContext) {}

// ExitLoop is called when production loop is exited.
func (s *BaseminigoListener) ExitLoop(ctx *LoopContext) {}

// EnterSwitch_stmt is called when production switch_stmt is entered.
func (s *BaseminigoListener) EnterSwitch_stmt(ctx *Switch_stmtContext) {}

// ExitSwitch_stmt is called when production switch_stmt is exited.
func (s *BaseminigoListener) ExitSwitch_stmt(ctx *Switch_stmtContext) {}

// EnterExpressionCaseClauseList is called when production expressionCaseClauseList is entered.
func (s *BaseminigoListener) EnterExpressionCaseClauseList(ctx *ExpressionCaseClauseListContext) {}

// ExitExpressionCaseClauseList is called when production expressionCaseClauseList is exited.
func (s *BaseminigoListener) ExitExpressionCaseClauseList(ctx *ExpressionCaseClauseListContext) {}

// EnterExpressionCaseClause is called when production expressionCaseClause is entered.
func (s *BaseminigoListener) EnterExpressionCaseClause(ctx *ExpressionCaseClauseContext) {}

// ExitExpressionCaseClause is called when production expressionCaseClause is exited.
func (s *BaseminigoListener) ExitExpressionCaseClause(ctx *ExpressionCaseClauseContext) {}

// EnterExpressionSwitchCase is called when production expressionSwitchCase is entered.
func (s *BaseminigoListener) EnterExpressionSwitchCase(ctx *ExpressionSwitchCaseContext) {}

// ExitExpressionSwitchCase is called when production expressionSwitchCase is exited.
func (s *BaseminigoListener) ExitExpressionSwitchCase(ctx *ExpressionSwitchCaseContext) {}
