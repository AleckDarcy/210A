package ast

import (
	"fmt"
	"testing"
)

var AssignStmt1 = &AssignStmt{
	declList: []Declaratorer{
		&Identifier{name: "a"},
		ListElementExpr1,
		&Identifier{name: "c"},
	},
	initList: []AssignIniter{
		aExprAdd1,
		ListInitExpr1, // list([][]int, 2+3)
		ListInitExpr2,
	},
}

func TestAssignStmt_String(t *testing.T) {
	s := AssignStmt1
	s.funcStatementer()
	s.definitioner()
	fmt.Println(s.String())
}
