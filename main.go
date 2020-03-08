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

	//e := runtime.New(2, 2)
	//_ = e
	//h := 12
	//_ = h
	//for i := 0; i < 2; i = i + 1 {
	//	_ = i
	//	for j := 0; j < 2; j = j + 1 {
	//		_ = j
	//		e.Get(i).Set(j, runtime.FloatData(i*2+j))
	//	}
	//}
	//f := e.MulMatrix(e)
	//_ = f
	//f1 := e.MulFloat(runtime.FloatData(h))
	//_ = f1
	//
	//fmt.Println(f1)
	//fmt.Println(f)

	fmt.Println(basic_function(1))
	//function1()
	gua()
}

func gua() {
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
	func basic_function(a float) (float) {
		c := a/3
		return c
	}
	`
	//
	//func basic_function(a int) (int) {
	//	a = a * 2
	//	b := a * 2
	//	print(a)
	//
	//	return a
	//}
	//func function(m1 matrix) (matrix) {
	//	m := m1 * transpose(m1)
	//	print(m)
	//
	//	return m
	//}

	//func function1() {
	//	m := matrix(3, 2)
	//	n := 1
	//	for i := 0; i < 3; i = i + 1 {
	//		for j := 0; j < 2; j = j + 1 {
	//			m[i][j] = n
	//			n = n + 1
	//		}
	//	}
	//
	//	m1 := function(m)
	//	print(m1)
	//}

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

	fmt.Println(f)
	if err != nil {
		fmt.Println(err)
	} else {
		tr := transformer.NewTransformer()
		//fmt.Println(f)
		fmt.Println(tr.WalkFile(f))
	}
}
func function(m1 *runtime.Data) *runtime.Data {
	m := m1.MulMatrix(m1.Transpose())
	_ = m
	fmt.Println(m)

	return m
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
	m1 := function(m)
	_ = m1
	fmt.Println(m1)
}

func basic_function(a float64) float64 {
	c := a / 3
	_ = c
	return c
}
