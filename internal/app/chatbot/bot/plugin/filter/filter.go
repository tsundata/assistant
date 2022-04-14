package filter

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
)

type Filter struct {
	Next bot.PluginHandler
}

func (a Filter) Run(ctx context.Context, ctrl *bot.Controller, input bot.PluginValue) (bot.PluginValue, error) {
	return bot.NextOrFailure(ctx, a.Name(), a.Next, ctrl, input)
}

func (a Filter) Name() string {
	return "filter"
}
