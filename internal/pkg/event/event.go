package event

type Subject string

const (
	EchoSubject        Subject = "echo"
	ChangeExpSubject   Subject = "change_exp"
	SendMessageSubject Subject = "send_message"
	RunWorkflowSubject Subject = "run_workflow"
)
