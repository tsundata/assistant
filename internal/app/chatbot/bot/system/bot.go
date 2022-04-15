package system

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/version"
)

const (
	DemoActionId = "demo"

	PushSwitchFormID      = "push_switch"
	SubscribeSwitchFormID = "subscribe_switch"
	WebhookSwitchFormID   = "webhook_switch"
	CronSwitchFormID      = "cron_switch"
)

var metadata = bot.Metadata{
	Name:       "System",
	Identifier: enum.SystemBot,
	Detail:     "",
	Avatar:     "",
}

var setting []bot.FieldItem

var workflowRules = bot.WorkflowRule{
	Plugin: []bot.PluginRule{
		{
			Name:  "expr",
			Param: []interface{}{"Trim(Value)"},
		},
		{
			Name:  "keyword",
			Param: []interface{}{"info", "version"},
		},
	},
	RunFunc: func(ctx context.Context, botCtx bot.Context, comp component.Component) []pb.MsgPayload {
		var result []pb.MsgPayload
		if a, ok := botCtx.Input.Stack[1].([]string); ok {
			for _, keyword := range a {
				switch keyword {
				case "info":
					result = append(result, pb.TextMsg{Text: botCtx.Input.Value})
				case "version":
					result = append(result, pb.TextMsg{Text: version.Info()})
				}
			}
		}
		return result
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

var tagRules = []bot.TagRule{
	{
		Tag: "test",
		TriggerFunc: func(ctx context.Context, botCtx bot.Context, comp component.Component) []pb.MsgPayload {
			return []pb.MsgPayload{
				pb.TextMsg{Text: "test tag"},
			}
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
