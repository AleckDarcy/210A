package ast

import (
	"fmt"
	"testing"
)

func TestAExprArith_String(t *testing.T) {
	e := &AExprArith{
		e1: &IntegerLiteral{value: 1},
		e2: &AExprArith{
			e1: &IntegerLiteral{value: 2},
			e2: &IntegerLiteral{value: 3},
			op: AExprArithAdd,
		},
		op: AExprArithAdd,
	}

	fmt.Println(e.String())
}