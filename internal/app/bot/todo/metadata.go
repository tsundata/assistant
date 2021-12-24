package todo

import (
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/pkg/robot"
)

var Metadata = robot.Metadata{
	Name:       "Todo",
	Identifier: enum.TodoBot,
	Detail:     "",
	Avatar:     "",
	Setting: []robot.SettingItem{
		{
			Key:      "report",
			Type:     robot.SettingItemTypeBool,
			Required: false,
			Value:    false,
		},
		{
			Key:      "time",
			Type:     robot.SettingItemTypeString,
			Required: true,
			Value:    "",
		},
	},
}
