package event

type Subject string

const (
	EchoSubject              Subject = "echo"
	RoleChangeExpSubject     Subject = "role_change_exp"
	RoleChangeAttrSubject    Subject = "role_change_attr"
	MessageSendSubject       Subject = "message_send"
	MessagePushSubject       Subject = "message_push"
	MessageChannelSubject    Subject = "message_channel"
	ScriptRunSubject         Subject = "script_run"
	BotActionSubject         Subject = "bot_action"
	BotFormSubject           Subject = "bot_form"
	BotHandleSubject         Subject = "bot_handle"
	BotRegisterSubject       Subject = "bot_register"
	CronRegisterSubject      Subject = "cron_register"
	SubscribeRegisterSubject Subject = "subscribe_register"
)
