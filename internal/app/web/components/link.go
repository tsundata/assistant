package components

import (
	"fmt"
	"html/template"
)

type Link struct {
	Name  string
	Title string
	URL   string
}

func (c *Link) GetContent() template.HTML {
	return template.HTML(fmt.Sprintf(`<h2>%s</h2>
            <p class="link">
                <a href="%s">
                    <span class="link-block">
                        <span class="link-icon"></span>
                        <span class="link-text">%s</span>
                    </span>
                </a>
            </p>`, c.Title, c.URL, c.Name))
}
