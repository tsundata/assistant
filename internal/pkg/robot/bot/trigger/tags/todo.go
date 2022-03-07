package tags

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

type Todo struct{}

func NewTodo() *Todo {
	return &Todo{}
}

func (t *Todo) Handle(ctx context.Context, comp component.Component, text string) {
	//// create
	//reply, err := comp.Todo.CreateTodo(ctx, &pb.TodoRequest{
	//	Todo: &pb.Todo{Content: text},
	//})
	//if err != nil {
	//	comp.Logger.Error(err)
	//	return
	//}
	//if !reply.GetState() {
	//	return
	//}
	//
	//// send message
	//err = comp.Bus.Publish(ctx, enum.Message, event.MessageSendSubject, pb.Message{Text: "Created Todo success"})
	//if err != nil {
	//	comp.Logger.Error(err)
	//	return
	//}
}
