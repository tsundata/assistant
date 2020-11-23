package interpreter

import (
	"testing"
)

func TestInterpreter_Expr(t *testing.T) {
	p, err := NewParser(NewLexer("7 + 3 * (10 / (12 / (3 + 1) - 1)) / (2 + 3) - 5 - 3 + (8)"))
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
