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
	logger log.Logger
	bot    *rulebot.RuleBot
	repo   repository.ChatbotRepository
	middle pb.MiddleSvcClient
	todo   pb.TodoSvcClient
}

func NewChatbot(
	logger log.Logger,
	repo repository.ChatbotRepository,
	middle pb.MiddleSvcClient,
	todo pb.TodoSvcClient,
	bot *rulebot.RuleBot) *Chatbot {
	return &Chatbot{
		logger: logger,
		bot:    bot,
		repo:   repo,
		middle: middle,
		todo:   todo,
	}
}

func (s *Chatbot) Handle(ctx context.Context, payload *pb.ChatbotRequest) (*pb.ChatbotReply, error) {
	s.bot.SetOptions(rule.Options...)
	out := s.bot.Process(ctx, payload.GetText()).MessageProviderOut()
	return &pb.ChatbotReply{
		Text: out,
	}, nil
}

func (s *Chatbot) GetBot(ctx context.Context, payload *pb.BotRequest) (*pb.BotReply, error) {
	bot, err := s.repo.GetByUUID(ctx, payload.Bot.GetUuid())
	if err != nil {
		return nil, err
	}
	return &pb.BotReply{Bot: bot}, nil
}

func (s *Chatbot) GetBots(ctx context.Context, _ *pb.BotRequest) (*pb.BotsReply, error) {
	bots, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.BotsReply{Bots: bots}, nil
}

func (s *Chatbot) UpdateBotSetting(_ context.Context, _ *pb.BotSettingRequest) (*pb.StateReply, error) {
	return &pb.StateReply{State: true}, nil
}
