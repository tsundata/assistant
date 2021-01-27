package components

import (
	"fmt"
	"html/template"
	"strings"
)

type Memo struct {
	Name    string
	Time    string
	Content Component
	Tags    []string
}

func (c *Memo) GetContent() template.HTML {
	var tags strings.Builder
	for _, tag := range c.Tags {
		tags.WriteString("<span>")
		tags.WriteString(tag)
		tags.WriteString("</span>")
	}

	return template.HTML(fmt.Sprintf(`
<div class="memo">
	<div class="time">%s</div>
	<div class="tags">%s</div>
	<div class="content">%s</div>
</div>
`, c.Time, tags.String(), c.Content.GetContent()))
}
