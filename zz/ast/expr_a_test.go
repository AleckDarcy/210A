package ast

import (
	"fmt"
	"testing"
)

func TestAExprArithOpType_String(t *testing.T) {
	fmt.Println(AExprArithAdd)
	fmt.Println(AExprArithOpType(-1))
}

func TestAExprSimple_String(t *testing.T) {
	e := AExprSimple1
	e.aExpr()
	e.assignIniter()
	e.funcReturnParaer()
	e.funcExecuteParaer()
	fmt.Println(e)
}

func TestAExprArith_String(t *testing.T) {
	e := AExprAdd1
	e.aExpr()
	e.assignIniter()
	e.funcReturnParaer()
	e.funcExecuteParaer()
	e.E1()
	e.E2()
	e.Op()
	fmt.Println(e)
}
