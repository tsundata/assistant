package component

import (
	"fmt"
	"html/template"
)

type Text struct {
	Name  string
	Title string
}

func (c *Text) GetContent() template.HTML {
	return template.HTML(fmt.Sprintf(`<div class="text">%s</div>`, c.Title))
}
