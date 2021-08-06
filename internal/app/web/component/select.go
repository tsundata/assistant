package component

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"math/big"
	"strings"
)

type Select struct {
	Name  string
	Title string
	Value map[string]string
}

func (c *Select) GetContent() template.HTML {
	n, _ := rand.Int(rand.Reader, big.NewInt(1000))

	var o strings.Builder
	for k, v := range c.Value {
		o.WriteString(fmt.Sprintf(`<option value="%s">%s</option>`, k, v))
	}

	return template.HTML(fmt.Sprintf(`
<div class="select">
	<label for="input-%d">%s:</label>
	<select name="%s" id="select-%d">%s</select>
</div>`, n.Int64(), c.Title, c.Name, n.Int64(), o.String()))
}
