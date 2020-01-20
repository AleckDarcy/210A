package ast

type BasicNoder interface {
	toString(ident string) string
	String() string
}

type AExpr interface {
	BasicNoder

	aExpr()
}

