package component

import (
	"fmt"
	"html/template"
)

type LinkButton struct {
	Name  string
	Title string
	URL   string
}

func (c *LinkButton) GetContent() template.HTML {
	return template.HTML(fmt.Sprintf(`
<div class="link-button">
	<h2>%s</h2>
	<p class="link-block">
		<a href="%s">
			<span class="link-content">
				<span class="link-icon"></span>
				<span class="link-text">%s</span>
			</span>
		</a>
	</p>
</div>
`, c.Title, c.URL, c.Name)) // #nosec
}
