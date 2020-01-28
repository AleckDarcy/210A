package ast

import (
	"fmt"
	"testing"
)

func TestBExprCompareOpType_String(t *testing.T) {
	fmt.Println(BExprCompareEQ)
	fmt.Println(BExprCompareLT)
	fmt.Println(BExprCompareOpType(-1))
}

func TestBExprCompare_String(t *testing.T) {
	e := BExprCompare1
	e.bExpr()
	e.funcReturnParaer()
	e.funcExecuteParaer()
	fmt.Println(e)
}

func TestBExprBinaryOpType_String(t *testing.T) {
	fmt.Println(BExprBinaryEQ)
	fmt.Println(BExprBinaryOpType(-1))
}

func TestBExprBinary_String(t *testing.T) {
	e := BExprBinary1
	e.bExpr()
	e.funcReturnParaer()
	e.funcExecuteParaer()
	fmt.Println(e)
}
