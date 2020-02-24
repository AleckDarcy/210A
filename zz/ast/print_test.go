package ast

import (
	"fmt"
	"testing"
)

func TestPrintStatement_String(t *testing.T) {
	p := PrintStatement1
	p.funcStatementer()
	p.List()
	fmt.Println(p)
}
