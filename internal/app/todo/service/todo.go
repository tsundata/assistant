package service

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/model"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/todo/repository"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"time"
)

type Todo struct {
	repo   repository.TodoRepository
	bus    event.Bus
	logger log.Logger
}

func NewTodo(bus event.Bus, logger log.Logger, repo repository.TodoRepository) *Todo {
	return &Todo{bus: bus, repo: repo, logger: logger}
}

func (s *Todo) CreateTodo(ctx context.Context, payload *pb.TodoRequest) (*pb.StateReply, error) {
	var err error
	var remindAt time.Time
	if payload.Todo.GetRemindAt() != "" {
		remindAt, err = time.ParseInLocation("2006-01-02 15:04", payload.Todo.GetRemindAt(), time.Local)
		if err != nil {
			return nil, err
		}
	}

	var endAt time.Time
	if payload.Todo.GetRepeatEndAt() != "" {
		endAt, err = time.ParseInLocation("2006-01-02 15:04", payload.Todo.GetRepeatEndAt(), time.Local)
		if err != nil {
			return nil, err
		}
	}

	fmt.Println(remindAt) // fixme
	fmt.Println(endAt)    // fixme

	todo := pb.Todo{
		Content:        payload.Todo.GetContent(),
		Priority:       payload.Todo.GetPriority(),
		IsRemindAtTime: payload.Todo.GetIsRemindAtTime(),
		RemindAt:       "", //&remindAt, fixme
		RepeatMethod:   payload.Todo.GetRepeatMethod(),
		RepeatRule:     payload.Todo.GetRepeatRule(),
		RepeatEndAt:    "", //&endAt, fixme
		Category:       "",
		Remark:         payload.Todo.GetRemark(),
		Complete:       false,
	}
	_, err = s.repo.CreateTodo(todo)
	if err != nil {
		return nil, err
	}

	if s.bus != nil {
		err = s.bus.Publish(ctx, event.ChangeExpSubject, pb.Role{UserId: model.SuperUserID, Exp: enum.TodoCreatedExp})
		if err != nil {
			s.logger.Error(err)
			return nil, err
		}
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Todo) GetTodo(_ context.Context, payload *pb.TodoRequest) (*pb.TodoReply, error) {
	find, err := s.repo.GetTodo(payload.Todo.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.TodoReply{
		Todo: &pb.Todo{
			Content:  find.Content,
			Priority: find.Priority,
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

	var res []*pb.Todo
	for _, item := range items {
		res = append(res, &pb.Todo{
			Content:  item.Content,
			Priority: item.Priority,
			Remark:   item.Remark,
			Complete: item.Complete,
		})
	}

	return &pb.TodosReply{Todos: res}, nil
}

func (s *Todo) DeleteTodo(_ context.Context, payload *pb.TodoRequest) (*pb.StateReply, error) {
	err := s.repo.DeleteTodo(payload.Todo.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Todo) UpdateTodo(_ context.Context, payload *pb.TodoRequest) (*pb.StateReply, error) {
	err := s.repo.UpdateTodo(pb.Todo{
		Id:      payload.Todo.GetId(),
		Content: payload.Todo.GetContent(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Todo) CompleteTodo(ctx context.Context, payload *pb.TodoRequest) (*pb.StateReply, error) {
	err := s.repo.CompleteTodo(payload.Todo.GetId())
	if err != nil {
		return nil, err
	}

	if s.bus != nil {
		err = s.bus.Publish(ctx, event.ChangeExpSubject, pb.Role{UserId: model.SuperUserID, Exp: enum.TodoCompletedExp})
		if err != nil {
			s.logger.Error(err)
			return nil, err
		}

		find, err := s.repo.GetTodo(payload.Todo.GetId())
		if err != nil {
			return nil, err
		}
		err = s.bus.Publish(ctx, event.ChangeAttrSubject, pb.AttrChange{UserId: model.SuperUserID, Content: find.Content})
		if err != nil {
			s.logger.Error(err)
			return nil, err
		}
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Todo) GetRemindTodos(_ context.Context, _ *pb.TodoRequest) (*pb.TodosReply, error) {
	items, err := s.repo.ListRemindTodos()
	if err != nil {
		return nil, err
	}

	var res []*pb.Todo
	for _, item := range items {
		res = append(res, &pb.Todo{
			Id:             item.Id,
			Content:        item.Content,
			Priority:       item.Priority,
			IsRemindAtTime: item.IsRemindAtTime,
			RemindAt:       item.RemindAt,
			RepeatMethod:   item.RepeatMethod,
			RepeatRule:     item.RepeatRule,
			RepeatEndAt:    item.RepeatEndAt,
			CreatedAt:      item.CreatedAt,
		})
	}

	return &pb.TodosReply{Todos: res}, nil
}
