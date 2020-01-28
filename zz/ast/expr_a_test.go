package ast

import (
	"fmt"
	"testing"
)

var aExprAdd1 = &AExprArith{
	e1: &IntegerLiteral{value: 2},
	e2: &IntegerLiteral{value: 3},
	op: AExprArithAdd,
}

var aExprAdd2 = &AExprArith{
	e1: &IntegerLiteral{value: 1},
	e2: aExprAdd1,
	op: AExprArithAdd,
}

func TestAExprArithOpType_String(t *testing.T) {
	fmt.Println(AExprArithAdd)
	fmt.Println(AExprArithSub)
	fmt.Println(AExprArithMul)
	fmt.Println(AExprArithOpType(-1))
}

func TestAExprArith_String(t *testing.T) {
	e := aExprAdd1
	e.aExpr()
	e.assignIniter()
	e.E1()
	e.E2()
	e.Op()
	fmt.Println(e)
}
