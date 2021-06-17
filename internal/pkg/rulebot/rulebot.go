package rulebot

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"github.com/tsundata/assistant/internal/pkg/version"
	"strings"
)

type Context struct {
	Conf   *config.AppConfig
	RDB    *redis.Client
	Client *rpc.Client
	Logger *logger.Logger
}

type RuleBot struct {
	Ctx         *Context
	name        string
	providerIn  string
	providerOut []string
	rules       []RuleParser
}

func New(ctx *Context, opts ...Option) *RuleBot {
	s := &RuleBot{
		name: ctx.Conf.Name,
		Ctx:  ctx,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *RuleBot) Name() string {
	return s.name
}

func (s *RuleBot) Process(in string) *RuleBot {
	s.Ctx.Logger.Info("plugin process event")

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
			s.Ctx.Logger.Error(fmt.Errorf("panic recovered when parsing message: %#v. Panic: %v", in, r))
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
		s.Ctx.Logger.Info(fmt.Sprintf("registering ruleset %T", rule))
		rule.Boot(s)
		s.rules = append(s.rules, rule)
	}
}
