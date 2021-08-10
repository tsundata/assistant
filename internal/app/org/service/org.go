package service

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/org/repository"
)

type Org struct {
	middle pb.MiddleSvcClient
	repo   repository.OrgRepository
}

func NewOrg(repo repository.OrgRepository, middle pb.MiddleSvcClient) *Org {
	return &Org{repo: repo, middle: middle}
}

func (o *Org) CreateObjective(ctx context.Context, payload *pb.ObjectiveRequest) (*pb.StateReply, error) {
	reply, err := o.middle.GetOrCreateTag(ctx, &pb.TagRequest{Tag: &pb.Tag{Name: payload.Objective.GetTag()}})
	if err != nil {
		return nil, err
	}
	item := pb.Objective{
		Name:  payload.Objective.GetName(),
		Tag:   payload.Objective.GetTag(),
		TagId: reply.Tag.GetId(),
	}

	_, err = o.repo.CreateObjective(item)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (o *Org) GetObjective(_ context.Context, payload *pb.ObjectiveRequest) (*pb.ObjectiveReply, error) {
	find, err := o.repo.GetObjectiveByID(payload.Objective.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.ObjectiveReply{
		Objective: &pb.Objective{
			Id:        find.Id,
			Name:      find.Name,
			Tag:       find.Tag,
			TagId:     find.TagId,
			CreatedAt: find.CreatedAt,
		},
	}, nil
}

func (o *Org) GetObjectives(_ context.Context, _ *pb.ObjectiveRequest) (*pb.ObjectivesReply, error) {
	items, err := o.repo.ListObjectives()
	if err != nil {
		return nil, err
	}

	var res []*pb.Objective
	for _, item := range items {
		res = append(res, &pb.Objective{
			Id:        item.Id,
			Name:      item.Name,
			Tag:       item.Tag,
			TagId:     item.TagId,
			CreatedAt: item.CreatedAt,
		})
	}

	return &pb.ObjectivesReply{Objective: res}, nil
}

func (o *Org) DeleteObjective(_ context.Context, payload *pb.ObjectiveRequest) (*pb.StateReply, error) {
	err := o.repo.DeleteObjective(payload.Objective.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (o *Org) CreateKeyResult(ctx context.Context, payload *pb.KeyResultRequest) (*pb.StateReply, error) {
	reply, err := o.middle.GetOrCreateTag(ctx, &pb.TagRequest{Tag: &pb.Tag{Name: payload.KeyResult.GetTag()}})
	if err != nil {
		return nil, err
	}
	item := pb.KeyResult{
		ObjectiveId: payload.KeyResult.GetObjectiveId(),
		Name:        payload.KeyResult.GetName(),
		Tag:         payload.KeyResult.GetTag(),
		TagId:       reply.Tag.GetId(),
	}

	_, err = o.repo.CreateKeyResult(item)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (o *Org) GetKeyResult(_ context.Context, payload *pb.KeyResultRequest) (*pb.KeyResultReply, error) {
	find, err := o.repo.GetKeyResultByID(payload.KeyResult.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.KeyResultReply{
		KeyResult: &pb.KeyResult{
			Id:          find.Id,
			Name:        find.Name,
			ObjectiveId: find.ObjectiveId,
			Tag:         find.Tag,
			TagId:       find.TagId,
			Complete:    find.Complete,
			CreatedAt:   find.CreatedAt,
			UpdatedAt:   find.UpdatedAt,
		},
	}, nil
}

func (o *Org) GetKeyResults(_ context.Context, _ *pb.KeyResultRequest) (*pb.KeyResultsReply, error) {
	items, err := o.repo.ListKeyResults()
	if err != nil {
		return nil, err
	}

	var res []*pb.KeyResult
	for _, item := range items {
		res = append(res, &pb.KeyResult{
			Id:          item.Id,
			Name:        item.Name,
			Tag:         item.Tag,
			TagId:       item.TagId,
			ObjectiveId: item.ObjectiveId,
			Complete:    item.Complete,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}

	return &pb.KeyResultsReply{Result: res}, nil
}

func (o *Org) DeleteKeyResult(_ context.Context, payload *pb.KeyResultRequest) (*pb.StateReply, error) {
	err := o.repo.DeleteKeyResult(payload.KeyResult.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}
