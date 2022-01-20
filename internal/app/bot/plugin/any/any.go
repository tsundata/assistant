package any

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/plugin"
	"log"
)

type Any struct {
	Next plugin.Handler
}

func (a Any) Run(ctx context.Context, input interface{}) (interface{}, error) {
	log.Println(a.Name())
	return plugin.NextOrFailure(a.Name(), a.Next, ctx, input)
}

func (a Any) Name() string {
	return "any"
}
