package component

import (
	"html/template"
	"strings"
)

type List struct {
	Items []Component
}

func (c *List) GetContent() template.HTML {
	buf := new(strings.Builder)
	for _, item := range c.Items {
		buf.WriteString(string(item.GetContent()))
	}
	return template.HTML(buf.String()) // #nosec
}
