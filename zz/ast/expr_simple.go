package ast

import "fmt"

type Declarator struct {
	Declaratorer
}

var DeclaratorHelper *Declarator

func (d *Declarator) New(decl Declaratorer) *Declarator {
	return &Declarator{Declaratorer: decl}
}

type Identifier struct {
	name string
}

var IdentifierHelper *Identifier

func (l *Identifier) New(name string) *Identifier {
	return &Identifier{name: name}
}

func (l *Identifier) aExpr() {}

func (i *Identifier) declaratorer() {}

func (i *Identifier) funcReturnParaer() {}

func (i *Identifier) toString(ident string) string {
	return fmt.Sprintf("%sIdentifier(%s)", ident, i.name)
}

func (i *Identifier) String() string {
	return i.toString("")
}

func (i *Identifier) Identifier() *Identifier {
	return i
}

func (i *Identifier) Name() string {
	return i.name
}

type IntegerLiteral struct {
	value int64
}

var IntegerLiteralHelper *IntegerLiteral

func (l *IntegerLiteral) New(value int64) *IntegerLiteral {
	return &IntegerLiteral{value: value}
}

func (l *IntegerLiteral) aExpr() {}

func (l *IntegerLiteral) assignIniter() {}

func (l *IntegerLiteral) toString(ident string) string {
	return fmt.Sprintf("%sInt(%d)", ident, l.value)
}

func (l *IntegerLiteral) String() string {
	return l.toString("")
}

func (l *IntegerLiteral) Value() int64 {
	return l.value
}

type FloatLiteral struct {
	value float64
}

var FloatLiteralHelper *FloatLiteral

func (l *FloatLiteral) New(value float64) *FloatLiteral {
	return &FloatLiteral{value: value}
}

func (l *FloatLiteral) aExpr() {}

func (l *FloatLiteral) assignIniter() {}

func (l *FloatLiteral) toString(ident string) string {
	return fmt.Sprintf("%sFloat(%f)", ident, l.value)
}

func (l *FloatLiteral) String() string {
	return l.toString("")
}

func (l *FloatLiteral) Value() float64 {
	return l.value
}

type BinaryLiteral struct {
	value bool
}

var BinaryLiteralHelper *BinaryLiteral

func (l *BinaryLiteral) New(value bool) *BinaryLiteral {
	return &BinaryLiteral{value: value}
}

func (l *BinaryLiteral) bExpr() {}

func (l *BinaryLiteral) toString(ident string) string {
	return fmt.Sprintf("%sBool(%v)", ident, l.value)
}

func (l *BinaryLiteral) String() string {
	return l.toString("")
}

func (l *BinaryLiteral) Value() bool {
	return l.value
}
