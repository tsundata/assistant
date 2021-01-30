package components

import (
	"fmt"
	"html/template"
)

type Button struct {
	Name  string
	Title string
	URL   string
}

func (c *Button) GetContent() template.HTML {
	return template.HTML(fmt.Sprintf(`<a href="%s" class="button">%s</a>`, c.URL, c.Title))
}
