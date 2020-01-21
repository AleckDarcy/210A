package ast

import "fmt"

type AssignStmt struct {
	declList []Declaratorer
	initList []AssignIniter
}

func (s *AssignStmt) statementer() {}

func (s *AssignStmt) funcStatementer() {}

func (s *AssignStmt) toString(ident string) string {
	initList := ""

	for i, init := range s.initList {
		if i == len(s.initList) - 1 {
			initList += fmt.Sprintf(""+
				"%s....[%d]:\n"+
				"%s", ident, i, init.toString(ident+"......"))
		} else {
			initList += fmt.Sprintf(""+
				"%s....[%d]:\n"+
				"%s\n", ident, i, init.toString(ident+"......"))
		}
	}

	return fmt.Sprintf("" +
		"%sAssignStatement\n" +
		"%s..DeclaratorList:\n" +
		"%s\n" +
		"%s..InitializerList:\n" +
		"%s", ident, ident,
		IterableToString(ident+"....", IteratableDeclaratorerList(s.declList)), ident, initList)
}

func (s *AssignStmt) String() string {
	return s.toString("")
}
