package end

import (
	"github.com/stretchr/testify/assert"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"testing"
)

func TestEnd(t *testing.T) {
	p := End{}
	input := "test"
	out, err := p.Run(bot.MockController(), input)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, input, out)
}
