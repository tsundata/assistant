package component

import "html/template"

type File struct {
	Name string
	URL  string
}

func (c *File) GetContent() template.HTML {
	return ""
}
