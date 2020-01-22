package ast

import "fmt"

type ListElementIndex struct {
	e AExpr
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

type ListElementExpr struct {
	name *Identifier
	list []*ListElementIndex
}

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

type TupleSizeList struct {
	list []AExpr
}

type ListInitExpr struct {
	typeSpecifier *ListTypeSpecifier
	size          AExpr
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
