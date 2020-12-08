package components

import "html/template"

type Script struct {
	Name string
	Code string
}

func (c *Script) GetContent() template.HTML {
	return ""
}
