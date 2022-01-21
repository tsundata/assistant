package end

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"log"
)

type End struct {
	Next bot.PluginHandler
}

func (a End) Run(_ context.Context, _ interface{}) (interface{}, error) {
	log.Println(a.Name())
	return nil, nil
}

func (a End) Name() string {
	return "end"
}
