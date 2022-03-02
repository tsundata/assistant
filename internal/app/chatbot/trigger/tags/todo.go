package tags

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/trigger/ctx"
	"github.com/tsundata/assistant/internal/pkg/event"
)

type Todo struct{}

func NewTodo() *Todo {
	return &Todo{}
}

func (t *Todo) Handle(ctx context.Context, comp *ctx.Component, text string) {
	// create
	reply, err := comp.Todo.CreateTodo(ctx, &pb.TodoRequest{
		Todo: &pb.Todo{Content: text},
	})
	if err != nil {
		comp.Logger.Error(err)
		return
	}
	if !reply.GetState() {
		return
	}

	// send message
	err = comp.Bus.Publish(ctx, enum.Message, event.MessageSendSubject, pb.Message{Text: "Created Todo success"})
	if err != nil {
		comp.Logger.Error(err)
		return
	}
}
