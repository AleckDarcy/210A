package ast

import (
	"fmt"
	"testing"
)

func TestIterationStmt_String(t *testing.T) {
	s := IterationStmt1
	s.funcStatementer()
	fmt.Println(s)
}

func TestIterationAssignStmt_String(t *testing.T) {
	s := IterationAssignStmt1
	fmt.Println(s)
}
