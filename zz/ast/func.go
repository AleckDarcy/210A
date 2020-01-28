package ast

import "fmt"

type ParaDeclarator struct {
	name *Identifier
}

var ParaDeclaratorHelper *ParaDeclarator

func (d *ParaDeclarator) New(name *Identifier) *ParaDeclarator {
	return &ParaDeclarator{name: name}
}

func (d *ParaDeclarator) toString(indent string) string {
	return fmt.Sprintf(""+
		"%sParaDeclarator {\n"+
		"%s..Name:\n"+
		"%s"+
		"%s}",
		indent, indent, d.name,
		indent,
	)
}

func (d *ParaDeclarator) String() string {
	return d.toString("")
}

func (d *ParaDeclarator) Name() *Identifier {
	return d.name
}

type ParaDeclaratorWithIdentity struct {
	declList      []*Identifier
	typeSpecifier TypeSpecifierer
}

var ParaDeclaratorWithIdentityHelper *ParaDeclaratorWithIdentity

func (s *ParaDeclaratorWithIdentity) New(declList []*Identifier, typeSpecifier TypeSpecifierer) *ParaDeclaratorWithIdentity {
	return &ParaDeclaratorWithIdentity{declList: declList, typeSpecifier: typeSpecifier}
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
	paraList   []*ParaDeclaratorWithIdentity
	returnList []TypeSpecifierer
}

var FuncTypeSpecifierHelper *FuncTypeSpecifier

func (s *FuncTypeSpecifier) New(paraList []*ParaDeclaratorWithIdentity, returnList []TypeSpecifierer) *FuncTypeSpecifier {
	return &FuncTypeSpecifier{paraList: paraList, returnList: returnList}
}

func (s *FuncTypeSpecifier) typeSpecifierer() {}

func (s *FuncTypeSpecifier) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sFuncTypeSpecifier {\n"+
		"%s..ParaDeclaratorWithIdentityList:\n"+
		"%s\n"+
		"%s..TypeSpecifierList:\n"+
		"%s\n"+
		"%s}",
		ident, ident, IterableToString(ident+"....", IteratableParaDeclaratorWithIdentityList(s.paraList)),
		ident, IterableToString(ident+"....", IteratableTypeSpecifierList(s.returnList)),
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

var FuncInitExprHelper *FuncInitExpr

func (e *FuncInitExpr) New(typeSpecifier *FuncTypeSpecifier, stmtList []FuncStatementer) *FuncInitExpr {
	return &FuncInitExpr{typeSpecifier: typeSpecifier, stmtList: stmtList}
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

type FuncIdentifier struct {
	name *Identifier
}

var FuncIdentityHelper *FuncIdentifier

func (i *FuncIdentifier) New(name *Identifier) *FuncIdentifier {
	return &FuncIdentifier{name: name}
}

func (i *FuncIdentifier) toString(indent string) string {
	return fmt.Sprintf(""+
		"%sFuncIdentifier {\n"+
		"%s..Identifier:\n"+
		"%s\n"+
		"%s}",
		indent, indent, i.name.toString(indent+"...."),
		indent,
	)
}

func (i *FuncIdentifier) String() string {
	return i.toString("")
}

func (i *FuncIdentifier) Name() *Identifier {
	return i.name
}

type FuncTypeSpecifierWithName struct {
	name       *Identifier
	paraList   []*ParaDeclaratorWithIdentity
	returnList []TypeSpecifierer
}

var FuncTypeSpecifierWithNameHelper *FuncTypeSpecifierWithName

func (s *FuncTypeSpecifierWithName) New(name *Identifier, paraList []*ParaDeclaratorWithIdentity, returnList []TypeSpecifierer) *FuncTypeSpecifierWithName {
	return &FuncTypeSpecifierWithName{name: name, paraList: paraList, returnList: returnList}
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
		ident, IterableToString(ident+"....", IteratableTypeSpecifierList(s.returnList)),
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

var FuncDefinitionHelper *FuncDefinition

func (e *FuncDefinition) New(typeSpecifier *FuncTypeSpecifierWithName, stmtList []FuncStatementer) *FuncDefinition {
	return &FuncDefinition{typeSpecifier: typeSpecifier, stmtList: stmtList}
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

type FuncReturnStatement struct {
	returnList []FuncReturnParaer
}

var FuncReturnStatementHelper *FuncReturnStatement

func (s *FuncReturnStatement) New(returnList []FuncReturnParaer) *FuncReturnStatement {
	return &FuncReturnStatement{returnList: returnList}
}

func (s *FuncReturnStatement) funcStatementer() {}

func (s *FuncReturnStatement) toString(indent string) string {
	return fmt.Sprintf(""+
		"%sFuncReturnStatement {\n"+
		"%s..ReturnList: \n"+
		"%s\n"+
		"%s}",
		indent, indent, IterableToString(indent+"....", IteratableFuncReturnParaList(s.returnList)),
		indent,
	)
}

func (s *FuncReturnStatement) String() string {
	return s.toString("")
}

type FuncExecutePara struct {
	para FuncExecuteParaer
}

var FuncExecuteParaHelper *FuncExecutePara

func (p *FuncExecutePara) New(para FuncExecuteParaer) *FuncExecutePara {
	return &FuncExecutePara{para: para}
}

func (p *FuncExecutePara) toString(indent string) string {
	return fmt.Sprintf(""+
		"%sFuncExecutePara {\n"+
		"%s..Para:\n"+
		"%s\n"+
		"%s}",
		indent, indent, p.para.toString(indent+"...."),
		indent,
	)
}

func (p *FuncExecutePara) String() string {
	return p.toString("")
}

type FuncExecuteExpression struct {
	name     *Identifier
	paraList []*FuncExecutePara
}

var FuncExecuteExpressionHelper *FuncExecuteExpression

func (e *FuncExecuteExpression) New(name *Identifier, paraList []*FuncExecutePara) *FuncExecuteExpression {
	return &FuncExecuteExpression{name: name, paraList: paraList}
}

func (e *FuncExecuteExpression) assignIniter() {}

func (e *FuncExecuteExpression) toString(indent string) string {
	return fmt.Sprintf(""+
		"%sFuncExecuteExpression {\n"+
		"%s..Name:\n"+
		"%s\n"+
		"%s..ParaList:\n"+
		"%s\n"+
		"%s}",
		indent, indent, e.name.toString(indent+"...."),
		indent, IterableToString(indent+"....", IteratableFuncExecuteParaList(e.paraList)),
		indent,
	)
}

func (e *FuncExecuteExpression) String() string {
	return e.toString("")
}
