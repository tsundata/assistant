package rules

import (
	"github.com/tsundata/assistant/internal/app/cron/agent"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"time"
)

var rules = []Rule{
	{
		Name: "heartbeat",
		When: "0 0 * * *",
		Action: func(b *rulebot.RuleBot) []string {
			return []string{
				"Plugin Cron Heartbeat: " + time.Now().String(),
			}
		},
	},
	{
		Name: "pocket",
		When: "* * * * *",
		Action: func(b *rulebot.RuleBot) []string {
			return agent.FetchPocket(b)
		},
	},
	{
		Name: "github_starred",
		When: "* * * * *",
		Action: func(b *rulebot.RuleBot) []string {
			return agent.FetchGithubStarred(b)
		},
	},
}

var Options = []rulebot.Option{
	rulebot.RegisterRuleset(New(rules)),
}
