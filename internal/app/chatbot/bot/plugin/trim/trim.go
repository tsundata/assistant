package trim

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"strings"
)

type Trim struct {
	Next bot.PluginHandler
}

func (a Trim) Run(ctx context.Context, ctrl *bot.Controller, input bot.PluginValue) (bot.PluginValue, error) {
	input.Value = strings.TrimSpace(input.Value)
	input.Stack[a.Name()] = input.Value
	return bot.NextOrFailure(ctx, a.Name(), a.Next, ctrl, input)
}

func (a Trim) Name() string {
	return "trim"
}
