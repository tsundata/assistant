package rulebot

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/version"
	"strings"
	"sync"
)

type RuleBot struct {
	onceOptions sync.Once
	Comp        component.Component
	name        string
	providerIn  string
	providerOut []string
	rules       []RuleParser
}

func New(comp component.Component) *RuleBot {
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

var ProviderSet = wire.NewSet(New)
