package todo

import (
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
)

var metadata = bot.Metadata{
	Name:       "Todo",
	Identifier: enum.TodoBot,
	Detail:     "",
	Avatar:     "",
}

var setting = []bot.SettingField{
	{
		Key:      "report",
		Type:     bot.SettingItemTypeBool,
		Required: false,
		Value:    false,
	},
	{
		Key:      "time",
		Type:     bot.SettingItemTypeString,
		Required: true,
		Value:    "",
	},
}

var workflowRules = []bot.PluginRule{
	{
		Name: "any",
	},
	{
		Name:  "filter",
		Param: []interface{}{1},
	},
	{
		Name: "save",
	},
}

var Bot *bot.Bot

func init() {
	var err error
	Bot, err = bot.NewBot(metadata, setting, workflowRules, commandRules)
	if err != nil {
		panic(err)
	}
}
