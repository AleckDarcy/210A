package ast

import "fmt"

type Iteraterable interface {
	Len() int
	Get(i int) BasicNoder
}

type IteratableTypeSpecifierList []TypeSpecifier

func (it IteratableTypeSpecifierList) Len() int {
	return len(it)
}

func (it IteratableTypeSpecifierList) Get(i int) BasicNoder {
	return it[i]
}

type IteratableIdentifierList []*Identifier

func (it IteratableIdentifierList) Len() int {
	return len(it)
}

func (it IteratableIdentifierList) Get(i int) BasicNoder {
	return it[i]
}

type IteratableDeclaratorerList []Declaratorer

func (it IteratableDeclaratorerList) Len() int {
	return len(it)
}

func (it IteratableDeclaratorerList) Get(i int) BasicNoder {
	return it[i]
}

type IteratableParaDeclaratorWithIdentityList []*ParaDeclaratorWithIdentity

func (it IteratableParaDeclaratorWithIdentityList) Len() int {
	return len(it)
}

func (it IteratableParaDeclaratorWithIdentityList) Get(i int) BasicNoder {
	return it[i]
}

type IteratableFuncStatementer []FuncStatementer

func (it IteratableFuncStatementer) Len() int {
	return len(it)
}

func (it IteratableFuncStatementer) Get(i int) BasicNoder {
	return it[i]
}

func IterableToString(ident string, list Iteraterable) string {
	result := ""
	for i := 0; i < list.Len(); i++ {
		elem := list.Get(i)
		if i == list.Len()-1 {
			result += fmt.Sprintf(""+
				"%s[%d]:\n"+
				"%s", ident, i, elem.toString(ident+".."))
		} else {
			result += fmt.Sprintf(""+
				"%s[%d]:\n"+
				"%s\n", ident, i, elem.toString(ident+".."))
		}
	}

	return result
}
