package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/mock"
	"reflect"
	"testing"

	"github.com/tsundata/assistant/api/pb"
)

func TestChatbot_Handle(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	bot, err := rulebot.CreateRuleBot(enum.Chatbot)
	if err != nil {
		t.Fatal(err)
	}

	middle := mock.NewMockMiddleSvcClient(ctl)
	todo := mock.NewMockTodoSvcClient(ctl)

	s := NewChatbot(nil, middle, todo, bot)

	type args struct {
		in0     context.Context
		payload *pb.ChatbotRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.ChatbotReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.ChatbotRequest{Text: ""}},
			&pb.ChatbotReply{Text: []string{}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Handle(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chatbot.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chatbot.Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}
