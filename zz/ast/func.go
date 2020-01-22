package ast

import "fmt"

type ParaDeclaratorWithIdentity struct {
	declList      []*Identifier
	typeSpecifier TypeSpecifier
}

func (s *ParaDeclaratorWithIdentity) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sParaDeclaratorWithIdentity {\n"+
		"%s..DeclaratorList:\n"+
		"%s\n"+
		"%s..TypeSpecifier:\n"+
		"%s\n"+
		"%s}",
		ident, ident, IterableToString(ident+"....", IteratableIdentifierList(s.declList)),
		ident, s.typeSpecifier.toString(ident+"...."),
		ident,
	)
}

func (s *ParaDeclaratorWithIdentity) String() string {
	return s.toString("")
}

type FuncTypeSpecifier struct {
	paraList          []*ParaDeclaratorWithIdentity
	typeSpecifierList []TypeSpecifier
}

func (s *FuncTypeSpecifier) typeSpecifier() {}

func (s *FuncTypeSpecifier) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sFuncTypeSpecifier {\n"+
		"%s..ParaDeclaratorWithIdentityList:\n"+
		"%s\n"+
		"%s..TypeSpecifierList:\n"+
		"%s\n"+
		"%s}",
		ident, ident, IterableToString(ident+"....", IteratableParaDeclaratorWithIdentityList(s.paraList)),
		ident, IterableToString(ident+"....", IteratableTypeSpecifierList(s.typeSpecifierList)),
		ident,
	)
}

func (s *FuncTypeSpecifier) String() string {
	return s.toString("")
}

type FuncInitExpr struct {
	typeSpecifier *FuncTypeSpecifier
	stmtList      []FuncStatementer
}

func (e *FuncInitExpr) assignIniter() {}

func (e *FuncInitExpr) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sFuncInitExpression {\n"+
		"%s..FuncTypeSpecifier:\n"+
		"%s\n"+
		"%s..FuncBody:\n"+
		"%s\n"+
		"%s}",
		ident, ident, e.typeSpecifier.toString(ident+"...."),
		ident, IterableToString(ident+"....", IteratableFuncStatementerList(e.stmtList)),
		ident,
	)
}

func (e *FuncInitExpr) String() string {
	return e.toString("")
}

type FuncTypeSpecifierWithName struct {
	name              *Identifier
	paraList          []*ParaDeclaratorWithIdentity
	typeSpecifierList []TypeSpecifier
}

func (s *FuncTypeSpecifierWithName) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sFuncTypeSpecifierWithName {\n"+
		"%s..Name:\n"+
		"%s\n"+
		"%s..ParaDeclaratorWithIdentityList:\n"+
		"%s\n"+
		"%s..TypeSpecifierList:\n"+
		"%s\n"+
		"%s}",
		ident, ident, s.name.toString(ident+"...."),
		ident, IterableToString(ident+"....", IteratableParaDeclaratorWithIdentityList(s.paraList)),
		ident, IterableToString(ident+"....", IteratableTypeSpecifierList(s.typeSpecifierList)),
		ident,
	)
}

func (s *FuncTypeSpecifierWithName) String() string {
	return s.toString("")
}

type FuncDefinition struct {
	typeSpecifier *FuncTypeSpecifierWithName
	stmtList      []FuncStatementer
}

func (e *FuncDefinition) definitioner() {}

func (e *FuncDefinition) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sFuncDefinition {\n"+
		"%s..FuncTypeSpecifierWithName:\n"+
		"%s\n"+
		"%s..FuncBody:\n"+
		"%s\n"+
		"%s}",
		ident, ident, e.typeSpecifier.toString(ident+"...."),
		ident, IterableToString(ident+"....", IteratableFuncStatementerList(e.stmtList)),
		ident,
	)
}

func (e *FuncDefinition) String() string {
	return e.toString("")
}
