package tags

type Issue struct {}

func NewIssue() *Issue {
	return &Issue{}
}

func (t *Issue) Handle(text string) {
	panic("implement me")
}
