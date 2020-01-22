package ast

import (
	"fmt"
	"testing"
)

var IterationStmt1 = &IterationStmt{
	initStmt: AssignStmt1,
	binExpr:  bExprBinary2,
	stmtList: FuncStmtList1,
}

func TestIterationStmt_String(t *testing.T) {
	s := IterationStmt1
	s.funcStatementer()
	fmt.Println(s)
}
