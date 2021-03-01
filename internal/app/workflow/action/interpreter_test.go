package action

import (
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
count
send "hello world"
`
	run(t, text)
}
