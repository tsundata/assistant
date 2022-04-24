package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/system/repository"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
	"gorm.io/gorm"
)

type System struct {
	repo   repository.SystemRepository
	logger log.Logger
	locker *global.Locker
}

func NewSystem(repo repository.SystemRepository, logger log.Logger, locker *global.Locker) pb.SystemSvcServer {
	return &System{repo: repo, logger: logger, locker: locker}
}

func (s *System) CreateCounter(ctx context.Context, payload *pb.CounterRequest) (*pb.StateReply, error) {
	id, _ := md.FromIncoming(ctx)
	payload.Counter.UserId = id
	_, err := s.repo.CreateCounter(ctx, payload.Counter)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (s *System) GetCounter(ctx context.Context, payload *pb.CounterRequest) (*pb.CounterReply, error) {
	find, err := s.repo.GetCounter(ctx, payload.Counter.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.CounterReply{Counter: &find}, nil
}

func (s *System) GetCounters(ctx context.Context, _ *pb.CounterRequest) (*pb.CountersReply, error) {
	id, _ := md.FromIncoming(ctx)
	list, err := s.repo.ListCounter(ctx, id)
	if err != nil {
		return nil, err
	}
	return &pb.CountersReply{Counters: list}, nil
}

func (s *System) ChangeCounter(ctx context.Context, payload *pb.CounterRequest) (*pb.CounterReply, error) {
	id, _ := md.FromIncoming(ctx)
	find, err := s.repo.GetCounterByFlag(ctx, id, payload.Counter.GetFlag())
	if err != nil {
		return nil, err
	}
	err = s.repo.IncreaseCounter(ctx, find.Id, payload.Counter.GetDigit())
	if err != nil {
		return nil, err
	}
	find, err = s.repo.GetCounter(ctx, find.Id)
	if err != nil {
		return nil, err
	}
	return &pb.CounterReply{Counter: &find}, nil
}

func (s *System) ResetCounter(ctx context.Context, payload *pb.CounterRequest) (*pb.CounterReply, error) {
	id, _ := md.FromIncoming(ctx)
	l, err := s.locker.Acquire(fmt.Sprintf("chatbot:system:counter:reset:%d", id))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = l.Release()
	}()

	find, err := s.repo.GetCounterByFlag(ctx, id, payload.Counter.GetFlag())
	if err != nil {
		return nil, err
	}
	err = s.repo.IncreaseCounter(ctx, find.Id, 1-find.Digit)
	if err != nil {
		return nil, err
	}
	find, err = s.repo.GetCounter(ctx, find.Id)
	if err != nil {
		return nil, err
	}
	return &pb.CounterReply{Counter: &find}, nil
}

func (s *System) GetCounterByFlag(ctx context.Context, payload *pb.CounterRequest) (*pb.CounterReply, error) {
	id, _ := md.FromIncoming(ctx)
	find, err := s.repo.GetCounterByFlag(ctx, id, payload.Counter.GetFlag())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &pb.CounterReply{Counter: &find}, nil
}
