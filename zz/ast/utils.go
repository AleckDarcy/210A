package ast

import "fmt"

type Iteraterable interface {
	Len() int
	Get(i int) BasicNoder
}

type IteratableAExprList []AExpr

func (it IteratableAExprList) Len() int {
	return len(it)
}

func (it IteratableAExprList) Get(i int) BasicNoder {
	return it[i]
}

type IteratableBasicNoderList []BasicNoder

func (it IteratableBasicNoderList) Len() int {
	return len(it)
}

func (it IteratableBasicNoderList) Get(i int) BasicNoder {
	return it[i]
}

type IteratableAssignIniterList []AssignIniter

func (it IteratableAssignIniterList) Len() int {
	return len(it)
}

func (it IteratableAssignIniterList) Get(i int) BasicNoder {
	return it[i]
}

type IteratableCollectionElementIndexList []*CollectionElementIndex

func (it IteratableCollectionElementIndexList) Len() int {
	return len(it)
}

func (it IteratableCollectionElementIndexList) Get(i int) BasicNoder {
	return it[i]
}

type IteratableTypeSpecifierList []TypeSpecifierer

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

type IteratableFuncStatementerList []FuncStatementer

func (it IteratableFuncStatementerList) Len() int {
	return len(it)
}

func (it IteratableFuncStatementerList) Get(i int) BasicNoder {
	return it[i]
}

type IteratableFuncReturnParaList []FuncReturnParaer

func (it IteratableFuncReturnParaList) Len() int {
	return len(it)
}

func (it IteratableFuncReturnParaList) Get(i int) BasicNoder {
	return it[i]
}

type IteratableFuncExecuteParaerList []FuncExecuteParaer

func (it IteratableFuncExecuteParaerList) Len() int {
	return len(it)
}

func (it IteratableFuncExecuteParaerList) Get(i int) BasicNoder {
	return it[i]
}

type IteratableIfExprList []*IfExpr

func (it IteratableIfExprList) Len() int {
	return len(it)
}

func (it IteratableIfExprList) Get(i int) BasicNoder {
	return it[i]
}

type IteratableDefinitioner []Definitioner

func (it IteratableDefinitioner) Len() int {
	return len(it)
}

func (it IteratableDefinitioner) Get(i int) BasicNoder {
	return it[i]
}

func IterableToString(ident string, list Iteraterable) string {
	if list == nil {
		return ident + "Null"
	} else if list.Len() == 0 {
		return ident + "..Empty"
	}

	result := ""
	for i := 0; i < list.Len(); i++ {
		elem := list.Get(i)
		if i == list.Len()-1 {
			result += fmt.Sprintf(""+
				"%s[%d]:\n"+
				"%s",
				ident, i, elem.toString(ident+".."),
			)
		} else {
			result += fmt.Sprintf(""+
				"%s[%d]:\n"+
				"%s\n",
				ident, i, elem.toString(ident+".."),
			)
		}
	}

	return result
}
