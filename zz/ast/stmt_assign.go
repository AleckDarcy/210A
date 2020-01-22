package ast

import "fmt"

type AssignStmt struct {
	declList []Declaratorer
	initList []AssignIniter
}

func (s *AssignStmt) funcStatementer() {}

func (s *AssignStmt) definitioner() {}

func (s *AssignStmt) toString(ident string) string {
	if s == nil {
		return ident + "Null"
	}

	return fmt.Sprintf(""+
		"%sAssignStatement {\n"+
		"%s..DeclaratorList:\n"+
		"%s\n"+
		"%s..InitializerList:\n"+
		"%s\n"+
		"%s}",
		ident, ident, IterableToString(ident+"....", IteratableDeclaratorerList(s.declList)),
		ident, IterableToString(ident+"....", IteratableAssignIniterList(s.initList)),
		ident,
	)
}

func (s *AssignStmt) String() string {
	return s.toString("")
}
