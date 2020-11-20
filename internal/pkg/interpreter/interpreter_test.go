package interpreter

import (
	"fmt"
	"testing"
)

func TestInterpreter_Expr(t *testing.T) {
	l := NewLexer("10 * 4  * 2 * 3 / 8")
	i, err := NewInterpreter(l)
	if err != nil {
		t.Fatal(err)
	}
	r, err := i.Parse()
	if err != nil {
		t.Fatal(err)
	}
	if r != 30 {
		fmt.Println(r, 30)
		t.Fatal("error expr")
	}
}
