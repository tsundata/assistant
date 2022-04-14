package keyword

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"strings"
)

type Keyword struct {
	Next bot.PluginHandler
}

func (a Keyword) Run(ctx context.Context, ctrl *bot.Controller, input bot.PluginValue) (bot.PluginValue, error) {
	var in []string
	for _, keyword := range bot.Param(ctrl, a) {
		if s, ok := keyword.(string); ok {
			if strings.Contains(input.Value, s) {
				in = append(in, s)
			}
		}
	}
	input.Stack[a.Name()] = in
	return bot.NextOrFailure(ctx, a.Name(), a.Next, ctrl, input)
}

func (a Keyword) Name() string {
	return "keyword"
}
