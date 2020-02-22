package main

import (
	"fmt"

	"github.com/AleckDarcy/210A/zz/runtime"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/AleckDarcy/210A/zz/ast/parser"
	"github.com/AleckDarcy/210A/zz/grammar"
	"github.com/AleckDarcy/210A/zz/transformer"
)

func main() {
	//input := `
	//	length := 3
	//
	//	func MatrixMultiplication(m1, m2 matrix) ([][]int) {
	//		result := matrix(3, 3)
	//
	//		for i := 0; i < length; i = i + 1 {
	//			result[i] = list([]int, length)
	//		}
	//
	//		for i := 0; i < length; {
	//			temp := 0
	//			for j := 0; j < length; j = j + 1 {
	//				for k := 0; k < length; k = k + 1 {
	//					temp = temp + m1[i][k] + m2[k][j]
	//				}
	//
	//				result[i][j] = temp
	//			}
	//
	//			i = i + 1
	//		}
	//
	//		return result
	//	}
	//
	//	func Ternary(a, b int) (int) {
	//		result := 2
	//		a < b ? result = 1 : result = 0
	//
	//	return result
	//}
	//
	//func function() (int) {
	//	return Ternary(1, 2)
	//}
	//`

	input := `
	func function(m1 matrix) {
		m2 := transpose(m1)
	
		m := m1 * m2
	
		print(m1, m2)
		print(m)
	}

	func function1() {
		m := matrix(3, 2)
		n := 1
		for i := 0; i < 3; i = i + 1 {
			for j := 0; j < 2; j = j + 1 {
				m[i][j] = n
				n = n + 1
			}
		}

		function(m)
	}
	`

	//m2 := matrix(2, 4)
	//
	//n := 1
	//for i := 0; i < 3; i = i + 1 {
	//	for j := 0; j < 2; j = j + 1 {
	//		m1[i][j] = n
	//		n = n + 1
	//	}
	//}
	//
	//for i := 0; i < 2; i = i + 1 {
	//	for j := 0; j < 4; j = j + 1 {
	//		m2[i][j] = n
	//		n = n + 1
	//	}
	//}
	//
	//a := m1 * (m2 * 0.5*2)

	inputStream := antlr.NewInputStream(input)
	lexer := grammar.NewZZLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := &parser.ParseTreeListener{}

	zzParser := grammar.NewZZParser(stream)
	zzParser.SetErrorHandler(antlr.NewDefaultErrorStrategy())
	zzParser.AddErrorListener(p)
	zzParser.AddParseListener(p)
	zzParser.BuildParseTrees = true
	zzParser.File()

	if p.ErrorFlag() {
		fmt.Println("error occurs when parsing")
		return
	}

	f, err := p.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		tr := transformer.NewTransformer()
		//fmt.Println(f)
		fmt.Println(tr.WalkFile(f))
	}
}
func function(m1 *runtime.Data) {
	m2 := m1.Transpose()
	_ = m2
	m := m1.MulMatrix(m2)
	_ = m
	fmt.Println(m1, m2)
	fmt.Println(m)
}

func function1() {
	m := runtime.New(3, 2)
	_ = m
	n := 1
	_ = n
	for i := 0; i < 3; i = i + 1 {
		_ = i
		for j := 0; j < 2; j = j + 1 {
			_ = j
			m.Get(i).Set(j, runtime.FloatData(n))
			n = n + 1
		}
	}
	function(m)
}
