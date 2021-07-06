package service

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/rule"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

type Chatbot struct {
	logger   *logger.Logger
	bot      *rulebot.RuleBot
	middle   pb.MiddleClient
	todo     pb.TodoClient
}

func NewChatbot(
	logger *logger.Logger,
	middle pb.MiddleClient,
	todo pb.TodoClient,
	bot *rulebot.RuleBot) *Chatbot {
	return &Chatbot{
		logger:   logger,
		bot:      bot,
		middle:   middle,
		todo:     todo,
	}
}

func (s *Chatbot) Handle(_ context.Context, payload *pb.ChatbotRequest) (*pb.ChatbotReply, error) {
	s.bot.SetOptions(rule.Options...)
	out := s.bot.Process(payload.GetText()).MessageProviderOut()
	return &pb.ChatbotReply{
		Text: out,
	}, nil
}
