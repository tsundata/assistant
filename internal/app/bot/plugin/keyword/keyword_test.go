package keyword

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/util"
	"testing"
)

func TestKeyword(t *testing.T) {
	p := Keyword{}
	input := bot.PluginValue{Value: "test", Stack: []interface{}{}}
	ctrl := &bot.Controller{}
	params := []util.Value{
		util.Variable("test"),
	}
	output, err := p.Run(context.Background(), ctrl, params, input)
	if err != nil {
		t.Fatal(err)
	}
	input.Stack = []interface{}{[]string{"test"}}
	assert.Equal(t, output.Stack[0], []string{"test"})
	assert.Equal(t, input, output)
}
