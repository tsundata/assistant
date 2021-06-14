package component

import (
	"fmt"
	"html/template"
	"math/rand"
	"strings"
	"time"
)

type Select struct {
	Name  string
	Title string
	Value map[string]string
}

func (c *Select) GetContent() template.HTML {
	rand.Seed(time.Now().Unix())
	id := rand.Int()

	var o strings.Builder
	for k, v := range c.Value {
		o.WriteString(fmt.Sprintf(`<option value="%s">%s</option>`, k, v))
	}

	return template.HTML(fmt.Sprintf(`
<div class="select">
	<label for="input-%d">%s:</label>
	<select name="%s" id="select-%d">%s</select>
</div>`, id, c.Title, c.Name, id, o.String()))
}
