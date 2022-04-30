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

	m := pb.Message{
		Id:       1,
		Sequence: 1,
		UserId:   1,
		Text:     "test",
	}

	message := mock.NewMockMessageSvcClient(ctl)
	gomock.InOrder(
		message.EXPECT().GetBySequence(gomock.Any(), gomock.Any()).Return(&pb.GetMessageReply{Message: &m}, nil),
		message.EXPECT().GetBySequence(gomock.Any(), gomock.Any()).Return(&pb.GetMessageReply{Message: &m}, nil),
	)

	comp := component.MockComponent(message)
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
		{"case1", args{variable: "Message(1)", expr: "Value.Text != \"test\""}, false},
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
