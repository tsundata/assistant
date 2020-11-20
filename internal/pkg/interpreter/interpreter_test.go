package interpreter

import (
	"fmt"
	"testing"
)

func TestInterpreter_Expr(t *testing.T) {
	l := NewLexer("14 + 2 * 3 - 6 / 2")
	i, err := NewInterpreter(l)
	if err != nil {
		t.Fatal(err)
	}
	r, err := i.Parse()
	if err != nil {
		t.Fatal(err)
	}
	if r != 17 {
		fmt.Println(r, 30)
		t.Fatal("error expr")
	}
}
