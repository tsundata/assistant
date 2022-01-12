package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/repository"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/mock"
	"gorm.io/gorm"
	"reflect"
	"testing"
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
	message := mock.NewMockMessageSvcClient(ctl)

	gomock.InOrder(
		message.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&pb.GetMessageReply{Message: &pb.MessageItem{Uuid: "test"}}, nil),
	)

	s := NewChatbot(nil, nil, message, middle, todo, bot)

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
	middle := mock.NewMockMiddleSvcClient(ctl)
	todo := mock.NewMockTodoSvcClient(ctl)

	item := pb.Bot{
		Id:        1,
		Uuid:      "1",
		Name:      "test",
		Avatar:    "test",
		CreatedAt: 0,
		UpdatedAt: 0,
	}
	gomock.InOrder(
		repo.EXPECT().GetByUUID(gomock.Any(), gomock.Any()).Return(item, nil),
	)

	s := NewChatbot(nil, repo, nil, middle, todo, bot)

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
	middle := mock.NewMockMiddleSvcClient(ctl)
	todo := mock.NewMockTodoSvcClient(ctl)

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
		repo.EXPECT().List(gomock.Any()).Return(items, nil),
	)

	s := NewChatbot(nil, repo, nil, middle, todo, bot)

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
	middle := mock.NewMockMiddleSvcClient(ctl)
	todo := mock.NewMockTodoSvcClient(ctl)

	s := NewChatbot(nil, repo, nil, middle, todo, bot)

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
		message.EXPECT().LastByGroup(gomock.Any(), gomock.Any()).Return(&pb.LastByGroupReply{Message: &pb.MessageItem{
			Sender:  "test",
			Message: "test",
		}}, nil),
	)

	s := NewChatbot(nil, repo, message, nil, nil, nil)

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

	s := NewChatbot(nil, repo, nil, nil, nil, nil)

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

	s := NewChatbot(nil, repo, nil, nil, nil, nil)

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
		middle pb.MiddleSvcClient
		todo   pb.TodoSvcClient
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
			if got := NewChatbot(tt.args.logger, tt.args.repo, nil, tt.args.middle, tt.args.todo, tt.args.bot); !reflect.DeepEqual(got, tt.want) {
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
