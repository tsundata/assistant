package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/user/repository"
)

type User struct {
	rdb  *redis.Client
	repo repository.UserRepository
}

func NewUser(rdb *redis.Client, repo repository.UserRepository) *User {
	return &User{rdb: rdb, repo: repo}
}

func (s *User) Authorization(ctx context.Context, payload *pb.TextRequest) (*pb.StateReply, error) {
	uuid, err := s.rdb.Get(ctx, "user:auth:token").Result()
	if err != nil {
		return &pb.StateReply{
			State: false,
		}, nil
	}

	return &pb.StateReply{
		State: payload.GetText() == uuid,
	}, nil
}

func (s *User) GetRole(_ context.Context, payload *pb.RoleRequest) (*pb.RoleReply, error) {
	find, err := s.repo.GetRole(int(payload.GetId()))
	if err != nil {
		return nil, err
	}
	return &pb.RoleReply{Role: &pb.Role{
		Profession:  find.Profession,
		Exp:         int64(find.Exp),
		Level:       int64(find.Level),
		Strength:    int64(find.Strength),
		Culture:     int64(find.Culture),
		Environment: int64(find.Environment),
		Charisma:    int64(find.Charisma),
		Talent:      int64(find.Talent),
		Intellect:   int64(find.Intellect),
	}}, nil
}
