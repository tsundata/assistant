package rulebot

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/version"
	"log"
	"strings"
)

type RuleBot struct {
	name        string
	providerIn  string
	providerOut []string
	rules       []RuleParser

	RDB        *redis.Client
	SubClient  pb.SubscribeClient
	MidClient  pb.MiddleClient
	MsgClient  pb.MessageClient
	WfClient   pb.WorkflowClient
	TaskClient pb.TaskClient
}

func New(name string, RDB *redis.Client,
	SubClient pb.SubscribeClient, MidClient pb.MiddleClient, MsgClient pb.MessageClient,
	WfClient pb.WorkflowClient, TaskClient pb.TaskClient,
	opts ...Option) *RuleBot {

	s := &RuleBot{
		name: name,
	}

	s.RDB = RDB
	s.SubClient = SubClient
	s.MidClient = MidClient
	s.MsgClient = MsgClient
	s.WfClient = WfClient
	s.TaskClient = TaskClient

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *RuleBot) Name() string {
	return s.name
}

func (s *RuleBot) Process(in string) *RuleBot {
	log.Println("plugin process event")

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
			log.Printf("panic recovered when parsing message: %#v. Panic: %v", in, r)
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
		log.Printf("registering ruleset %T", rule)
		rule.Boot(s)
		s.rules = append(s.rules, rule)
	}
}
