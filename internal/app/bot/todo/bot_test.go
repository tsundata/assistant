package todo

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/command"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/mock"
	"testing"
)

func parseCommand(t *testing.T, comp command.Component, in string) []string {
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

	return []string{}
}

func TestTodoList(t *testing.T) {
	t.SkipNow()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	todo := mock.NewMockTodoSvcClient(ctl)
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
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"  ID | PRIORITY | CONTENT | COMPLETE  \n-----+----------+---------+-----------\n   1 |        1 | todo    | true      \n"}, res)
}

func TestTodoCommand(t *testing.T) {
	t.SkipNow()
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	todo := mock.NewMockTodoSvcClient(ctl)
	gomock.InOrder(
		todo.EXPECT().CreateTodo(gomock.Any(), gomock.Any()).Return(&pb.StateReply{State: true}, nil),
	)

	cmd := "todo test1"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{"success"}, res)
}

func TestRemindCommand(t *testing.T) {
	cmd := "remind test 19:50"
	comp := rulebot.NewComponent(nil, nil, nil, nil, nil, nil,
		nil, nil, nil)
	res := parseCommand(t, comp, cmd)
	require.Equal(t, []string{}, res)
}
