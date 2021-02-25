package components

import (
	"fmt"
	"html/template"
)

type Script struct {
	Name    string
	ID      int
	UUID    string
	Content string
}

func (c *Script) GetContent() template.HTML {
	return template.HTML(fmt.Sprintf(`<div class="script"> 
	<div class="title"><div class="id">ID: %d</div> <div class="run-btn" onclick="if(confirm('Is it running?')){ window.location = '/script/%s/run?id=%d' }">Run</div></div> 
	<pre>%s</pre>
</div>`, c.ID, c.UUID, c.ID, c.Content))
}
