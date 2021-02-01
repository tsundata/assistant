package components

import (
	"fmt"
	"html/template"
)

type Link struct {
	Name  string
	Title string
	URL   string
}

func (c *Link) GetContent() template.HTML {
	return template.HTML(fmt.Sprintf(`<a href="%s" class="link">%s</a>`, c.URL, c.Title))
}
