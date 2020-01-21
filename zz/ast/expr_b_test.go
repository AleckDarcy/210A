package ast

import (
	"fmt"
	"testing"
)

var bExprCompare1 = &BExprCompare{
	e1: aExprAdd1,
	e2: &IntegerLiteral{value: 5},
	op: BExprCompareEQ,
}

var bExprCompare2 = &BExprCompare{
	e1: &IntegerLiteral{value: 2},
	e2: &IntegerLiteral{value: 2},
	op: BExprCompareEQ,
}

var bExprBinary1 = &BExprBinary{
	e1: bExprCompare1,
	e2: bExprCompare2,
	op: BExprBinaryEQ,
}

var bExprBinary2 = &BExprBinary{
	e1: &BinaryLiteral{value: true},
	e2: bExprCompare2,
	op: BExprBinaryEQ,
}

func TestBExprCompareOpType_String(t *testing.T) {
	fmt.Println(BExprCompareEQ)
	fmt.Println(BExprCompareOpType(-1))
}

func TestBExprCompare_String(t *testing.T) {
	e := bExprCompare1
	e.bExpr()
	fmt.Println(e)
}

func TestBExprBinaryOpType_String(t *testing.T) {
	fmt.Println(BExprBinaryEQ)
	fmt.Println(BExprBinaryOpType(-1))
}

func TestBExprBinary_String(t *testing.T) {
	e := bExprBinary1
	e.bExpr()
	fmt.Println(e)
}
