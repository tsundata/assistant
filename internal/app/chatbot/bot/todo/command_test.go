package todo

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/mock"
	"testing"
)

func parseCommand(t *testing.T, comp component.Component, in string) []pb.MsgPayload {
	for _, rule := range Bot.CommandRule {
		tokens, err := command.ParseCommand(in)
		if err != nil {
			t.Fatal(err)
		}
		check, err := command.SyntaxCheck(rule.Define, tokens)
		if err != nil {
			t.Fatal(err)
		}
		if !check {
			continue
		}

		if ret := rule.Parse(context.Background(), comp, tokens); len(ret) > 0 {
			return ret
		}
	}

	return []pb.MsgPayload{}
}

func TestTodoList(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	todo := mock.NewMockTodoSvcServer(ctl)
	gomock.InOrder(
		todo.EXPECT().GetTodos(gomock.Any(), gomock.Any()).Return(&pb.TodosReply{Todos: []*pb.Todo{
			{
				Id:        1,
				Priority:  1,
				Content:   "todo",
				Complete:  true,
				UpdatedAt: 946659600,
			},
		}}, nil),
	)

	cmd := "todo list"
	comp := component.MockComponent(todo)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []pb.MsgPayload{pb.TableMsg{
		Title:  "Todo",
		Header: []string{"Id", "Priority", "Content", "Complete"},
		Row: [][]interface{}{
			{"1", "1", "todo", "true"},
		},
	}}, res)
}

func TestTodoCommand(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	todo := mock.NewMockTodoSvcServer(ctl)
	gomock.InOrder(
		todo.EXPECT().CreateTodo(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	cmd := "todo create test1"
	comp := component.MockComponent(todo)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []pb.MsgPayload{pb.TextMsg{Text: "success"}}, res)
}

func TestRemindCommand(t *testing.T) {
	cmd := "remind test 19:50"
	comp := component.MockComponent()
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []pb.MsgPayload{}, res)
}
