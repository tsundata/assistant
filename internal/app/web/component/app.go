package component

import (
	"fmt"
	"html/template"
)

type App struct {
	Name string
	Icon string
	Text string
	URL  string
}

func (c *App) GetContent() template.HTML {
	return template.HTML(fmt.Sprintf(`
<div class="app">
	<a href="%s"><i class="fa fa-%s fa-4x"></i>
	<span>%s</span></a>
</div>`, c.URL, c.Icon, c.Text)) // #nosec
}
