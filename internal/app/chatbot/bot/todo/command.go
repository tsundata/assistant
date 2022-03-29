package todo

import (
	"context"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/util"
	"strconv"
	"strings"
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

			tableString := &strings.Builder{}
			if len(reply.Todos) > 0 {
				table := tablewriter.NewWriter(tableString)
				table.SetBorder(false)
				table.SetHeader([]string{"Id", "Priority", "Content", "Complete"})
				for _, v := range reply.Todos {
					table.Append([]string{strconv.Itoa(int(v.Id)), strconv.Itoa(int(v.Priority)), v.Content, util.BoolToString(v.Complete)})
				}
				table.Render()
			}
			if tableString.String() == "" {
				return []pb.MsgPayload{pb.TextMsg{Text: "Empty"}}
			}

			return []pb.MsgPayload{pb.TextMsg{Text: tableString.String()}}
		},
	},
	{
		Define: `todo [string]`,
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
	{
		Define: `remind [string] [string]`,
		Help:   `Remind something`,
		Parse: func(ctx context.Context, comp component.Component, tokens []*command.Token) []pb.MsgPayload {
			arg1 := tokens[1].Value
			arg2 := tokens[2].Value
			fmt.Println(arg1, arg2) // todo remind message

			return []pb.MsgPayload{}
		},
	},
}