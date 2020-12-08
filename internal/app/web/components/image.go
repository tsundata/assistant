package components

import "html/template"

type Image struct {
	Name string
	URL  string
}

func (c *Image) GetContent() template.HTML {
	return ""
}
