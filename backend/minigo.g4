grammar minigo;

// ---------------------------------------------------------
// 0. REGLA PRINCIPAL
// ---------------------------------------------------------

root
    : 'package' IDENTIFIER ';' topDeclarationList EOF
    ;

// ---------------------------------------------------------
// 1. DECLARACIONES GLOBALES
// ---------------------------------------------------------

topDeclarationList
    : (variableDecl | typeDecl | funcDecl)*
    ;

variableDecl
    : 'var' singleVarDecl ';'
    | 'var' '(' innerVarDecls ')' ';'
    | 'var' '(' ')' ';'
    ;

innerVarDecls
    : singleVarDecl ';' (singleVarDecl ';')*
    ;

singleVarDecl
    : identifierList declType '=' expressionList
    | identifierList '=' expressionList
    | singleVarDeclNoExps
    ;

singleVarDeclNoExps
    : identifierList declType
    ;

// ---------------------------------------------------------
// 2. DECLARACIÓN DE TIPOS
// ---------------------------------------------------------

typeDecl
    : 'type' singleTypeDecl ';'
    | 'type' '(' innerTypeDecls ')' ';'
    | 'type' '(' ')' ';'
    ;

innerTypeDecls
    : singleTypeDecl ';' (singleTypeDecl ';')*
    ;

singleTypeDecl
    : IDENTIFIER declType
    ;

declType
    : '(' declType ')'
    | IDENTIFIER
    | sliceDeclType
    | arrayDeclType
    | structDeclType
    ;

sliceDeclType
    : '[' ']' declType
    ;

arrayDeclType
    : '[' INTLITERAL ']' declType
    ;

structDeclType
    : 'struct' '{' (structMemDecls | ) '}'
    ;

structMemDecls
    : singleVarDeclNoExps ';' (singleVarDeclNoExps ';')*
    ;

// ---------------------------------------------------------
// 3. DECLARACIÓN DE FUNCIONES
// ---------------------------------------------------------

funcDecl
    : funcFrontDecl block ';'
    ;

funcFrontDecl
    : 'func' IDENTIFIER '(' (funcArgDecls | ) ')' (declType | )
    ;

funcArgDecls
    : singleVarDeclNoExps (',' singleVarDeclNoExps)*
    ;

// ---------------------------------------------------------
// 4. EXPRESIONES Y OPERADORES
// ---------------------------------------------------------

identifierList
    : IDENTIFIER (',' IDENTIFIER)*
    ;

expressionList
    : expression (',' expression)*
    ;

expression
    : primaryExpression                                     # primaryExpr
    | expression op=( '*' | '/' | '%' | '<<' | '>>' | '&' | BIT_CLEAR ) expression  # mulExpr
    | expression op=( '+' | '-' | '|' | '^' ) expression    # addExpr
    | expression op=( '==' | '!=' | '<' | '<=' | '>' | '>=' ) expression # relExpr
    | expression '&&' expression                            # andExpr
    | expression '||' expression                            # orExpr
    | '+' expression                                        # unaryExpr
    | '-' expression                                        # unaryExpr
    | '!' expression                                        # unaryExpr
    | '^' expression                                        # unaryExpr
    ;

primaryExpression
    : operand
    | primaryExpression selector
    | primaryExpression index
    | primaryExpression arguments
    | appendExpression
    | lengthExpression
    | capExpression
    ;

operand
    : literal
    | IDENTIFIER
    | '(' expression ')'
    ;

literal
    : INTLITERAL
    | FLOATLITERAL
    | RUNELITERAL
    | RAWSTRINGLITERAL
    | INTERPRETEDSTRINGLITERAL
    ;

index
    : '[' expression ']'
    ;

arguments
    : '(' (expressionList | ) ')'
    ;

selector
    : '.' IDENTIFIER
    ;

// ---------------------------------------------------------
// 5. FUNCIONES INCORPORADAS (BUILT-IN)
// ---------------------------------------------------------

appendExpression
    : 'append' '(' expression ',' expression ')'
    | 'append' '(' expression ',' expression ',' ')'
    ;

lengthExpression
    : 'len' '(' expression ')'
    ;

capExpression
    : 'cap' '(' expression ')'
    ;

// ---------------------------------------------------------
// 6. SENTENCIAS (STATEMENTS) Y BLOQUES
// ---------------------------------------------------------

block
    : '{' statementList '}'
    ;

statementList
    : statement*
    ;

statement
    : 'print' '(' (expressionList | ) ')' ';'
    | 'println' '(' (expressionList | ) ')' ';'
    | 'return' (expression | ) ';'
    | 'break' ';'
    | 'continue' ';'
    | simpleStatement ';'
    | block ';'
    | switch_stmt ';'
    | ifStatement ';'
    | loop ';'
    | typeDecl
    | variableDecl
    ;

simpleStatement
    : // epsilon
    | expression ( '++' | '--' | )
    | assignmentStatement
    | expressionList ':=' expressionList
    ;

assignmentStatement
    : expressionList '=' expressionList
    | expression op=( '+=' | '&=' | '-=' | '|=' | '*=' | '^=' | '<<=' | '>>=' | BIT_CLEAR_ASSIGN | '%=' | '/=' ) expression
    ;

// ---------------------------------------------------------
// 7. ESTRUCTURAS DE CONTROL
// ---------------------------------------------------------

ifStatement
    : 'if' expression block
    | 'if' expression block 'else' ifStatement
    | 'if' expression block 'else' block
    | 'if' simpleStatement ';' expression block
    | 'if' simpleStatement ';' expression block 'else' ifStatement
    | 'if' simpleStatement ';' expression block 'else' block
    ;

loop
    : 'for' block
    | 'for' expression block
    | 'for' simpleStatement ';' expression ';' simpleStatement block
    | 'for' simpleStatement ';' ';' simpleStatement block
    ;

switch_stmt
    : 'switch' simpleStatement ';' expression '{' expressionCaseClauseList '}'
    | 'switch' expression '{' expressionCaseClauseList '}'
    | 'switch' simpleStatement ';' '{' expressionCaseClauseList '}'
    | 'switch' '{' expressionCaseClauseList '}'
    ;

expressionCaseClauseList
    : expressionCaseClause*
    ;

expressionCaseClause
    : expressionSwitchCase ':' statementList
    ;

expressionSwitchCase
    : 'case' expressionList
    | 'default'
    ;

// LEXER RULES

BIT_CLEAR: '&^' ;
BIT_CLEAR_ASSIGN: '&^=' ;

IDENTIFIER: [a-zA-Z_] [a-zA-Z0-9_]* ;

INTLITERAL: [0-9]+ ;

FLOATLITERAL: [0-9]+ '.' [0-9]* | '.' [0-9]+ ;

RUNELITERAL: '\'' ( ~['\\] | '\\' . ) '\'' ;

RAWSTRINGLITERAL: '`' .*? '`' ;

INTERPRETEDSTRINGLITERAL: '"' ( ~["\\] | '\\' . )* '"' ;

WS: [ \t\r\n]+ -> skip ;

COMMENT: '//' ~[\r\n]* -> skip ;
MULTILINE_COMMENT: '/*' .*? '*/' -> skip ;
