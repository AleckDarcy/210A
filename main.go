package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/AleckDarcy/210A/zz/ast/parser"
	"github.com/AleckDarcy/210A/zz/grammar"
	"github.com/AleckDarcy/210A/zz/runtime"
	"github.com/AleckDarcy/210A/zz/transformer"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	var implType string
	var filePath string

	flag.StringVar(&implType, "t", "i", "i: interpreted; c: complied")
	flag.StringVar(&filePath, "f", "", "file path for .zz file")
	flag.Parse()

	fmt.Printf("arguments: [-t %s] [-f %s]\n", implType, filePath)

	if !strings.HasSuffix(filePath, ".zz") {
		fmt.Println("invalid file path")
		return
	} else if implType != "i" && implType != "c" {
		fmt.Println("invalid implement type")
		return
	}

	fmt.Println("reading .zz file")
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println("transforming code")
	code, err := generate(string(bytes))

	code = fmt.Sprintf(
		`package main

import(
	"fmt"

	"github.com/AleckDarcy/210A/zz/runtime"
)

var _ = fmt.Formatter(nil)
var _ = runtime.FloatData(0)

%s
`,
		code,
	)

	absPath, _ := filepath.Abs("output")
	if ok, err := pathExists("output"); err != nil {
		fmt.Printf("check output directory err: %s\n", err)
	} else if ok {
		if err = filepath.Walk(absPath, func(path string, fi os.FileInfo, err error) error {
			if nil == fi {
				return err
			} else if fi.IsDir() {
				return nil
			}

			if strings.HasSuffix(path, ".go") {
				os.Remove(path)
			}

			return nil
		}); err != nil {
			fmt.Printf("clear output directory err: %s\n", err)
			return
		}
	} else if err = os.Mkdir("output", os.ModePerm); err != nil {
		fmt.Printf("make output directory err: %s\n", err)
		return
	}

	f, err := os.OpenFile("output/zz.go", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Printf("make output file err: %s\n", err)
		return
	}

	fmt.Println("writing code")

	if _, err = f.Write([]byte(code)); err != nil {
		fmt.Printf("write output file err: %s\n", err)
		return
	}

	if implType == "i" {
		fmt.Println("executing ZZ as an Interpreted Language")
		cmd := exec.Command("go", "run", "output/zz.go")
		bytes, err := cmd.Output()
		if err != nil {
			fmt.Printf("executing ZZ as an Interpreted Language err: %s\n", err)
			return
		}

		fmt.Printf("output:\n%s\nfinished\n", string(bytes))
	} else {
		fmt.Println("compiling ZZ as a Compiled Language")
		cmd := exec.Command("go", "build", "-o", "output/main", "output/zz.go")
		_, err := cmd.Output()
		if err != nil {
			fmt.Printf("executing ZZ as an Interpreted Language err: %s\n", err)
			return
		}

		fmt.Println("executable file: output/main")
		fmt.Println("finished")
	}
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func generate(input string) (string, error) {
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
		return "", errors.New("error occurs when parsing")
	}

	f, err := p.Pop()

	//fmt.Println(f)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	tr := transformer.NewTransformer()
	//fmt.Println(f)
	return tr.WalkFile(f), nil
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

	//// basic function
	//input := `
	//func basicFunction(a int) (int) {
	//	return a
	//}
	//`

	//// test int1
	//input := `
	//func testInt1(a int) (int) {
	//	b := a + 1
	//	return b
	//}
	//`

	//// test int2
	//input := `
	//func testInt2(a int) (int) {
	//	a = a + 1
	//	return a
	//}
	//`

	//// test float
	//input := `
	//func testFloat(a float) (float) {
	//	b := a/3
	//	return b
	//}
	//`

	//// test matrixNew0
	//input := `
	//	func matrixNew0() (matrix) {
	//		result := matrix(1,3)
	//		result[0][0] = 1
	//		result[0][1] = 1
	//		result[0][2] = 1
	//		return result
	//	}
	//`

	//// test matrixNew1
	//input := `
	//	func matrixNew1() (matrix) {
	//		result := matrix(3)
	//		result[0] = 1
	//		result[1] = 1
	//		result[2] = 1
	//		return result
	//	}
	//`

	//// test matrixNew2
	//input := `
	//	func matrixNew2() (matrix) {
	//		result := matrix(3,1)
	//		result[0][0] = 1
	//		result[1][0] = 1
	//		result[2][0] = 1
	//		return result
	//	}
	//`

	//// test matrixNew3
	//input := `
	//	func matrixNew3() (matrix) {
	//		result := matrix(2,2)
	//		result[0][0] = 1
	//		result[0][1] = 2
	//		result[1][0] = 3
	//		result[1][1] = 4
	//		return result
	//	}
	//`

	//// test matrixAdd
	//input := `
	//	func matrixAdd(m1, m2 matrix) (matrix) {
	//		result := m1 + m2
	//		return result
	//	}
	//`

	//// test matrixMul
	//input := `
	//	func matrixMul(m1, m2 matrix) (matrix) {
	//		result := m1 * m2
	//		return result
	//	}
	//`

	// test matrixTranspose
	input := `
		func matrixT(m matrix) (matrix) {
			result := transpose(m)
			return result
		}
	`

	//func basic_function(a float) (float) {
	//	c := a/3
	//	return c
	//}
	//

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

func basicFunction(a int) int {
	return a
}

func testInt1(a int) int {
	b := a + 1
	_ = b
	return b
}

func testInt2(a int) int {
	a = a + 1
	return a
}

func testFloat(a float64) float64 {
	b := a / 3
	_ = b
	return b
}

func matrixNew1() *runtime.Data {
	result := runtime.New(3)
	_ = result
	result.Set(0, runtime.FloatData(1))
	result.Set(1, runtime.FloatData(1))
	result.Set(2, runtime.FloatData(1))
	return result
}

func matrixNew2() *runtime.Data {
	result := runtime.New(3, 1)
	_ = result
	result.Get(0).Set(0, runtime.FloatData(1))
	result.Get(1).Set(0, runtime.FloatData(1))
	result.Get(2).Set(0, runtime.FloatData(1))
	return result
}

func matrixNew3() *runtime.Data {
	result := runtime.New(2, 2)
	_ = result
	result.Get(0).Set(0, runtime.FloatData(1))
	result.Get(0).Set(1, runtime.FloatData(2))
	result.Get(1).Set(0, runtime.FloatData(3))
	result.Get(1).Set(1, runtime.FloatData(4))
	return result
}

func matrixAdd(m1, m2 *runtime.Data) *runtime.Data {
	result := m1.AddMatrix(m2)
	_ = result
	return result
}

func matrixMul(m1, m2 *runtime.Data) *runtime.Data {
	result := m1.MulMatrix(m2)
	_ = result
	return result
}

func matrixNew0() *runtime.Data {
	result := runtime.New(1, 3)
	_ = result
	result.Get(0).Set(0, runtime.FloatData(1))
	result.Get(0).Set(1, runtime.FloatData(1))
	result.Get(0).Set(2, runtime.FloatData(1))
	return result
}

func matrixT(m *runtime.Data) *runtime.Data {
	result := m.Transpose()
	_ = result
	return result
}
