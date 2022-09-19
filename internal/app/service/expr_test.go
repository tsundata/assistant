package service

import (
	"github.com/antonmedv/expr"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/mock"
	"testing"
)

func TestExprEnv_Run(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	message := mock.NewMockMessageSvcClient(ctl)
	middle := mock.NewMockMiddleSvcClient(ctl)
	todo := mock.NewMockTodoSvcServer(ctl)
	okr := mock.NewMockOkrSvcServer(ctl)
	gomock.InOrder(
		message.EXPECT().GetBySequence(gomock.Any(), gomock.Any()).Return(&pb.GetMessageReply{Message: &pb.Message{
			Id:       1,
			UserId:   1,
			Sequence: 1,
			Text:     "test",
		}}, nil),
		middle.EXPECT().GetCounterByFlag(gomock.Any(), gomock.Any()).Return(&pb.CounterReply{Counter: &pb.Counter{
			Id:     1,
			UserId: 1,
			Flag:   "test",
			Digit:  1,
		}}, nil),
		todo.EXPECT().GetTodo(gomock.Any(), gomock.Any()).Return(&pb.TodoReply{Todo: &pb.Todo{
			Id:       1,
			UserId:   1,
			Sequence: 1,
			Content:  "test",
			Complete: true,
		}}, nil),
		okr.EXPECT().GetKeyResult(gomock.Any(), gomock.Any()).Return(&pb.KeyResultReply{KeyResult: &pb.KeyResult{
			Id:           1,
			UserId:       1,
			Sequence:     1,
			Title:        "test",
			CurrentValue: 1,
			TargetValue:  10,
		}}, nil),
	)

	comp := component.MockComponent(message, middle, todo, okr)
	type args struct {
		variable string
		expr     string
	}
	tests := []struct {
		name   string
		args   args
		output interface{}
	}{
		{"case1", args{variable: "Message(1)", expr: "Value.Text == \"test\""}, true},
		{"case2", args{variable: "Counter(\"test\")", expr: "Value.Digit == 1"}, true},
		{"case3", args{variable: "Todo(1)", expr: "Value.Complete == true"}, true},
		{"case4", args{variable: "KeyResult(1)", expr: "Value.CurrentValue == 1"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// run variable
			program, err := expr.Compile(tt.args.variable, expr.Env(ExprEnv{}))
			if err != nil {
				t.Fatal(err)
			}
			value, err := expr.Run(program, ExprEnv{Comp: comp})
			if err != nil {
				t.Fatal(err)
			}

			// run expr
			program, err = expr.Compile(tt.args.expr, expr.Env(ExprEnv{}))
			if err != nil {
				t.Fatal(err)
			}
			output, err := expr.Run(program, ExprEnv{Comp: comp, Value: value})
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(t, tt.output, output)
		})
	}
}
