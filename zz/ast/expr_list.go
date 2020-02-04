package ast

import "fmt"

type ListElementIndex struct {
	e AExpr
}

var ListElementIndexHelper *ListElementIndex

func (i *ListElementIndex) New(e AExpr) *ListElementIndex {
	return &ListElementIndex{e: e}
}

func (i *ListElementIndex) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sListElementIndex {\n"+
		"%s\n"+
		"%s}",
		ident, i.e.toString(ident+".."),
		ident,
	)
}

func (i *ListElementIndex) String() string {
	return i.toString("")
}

func (i *ListElementIndex) E() AExpr {
	return i.e
}

type ListElementExpr struct {
	name *Identifier
	list []*ListElementIndex
}

var ListElementExprHelper *ListElementExpr

func (e *ListElementExpr) New(name *Identifier, list []*ListElementIndex) *ListElementExpr {
	return &ListElementExpr{
		name: name,
		list: list,
	}
}

func (e *ListElementExpr) aExpr() {}

func (e *ListElementExpr) declaratorer() {}

func (e *ListElementExpr) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sListElementExpression {\n"+
		"%s..Name: %s\n"+
		"%s..List:\n"+
		"%s\n"+
		"%s}",
		ident, ident, e.name,
		ident, IterableToString(ident+"....", IteratableListElementIndexList(e.list)),
		ident,
	)
}

func (e *ListElementExpr) String() string {
	return e.toString("")
}

func (e *ListElementExpr) Name() *Identifier {
	return e.name
}

func (e *ListElementExpr) List() []*ListElementIndex {
	return e.list
}

func (e *ListElementExpr) Identifier() *Identifier {
	return e.name
}

type TupleSizeList struct {
	list []AExpr
}

type ListInitExpr struct {
	typeSpecifier *ListTypeSpecifier
	size          AExpr
}

var ListInitExprHelper *ListInitExpr

func (e *ListInitExpr) New(typeSpecifier *ListTypeSpecifier, size AExpr) *ListInitExpr {
	return &ListInitExpr{typeSpecifier: typeSpecifier, size: size}
}

func (e *ListInitExpr) assignIniter() {}

func (e *ListInitExpr) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sListInitExpression {\n"+
		"%s..TypeSepcifier:\n"+
		"%s\n"+
		"%s..Size:\n"+
		"%s\n"+
		"%s}",
		ident, ident, e.typeSpecifier.toString(ident+"...."),
		ident, e.size.toString(ident+"...."),
		ident,
	)
}

func (e *ListInitExpr) String() string {
	return e.toString("")
}

func (e *ListInitExpr) TypeSpecifier() *ListTypeSpecifier {
	return e.typeSpecifier
}

func (e *ListInitExpr) Size() AExpr {
	return e.size
}
