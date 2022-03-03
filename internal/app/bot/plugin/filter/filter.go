package filter

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"log"
)

type Filter struct {
	Next bot.PluginHandler
}

func (a Filter) Run(ctx context.Context, input interface{}) (interface{}, error) {
	log.Println(a.Name())
	return bot.NextOrFailure(a.Name(), a.Next, ctx, input)
}

func (a Filter) Name() string {
	return "filter"
}
