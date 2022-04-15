package keyword

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"strings"
)

type Keyword struct {}

func (a Keyword) Run(_ context.Context, _ *bot.Controller, param []interface{}, input bot.PluginValue) (bot.PluginValue, error) {
	var in []string
	for _, keyword := range param {
		if s, ok := keyword.(string); ok {
			if strings.Contains(input.Value, s) {
				in = append(in, s)
			}
		}
	}
	input.Stack = append(input.Stack, in)
	return input, nil
}

func (a Keyword) Name() string {
	return "keyword"
}
