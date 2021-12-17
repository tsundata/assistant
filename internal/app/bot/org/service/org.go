package service

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/bot/org/repository"
)

type Org struct {
	middle pb.MiddleSvcClient
	repo   repository.OrgRepository
}

func NewOrg(repo repository.OrgRepository, middle pb.MiddleSvcClient) *Org {
	return &Org{repo: repo, middle: middle}
}

func (o *Org) CreateObjective(ctx context.Context, payload *pb.ObjectiveRequest) (*pb.StateReply, error) {
	reply, err := o.middle.GetOrCreateTag(ctx, &pb.TagRequest{Tag: &pb.Tag{Name: payload.Tag}})
	if err != nil {
		return nil, err
	}
	item := pb.Objective{
		Name:  payload.Objective.GetName(),
		TagId: reply.Tag.GetId(),
	}

	_, err = o.repo.CreateObjective(ctx, &item)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (o *Org) GetObjective(ctx context.Context, payload *pb.ObjectiveRequest) (*pb.ObjectiveReply, error) {
	find, err := o.repo.GetObjectiveByID(ctx, payload.Objective.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.ObjectiveReply{
		Objective: &pb.Objective{
			Id:        find.Id,
			Name:      find.Name,
			TagId:     find.TagId,
			CreatedAt: find.CreatedAt,
		},
	}, nil
}

func (o *Org) GetObjectives(ctx context.Context, _ *pb.ObjectiveRequest) (*pb.ObjectivesReply, error) {
	items, err := o.repo.ListObjectives(ctx)
	if err != nil {
		return nil, err
	}

	var res []*pb.Objective
	for _, item := range items {
		res = append(res, &pb.Objective{
			Id:        item.Id,
			Name:      item.Name,
			TagId:     item.TagId,
			CreatedAt: item.CreatedAt,
		})
	}

	return &pb.ObjectivesReply{Objective: res}, nil
}

func (o *Org) DeleteObjective(ctx context.Context, payload *pb.ObjectiveRequest) (*pb.StateReply, error) {
	err := o.repo.DeleteObjective(ctx, payload.Objective.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (o *Org) CreateKeyResult(ctx context.Context, payload *pb.KeyResultRequest) (*pb.StateReply, error) {
	reply, err := o.middle.GetOrCreateTag(ctx, &pb.TagRequest{Tag: &pb.Tag{Name: payload.Tag}})
	if err != nil {
		return nil, err
	}
	item := pb.KeyResult{
		ObjectiveId: payload.KeyResult.GetObjectiveId(),
		Name:        payload.KeyResult.GetName(),
		TagId:       reply.Tag.GetId(),
	}

	_, err = o.repo.CreateKeyResult(ctx, &item)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (o *Org) GetKeyResult(ctx context.Context, payload *pb.KeyResultRequest) (*pb.KeyResultReply, error) {
	find, err := o.repo.GetKeyResultByID(ctx, payload.KeyResult.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.KeyResultReply{
		KeyResult: &pb.KeyResult{
			Id:          find.Id,
			Name:        find.Name,
			ObjectiveId: find.ObjectiveId,
			TagId:       find.TagId,
			Complete:    find.Complete,
			CreatedAt:   find.CreatedAt,
			UpdatedAt:   find.UpdatedAt,
		},
	}, nil
}

func (o *Org) GetKeyResults(ctx context.Context, _ *pb.KeyResultRequest) (*pb.KeyResultsReply, error) {
	items, err := o.repo.ListKeyResults(ctx)
	if err != nil {
		return nil, err
	}

	var res []*pb.KeyResult
	for _, item := range items {
		res = append(res, &pb.KeyResult{
			Id:          item.Id,
			Name:        item.Name,
			TagId:       item.TagId,
			ObjectiveId: item.ObjectiveId,
			Complete:    item.Complete,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}

	return &pb.KeyResultsReply{Result: res}, nil
}

func (o *Org) DeleteKeyResult(ctx context.Context, payload *pb.KeyResultRequest) (*pb.StateReply, error) {
	err := o.repo.DeleteKeyResult(ctx, payload.KeyResult.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}
