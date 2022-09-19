package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/mock"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestMessage_List(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockMessageRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().List(gomock.Any()).Return([]*pb.Message{{
			Text: "test",
		}}, nil),
	)

	s := NewMessage(nil, nil, nil, nil, repo, nil, nil, nil)

	type args struct {
		in0 context.Context
		in1 *pb.MessageRequest
	}
	tests := []struct {
		name    string
		m       *Message
		args    args
		want    int
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.MessageRequest{}},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.List(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Message.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && len(got.Messages) != tt.want {
				t.Errorf("Message.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_Get(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockMessageRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(pb.Message{
			Id:   1,
			Text: "test",
			Type: "text",
		}, nil),
	)

	s := NewMessage(nil, nil, nil, nil, repo, nil, nil, nil)

	type args struct {
		in0     context.Context
		payload *pb.MessageRequest
	}
	tests := []struct {
		name    string
		m       *Message
		args    args
		want    *pb.MessageReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.MessageRequest{Message: &pb.Message{Id: 1}}},
			&pb.MessageReply{
				Message: &pb.Message{
					Id:   1,
					Text: "test",
					Type: "text",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.GetById(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Message.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && (got.Message.GetText() != tt.want.Message.GetText() || got.Message.Type != tt.want.Message.Type) {
				t.Errorf("Message.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_Create(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mq, err := event.CreateRabbitmq(enum.Message)
	if err != nil {
		t.Fatal(err)
	}
	bus := event.NewRabbitmqBus(mq, nil)

	repo := mock.NewMockMessageRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(int64(1), nil),
	)

	s := NewMessage(bus, nil, nil, nil, repo, nil, nil, nil)

	type args struct {
		in0     context.Context
		payload *pb.MessageRequest
	}
	tests := []struct {
		name    string
		m       *Message
		args    args
		want    *pb.MessageReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.MessageRequest{Message: &pb.Message{Text: "demo2", Payload: "{}"}}},
			&pb.MessageReply{Message: &pb.Message{Type: "text", Text: "demo2", Payload: "{}", SenderType: enum.MessageUserType, ReceiverType: enum.MessageGroupType}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Create(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Message.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Message.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_Delete(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockMessageRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil),
		repo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(gorm.ErrRecordNotFound),
	)

	s := NewMessage(nil, nil, nil, nil, repo, nil, nil, nil)

	type args struct {
		in0     context.Context
		payload *pb.MessageRequest
	}
	tests := []struct {
		name    string
		m       *Message
		args    args
		want    *pb.TextReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.MessageRequest{Message: &pb.Message{Id: 1}}},
			&pb.TextReply{Text: ""},
			false,
		},
		{
			"case2",
			s,
			args{context.Background(), &pb.MessageRequest{Message: &pb.Message{Id: 2}}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Delete(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Message.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Message.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_Run(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	chatbot := mock.NewMockChatbotSvcClient(ctl)
	repo := mock.NewMockMessageRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().GetByID(gomock.Any(), gomock.Any()).
			Return(pb.Message{Id: 1, Text: "test", Type: string(enum.MessageTypeScript)}, nil),
		chatbot.EXPECT().RunActionScript(gomock.Any(), gomock.Any()).
			Return(&pb.WorkflowReply{Text: "ok"}, nil),

		repo.EXPECT().GetByID(gomock.Any(), gomock.Any()).
			Return(pb.Message{Id: 1, Text: "test", Type: "other"}, nil),
	)

	s := NewMessage(nil, nil, nil, nil, repo, chatbot, nil, nil)

	type args struct {
		ctx     context.Context
		payload *pb.MessageRequest
	}
	tests := []struct {
		name    string
		m       *Message
		args    args
		want    *pb.TextReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.MessageRequest{Message: &pb.Message{Id: 1}}},
			&pb.TextReply{Text: "ok"},
			false,
		},
		{
			"case2",
			s,
			args{context.Background(), &pb.MessageRequest{Message: &pb.Message{Id: 2}}},
			&pb.TextReply{Text: "Not running"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Run(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Message.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Message.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
