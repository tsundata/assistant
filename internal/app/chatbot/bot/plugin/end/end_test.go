package end

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"testing"
)

func TestEnd(t *testing.T) {
	p := End{}
	input := bot.PluginValue{Value: "test"}
	out, err := p.Run(context.Background(), bot.MockController(), input)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, input, out)
}
