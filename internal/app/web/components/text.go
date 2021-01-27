package components

import (
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"html/template"
	"regexp"
	"strings"
)

type Text struct {
	Name  string
	Title string
}

func (c *Text) GetContent() template.HTML {
	// a, br
	re, _ := regexp.Compile(utils.UrlRegex)
	s := re.FindString(c.Title)
	if s != "" {
		c.Title = strings.ReplaceAll(c.Title, s, fmt.Sprintf(`<a href="%s" target="_blank">%s</a>`, s, s))
	}
	c.Title = strings.ReplaceAll(c.Title, "\n", "<br>")

	return template.HTML(fmt.Sprintf(`<h2>%s</h2>`, c.Title))
}
