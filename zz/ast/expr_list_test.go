package ast

import (
	"fmt"
	"testing"
)

func TestListElementIndex_String(t *testing.T) {
	i := &ListElementIndex{e: AExprAdd1}
	fmt.Println(i)
}

func TestListElementExpr_String(t *testing.T) {
	e := ListElementExpr2
	e.aExpr()
	e.declaratorer()
	fmt.Println(e)
}

func TestListInitExpr_String(t *testing.T) {
	e := ListInitExpr1
	e.assignIniter()
	fmt.Println(e)
}
