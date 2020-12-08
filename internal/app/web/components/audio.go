package components

import "html/template"

type Audio struct {
	Name string
	URL  string
}

func (c *Audio) GetContent() template.HTML {
	return ""
}
