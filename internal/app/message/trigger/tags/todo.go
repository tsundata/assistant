package tags

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/message/trigger/ctx"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc/rpcclient"
)

type Todo struct{}

func NewTodo() *Todo {
	return &Todo{}
}

func (t *Todo) Handle(ctx *ctx.Context, text string) {
	// create
	reply, err := rpcclient.GetTodoClient(ctx.Client).CreateTodo(context.Background(), &pb.TodoRequest{
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
	_, err = rpcclient.GetMessageClient(ctx.Client).Send(context.Background(), &pb.MessageRequest{Text: "Created Todo success"})
	if err != nil {
		ctx.Logger.Error(err)
		return
	}
}
