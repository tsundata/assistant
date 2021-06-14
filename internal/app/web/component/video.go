package component

import "html/template"

type Video struct {
	Name string
	URL  string
}

func (c *Video) GetContent() template.HTML {
	return ""
}
