package save

import (
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"log"
)

type Save struct {
	Next bot.PluginHandler
}

func (a Save) Run(ctrl *bot.Controller, input interface{}) (interface{}, error) {
	log.Println(a.Name())
	return bot.NextOrFailure(a.Name(), a.Next, ctrl, input)
}

func (a Save) Name() string {
	return "save"
}
