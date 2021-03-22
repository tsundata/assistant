package rules

import (
	"github.com/tsundata/assistant/internal/app/cron/agent"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

var rules = []Rule{
	{
		Name: "pocket",
		When: "* * * * *",
		Action: func(b *rulebot.RuleBot) []result.Result {
			return agent.FetchPocket(b)
		},
	},
	{
		Name: "github_starred",
		When: "* * * * *",
		Action: func(b *rulebot.RuleBot) []result.Result {
			return agent.FetchGithubStarred(b)
		},
	},
	{
		Name: "backup",
		When: "0 0 * * *",
		Action: func(b *rulebot.RuleBot) []result.Result {
			return agent.Backup(b)
		},
	},
	{
		Name: "workflow_cron",
		When: "* * * * *",
		Action: func(b *rulebot.RuleBot) []result.Result {
			return agent.WorkflowCron(b)
		},
	},
}

var Options = []rulebot.Option{
	rulebot.RegisterRuleset(New(rules)),
}
