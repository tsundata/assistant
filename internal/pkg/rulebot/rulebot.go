package rulebot

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/version"
	"strings"
	"sync"
)

type Component struct {
	Conf   *config.AppConfig
	RDB    *redis.Client
	Logger log.Logger

	MessageClient     pb.MessageSvcClient
	MiddleClient      pb.MiddleSvcClient
	WorkflowSvcClient pb.WorkflowSvcClient
	StorageClient     pb.StorageSvcClient
	TodoClient        pb.TodoSvcClient
	UserClient        pb.UserSvcClient
	NLPClient         pb.NLPSvcClient
}

func (c Component) Message() pb.MessageSvcClient {
	return c.MessageClient
}

func (c Component) Middle() pb.MiddleSvcClient {
	return c.MiddleClient
}

func (c Component) Workflow() pb.WorkflowSvcClient {
	return c.WorkflowSvcClient
}

func (c Component) Storage() pb.StorageSvcClient {
	return c.StorageClient
}

func (c Component) Todo() pb.TodoSvcClient {
	return c.TodoClient
}

func (c Component) User() pb.UserSvcClient {
	return c.UserClient
}

func (c Component) NLP() pb.NLPSvcClient {
	return c.NLPClient
}

func (c Component) GetConfig() *config.AppConfig {
	return c.Conf
}

func (c Component) GetRedis() *redis.Client {
	return c.RDB
}

func (c Component) GetLogger() log.Logger {
	return c.Logger
}

type IComponent interface {
	GetConfig() *config.AppConfig
	GetRedis() *redis.Client
	GetLogger() log.Logger
	Message() pb.MessageSvcClient
	Middle() pb.MiddleSvcClient
	Workflow() pb.WorkflowSvcClient
	Storage() pb.StorageSvcClient
	Todo() pb.TodoSvcClient
	User() pb.UserSvcClient
	NLP() pb.NLPSvcClient
}

func NewComponent(
	conf *config.AppConfig,
	rdb *redis.Client,
	logger log.Logger,

	messageClient pb.MessageSvcClient,
	middleClient pb.MiddleSvcClient,
	workflowClient pb.WorkflowSvcClient,
	storageClient pb.StorageSvcClient,
	todoClient pb.TodoSvcClient,
	userClient pb.UserSvcClient,
	nlpClient pb.NLPSvcClient,
) IComponent {
	return Component{
		Conf:              conf,
		RDB:               rdb,
		Logger:            logger,
		MessageClient:     messageClient,
		MiddleClient:      middleClient,
		WorkflowSvcClient: workflowClient,
		StorageClient:     storageClient,
		TodoClient:        todoClient,
		UserClient:        userClient,
		NLPClient:         nlpClient,
	}
}

type RuleBot struct {
	onceOptions sync.Once
	Comp        IComponent
	name        string
	providerIn  string
	providerOut []string
	rules       []RuleParser
}

func New(comp IComponent) *RuleBot {
	name := ""
	if comp != nil {
		name = comp.GetConfig().Name
	}
	s := &RuleBot{
		name: name,
		Comp: comp,
	}

	return s
}

func (s *RuleBot) SetOptions(opts ...Option) {
	s.onceOptions.Do(func() {
		for _, opt := range opts {
			opt(s)
		}
	})
}

func (s *RuleBot) Name() string {
	return s.name
}

func (s *RuleBot) Process(ctx context.Context, in string) *RuleBot {
	if s.Comp != nil && s.Comp.GetLogger() != nil {
		s.Comp.GetLogger().Debug("plugin process event")
	}

	s.providerIn = in
	s.providerOut = []string{}
	if strings.ToLower(in) == "help" {
		helpMsg := fmt.Sprintf("available commands (v%s):\n", version.Version)
		for _, rule := range s.rules {
			helpMsg = fmt.Sprintln(helpMsg, rule.HelpRule(s, in))
		}
		s.providerOut = append(s.providerOut, helpMsg)
		return s
	}

	defer func() {
		if r := recover(); r != nil {
			s.Comp.GetLogger().Error(fmt.Errorf("panic recovered when parsing message: %#v. Panic: %v", in, r))
		}
	}()
	for _, rule := range s.rules {
		responses := rule.ParseRule(ctx, s, in)
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
	ParseRule(context.Context, *RuleBot, string) []string
	HelpRule(*RuleBot, string) string
}

func RegisterRuleset(rule RuleParser) Option {
	return func(s *RuleBot) {
		if s.Comp != nil && s.Comp.GetLogger() != nil {
			s.Comp.GetLogger().Info(fmt.Sprintf("registering ruleset %T", rule))
		}
		rule.Boot(s)
		s.rules = append(s.rules, rule)
	}
}

var ProviderSet = wire.NewSet(NewComponent, New)
