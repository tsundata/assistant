package rule

import (
	"github.com/tsundata/assistant/internal/app/cron/agent"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

var rules = []Rule{
	{
		Name: "pocket",
		When: "* * * * *",
		Action: func(ctx rulebot.IContext) []result.Result {
			return agent.FetchPocket(ctx)
		},
	},
	{
		Name: "github_starred",
		When: "* * * * *",
		Action: func(ctx rulebot.IContext) []result.Result {
			return agent.FetchGithubStarred(ctx)
		},
	},
	{
		Name: "backup",
		When: "0 0 * * *",
		Action: func(ctx rulebot.IContext) []result.Result {
			return agent.Backup(ctx)
		},
	},
	{
		Name: "workflow_cron",
		When: "* * * * *",
		Action: func(ctx rulebot.IContext) []result.Result {
			return agent.WorkflowCron(ctx)
		},
	},
	{
		Name: "cloudflare_report",
		When: "0 0 * * 0",
		Action: func(ctx rulebot.IContext) []result.Result {
			return agent.DomainAnalyticsReport(ctx)
		},
	},
}

var Options = []rulebot.Option{
	rulebot.RegisterRuleset(New(rules)),
}
