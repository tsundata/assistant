package system

import (
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
)

var metadata = bot.Metadata{
	Name:       "System",
	Identifier: enum.SystemBot,
	Detail:     "",
	Avatar:     "",
}

var setting []bot.SettingField

var workflowRules []bot.PluginRule

var Bot *bot.Bot

func init() {
	var err error
	Bot, err = bot.NewBot(metadata, setting, workflowRules, commandRules)
	if err != nil {
		panic(err)
	}
}
