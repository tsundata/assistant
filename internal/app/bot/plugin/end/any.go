package end

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/plugin"
	"log"
)

type End struct {
	Next plugin.Handler
}

func (a End) Run(_ context.Context, _ interface{}) (interface{}, error) {
	log.Println(a.Name())
	return nil, nil
}

func (a End) Name() string {
	return "end"
}
