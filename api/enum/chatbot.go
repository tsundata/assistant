package enum

const (
	SystemBot  = "system_bot"
	FinanceBot = "finance_bot"
	OrgBot     = "org_bot"
	TodoBot    = "todo_bot"
)

const (
	TriggerEnable = iota + 1
	TriggerDisable
)

const (
	TriggerWebhookType = "webhook"
	TriggerCronType    = "cron"
)
