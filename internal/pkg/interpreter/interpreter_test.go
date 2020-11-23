package interpreter

import (
	"testing"
)

func TestInterpreter_Expr(t *testing.T) {
	p, err := NewParser(NewLexer("5 - - - + - (3 + 4) - +2"))
	if err != nil {
		t.Fatal(err)
	}
	i := NewInterpreter(p)
	r, err := i.Interpret()
	if err != nil {
		t.Fatal(err)
	}
	if r != 10 {
		t.Fatal("error expr")
	}
}
