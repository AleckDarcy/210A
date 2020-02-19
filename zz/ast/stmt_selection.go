package ast

import "fmt"

type IfExpr struct {
	binExpr  BExpr
	stmtList []FuncStatementer
}

var IfExprHelper *IfExpr

func (e *IfExpr) New(binExpr BExpr, stmtList []FuncStatementer) *IfExpr {
	return &IfExpr{binExpr: binExpr, stmtList: stmtList}
}

func (e *IfExpr) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sIfExpression {\n"+
		"%s..BinaryExpression:\n"+
		"%s\n"+
		"%s..IfBody:\n"+
		"%s\n"+
		"%s}",
		ident, ident, e.binExpr.toString(ident+"...."),
		ident, IterableToString(ident+"....", IteratableFuncStatementerList(e.stmtList)),
		ident,
	)
}

func (e *IfExpr) String() string {
	return e.toString("")
}

func (e *IfExpr) BinExpr() BExpr {
	return e.binExpr
}

func (e *IfExpr) StmtList() []FuncStatementer {
	return e.stmtList
}

type ElseExpr struct {
	stmtList []FuncStatementer
}

var ElseExprHelper *ElseExpr

func (e *ElseExpr) New(stmtList []FuncStatementer) *ElseExpr {
	return &ElseExpr{stmtList: stmtList}
}

func (e *ElseExpr) toString(ident string) string {
	if e == nil {
		return fmt.Sprintf(""+
			"%sElseExpression {\n"+
			"%s..ElseBody:\n"+
			"%s\n"+
			"%s}",
			ident, ident, IterableToString(ident+"....", nil),
			ident)
	}

	return fmt.Sprintf(""+
		"%sElseExpression {\n"+
		"%s..ElseBody:\n"+
		"%s\n"+
		"%s}",
		ident, ident, IterableToString(ident+"....", IteratableFuncStatementerList(e.stmtList)),
		ident,
	)
}

func (e *ElseExpr) String() string {
	return e.toString("")
}

func (e *ElseExpr) StmtList() []FuncStatementer {
	return e.stmtList
}

type SelectionStmt struct {
	ifExprList []*IfExpr
	elseExpr   *ElseExpr
}

var SelectionStmtHelper *SelectionStmt

func (s *SelectionStmt) New(ifExprList []*IfExpr, elseExpr *ElseExpr) *SelectionStmt {
	return &SelectionStmt{ifExprList: ifExprList, elseExpr: elseExpr}
}

func (s *SelectionStmt) funcStatementer() {}

func (s *SelectionStmt) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sSelectionExpression {\n"+
		"%s..IfExpressionList:\n"+
		"%s\n"+
		"%s..ElseExpression:\n"+
		"%s\n"+
		"%s}",
		ident, ident, IterableToString(ident+"....", IteratableIfExprList(s.ifExprList)),
		ident, s.elseExpr.toString(ident+"...."),
		ident,
	)
}

func (s *SelectionStmt) String() string {
	return s.toString("")
}

func (s *SelectionStmt) IfExprList() []*IfExpr {
	return s.ifExprList
}

func (s *SelectionStmt) ElseExpr() *ElseExpr {
	return s.elseExpr
}
