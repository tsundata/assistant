package components

import (
	"fmt"
	"html/template"
)

type Page struct {
	Title   string
	Content Component
}

func (c *Page) GetContent() template.HTML {
	return template.HTML(fmt.Sprintf(`<div class="container">
        <div class="title">
            %s
        </div>
        <div class="content">
            %s
        </div>
    </div>`, c.Title, c.Content.GetContent()))
}
