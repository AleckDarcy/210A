package ast

import (
	"fmt"
	"testing"
)

func TestIdentifier_String(t *testing.T) {
	i := &Identifier{name: "test"}
	i.aExpr()
	i.declaratorer()
	i.funcReturnParaer()
	fmt.Println(i)
}

func TestIntegerLiteral_String(t *testing.T) {
	l := &IntegerLiteral{value: 2}
	l.aExpr()
	l.assignIniter()
	l.Value()
	fmt.Println(l)
}

func TestFloatLiteral_String(t *testing.T) {
	l := &FloatLiteral{value: 2}
	l.aExpr()
	l.assignIniter()
	fmt.Println(l)
}

func TestBinaryLiteral_String(t *testing.T) {
	l := &BinaryLiteral{value: true}
	l.bExpr()
	fmt.Println(l)
}
