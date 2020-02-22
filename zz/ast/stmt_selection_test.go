package ast

import (
	"fmt"
	"testing"
)

func TestIfExpr_String(t *testing.T) {
	e := IFExpr1
	e.BinExpr()
	e.StmtList()
	fmt.Println(e)
}

func TestElseExpr_String(t *testing.T) {
	e := ElseExpr1
	e.StmtList()
	fmt.Println(e)
}

func TestSelectionStmt_String(t *testing.T) {
	s := SelectionStmt1
	s.funcStatementer()
	s.IfExprList()
	s.ElseExpr()
	fmt.Println(s)

	fmt.Println(SelectionStmt2)
	fmt.Println(SelectionStmt3)
}
