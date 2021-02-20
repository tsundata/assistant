package agent

import "github.com/tsundata/assistant/internal/pkg/rulebot"

type Pocket struct {
}

func NewPocket() *Pocket {
	return &Pocket{}
}

func (a *Pocket) Fetch(_ *rulebot.RuleBot) []string {
	return nil
}
