package service

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
)

type Todo struct{}

func NewTodo() *Todo {
	return &Todo{}
}

func (t Todo) CreateTodo(ctx context.Context, req *pb.TodoRequest) (*pb.StateReply, error) {
	panic("implement me")
}

func (t Todo) GetTodo(ctx context.Context, req *pb.TodoRequest) (*pb.TodoReply, error) {
	panic("implement me")
}

func (t Todo) GetTodos(ctx context.Context, req *pb.TodoRequest) (*pb.TodosReply, error) {
	panic("implement me")
}

func (t Todo) DeleteTodo(ctx context.Context, req *pb.TodoRequest) (*pb.StateReply, error) {
	panic("implement me")
}

func (t Todo) UpdateTodo(ctx context.Context, req *pb.TodoRequest) (*pb.StateReply, error) {
	panic("implement me")
}

func (t Todo) CompleteTodo(ctx context.Context, req *pb.TodoRequest) (*pb.StateReply, error) {
	panic("implement me")
}
