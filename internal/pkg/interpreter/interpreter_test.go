package interpreter

import (
	"fmt"
	"testing"
)

func TestInterpreter_Expr(t *testing.T) {
	text := `BEGIN

    BEGIN
        number := 2;
        a := number;
        b := 10 * a + 10 * number / 4;
        c := a - - b
    END;

    x := 11;
END.`
	p, err := NewParser(NewLexer([]rune(text)))
	if err != nil {
		t.Fatal(err)
	}
	i := NewInterpreter(p)
	r, err := i.Interpret()
	if err != nil {
		t.Fatal(err)
	}
	if r != 0 {
		t.Fatal("error expr")
	}
	fmt.Println(i.GlobalScope)
}
