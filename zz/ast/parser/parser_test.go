package parser

import (
	"reflect"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/AleckDarcy/210A/zz/ast"
	"github.com/AleckDarcy/210A/zz/grammar"
)

type TestCase struct {
	input  string
	expect ast.BasicNoder
}

func GenerateParser(input string) (*grammar.ZZParser, *ParseTreeListener) {
	inputStream := antlr.NewInputStream(input)
	lexer := grammar.NewZZLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	parser := &ParseTreeListener{}

	zzParser := grammar.NewZZParser(stream)
	zzParser.SetErrorHandler(antlr.NewDefaultErrorStrategy())
	zzParser.AddErrorListener(antlr.NewDefaultErrorListener())
	zzParser.AddParseListener(parser)
	zzParser.BuildParseTrees = true

	return zzParser, parser
}

func JudgeResult(t *testing.T, parser *ParseTreeListener, expect ast.BasicNoder) (ast.BasicNoder, error) {
	node, err := parser.stack.PopByType(ast.NoderBasic)
	if err != nil {
		t.Error(err)
	} else if reflect.DeepEqual(node, expect) {
		t.Log("success")
	} else {
		t.Errorf(""+
			"Judge fail\n"+
			"Expect:\n"+
			"%s\n"+
			"Get:\n"+
			"%s\n", expect, node)
	}

	return node, err
}

func PopAll(t *testing.T, parser *ParseTreeListener) {
	for {
		item, _ := parser.stack.PopByType(ast.NoderBasic)
		if item == nil {
			return
		}

		t.Log(item)
	}
}

func TestParseTreeListener_AssignStatement(t *testing.T) {
	testCases := []TestCase{
		{input: "a = list([]int, 4)", expect: ast.AssignStmt1},
		{input: "a[1], b = 2, 3", expect: ast.AssignStmt2},
	}

	for i, testCase := range testCases {
		t.Logf("testing %d:\n%s\n", i, testCase.input)
		zzParser, parser := GenerateParser(testCase.input)
		zzParser.AssignStatement()
		_, _ = JudgeResult(t, parser, testCase.expect)
	}
}

func TestParseTreeListener_SelectionStatement(t *testing.T) {
	testCases := []TestCase{
		{
			input: "" +
				"if a[1] == b {\n" +
				"	a[1] = a[1] + 1\n" +
				"} elsif a[1] == b + 1 {\n" +
				"	a[1] = a[1] + 2\n" +
				"} else {\n" +
				"	a[1] = 0\n" +
				"}",
			expect: ast.SelectionStmt1,
		},
		{
			input: "" +
				"if a[1] == b {\n" +
				"	a[1] = a[1] + 1\n" +
				"} else {\n" +
				"	a[1] = 0\n" +
				"}",
			expect: ast.SelectionStmt2,
		},
		{input: "a[1] == b ? a[1] = a[1] + 1 : a[1] = 0", expect: ast.SelectionStmt2},
		{
			input: "" +
				"if a[1] == b {\n" +
				"	a[1] = a[1] + 1\n" +
				"}",
			expect: ast.SelectionStmt3,
		},
	}

	for i, testCase := range testCases {
		t.Logf("testing %d:\n%s\n", i, testCase.input)
		zzParser, parser := GenerateParser(testCase.input)
		zzParser.SelectionStatement()
		_, _ = JudgeResult(t, parser, testCase.expect)
	}
}

func TestParseTreeListener_FuncDefinition(t *testing.T) {
	testCases := []TestCase{
		{input: "func function() {}", expect: ast.FuncDefinition1},
		{
			input: "" +
				"func function1(x, y int, z []float) ([]int, float, int) {\n" +
				"	a = list([]int, 4)\n" +
				"	b = x + y\n" +
				"	return a, b + 1, 1\n" +
				"}",
			expect: ast.FuncDefinition2,
		},
	}

	for i, testCase := range testCases {
		t.Logf("testing %d:\n%s\n", i, testCase.input)
		zzParser, parser := GenerateParser(testCase.input)
		zzParser.FuncDefinition()
		_, _ = JudgeResult(t, parser, testCase.expect)
	}
}

func TestTemplate(t *testing.T) {
	testCases := []TestCase{
		{input: "a = list([]int, 4)", expect: ast.AssignStmt1},
		{input: "a[1], b = 2, 3", expect: ast.AssignStmt2},
	}

	for i, testCase := range testCases {
		t.Logf("testing %d:\n%s\n", i, testCase.input)
		//zzParser, parser := GenerateParser(testCase.input)
		//zzParser.AssignStatement()
		//_, _ = JudgeResult(t, parser, testCase.expect)
	}
}
