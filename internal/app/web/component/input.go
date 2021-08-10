package component

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"math/big"
)

type Input struct {
	Name  string
	Title string
	Type  string
	Value string
}

func (c *Input) GetContent() template.HTML {
	n, _ := rand.Int(rand.Reader, big.NewInt(1000))
	return template.HTML(fmt.Sprintf(`
<div class="input">
	<label for="input-%d">%s:</label>
  	<input type="%s" id="input-%d" value="%s" name="%s">
</div>`, n.Int64(), c.Title, c.Type, n.Int64(), c.Value, c.Name))
}
