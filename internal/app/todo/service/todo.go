package service

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/todo/repository"
	"github.com/tsundata/assistant/internal/pkg/model"
	"time"
)

type Todo struct {
	repo repository.TodoRepository
}

func NewTodo(repo repository.TodoRepository) *Todo {
	return &Todo{repo: repo}
}

func (s *Todo) CreateTodo(_ context.Context, payload *pb.TodoRequest) (*pb.StateReply, error) {
	todo := model.Todo{
		Content:        payload.GetContent(),
		Priority:       0,
		IsRemindAtTime: false,
		RemindAt:       nil,
		RepeatMethod:   "",
		RepeatRule:     "",
		Category:       "",
		Remark:         "",
		Complete:       false,
		Time:           time.Now(),
	}
	_, err := s.repo.CreateTodo(todo)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *Todo) GetTodo(_ context.Context, payload *pb.TodoRequest) (*pb.TodoReply, error) {
	find, err := s.repo.GetTodo(int(payload.GetId()))
	if err != nil {
		return nil, err
	}

	return &pb.TodoReply{
		Todo: &pb.TodoItem{
			Content:  find.Content,
			Priority: int64(find.Priority),
			Remark:   find.Remark,
			Complete: find.Complete,
		},
	}, nil
}

func (s *Todo) GetTodos(_ context.Context, _ *pb.TodoRequest) (*pb.TodosReply, error) {
	items, err := s.repo.ListTodos()
	if err != nil {
		return nil, err
	}

	var res []*pb.TodoItem
	for _, item := range items {
		res = append(res, &pb.TodoItem{
			Content:  item.Content,
			Priority: int64(item.Priority),
			Remark:   item.Remark,
			Complete: item.Complete,
		})
	}

	return &pb.TodosReply{Todos: res}, nil
}

func (s *Todo) DeleteTodo(_ context.Context, payload *pb.TodoRequest) (*pb.StateReply, error) {
	err := s.repo.DeleteTodo(int(payload.GetId()))
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Todo) UpdateTodo(_ context.Context, payload *pb.TodoRequest) (*pb.StateReply, error) {
	err := s.repo.UpdateTodo(model.Todo{
		ID:      int(payload.GetId()),
		Content: payload.GetContent(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Todo) CompleteTodo(_ context.Context, payload *pb.TodoRequest) (*pb.StateReply, error) {
	err := s.repo.CompleteTodo(int(payload.GetId()))
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}
