package command

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"go.uber.org/zap"
	"strings"
	"unicode"
)

type Rule struct {
	Define string
	Help   string
	Parse  func(context.Context, component.Component, []*Token) []pb.MsgPayload
}

type Ruleset struct {
	rules []Rule
}

func (r Ruleset) Help(in string) string {
	if strings.ToLower(in) == "help" {
		var helpMsg string
		for _, rule := range r.rules {
			helpMsg = fmt.Sprintf("%s%s%s%s\n", helpMsg, rule.Define, " :: ", rule.Help)
		}
		return strings.TrimLeftFunc(helpMsg, unicode.IsSpace)
	}
	return ""
}

func (r Ruleset) ProcessCommand(ctx context.Context, comp component.Component, in string) ([]pb.MsgPayload, error) {
	var result []pb.MsgPayload
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
			result = ret
		}
	}

	return result, nil
}

func New(rules []Rule) *Ruleset {
	return &Ruleset{
		rules: rules,
	}
}
