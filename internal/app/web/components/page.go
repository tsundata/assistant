package components

import (
	"fmt"
	"html/template"
)

type Page struct {
	Title   string
	Action  Component
	Content Component
}

func (c *Page) GetContent() template.HTML {
	var action template.HTML
	if c.Action != nil {
		action = c.Action.GetContent()
	}
	return template.HTML(fmt.Sprintf(`
<div class="page">
	<div class="title">
		%s
		<span> %s </span>
	</div>
	<div class="content">
		%s
	</div>
</div>`, c.Title, action, c.Content.GetContent()))
}
