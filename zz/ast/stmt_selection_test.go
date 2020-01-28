package ast

import (
	"fmt"
	"testing"
)

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

	fmt.Println(SelectionStmt2)
	fmt.Println(SelectionStmt3)
}
