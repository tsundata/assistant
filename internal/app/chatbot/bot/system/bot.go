package system

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/vendors/bark"
	"github.com/tsundata/assistant/internal/pkg/vendors/pushover"
)

const PushSwitchFormID = "push_switch"

var metadata = bot.Metadata{
	Name:       "System",
	Identifier: enum.SystemBot,
	Detail:     "",
	Avatar:     "",
}

var setting []bot.FieldItem

var workflowRules []bot.PluginRule

var formRules = []bot.FormRule{
	{
		ID:    PushSwitchFormID,
		Title: "Push notification switch",
		Field: []bot.FieldItem{
			{
				Key:      enum.SystemBot,
				Type:     bot.FieldItemTypeBool,
				Required: true,
			},
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
		SubmitFunc: func(ctx context.Context, c component.Component, form []bot.FieldItem) []pb.MsgPayload {
			for _, item := range form {
				c.GetRedis().Set(ctx, fmt.Sprintf("system:push:%s", item.Key), item.Value == "1", redis.KeepTTL)
			}
			return []pb.MsgPayload{
				pb.TextMsg{Text: "switch success"},
			}
		},
	},
}

var Bot *bot.Bot

func init() {
	var err error
	Bot, err = bot.NewBot(metadata, setting, workflowRules, commandRules, nil, formRules)
	if err != nil {
		panic(err)
	}
}
