package enum

const (
	SystemBot  = "system_bot"
	FinanceBot = "finance_bot"
	OkrBot     = "okr_bot"
	TodoBot    = "todo_bot"
	GithubBot  = "github_bot"
)

const (
	TriggerEnable = iota + 1
	TriggerDisable
)

const (
	TriggerWebhookType = "webhook"
	TriggerCronType    = "cron"
	TriggerWatchType   = "watch"
)

const (
	ValueSumMode  = "sum"
	ValueLastMode = "last"
	ValueAvgMode  = "avg"
	ValueMaxMode  = "max"
)

const DefaultGroupName = "System"
