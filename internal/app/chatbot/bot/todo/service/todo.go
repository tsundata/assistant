package service

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/todo/repository"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
)

type Todo struct {
	repo   repository.TodoRepository
	bus    event.Bus
	logger log.Logger
}

func NewTodo(bus event.Bus, logger log.Logger, repo repository.TodoRepository) pb.TodoSvcServer {
	return &Todo{bus: bus, repo: repo, logger: logger}
}

func (s *Todo) CreateTodo(ctx context.Context, payload *pb.TodoRequest) (*pb.StateReply, error) {
	id, _ := md.FromIncoming(ctx)

	var err error
	todo := pb.Todo{
		Content:        payload.Todo.GetContent(),
		Priority:       payload.Todo.GetPriority(),
		IsRemindAtTime: payload.Todo.GetIsRemindAtTime(),
		RemindAt:       0, //&remindAt, fixme
		RepeatMethod:   payload.Todo.GetRepeatMethod(),
		RepeatRule:     payload.Todo.GetRepeatRule(),
		RepeatEndAt:    0, //&endAt, fixme
		Category:       "",
		Remark:         payload.Todo.GetRemark(),
		Complete:       false,
	}
	_, err = s.repo.CreateTodo(ctx, &todo)
	if err != nil {
		return nil, err
	}

	if s.bus != nil {
		err = s.bus.Publish(ctx, enum.User, event.RoleChangeExpSubject, pb.Role{UserId: id, Exp: enum.TodoCreatedExp})
		if err != nil {
			s.logger.Error(err)
			return nil, err
		}
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Todo) GetTodo(ctx context.Context, payload *pb.TodoRequest) (*pb.TodoReply, error) {
	find, err := s.repo.GetTodo(ctx, payload.Todo.GetId())
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

func (s *Todo) GetTodos(ctx context.Context, _ *pb.TodoRequest) (*pb.TodosReply, error) {
	items, err := s.repo.ListTodos(ctx)
	if err != nil {
		return nil, err
	}

	var res []*pb.Todo
	for _, item := range items {
		res = append(res, &pb.Todo{
			Id:       item.Id,
			Content:  item.Content,
			Priority: item.Priority,
			Remark:   item.Remark,
			Complete: item.Complete,
		})
	}

	return &pb.TodosReply{Todos: res}, nil
}

func (s *Todo) DeleteTodo(ctx context.Context, payload *pb.TodoRequest) (*pb.StateReply, error) {
	err := s.repo.DeleteTodo(ctx, payload.Todo.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Todo) UpdateTodo(ctx context.Context, payload *pb.TodoRequest) (*pb.StateReply, error) {
	err := s.repo.UpdateTodo(ctx, &pb.Todo{
		Id:      payload.Todo.GetId(),
		Content: payload.Todo.GetContent(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Todo) CompleteTodo(ctx context.Context, payload *pb.TodoRequest) (*pb.StateReply, error) {
	id, _ := md.FromIncoming(ctx)

	err := s.repo.CompleteTodo(ctx, payload.Todo.GetId())
	if err != nil {
		return nil, err
	}

	if s.bus != nil {
		err = s.bus.Publish(ctx, enum.User, event.RoleChangeExpSubject, pb.Role{UserId: id, Exp: enum.TodoCompletedExp})
		if err != nil {
			s.logger.Error(err)
			return nil, err
		}

		find, err := s.repo.GetTodo(ctx, payload.Todo.GetId())
		if err != nil {
			return nil, err
		}
		err = s.bus.Publish(ctx, enum.User, event.RoleChangeAttrSubject, pb.AttrChange{UserId: id, Content: find.Content})
		if err != nil {
			s.logger.Error(err)
			return nil, err
		}
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Todo) GetRemindTodos(ctx context.Context, _ *pb.TodoRequest) (*pb.TodosReply, error) {
	items, err := s.repo.ListRemindTodos(ctx)
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
