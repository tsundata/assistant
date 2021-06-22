package rulebot

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/version"
	"strings"
)

type Context struct {
	Conf   *config.AppConfig
	RDB    *redis.Client
	Logger *logger.Logger

	MessageClient   pb.MessageClient
	MiddleClient    pb.MiddleClient
	SubscribeClient pb.SubscribeClient
	WorkflowClient  pb.WorkflowClient
	StorageClient   pb.StorageClient
	TodoClient      pb.TodoClient
	UserClient      pb.UserClient
	NLPClient       pb.NLPClient
}

func (c Context) Message() pb.MessageClient {
	return c.MessageClient
}

func (c Context) Middle() pb.MiddleClient {
	return c.MiddleClient
}

func (c Context) Subscribe() pb.SubscribeClient {
	return c.SubscribeClient
}

func (c Context) Workflow() pb.WorkflowClient {
	return c.WorkflowClient
}

func (c Context) Storage() pb.StorageClient {
	return c.StorageClient
}

func (c Context) Todo() pb.TodoClient {
	return c.TodoClient
}

func (c Context) User() pb.UserClient {
	return c.UserClient
}

func (c Context) NLP() pb.NLPClient {
	return c.NLPClient
}

func (c Context) GetConfig() *config.AppConfig {
	return c.Conf
}

func (c Context) GetRedis() *redis.Client {
	return c.RDB
}

func (c Context) GetLogger() *logger.Logger {
	return c.Logger
}

type IContext interface {
	GetConfig() *config.AppConfig
	GetRedis() *redis.Client
	GetLogger() *logger.Logger
	Message() pb.MessageClient
	Middle() pb.MiddleClient
	Subscribe() pb.SubscribeClient
	Workflow() pb.WorkflowClient
	Storage() pb.StorageClient
	Todo() pb.TodoClient
	User() pb.UserClient
	NLP() pb.NLPClient
}

func NewContext(
	conf *config.AppConfig,
	rdb *redis.Client,
	logger *logger.Logger,

	messageClient pb.MessageClient,
	middleClient pb.MiddleClient,
	subscribeClient pb.SubscribeClient,
	workflowClient pb.WorkflowClient,
	storageClient pb.StorageClient,
	todoClient pb.TodoClient,
	userClient pb.UserClient,
	nlpClient pb.NLPClient,
) IContext {
	return Context{
		Conf:            conf,
		RDB:             rdb,
		Logger:          logger,
		MessageClient:   messageClient,
		MiddleClient:    middleClient,
		SubscribeClient: subscribeClient,
		WorkflowClient:  workflowClient,
		StorageClient:   storageClient,
		TodoClient:      todoClient,
		UserClient:      userClient,
		NLPClient:       nlpClient,
	}
}

type RuleBot struct {
	Ctx         IContext
	name        string
	providerIn  string
	providerOut []string
	rules       []RuleParser
}

func New(ctx IContext) *RuleBot {
	s := &RuleBot{
		name: ctx.GetConfig().Name,
		Ctx:  ctx,
	}

	return s
}

func (s *RuleBot) SetOptions(opts ...Option) {
	for _, opt := range opts {
		opt(s)
	}
}

func (s *RuleBot) Name() string {
	return s.name
}

func (s *RuleBot) Process(in string) *RuleBot {
	s.Ctx.GetLogger().Info("plugin process event")

	s.providerIn = in
	s.providerOut = []string{}
	if strings.ToLower(in) == "help" {
		helpMsg := fmt.Sprintf("available commands (v%s):\n", version.Version)
		for _, rule := range s.rules {
			helpMsg = fmt.Sprintln(helpMsg, rule.HelpMessage(s, in))
		}
		s.providerOut = append(s.providerOut, helpMsg)
		return s
	}

	defer func() {
		if r := recover(); r != nil {
			s.Ctx.GetLogger().Error(fmt.Errorf("panic recovered when parsing message: %#v. Panic: %v", in, r))
		}
	}()
	for _, rule := range s.rules {
		responses := rule.ParseMessage(s, in)
		s.providerOut = append(s.providerOut, responses...)
	}
	return s
}

func (s *RuleBot) MessageProviderOut() []string {
	return s.providerOut
}

type Option func(*RuleBot)

type RuleParser interface {
	Name() string
	Boot(*RuleBot)
	ParseMessage(*RuleBot, string) []string
	HelpMessage(*RuleBot, string) string
}

func RegisterRuleset(rule RuleParser) Option {
	return func(s *RuleBot) {
		s.Ctx.GetLogger().Info(fmt.Sprintf("registering ruleset %T", rule))
		rule.Boot(s)
		s.rules = append(s.rules, rule)
	}
}

var ProviderSet = wire.NewSet(NewContext, New)
