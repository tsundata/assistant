package interpreter

import (
	"fmt"
	"testing"
)

func TestInterpreter(t *testing.T) {
	text := `PROGRAM Part10;
VAR
   number     : INTEGER;
   a, b, c, x : INTEGER;
   y          : REAL;

BEGIN {Part10}
   BEGIN
      number := 2;
      a := number;
      b := 10 * a + 10 * number DIV 4;
      c := a - - b
   END;
   x := 11;
   y := 20 / 7 + 3.14;
   { writeln('a = ', a); }
   { writeln('b = ', b); }
   { writeln('c = ', c); }
   { writeln('number = ', number); }
   { writeln('x = ', x); }
   { writeln('y = ', y); }
END.  {Part10}`
	p, err := NewParser(NewLexer([]rune(text)))
	if err != nil {
		t.Fatal(err)
	}
	tree, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	symbolTable := NewSymbolTableBuilder()
	symbolTable.Visit(tree)
	fmt.Println(symbolTable.symbolTable)

	i := NewInterpreter(tree)
	r, err := i.Interpret()
	if err != nil {
		t.Fatal(err)
	}
	if r != 0 {
		t.Fatal("error expr")
	}
	fmt.Println(i.GlobalMemory)
}

func TestInterpreterProcedure(t *testing.T) {
	text := `PROGRAM Part12;
VAR
   a : INTEGER;

PROCEDURE P1;
VAR
   a : REAL;
   k : INTEGER;

   PROCEDURE P2;
   VAR
      a, z : INTEGER;
   BEGIN {P2}
      z := 777;
   END;  {P2}

BEGIN {P1}

END;  {P1}

BEGIN {Part12}
   a := 10;
END.  {Part12}`
	p, err := NewParser(NewLexer([]rune(text)))
	if err != nil {
		t.Fatal(err)
	}
	tree, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	symbolTable := NewSymbolTableBuilder()
	symbolTable.Visit(tree)
	fmt.Println(symbolTable.symbolTable)

	i := NewInterpreter(tree)
	r, err := i.Interpret()
	if err != nil {
		t.Fatal(err)
	}
	if r != 0 {
		t.Fatal("error expr")
	}
	fmt.Println(i.GlobalMemory)
}

func TestInterpreterNameError1(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic, NameError1")
		}
	}()

	text := `PROGRAM NameError1;
VAR
   a : INTEGER;

BEGIN
   a := 2 + b;
END.`
	p, err := NewParser(NewLexer([]rune(text)))
	if err != nil {
		t.Fatal(err)
	}
	tree, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	symbolTable := NewSymbolTableBuilder()
	symbolTable.Visit(tree)
	fmt.Println(symbolTable.symbolTable)

	i := NewInterpreter(tree)
	r, err := i.Interpret()
	if err != nil {
		t.Fatal(err)
	}
	if r != 0 {
		t.Fatal("error expr")
	}
	fmt.Println(i.GlobalMemory)
}

func TestInterpreterNameError2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic, NameError2")
		}
	}()

	text := `PROGRAM NameError2;
VAR
   b : INTEGER;

BEGIN
   b := 1;
   a := b + 2;
END.`
	p, err := NewParser(NewLexer([]rune(text)))
	if err != nil {
		t.Fatal(err)
	}
	tree, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	symbolTable := NewSymbolTableBuilder()
	symbolTable.Visit(tree)
	fmt.Println(symbolTable.symbolTable)

	i := NewInterpreter(tree)
	r, err := i.Interpret()
	if err != nil {
		t.Fatal(err)
	}
	if r != 0 {
		t.Fatal("error expr")
	}
	fmt.Println(i.GlobalMemory)
}
