package ast

import "fmt"

type SimpleTypeSpecifier struct {
	name string
}

var SimpleTypeSpecifierHelper *SimpleTypeSpecifier

func (s *SimpleTypeSpecifier) New(name string) *SimpleTypeSpecifier {
	return &SimpleTypeSpecifier{name: name}
}

func (s *SimpleTypeSpecifier) typeSpecifierer() {}

func (s *SimpleTypeSpecifier) listElementTypeSpecifierer() {}

func (s *SimpleTypeSpecifier) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sSimpleTypeSpecifier:\n"+
		"%s..%s", ident, ident, s.name)
}

func (s *SimpleTypeSpecifier) String() string {
	return s.name
}

func (s *SimpleTypeSpecifier) Name() string {
	return s.name
}

type ListElementTypeSpecifierType int64

const (
	ListElementTypeSpecifierSimple ListElementTypeSpecifierType = iota
	ListElementTypeSpecifierNested
)

func (t ListElementTypeSpecifierType) toString(ident string) string {
	switch t {
	case ListElementTypeSpecifierSimple:
		return ident + "Simple"
	case ListElementTypeSpecifierNested:
		return ident + "Nested"
	default:
		return ident + "undefined"
	}
}

func (t ListElementTypeSpecifierType) String() string {
	return t.toString("")
}

type ListElementTypeSpecifier struct {
	elem ListElementTypeSpecifierer
	typ  ListElementTypeSpecifierType
}

var ListElementTypeSpecifierHelper *ListElementTypeSpecifier

func (s *ListElementTypeSpecifier) New(elem ListElementTypeSpecifierer, typ ListElementTypeSpecifierType) *ListElementTypeSpecifier {
	return &ListElementTypeSpecifier{elem: elem, typ: typ}
}

func (s *ListElementTypeSpecifier) listElementTypeSpecifierer() {}

func (s *ListElementTypeSpecifier) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sListElementTypeSpecifier:\n"+
		"%s", ident, s.elem.toString(ident+".."))
}

func (s *ListElementTypeSpecifier) String() string {
	return s.toString("")
}

func (s *ListElementTypeSpecifier) Elem() ListElementTypeSpecifierer {
	return s.elem
}

type ListTypeSpecifier struct {
	elem *ListElementTypeSpecifier
}

var ListTypeSpecifierHelper *ListTypeSpecifier

func (s *ListTypeSpecifier) New(elem *ListElementTypeSpecifier) *ListTypeSpecifier {
	return &ListTypeSpecifier{elem: elem}
}

func (s *ListTypeSpecifier) typeSpecifierer() {}

func (s *ListTypeSpecifier) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sListTypeSpecifier {\n"+
		"%s..Element:\n"+
		"%s\n"+
		"%s}",
		ident, ident, s.elem.toString(ident+"...."),
		ident,
	)
}

func (s *ListTypeSpecifier) Elem() *ListElementTypeSpecifier {
	return s.elem
}

func (s *ListTypeSpecifier) String() string {
	return s.toString("")
}
