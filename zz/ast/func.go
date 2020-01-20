package ast

import "fmt"

type ParaDeclaratorWithIdentity struct {
	declList      []*Identifier
	typeSpecifier TypeSpecifier
}

func (s *ParaDeclaratorWithIdentity) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sParaDeclaratorWithIdentity\n"+
		"%s..DeclaratorList:\n"+
		"%s\n"+
		"%s..TypeSpecifier:\n"+
		"%s",
		ident, ident,
		IterableToString(ident+"....", IteratableIdentifierList(s.declList)),
		ident, s.typeSpecifier.toString(ident+"...."))
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
		"%sFuncTypeSpecifier:\n"+
		"%s..ParaDeclaratorWithIdentityList:\n"+
		"%s\n"+
		"%s..TypeSpecifierList:\n"+
		"%s",
		ident, ident,
		IterableToString(ident+"....", IteratableParaDeclaratorWithIdentityList(s.paraList)),
		ident,
		IterableToString(ident+"....", IteratableTypeSpecifierList(s.typeSpecifierList)),
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
		"%sFuncInitExpr\n"+
		"%s..FuncTypeSpecifier:\n"+
		"%s\n"+
		"%s..FuncBody:\n"+
		"%s",
		ident, ident, e.typeSpecifier.toString(ident+".."),
		ident, IterableToString(ident+"....", IteratableFuncStatementer(e.stmtList)),
	)
}

func (e *FuncInitExpr) String() string {
	return e.toString("")
}
