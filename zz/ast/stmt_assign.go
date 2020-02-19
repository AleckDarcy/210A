package ast

import "fmt"

type AssignStmtFlag int64

const (
	AssignStmtFlagNormal AssignStmtFlag = 1 << iota >> 1
	AssignStmtFlagInit
	AssignStmtFlagGlobal
	AssignStmtFlagMatrix
)

func (f AssignStmtFlag) toString(indent string) string {
	if f.IsNormal() {
		return indent + "Normal"
	}

	var result string
	if f.IsInit() {
		result += "Init "
	}
	if f.IsGlobal() {
		result += "Global "
	}

	return indent + result
}

func (f AssignStmtFlag) String() string {
	return f.toString("")
}

func (f AssignStmtFlag) IsNormal() bool {
	return f == AssignStmtFlagNormal
}

func (f AssignStmtFlag) IsInit() bool {
	return f&AssignStmtFlagInit != 0
}

func (f AssignStmtFlag) IsGlobal() bool {
	return f&AssignStmtFlagGlobal != 0
}

func (f AssignStmtFlag) IsMatrix() bool {
	return f&AssignStmtFlagMatrix != 0
}

type AssignStmt struct {
	flag     AssignStmtFlag
	declList []Declaratorer
	initList []AssignIniter
}

var AssignStmtHelper *AssignStmt

func (s *AssignStmt) New(flag AssignStmtFlag, declList []Declaratorer, initList []AssignIniter) *AssignStmt {
	return &AssignStmt{
		flag:     flag,
		declList: declList,
		initList: initList,
	}
}

func (s *AssignStmt) funcStatementer() {}

func (s *AssignStmt) definitioner() {}

func (s *AssignStmt) toString(ident string) string {
	if s == nil {
		return ident + "Null"
	}

	return fmt.Sprintf(""+
		"%sAssignStatement {\n"+
		"%s..InitFlag: \n"+
		"%s\n"+
		"%s..DeclaratorList:\n"+
		"%s\n"+
		"%s..InitializerList:\n"+
		"%s\n"+
		"%s}",
		ident, ident, s.flag.toString(ident+"...."),
		ident, IterableToString(ident+"....", IteratableDeclaratorerList(s.declList)),
		ident, IterableToString(ident+"....", IteratableAssignIniterList(s.initList)),
		ident,
	)
}

func (s *AssignStmt) String() string {
	return s.toString("")
}

func (s *AssignStmt) Flag() AssignStmtFlag {
	return s.flag
}

func (s *AssignStmt) DeclList() []Declaratorer {
	return s.declList
}

func (s *AssignStmt) InitList() []AssignIniter {
	return s.initList
}

func (s *AssignStmt) AddFlag(flag AssignStmtFlag) {
	s.flag |= flag
}
