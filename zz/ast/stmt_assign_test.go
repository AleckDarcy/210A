package ast

import (
	"fmt"
	"testing"
)

func TestAssignStmt_String(t *testing.T) {
	s := AssignStmt1
	s.funcStatementer()
	s.definitioner()
	fmt.Println(s.String())
}
