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

	// i := NewInterpreter(tree, nil)
	// r, err := i.Interpret()
	// if err != nil {
	// 	 t.Fatal(err)
	// }
	// if r != 0 {
	// 	 t.Fatal("error expr")
	// }
	// log.Println(i.callStack)
	// log.Println(i.Stdout())
}

func TestInterpreter(t *testing.T) {
	text := `
#!/usr/bin/env flowscript

node abc (cron):
	with: {
			"mode": "custom",
			"cron_expression": "* * * * *"
		}
	secret: github_key
end

node xkcd (http):
	with: { 
			"method": "GET",
	 		"url": "https://xkcd.com",
	 		"response_format": "html",
	 		"headers":  { "X-FOO": "BAR" },
	 		"query": { "foo": "bar"},
			"extract": {
				"url": {
				  "css": "#comic img",
				  "value": "@src"
				},
				"title": {
				  "css": "#comic img",
				  "value": "@alt"
				},
				"hovertext": {
				  "css": "#comic img",
				  "value": "@title"
				}
			}
	 	}
end

node httpbin (http):
	with: { 
			"method": "GET",
	 		"url": "https://httpbin.org/json",
	 		"response_format": "json",
	 		"headers":  { "X-FOO": "BAR" },
	 		"query": { "foo": "bar" },
			"extract": {
			  "title": { "path": "slideshow.slides.#.title", "value": "string(.)" },
			  "type": { "path": "slideshow.slides.#.type", "value": "string(.)" }
			}
	 	}
end

node hi (http):
	with: { 
			"method": "GET",
	 		"url": "https://httpbin.org/uuid",
	 		"response_format": "text",
			"extract": {
			  "uuid": { "regexp": "(\w{8}\-\w{4}\-\w{4}\-\w{4}\-\w{12})", "index": 1, "value": "string(.)" }
			}
	 	}
end

node notice (pushover):
	with: {
		"title": "title - {{0.title}}",
		"message": "message - {{0.title}}",
		"url": "{{0.url}}"
	}
	secret: pushover
end

workflow demo:
    @xkcd -> @httpbin
end

workflow main:
    @abc -> @xkcd -> @httpbin -> @hi -> @notice;
	@xkcd -> @notice
end
`
	run(t, text)
}
