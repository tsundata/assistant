package end

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
)

type End struct {
	Next bot.PluginHandler
}

func (a End) Run(_ context.Context, _ *bot.Controller, input bot.PluginValue) (bot.PluginValue, error) {
	return input, nil
}

func (a End) Name() string {
	return "end"
}
