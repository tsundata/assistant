package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/pb"
)

type User struct {
	rdb *redis.Client
}

func NewUser(rdb *redis.Client) *User {
	return &User{rdb: rdb}
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

func (s *User) CreateRole(ctx context.Context, request *pb.RoleRequest) (*pb.StateReply, error) {
	panic("implement me")
}

func (s *User) GetRole(ctx context.Context, request *pb.RoleRequest) (*pb.RoleReply, error) {
	panic("implement me")
}

func (s *User) GetRoles(ctx context.Context, request *pb.RoleRequest) (*pb.RolesReply, error) {
	panic("implement me")
}

func (s *User) UpdateRoles(ctx context.Context, request *pb.RoleRequest) (*pb.StateReply, error) {
	panic("implement me")
}
