package ast

type BasicNoder interface {
	toString(ident string) string
	String() string
}

// AExpr is an interface of aExp
type AExpr interface {
	BasicNoder

	aExpr()
}

// BExpr is an interface of bExp
type BExpr interface {
	BasicNoder

	bExpr()
}

// TypeSpecifier is an interface of typeSpecifier
type TypeSpecifier interface {
	BasicNoder

	typeSpecifier()
}

// Declaratorer is an interface of declarator
type Declaratorer interface {
	BasicNoder

	declaratorer()
}

// AssignIniter is an interface of assignInit
type AssignIniter interface {
	BasicNoder

	assignIniter()
}

// FuncStatementer is an interface of functionStatement
type FuncStatementer interface {
	BasicNoder

	funcStatementer()
}

// Definitioner is an interface of definition
type Definitioner interface {
	BasicNoder

	definitioner()
}
