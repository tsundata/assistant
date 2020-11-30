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
	fmt.Println(i.Stdout())
}

func TestInterpreter(t *testing.T) {
	text := `PROGRAM Part10;
VAR
   number      INT;
   a, b, c, x  INT;
   y           FLOAT;

BEGIN //Part10
   BEGIN
      number := 2;
      a := number;
      b := 10 * a + 10 * number DIV 4;
      c := a - - b
   END;
   x := 11;
   y := 20 / 7 + 3.14;
   // writeln('a = ', a); 
   // writeln('b = ', b); 
   // writeln('c = ', c); 
   // writeln('number = ', number); 
   // writeln('x = ', x); 
   // writeln('y = ', y); 
END.  //Part10`
	run(t, text)
}

func TestInterpreterFunction(t *testing.T) {
	text := `PROGRAM Part12;
VAR
   a  INT;

FUNC P1;
VAR
   a  FLOAT;
   k  INT;

   FUNC P2;
   VAR
      a, z  INT;
   BEGIN //P2
      z := 777;
   END;  //P2

BEGIN //P1

END;  //P1

BEGIN //Part12
   a := 10;
END.  //Part12`
	run(t, text)
}

func TestInterpreterNestedScopes(t *testing.T) {
	text := `program Main;
   var b, x, y  FLOAT;
   var z  INT;

   FUNC AlphaA(a : INT);
      var b  INT;

      FUNC Beta(c : INT);
         var y  INT;

         FUNC Gamma(c : INT);
            var x  INT;
         begin // Gamma 
            x := a + b + c + x + y + z;
         end;  // Gamma 

      begin // Beta 

      end;  // Beta 

   begin // AlphaA 

   end;  // AlphaA 

   FUNC AlphaB(a : INT);
      var c  FLOAT;
   begin // AlphaB 
      c := a + b;
   end;  // AlphaB 

begin // Main 
end.  // Main `
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
   a  INT;

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
   b  INT;

BEGIN
   b := 1;
   a := b + 2;
END.`
	run(t, text)
}

func TestInterpreterFunctionCall(t *testing.T) {
	text := `program Main;

FUNC Alpha(a : INT; b : INT);
var x  INT;
begin
   x := (a + b ) * 2;
end;

begin // Main 

   Alpha(3 + 5, 7);  // FUNC call 

end.  // Main `
	run(t, text)
}

func TestInterpreterFunctionCall2(t *testing.T) {
	text := `program exFunc;
var a, b, c,  min INT;

FUNC findMin(x, y, z: INT;  m: INT);  // Finds the minimum of the 3 values 
begin
   if x < y then
      m := x
   else
      m := y
   end;

   if z < m then
      m := z
   end;
end; // end of FUNC findMin   

begin
   	a := 1;
	b := 2;
	c := 3;
   	findMin(a, b, c, min); // FUNC call 
end.`
	run(t, text)
}

func TestInterpreterCallStack(t *testing.T) {
	text := `program Main;
var x, y  INT;
begin // Main 
   y := 7;
   x := (y + 3) * 3;
end.  // Main `
	run(t, text)
}

func TestInterpreterExecutingFunctionCalls(t *testing.T) {
	text := `program Main;

FUNC Alpha(a : INT; b : INT);
var x  INT;
begin
   x := (a + b ) * 2;
end;

begin // Main 

   Alpha(3 + 5, 7);  // FUNC call 

end.  // Main `
	run(t, text)
}

func TestInterpreterNestedFunctionCalls(t *testing.T) {
	text := `program Main;

FUNC Alpha(a : INT; b : INT);
var x  INT;

   FUNC Beta(a : INT; b : INT);
   var x  INT;
   begin
      x := a * 10 + b * 2;
   end;

begin
   x := (a + b ) * 2;

   Beta(5, 10);      // FUNC call 
end;

begin // Main 

   Alpha(3 + 5, 7);  // FUNC call 

end.  // Main `
	run(t, text)
}

func TestInterpreterIf(t *testing.T) {
	text := `program Main;
var x, y  INT;
begin // Main 
   	y := 7;
   	x := (y + 3) * 3;
	if x < y and x > y or y > 0 then
		// then branch 
	else
		// else branch 
	end
end.  // Main `
	run(t, text)
}

func TestInterpreterWhile(t *testing.T) {
	text := `program whileLoop;
var a INT;

begin
   a := 1;
   while  a < 20  do
      a := a + a
   end;
end.`
	run(t, text)
}

func TestInterpreterString(t *testing.T) {
	text := `program stringTest;
var a, b string;

begin
   	a := "abc";
	b := "foobar";
end.`
	run(t, text)
}

func TestInterpreterBoolean(t *testing.T) {
	text := `program booleanTest;
var a, b bool;

begin
   	a := true;
	b := false;
	if a then
		// then branch 
	else
		// else branch 
	end;
end.`
	run(t, text)
}

func TestInterpreterPrint(t *testing.T) {
	text := `program booleanTest;
var a, b bool;
var s string;

begin
   	a := true;
	b := false;
	s := "hi";
	if b then
		// then branch 
	else
		// else branch 
		print a ;
		print b ;
		print s ;
	end;
end.`
	run(t, text)
}

func TestInterpreterFunctionReturn(t *testing.T) {
	text := `program exFunc;
var a, b, c,  min  int;

func findMin(x, y, z: INT) : int ;  // Finds the minimum of the 3 values 
var m int;
begin
	if x < y then
	  m := x
	else
	  m := y
	end;
	
	if z < m then
	  m := z
	end;
	
	print 111;
	return m;
	print 222;
	return m;
	print 333;	
	return m;
	
end; // end of FUNC findMin   

begin
   	a := 1;
	b := 2;
	c := 3;
   	min := findMin(a, b, c) + findMin(a, b, c) + findMin(a, b, c); // FUNC call 
	print min;
end.`
	run(t, text)
}

func TestInterpreterBuildInFunction(t *testing.T) {
	text := `program exFunc;
var a, b, c,  min  int;
    s string;
func findMin(x, y, z: INT) : int ;  // Finds the minimum of the 3 values 
var m int;
begin
	if x < y then
	  m := x
	else
	  m := y
	end;
	
	if z < m then
	  m := z
	end;

	return m;
	return m + 1;
	return m + 2;
	
end; // end of FUNC findMin   

begin
	s := "abc";
   	a := 1;
	b := 2;
	c := 3;
   	min := findMin(a, b, c) + findMin(a, b, c) + findMin(a, b, c); // FUNC call 
	print min;
	print findMin(a, b, c) + len(s);
end.`
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
