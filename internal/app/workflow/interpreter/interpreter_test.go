package interpreter

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

	symbolTable := NewSemanticAnalyzer()
	symbolTable.Visit(tree)
	log.Println(symbolTable.CurrentScope)

	i := NewInterpreter(tree)
	r, err := i.Interpret()
	if err != nil {
		t.Fatal(err)
	}
	if r != 0 {
		t.Fatal("error expr")
	}
	log.Println(i.callStack)
	log.Println(i.Stdout())
}

func TestInterpreter(t *testing.T) {
	text := `
node abc (cron):
	with: {
			"mode": "custom",
			"cron_expression": "* * * * *"
		}
	secret: github_key
end

node news (http):
	with: { 
			"method": "GET",
	 		"url": "https://httpbin.org/get",
	 		"response_format": "json",
	 		"headers":  { "X-FOO": "BAR" },
	 		"query": { "foo": "bar"}
	 	}
end

node notice (http):
	with: { 
			"method": "GET",
	 		"url": "https://httpbin.org/get",
	 		"response_format": "json",
	 		"headers":  { "X-FOO": "BAR" },
	 		"query": { "foo": "bar" }
	 	}
end

workflow demo:
    @notice -> @news
end

workflow main:
    @abc -> @news -> @notice
end
`
	run(t, text)
}
