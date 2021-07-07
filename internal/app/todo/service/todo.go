package service

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/todo/repository"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/model"
	"time"
)

type Todo struct {
	repo   repository.TodoRepository
	bus    *event.Bus
	logger *logger.Logger
}

func NewTodo(bus *event.Bus, logger *logger.Logger, repo repository.TodoRepository) *Todo {
	return &Todo{bus: bus, repo: repo, logger: logger}
}

func (s *Todo) CreateTodo(_ context.Context, payload *pb.TodoRequest) (*pb.StateReply, error) {
	var err error
	var remindAt time.Time
	if payload.GetRemindAt() != "" {
		remindAt, err = time.ParseInLocation("2006-01-02 15:04", payload.GetRemindAt(), time.Local)
		if err != nil {
			return nil, err
		}
	}

	var endAt time.Time
	if payload.GetRepeatEndAt() != "" {
		endAt, err = time.ParseInLocation("2006-01-02 15:04", payload.GetRepeatEndAt(), time.Local)
		if err != nil {
			return nil, err
		}
	}

	todo := model.Todo{
		Content:        payload.GetContent(),
		Priority:       payload.GetPriority(),
		IsRemindAtTime: payload.GetIsRemindAtTime(),
		RemindAt:       &remindAt,
		RepeatMethod:   payload.GetRepeatMethod(),
		RepeatRule:     payload.GetRepeatRule(),
		RepeatEndAt:    &endAt,
		Category:       "",
		Remark:         payload.GetRemark(),
		Complete:       false,
	}
	_, err = s.repo.CreateTodo(todo)
	if err != nil {
		return nil, err
	}

	if s.bus != nil {
		err = s.bus.Publish(event.ChangeExpSubject, model.Role{UserID: model.SuperUserID, Exp: model.TodoCreatedExp})
		if err != nil {
			s.logger.Error(err)
			return nil, err
		}
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Todo) GetTodo(_ context.Context, payload *pb.TodoRequest) (*pb.TodoReply, error) {
	find, err := s.repo.GetTodo(payload.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.TodoReply{
		Todo: &pb.TodoItem{
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

	var res []*pb.TodoItem
	for _, item := range items {
		res = append(res, &pb.TodoItem{
			Content:  item.Content,
			Priority: item.Priority,
			Remark:   item.Remark,
			Complete: item.Complete,
		})
	}

	return &pb.TodosReply{Todos: res}, nil
}

func (s *Todo) DeleteTodo(_ context.Context, payload *pb.TodoRequest) (*pb.StateReply, error) {
	err := s.repo.DeleteTodo(payload.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Todo) UpdateTodo(_ context.Context, payload *pb.TodoRequest) (*pb.StateReply, error) {
	err := s.repo.UpdateTodo(model.Todo{
		ID:      payload.GetId(),
		Content: payload.GetContent(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Todo) CompleteTodo(_ context.Context, payload *pb.TodoRequest) (*pb.StateReply, error) {
	err := s.repo.CompleteTodo(payload.GetId())
	if err != nil {
		return nil, err
	}

	if s.bus != nil {
		err = s.bus.Publish(event.ChangeExpSubject, model.Role{UserID: model.SuperUserID, Exp: model.TodoCompletedExp})
		if err != nil {
			s.logger.Error(err)
			return nil, err
		}

		find, err := s.repo.GetTodo(payload.GetId())
		if err != nil {
			return nil, err
		}
		err = s.bus.Publish(event.ChangeAttrSubject, model.AttrChange{UserID: model.SuperUserID, Content: find.Content})
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

	layout := "2006-01-02 15:04"
	var res []*pb.TodoItem
	for _, item := range items {
		remindAt := ""
		if item.RemindAt != nil {
			remindAt = item.RemindAt.Format(layout)
		}
		endAt := ""
		if item.RepeatEndAt != nil {
			endAt = item.RepeatEndAt.Format(layout)
		}

		res = append(res, &pb.TodoItem{
			Id:             item.ID,
			Content:        item.Content,
			Priority:       item.Priority,
			IsRemindAtTime: item.IsRemindAtTime,
			RemindAt:       remindAt,
			RepeatMethod:   item.RepeatMethod,
			RepeatRule:     item.RepeatRule,
			RepeatEndAt:    endAt,
			CreatedAt:      item.CreatedAt.Format(layout),
		})
	}

	return &pb.TodosReply{Todos: res}, nil
}
