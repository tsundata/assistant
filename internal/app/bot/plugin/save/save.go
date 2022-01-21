package save

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"log"
)

type Save struct {
	Next bot.PluginHandler
}

func (a Save) Run(ctx context.Context, input interface{}) (interface{}, error) {
	log.Println(a.Name())
	return bot.NextOrFailure(a.Name(), a.Next, ctx, input)
}

func (a Save) Name() string {
	return "save"
}
