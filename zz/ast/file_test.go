package ast

import (
	"fmt"
	"testing"
)

var File1 = &File{
	defList: []Definitioner{
		AssignStmt1,
		FuncDefinition1,
	},
}

func TestFile_String(t *testing.T) {
	f := File1
	fmt.Println(f)
}
