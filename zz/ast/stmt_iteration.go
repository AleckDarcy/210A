package ast

import "fmt"

type IterationStmt struct {
	initStmt  *AssignStmt
	binExpr   BExpr
	increStmt *AssignStmt
	stmtList  []FuncStatementer
}

var IterationStmtHelper *IterationStmt

func (s *IterationStmt) New(initStmt *AssignStmt, binExpr BExpr, increStmt *AssignStmt, stmtList []FuncStatementer) *IterationStmt {
	return &IterationStmt{initStmt: initStmt, binExpr: binExpr, increStmt: increStmt, stmtList: stmtList}
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

type IterationAssignStmt struct {
	stmt *AssignStmt
}

var IterationAssignStmtHelper *IterationAssignStmt

func (s *IterationAssignStmt) New(stmt *AssignStmt) *IterationAssignStmt {
	return &IterationAssignStmt{stmt: stmt}
}

func (s *IterationAssignStmt) toString(indent string) string {
	return fmt.Sprintf(""+
		"%sIterationAssignStatement {\n"+
		"%s..AssignStatement:\n"+
		"%s\n"+
		"%s}",
		indent, indent, s.stmt.toString(indent+"...."),
		indent,
	)
}

func (s *IterationAssignStmt) String() string {
	return s.toString("")
}

func (s *IterationAssignStmt) AssignStmt() *AssignStmt {
	return s.stmt
}
