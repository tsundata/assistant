package tags

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/message/trigger/ctx"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/model"
)

type Todo struct{}

func NewTodo() *Todo {
	return &Todo{}
}

func (t *Todo) Handle(ctx *ctx.Context, text string) {
	// create
	reply, err := ctx.Todo.CreateTodo(context.Background(), &pb.TodoRequest{
		Content: text,
	})
	if err != nil {
		ctx.Logger.Error(err)
		return
	}
	if !reply.GetState() {
		return
	}

	// send message
	err = ctx.Bus.Publish(event.SendMessageSubject, model.Message{Text: "Created Todo success"})
	if err != nil {
		ctx.Logger.Error(err)
		return
	}
}
