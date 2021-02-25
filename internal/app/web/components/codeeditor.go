package components

import (
	"fmt"
	"html/template"
)

type CodeEditor struct {
	Name string
}

func (c *CodeEditor) GetContent() template.HTML {
	return template.HTML(fmt.Sprintf(`<textarea id="code" name="%s" placeholder="Code goes here..."></textarea>`, c.Name))
}
