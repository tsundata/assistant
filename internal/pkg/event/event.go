package event

type Subject string

const (
	EchoSubject           Subject = "echo"
	RoleChangeExpSubject  Subject = "role_change_exp"
	RoleChangeAttrSubject Subject = "role_change_attr"
	MessageSendSubject    Subject = "message_send"
	WorkflowRunSubject    Subject = "workflow_run"
	MessageTriggerSubject Subject = "message_trigger"
)
