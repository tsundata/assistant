package component

import (
	"fmt"
	"html/template"
	"math/rand"
	"time"
)

type Input struct {
	Name  string
	Title string
	Type  string
	Value string
}

func (c *Input) GetContent() template.HTML {
	rand.Seed(time.Now().Unix())
	id := rand.Int()
	return template.HTML(fmt.Sprintf(`
<div class="input">
	<label for="input-%d">%s:</label>
  	<input type="%s" id="input-%d" value="%s" name="%s">
</div>`, id, c.Title, c.Type, id, c.Value, c.Name))
}
