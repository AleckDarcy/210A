package ast

import (
	"fmt"
	"testing"
)

func TestIdentifier_String(t *testing.T) {
	i := &Identifier{name: "test"}
	i.declaratorer()
	fmt.Println(i)
}

func TestIntegerLiteral_String(t *testing.T) {
	l := &IntegerLiteral{value: 2}
	l.aExpr()
	l.Value()
	fmt.Println(l)
}
