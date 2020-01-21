package ast

import "fmt"

type IfExpr struct {
	e        BExpr
	stmtList []FuncStatementer
}

func (e *IfExpr) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sIfExpression:\n"+
		"%s..BinaryExpression:\n"+
		"%s\n"+
		"%s..IfBody:\n"+
		"%s",
		ident, ident, e.e.toString(ident+"...."),
		ident, IterableToString(ident+"....", IteratableFuncStatementerList(e.stmtList)))
}

func (e *IfExpr) String() string {
	return e.toString("")
}

type ElseExpr struct {
	stmtList []FuncStatementer
}

func (e *ElseExpr) toString(ident string) string {
	if e == nil {
		return fmt.Sprintf(""+
			"%sElseExpression:\n"+
			"%s..ElseBody:\n"+
			"%s",
			ident, ident, IterableToString(ident+"....", nil))
	}

	return fmt.Sprintf(""+
		"%sElseExpression:\n"+
		"%s..ElseBody:\n"+
		"%s",
		ident, ident, IterableToString(ident+"....", IteratableFuncStatementerList(e.stmtList)))
}

func (e *ElseExpr) String() string {
	return e.toString("")
}

type SelectionStmt struct {
	ifExprList []*IfExpr
	elseExpr   *ElseExpr
}

func (s *SelectionStmt) statementer() {}

func (s *SelectionStmt) funcStatementer() {}

func (s *SelectionStmt) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sSelectionExpression:\n"+
		"%s..IfExpressionList:\n"+
		"%s\n"+
		"%s..ElseExpression:\n"+
		"%s",
		ident, ident, IterableToString(ident+"....", IteratableIfExprList(s.ifExprList)),
		ident, s.elseExpr.toString(ident+"...."))
}

func (s *SelectionStmt) String() string {
	return s.toString("")
}
