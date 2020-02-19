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

func TestAExprSimple_String_1(t *testing.T) {
	e := AExprSimple1
	e.aExpr()
	e.assignIniter()
	e.funcReturnParaer()
	e.funcExecuteParaer()
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
