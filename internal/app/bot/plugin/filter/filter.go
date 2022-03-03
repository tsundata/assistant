package filter

import (
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"log"
)

type Filter struct {
	Next bot.PluginHandler
}

func (a Filter) Run(ctrl *bot.Controller, input interface{}) (interface{}, error) {
	log.Println(a.Name())
	param := bot.Param(ctrl, a)
	log.Println(param)
	return bot.NextOrFailure(a.Name(), a.Next, ctrl, input)
}

func (a Filter) Name() string {
	return "filter"
}
