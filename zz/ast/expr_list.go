package ast

import "fmt"

type ListElementIndex struct {
	e AExpr
}

func (i *ListElementIndex) toString(ident string) string {
	//panic("")
	return fmt.Sprintf(""+
		"%sListElementIndex:\n"+
		"%s",
		ident, i.e.toString(ident+".."))
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
	list := ""
	for i, elem := range e.list {
		if i == len(e.list)-1 {
			list += fmt.Sprintf(""+
				"%s[%d]:\n"+
				"%s", ident+"....", i, elem.toString(ident+"......"))
		} else {
			list += fmt.Sprintf(""+
				"%s[%d]:\n"+
				"%s\n", ident+"....", i, elem.toString(ident+"......"))
		}
	}

	return fmt.Sprintf(""+
		"%sListElementExpression:\n"+
		"%s..Name: %s\n"+
		"%s..List:\n"+
		"%s", ident, ident, e.name.String(), ident, list)
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
	return fmt.Sprintf("" +
		"%sListInitExpression:\n"+
		"%s..TypeSepcifier:\n" +
		"%s\n" +
		"%s..Size:\n"+
		"%s", ident, ident, e.typeSpecifier.toString(ident+"...."),
		ident, e.size.toString(ident+"...."))
}

func (e *ListInitExpr) String() string {
	return e.toString("")
}