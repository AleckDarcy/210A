package ast

import "fmt"

type SimpleTypeSpecifier struct {
	name string
}

func (s *SimpleTypeSpecifier) typeSpecifier() {}

func (s *SimpleTypeSpecifier) toString(ident string) string {
	return fmt.Sprintf("" +
		"%sSimpleTypeSpecifier:\n" +
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
	elem TypeSpecifier
	typ  ListElementTypeSpecifierType
}

func (s *ListElementTypeSpecifier) typeSpecifier() {}

func (s *ListElementTypeSpecifier) toString(ident string) string {
	return fmt.Sprintf("" +
		"%sListElementTypeSpecifier:\n" +
		"%s", ident,  s.elem.toString(ident + ".."))
}

func (s *ListElementTypeSpecifier) String() string {
	return s.toString("")
}

type ListTypeSpecifier struct {
	elem *ListElementTypeSpecifier
}

func (s *ListTypeSpecifier) typeSpecifier() {}

func (s *ListTypeSpecifier) toString(ident string) string {
	return fmt.Sprintf("" +
		"%sListTypeSpecifier:\n" +
		"%s..Element:\n" +
		"%s",
		ident, ident, s.elem.toString(ident + "...."))
}

func (s *ListTypeSpecifier) String() string {
	return s.toString("")
}
