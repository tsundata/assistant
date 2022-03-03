package any

import (
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"log"
)

type Any struct {
	Next bot.PluginHandler
}

func (a Any) Run(ctrl *bot.Controller, input interface{}) (interface{}, error) {
	log.Println(a.Name())
	return bot.NextOrFailure(a.Name(), a.Next, ctrl, input)
}

func (a Any) Name() string {
	return "any"
}
