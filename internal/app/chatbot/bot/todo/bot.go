package todo

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
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

var actionRules = []bot.ActionRule{
	{
		ID: "demo",
		Title: "demo?",
		OptionFunc: map[string]bot.OptionFunc{
			"true": func(ctx context.Context, component component.Component) []pb.MsgPayload {
				return []pb.MsgPayload{
					pb.TextMsg{Text: "true"},
				}
			},
			"false": func(ctx context.Context, component component.Component) []pb.MsgPayload {
				return []pb.MsgPayload{
					pb.TextMsg{Text: "false"},
				}
			},
		},
	},
}

var Bot *bot.Bot

func init() {
	var err error
	Bot, err = bot.NewBot(metadata, setting, workflowRules, commandRules, actionRules)
	if err != nil {
		panic(err)
	}
}
