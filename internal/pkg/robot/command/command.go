package command

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"strings"
	"unicode"
)

type Rule struct {
	Define string
	Help   string
	Parse  func(context.Context, Component, []*Token) []string
}

type Ruleset struct {
	rules []Rule
}

func (r Ruleset) Help(in string) string {
	if strings.ToLower(in) == "help" {
		var helpMsg string
		for _, rule := range r.rules {
			helpMsg = fmt.Sprintln(helpMsg, rule.Define, " : ", rule.Help)
		}
		return strings.TrimLeftFunc(helpMsg, unicode.IsSpace)
	}
	return ""
}

func (r Ruleset) ParseCommand(ctx context.Context, comp Component, in string) ([]string, error) {
	for _, rule := range r.rules {
		tokens, err := ParseCommand(in)
		if err != nil {
			if comp.GetLogger() != nil {
				comp.GetLogger().Error(err, zap.Any("rule", in))
			}
		}
		check, err := SyntaxCheck(rule.Define, tokens)
		if err != nil {
			if comp.GetLogger() != nil {
				comp.GetLogger().Error(err, zap.Any("rule", in))
			}
		}
		if !check {
			continue
		}

		if ret := rule.Parse(ctx, comp, tokens); len(ret) > 0 {
			return ret, nil
		}
	}

	return []string{}, nil
}

func New(rules []Rule) *Ruleset {
	return &Ruleset{
		rules: rules,
	}
}
