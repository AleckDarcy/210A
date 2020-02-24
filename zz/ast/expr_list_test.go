package ast

import (
	"fmt"
	"testing"
)

//
//var ListElementTypeSpecifier1 = &ListElementTypeSpecifier{ // [][]int
//	elem: &ListElementTypeSpecifier{ // []int
//		elem: &SimpleTypeSpecifier{name: "int"},
//		typ:  ListElementTypeSpecifierSimple,
//	},
//	typ: ListElementTypeSpecifierNested,
//}
//
//var ListTypeSpecifier1 = &ListTypeSpecifier{
//	elem: ListElementTypeSpecifier1, // [][]int
//}
//
//var ListElementExpr1 = &ListElementExpr{
//	name: &Identifier{name: "b"},
//	list: []*ListElementIndex{
//		{e: aExprAdd1},
//		{e: &IntegerLiteral{value: 3}},
//	},
//}
//
//var ListInitExpr1 = &ListInitExpr{
//	typeSpecifier: ListTypeSpecifier1,// [][]int
//	size: aExprAdd1, // 2+3
//}
//
//var ListInitExpr2 = &ListInitExpr{
//	typeSpecifier: &ListTypeSpecifier{
//		elem: &ListElementTypeSpecifier{
//			elem: &SimpleTypeSpecifier{name: "int"},
//		},
//	},
//	size: &IntegerLiteral{value: 3},
//}

func TestCollectionType_String(t *testing.T) {
	fmt.Println(CollectionList)
	fmt.Println(CollectionMatrix)
	fmt.Println(CollectionType(-1))
}

func TestCollectionElementIndex_String(t *testing.T) {
	i := CollectionElementIndex1
	i.E()
	fmt.Println(i)
}

//func TestListElementIndex_String(t *testing.T) {
//	i := &CollectionElementIndex{e: AExprAdd1}
//	fmt.Println(i)
//}

func TestListElementExpr_String(t *testing.T) {
	e := ListElementExpr2
	e.aExpr()
	e.declaratorer()
	e.Type()
	e.Name()
	e.List()
	e.Identifier()
	fmt.Println(e)
}

func TestListInitExpr_String(t *testing.T) {
	e := ListInitExpr1
	e.assignIniter()
	e.TypeSpecifier()
	e.Size()
	fmt.Println(e)
}

func TestMatrixInitExpr_String(t *testing.T) {
	e := MatrixInitExpr1
	e.assignIniter()
	e.Sizes()
	fmt.Println(e)
}
