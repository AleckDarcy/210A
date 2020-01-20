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

// Statementer is an interface of statement
type Statementer interface {
	BasicNoder

	statementer()
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

type FuncStatementer interface {
	BasicNoder

	funcStatementer()
}
