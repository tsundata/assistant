package any

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
)

type Any struct {
	Next bot.PluginHandler
}

func (a Any) Run(ctx context.Context, ctrl *bot.Controller, input bot.PluginValue) (bot.PluginValue, error) {
	return bot.NextOrFailure(ctx, a.Name(), a.Next, ctrl, input)
}

func (a Any) Name() string {
	return "any"
}
