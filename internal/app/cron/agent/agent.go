package agent

import "github.com/tsundata/assistant/internal/pkg/rulebot"

type Agenter interface {
	Fetch(*rulebot.RuleBot) []string
}

func Deduplication(name string, data []string) []string {
	return data
}
