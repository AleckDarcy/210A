// Code generated from zz/grammar/ZZ.g4 by ANTLR 4.7.2. DO NOT EDIT.

package grammar

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 42, 225,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3,
	27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 38, 3, 38, 7, 38, 200, 10, 38, 12, 38, 14, 38, 203, 11, 38,
	3, 39, 3, 39, 3, 39, 7, 39, 208, 10, 39, 12, 39, 14, 39, 211, 11, 39, 5,
	39, 213, 10, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 6, 41, 220, 10, 41,
	13, 41, 14, 41, 221, 3, 41, 3, 41, 2, 2, 42, 3, 3, 5, 4, 7, 5, 9, 6, 11,
	7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16,
	31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25,
	49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34,
	67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 3, 2, 7,
	4, 2, 67, 92, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 51,
	59, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34, 2, 228, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2,
	2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3,
	2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35,
	3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2,
	43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2,
	2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2,
	2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2,
	2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3,
	2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81,
	3, 2, 2, 2, 3, 83, 3, 2, 2, 2, 5, 85, 3, 2, 2, 2, 7, 92, 3, 2, 2, 2, 9,
	94, 3, 2, 2, 2, 11, 96, 3, 2, 2, 2, 13, 98, 3, 2, 2, 2, 15, 100, 3, 2,
	2, 2, 17, 102, 3, 2, 2, 2, 19, 104, 3, 2, 2, 2, 21, 106, 3, 2, 2, 2, 23,
	108, 3, 2, 2, 2, 25, 111, 3, 2, 2, 2, 27, 113, 3, 2, 2, 2, 29, 115, 3,
	2, 2, 2, 31, 118, 3, 2, 2, 2, 33, 121, 3, 2, 2, 2, 35, 124, 3, 2, 2, 2,
	37, 127, 3, 2, 2, 2, 39, 130, 3, 2, 2, 2, 41, 132, 3, 2, 2, 2, 43, 134,
	3, 2, 2, 2, 45, 137, 3, 2, 2, 2, 47, 139, 3, 2, 2, 2, 49, 141, 3, 2, 2,
	2, 51, 147, 3, 2, 2, 2, 53, 152, 3, 2, 2, 2, 55, 154, 3, 2, 2, 2, 57, 156,
	3, 2, 2, 2, 59, 160, 3, 2, 2, 2, 61, 162, 3, 2, 2, 2, 63, 169, 3, 2, 2,
	2, 65, 173, 3, 2, 2, 2, 67, 177, 3, 2, 2, 2, 69, 183, 3, 2, 2, 2, 71, 188,
	3, 2, 2, 2, 73, 193, 3, 2, 2, 2, 75, 197, 3, 2, 2, 2, 77, 212, 3, 2, 2,
	2, 79, 214, 3, 2, 2, 2, 81, 219, 3, 2, 2, 2, 83, 84, 7, 46, 2, 2, 84, 4,
	3, 2, 2, 2, 85, 86, 7, 117, 2, 2, 86, 87, 7, 118, 2, 2, 87, 88, 7, 116,
	2, 2, 88, 89, 7, 107, 2, 2, 89, 90, 7, 112, 2, 2, 90, 91, 7, 105, 2, 2,
	91, 6, 3, 2, 2, 2, 92, 93, 7, 93, 2, 2, 93, 8, 3, 2, 2, 2, 94, 95, 7, 95,
	2, 2, 95, 10, 3, 2, 2, 2, 96, 97, 7, 42, 2, 2, 97, 12, 3, 2, 2, 2, 98,
	99, 7, 43, 2, 2, 99, 14, 3, 2, 2, 2, 100, 101, 7, 44, 2, 2, 101, 16, 3,
	2, 2, 2, 102, 103, 7, 49, 2, 2, 103, 18, 3, 2, 2, 2, 104, 105, 7, 45, 2,
	2, 105, 20, 3, 2, 2, 2, 106, 107, 7, 47, 2, 2, 107, 22, 3, 2, 2, 2, 108,
	109, 7, 63, 2, 2, 109, 110, 7, 63, 2, 2, 110, 24, 3, 2, 2, 2, 111, 112,
	7, 62, 2, 2, 112, 26, 3, 2, 2, 2, 113, 114, 7, 64, 2, 2, 114, 28, 3, 2,
	2, 2, 115, 116, 7, 64, 2, 2, 116, 117, 7, 63, 2, 2, 117, 30, 3, 2, 2, 2,
	118, 119, 7, 62, 2, 2, 119, 120, 7, 63, 2, 2, 120, 32, 3, 2, 2, 2, 121,
	122, 7, 35, 2, 2, 122, 123, 7, 63, 2, 2, 123, 34, 3, 2, 2, 2, 124, 125,
	7, 40, 2, 2, 125, 126, 7, 40, 2, 2, 126, 36, 3, 2, 2, 2, 127, 128, 7, 126,
	2, 2, 128, 129, 7, 126, 2, 2, 129, 38, 3, 2, 2, 2, 130, 131, 7, 35, 2,
	2, 131, 40, 3, 2, 2, 2, 132, 133, 7, 63, 2, 2, 133, 42, 3, 2, 2, 2, 134,
	135, 7, 107, 2, 2, 135, 136, 7, 104, 2, 2, 136, 44, 3, 2, 2, 2, 137, 138,
	7, 125, 2, 2, 138, 46, 3, 2, 2, 2, 139, 140, 7, 127, 2, 2, 140, 48, 3,
	2, 2, 2, 141, 142, 7, 103, 2, 2, 142, 143, 7, 110, 2, 2, 143, 144, 7, 117,
	2, 2, 144, 145, 7, 107, 2, 2, 145, 146, 7, 104, 2, 2, 146, 50, 3, 2, 2,
	2, 147, 148, 7, 103, 2, 2, 148, 149, 7, 110, 2, 2, 149, 150, 7, 117, 2,
	2, 150, 151, 7, 103, 2, 2, 151, 52, 3, 2, 2, 2, 152, 153, 7, 65, 2, 2,
	153, 54, 3, 2, 2, 2, 154, 155, 7, 60, 2, 2, 155, 56, 3, 2, 2, 2, 156, 157,
	7, 104, 2, 2, 157, 158, 7, 113, 2, 2, 158, 159, 7, 116, 2, 2, 159, 58,
	3, 2, 2, 2, 160, 161, 7, 61, 2, 2, 161, 60, 3, 2, 2, 2, 162, 163, 7, 116,
	2, 2, 163, 164, 7, 103, 2, 2, 164, 165, 7, 118, 2, 2, 165, 166, 7, 119,
	2, 2, 166, 167, 7, 116, 2, 2, 167, 168, 7, 112, 2, 2, 168, 62, 3, 2, 2,
	2, 169, 170, 7, 120, 2, 2, 170, 171, 7, 99, 2, 2, 171, 172, 7, 116, 2,
	2, 172, 64, 3, 2, 2, 2, 173, 174, 7, 107, 2, 2, 174, 175, 7, 112, 2, 2,
	175, 176, 7, 118, 2, 2, 176, 66, 3, 2, 2, 2, 177, 178, 7, 104, 2, 2, 178,
	179, 7, 110, 2, 2, 179, 180, 7, 113, 2, 2, 180, 181, 7, 99, 2, 2, 181,
	182, 7, 118, 2, 2, 182, 68, 3, 2, 2, 2, 183, 184, 7, 104, 2, 2, 184, 185,
	7, 119, 2, 2, 185, 186, 7, 112, 2, 2, 186, 187, 7, 101, 2, 2, 187, 70,
	3, 2, 2, 2, 188, 189, 7, 110, 2, 2, 189, 190, 7, 107, 2, 2, 190, 191, 7,
	117, 2, 2, 191, 192, 7, 118, 2, 2, 192, 72, 3, 2, 2, 2, 193, 194, 7, 112,
	2, 2, 194, 195, 7, 107, 2, 2, 195, 196, 7, 110, 2, 2, 196, 74, 3, 2, 2,
	2, 197, 201, 9, 2, 2, 2, 198, 200, 9, 3, 2, 2, 199, 198, 3, 2, 2, 2, 200,
	203, 3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 76, 3,
	2, 2, 2, 203, 201, 3, 2, 2, 2, 204, 213, 7, 50, 2, 2, 205, 209, 9, 4, 2,
	2, 206, 208, 9, 5, 2, 2, 207, 206, 3, 2, 2, 2, 208, 211, 3, 2, 2, 2, 209,
	207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 213, 3, 2, 2, 2, 211, 209,
	3, 2, 2, 2, 212, 204, 3, 2, 2, 2, 212, 205, 3, 2, 2, 2, 213, 78, 3, 2,
	2, 2, 214, 215, 7, 50, 2, 2, 215, 216, 7, 48, 2, 2, 216, 217, 7, 50, 2,
	2, 217, 80, 3, 2, 2, 2, 218, 220, 9, 6, 2, 2, 219, 218, 3, 2, 2, 2, 220,
	221, 3, 2, 2, 2, 221, 219, 3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222, 223,
	3, 2, 2, 2, 223, 224, 8, 41, 2, 2, 224, 82, 3, 2, 2, 2, 7, 2, 201, 209,
	212, 221, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "'string'", "'['", "']'", "'('", "')'", "'*'", "'/'", "'+'",
	"'-'", "'=='", "'<'", "'>'", "'>='", "'<='", "'!='", "'&&'", "'||'", "'!'",
	"'='", "'if'", "'{'", "'}'", "'elsif'", "'else'", "'?'", "':'", "'for'",
	"';'", "'return'", "'var'", "'int'", "'float'", "'func'", "'list'", "'nil'",
	"", "", "'0.0'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "Var", "Int", "Float",
	"Func", "List", "Nil", "Identifier", "IntegerLiteral", "FloatLiteral",
	"WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "Var", "Int", "Float", "Func",
	"List", "Nil", "Identifier", "IntegerLiteral", "FloatLiteral", "WS",
}

type ZZLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewZZLexer(input antlr.CharStream) *ZZLexer {

	l := new(ZZLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "ZZ.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ZZLexer tokens.
const (
	ZZLexerT__0           = 1
	ZZLexerT__1           = 2
	ZZLexerT__2           = 3
	ZZLexerT__3           = 4
	ZZLexerT__4           = 5
	ZZLexerT__5           = 6
	ZZLexerT__6           = 7
	ZZLexerT__7           = 8
	ZZLexerT__8           = 9
	ZZLexerT__9           = 10
	ZZLexerT__10          = 11
	ZZLexerT__11          = 12
	ZZLexerT__12          = 13
	ZZLexerT__13          = 14
	ZZLexerT__14          = 15
	ZZLexerT__15          = 16
	ZZLexerT__16          = 17
	ZZLexerT__17          = 18
	ZZLexerT__18          = 19
	ZZLexerT__19          = 20
	ZZLexerT__20          = 21
	ZZLexerT__21          = 22
	ZZLexerT__22          = 23
	ZZLexerT__23          = 24
	ZZLexerT__24          = 25
	ZZLexerT__25          = 26
	ZZLexerT__26          = 27
	ZZLexerT__27          = 28
	ZZLexerT__28          = 29
	ZZLexerT__29          = 30
	ZZLexerVar            = 31
	ZZLexerInt            = 32
	ZZLexerFloat          = 33
	ZZLexerFunc           = 34
	ZZLexerList           = 35
	ZZLexerNil            = 36
	ZZLexerIdentifier     = 37
	ZZLexerIntegerLiteral = 38
	ZZLexerFloatLiteral   = 39
	ZZLexerWS             = 40
)
