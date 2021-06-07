package event

type Subject string

const (
	EchoSubject    Subject = "echo"
	IncrExpSubject Subject = "incr_exp"
	DecrExpSubject Subject = "decr_exp"
)
