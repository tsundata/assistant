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

var workflowRules = bot.WorkflowRule{
	Plugin: []bot.PluginRule{
		{
			Name:  "expr",
			Param: []interface{}{"len(Value)"},
		},
	},
	RunFunc: func(ctx context.Context, botCtx bot.Context, comp component.Component) []pb.MsgPayload {
		return []pb.MsgPayload{
			pb.TextMsg{Text: "todo workflow run"},
		}
	},
}

var actionRules []bot.ActionRule

var formRules []bot.FormRule

var tagRules []bot.TagRule

var Bot *bot.Bot

func init() {
	var err error
	Bot, err = bot.NewBot(metadata, setting, nil, workflowRules, commandRules, actionRules, formRules, tagRules)
	if err != nil {
		panic(err)
	}
}
