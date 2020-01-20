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

type DeclaratorType int64

const (
	DeclaratorIdentifier DeclaratorType = iota
	DeclaratorListElement
)

type Declarator struct {
	elem Declaratorer
	typ DeclaratorType
}

func (d *Declarator) declarator() {}

func (d *Declarator) toString(ident string) string {
	return d.elem.toString(ident)
}

func (d *Declarator) String() string {
	return d.toString("")
}
