package trim

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/plugin/end"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"testing"
)

func TestKeyword(t *testing.T) {
	p := Trim{
		Next: end.End{},
	}
	input := bot.PluginValue{Value: " test  ", Stack: make(map[string]interface{})}
	ctrl := bot.MockController(map[string][]interface{}{})
	out, err := p.Run(context.Background(), ctrl, input)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, input.Stack[p.Name()], "test")
	input.Value = "test"
	assert.Equal(t, input, out)
}
