package service

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/rule"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

type Chatbot struct {
	logger log.Logger
	bot    *rulebot.RuleBot
	middle pb.MiddleClient
	todo   pb.TodoSvcClient
}

func NewChatbot(
	logger log.Logger,
	middle pb.MiddleClient,
	todo pb.TodoSvcClient,
	bot *rulebot.RuleBot) *Chatbot {
	return &Chatbot{
		logger:   logger,
		bot:      bot,
		middle:   middle,
		todo:     todo,
	}
}

func (s *Chatbot) Handle(ctx context.Context, payload *pb.ChatbotRequest) (*pb.ChatbotReply, error) {
	s.bot.SetOptions(rule.Options...)
	out := s.bot.Process(ctx, payload.GetText()).MessageProviderOut()
	return &pb.ChatbotReply{
		Text: out,
	}, nil
}
