package ast

import (
	"fmt"
	"testing"
)

func TestSimpleTypeSpecifier_String(t *testing.T) {
	s := &SimpleTypeSpecifier{name:"int"}
	s.typeSpecifier()
	s.Name()
	fmt.Println(s)
}

func TestListElementTypeSpecifierType_String(t *testing.T) {
	fmt.Println(ListElementTypeSpecifierSimple)
	fmt.Println(ListElementTypeSpecifierNested)
	fmt.Println(ListElementTypeSpecifierType(-1))
}

func TestListElementTypeSpecifier_String(t *testing.T) {
	s := ListElementTypeSpecifier1
	s.typeSpecifier()
	fmt.Println(s)
}

func TestListTypeSpecifier_String(t *testing.T) {
	s := ListTypeSpecifier1
	s.typeSpecifier()
	fmt.Println(s)
}
