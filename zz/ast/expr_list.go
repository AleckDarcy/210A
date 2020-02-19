package ast

import "fmt"

type CollectionType int64

const (
	CollectionList CollectionType = iota
	CollectionMatrix
)

func (t CollectionType) toString(indent string) string {
	switch t {
	case CollectionList:
		return indent + "List"
	case CollectionMatrix:
		return indent + "Matrix"
	default:
		return indent + "undefined"
	}
}

func (t CollectionType) String() string {
	return t.toString("")
}

type CollectionElementIndex struct {
	e AExpr
}

var CollectionElementIndexHelper *CollectionElementIndex

func (i *CollectionElementIndex) New(e AExpr) *CollectionElementIndex {
	return &CollectionElementIndex{e: e}
}

func (i *CollectionElementIndex) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sCollectionElementIndex {\n"+
		"%s\n"+
		"%s}",
		ident, i.e.toString(ident+".."),
		ident,
	)
}

func (i *CollectionElementIndex) String() string {
	return i.toString("")
}

func (i *CollectionElementIndex) E() AExpr {
	return i.e
}

type CollectionElementExpr struct {
	typ  CollectionType
	name *Identifier
	list []*CollectionElementIndex
}

var CollectionElementExprHelper *CollectionElementExpr

func (e *CollectionElementExpr) New(name *Identifier, list []*CollectionElementIndex) *CollectionElementExpr {
	return &CollectionElementExpr{name: name, list: list}
}

func (e *CollectionElementExpr) aExpr() {}

func (e *CollectionElementExpr) declaratorer() {}

func (e *CollectionElementExpr) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sCollectionElementExpression {\n"+
		"%s..Name: %s\n"+
		"%s..List:\n"+
		"%s\n"+
		"%s}",
		ident, ident, e.name,
		ident, IterableToString(ident+"....", IteratableCollectionElementIndexList(e.list)),
		ident,
	)
}

func (e *CollectionElementExpr) String() string {
	return e.toString("")
}

func (e *CollectionElementExpr) Type() CollectionType {
	return e.typ
}

func (e *CollectionElementExpr) Name() *Identifier {
	return e.name
}

func (e *CollectionElementExpr) List() []*CollectionElementIndex {
	return e.list
}

func (e *CollectionElementExpr) Identifier() *Identifier {
	return e.name
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

type MatrixInitExpr struct {
	sizes []AExpr
}

var MatrixInitExprHelper *MatrixInitExpr

func (e *MatrixInitExpr) New(sizes []AExpr) *MatrixInitExpr {
	return &MatrixInitExpr{sizes: sizes}
}

func (e *MatrixInitExpr) assignIniter() {}

func (e *MatrixInitExpr) toString(indent string) string {
	return fmt.Sprintf(""+
		"%sMatrixInitExpression {\n"+
		"%s..Size:\n"+
		"%s\n"+
		"%s}",
		indent, indent, IterableToString(indent+"....", IteratableAExprList(e.sizes)),
		indent,
	)
}

func (e *MatrixInitExpr) String() string {
	return e.toString("")
}

func (e *MatrixInitExpr) Sizes() []AExpr {
	return e.sizes
}
