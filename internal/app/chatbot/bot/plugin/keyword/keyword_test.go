package keyword

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/plugin/end"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"testing"
)

func TestKeyword(t *testing.T) {
	p := Keyword{
		Next: end.End{},
	}
	input := bot.PluginValue{Value: "test", Stack: make(map[string]interface{})}
	ctrl := bot.MockController(map[string][]interface{}{
		"keyword": {"test"},
	})
	output, err := p.Run(context.Background(), ctrl, input)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, input.Stack[p.Name()], 1)
	assert.Equal(t, input.Stack[p.Name()], []string{"test"})
	assert.Equal(t, input, output)
}
