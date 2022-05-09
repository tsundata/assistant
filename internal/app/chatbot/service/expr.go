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

func (e ExprEnv) Todo(sequence int64) pb.Todo {
	reply, err := e.Comp.Todo().GetTodo(e.Ctx, &pb.TodoRequest{Todo: &pb.Todo{Sequence: sequence}})
	if err != nil {
		log.Println(err)
		return pb.Todo{}
	}
	return *reply.Todo
}

func (e ExprEnv) KeyResult(sequence int64) pb.KeyResult {
	reply, err := e.Comp.Okr().GetKeyResult(e.Ctx, &pb.KeyResultRequest{KeyResult: &pb.KeyResult{Sequence: sequence}})
	if err != nil {
		return pb.KeyResult{}
	}
	return *reply.KeyResult
}

func (e ExprEnv) Counter(flag string) pb.Counter {
	reply, err := e.Comp.Middle().GetCounterByFlag(e.Ctx, &pb.CounterRequest{Counter: &pb.Counter{Flag: flag}})
	if err != nil {
		log.Println(err)
		return pb.Counter{}
	}
	return *reply.Counter
}
