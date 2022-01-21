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

var Setting = []bot.SettingItem{
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
