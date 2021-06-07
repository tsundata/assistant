package service

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (u User) CreateRole(ctx context.Context, request *pb.RoleRequest) (*pb.StateReply, error) {
	panic("implement me")
}

func (u User) GetRole(ctx context.Context, request *pb.RoleRequest) (*pb.RoleReply, error) {
	panic("implement me")
}

func (u User) GetRoles(ctx context.Context, request *pb.RoleRequest) (*pb.RolesReply, error) {
	panic("implement me")
}

func (u User) UpdateRoles(ctx context.Context, request *pb.RoleRequest) (*pb.StateReply, error) {
	panic("implement me")
}
