package todo

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

const (
	DemoActionId = "demo"
	DemoFormId   = "demo"
)

var metadata = bot.Metadata{
	Name:       "Todo",
	Identifier: enum.TodoBot,
	Detail:     "",
	Avatar:     "",
}

var setting = []bot.FieldItem{
	{
		Key:      "report",
		Type:     bot.FieldItemTypeBool,
		Required: false,
		Value:    false,
	},
	{
		Key:      "time",
		Type:     bot.FieldItemTypeString,
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
		ID:    DemoActionId,
		Title: "demo action?",
		OptionFunc: map[string]bot.ActFunc{
			"true": func(ctx context.Context, botCtx bot.Context, component component.Component) []pb.MsgPayload {
				return []pb.MsgPayload{
					pb.TextMsg{Text: "true"},
				}
			},
			"false": func(ctx context.Context, botCtx bot.Context, component component.Component) []pb.MsgPayload {
				return []pb.MsgPayload{
					pb.TextMsg{Text: "false"},
				}
			},
		},
	},
}

var formRules = []bot.FormRule{
	{
		ID:    DemoFormId,
		Title: "demo form?",
		Field: []bot.FieldItem{
			{
				Key:      "title",
				Type:     bot.FieldItemTypeString,
				Required: true,
			},
		},
		SubmitFunc: func(ctx context.Context, botCtx bot.Context, component component.Component) []pb.MsgPayload {
			return []pb.MsgPayload{
				pb.TextMsg{Text: "submit!"},
			}
		},
	},
}

var tagRules = []bot.TagRule{
	{
		Tag: "todo",
		TriggerFunc: func(ctx context.Context, botCtx bot.Context, comp component.Component) []pb.MsgPayload {
			return nil
		},
	},
}

var Bot *bot.Bot

func init() {
	var err error
	Bot, err = bot.NewBot(metadata, setting, workflowRules, commandRules, actionRules, formRules, tagRules)
	if err != nil {
		panic(err)
	}
}
