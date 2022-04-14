package save

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
)

type Save struct {
	Next bot.PluginHandler
}

func (a Save) Run(ctx context.Context, ctrl *bot.Controller, input bot.PluginValue) (bot.PluginValue, error) {
	return bot.NextOrFailure(ctx, a.Name(), a.Next, ctrl, input)
}

func (a Save) Name() string {
	return "save"
}
