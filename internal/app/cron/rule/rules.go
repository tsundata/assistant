package rule

import (
	"context"
	"github.com/tsundata/assistant/internal/app/cron/agent"
	"github.com/tsundata/assistant/internal/app/cron/pipeline/result"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
)

var rules = []Rule{
	{
		Name: "pocket",
		When: "* * * * *",
		Action: func(ctx context.Context, comp component.Component) []result.Result {
			return agent.FetchPocket(ctx, comp)
		},
	},
	{
		Name: "github_starred",
		When: "* * * * *",
		Action: func(ctx context.Context, comp component.Component) []result.Result {
			return agent.FetchGithubStarred(ctx, comp)
		},
	},
	{
		Name: "github_stargazers",
		When: "* * * * *",
		Action: func(ctx context.Context, comp component.Component) []result.Result {
			return agent.FetchGithubStargazers(ctx, comp)
		},
	},
	{
		Name: "backup",
		When: "0 0 * * *",
		Action: func(ctx context.Context, comp component.Component) []result.Result {
			return agent.Backup(ctx, comp)
		},
	},
	{
		Name: "script_cron",
		When: "* * * * *",
		Action: func(ctx context.Context, comp component.Component) []result.Result {
			return agent.ScriptCron(ctx, comp)
		},
	},
	{
		Name: "script_watch",
		When: "* * * * *",
		Action: func(ctx context.Context, comp component.Component) []result.Result {
			return agent.ScriptWatch(ctx, comp)
		},
	},
	{
		Name: "cloudflare_report",
		When: "0 0 * * 0",
		Action: func(ctx context.Context, comp component.Component) []result.Result {
			return agent.DomainAnalyticsReport(ctx, comp)
		},
	},
	{
		Name: "todo_remind",
		When: "* * * * *",
		Action: func(ctx context.Context, comp component.Component) []result.Result {
			return agent.TodoRemind(ctx, comp)
		},
	},
	{
		Name: "cloudcone_billing",
		When: "50 23 * * SUN",
		Action: func(ctx context.Context, comp component.Component) []result.Result {
			return agent.CloudconeWeeklyBilling(ctx, comp)
		},
	},
	{
		Name: "search_metadata",
		When: "* */1 * * *",
		Action: func(ctx context.Context, comp component.Component) []result.Result {
			return agent.SearchMetadata(ctx, comp)
		},
	},
}

var Options = []rulebot.Option{
	rulebot.RegisterRuleset(newCronRuleset(rules)),
}
