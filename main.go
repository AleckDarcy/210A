package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/AleckDarcy/210A/zz/ast/parser"
	"github.com/AleckDarcy/210A/zz/grammar"
	"github.com/AleckDarcy/210A/zz/transformer"
)

func main() {
	input := `
	length := 3
	
	func MatrixMultiplication(m1, m2 [][]int) ([][]int) {
		result := list([][]int, length)
		for i := 0; i < length; i = i + 1 {
			result[i] = list([]int, length)
		}
	
		for i := 0; i < length; i = i + 1 {
			temp := 0
			for j := 0; j < length; j = j + 1 {
				for k := 0; k < length; k = k + 1 {
					temp = temp + m1[i][k] + m2[k][j]
				}
	
				result[i][j] = temp
			}
		}
	
		return result
	}
	
	func Ternary(a, b int) (int) {
		result := 2
		a < b ? result = 1 : result = 0

	return result
}

func function() (int) {
	return Ternary(1, 2)
}
`

	inputStream := antlr.NewInputStream(input)
	lexer := grammar.NewZZLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := &parser.ParseTreeListener{}

	zzParser := grammar.NewZZParser(stream)
	zzParser.SetErrorHandler(antlr.NewDefaultErrorStrategy())
	zzParser.AddErrorListener(antlr.NewDefaultErrorListener())
	zzParser.AddParseListener(p)
	zzParser.BuildParseTrees = true
	zzParser.File()

	f, err := p.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		tr := transformer.NewTransformer()
		//fmt.Println(f)
		fmt.Println(tr.WalkFile(f))
	}
}

var length = 3

func MatrixMultiplication(m1, m2 [][]int) [][]int {
	result := make([][]int, length)
	for i := 0; i < length; i = i + 1 {
		result[i] = make([]int, length)
	}
	for i := 0; i < length; i = i + 1 {
		temp := 0
		for j := 0; j < length; j = j + 1 {
			for k := 0; k < length; k = k + 1 {
				temp = temp + m1[i][k] + m2[k][j]
			}
			result[i][j] = temp
		}
	}
	return result
}

func Ternary(a, b int) int {
	result := 1
	if a < b {
		result = 1
	} else {
		result = 0
	}
	return result
}
