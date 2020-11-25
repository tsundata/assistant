package interpreter

import (
	"fmt"
	"testing"
)

func run(t *testing.T, text string) {
	p, err := NewParser(NewLexer([]rune(text)))
	if err != nil {
		t.Fatal(err)
	}
	tree, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	symbolTable := NewSemanticAnalyzer()
	symbolTable.Visit(tree)
	fmt.Println(symbolTable.CurrentScope)

	i := NewInterpreter(tree)
	r, err := i.Interpret()
	if err != nil {
		t.Fatal(err)
	}
	if r != 0 {
		t.Fatal("error expr")
	}
	fmt.Println(i.callStack)
}

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
	run(t, text)
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
	run(t, text)
}

func TestInterpreterNestedScopes(t *testing.T) {
	text := `program Main;
   var b, x, y : real;
   var z : integer;

   procedure AlphaA(a : integer);
      var b : integer;

      procedure Beta(c : integer);
         var y : integer;

         procedure Gamma(c : integer);
            var x : integer;
         begin { Gamma }
            x := a + b + c + x + y + z;
         end;  { Gamma }

      begin { Beta }

      end;  { Beta }

   begin { AlphaA }

   end;  { AlphaA }

   procedure AlphaB(a : integer);
      var c : real;
   begin { AlphaB }
      c := a + b;
   end;  { AlphaB }

begin { Main }
end.  { Main }`
	run(t, text)
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
	run(t, text)
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
	run(t, text)
}

func TestInterpreterProcedureCall(t *testing.T) {
	text := `program Main;

procedure Alpha(a : integer; b : integer);
var x : integer;
begin
   x := (a + b ) * 2;
end;

begin { Main }

   Alpha(3 + 5, 7);  { procedure call }

end.  { Main }`
	run(t, text)
}

func TestInterpreterCallStack(t *testing.T) {
	text := `program Main;
var x, y : integer;
begin { Main }
   y := 7;
   x := (y + 3) * 3;
end.  { Main }`
	run(t, text)
}

func TestCallStack(t *testing.T) {
	s := NewCallStack()
	s.Push(NewActivationRecord("a", ARTypeProgram, 1))
	s.Push(NewActivationRecord("b", ARTypeProgram, 1))
	s.Push(NewActivationRecord("c", ARTypeProgram, 1))
	fmt.Println(s)
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s)
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s)
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s)
	fmt.Println(s.Peek())
}