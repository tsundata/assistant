package rules

import (
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"time"
)

var rules = map[string]Rule{
	"heartbeat": {
		When: "0 0 * * *",
		Action: func(b *rulebot.RuleBot) []string {
			return []string{
				"Plugin Cron Heartbeat: " + time.Now().String(),
			}
		},
	},
	"pocket": {
		When: "* * * * *",
		Action: func(b *rulebot.RuleBot) []string {
			return []string{
				time.Now().String(),
			}
		},
	},
}

var Options = []rulebot.Option{
	rulebot.RegisterRuleset(New(rules)),
}
