package components

import (
	"fmt"
	"html/template"
	"strings"
)

type Form struct {
	Name   string
	Action string
	Method string
	Inputs []Component
}

func (c *Form) GetContent() template.HTML {
	buf := new(strings.Builder)
	for _, item := range c.Inputs {
		buf.WriteString(string(item.GetContent()))
	}
	return template.HTML(fmt.Sprintf(`<form action="%s" method="%s" class="form" autocomplete="off">
%s
<div class="button">
	<button type="submit">Submit</button>
	<button type="reset">Reset</button>
</div>
</form>`, c.Action, c.Method, buf.String()))
}
