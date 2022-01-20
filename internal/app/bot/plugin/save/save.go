package save

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/plugin"
	"log"
)

type Save struct {
	Next plugin.Handler
}

func (a Save) Run(ctx context.Context, input interface{}) (interface{}, error) {
	log.Println(a.Name())
	return plugin.NextOrFailure(a.Name(), a.Next, ctx, input)
}

func (a Save) Name() string {
	return "save"
}
