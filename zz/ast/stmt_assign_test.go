package ast

import (
	"fmt"
	"testing"
)

//
//var AssignStmt1 = &AssignStmt{
//	declList: []Declaratorer{
//		&Identifier{name: "a"},
//		ListElementExpr1,
//		&Identifier{name: "c"},
//	},
//	initList: []AssignIniter{
//		aExprAdd1,
//		ListInitExpr1, // list([][]int, 2+3)
//		ListInitExpr2,
//	},
//}

func TestAssignStmt_String(t *testing.T) {
	s := AssignStmt3
	s.funcStatementer()
	s.definitioner()
	s.Flag()
	s.DeclList()
	s.InitList()
	s.AddFlag(AssignStmtFlagNormal)
	fmt.Println(s)
	fmt.Println((*AssignStmt)(nil))
}

func TestAssignStmtFlag_String(t *testing.T) {
	fmt.Println(AssignStmtFlagNormal)
	fmt.Println(AssignStmtFlagInit)
	fmt.Println(AssignStmtFlagGlobal)
	fmt.Println(AssignStmtFlagMatrix)

}
