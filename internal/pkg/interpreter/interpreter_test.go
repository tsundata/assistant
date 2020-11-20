package interpreter

import "testing"

func TestInterpreter_Expr(t *testing.T) {
	i := NewInterpreter("9 - 5 + 3 + 11")
	r, err := i.Expr()
	if err != nil {
		t.Fatal(err)
	}
	if r != 18 {
		t.Fatal("error expr")
	}

	i2 := NewInterpreter("2  * 3")
	r2, err := i2.Expr()
	if err != nil {
		t.Fatal(err)
	}
	if r2 != 6 {
		t.Fatal("error expr")
	}

	i3 := NewInterpreter("66 / 33")
	r3, err := i3.Expr()
	if err != nil {
		t.Fatal(err)
	}
	if r3 != 2 {
		t.Fatal("error expr")
	}
}
