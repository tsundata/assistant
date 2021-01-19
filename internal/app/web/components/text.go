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
	c.Title = strings.Replace(c.Title, s, fmt.Sprintf(`<a href="%s" target="_blank">%s<a>`, s, s), -1)
	c.Title = strings.Replace(c.Title, "\n", "<br>", -1)

	return template.HTML(fmt.Sprintf(`<h2>%s</h2>`, c.Title))
}
