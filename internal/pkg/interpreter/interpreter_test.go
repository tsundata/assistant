package interpreter

import "testing"

func TestInterpreter_Expr(t *testing.T) {
	i := NewInterpreter("12+33")
	r, err := i.Expr()
	if err != nil {
		t.Fatal(err)
	}
	if r != 45 {
		t.Fatal("error expr")
	}

	i2 := NewInterpreter("1  + 1")
	r2, err := i2.Expr()
	if err != nil {
		t.Fatal(err)
	}
	if r2 != 2 {
		t.Fatal("error expr")
	}

	i3 := NewInterpreter("66 - 33")
	r3, err := i3.Expr()
	if err != nil {
		t.Fatal(err)
	}
	if r3 != 33 {
		t.Fatal("error expr")
	}
}
