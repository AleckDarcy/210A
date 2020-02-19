package ast

import (
	"fmt"
	"testing"
)

func TestBExprCompareOpType_String(t *testing.T) {
	fmt.Println(BExprCompareEQ)
	fmt.Println(BExprCompareLT)
	fmt.Println(BExprCompareGT)
	fmt.Println(BExprCompareLEQ)
	fmt.Println(BExprCompareGEQ)
	fmt.Println(BExprCompareNEQ)
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
	fmt.Println(BExprBinaryAND)
	fmt.Println(BExprBinaryOR)
	fmt.Println(BExprBinaryNEQ)
	fmt.Println(BExprBinaryNOT)
	fmt.Println(BExprBinaryOpType(-1))
}

func TestBExprBinary_String(t *testing.T) {
	e := BExprCompare7
	e.bExpr()
	e.funcReturnParaer()
	e.funcExecuteParaer()
	fmt.Println(e)
}
