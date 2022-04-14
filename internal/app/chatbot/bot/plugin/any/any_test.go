package any

import (
	"context"
	"github.com/stretchr/testify/assert"
	end2 "github.com/tsundata/assistant/internal/app/chatbot/bot/plugin/end"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"testing"
)

func TestAny(t *testing.T) {
	p := Any{
		Next: end2.End{},
	}
	input := bot.PluginValue{Value: "test"}
	out, err := p.Run(context.Background(), bot.MockController(), input)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, input, out)
}
