package ast

import (
	"fmt"
	"testing"
)

var IFExpr1 = &IfExpr{
	e:        bExprBinary1,
	stmtList: FuncStmtList1,
}

var IFExpr2 = &IfExpr{
	e:        bExprBinary2,
	stmtList: nil,
}

var ElseExpr1 = &ElseExpr{
	stmtList: FuncStmtList1,
}

var SelectionStmt1 = &SelectionStmt{
	ifExprList: []*IfExpr{
		IFExpr1,
		IFExpr2,
	},
	elseExpr: ElseExpr1,
}

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
	s.statementer()
	s.funcStatementer()
	fmt.Println(s)

	s = SelectionStmt2
	fmt.Println(s)
}
