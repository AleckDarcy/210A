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

func TestAExprArith_String(t *testing.T) {
	e := aExprAdd1

	fmt.Println(e.String())
}