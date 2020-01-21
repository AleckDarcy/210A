package ast

import "fmt"

type Identifier struct {
	name string
}

func (i *Identifier) declaratorer() {}

func (i *Identifier) toString(ident string) string {
	return fmt.Sprintf("%sIdentifier(%s)", ident, i.name)
}

func (i *Identifier) String() string {
	return i.toString("")
}

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

type BinaryLiteral struct {
	value bool
}

func (l *BinaryLiteral) bExpr() {}

func (l *BinaryLiteral) toString(ident string) string {
	return fmt.Sprintf("%sBool(%v)", ident, l.value)
}

func (l *BinaryLiteral) String() string {
	return l.toString("")
}
