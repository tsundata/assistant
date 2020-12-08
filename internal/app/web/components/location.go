package components

import "html/template"

type Location struct {
	Name string
	URL  string
}

func (c *Location) GetContent() template.HTML {
	return ""
}
