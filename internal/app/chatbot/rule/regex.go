package rule

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"go.uber.org/zap"
	"strings"
	"unicode"
)

type Rule struct {
	Define string
	Help   string
	Parse  func(context.Context, rulebot.IComponent, []*Token) []string
}

type regexRuleset struct {
	rules []Rule
}

func (r regexRuleset) Name() string {
	return "Regex Ruleset"
}

func (r regexRuleset) Boot(_ *rulebot.RuleBot) {}

func (r regexRuleset) HelpRule(_ *rulebot.RuleBot, _ string) string {
	var helpMsg string
	for _, rule := range r.rules {
		helpMsg = fmt.Sprintln(helpMsg, rule.Define, " : ", rule.Help)
	}
	return strings.TrimLeftFunc(helpMsg, unicode.IsSpace)
}

func (r regexRuleset) ParseRule(ctx context.Context, b *rulebot.RuleBot, in string) []string {
	for _, rule := range r.rules {
		tokens, err := ParseCommand(in)
		if err != nil {
			if b.Comp.GetLogger() != nil {
				b.Comp.GetLogger().Error(err, zap.Any("rule", in))
			}
		}
		check, err := SyntaxCheck(rule.Define, tokens)
		if err != nil {
			if b.Comp.GetLogger() != nil {
				b.Comp.GetLogger().Error(err, zap.Any("rule", in))
			}
		}
		if !check {
			continue
		}

		if ret := rule.Parse(ctx, b.Comp, tokens); len(ret) > 0 {
			return ret
		}
	}

	return []string{}
}

func New(rules []Rule) *regexRuleset {
	return &regexRuleset{
		rules: rules,
	}
}
