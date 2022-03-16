package service

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/repository"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/internal/pkg/vendors"
	"github.com/tsundata/assistant/mock"
	"gorm.io/gorm"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestChatbot_Handle(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mq, err := event.CreateRabbitmq(enum.Chatbot)
	if err != nil {
		t.Fatal(err)
	}
	bus := event.NewRabbitmqBus(mq, nil)

	bot, err := rulebot.CreateRuleBot(enum.Chatbot)
	if err != nil {
		t.Fatal(err)
	}

	comp := component.MockComponent()

	message := mock.NewMockMessageSvcClient(ctl)
	middle := mock.NewMockMiddleSvcClient(ctl)
	repo := mock.NewMockChatbotRepository(ctl)

	gomock.InOrder(
		message.EXPECT().GetByUuid(gomock.Any(), gomock.Any()).Return(&pb.GetMessageReply{Message: &pb.Message{Uuid: "test", Text: "text @System /version/ #tag1"}}, nil),
		repo.EXPECT().TouchGroupUpdatedAt(gomock.Any(), gomock.Any()).Return(nil),
		repo.EXPECT().ListGroupBot(gomock.Any(), gomock.Any()).Return([]*pb.Bot{
			{
				Name:       "System",
				Identifier: "system_bot",
			},
		}, nil),
		repo.EXPECT().GetGroupSetting(gomock.Any(), gomock.Any()).Return(nil, nil),
		repo.EXPECT().GetGroupBotSettingByGroup(gomock.Any(), gomock.Any()).Return(nil, nil),
		repo.EXPECT().GetBotsByText(gomock.Any(), gomock.Any()).Return(map[string]*pb.Bot{
			"System": {
				Name:       "System",
				Identifier: "system_bot",
			},
		}, nil),
		middle.EXPECT().SaveModelTag(gomock.Any(), gomock.Any()).Return(&pb.ModelTagReply{}, nil),
		message.EXPECT().Save(gomock.Any(), gomock.Any()).Return(&pb.MessageReply{}, nil),
	)

	s := NewChatbot(nil, bus, nil, repo, message, middle, bot, comp)

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
			args{context.Background(), &pb.ChatbotRequest{MessageId: 1}},
			&pb.ChatbotReply{State: true},
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

func TestChatbot_GetBot(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	bot, err := rulebot.CreateRuleBot(enum.Chatbot)
	if err != nil {
		t.Fatal(err)
	}

	repo := mock.NewMockChatbotRepository(ctl)

	item := pb.Bot{
		Id:        1,
		Uuid:      "1",
		Name:      "test",
		Avatar:    "test",
		CreatedAt: 0,
		UpdatedAt: 0,
	}
	gomock.InOrder(
		repo.EXPECT().GetGroupBot(gomock.Any(), gomock.Any(), gomock.Any()).Return(item, nil),
	)

	s := NewChatbot(nil, nil, nil, repo, nil, nil, bot, nil)

	type args struct {
		ctx     context.Context
		payload *pb.BotRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.BotReply
		wantErr bool
	}{
		{"case1", s, args{context.Background(), &pb.BotRequest{Bot: &pb.Bot{Uuid: "1"}}},
			&pb.BotReply{Bot: &item}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetBot(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chatbot.GetBot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chatbot.GetBot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatbot_GetBots(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	bot, err := rulebot.CreateRuleBot(enum.Chatbot)
	if err != nil {
		t.Fatal(err)
	}

	repo := mock.NewMockChatbotRepository(ctl)

	items := []*pb.Bot{
		{
			Id:        1,
			Uuid:      "1",
			Name:      "test",
			Avatar:    "test",
			CreatedAt: 0,
			UpdatedAt: 0,
		},
	}
	gomock.InOrder(
		repo.EXPECT().GetBotsByGroupUuid(gomock.Any(), gomock.Any()).Return(items, nil),
	)

	s := NewChatbot(nil, nil, nil, repo, nil, nil, bot, nil)

	type args struct {
		ctx     context.Context
		payload *pb.BotsRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.BotsReply
		wantErr bool
	}{
		{"case1", s, args{context.Background(), &pb.BotsRequest{}}, &pb.BotsReply{Bots: items}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetBots(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chatbot.GetBots() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chatbot.GetBots() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatbot_UpdateBotSetting(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	bot, err := rulebot.CreateRuleBot(enum.Chatbot)
	if err != nil {
		t.Fatal(err)
	}

	repo := mock.NewMockChatbotRepository(ctl)

	s := NewChatbot(nil, nil, nil, repo, nil, nil, bot, nil)

	type args struct {
		ctx     context.Context
		payload *pb.BotSettingRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{"case1", s, args{context.Background(), &pb.BotSettingRequest{}}, &pb.StateReply{State: true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UpdateBotSetting(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chatbot.UpdateBotSetting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chatbot.UpdateBotSetting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_GetGroups(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockChatbotRepository(ctl)
	message := mock.NewMockMessageSvcClient(ctl)
	gomock.InOrder(
		repo.EXPECT().GetGroupByName(gomock.Any(), gomock.Any(), gomock.Any()).Return(pb.Group{}, gorm.ErrRecordNotFound),
		repo.EXPECT().CreateGroup(gomock.Any(), gomock.Any()).Return(int64(1), nil),
		repo.EXPECT().GetBotsByUser(gomock.Any(), gomock.Any()).Return([]*pb.Bot{}, nil),
		repo.EXPECT().ListGroup(gomock.Any(), gomock.Any()).Return([]*pb.Group{{
			Id:     1,
			UserId: 1,
			Name:   "test",
		}}, nil),
		message.EXPECT().LastByGroup(gomock.Any(), gomock.Any()).Return(&pb.LastByGroupReply{Message: &pb.Message{
			SenderName: "test",
			Text:       "test",
		}}, nil),
	)

	s := NewChatbot(nil, nil, nil, repo, message, nil, nil, nil)

	type args struct {
		in0 context.Context
		in1 *pb.GroupRequest
	}
	tests := []struct {
		name    string
		m       *Chatbot
		args    args
		want    int
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.GroupRequest{Group: &pb.Group{UserId: 1}}},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.GetGroups(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Message.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && len(got.Groups) != tt.want {
				t.Errorf("Message.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_GetGroup(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockChatbotRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().GetGroupByUUID(gomock.Any(), gomock.Any()).Return(pb.Group{
			Uuid: "test",
		}, nil),
	)

	s := NewChatbot(nil, nil, nil, repo, nil, nil, nil, nil)

	type args struct {
		in0     context.Context
		payload *pb.GroupRequest
	}
	tests := []struct {
		name    string
		m       *Chatbot
		args    args
		want    *pb.GroupReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.GroupRequest{Group: &pb.Group{Uuid: "1"}}},
			&pb.GroupReply{
				Group: &pb.Group{
					Uuid: "test",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.GetGroup(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Message.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && (got.Group.Uuid != tt.want.Group.Uuid) {
				t.Errorf("Message.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_CreateGroup(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockChatbotRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().CreateGroup(gomock.Any(), gomock.Any()).Return(int64(1), nil),
		repo.EXPECT().CreateGroup(gomock.Any(), gomock.Any()).Return(int64(2), nil),
	)

	s := NewChatbot(nil, nil, nil, repo, nil, nil, nil, nil)

	type args struct {
		in0     context.Context
		payload *pb.GroupRequest
	}
	tests := []struct {
		name    string
		m       *Chatbot
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.GroupRequest{Group: &pb.Group{Name: "demo1", UserId: 1}}},
			&pb.StateReply{State: true},
			false,
		},
		{
			"case2",
			s,
			args{context.Background(), &pb.GroupRequest{Group: &pb.Group{Name: "demo2", UserId: 2}}},
			&pb.StateReply{State: true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.CreateGroup(tt.args.in0, tt.args.payload)
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

func TestNewChatbot(t *testing.T) {
	type args struct {
		logger log.Logger
		repo   repository.ChatbotRepository
		bot    *rulebot.RuleBot
	}
	tests := []struct {
		name string
		args args
		want *Chatbot
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChatbot(tt.args.logger, nil, nil, tt.args.repo, nil, nil, tt.args.bot, nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChatbot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatbot_GetGroups(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload *pb.GroupRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.GetGroupsReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetGroups(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chatbot.GetGroups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chatbot.GetGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatbot_CreateGroup(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload *pb.GroupRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateGroup(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chatbot.CreateGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chatbot.CreateGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatbot_GetGroup(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload *pb.GroupRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.GroupReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetGroup(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chatbot.GetGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chatbot.GetGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatbot_CreateGroupBot(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload *pb.GroupBotRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateGroupBot(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chatbot.CreateGroupBot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chatbot.CreateGroupBot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatbot_DeleteGroupBot(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload *pb.GroupBotRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.DeleteGroupBot(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chatbot.DeleteGroupBot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chatbot.DeleteGroupBot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatbot_UpdateGroupBotSetting(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload *pb.BotSettingRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UpdateGroupBotSetting(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chatbot.UpdateGroupBotSetting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chatbot.UpdateGroupBotSetting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatbot_UpdateGroupSetting(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload *pb.GroupSettingRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UpdateGroupSetting(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chatbot.UpdateGroupSetting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chatbot.UpdateGroupSetting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatbot_DeleteGroup(t *testing.T) {
	type args struct {
		ctx     context.Context
		payload *pb.GroupRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.DeleteGroup(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chatbot.DeleteGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chatbot.DeleteGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatbot_UpdateGroup(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *pb.GroupRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UpdateGroup(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chatbot.UpdateGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chatbot.UpdateGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatbot_SyntaxCheck(t *testing.T) {
	s := NewChatbot(nil, nil, nil, nil, nil, nil, nil, nil)

	type args struct {
		in0     context.Context
		payload *pb.WorkflowRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"ok syntax",
			s,
			args{context.Background(), &pb.WorkflowRequest{
				Type: string(enum.MessageTypeAction),
				Text: `get 1
echo 1
message "11"
`}},
			&pb.StateReply{State: true},
			false,
		},
		{
			"error syntax",
			s,
			args{context.Background(), &pb.WorkflowRequest{
				Type: string(enum.MessageTypeAction),
				Text: `get a a;`}},
			&pb.StateReply{State: false},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.SyntaxCheck(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workflow.SyntaxCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Workflow.SyntaxCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatbot_RunAction(t *testing.T) {
	s := NewChatbot(nil, nil, nil, nil, nil, nil, nil, nil)

	type args struct {
		in0     context.Context
		payload *pb.WorkflowRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.WorkflowReply
		wantErr bool
	}{
		{
			"run action",
			s,
			args{context.Background(), &pb.WorkflowRequest{Message: &pb.Message{Id: 1, Text: `echo "ok"`}}},
			&pb.WorkflowReply{Text: ""},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.RunAction(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workflow.RunAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Workflow.RunAction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatbot_WebhookTrigger(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mq, err := event.CreateRabbitmq(enum.Chatbot)
	if err != nil {
		t.Fatal(err)
	}
	bus := event.NewRabbitmqBus(mq, nil)

	message := mock.NewMockMessageSvcClient(ctl)
	repo := mock.NewMockChatbotRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().GetTriggerByFlag(gomock.Any(), gomock.Any(), gomock.Any()).Return(pb.Trigger{MessageId: 1, Secret: "test"}, nil),
		message.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(&pb.GetMessageReply{Message: &pb.Message{Id: 1, Text: ""}}, nil),
	)

	s := NewChatbot(nil, bus, nil, repo, message, nil, nil, nil)

	type args struct {
		ctx     context.Context
		payload *pb.TriggerRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.WorkflowReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TriggerRequest{Trigger: &pb.Trigger{Type: "webhook", Secret: "test"}}},
			&pb.WorkflowReply{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.WebhookTrigger(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workflow.WebhookTrigger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Workflow.WebhookTrigger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatbot_CronTrigger(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	rdb, err := vendors.CreateRedisClient(enum.User)
	if err != nil {
		t.Fatal(err)
	}

	mq, err := event.CreateRabbitmq(enum.Chatbot)
	if err != nil {
		t.Fatal(err)
	}
	bus := event.NewRabbitmqBus(mq, nil)

	messageID := rand.Int63()

	rdb.Set(context.Background(), fmt.Sprintf("workflow:cron:%d:time", messageID), time.Now().Add(-2*time.Minute).Format("2006-01-02 15:04:05"), redis.KeepTTL)

	message := mock.NewMockMessageSvcClient(ctl)
	repo := mock.NewMockChatbotRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().ListTriggersByType(gomock.Any(), "cron").Return([]*pb.Trigger{{MessageId: messageID, When: "* * * * *"}}, nil),
		message.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(&pb.GetMessageReply{Message: &pb.Message{Id: messageID, Text: ""}}, nil),
		repo.EXPECT().ListTriggersByType(gomock.Any(), "cron").Return([]*pb.Trigger{{MessageId: messageID, When: "* * * * *"}}, nil),
	)

	s := NewChatbot(nil, bus, rdb, repo, message, nil, nil, nil)

	type args struct {
		ctx context.Context
		in1 *pb.TriggerRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.WorkflowReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TriggerRequest{}},
			&pb.WorkflowReply{},
			false,
		},
		{
			"case2",
			s,
			args{context.Background(), &pb.TriggerRequest{}},
			&pb.WorkflowReply{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CronTrigger(tt.args.ctx, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workflow.CronTrigger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Workflow.CronTrigger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatbot_CreateTrigger(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockChatbotRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().CreateTrigger(gomock.Any(), gomock.Any()).Return(int64(1), nil),
		repo.EXPECT().GetTriggerByFlag(gomock.Any(), gomock.Any(), gomock.Any()).Return(pb.Trigger{}, nil),
		repo.EXPECT().CreateTrigger(gomock.Any(), gomock.Any()).Return(int64(1), nil),
	)

	s := NewChatbot(nil, nil, nil, repo, nil, nil, nil, nil)

	type args struct {
		in0     context.Context
		payload *pb.TriggerRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TriggerRequest{
				Trigger: &pb.Trigger{
					Kind: string(enum.MessageTypeAction),
					Type:      "webhook",
					MessageId: 1,
				},
				Info: &pb.TriggerInfo{
					MessageText: `
					cron "* * * * *"
					webhook "test"
					`,
				},
			}},
			&pb.StateReply{State: true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateTrigger(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workflow.CreateTrigger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Workflow.CreateTrigger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatbot_DeleteTrigger(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockChatbotRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().DeleteTriggerByMessageID(gomock.Any(), gomock.Any()).Return(nil),
		repo.EXPECT().DeleteTriggerByMessageID(gomock.Any(), gomock.Any()).Return(errors.New("not record")),
	)

	s := NewChatbot(nil, nil, nil, repo, nil, nil, nil, nil)

	type args struct {
		in0     context.Context
		payload *pb.TriggerRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TriggerRequest{Trigger: &pb.Trigger{MessageId: 1}}},
			&pb.StateReply{State: true},
			false,
		},
		{
			"case1",
			s,
			args{context.Background(), &pb.TriggerRequest{Trigger: &pb.Trigger{MessageId: 2}}},
			&pb.StateReply{State: false},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.DeleteTrigger(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workflow.DeleteTrigger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Workflow.DeleteTrigger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatbot_ActionDoc(t *testing.T) {
	s := NewChatbot(nil, nil, nil, nil, nil, nil, nil, nil)

	type args struct {
		in0     context.Context
		payload *pb.WorkflowRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.WorkflowReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.WorkflowRequest{}},
			&pb.WorkflowReply{Text: ""},
			false,
		},
		{
			"case1",
			s,
			args{context.Background(), &pb.WorkflowRequest{Text: "webhook"}},
			&pb.WorkflowReply{Text: "webhook [string] [string]?"},
			false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ActionDoc(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workflow.ActionDoc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if i == 0 {
				if got != nil && len(got.Text) == 0 {
					t.Errorf("Workflow.ActionDoc() = %v, want %v", got, tt.want)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Workflow.ActionDoc() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestChatbot_ListWebhook(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockChatbotRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().ListTriggersByType(gomock.Any(), gomock.Any()).Return([]*pb.Trigger{{Flag: "test1"}, {Flag: "test2"}}, nil),
	)

	s := NewChatbot(nil, nil, nil, repo, nil, nil, nil, nil)

	type args struct {
		in0     context.Context
		payload *pb.WorkflowRequest
	}
	tests := []struct {
		name    string
		s       *Chatbot
		args    args
		want    *pb.WebhooksReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.WorkflowRequest{}},
			&pb.WebhooksReply{Flag: []string{"test1", "test2"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ListWebhook(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workflow.ListWebhook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Workflow.ListWebhook() = %v, want %v", got, tt.want)
			}
		})
	}
}
