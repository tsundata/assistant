package service

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
)

type Org struct{}

func NewOrg() *Org {
	return &Org{}
}

func (o Org) CreateObjective(ctx context.Context, request *pb.ObjectiveRequest) (*pb.StateReply, error) {
	panic("implement me")
}

func (o Org) GetObjective(ctx context.Context, request *pb.ObjectiveRequest) (*pb.ObjectiveReply, error) {
	panic("implement me")
}

func (o Org) GetObjectives(ctx context.Context, request *pb.ObjectiveRequest) (*pb.ObjectiveReply, error) {
	panic("implement me")
}

func (o Org) DeleteObjective(ctx context.Context, request *pb.ObjectiveRequest) (*pb.StateReply, error) {
	panic("implement me")
}
