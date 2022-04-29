package todo

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/bot/msg"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

var commandRules = []command.Rule{
	{
		Define: `todo list`,
		Help:   `List todo`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Todo() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}

			reply, err := comp.Todo().GetTodos(ctx, &pb.TodoRequest{})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}

			return []pb.MsgPayload{pb.TodoMsg{
				Title: "Todo",
				Todo:  reply.Todos,
			}}
		},
	},
	{
		Define: `todo create`,
		Help:   "Create Todo something",
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			return []pb.MsgPayload{msg.BotFormMsg(Bot.FormRule, CreateTodoFormID)}
		},
	},
	{
		Define: `todo update [number]`,
		Help:   "Update Todo something",
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Todo() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}
			sequence, _ := tokens[2].Value.Int64()

			reply, err := comp.Todo().GetTodo(ctx, &pb.TodoRequest{
				Todo: &pb.Todo{Sequence: sequence},
			})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}

			return []pb.MsgPayload{pb.FormMsg{
				ID:    UpdateTodoFormID,
				Title: fmt.Sprintf("Update Todo #%d", sequence),
				Field: []pb.FormField{
					{
						Key:      "sequence",
						Type:     string(bot.FieldItemTypeInt),
						Required: true,
						Value:    reply.Todo.Sequence,
					},
					{
						Key:      "content",
						Type:     string(bot.FieldItemTypeString),
						Required: true,
						Value:    reply.Todo.Content,
					},
					{
						Key:      "category",
						Type:     string(bot.FieldItemTypeString),
						Required: true,
						Value:    reply.Todo.Category,
					},
					{
						Key:      "remark",
						Type:     string(bot.FieldItemTypeString),
						Required: true,
						Value:    reply.Todo.Remark,
					},
					{
						Key:      "priority",
						Type:     string(bot.FieldItemTypeInt),
						Required: true,
						Value:    reply.Todo.Priority,
					},
				},
			}}
		},
	},
}
