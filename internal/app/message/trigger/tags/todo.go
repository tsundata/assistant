package tags

type Todo struct {}

func NewTodo() *Todo {
	return &Todo{}
}

func (t *Todo) Handle(text string) {
	panic("implement me")
}
