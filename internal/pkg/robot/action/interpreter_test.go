package action

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
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
		fmt.Println(symbolTable.CurrentScope)
	}

	mq, err := event.CreateRabbitmq(enum.Chatbot)
	if err != nil {
		t.Fatal(err)
	}
	bus := event.NewRabbitmqBus(mq, nil)

	comp := component.MockComponent(bus)
	i := NewInterpreter(context.Background(), tree)
	i.SetComponent(comp)
	i.SetMessage(pb.Message{UserId: 1, GroupId: 1})
	_, err = i.Interpret()
	if err != nil {
		t.Fatal(err)
	}
}

func TestInterpreter(t *testing.T) {
	text := `
get "https://httpbin.org/get"
json
count
pdf
echo "ok"
echo 1 1.2 "hi" #1
`
	run(t, text)
}

func TestInterpreter2(t *testing.T) {
	text := `get "https://httpbin.org/get"
json
count
echo "hello world"
`
	run(t, text)
}

func TestInterpreter3(t *testing.T) {
	text := `
set "[1, 2]"
json
if
echo "ok"
else
echo "error"

set 1
if
echo "ok"
else
echo "error"

set "a"
if
echo "ok"
else
echo "error"

set 1.2
if
echo "ok"
else
echo "error"

set true
if
echo "ok"
else
echo "error"
`
	run(t, text)
}

func TestInterpreter4(t *testing.T) {
	text := `
// status 
status "http" "https://www.example.com"
status "tcp" "www.example.com:80"
status "dns" "8.8.8.8:53"
status "tls" "www.example.com:443"
`
	run(t, text)
}

func TestInterpreter5(t *testing.T) {
	text := `set "aaa" "bbb" "ccc" "aaa" "ccc"
dedupe
message`
	run(t, text)
}

func TestInterpreter6(t *testing.T) {
	text := `
get "https://httpbin.org/html"
query "css" "h1" "text"

get "https://httpbin.org/get"
query "json" "headers.Host"

get "https://httpbin.org/robots.txt"
query "regex" "^Disallow: .*$"
`
	run(t, text)
}

func TestInterpreter7(t *testing.T) {
	text := `
counter "abc"
increase "abc"
decrease "abc"
`
	run(t, text)
}

func TestInterpreter8(t *testing.T) {
	text := `
counter "abc"
watch "counter" "change(value)"
`
	run(t, text)
}
