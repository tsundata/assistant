package save

import (
	"github.com/stretchr/testify/assert"
	"github.com/tsundata/assistant/internal/app/bot/plugin/end"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"testing"
)

func TestSave(t *testing.T) {
	p := Save{
		Next: end.End{},
	}
	input := "test"
	out, err := p.Run(bot.MockController(), input)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, input, out)
}
