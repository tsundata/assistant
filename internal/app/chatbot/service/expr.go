package service

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"log"
)

type ExprEnv struct {
	Ctx   context.Context
	Comp  component.Component
	Value interface{}
}

func (e ExprEnv) Message(sequence int64) pb.Message {
	reply, err := e.Comp.Message().GetBySequence(e.Ctx, &pb.MessageRequest{Message: &pb.Message{Sequence: sequence}})
	if err != nil {
		log.Println(err)
		return pb.Message{}
	}
	return *reply.Message
}
