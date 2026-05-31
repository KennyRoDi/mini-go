// Code generated from AlphaCompiler.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // AlphaCompiler

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type AlphaCompilerParser struct {
	*antlr.BaseParser
}

var AlphaCompilerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func alphacompilerParserInit() {
	staticData := &AlphaCompilerParserStaticData
	staticData.LiteralNames = []string{
		"", "'package'", "';'", "'var'", "'('", "')'", "'='", "'type'", "'func'",
		"','", "'['", "']'", "'struct'", "'{'", "'}'", "'+'", "'-'", "'!'",
		"'^'", "'*'", "'/'", "'%'", "'<<'", "'>>'", "'&'", "'&^'", "'|'", "'=='",
		"'!='", "'<'", "'<='", "'>'", "'>='", "'&&'", "'||'", "'.'", "'append'",
		"'len'", "'cap'", "'print'", "'println'", "'return'", "'break'", "'continue'",
		"'++'", "'--'", "':='", "'+='", "'&='", "'-='", "'|='", "'*='", "'^='",
		"'<<='", "'>>='", "'&^='", "'%='", "'/='", "'if'", "'else'", "'for'",
		"'switch'", "':'", "'case'", "'default'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "IDENTIFIER",
		"INTLITERAL", "FLOATLITERAL", "RUNELITERAL", "RAWSTRINGLITERAL", "INTERPRETEDSTRINGLITERAL",
		"WS", "COMMENT", "MULTILINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"root", "topDeclarationList", "variableDecl", "innerVarDecls", "singleVarDecl",
		"singleVarDeclNoExps", "typeDecl", "innerTypeDecls", "singleTypeDecl",
		"funcDecl", "funcFrontDecl", "funcArgDecls", "declType", "sliceDeclType",
		"arrayDeclType", "structDeclType", "structMemDecls", "identifierList",
		"expression", "expressionList", "primaryExpression", "operand", "literal",
		"index", "arguments", "selector", "appendExpression", "lengthExpression",
		"capExpression", "statementList", "block", "statement", "simpleStatement",
		"assignmentStatement", "ifStatement", "loop", "switch_stmt", "expressionCaseClauseList",
		"expressionCaseClause", "expressionSwitchCase",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 73, 484, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 5, 1, 90, 8, 1, 10, 1, 12, 1, 93, 9, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 3, 2, 109, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 116, 8, 3, 10,
		3, 12, 3, 119, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 3, 4, 131, 8, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 150, 8,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 157, 8, 7, 10, 7, 12, 7, 160, 9,
		7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 3, 10, 174, 8, 10, 1, 10, 1, 10, 1, 10, 3, 10, 179, 8, 10, 1, 11,
		1, 11, 1, 11, 5, 11, 184, 8, 11, 10, 11, 12, 11, 187, 9, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 197, 8, 12, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 15, 3, 15, 212, 8, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 5, 16, 221, 8, 16, 10, 16, 12, 16, 224, 9, 16, 1, 17, 1, 17, 1,
		17, 5, 17, 229, 8, 17, 10, 17, 12, 17, 232, 9, 17, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 244, 8, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 261, 8, 18, 10, 18, 12, 18, 264, 9,
		18, 1, 19, 1, 19, 1, 19, 5, 19, 269, 8, 19, 10, 19, 12, 19, 272, 9, 19,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 279, 8, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 5, 20, 287, 8, 20, 10, 20, 12, 20, 290, 9, 20,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 298, 8, 21, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 3, 24, 309, 8, 24,
		1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 29, 5, 29, 334, 8, 29, 10, 29, 12, 29, 337, 9, 29, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 347, 8, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 355, 8, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 3, 31, 362, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 386, 8, 31, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 393, 8, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 3, 32, 400, 8, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 3, 33, 410, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 3, 34, 418, 8, 34, 3, 34, 420, 8, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 3, 34, 430, 8, 34, 3, 34, 432, 8, 34, 3, 34, 434,
		8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 3, 35, 446, 8, 35, 1, 35, 1, 35, 3, 35, 450, 8, 35, 1, 35, 1, 35, 3,
		35, 454, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 460, 8, 36, 1, 36, 3,
		36, 463, 8, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 5, 37, 470, 8, 37, 10,
		37, 12, 37, 473, 9, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39,
		3, 39, 482, 8, 39, 1, 39, 0, 2, 36, 40, 40, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
		52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 0, 5, 1, 0, 19,
		25, 3, 0, 15, 16, 18, 18, 26, 26, 1, 0, 27, 32, 1, 0, 66, 70, 1, 0, 47,
		57, 517, 0, 80, 1, 0, 0, 0, 2, 91, 1, 0, 0, 0, 4, 108, 1, 0, 0, 0, 6, 110,
		1, 0, 0, 0, 8, 130, 1, 0, 0, 0, 10, 132, 1, 0, 0, 0, 12, 149, 1, 0, 0,
		0, 14, 151, 1, 0, 0, 0, 16, 161, 1, 0, 0, 0, 18, 164, 1, 0, 0, 0, 20, 168,
		1, 0, 0, 0, 22, 180, 1, 0, 0, 0, 24, 196, 1, 0, 0, 0, 26, 198, 1, 0, 0,
		0, 28, 202, 1, 0, 0, 0, 30, 207, 1, 0, 0, 0, 32, 215, 1, 0, 0, 0, 34, 225,
		1, 0, 0, 0, 36, 243, 1, 0, 0, 0, 38, 265, 1, 0, 0, 0, 40, 278, 1, 0, 0,
		0, 42, 297, 1, 0, 0, 0, 44, 299, 1, 0, 0, 0, 46, 301, 1, 0, 0, 0, 48, 305,
		1, 0, 0, 0, 50, 312, 1, 0, 0, 0, 52, 315, 1, 0, 0, 0, 54, 322, 1, 0, 0,
		0, 56, 327, 1, 0, 0, 0, 58, 335, 1, 0, 0, 0, 60, 338, 1, 0, 0, 0, 62, 385,
		1, 0, 0, 0, 64, 399, 1, 0, 0, 0, 66, 409, 1, 0, 0, 0, 68, 433, 1, 0, 0,
		0, 70, 453, 1, 0, 0, 0, 72, 455, 1, 0, 0, 0, 74, 471, 1, 0, 0, 0, 76, 474,
		1, 0, 0, 0, 78, 481, 1, 0, 0, 0, 80, 81, 5, 1, 0, 0, 81, 82, 5, 65, 0,
		0, 82, 83, 5, 2, 0, 0, 83, 84, 3, 2, 1, 0, 84, 85, 5, 0, 0, 1, 85, 1, 1,
		0, 0, 0, 86, 90, 3, 4, 2, 0, 87, 90, 3, 12, 6, 0, 88, 90, 3, 18, 9, 0,
		89, 86, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 88, 1, 0, 0, 0, 90, 93, 1,
		0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 3, 1, 0, 0, 0, 93,
		91, 1, 0, 0, 0, 94, 95, 5, 3, 0, 0, 95, 96, 3, 8, 4, 0, 96, 97, 5, 2, 0,
		0, 97, 109, 1, 0, 0, 0, 98, 99, 5, 3, 0, 0, 99, 100, 5, 4, 0, 0, 100, 101,
		3, 6, 3, 0, 101, 102, 5, 5, 0, 0, 102, 103, 5, 2, 0, 0, 103, 109, 1, 0,
		0, 0, 104, 105, 5, 3, 0, 0, 105, 106, 5, 4, 0, 0, 106, 107, 5, 5, 0, 0,
		107, 109, 5, 2, 0, 0, 108, 94, 1, 0, 0, 0, 108, 98, 1, 0, 0, 0, 108, 104,
		1, 0, 0, 0, 109, 5, 1, 0, 0, 0, 110, 111, 3, 8, 4, 0, 111, 117, 5, 2, 0,
		0, 112, 113, 3, 8, 4, 0, 113, 114, 5, 2, 0, 0, 114, 116, 1, 0, 0, 0, 115,
		112, 1, 0, 0, 0, 116, 119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118,
		1, 0, 0, 0, 118, 7, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 121, 3, 34,
		17, 0, 121, 122, 3, 24, 12, 0, 122, 123, 5, 6, 0, 0, 123, 124, 3, 38, 19,
		0, 124, 131, 1, 0, 0, 0, 125, 126, 3, 34, 17, 0, 126, 127, 5, 6, 0, 0,
		127, 128, 3, 38, 19, 0, 128, 131, 1, 0, 0, 0, 129, 131, 3, 10, 5, 0, 130,
		120, 1, 0, 0, 0, 130, 125, 1, 0, 0, 0, 130, 129, 1, 0, 0, 0, 131, 9, 1,
		0, 0, 0, 132, 133, 3, 34, 17, 0, 133, 134, 3, 24, 12, 0, 134, 11, 1, 0,
		0, 0, 135, 136, 5, 7, 0, 0, 136, 137, 3, 16, 8, 0, 137, 138, 5, 2, 0, 0,
		138, 150, 1, 0, 0, 0, 139, 140, 5, 7, 0, 0, 140, 141, 5, 4, 0, 0, 141,
		142, 3, 14, 7, 0, 142, 143, 5, 5, 0, 0, 143, 144, 5, 2, 0, 0, 144, 150,
		1, 0, 0, 0, 145, 146, 5, 7, 0, 0, 146, 147, 5, 4, 0, 0, 147, 148, 5, 5,
		0, 0, 148, 150, 5, 2, 0, 0, 149, 135, 1, 0, 0, 0, 149, 139, 1, 0, 0, 0,
		149, 145, 1, 0, 0, 0, 150, 13, 1, 0, 0, 0, 151, 152, 3, 16, 8, 0, 152,
		158, 5, 2, 0, 0, 153, 154, 3, 16, 8, 0, 154, 155, 5, 2, 0, 0, 155, 157,
		1, 0, 0, 0, 156, 153, 1, 0, 0, 0, 157, 160, 1, 0, 0, 0, 158, 156, 1, 0,
		0, 0, 158, 159, 1, 0, 0, 0, 159, 15, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0,
		161, 162, 5, 65, 0, 0, 162, 163, 3, 24, 12, 0, 163, 17, 1, 0, 0, 0, 164,
		165, 3, 20, 10, 0, 165, 166, 3, 60, 30, 0, 166, 167, 5, 2, 0, 0, 167, 19,
		1, 0, 0, 0, 168, 169, 5, 8, 0, 0, 169, 170, 5, 65, 0, 0, 170, 173, 5, 4,
		0, 0, 171, 174, 3, 22, 11, 0, 172, 174, 1, 0, 0, 0, 173, 171, 1, 0, 0,
		0, 173, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 178, 5, 5, 0, 0, 176,
		179, 3, 24, 12, 0, 177, 179, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 178, 177,
		1, 0, 0, 0, 179, 21, 1, 0, 0, 0, 180, 185, 3, 10, 5, 0, 181, 182, 5, 9,
		0, 0, 182, 184, 3, 10, 5, 0, 183, 181, 1, 0, 0, 0, 184, 187, 1, 0, 0, 0,
		185, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 23, 1, 0, 0, 0, 187, 185,
		1, 0, 0, 0, 188, 189, 5, 4, 0, 0, 189, 190, 3, 24, 12, 0, 190, 191, 5,
		5, 0, 0, 191, 197, 1, 0, 0, 0, 192, 197, 5, 65, 0, 0, 193, 197, 3, 26,
		13, 0, 194, 197, 3, 28, 14, 0, 195, 197, 3, 30, 15, 0, 196, 188, 1, 0,
		0, 0, 196, 192, 1, 0, 0, 0, 196, 193, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0,
		196, 195, 1, 0, 0, 0, 197, 25, 1, 0, 0, 0, 198, 199, 5, 10, 0, 0, 199,
		200, 5, 11, 0, 0, 200, 201, 3, 24, 12, 0, 201, 27, 1, 0, 0, 0, 202, 203,
		5, 10, 0, 0, 203, 204, 5, 66, 0, 0, 204, 205, 5, 11, 0, 0, 205, 206, 3,
		24, 12, 0, 206, 29, 1, 0, 0, 0, 207, 208, 5, 12, 0, 0, 208, 211, 5, 13,
		0, 0, 209, 212, 3, 32, 16, 0, 210, 212, 1, 0, 0, 0, 211, 209, 1, 0, 0,
		0, 211, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 214, 5, 14, 0, 0, 214,
		31, 1, 0, 0, 0, 215, 216, 3, 10, 5, 0, 216, 222, 5, 2, 0, 0, 217, 218,
		3, 10, 5, 0, 218, 219, 5, 2, 0, 0, 219, 221, 1, 0, 0, 0, 220, 217, 1, 0,
		0, 0, 221, 224, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0,
		223, 33, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 225, 230, 5, 65, 0, 0, 226,
		227, 5, 9, 0, 0, 227, 229, 5, 65, 0, 0, 228, 226, 1, 0, 0, 0, 229, 232,
		1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 35, 1, 0,
		0, 0, 232, 230, 1, 0, 0, 0, 233, 234, 6, 18, -1, 0, 234, 244, 3, 40, 20,
		0, 235, 236, 5, 15, 0, 0, 236, 244, 3, 36, 18, 9, 237, 238, 5, 16, 0, 0,
		238, 244, 3, 36, 18, 8, 239, 240, 5, 17, 0, 0, 240, 244, 3, 36, 18, 7,
		241, 242, 5, 18, 0, 0, 242, 244, 3, 36, 18, 6, 243, 233, 1, 0, 0, 0, 243,
		235, 1, 0, 0, 0, 243, 237, 1, 0, 0, 0, 243, 239, 1, 0, 0, 0, 243, 241,
		1, 0, 0, 0, 244, 262, 1, 0, 0, 0, 245, 246, 10, 5, 0, 0, 246, 247, 7, 0,
		0, 0, 247, 261, 3, 36, 18, 6, 248, 249, 10, 4, 0, 0, 249, 250, 7, 1, 0,
		0, 250, 261, 3, 36, 18, 5, 251, 252, 10, 3, 0, 0, 252, 253, 7, 2, 0, 0,
		253, 261, 3, 36, 18, 4, 254, 255, 10, 2, 0, 0, 255, 256, 5, 33, 0, 0, 256,
		261, 3, 36, 18, 3, 257, 258, 10, 1, 0, 0, 258, 259, 5, 34, 0, 0, 259, 261,
		3, 36, 18, 2, 260, 245, 1, 0, 0, 0, 260, 248, 1, 0, 0, 0, 260, 251, 1,
		0, 0, 0, 260, 254, 1, 0, 0, 0, 260, 257, 1, 0, 0, 0, 261, 264, 1, 0, 0,
		0, 262, 260, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 37, 1, 0, 0, 0, 264,
		262, 1, 0, 0, 0, 265, 270, 3, 36, 18, 0, 266, 267, 5, 9, 0, 0, 267, 269,
		3, 36, 18, 0, 268, 266, 1, 0, 0, 0, 269, 272, 1, 0, 0, 0, 270, 268, 1,
		0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 39, 1, 0, 0, 0, 272, 270, 1, 0, 0,
		0, 273, 274, 6, 20, -1, 0, 274, 279, 3, 42, 21, 0, 275, 279, 3, 52, 26,
		0, 276, 279, 3, 54, 27, 0, 277, 279, 3, 56, 28, 0, 278, 273, 1, 0, 0, 0,
		278, 275, 1, 0, 0, 0, 278, 276, 1, 0, 0, 0, 278, 277, 1, 0, 0, 0, 279,
		288, 1, 0, 0, 0, 280, 281, 10, 6, 0, 0, 281, 287, 3, 50, 25, 0, 282, 283,
		10, 5, 0, 0, 283, 287, 3, 46, 23, 0, 284, 285, 10, 4, 0, 0, 285, 287, 3,
		48, 24, 0, 286, 280, 1, 0, 0, 0, 286, 282, 1, 0, 0, 0, 286, 284, 1, 0,
		0, 0, 287, 290, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0,
		289, 41, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 291, 298, 3, 44, 22, 0, 292,
		298, 5, 65, 0, 0, 293, 294, 5, 4, 0, 0, 294, 295, 3, 36, 18, 0, 295, 296,
		5, 5, 0, 0, 296, 298, 1, 0, 0, 0, 297, 291, 1, 0, 0, 0, 297, 292, 1, 0,
		0, 0, 297, 293, 1, 0, 0, 0, 298, 43, 1, 0, 0, 0, 299, 300, 7, 3, 0, 0,
		300, 45, 1, 0, 0, 0, 301, 302, 5, 10, 0, 0, 302, 303, 3, 36, 18, 0, 303,
		304, 5, 11, 0, 0, 304, 47, 1, 0, 0, 0, 305, 308, 5, 4, 0, 0, 306, 309,
		3, 38, 19, 0, 307, 309, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 308, 307, 1,
		0, 0, 0, 309, 310, 1, 0, 0, 0, 310, 311, 5, 5, 0, 0, 311, 49, 1, 0, 0,
		0, 312, 313, 5, 35, 0, 0, 313, 314, 5, 65, 0, 0, 314, 51, 1, 0, 0, 0, 315,
		316, 5, 36, 0, 0, 316, 317, 5, 4, 0, 0, 317, 318, 3, 36, 18, 0, 318, 319,
		5, 9, 0, 0, 319, 320, 3, 36, 18, 0, 320, 321, 5, 5, 0, 0, 321, 53, 1, 0,
		0, 0, 322, 323, 5, 37, 0, 0, 323, 324, 5, 4, 0, 0, 324, 325, 3, 36, 18,
		0, 325, 326, 5, 5, 0, 0, 326, 55, 1, 0, 0, 0, 327, 328, 5, 38, 0, 0, 328,
		329, 5, 4, 0, 0, 329, 330, 3, 36, 18, 0, 330, 331, 5, 5, 0, 0, 331, 57,
		1, 0, 0, 0, 332, 334, 3, 62, 31, 0, 333, 332, 1, 0, 0, 0, 334, 337, 1,
		0, 0, 0, 335, 333, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336, 59, 1, 0, 0,
		0, 337, 335, 1, 0, 0, 0, 338, 339, 5, 13, 0, 0, 339, 340, 3, 58, 29, 0,
		340, 341, 5, 14, 0, 0, 341, 61, 1, 0, 0, 0, 342, 343, 5, 39, 0, 0, 343,
		346, 5, 4, 0, 0, 344, 347, 3, 38, 19, 0, 345, 347, 1, 0, 0, 0, 346, 344,
		1, 0, 0, 0, 346, 345, 1, 0, 0, 0, 347, 348, 1, 0, 0, 0, 348, 349, 5, 5,
		0, 0, 349, 386, 5, 2, 0, 0, 350, 351, 5, 40, 0, 0, 351, 354, 5, 4, 0, 0,
		352, 355, 3, 38, 19, 0, 353, 355, 1, 0, 0, 0, 354, 352, 1, 0, 0, 0, 354,
		353, 1, 0, 0, 0, 355, 356, 1, 0, 0, 0, 356, 357, 5, 5, 0, 0, 357, 386,
		5, 2, 0, 0, 358, 361, 5, 41, 0, 0, 359, 362, 3, 36, 18, 0, 360, 362, 1,
		0, 0, 0, 361, 359, 1, 0, 0, 0, 361, 360, 1, 0, 0, 0, 362, 363, 1, 0, 0,
		0, 363, 386, 5, 2, 0, 0, 364, 365, 5, 42, 0, 0, 365, 386, 5, 2, 0, 0, 366,
		367, 5, 43, 0, 0, 367, 386, 5, 2, 0, 0, 368, 369, 3, 64, 32, 0, 369, 370,
		5, 2, 0, 0, 370, 386, 1, 0, 0, 0, 371, 372, 3, 60, 30, 0, 372, 373, 5,
		2, 0, 0, 373, 386, 1, 0, 0, 0, 374, 375, 3, 72, 36, 0, 375, 376, 5, 2,
		0, 0, 376, 386, 1, 0, 0, 0, 377, 378, 3, 68, 34, 0, 378, 379, 5, 2, 0,
		0, 379, 386, 1, 0, 0, 0, 380, 381, 3, 70, 35, 0, 381, 382, 5, 2, 0, 0,
		382, 386, 1, 0, 0, 0, 383, 386, 3, 12, 6, 0, 384, 386, 3, 4, 2, 0, 385,
		342, 1, 0, 0, 0, 385, 350, 1, 0, 0, 0, 385, 358, 1, 0, 0, 0, 385, 364,
		1, 0, 0, 0, 385, 366, 1, 0, 0, 0, 385, 368, 1, 0, 0, 0, 385, 371, 1, 0,
		0, 0, 385, 374, 1, 0, 0, 0, 385, 377, 1, 0, 0, 0, 385, 380, 1, 0, 0, 0,
		385, 383, 1, 0, 0, 0, 385, 384, 1, 0, 0, 0, 386, 63, 1, 0, 0, 0, 387, 400,
		1, 0, 0, 0, 388, 392, 3, 36, 18, 0, 389, 393, 5, 44, 0, 0, 390, 393, 5,
		45, 0, 0, 391, 393, 1, 0, 0, 0, 392, 389, 1, 0, 0, 0, 392, 390, 1, 0, 0,
		0, 392, 391, 1, 0, 0, 0, 393, 400, 1, 0, 0, 0, 394, 400, 3, 66, 33, 0,
		395, 396, 3, 38, 19, 0, 396, 397, 5, 46, 0, 0, 397, 398, 3, 38, 19, 0,
		398, 400, 1, 0, 0, 0, 399, 387, 1, 0, 0, 0, 399, 388, 1, 0, 0, 0, 399,
		394, 1, 0, 0, 0, 399, 395, 1, 0, 0, 0, 400, 65, 1, 0, 0, 0, 401, 402, 3,
		38, 19, 0, 402, 403, 5, 6, 0, 0, 403, 404, 3, 38, 19, 0, 404, 410, 1, 0,
		0, 0, 405, 406, 3, 36, 18, 0, 406, 407, 7, 4, 0, 0, 407, 408, 3, 36, 18,
		0, 408, 410, 1, 0, 0, 0, 409, 401, 1, 0, 0, 0, 409, 405, 1, 0, 0, 0, 410,
		67, 1, 0, 0, 0, 411, 412, 5, 58, 0, 0, 412, 413, 3, 36, 18, 0, 413, 419,
		3, 60, 30, 0, 414, 417, 5, 59, 0, 0, 415, 418, 3, 68, 34, 0, 416, 418,
		3, 60, 30, 0, 417, 415, 1, 0, 0, 0, 417, 416, 1, 0, 0, 0, 418, 420, 1,
		0, 0, 0, 419, 414, 1, 0, 0, 0, 419, 420, 1, 0, 0, 0, 420, 434, 1, 0, 0,
		0, 421, 422, 5, 58, 0, 0, 422, 423, 3, 64, 32, 0, 423, 424, 5, 2, 0, 0,
		424, 425, 3, 36, 18, 0, 425, 431, 3, 60, 30, 0, 426, 429, 5, 59, 0, 0,
		427, 430, 3, 68, 34, 0, 428, 430, 3, 60, 30, 0, 429, 427, 1, 0, 0, 0, 429,
		428, 1, 0, 0, 0, 430, 432, 1, 0, 0, 0, 431, 426, 1, 0, 0, 0, 431, 432,
		1, 0, 0, 0, 432, 434, 1, 0, 0, 0, 433, 411, 1, 0, 0, 0, 433, 421, 1, 0,
		0, 0, 434, 69, 1, 0, 0, 0, 435, 436, 5, 60, 0, 0, 436, 454, 3, 60, 30,
		0, 437, 438, 5, 60, 0, 0, 438, 439, 3, 36, 18, 0, 439, 440, 3, 60, 30,
		0, 440, 454, 1, 0, 0, 0, 441, 442, 5, 60, 0, 0, 442, 443, 3, 64, 32, 0,
		443, 445, 5, 2, 0, 0, 444, 446, 3, 36, 18, 0, 445, 444, 1, 0, 0, 0, 445,
		446, 1, 0, 0, 0, 446, 447, 1, 0, 0, 0, 447, 449, 5, 2, 0, 0, 448, 450,
		3, 64, 32, 0, 449, 448, 1, 0, 0, 0, 449, 450, 1, 0, 0, 0, 450, 451, 1,
		0, 0, 0, 451, 452, 3, 60, 30, 0, 452, 454, 1, 0, 0, 0, 453, 435, 1, 0,
		0, 0, 453, 437, 1, 0, 0, 0, 453, 441, 1, 0, 0, 0, 454, 71, 1, 0, 0, 0,
		455, 459, 5, 61, 0, 0, 456, 457, 3, 64, 32, 0, 457, 458, 5, 2, 0, 0, 458,
		460, 1, 0, 0, 0, 459, 456, 1, 0, 0, 0, 459, 460, 1, 0, 0, 0, 460, 462,
		1, 0, 0, 0, 461, 463, 3, 36, 18, 0, 462, 461, 1, 0, 0, 0, 462, 463, 1,
		0, 0, 0, 463, 464, 1, 0, 0, 0, 464, 465, 5, 13, 0, 0, 465, 466, 3, 74,
		37, 0, 466, 467, 5, 14, 0, 0, 467, 73, 1, 0, 0, 0, 468, 470, 3, 76, 38,
		0, 469, 468, 1, 0, 0, 0, 470, 473, 1, 0, 0, 0, 471, 469, 1, 0, 0, 0, 471,
		472, 1, 0, 0, 0, 472, 75, 1, 0, 0, 0, 473, 471, 1, 0, 0, 0, 474, 475, 3,
		78, 39, 0, 475, 476, 5, 62, 0, 0, 476, 477, 3, 58, 29, 0, 477, 77, 1, 0,
		0, 0, 478, 479, 5, 63, 0, 0, 479, 482, 3, 38, 19, 0, 480, 482, 5, 64, 0,
		0, 481, 478, 1, 0, 0, 0, 481, 480, 1, 0, 0, 0, 482, 79, 1, 0, 0, 0, 43,
		89, 91, 108, 117, 130, 149, 158, 173, 178, 185, 196, 211, 222, 230, 243,
		260, 262, 270, 278, 286, 288, 297, 308, 335, 346, 354, 361, 385, 392, 399,
		409, 417, 419, 429, 431, 433, 445, 449, 453, 459, 462, 471, 481,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// AlphaCompilerParserInit initializes any static state used to implement AlphaCompilerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAlphaCompilerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AlphaCompilerParserInit() {
	staticData := &AlphaCompilerParserStaticData
	staticData.once.Do(alphacompilerParserInit)
}

// NewAlphaCompilerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAlphaCompilerParser(input antlr.TokenStream) *AlphaCompilerParser {
	AlphaCompilerParserInit()
	this := new(AlphaCompilerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &AlphaCompilerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "AlphaCompiler.g4"

	return this
}

// AlphaCompilerParser tokens.
const (
	AlphaCompilerParserEOF                      = antlr.TokenEOF
	AlphaCompilerParserT__0                     = 1
	AlphaCompilerParserT__1                     = 2
	AlphaCompilerParserT__2                     = 3
	AlphaCompilerParserT__3                     = 4
	AlphaCompilerParserT__4                     = 5
	AlphaCompilerParserT__5                     = 6
	AlphaCompilerParserT__6                     = 7
	AlphaCompilerParserT__7                     = 8
	AlphaCompilerParserT__8                     = 9
	AlphaCompilerParserT__9                     = 10
	AlphaCompilerParserT__10                    = 11
	AlphaCompilerParserT__11                    = 12
	AlphaCompilerParserT__12                    = 13
	AlphaCompilerParserT__13                    = 14
	AlphaCompilerParserT__14                    = 15
	AlphaCompilerParserT__15                    = 16
	AlphaCompilerParserT__16                    = 17
	AlphaCompilerParserT__17                    = 18
	AlphaCompilerParserT__18                    = 19
	AlphaCompilerParserT__19                    = 20
	AlphaCompilerParserT__20                    = 21
	AlphaCompilerParserT__21                    = 22
	AlphaCompilerParserT__22                    = 23
	AlphaCompilerParserT__23                    = 24
	AlphaCompilerParserT__24                    = 25
	AlphaCompilerParserT__25                    = 26
	AlphaCompilerParserT__26                    = 27
	AlphaCompilerParserT__27                    = 28
	AlphaCompilerParserT__28                    = 29
	AlphaCompilerParserT__29                    = 30
	AlphaCompilerParserT__30                    = 31
	AlphaCompilerParserT__31                    = 32
	AlphaCompilerParserT__32                    = 33
	AlphaCompilerParserT__33                    = 34
	AlphaCompilerParserT__34                    = 35
	AlphaCompilerParserT__35                    = 36
	AlphaCompilerParserT__36                    = 37
	AlphaCompilerParserT__37                    = 38
	AlphaCompilerParserT__38                    = 39
	AlphaCompilerParserT__39                    = 40
	AlphaCompilerParserT__40                    = 41
	AlphaCompilerParserT__41                    = 42
	AlphaCompilerParserT__42                    = 43
	AlphaCompilerParserT__43                    = 44
	AlphaCompilerParserT__44                    = 45
	AlphaCompilerParserT__45                    = 46
	AlphaCompilerParserT__46                    = 47
	AlphaCompilerParserT__47                    = 48
	AlphaCompilerParserT__48                    = 49
	AlphaCompilerParserT__49                    = 50
	AlphaCompilerParserT__50                    = 51
	AlphaCompilerParserT__51                    = 52
	AlphaCompilerParserT__52                    = 53
	AlphaCompilerParserT__53                    = 54
	AlphaCompilerParserT__54                    = 55
	AlphaCompilerParserT__55                    = 56
	AlphaCompilerParserT__56                    = 57
	AlphaCompilerParserT__57                    = 58
	AlphaCompilerParserT__58                    = 59
	AlphaCompilerParserT__59                    = 60
	AlphaCompilerParserT__60                    = 61
	AlphaCompilerParserT__61                    = 62
	AlphaCompilerParserT__62                    = 63
	AlphaCompilerParserT__63                    = 64
	AlphaCompilerParserIDENTIFIER               = 65
	AlphaCompilerParserINTLITERAL               = 66
	AlphaCompilerParserFLOATLITERAL             = 67
	AlphaCompilerParserRUNELITERAL              = 68
	AlphaCompilerParserRAWSTRINGLITERAL         = 69
	AlphaCompilerParserINTERPRETEDSTRINGLITERAL = 70
	AlphaCompilerParserWS                       = 71
	AlphaCompilerParserCOMMENT                  = 72
	AlphaCompilerParserMULTILINE_COMMENT        = 73
)

// AlphaCompilerParser rules.
const (
	AlphaCompilerParserRULE_root                     = 0
	AlphaCompilerParserRULE_topDeclarationList       = 1
	AlphaCompilerParserRULE_variableDecl             = 2
	AlphaCompilerParserRULE_innerVarDecls            = 3
	AlphaCompilerParserRULE_singleVarDecl            = 4
	AlphaCompilerParserRULE_singleVarDeclNoExps      = 5
	AlphaCompilerParserRULE_typeDecl                 = 6
	AlphaCompilerParserRULE_innerTypeDecls           = 7
	AlphaCompilerParserRULE_singleTypeDecl           = 8
	AlphaCompilerParserRULE_funcDecl                 = 9
	AlphaCompilerParserRULE_funcFrontDecl            = 10
	AlphaCompilerParserRULE_funcArgDecls             = 11
	AlphaCompilerParserRULE_declType                 = 12
	AlphaCompilerParserRULE_sliceDeclType            = 13
	AlphaCompilerParserRULE_arrayDeclType            = 14
	AlphaCompilerParserRULE_structDeclType           = 15
	AlphaCompilerParserRULE_structMemDecls           = 16
	AlphaCompilerParserRULE_identifierList           = 17
	AlphaCompilerParserRULE_expression               = 18
	AlphaCompilerParserRULE_expressionList           = 19
	AlphaCompilerParserRULE_primaryExpression        = 20
	AlphaCompilerParserRULE_operand                  = 21
	AlphaCompilerParserRULE_literal                  = 22
	AlphaCompilerParserRULE_index                    = 23
	AlphaCompilerParserRULE_arguments                = 24
	AlphaCompilerParserRULE_selector                 = 25
	AlphaCompilerParserRULE_appendExpression         = 26
	AlphaCompilerParserRULE_lengthExpression         = 27
	AlphaCompilerParserRULE_capExpression            = 28
	AlphaCompilerParserRULE_statementList            = 29
	AlphaCompilerParserRULE_block                    = 30
	AlphaCompilerParserRULE_statement                = 31
	AlphaCompilerParserRULE_simpleStatement          = 32
	AlphaCompilerParserRULE_assignmentStatement      = 33
	AlphaCompilerParserRULE_ifStatement              = 34
	AlphaCompilerParserRULE_loop                     = 35
	AlphaCompilerParserRULE_switch_stmt              = 36
	AlphaCompilerParserRULE_expressionCaseClauseList = 37
	AlphaCompilerParserRULE_expressionCaseClause     = 38
	AlphaCompilerParserRULE_expressionSwitchCase     = 39
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	TopDeclarationList() ITopDeclarationListContext
	EOF() antlr.TerminalNode

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_root
	return p
}

func InitEmptyRootContext(p *RootContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_root
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AlphaCompilerParserIDENTIFIER, 0)
}

func (s *RootContext) TopDeclarationList() ITopDeclarationListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopDeclarationListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopDeclarationListContext)
}

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(AlphaCompilerParserEOF, 0)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (s *RootContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitRoot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AlphaCompilerParserRULE_root)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(AlphaCompilerParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Match(AlphaCompilerParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Match(AlphaCompilerParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.TopDeclarationList()
	}
	{
		p.SetState(84)
		p.Match(AlphaCompilerParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITopDeclarationListContext is an interface to support dynamic dispatch.
type ITopDeclarationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVariableDecl() []IVariableDeclContext
	VariableDecl(i int) IVariableDeclContext
	AllTypeDecl() []ITypeDeclContext
	TypeDecl(i int) ITypeDeclContext
	AllFuncDecl() []IFuncDeclContext
	FuncDecl(i int) IFuncDeclContext

	// IsTopDeclarationListContext differentiates from other interfaces.
	IsTopDeclarationListContext()
}

type TopDeclarationListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopDeclarationListContext() *TopDeclarationListContext {
	var p = new(TopDeclarationListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_topDeclarationList
	return p
}

func InitEmptyTopDeclarationListContext(p *TopDeclarationListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_topDeclarationList
}

func (*TopDeclarationListContext) IsTopDeclarationListContext() {}

func NewTopDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopDeclarationListContext {
	var p = new(TopDeclarationListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_topDeclarationList

	return p
}

func (s *TopDeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *TopDeclarationListContext) AllVariableDecl() []IVariableDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableDeclContext); ok {
			len++
		}
	}

	tst := make([]IVariableDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableDeclContext); ok {
			tst[i] = t.(IVariableDeclContext)
			i++
		}
	}

	return tst
}

func (s *TopDeclarationListContext) VariableDecl(i int) IVariableDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclContext)
}

func (s *TopDeclarationListContext) AllTypeDecl() []ITypeDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeDeclContext); ok {
			len++
		}
	}

	tst := make([]ITypeDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeDeclContext); ok {
			tst[i] = t.(ITypeDeclContext)
			i++
		}
	}

	return tst
}

func (s *TopDeclarationListContext) TypeDecl(i int) ITypeDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *TopDeclarationListContext) AllFuncDecl() []IFuncDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncDeclContext); ok {
			len++
		}
	}

	tst := make([]IFuncDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncDeclContext); ok {
			tst[i] = t.(IFuncDeclContext)
			i++
		}
	}

	return tst
}

func (s *TopDeclarationListContext) FuncDecl(i int) IFuncDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDeclContext)
}

func (s *TopDeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopDeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopDeclarationListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterTopDeclarationList(s)
	}
}

func (s *TopDeclarationListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitTopDeclarationList(s)
	}
}

func (s *TopDeclarationListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitTopDeclarationList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) TopDeclarationList() (localctx ITopDeclarationListContext) {
	localctx = NewTopDeclarationListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AlphaCompilerParserRULE_topDeclarationList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&392) != 0 {
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case AlphaCompilerParserT__2:
			{
				p.SetState(86)
				p.VariableDecl()
			}

		case AlphaCompilerParserT__6:
			{
				p.SetState(87)
				p.TypeDecl()
			}

		case AlphaCompilerParserT__7:
			{
				p.SetState(88)
				p.FuncDecl()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableDeclContext is an interface to support dynamic dispatch.
type IVariableDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SingleVarDecl() ISingleVarDeclContext
	InnerVarDecls() IInnerVarDeclsContext

	// IsVariableDeclContext differentiates from other interfaces.
	IsVariableDeclContext()
}

type VariableDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclContext() *VariableDeclContext {
	var p = new(VariableDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_variableDecl
	return p
}

func InitEmptyVariableDeclContext(p *VariableDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_variableDecl
}

func (*VariableDeclContext) IsVariableDeclContext() {}

func NewVariableDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclContext {
	var p = new(VariableDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_variableDecl

	return p
}

func (s *VariableDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclContext) SingleVarDecl() ISingleVarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclContext)
}

func (s *VariableDeclContext) InnerVarDecls() IInnerVarDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInnerVarDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInnerVarDeclsContext)
}

func (s *VariableDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterVariableDecl(s)
	}
}

func (s *VariableDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitVariableDecl(s)
	}
}

func (s *VariableDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitVariableDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) VariableDecl() (localctx IVariableDeclContext) {
	localctx = NewVariableDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AlphaCompilerParserRULE_variableDecl)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.Match(AlphaCompilerParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.SingleVarDecl()
		}
		{
			p.SetState(96)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(98)
			p.Match(AlphaCompilerParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Match(AlphaCompilerParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.InnerVarDecls()
		}
		{
			p.SetState(101)
			p.Match(AlphaCompilerParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(104)
			p.Match(AlphaCompilerParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Match(AlphaCompilerParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(AlphaCompilerParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInnerVarDeclsContext is an interface to support dynamic dispatch.
type IInnerVarDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSingleVarDecl() []ISingleVarDeclContext
	SingleVarDecl(i int) ISingleVarDeclContext

	// IsInnerVarDeclsContext differentiates from other interfaces.
	IsInnerVarDeclsContext()
}

type InnerVarDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerVarDeclsContext() *InnerVarDeclsContext {
	var p = new(InnerVarDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_innerVarDecls
	return p
}

func InitEmptyInnerVarDeclsContext(p *InnerVarDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_innerVarDecls
}

func (*InnerVarDeclsContext) IsInnerVarDeclsContext() {}

func NewInnerVarDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerVarDeclsContext {
	var p = new(InnerVarDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_innerVarDecls

	return p
}

func (s *InnerVarDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerVarDeclsContext) AllSingleVarDecl() []ISingleVarDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleVarDeclContext); ok {
			len++
		}
	}

	tst := make([]ISingleVarDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleVarDeclContext); ok {
			tst[i] = t.(ISingleVarDeclContext)
			i++
		}
	}

	return tst
}

func (s *InnerVarDeclsContext) SingleVarDecl(i int) ISingleVarDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclContext)
}

func (s *InnerVarDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerVarDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InnerVarDeclsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterInnerVarDecls(s)
	}
}

func (s *InnerVarDeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitInnerVarDecls(s)
	}
}

func (s *InnerVarDeclsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitInnerVarDecls(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) InnerVarDecls() (localctx IInnerVarDeclsContext) {
	localctx = NewInnerVarDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AlphaCompilerParserRULE_innerVarDecls)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.SingleVarDecl()
	}
	{
		p.SetState(111)
		p.Match(AlphaCompilerParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == AlphaCompilerParserIDENTIFIER {
		{
			p.SetState(112)
			p.SingleVarDecl()
		}
		{
			p.SetState(113)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingleVarDeclContext is an interface to support dynamic dispatch.
type ISingleVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentifierList() IIdentifierListContext
	DeclType() IDeclTypeContext
	ExpressionList() IExpressionListContext
	SingleVarDeclNoExps() ISingleVarDeclNoExpsContext

	// IsSingleVarDeclContext differentiates from other interfaces.
	IsSingleVarDeclContext()
}

type SingleVarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleVarDeclContext() *SingleVarDeclContext {
	var p = new(SingleVarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_singleVarDecl
	return p
}

func InitEmptySingleVarDeclContext(p *SingleVarDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_singleVarDecl
}

func (*SingleVarDeclContext) IsSingleVarDeclContext() {}

func NewSingleVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleVarDeclContext {
	var p = new(SingleVarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_singleVarDecl

	return p
}

func (s *SingleVarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleVarDeclContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *SingleVarDeclContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *SingleVarDeclContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *SingleVarDeclContext) SingleVarDeclNoExps() ISingleVarDeclNoExpsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclNoExpsContext)
}

func (s *SingleVarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleVarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterSingleVarDecl(s)
	}
}

func (s *SingleVarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitSingleVarDecl(s)
	}
}

func (s *SingleVarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitSingleVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) SingleVarDecl() (localctx ISingleVarDeclContext) {
	localctx = NewSingleVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AlphaCompilerParserRULE_singleVarDecl)
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.IdentifierList()
		}
		{
			p.SetState(121)
			p.DeclType()
		}
		{
			p.SetState(122)
			p.Match(AlphaCompilerParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.ExpressionList()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.IdentifierList()
		}
		{
			p.SetState(126)
			p.Match(AlphaCompilerParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.ExpressionList()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(129)
			p.SingleVarDeclNoExps()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingleVarDeclNoExpsContext is an interface to support dynamic dispatch.
type ISingleVarDeclNoExpsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentifierList() IIdentifierListContext
	DeclType() IDeclTypeContext

	// IsSingleVarDeclNoExpsContext differentiates from other interfaces.
	IsSingleVarDeclNoExpsContext()
}

type SingleVarDeclNoExpsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleVarDeclNoExpsContext() *SingleVarDeclNoExpsContext {
	var p = new(SingleVarDeclNoExpsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_singleVarDeclNoExps
	return p
}

func InitEmptySingleVarDeclNoExpsContext(p *SingleVarDeclNoExpsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_singleVarDeclNoExps
}

func (*SingleVarDeclNoExpsContext) IsSingleVarDeclNoExpsContext() {}

func NewSingleVarDeclNoExpsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleVarDeclNoExpsContext {
	var p = new(SingleVarDeclNoExpsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_singleVarDeclNoExps

	return p
}

func (s *SingleVarDeclNoExpsContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleVarDeclNoExpsContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *SingleVarDeclNoExpsContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *SingleVarDeclNoExpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclNoExpsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleVarDeclNoExpsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterSingleVarDeclNoExps(s)
	}
}

func (s *SingleVarDeclNoExpsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitSingleVarDeclNoExps(s)
	}
}

func (s *SingleVarDeclNoExpsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitSingleVarDeclNoExps(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) SingleVarDeclNoExps() (localctx ISingleVarDeclNoExpsContext) {
	localctx = NewSingleVarDeclNoExpsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AlphaCompilerParserRULE_singleVarDeclNoExps)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.IdentifierList()
	}
	{
		p.SetState(133)
		p.DeclType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeDeclContext is an interface to support dynamic dispatch.
type ITypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SingleTypeDecl() ISingleTypeDeclContext
	InnerTypeDecls() IInnerTypeDeclsContext

	// IsTypeDeclContext differentiates from other interfaces.
	IsTypeDeclContext()
}

type TypeDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclContext() *TypeDeclContext {
	var p = new(TypeDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_typeDecl
	return p
}

func InitEmptyTypeDeclContext(p *TypeDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_typeDecl
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_typeDecl

	return p
}

func (s *TypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclContext) SingleTypeDecl() ISingleTypeDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleTypeDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleTypeDeclContext)
}

func (s *TypeDeclContext) InnerTypeDecls() IInnerTypeDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInnerTypeDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInnerTypeDeclsContext)
}

func (s *TypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterTypeDecl(s)
	}
}

func (s *TypeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitTypeDecl(s)
	}
}

func (s *TypeDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitTypeDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) TypeDecl() (localctx ITypeDeclContext) {
	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AlphaCompilerParserRULE_typeDecl)
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.Match(AlphaCompilerParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.SingleTypeDecl()
		}
		{
			p.SetState(137)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(139)
			p.Match(AlphaCompilerParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.Match(AlphaCompilerParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.InnerTypeDecls()
		}
		{
			p.SetState(142)
			p.Match(AlphaCompilerParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(145)
			p.Match(AlphaCompilerParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Match(AlphaCompilerParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Match(AlphaCompilerParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(148)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInnerTypeDeclsContext is an interface to support dynamic dispatch.
type IInnerTypeDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSingleTypeDecl() []ISingleTypeDeclContext
	SingleTypeDecl(i int) ISingleTypeDeclContext

	// IsInnerTypeDeclsContext differentiates from other interfaces.
	IsInnerTypeDeclsContext()
}

type InnerTypeDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerTypeDeclsContext() *InnerTypeDeclsContext {
	var p = new(InnerTypeDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_innerTypeDecls
	return p
}

func InitEmptyInnerTypeDeclsContext(p *InnerTypeDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_innerTypeDecls
}

func (*InnerTypeDeclsContext) IsInnerTypeDeclsContext() {}

func NewInnerTypeDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerTypeDeclsContext {
	var p = new(InnerTypeDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_innerTypeDecls

	return p
}

func (s *InnerTypeDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerTypeDeclsContext) AllSingleTypeDecl() []ISingleTypeDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleTypeDeclContext); ok {
			len++
		}
	}

	tst := make([]ISingleTypeDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleTypeDeclContext); ok {
			tst[i] = t.(ISingleTypeDeclContext)
			i++
		}
	}

	return tst
}

func (s *InnerTypeDeclsContext) SingleTypeDecl(i int) ISingleTypeDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleTypeDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleTypeDeclContext)
}

func (s *InnerTypeDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerTypeDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InnerTypeDeclsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterInnerTypeDecls(s)
	}
}

func (s *InnerTypeDeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitInnerTypeDecls(s)
	}
}

func (s *InnerTypeDeclsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitInnerTypeDecls(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) InnerTypeDecls() (localctx IInnerTypeDeclsContext) {
	localctx = NewInnerTypeDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AlphaCompilerParserRULE_innerTypeDecls)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.SingleTypeDecl()
	}
	{
		p.SetState(152)
		p.Match(AlphaCompilerParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == AlphaCompilerParserIDENTIFIER {
		{
			p.SetState(153)
			p.SingleTypeDecl()
		}
		{
			p.SetState(154)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingleTypeDeclContext is an interface to support dynamic dispatch.
type ISingleTypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	DeclType() IDeclTypeContext

	// IsSingleTypeDeclContext differentiates from other interfaces.
	IsSingleTypeDeclContext()
}

type SingleTypeDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleTypeDeclContext() *SingleTypeDeclContext {
	var p = new(SingleTypeDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_singleTypeDecl
	return p
}

func InitEmptySingleTypeDeclContext(p *SingleTypeDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_singleTypeDecl
}

func (*SingleTypeDeclContext) IsSingleTypeDeclContext() {}

func NewSingleTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleTypeDeclContext {
	var p = new(SingleTypeDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_singleTypeDecl

	return p
}

func (s *SingleTypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleTypeDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AlphaCompilerParserIDENTIFIER, 0)
}

func (s *SingleTypeDeclContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *SingleTypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleTypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleTypeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterSingleTypeDecl(s)
	}
}

func (s *SingleTypeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitSingleTypeDecl(s)
	}
}

func (s *SingleTypeDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitSingleTypeDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) SingleTypeDecl() (localctx ISingleTypeDeclContext) {
	localctx = NewSingleTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AlphaCompilerParserRULE_singleTypeDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(AlphaCompilerParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.DeclType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncDeclContext is an interface to support dynamic dispatch.
type IFuncDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FuncFrontDecl() IFuncFrontDeclContext
	Block() IBlockContext

	// IsFuncDeclContext differentiates from other interfaces.
	IsFuncDeclContext()
}

type FuncDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclContext() *FuncDeclContext {
	var p = new(FuncDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_funcDecl
	return p
}

func InitEmptyFuncDeclContext(p *FuncDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_funcDecl
}

func (*FuncDeclContext) IsFuncDeclContext() {}

func NewFuncDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclContext {
	var p = new(FuncDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_funcDecl

	return p
}

func (s *FuncDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclContext) FuncFrontDecl() IFuncFrontDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncFrontDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncFrontDeclContext)
}

func (s *FuncDeclContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterFuncDecl(s)
	}
}

func (s *FuncDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitFuncDecl(s)
	}
}

func (s *FuncDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitFuncDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) FuncDecl() (localctx IFuncDeclContext) {
	localctx = NewFuncDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AlphaCompilerParserRULE_funcDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.FuncFrontDecl()
	}
	{
		p.SetState(165)
		p.Block()
	}
	{
		p.SetState(166)
		p.Match(AlphaCompilerParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncFrontDeclContext is an interface to support dynamic dispatch.
type IFuncFrontDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	FuncArgDecls() IFuncArgDeclsContext
	DeclType() IDeclTypeContext

	// IsFuncFrontDeclContext differentiates from other interfaces.
	IsFuncFrontDeclContext()
}

type FuncFrontDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncFrontDeclContext() *FuncFrontDeclContext {
	var p = new(FuncFrontDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_funcFrontDecl
	return p
}

func InitEmptyFuncFrontDeclContext(p *FuncFrontDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_funcFrontDecl
}

func (*FuncFrontDeclContext) IsFuncFrontDeclContext() {}

func NewFuncFrontDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncFrontDeclContext {
	var p = new(FuncFrontDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_funcFrontDecl

	return p
}

func (s *FuncFrontDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncFrontDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AlphaCompilerParserIDENTIFIER, 0)
}

func (s *FuncFrontDeclContext) FuncArgDecls() IFuncArgDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncArgDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncArgDeclsContext)
}

func (s *FuncFrontDeclContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *FuncFrontDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncFrontDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncFrontDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterFuncFrontDecl(s)
	}
}

func (s *FuncFrontDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitFuncFrontDecl(s)
	}
}

func (s *FuncFrontDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitFuncFrontDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) FuncFrontDecl() (localctx IFuncFrontDeclContext) {
	localctx = NewFuncFrontDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AlphaCompilerParserRULE_funcFrontDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(AlphaCompilerParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(AlphaCompilerParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(AlphaCompilerParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AlphaCompilerParserIDENTIFIER:
		{
			p.SetState(171)
			p.FuncArgDecls()
		}

	case AlphaCompilerParserT__4:

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(175)
		p.Match(AlphaCompilerParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AlphaCompilerParserT__3, AlphaCompilerParserT__9, AlphaCompilerParserT__11, AlphaCompilerParserIDENTIFIER:
		{
			p.SetState(176)
			p.DeclType()
		}

	case AlphaCompilerParserT__12:

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncArgDeclsContext is an interface to support dynamic dispatch.
type IFuncArgDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext
	SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext

	// IsFuncArgDeclsContext differentiates from other interfaces.
	IsFuncArgDeclsContext()
}

type FuncArgDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncArgDeclsContext() *FuncArgDeclsContext {
	var p = new(FuncArgDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_funcArgDecls
	return p
}

func InitEmptyFuncArgDeclsContext(p *FuncArgDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_funcArgDecls
}

func (*FuncArgDeclsContext) IsFuncArgDeclsContext() {}

func NewFuncArgDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncArgDeclsContext {
	var p = new(FuncArgDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_funcArgDecls

	return p
}

func (s *FuncArgDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncArgDeclsContext) AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			len++
		}
	}

	tst := make([]ISingleVarDeclNoExpsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			tst[i] = t.(ISingleVarDeclNoExpsContext)
			i++
		}
	}

	return tst
}

func (s *FuncArgDeclsContext) SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclNoExpsContext)
}

func (s *FuncArgDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncArgDeclsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterFuncArgDecls(s)
	}
}

func (s *FuncArgDeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitFuncArgDecls(s)
	}
}

func (s *FuncArgDeclsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitFuncArgDecls(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) FuncArgDecls() (localctx IFuncArgDeclsContext) {
	localctx = NewFuncArgDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AlphaCompilerParserRULE_funcArgDecls)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.SingleVarDeclNoExps()
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == AlphaCompilerParserT__8 {
		{
			p.SetState(181)
			p.Match(AlphaCompilerParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.SingleVarDeclNoExps()
		}

		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclTypeContext is an interface to support dynamic dispatch.
type IDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DeclType() IDeclTypeContext
	IDENTIFIER() antlr.TerminalNode
	SliceDeclType() ISliceDeclTypeContext
	ArrayDeclType() IArrayDeclTypeContext
	StructDeclType() IStructDeclTypeContext

	// IsDeclTypeContext differentiates from other interfaces.
	IsDeclTypeContext()
}

type DeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclTypeContext() *DeclTypeContext {
	var p = new(DeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_declType
	return p
}

func InitEmptyDeclTypeContext(p *DeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_declType
}

func (*DeclTypeContext) IsDeclTypeContext() {}

func NewDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclTypeContext {
	var p = new(DeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_declType

	return p
}

func (s *DeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclTypeContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *DeclTypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AlphaCompilerParserIDENTIFIER, 0)
}

func (s *DeclTypeContext) SliceDeclType() ISliceDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceDeclTypeContext)
}

func (s *DeclTypeContext) ArrayDeclType() IArrayDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayDeclTypeContext)
}

func (s *DeclTypeContext) StructDeclType() IStructDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclTypeContext)
}

func (s *DeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterDeclType(s)
	}
}

func (s *DeclTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitDeclType(s)
	}
}

func (s *DeclTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitDeclType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) DeclType() (localctx IDeclTypeContext) {
	localctx = NewDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AlphaCompilerParserRULE_declType)
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.Match(AlphaCompilerParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.DeclType()
		}
		{
			p.SetState(190)
			p.Match(AlphaCompilerParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(192)
			p.Match(AlphaCompilerParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(193)
			p.SliceDeclType()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(194)
			p.ArrayDeclType()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(195)
			p.StructDeclType()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISliceDeclTypeContext is an interface to support dynamic dispatch.
type ISliceDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DeclType() IDeclTypeContext

	// IsSliceDeclTypeContext differentiates from other interfaces.
	IsSliceDeclTypeContext()
}

type SliceDeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceDeclTypeContext() *SliceDeclTypeContext {
	var p = new(SliceDeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_sliceDeclType
	return p
}

func InitEmptySliceDeclTypeContext(p *SliceDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_sliceDeclType
}

func (*SliceDeclTypeContext) IsSliceDeclTypeContext() {}

func NewSliceDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceDeclTypeContext {
	var p = new(SliceDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_sliceDeclType

	return p
}

func (s *SliceDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceDeclTypeContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *SliceDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceDeclTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterSliceDeclType(s)
	}
}

func (s *SliceDeclTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitSliceDeclType(s)
	}
}

func (s *SliceDeclTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitSliceDeclType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) SliceDeclType() (localctx ISliceDeclTypeContext) {
	localctx = NewSliceDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AlphaCompilerParserRULE_sliceDeclType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(AlphaCompilerParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.Match(AlphaCompilerParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
		p.DeclType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayDeclTypeContext is an interface to support dynamic dispatch.
type IArrayDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTLITERAL() antlr.TerminalNode
	DeclType() IDeclTypeContext

	// IsArrayDeclTypeContext differentiates from other interfaces.
	IsArrayDeclTypeContext()
}

type ArrayDeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayDeclTypeContext() *ArrayDeclTypeContext {
	var p = new(ArrayDeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_arrayDeclType
	return p
}

func InitEmptyArrayDeclTypeContext(p *ArrayDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_arrayDeclType
}

func (*ArrayDeclTypeContext) IsArrayDeclTypeContext() {}

func NewArrayDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayDeclTypeContext {
	var p = new(ArrayDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_arrayDeclType

	return p
}

func (s *ArrayDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayDeclTypeContext) INTLITERAL() antlr.TerminalNode {
	return s.GetToken(AlphaCompilerParserINTLITERAL, 0)
}

func (s *ArrayDeclTypeContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *ArrayDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayDeclTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterArrayDeclType(s)
	}
}

func (s *ArrayDeclTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitArrayDeclType(s)
	}
}

func (s *ArrayDeclTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitArrayDeclType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) ArrayDeclType() (localctx IArrayDeclTypeContext) {
	localctx = NewArrayDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, AlphaCompilerParserRULE_arrayDeclType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Match(AlphaCompilerParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
		p.Match(AlphaCompilerParserINTLITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.Match(AlphaCompilerParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.DeclType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructDeclTypeContext is an interface to support dynamic dispatch.
type IStructDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StructMemDecls() IStructMemDeclsContext

	// IsStructDeclTypeContext differentiates from other interfaces.
	IsStructDeclTypeContext()
}

type StructDeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclTypeContext() *StructDeclTypeContext {
	var p = new(StructDeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_structDeclType
	return p
}

func InitEmptyStructDeclTypeContext(p *StructDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_structDeclType
}

func (*StructDeclTypeContext) IsStructDeclTypeContext() {}

func NewStructDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclTypeContext {
	var p = new(StructDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_structDeclType

	return p
}

func (s *StructDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclTypeContext) StructMemDecls() IStructMemDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructMemDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructMemDeclsContext)
}

func (s *StructDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterStructDeclType(s)
	}
}

func (s *StructDeclTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitStructDeclType(s)
	}
}

func (s *StructDeclTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitStructDeclType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) StructDeclType() (localctx IStructDeclTypeContext) {
	localctx = NewStructDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, AlphaCompilerParserRULE_structDeclType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Match(AlphaCompilerParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.Match(AlphaCompilerParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AlphaCompilerParserIDENTIFIER:
		{
			p.SetState(209)
			p.StructMemDecls()
		}

	case AlphaCompilerParserT__13:

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(213)
		p.Match(AlphaCompilerParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructMemDeclsContext is an interface to support dynamic dispatch.
type IStructMemDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext
	SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext

	// IsStructMemDeclsContext differentiates from other interfaces.
	IsStructMemDeclsContext()
}

type StructMemDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructMemDeclsContext() *StructMemDeclsContext {
	var p = new(StructMemDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_structMemDecls
	return p
}

func InitEmptyStructMemDeclsContext(p *StructMemDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_structMemDecls
}

func (*StructMemDeclsContext) IsStructMemDeclsContext() {}

func NewStructMemDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructMemDeclsContext {
	var p = new(StructMemDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_structMemDecls

	return p
}

func (s *StructMemDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *StructMemDeclsContext) AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			len++
		}
	}

	tst := make([]ISingleVarDeclNoExpsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			tst[i] = t.(ISingleVarDeclNoExpsContext)
			i++
		}
	}

	return tst
}

func (s *StructMemDeclsContext) SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclNoExpsContext)
}

func (s *StructMemDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructMemDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructMemDeclsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterStructMemDecls(s)
	}
}

func (s *StructMemDeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitStructMemDecls(s)
	}
}

func (s *StructMemDeclsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitStructMemDecls(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) StructMemDecls() (localctx IStructMemDeclsContext) {
	localctx = NewStructMemDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, AlphaCompilerParserRULE_structMemDecls)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.SingleVarDeclNoExps()
	}
	{
		p.SetState(216)
		p.Match(AlphaCompilerParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == AlphaCompilerParserIDENTIFIER {
		{
			p.SetState(217)
			p.SingleVarDeclNoExps()
		}
		{
			p.SetState(218)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierListContext is an interface to support dynamic dispatch.
type IIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode

	// IsIdentifierListContext differentiates from other interfaces.
	IsIdentifierListContext()
}

type IdentifierListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierListContext() *IdentifierListContext {
	var p = new(IdentifierListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_identifierList
	return p
}

func InitEmptyIdentifierListContext(p *IdentifierListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_identifierList
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(AlphaCompilerParserIDENTIFIER)
}

func (s *IdentifierListContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(AlphaCompilerParserIDENTIFIER, i)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterIdentifierList(s)
	}
}

func (s *IdentifierListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitIdentifierList(s)
	}
}

func (s *IdentifierListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitIdentifierList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) IdentifierList() (localctx IIdentifierListContext) {
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, AlphaCompilerParserRULE_identifierList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.Match(AlphaCompilerParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == AlphaCompilerParserT__8 {
		{
			p.SetState(226)
			p.Match(AlphaCompilerParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)
			p.Match(AlphaCompilerParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UnaryExprContext struct {
	ExpressionContext
}

func NewUnaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExprContext {
	var p = new(UnaryExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterUnaryExpr(s)
	}
}

func (s *UnaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitUnaryExpr(s)
	}
}

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExprContext struct {
	ExpressionContext
}

func NewPrimaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddExprContext struct {
	ExpressionContext
	op antlr.Token
}

func NewAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExprContext {
	var p = new(AddExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AddExprContext) GetOp() antlr.Token { return s.op }

func (s *AddExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterAddExpr(s)
	}
}

func (s *AddExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitAddExpr(s)
	}
}

func (s *AddExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitAddExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulExprContext struct {
	ExpressionContext
	op antlr.Token
}

func NewMulExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulExprContext {
	var p = new(MulExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MulExprContext) GetOp() antlr.Token { return s.op }

func (s *MulExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MulExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterMulExpr(s)
	}
}

func (s *MulExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitMulExpr(s)
	}
}

func (s *MulExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitMulExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExprContext struct {
	ExpressionContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type RelExprContext struct {
	ExpressionContext
	op antlr.Token
}

func NewRelExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelExprContext {
	var p = new(RelExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RelExprContext) GetOp() antlr.Token { return s.op }

func (s *RelExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RelExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RelExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterRelExpr(s)
	}
}

func (s *RelExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitRelExpr(s)
	}
}

func (s *RelExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitRelExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndExprContext struct {
	ExpressionContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *AlphaCompilerParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, AlphaCompilerParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AlphaCompilerParserT__3, AlphaCompilerParserT__35, AlphaCompilerParserT__36, AlphaCompilerParserT__37, AlphaCompilerParserIDENTIFIER, AlphaCompilerParserINTLITERAL, AlphaCompilerParserFLOATLITERAL, AlphaCompilerParserRUNELITERAL, AlphaCompilerParserRAWSTRINGLITERAL, AlphaCompilerParserINTERPRETEDSTRINGLITERAL:
		localctx = NewPrimaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(234)
			p.primaryExpression(0)
		}

	case AlphaCompilerParserT__14:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(235)
			p.Match(AlphaCompilerParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)
			p.expression(9)
		}

	case AlphaCompilerParserT__15:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(237)
			p.Match(AlphaCompilerParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.expression(8)
		}

	case AlphaCompilerParserT__16:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(239)
			p.Match(AlphaCompilerParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(240)
			p.expression(7)
		}

	case AlphaCompilerParserT__17:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(241)
			p.Match(AlphaCompilerParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.expression(6)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(260)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AlphaCompilerParserRULE_expression)
				p.SetState(245)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(246)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&66584576) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(247)
					p.expression(6)
				}

			case 2:
				localctx = NewAddExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AlphaCompilerParserRULE_expression)
				p.SetState(248)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(249)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67469312) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(250)
					p.expression(5)
				}

			case 3:
				localctx = NewRelExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AlphaCompilerParserRULE_expression)
				p.SetState(251)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(252)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8455716864) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(253)
					p.expression(4)
				}

			case 4:
				localctx = NewAndExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AlphaCompilerParserRULE_expression)
				p.SetState(254)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(255)
					p.Match(AlphaCompilerParserT__32)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(256)
					p.expression(3)
				}

			case 5:
				localctx = NewOrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AlphaCompilerParserRULE_expression)
				p.SetState(257)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(258)
					p.Match(AlphaCompilerParserT__33)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(259)
					p.expression(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_expressionList
	return p
}

func InitEmptyExpressionListContext(p *ExpressionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_expressionList
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, AlphaCompilerParserRULE_expressionList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.expression(0)
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == AlphaCompilerParserT__8 {
		{
			p.SetState(266)
			p.Match(AlphaCompilerParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
			p.expression(0)
		}

		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Operand() IOperandContext
	AppendExpression() IAppendExpressionContext
	LengthExpression() ILengthExpressionContext
	CapExpression() ICapExpressionContext
	PrimaryExpression() IPrimaryExpressionContext
	Selector() ISelectorContext
	Index() IIndexContext
	Arguments() IArgumentsContext

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *PrimaryExpressionContext) AppendExpression() IAppendExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAppendExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAppendExpressionContext)
}

func (s *PrimaryExpressionContext) LengthExpression() ILengthExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILengthExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILengthExpressionContext)
}

func (s *PrimaryExpressionContext) CapExpression() ICapExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICapExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICapExpressionContext)
}

func (s *PrimaryExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PrimaryExpressionContext) Selector() ISelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *PrimaryExpressionContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *PrimaryExpressionContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitPrimaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	return p.primaryExpression(0)
}

func (p *AlphaCompilerParser) primaryExpression(_p int) (localctx IPrimaryExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 40
	p.EnterRecursionRule(localctx, 40, AlphaCompilerParserRULE_primaryExpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AlphaCompilerParserT__3, AlphaCompilerParserIDENTIFIER, AlphaCompilerParserINTLITERAL, AlphaCompilerParserFLOATLITERAL, AlphaCompilerParserRUNELITERAL, AlphaCompilerParserRAWSTRINGLITERAL, AlphaCompilerParserINTERPRETEDSTRINGLITERAL:
		{
			p.SetState(274)
			p.Operand()
		}

	case AlphaCompilerParserT__35:
		{
			p.SetState(275)
			p.AppendExpression()
		}

	case AlphaCompilerParserT__36:
		{
			p.SetState(276)
			p.LengthExpression()
		}

	case AlphaCompilerParserT__37:
		{
			p.SetState(277)
			p.CapExpression()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(286)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPrimaryExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, AlphaCompilerParserRULE_primaryExpression)
				p.SetState(280)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(281)
					p.Selector()
				}

			case 2:
				localctx = NewPrimaryExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, AlphaCompilerParserRULE_primaryExpression)
				p.SetState(282)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(283)
					p.Index()
				}

			case 3:
				localctx = NewPrimaryExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, AlphaCompilerParserRULE_primaryExpression)
				p.SetState(284)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(285)
					p.Arguments()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() ILiteralContext
	IDENTIFIER() antlr.TerminalNode
	Expression() IExpressionContext

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_operand
	return p
}

func InitEmptyOperandContext(p *OperandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_operand
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *OperandContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AlphaCompilerParserIDENTIFIER, 0)
}

func (s *OperandContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitOperand(s)
	}
}

func (s *OperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, AlphaCompilerParserRULE_operand)
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AlphaCompilerParserINTLITERAL, AlphaCompilerParserFLOATLITERAL, AlphaCompilerParserRUNELITERAL, AlphaCompilerParserRAWSTRINGLITERAL, AlphaCompilerParserINTERPRETEDSTRINGLITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(291)
			p.Literal()
		}

	case AlphaCompilerParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(292)
			p.Match(AlphaCompilerParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AlphaCompilerParserT__3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(293)
			p.Match(AlphaCompilerParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(294)
			p.expression(0)
		}
		{
			p.SetState(295)
			p.Match(AlphaCompilerParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTLITERAL() antlr.TerminalNode
	FLOATLITERAL() antlr.TerminalNode
	RUNELITERAL() antlr.TerminalNode
	RAWSTRINGLITERAL() antlr.TerminalNode
	INTERPRETEDSTRINGLITERAL() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) INTLITERAL() antlr.TerminalNode {
	return s.GetToken(AlphaCompilerParserINTLITERAL, 0)
}

func (s *LiteralContext) FLOATLITERAL() antlr.TerminalNode {
	return s.GetToken(AlphaCompilerParserFLOATLITERAL, 0)
}

func (s *LiteralContext) RUNELITERAL() antlr.TerminalNode {
	return s.GetToken(AlphaCompilerParserRUNELITERAL, 0)
}

func (s *LiteralContext) RAWSTRINGLITERAL() antlr.TerminalNode {
	return s.GetToken(AlphaCompilerParserRAWSTRINGLITERAL, 0)
}

func (s *LiteralContext) INTERPRETEDSTRINGLITERAL() antlr.TerminalNode {
	return s.GetToken(AlphaCompilerParserINTERPRETEDSTRINGLITERAL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, AlphaCompilerParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&31) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_index
	return p
}

func InitEmptyIndexContext(p *IndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_index
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitIndex(s)
	}
}

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, AlphaCompilerParserRULE_index)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.Match(AlphaCompilerParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(302)
		p.expression(0)
	}
	{
		p.SetState(303)
		p.Match(AlphaCompilerParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionList() IExpressionListContext

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, AlphaCompilerParserRULE_arguments)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Match(AlphaCompilerParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(308)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AlphaCompilerParserT__3, AlphaCompilerParserT__14, AlphaCompilerParserT__15, AlphaCompilerParserT__16, AlphaCompilerParserT__17, AlphaCompilerParserT__35, AlphaCompilerParserT__36, AlphaCompilerParserT__37, AlphaCompilerParserIDENTIFIER, AlphaCompilerParserINTLITERAL, AlphaCompilerParserFLOATLITERAL, AlphaCompilerParserRUNELITERAL, AlphaCompilerParserRAWSTRINGLITERAL, AlphaCompilerParserINTERPRETEDSTRINGLITERAL:
		{
			p.SetState(306)
			p.ExpressionList()
		}

	case AlphaCompilerParserT__4:

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(310)
		p.Match(AlphaCompilerParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsSelectorContext differentiates from other interfaces.
	IsSelectorContext()
}

type SelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorContext() *SelectorContext {
	var p = new(SelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_selector
	return p
}

func InitEmptySelectorContext(p *SelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_selector
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(AlphaCompilerParserIDENTIFIER, 0)
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterSelector(s)
	}
}

func (s *SelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitSelector(s)
	}
}

func (s *SelectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitSelector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, AlphaCompilerParserRULE_selector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Match(AlphaCompilerParserT__34)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(313)
		p.Match(AlphaCompilerParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAppendExpressionContext is an interface to support dynamic dispatch.
type IAppendExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsAppendExpressionContext differentiates from other interfaces.
	IsAppendExpressionContext()
}

type AppendExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAppendExpressionContext() *AppendExpressionContext {
	var p = new(AppendExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_appendExpression
	return p
}

func InitEmptyAppendExpressionContext(p *AppendExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_appendExpression
}

func (*AppendExpressionContext) IsAppendExpressionContext() {}

func NewAppendExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AppendExpressionContext {
	var p = new(AppendExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_appendExpression

	return p
}

func (s *AppendExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AppendExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AppendExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AppendExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AppendExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterAppendExpression(s)
	}
}

func (s *AppendExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitAppendExpression(s)
	}
}

func (s *AppendExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitAppendExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) AppendExpression() (localctx IAppendExpressionContext) {
	localctx = NewAppendExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, AlphaCompilerParserRULE_appendExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(315)
		p.Match(AlphaCompilerParserT__35)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(316)
		p.Match(AlphaCompilerParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(317)
		p.expression(0)
	}
	{
		p.SetState(318)
		p.Match(AlphaCompilerParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(319)
		p.expression(0)
	}
	{
		p.SetState(320)
		p.Match(AlphaCompilerParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILengthExpressionContext is an interface to support dynamic dispatch.
type ILengthExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsLengthExpressionContext differentiates from other interfaces.
	IsLengthExpressionContext()
}

type LengthExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthExpressionContext() *LengthExpressionContext {
	var p = new(LengthExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_lengthExpression
	return p
}

func InitEmptyLengthExpressionContext(p *LengthExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_lengthExpression
}

func (*LengthExpressionContext) IsLengthExpressionContext() {}

func NewLengthExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthExpressionContext {
	var p = new(LengthExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_lengthExpression

	return p
}

func (s *LengthExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LengthExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterLengthExpression(s)
	}
}

func (s *LengthExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitLengthExpression(s)
	}
}

func (s *LengthExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitLengthExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) LengthExpression() (localctx ILengthExpressionContext) {
	localctx = NewLengthExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, AlphaCompilerParserRULE_lengthExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)
		p.Match(AlphaCompilerParserT__36)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(323)
		p.Match(AlphaCompilerParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(324)
		p.expression(0)
	}
	{
		p.SetState(325)
		p.Match(AlphaCompilerParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICapExpressionContext is an interface to support dynamic dispatch.
type ICapExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsCapExpressionContext differentiates from other interfaces.
	IsCapExpressionContext()
}

type CapExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCapExpressionContext() *CapExpressionContext {
	var p = new(CapExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_capExpression
	return p
}

func InitEmptyCapExpressionContext(p *CapExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_capExpression
}

func (*CapExpressionContext) IsCapExpressionContext() {}

func NewCapExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CapExpressionContext {
	var p = new(CapExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_capExpression

	return p
}

func (s *CapExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *CapExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CapExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CapExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CapExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterCapExpression(s)
	}
}

func (s *CapExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitCapExpression(s)
	}
}

func (s *CapExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitCapExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) CapExpression() (localctx ICapExpressionContext) {
	localctx = NewCapExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, AlphaCompilerParserRULE_capExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(327)
		p.Match(AlphaCompilerParserT__37)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(328)
		p.Match(AlphaCompilerParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(329)
		p.expression(0)
	}
	{
		p.SetState(330)
		p.Match(AlphaCompilerParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_statementList
	return p
}

func InitEmptyStatementListContext(p *StatementListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_statementList
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementListContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitStatementList(s)
	}
}

func (s *StatementListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitStatementList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, AlphaCompilerParserRULE_statementList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3747012413439320220) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&63) != 0) {
		{
			p.SetState(332)
			p.Statement()
		}

		p.SetState(337)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StatementList() IStatementListContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, AlphaCompilerParserRULE_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(338)
		p.Match(AlphaCompilerParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(339)
		p.StatementList()
	}
	{
		p.SetState(340)
		p.Match(AlphaCompilerParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionList() IExpressionListContext
	Expression() IExpressionContext
	SimpleStatement() ISimpleStatementContext
	Block() IBlockContext
	Switch_stmt() ISwitch_stmtContext
	IfStatement() IIfStatementContext
	Loop() ILoopContext
	TypeDecl() ITypeDeclContext
	VariableDecl() IVariableDeclContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *StatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) Switch_stmt() ISwitch_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_stmtContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) Loop() ILoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *StatementContext) TypeDecl() ITypeDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *StatementContext) VariableDecl() IVariableDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, AlphaCompilerParserRULE_statement)
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AlphaCompilerParserT__38:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(342)
			p.Match(AlphaCompilerParserT__38)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.Match(AlphaCompilerParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case AlphaCompilerParserT__3, AlphaCompilerParserT__14, AlphaCompilerParserT__15, AlphaCompilerParserT__16, AlphaCompilerParserT__17, AlphaCompilerParserT__35, AlphaCompilerParserT__36, AlphaCompilerParserT__37, AlphaCompilerParserIDENTIFIER, AlphaCompilerParserINTLITERAL, AlphaCompilerParserFLOATLITERAL, AlphaCompilerParserRUNELITERAL, AlphaCompilerParserRAWSTRINGLITERAL, AlphaCompilerParserINTERPRETEDSTRINGLITERAL:
			{
				p.SetState(344)
				p.ExpressionList()
			}

		case AlphaCompilerParserT__4:

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(348)
			p.Match(AlphaCompilerParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(349)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AlphaCompilerParserT__39:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(350)
			p.Match(AlphaCompilerParserT__39)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)
			p.Match(AlphaCompilerParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case AlphaCompilerParserT__3, AlphaCompilerParserT__14, AlphaCompilerParserT__15, AlphaCompilerParserT__16, AlphaCompilerParserT__17, AlphaCompilerParserT__35, AlphaCompilerParserT__36, AlphaCompilerParserT__37, AlphaCompilerParserIDENTIFIER, AlphaCompilerParserINTLITERAL, AlphaCompilerParserFLOATLITERAL, AlphaCompilerParserRUNELITERAL, AlphaCompilerParserRAWSTRINGLITERAL, AlphaCompilerParserINTERPRETEDSTRINGLITERAL:
			{
				p.SetState(352)
				p.ExpressionList()
			}

		case AlphaCompilerParserT__4:

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(356)
			p.Match(AlphaCompilerParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AlphaCompilerParserT__40:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(358)
			p.Match(AlphaCompilerParserT__40)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(361)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case AlphaCompilerParserT__3, AlphaCompilerParserT__14, AlphaCompilerParserT__15, AlphaCompilerParserT__16, AlphaCompilerParserT__17, AlphaCompilerParserT__35, AlphaCompilerParserT__36, AlphaCompilerParserT__37, AlphaCompilerParserIDENTIFIER, AlphaCompilerParserINTLITERAL, AlphaCompilerParserFLOATLITERAL, AlphaCompilerParserRUNELITERAL, AlphaCompilerParserRAWSTRINGLITERAL, AlphaCompilerParserINTERPRETEDSTRINGLITERAL:
			{
				p.SetState(359)
				p.expression(0)
			}

		case AlphaCompilerParserT__1:

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(363)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AlphaCompilerParserT__41:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(364)
			p.Match(AlphaCompilerParserT__41)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(365)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AlphaCompilerParserT__42:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(366)
			p.Match(AlphaCompilerParserT__42)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(367)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AlphaCompilerParserT__1, AlphaCompilerParserT__3, AlphaCompilerParserT__14, AlphaCompilerParserT__15, AlphaCompilerParserT__16, AlphaCompilerParserT__17, AlphaCompilerParserT__35, AlphaCompilerParserT__36, AlphaCompilerParserT__37, AlphaCompilerParserIDENTIFIER, AlphaCompilerParserINTLITERAL, AlphaCompilerParserFLOATLITERAL, AlphaCompilerParserRUNELITERAL, AlphaCompilerParserRAWSTRINGLITERAL, AlphaCompilerParserINTERPRETEDSTRINGLITERAL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(368)
			p.SimpleStatement()
		}
		{
			p.SetState(369)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AlphaCompilerParserT__12:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(371)
			p.Block()
		}
		{
			p.SetState(372)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AlphaCompilerParserT__60:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(374)
			p.Switch_stmt()
		}
		{
			p.SetState(375)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AlphaCompilerParserT__57:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(377)
			p.IfStatement()
		}
		{
			p.SetState(378)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AlphaCompilerParserT__59:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(380)
			p.Loop()
		}
		{
			p.SetState(381)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case AlphaCompilerParserT__6:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(383)
			p.TypeDecl()
		}

	case AlphaCompilerParserT__2:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(384)
			p.VariableDecl()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimpleStatementContext is an interface to support dynamic dispatch.
type ISimpleStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	AssignmentStatement() IAssignmentStatementContext
	AllExpressionList() []IExpressionListContext
	ExpressionList(i int) IExpressionListContext

	// IsSimpleStatementContext differentiates from other interfaces.
	IsSimpleStatementContext()
}

type SimpleStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleStatementContext() *SimpleStatementContext {
	var p = new(SimpleStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_simpleStatement
	return p
}

func InitEmptySimpleStatementContext(p *SimpleStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_simpleStatement
}

func (*SimpleStatementContext) IsSimpleStatementContext() {}

func NewSimpleStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleStatementContext {
	var p = new(SimpleStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_simpleStatement

	return p
}

func (s *SimpleStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SimpleStatementContext) AssignmentStatement() IAssignmentStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *SimpleStatementContext) AllExpressionList() []IExpressionListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionListContext); ok {
			len++
		}
	}

	tst := make([]IExpressionListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionListContext); ok {
			tst[i] = t.(IExpressionListContext)
			i++
		}
	}

	return tst
}

func (s *SimpleStatementContext) ExpressionList(i int) IExpressionListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *SimpleStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterSimpleStatement(s)
	}
}

func (s *SimpleStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitSimpleStatement(s)
	}
}

func (s *SimpleStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitSimpleStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) SimpleStatement() (localctx ISimpleStatementContext) {
	localctx = NewSimpleStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, AlphaCompilerParserRULE_simpleStatement)
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(388)
			p.expression(0)
		}
		p.SetState(392)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case AlphaCompilerParserT__43:
			{
				p.SetState(389)
				p.Match(AlphaCompilerParserT__43)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case AlphaCompilerParserT__44:
			{
				p.SetState(390)
				p.Match(AlphaCompilerParserT__44)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case AlphaCompilerParserT__1, AlphaCompilerParserT__12:

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(394)
			p.AssignmentStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(395)
			p.ExpressionList()
		}
		{
			p.SetState(396)
			p.Match(AlphaCompilerParserT__45)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)
			p.ExpressionList()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentStatementContext is an interface to support dynamic dispatch.
type IAssignmentStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	AllExpressionList() []IExpressionListContext
	ExpressionList(i int) IExpressionListContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsAssignmentStatementContext differentiates from other interfaces.
	IsAssignmentStatementContext()
}

type AssignmentStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyAssignmentStatementContext() *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_assignmentStatement
	return p
}

func InitEmptyAssignmentStatementContext(p *AssignmentStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_assignmentStatement
}

func (*AssignmentStatementContext) IsAssignmentStatementContext() {}

func NewAssignmentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_assignmentStatement

	return p
}

func (s *AssignmentStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentStatementContext) GetOp() antlr.Token { return s.op }

func (s *AssignmentStatementContext) SetOp(v antlr.Token) { s.op = v }

func (s *AssignmentStatementContext) AllExpressionList() []IExpressionListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionListContext); ok {
			len++
		}
	}

	tst := make([]IExpressionListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionListContext); ok {
			tst[i] = t.(IExpressionListContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentStatementContext) ExpressionList(i int) IExpressionListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *AssignmentStatementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentStatementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitAssignmentStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) AssignmentStatement() (localctx IAssignmentStatementContext) {
	localctx = NewAssignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, AlphaCompilerParserRULE_assignmentStatement)
	var _la int

	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(401)
			p.ExpressionList()
		}
		{
			p.SetState(402)
			p.Match(AlphaCompilerParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)
			p.ExpressionList()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(405)
			p.expression(0)
		}
		{
			p.SetState(406)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AssignmentStatementContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&288089638663356416) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AssignmentStatementContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(407)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	IfStatement() IIfStatementContext
	SimpleStatement() ISimpleStatementContext

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfStatementContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, AlphaCompilerParserRULE_ifStatement)
	var _la int

	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(411)
			p.Match(AlphaCompilerParserT__57)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(412)
			p.expression(0)
		}
		{
			p.SetState(413)
			p.Block()
		}
		p.SetState(419)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AlphaCompilerParserT__58 {
			{
				p.SetState(414)
				p.Match(AlphaCompilerParserT__58)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(417)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case AlphaCompilerParserT__57:
				{
					p.SetState(415)
					p.IfStatement()
				}

			case AlphaCompilerParserT__12:
				{
					p.SetState(416)
					p.Block()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(421)
			p.Match(AlphaCompilerParserT__57)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(422)
			p.SimpleStatement()
		}
		{
			p.SetState(423)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(424)
			p.expression(0)
		}
		{
			p.SetState(425)
			p.Block()
		}
		p.SetState(431)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == AlphaCompilerParserT__58 {
			{
				p.SetState(426)
				p.Match(AlphaCompilerParserT__58)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(429)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case AlphaCompilerParserT__57:
				{
					p.SetState(427)
					p.IfStatement()
				}

			case AlphaCompilerParserT__12:
				{
					p.SetState(428)
					p.Block()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	Expression() IExpressionContext
	AllSimpleStatement() []ISimpleStatementContext
	SimpleStatement(i int) ISimpleStatementContext

	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_loop
	return p
}

func InitEmptyLoopContext(p *LoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_loop
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LoopContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LoopContext) AllSimpleStatement() []ISimpleStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			len++
		}
	}

	tst := make([]ISimpleStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleStatementContext); ok {
			tst[i] = t.(ISimpleStatementContext)
			i++
		}
	}

	return tst
}

func (s *LoopContext) SimpleStatement(i int) ISimpleStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterLoop(s)
	}
}

func (s *LoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitLoop(s)
	}
}

func (s *LoopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitLoop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) Loop() (localctx ILoopContext) {
	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, AlphaCompilerParserRULE_loop)
	var _la int

	p.SetState(453)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(435)
			p.Match(AlphaCompilerParserT__59)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(436)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(437)
			p.Match(AlphaCompilerParserT__59)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(438)
			p.expression(0)
		}
		{
			p.SetState(439)
			p.Block()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(441)
			p.Match(AlphaCompilerParserT__59)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(442)
			p.SimpleStatement()
		}
		{
			p.SetState(443)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(445)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&481036828688) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&63) != 0) {
			{
				p.SetState(444)
				p.expression(0)
			}

		}
		{
			p.SetState(447)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(449)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(448)
				p.SimpleStatement()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(451)
			p.Block()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitch_stmtContext is an interface to support dynamic dispatch.
type ISwitch_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionCaseClauseList() IExpressionCaseClauseListContext
	SimpleStatement() ISimpleStatementContext
	Expression() IExpressionContext

	// IsSwitch_stmtContext differentiates from other interfaces.
	IsSwitch_stmtContext()
}

type Switch_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_stmtContext() *Switch_stmtContext {
	var p = new(Switch_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_switch_stmt
	return p
}

func InitEmptySwitch_stmtContext(p *Switch_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_switch_stmt
}

func (*Switch_stmtContext) IsSwitch_stmtContext() {}

func NewSwitch_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_stmtContext {
	var p = new(Switch_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_switch_stmt

	return p
}

func (s *Switch_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_stmtContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *Switch_stmtContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *Switch_stmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Switch_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Switch_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterSwitch_stmt(s)
	}
}

func (s *Switch_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitSwitch_stmt(s)
	}
}

func (s *Switch_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitSwitch_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) Switch_stmt() (localctx ISwitch_stmtContext) {
	localctx = NewSwitch_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, AlphaCompilerParserRULE_switch_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(455)
		p.Match(AlphaCompilerParserT__60)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(459)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(456)
			p.SimpleStatement()
		}
		{
			p.SetState(457)
			p.Match(AlphaCompilerParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&481036828688) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&63) != 0) {
		{
			p.SetState(461)
			p.expression(0)
		}

	}
	{
		p.SetState(464)
		p.Match(AlphaCompilerParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(465)
		p.ExpressionCaseClauseList()
	}
	{
		p.SetState(466)
		p.Match(AlphaCompilerParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionCaseClauseListContext is an interface to support dynamic dispatch.
type IExpressionCaseClauseListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpressionCaseClause() []IExpressionCaseClauseContext
	ExpressionCaseClause(i int) IExpressionCaseClauseContext

	// IsExpressionCaseClauseListContext differentiates from other interfaces.
	IsExpressionCaseClauseListContext()
}

type ExpressionCaseClauseListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionCaseClauseListContext() *ExpressionCaseClauseListContext {
	var p = new(ExpressionCaseClauseListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_expressionCaseClauseList
	return p
}

func InitEmptyExpressionCaseClauseListContext(p *ExpressionCaseClauseListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_expressionCaseClauseList
}

func (*ExpressionCaseClauseListContext) IsExpressionCaseClauseListContext() {}

func NewExpressionCaseClauseListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionCaseClauseListContext {
	var p = new(ExpressionCaseClauseListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_expressionCaseClauseList

	return p
}

func (s *ExpressionCaseClauseListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionCaseClauseListContext) AllExpressionCaseClause() []IExpressionCaseClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionCaseClauseContext); ok {
			len++
		}
	}

	tst := make([]IExpressionCaseClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionCaseClauseContext); ok {
			tst[i] = t.(IExpressionCaseClauseContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionCaseClauseListContext) ExpressionCaseClause(i int) IExpressionCaseClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseContext)
}

func (s *ExpressionCaseClauseListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionCaseClauseListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionCaseClauseListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterExpressionCaseClauseList(s)
	}
}

func (s *ExpressionCaseClauseListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitExpressionCaseClauseList(s)
	}
}

func (s *ExpressionCaseClauseListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitExpressionCaseClauseList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) ExpressionCaseClauseList() (localctx IExpressionCaseClauseListContext) {
	localctx = NewExpressionCaseClauseListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, AlphaCompilerParserRULE_expressionCaseClauseList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(471)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == AlphaCompilerParserT__62 || _la == AlphaCompilerParserT__63 {
		{
			p.SetState(468)
			p.ExpressionCaseClause()
		}

		p.SetState(473)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionCaseClauseContext is an interface to support dynamic dispatch.
type IExpressionCaseClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionSwitchCase() IExpressionSwitchCaseContext
	StatementList() IStatementListContext

	// IsExpressionCaseClauseContext differentiates from other interfaces.
	IsExpressionCaseClauseContext()
}

type ExpressionCaseClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionCaseClauseContext() *ExpressionCaseClauseContext {
	var p = new(ExpressionCaseClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_expressionCaseClause
	return p
}

func InitEmptyExpressionCaseClauseContext(p *ExpressionCaseClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_expressionCaseClause
}

func (*ExpressionCaseClauseContext) IsExpressionCaseClauseContext() {}

func NewExpressionCaseClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionCaseClauseContext {
	var p = new(ExpressionCaseClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_expressionCaseClause

	return p
}

func (s *ExpressionCaseClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionCaseClauseContext) ExpressionSwitchCase() IExpressionSwitchCaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSwitchCaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSwitchCaseContext)
}

func (s *ExpressionCaseClauseContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *ExpressionCaseClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionCaseClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionCaseClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterExpressionCaseClause(s)
	}
}

func (s *ExpressionCaseClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitExpressionCaseClause(s)
	}
}

func (s *ExpressionCaseClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitExpressionCaseClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) ExpressionCaseClause() (localctx IExpressionCaseClauseContext) {
	localctx = NewExpressionCaseClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, AlphaCompilerParserRULE_expressionCaseClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(474)
		p.ExpressionSwitchCase()
	}
	{
		p.SetState(475)
		p.Match(AlphaCompilerParserT__61)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(476)
		p.StatementList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionSwitchCaseContext is an interface to support dynamic dispatch.
type IExpressionSwitchCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionList() IExpressionListContext

	// IsExpressionSwitchCaseContext differentiates from other interfaces.
	IsExpressionSwitchCaseContext()
}

type ExpressionSwitchCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionSwitchCaseContext() *ExpressionSwitchCaseContext {
	var p = new(ExpressionSwitchCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_expressionSwitchCase
	return p
}

func InitEmptyExpressionSwitchCaseContext(p *ExpressionSwitchCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AlphaCompilerParserRULE_expressionSwitchCase
}

func (*ExpressionSwitchCaseContext) IsExpressionSwitchCaseContext() {}

func NewExpressionSwitchCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionSwitchCaseContext {
	var p = new(ExpressionSwitchCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlphaCompilerParserRULE_expressionSwitchCase

	return p
}

func (s *ExpressionSwitchCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionSwitchCaseContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ExpressionSwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSwitchCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionSwitchCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.EnterExpressionSwitchCase(s)
	}
}

func (s *ExpressionSwitchCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlphaCompilerListener); ok {
		listenerT.ExitExpressionSwitchCase(s)
	}
}

func (s *ExpressionSwitchCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AlphaCompilerVisitor:
		return t.VisitExpressionSwitchCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *AlphaCompilerParser) ExpressionSwitchCase() (localctx IExpressionSwitchCaseContext) {
	localctx = NewExpressionSwitchCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, AlphaCompilerParserRULE_expressionSwitchCase)
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case AlphaCompilerParserT__62:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(478)
			p.Match(AlphaCompilerParserT__62)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(479)
			p.ExpressionList()
		}

	case AlphaCompilerParserT__63:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(480)
			p.Match(AlphaCompilerParserT__63)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *AlphaCompilerParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 20:
		var t *PrimaryExpressionContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExpressionContext)
		}
		return p.PrimaryExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *AlphaCompilerParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *AlphaCompilerParser) PrimaryExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
