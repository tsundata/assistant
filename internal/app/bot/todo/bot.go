package todo

import (
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
)

var Metadata = bot.Metadata{
	Name:       "Todo",
	Identifier: enum.TodoBot,
	Detail:     "",
	Avatar:     "",
}

var Setting = []bot.SettingField{
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

var PluginRules = []bot.PluginRule{
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
	Bot, err = bot.NewBot(Metadata, Setting, PluginRules)
	if err != nil {
		panic(err)
	}
}
