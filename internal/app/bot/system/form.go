package system

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/vendors/bark"
	"github.com/tsundata/assistant/internal/pkg/vendors/pushover"
)

var formRules = []bot.FormRule{
	{
		ID:    PushSwitchFormID,
		Title: "Push notification switch",
		Field: []bot.FieldItem{
			{
				Key:      pushover.ID,
				Type:     bot.FieldItemTypeBool,
				Required: true,
			},
			{
				Key:      bark.ID,
				Type:     bot.FieldItemTypeBool,
				Required: true,
			},
		},
		SubmitFunc: func(ctx context.Context, botCtx bot.Context, comp component.Component) []pb.MsgPayload {
			for _, item := range botCtx.FieldItem {
				comp.GetRedis().HSet(ctx, "system:push:switch", item.Key, item.Value == "1")
			}
			return []pb.MsgPayload{
				pb.TextMsg{Text: "switch success"},
			}
		},
	},
	{
		ID:    SubscribeSwitchFormID,
		Title: "Subscribe switch",
		Field: []bot.FieldItem{},
		SubmitFunc: func(ctx context.Context, botCtx bot.Context, comp component.Component) []pb.MsgPayload {
			var kv []*pb.KV
			for _, item := range botCtx.FieldItem {
				kv = append(kv, &pb.KV{
					Key:   item.Key,
					Value: item.Value.(string),
				})
			}
			reply, err := comp.Middle().SwitchUserSubscribe(ctx, &pb.SwitchUserSubscribeRequest{Subscribe: kv})
			if err != nil {
				return []pb.MsgPayload{
					pb.TextMsg{Text: err.Error()},
				}
			}
			if reply.State {
				return []pb.MsgPayload{
					pb.TextMsg{Text: "switch success"},
				}
			}
			return []pb.MsgPayload{
				pb.TextMsg{Text: "switch failed"},
			}
		},
	},
	{
		ID:    WebhookSwitchFormID,
		Title: "Script Webhook switch",
		Field: []bot.FieldItem{},
		SubmitFunc: func(ctx context.Context, botCtx bot.Context, comp component.Component) []pb.MsgPayload {
			var kv []*pb.KV
			for _, item := range botCtx.FieldItem {
				kv = append(kv, &pb.KV{
					Key:   item.Key,
					Value: item.Value.(string),
				})
			}
			reply, err := comp.Chatbot().SwitchTriggers(ctx, &pb.SwitchTriggersRequest{Triggers: kv})
			if err != nil {
				return []pb.MsgPayload{
					pb.TextMsg{Text: err.Error()},
				}
			}
			if reply.State {
				return []pb.MsgPayload{
					pb.TextMsg{Text: "switch success"},
				}
			}
			return []pb.MsgPayload{
				pb.TextMsg{Text: "switch failed"},
			}
		},
	},
	{
		ID:    CronSwitchFormID,
		Title: "Script Cron switch",
		Field: []bot.FieldItem{},
		SubmitFunc: func(ctx context.Context, botCtx bot.Context, comp component.Component) []pb.MsgPayload {
			var kv []*pb.KV
			for _, item := range botCtx.FieldItem {
				kv = append(kv, &pb.KV{
					Key:   item.Key,
					Value: item.Value.(string),
				})
			}
			reply, err := comp.Chatbot().SwitchTriggers(ctx, &pb.SwitchTriggersRequest{Triggers: kv})
			if err != nil {
				return []pb.MsgPayload{
					pb.TextMsg{Text: err.Error()},
				}
			}
			if reply.State {
				return []pb.MsgPayload{
					pb.TextMsg{Text: "switch success"},
				}
			}
			return []pb.MsgPayload{
				pb.TextMsg{Text: "switch failed"},
			}
		},
	},
}
