grammar AlphaCompiler;

/*
 * SECTION 2: THE MINI-GO FORMAL GRAMMAR
 * Translating EBNF to ANTLR4 with native locals for metadata.
 */

// REGLA PRINCIPAL: "root"
root
    : 'package' IDENTIFIER ';' topDeclarationList EOF
    ;

topDeclarationList
    : ( variableDecl | typeDecl | funcDecl )*
    ;

variableDecl
    : 'var' singleVarDecl ';'
    | 'var' '(' innerVarDecls ')' ';'
    | 'var' '(' ')' ';'
    ;

innerVarDecls
    : singleVarDecl ';' ( singleVarDecl ';' )*
    ;

singleVarDecl
    : identifierList declType '=' expressionList
    | identifierList '=' expressionList
    | singleVarDeclNoExps
    ;

singleVarDeclNoExps
    : identifierList declType
    ;

typeDecl
    : 'type' singleTypeDecl ';'
    | 'type' '(' innerTypeDecls ')' ';'
    | 'type' '(' ')' ';'
    ;

innerTypeDecls
    : singleTypeDecl ';' ( singleTypeDecl ';' )*
    ;

singleTypeDecl
    : IDENTIFIER declType
    ;

funcDecl
    : funcFrontDecl block ';'
    ;

funcFrontDecl
    : 'func' IDENTIFIER '(' ( funcArgDecls | ) ')' ( declType | )
    ;

funcArgDecls
    : singleVarDeclNoExps ( ',' singleVarDeclNoExps )*
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
    : 'struct' '{' ( structMemDecls | ) '}'
    ;

structMemDecls
    : singleVarDeclNoExps ';' ( singleVarDeclNoExps ';' )*
    ;

identifierList
    : IDENTIFIER ( ',' IDENTIFIER )*
    ;

expression
    : primaryExpression                                     # primaryExpr
    | '+' expression                                        # unaryExpr
    | '-' expression                                        # unaryExpr
    | '!' expression                                        # unaryExpr
    | '^' expression                                        # unaryExpr
    | expression op=( '*' | '/' | '%' | '<<' | '>>' | '&' | '&^' ) expression  # mulExpr
    | expression op=( '+' | '-' | '|' | '^' ) expression    # addExpr
    | expression op=( '==' | '!=' | '<' | '<=' | '>' | '>=' ) expression # relExpr
    | expression '&&' expression                            # andExpr
    | expression '||' expression                            # orExpr
    ;

expressionList
    : expression ( ',' expression )*
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
    : '(' ( expressionList | ) ')'
    ;

selector
    : '.' IDENTIFIER
    ;

appendExpression
    : 'append' '(' expression ',' expression ')'
    ;

lengthExpression
    : 'len' '(' expression ')'
    ;

capExpression
    : 'cap' '(' expression ')'
    ;

statementList
    : statement*
    ;

block
    : '{' statementList '}'
    ;

statement
    : 'print' '(' ( expressionList | ) ')' ';'
    | 'println' '(' ( expressionList | ) ')' ';'
    | 'return' ( expression | ) ';'
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
    | expression op=( '+=' | '&=' | '-=' | '|=' | '*=' | '^=' | '<<=' | '>>=' | '&^=' | '%=' | '/=' ) expression
    ;

ifStatement
    : 'if' expression block ( 'else' ( ifStatement | block ) )?
    | 'if' simpleStatement ';' expression block ( 'else' ( ifStatement | block ) )?
    ;

loop
    : 'for' block
    | 'for' expression block
    | 'for' simpleStatement ';' ( expression )? ';' ( simpleStatement )? block
    ;

switch_stmt
    : 'switch' ( simpleStatement ';' )? ( expression )? '{' expressionCaseClauseList '}'
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

IDENTIFIER: [a-zA-Z_] [a-zA-Z0-9_]* ;

INTLITERAL: [0-9]+ ;

FLOATLITERAL: [0-9]+ '.' [0-9]* | '.' [0-9]+ ;

RUNELITERAL: '\'' ( ~['\\] | '\\' . ) '\'' ;

RAWSTRINGLITERAL: '`' .*? '`' ;

INTERPRETEDSTRINGLITERAL: '"' ( ~["\\] | '\\' . )* '"' ;

WS: [ \t\r\n]+ -> skip ;

COMMENT: '//' ~[\r\n]* -> skip ;
MULTILINE_COMMENT: '/*' .*? '*/' -> skip ;
