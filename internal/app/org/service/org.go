package service

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
)

type Org struct{}

func NewOrg() *Org {
	return &Org{}
}

func (o Org) CreateObjective(ctx context.Context, payload *pb.ObjectiveRequest) (*pb.StateReply, error) {
	panic("implement me")
}

func (o Org) GetObjective(ctx context.Context, payload *pb.ObjectiveRequest) (*pb.ObjectiveReply, error) {
	panic("implement me")
}

func (o Org) GetObjectives(ctx context.Context, payload *pb.ObjectiveRequest) (*pb.ObjectivesReply, error) {
	panic("implement me")
}

func (o Org) DeleteObjective(ctx context.Context, payload *pb.ObjectiveRequest) (*pb.StateReply, error) {
	panic("implement me")
}

func (o Org) CreateKeyResult(ctx context.Context, payload *pb.KeyResultRequest) (*pb.StateReply, error) {
	panic("implement me")
}

func (o Org) GetKeyResult(ctx context.Context, payload *pb.KeyResultRequest) (*pb.KeyResultReply, error) {
	panic("implement me")
}

func (o Org) GetKeyResults(ctx context.Context, payload *pb.KeyResultRequest) (*pb.KeyResultsReply, error) {
	panic("implement me")
}

func (o Org) DeleteKeyResult(ctx context.Context, payload *pb.KeyResultRequest) (*pb.StateReply, error) {
	panic("implement me")
}
