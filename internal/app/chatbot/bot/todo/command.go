package todo

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/util"
	"strconv"
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

			var header []string
			var row [][]interface{}
			if len(reply.Todos) > 0 {
				header = []string{"Id", "Priority", "Content", "Complete"}
				for _, v := range reply.Todos {
					row = append(row, []interface{}{strconv.Itoa(int(v.Id)), strconv.Itoa(int(v.Priority)), v.Content, util.BoolToString(v.Complete)})
				}
			}
			if len(row) == 0 {
				return []pb.MsgPayload{pb.TextMsg{Text: "Empty"}}
			}

			return []pb.MsgPayload{pb.TableMsg{
				Title:  "Todo",
				Header: header,
				Row:    row,
			}}
		},
	},
	{
		Define: `todo create [string]`,
		Help:   "Todo something",
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			if comp.Todo() == nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "empty client"}}
			}
			reply, err := comp.Todo().CreateTodo(ctx, &pb.TodoRequest{
				Todo: &pb.Todo{Content: tokens[1].Value.(string)},
			})
			if err != nil {
				return []pb.MsgPayload{pb.TextMsg{Text: "error call: " + err.Error()}}
			}
			if !reply.GetState() {
				return []pb.MsgPayload{pb.TextMsg{Text: "failed"}}
			}
			return []pb.MsgPayload{pb.TextMsg{Text: "success"}}
		},
	},
}
