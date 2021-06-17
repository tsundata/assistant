package service

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/pb"
)

const RuleKey = "subscribe:rule"

type Subscribe struct {
	rdb *redis.Client
}

func NewSubscribe(rdb *redis.Client) *Subscribe {
	return &Subscribe{rdb}
}

func (s *Subscribe) List(ctx context.Context, _ *pb.SubscribeRequest) (*pb.SubscribeReply, error) {
	res, err := s.rdb.HGetAll(ctx, RuleKey).Result()
	if err != nil {
		return nil, err
	}

	var result []string
	for source, isSubscribe := range res {
		result = append(result, fmt.Sprintf("%s [Subscribe:%v]", source, isSubscribe))
	}

	return &pb.SubscribeReply{
		Text: result,
	}, nil
}

func (s *Subscribe) Register(ctx context.Context, payload *pb.SubscribeRequest) (*pb.StateReply, error) {
	resp, err := s.rdb.HMGet(ctx, RuleKey, payload.GetText()).Result()
	if err != nil {
		return nil, err
	}

	if len(resp) == 0 {
		_, err = s.rdb.HMSet(ctx, RuleKey, payload.GetText(), "true").Result()
		if err != nil {
			return nil, err
		}
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Subscribe) Open(ctx context.Context, payload *pb.SubscribeRequest) (*pb.StateReply, error) {
	_, err := s.rdb.HMSet(ctx, RuleKey, payload.GetText(), "true").Result()
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Subscribe) Close(ctx context.Context, payload *pb.SubscribeRequest) (*pb.StateReply, error) {
	_, err := s.rdb.HMSet(ctx, RuleKey, payload.GetText(), "false").Result()
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (s *Subscribe) Status(ctx context.Context, payload *pb.SubscribeRequest) (*pb.StateReply, error) {
	resp, err := s.rdb.HGetAll(ctx, RuleKey).Result()
	if err != nil {
		return nil, err
	}
	for k, v := range resp {
		if k == payload.GetText() {
			return &pb.StateReply{
				State: v == "true",
			}, nil
		}
	}
	return &pb.StateReply{
		State: false,
	}, nil
}
