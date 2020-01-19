package ast

import "fmt"

type IntegerLiteral struct {
	value int64
}

func (l *IntegerLiteral) aExpr() {}

func (l *IntegerLiteral) toString(ident string) string {
	return fmt.Sprintf("%sInt(%d)", ident, l.value)
}

func (l *IntegerLiteral) String() string {
	return l.toString("")
}

func (l *IntegerLiteral) Value() int64 {
	return l.value
}

type AExprArithOpType int64

const (
	AExprArithAdd AExprArithOpType = iota
)

func (t AExprArithOpType) toString(ident string) string {
	switch t {
	case AExprArithAdd:
		return ident + "Add"
	default:
		return ident + "undefined"
	}
}

func (t AExprArithOpType) String() string {
	switch t {
	case AExprArithAdd:
		return "Add"
	default:
		return "undefined"
	}
}

type AExprArith struct {
	e1, e2 AExpr
	op     AExprArithOpType
}

func (e *AExprArith) aExpr() {}

func (e *AExprArith) toString(ident string) string {
	return fmt.Sprintf("" +
		"%sAExprArith\n" +
		"%s..E1:\n" +
		"%s\n" +
		"%s..E2:\n" +
		"%s\n" +
		"%s..Op: %s",

		ident, ident, e.e1.toString("...."+ident),
		ident, e.e2.toString("...."+ident),
		ident, e.op)
}

func (e *AExprArith) String() string {
	return e.toString("")
}

func (e *AExprArith) E1() AExpr {
	return e.e1
}

func (e *AExprArith) E2() AExpr {
	return e.e2
}

func (e *AExprArith) Op() AExprArithOpType {
	return e.op
}

