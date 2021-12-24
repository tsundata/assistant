package service

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/repository"
	"github.com/tsundata/assistant/internal/app/chatbot/rule"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

type Chatbot struct {
	logger  log.Logger
	bot     *rulebot.RuleBot
	repo    repository.ChatbotRepository
	message pb.MessageSvcClient
	middle  pb.MiddleSvcClient
	todo    pb.TodoSvcClient
}

func NewChatbot(
	logger log.Logger,
	repo repository.ChatbotRepository,
	message pb.MessageSvcClient,
	middle pb.MiddleSvcClient,
	todo pb.TodoSvcClient,
	bot *rulebot.RuleBot) *Chatbot {
	return &Chatbot{
		logger:  logger,
		bot:     bot,
		repo:    repo,
		message: message,
		middle:  middle,
		todo:    todo,
	}
}

func (s *Chatbot) Handle(ctx context.Context, payload *pb.ChatbotRequest) (*pb.ChatbotReply, error) {
	reply, err := s.message.Get(ctx, &pb.MessageRequest{Message: &pb.Message{Id: payload.MessageId}})
	if err != nil {
		return nil, err
	}
	s.bot.SetOptions(rule.Options...)
	s.bot.Process(ctx, reply.Message.Text).MessageProviderOut()
	return &pb.ChatbotReply{
		State: true,
	}, nil
}

func (s *Chatbot) Register(ctx context.Context, request *pb.BotRequest) (*pb.StateReply, error) {
	bot, err := s.repo.GetByIdentifier(ctx, request.Bot.Identifier)
	if err != nil {
		return nil, err
	}
	if bot.Id > 0 {
		return &pb.StateReply{State: true}, nil
	}
	_, err = s.repo.Create(ctx, request.Bot)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) GetBot(ctx context.Context, payload *pb.BotRequest) (*pb.BotReply, error) {
	bot, err := s.repo.GetByUUID(ctx, payload.Bot.GetUuid())
	if err != nil {
		return nil, err
	}
	return &pb.BotReply{Bot: &bot}, nil
}

func (s *Chatbot) GetBots(ctx context.Context, _ *pb.BotsRequest) (*pb.BotsReply, error) {
	bots, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.BotsReply{Bots: bots}, nil
}

func (s *Chatbot) UpdateBotSetting(_ context.Context, _ *pb.BotSettingRequest) (*pb.StateReply, error) {
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) GetGroups(ctx context.Context, payload *pb.GroupRequest) (*pb.GroupsReply, error) {
	groups, err := s.repo.ListGroup(ctx, payload.Group.GetUserId())
	if err != nil {
		return nil, err
	}
	return &pb.GroupsReply{Groups: groups}, nil
}

func (s *Chatbot) CreateGroup(ctx context.Context, payload *pb.GroupRequest) (*pb.StateReply, error) {
	_, err := s.repo.CreateGroup(ctx, payload.Group)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) GetGroup(ctx context.Context, payload *pb.GroupRequest) (*pb.GroupReply, error) {
	group, err := s.repo.GetGroupByUUID(ctx, payload.Group.GetUuid())
	if err != nil {
		return nil, err
	}
	return &pb.GroupReply{Group: &group}, nil
}

func (s *Chatbot) CreateGroupBot(ctx context.Context, payload *pb.GroupBotRequest) (*pb.StateReply, error) {
	err := s.repo.CreateGroupBot(ctx, payload.GroupId, payload.Bot)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) DeleteGroupBot(ctx context.Context, payload *pb.GroupBotRequest) (*pb.StateReply, error) {
	err := s.repo.DeleteGroupBot(ctx, payload.GroupId, payload.Bot.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) UpdateGroupBotSetting(ctx context.Context, payload *pb.BotSettingRequest) (*pb.StateReply, error) {
	err := s.repo.UpdateGroupBotSetting(ctx, payload.GroupId, payload.BotId, payload.Kvs)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) UpdateGroupSetting(ctx context.Context, payload *pb.GroupSettingRequest) (*pb.StateReply, error) {
	err := s.repo.UpdateGroupSetting(ctx, payload.GroupId, payload.Kvs)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) DeleteGroup(ctx context.Context, payload *pb.GroupRequest) (*pb.StateReply, error) {
	err := s.repo.DeleteGroup(ctx, payload.Group.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) UpdateGroup(ctx context.Context, payload *pb.GroupRequest) (*pb.StateReply, error) {
	err := s.repo.UpdateGroup(ctx, payload.Group)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Chatbot) GetGroupBotSetting(ctx context.Context, request *pb.BotSettingRequest) (*pb.BotSettingReply, error) {
	kv, err := s.repo.GetGroupBotSetting(ctx, request.GroupId, request.BotId)
	if err != nil {
		return nil, err
	}
	return &pb.BotSettingReply{Kvs: kv}, nil
}

func (s *Chatbot) GetGroupSetting(ctx context.Context, request *pb.GroupSettingRequest) (*pb.GroupSettingReply, error) {
	kv, err := s.repo.GetGroupSetting(ctx, request.GroupId)
	if err != nil {
		return nil, err
	}
	return &pb.GroupSettingReply{Kvs: kv}, nil
}
