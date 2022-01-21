package any

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"log"
)

type Any struct {
	Next bot.PluginHandler
}

func (a Any) Run(ctx context.Context, input interface{}) (interface{}, error) {
	log.Println(a.Name())
	return bot.NextOrFailure(a.Name(), a.Next, ctx, input)
}

func (a Any) Name() string {
	return "any"
}
