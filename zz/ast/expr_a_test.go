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

var aExprSub1 = &AExprArith{
	e1: &IntegerLiteral{value: 2},
	e2: aExprAdd2,
	op: AExprArithSub,
}

var aExprSub2 = &AExprArith{
	e1: &IntegerLiteral{value: 1},
	e2: aExprSub1,
	op: AExprArithSub,
}


var aExprMul = &AExprArith{
	e1: &IntegerLiteral{value: 2},
	e2: aExprSub2,
	op: AExprArithMul,
}

var aExprDiv = &AExprArith{
	e1: &IntegerLiteral{value: 1},
	e2: aExprMul,
	op: AExprArithDiv,
}

func TestAExprArithOpType_String(t *testing.T) {
	fmt.Println(AExprArithAdd)
	fmt.Println(AExprArithSub)
	fmt.Println(AExprArithMul)
	fmt.Println(AExprArithDiv)
	fmt.Println(AExprArithOpType(-1))
}

func TestAExprArith_String(t *testing.T) {
	e := aExprDiv
	e.aExpr()
	e.assignIniter()
	e.E1()
	e.E2()
	e.Op()
	fmt.Println(e)
}
