package ast

type NoderType int64

const (
	NoderBasic NoderType = iota
	NoderAExpr
	NoderBExpr
	NoderTypeSpecifierer
	NoderListElementTypeSpecifierer
	NoderDeclarator
	NoderAssignIniter
	NoderFuncStatementer
	NoderFuncReturnPara
	NoderDefinitioner

	NoderIdentifier

	NoderAssignStatment

	NoderListElementIndex
	NoderListElementTypeSpecifier
	NoderListTypeSpecifier

	NoderIfExpr
	NoderElseExpr

	NoderIterationAssignStatement
	NoderParaDeclarator
	NoderParaDeclaratorWithIdentity
	NoderFuncIdentifier
	NoderFuncTypeSpecifier
	NoderFuncTypeSpecifierWithName
)

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

// TypeSpecifierer is an interface of typeSpecifier
type TypeSpecifierer interface {
	BasicNoder

	typeSpecifierer()
}

type ListElementTypeSpecifierer interface {
	BasicNoder

	listElementTypeSpecifierer()
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

// FuncReturnParaer is an interface of funcReturnPara
type FuncReturnParaer interface {
	BasicNoder

	funcReturnParaer()
}

// FuncExecuteParaer is an interface of funcExecutePara
type FuncExecuteParaer interface {
	BasicNoder

	funcExecuteParaer()
}

// Definitioner is an interface of definition
type Definitioner interface {
	BasicNoder

	definitioner()
}
