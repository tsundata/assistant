package any

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"log"
)

type Any struct {
	Next bot.PluginHandler
}

func (a Any) Run(ctx context.Context, ctrl *bot.Controller, input interface{}) (interface{}, error) {
	log.Println(a.Name())
	return bot.NextOrFailure(ctx, a.Name(), a.Next, ctrl, input)
}

func (a Any) Name() string {
	return "any"
}
