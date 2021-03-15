package action

import (
	"log"
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

	Debug = true
	symbolTable := NewSemanticAnalyzer()
	err = symbolTable.Visit(tree)
	if err != nil {
		t.Fatal(err)
	}
	if Debug {
		log.Println(symbolTable.CurrentScope)
	}

	i := NewInterpreter(tree)
	r, err := i.Interpret()
	if err != nil {
		t.Fatal(err)
	}
	if r != 0 {
		t.Fatal("error expr")
	}
}

func TestInterpreter(t *testing.T) {
	text := `
get "https://httpbin.org/get"
json
count
pdf
send "success"
echo 1 1.2 "hi" #1
`
	run(t, text)
}

func TestInterpreter2(t *testing.T) {
	text := `get "https://httpbin.org/get"
json
count
send "hello world"
`
	run(t, text)
}

func TestInterpreter3(t *testing.T) {
	text := `
set "[1, 2]"
json
if
send "success"
else
send "error"

set 1
if
send "success"
else
send "error"

set "a"
if
send "success"
else
send "error"

set 1.2
if
send "success"
else
send "error"

set true
if
send "success"
else
send "error"
`
	run(t, text)
}
