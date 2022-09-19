package todo

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"strconv"
)

var formRules = []bot.FormRule{
	{
		ID:    CreateTodoFormID,
		Title: "Create todo",
		Field: []bot.FieldItem{
			{
				Key:      "content",
				Type:     bot.FieldItemTypeString,
				Required: true,
			},
			{
				Key:      "category",
				Type:     bot.FieldItemTypeString,
				Required: true,
			},
			{
				Key:      "remark",
				Type:     bot.FieldItemTypeString,
				Required: true,
			},
			{
				Key:      "priority",
				Type:     bot.FieldItemTypeInt,
				Required: true,
			},
		},
		SubmitFunc: func(ctx context.Context, botCtx bot.Context, comp component.Component) []pb.MsgPayload {
			if comp.Todo() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}

			var todo pb.Todo
			for _, item := range botCtx.FieldItem {
				switch item.Key {
				case "content":
					todo.Content = item.Value.(string)
				case "category":
					todo.Category = item.Value.(string)
				case "remark":
					todo.Remark = item.Value.(string)
				case "priority":
					priority, _ := strconv.ParseInt(item.Value.(string), 10, 64)
					todo.Priority = priority
				}
			}

			reply, err := comp.Todo().CreateTodo(ctx, &pb.TodoRequest{
				Todo: &todo,
			})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			if reply.GetState() {
				return []pb.MsgPayload{pb.TextMsg{Text: "ok"}}
			}

			return []pb.MsgPayload{}
		},
	},
	{
		ID:    UpdateTodoFormID,
		Title: "Update todo",
		Field: []bot.FieldItem{
			{
				Key:      "sequence",
				Type:     bot.FieldItemTypeInt,
				Required: true,
			},
			{
				Key:      "content",
				Type:     bot.FieldItemTypeString,
				Required: true,
			},
			{
				Key:      "category",
				Type:     bot.FieldItemTypeString,
				Required: true,
			},
			{
				Key:      "remark",
				Type:     bot.FieldItemTypeString,
				Required: true,
			},
			{
				Key:      "priority",
				Type:     bot.FieldItemTypeInt,
				Required: true,
			},
		},
		SubmitFunc: func(ctx context.Context, botCtx bot.Context, comp component.Component) []pb.MsgPayload {
			if comp.Todo() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}

			var todo pb.Todo
			for _, item := range botCtx.FieldItem {
				switch item.Key {
				case "sequence":
					sequence, _ := strconv.ParseInt(item.Value.(string), 10, 64)
					todo.Sequence = sequence
				case "content":
					todo.Content = item.Value.(string)
				case "category":
					todo.Category = item.Value.(string)
				case "remark":
					todo.Remark = item.Value.(string)
				case "priority":
					priority, _ := strconv.ParseInt(item.Value.(string), 10, 64)
					todo.Priority = priority
				}
			}

			reply, err := comp.Todo().UpdateTodo(ctx, &pb.TodoRequest{
				Todo: &todo,
			})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			if reply.GetState() {
				return []pb.MsgPayload{pb.TextMsg{Text: "ok"}}
			}

			return []pb.MsgPayload{}
		},
	},
}
