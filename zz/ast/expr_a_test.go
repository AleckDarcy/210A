package ast

import (
	"fmt"
	"testing"
)

func TestAExprArithOpType_String(t *testing.T) {
	fmt.Println(AExprArithAdd)
	fmt.Println(AExprArithSub)
	fmt.Println(AExprArithMul)
	fmt.Println(AExprArithDiv)
	fmt.Println(AExprArithOpType(-1))
}

func TestAExprArithOpType_PriorTo_String(t *testing.T) {
	fmt.Println(AExprArithDiv.PriorTo(AExprArithAdd))
	fmt.Println(AExprArithMul.PriorTo(AExprArithAdd))
	fmt.Println(AExprArithAdd.PriorTo(AExprArithSub))

	test := func() {
		defer func() {
			if err := recover(); err == nil {
				panic("fail")
			} else {
				t.Log(err)
			}
		}()
		fmt.Println(AExprArithDiv.PriorTo(AExprArithOpType(-1)))
	}
	test()
}

func TestAExprSimple_String(t *testing.T) {
	e := AExprSimple1
	e.aExpr()
	e.assignIniter()
	e.funcReturnParaer()
	e.funcExecuteParaer()
	e.E()
	fmt.Println(e)
}

func TestAExprArith_String(t *testing.T) {
	e := AExprDiv
	e.aExpr()
	e.assignIniter()
	e.funcReturnParaer()
	e.funcExecuteParaer()
	e.E1()
	e.E2()
	e.Op()
	fmt.Println(e)
}

func TestAExprArithTranspose_String(t *testing.T) {
	e := AExprTranspose1
	e.aExpr()
	e.assignIniter()
	e.E()
	fmt.Println(e)
}
