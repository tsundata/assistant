package filter

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/plugin"
	"log"
)

type Filter struct {
	Next plugin.Handler
}

func (a Filter) Run(ctx context.Context, input interface{}) (interface{}, error) {
	log.Println(a.Name())
	return plugin.NextOrFailure(a.Name(), a.Next, ctx, input)
}

func (a Filter) Name() string {
	return "filter"
}
