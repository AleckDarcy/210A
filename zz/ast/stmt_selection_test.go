package ast

import (
	"fmt"
	"testing"
)

var SelectionStmt2 = &SelectionStmt{
	ifExprList: []*IfExpr{
		IFExpr1,
		IFExpr2,
	},
	elseExpr: nil,
}

func TestIfExpr_String(t *testing.T) {
	e := IFExpr1
	fmt.Println(e)
}

func TestElseExpr_String(t *testing.T) {
	e := ElseExpr1
	fmt.Println(e)
}

func TestSelectionStmt_String(t *testing.T) {
	s := SelectionStmt1
	s.funcStatementer()
	fmt.Println(s)

	s = SelectionStmt2
	fmt.Println(s)
}
