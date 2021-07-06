package event

type Subject string

const (
	EchoSubject           Subject = "echo"
	ChangeExpSubject      Subject = "change_exp"
	ChangeAttrSubject     Subject = "change_attr"
	SendMessageSubject    Subject = "send_message"
	RunWorkflowSubject    Subject = "run_workflow"
	MessageTriggerSubject Subject = "message_trigger"
)
