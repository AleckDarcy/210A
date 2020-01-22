package ast

import "fmt"

type IterationStmt struct {
	initStmt  *AssignStmt
	binExpr   BExpr
	increStmt *AssignStmt
	stmtList  []FuncStatementer
}

func (s *IterationStmt) funcStatementer() {}

func (s *IterationStmt) toString(ident string) string {
	return fmt.Sprintf(""+
		"%sIterationStatement {\n"+
		"%s..InitializationStatement:\n"+
		"%s\n"+
		"%s..BinaryExpression:\n"+
		"%s\n"+
		"%s..IncrementStatement:\n"+
		"%s\n"+
		"%s..ForBody:\n"+
		"%s\n"+
		"%s}",
		ident, ident, s.initStmt.toString(ident+"...."),
		ident, s.binExpr.toString(ident+"...."),
		ident, s.increStmt.toString(ident+"...."),
		ident, IterableToString(ident+"....", IteratableFuncStatementerList(s.stmtList)),
		ident,
	)
}

func (s *IterationStmt) String() string {
	return s.toString("")
}
