package service

import (
	"context"
	"errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/okr/repository"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
	"github.com/tsundata/assistant/internal/pkg/util"
)

type Okr struct {
	middle pb.MiddleSvcClient
	repo   repository.OkrRepository
}

func NewOkr(repo repository.OkrRepository, middle pb.MiddleSvcClient) pb.OkrSvcServer {
	return &Okr{repo: repo, middle: middle}
}

func (o *Okr) CreateObjective(ctx context.Context, payload *pb.ObjectiveRequest) (*pb.StateReply, error) {
	id, _ := md.FromIncoming(ctx)
	payload.Objective.UserId = id

	objectiveId, err := o.repo.CreateObjective(ctx, payload.Objective)
	if err != nil {
		return nil, err
	}

	if payload.Objective.GetTag() != "" {
		_, err = o.middle.SaveModelTag(ctx, &pb.ModelTagRequest{
			Model: &pb.ModelTag{
				Service: enum.Chatbot,
				Model:   util.ModelName(pb.Objective{}),
				ModelId: objectiveId,
			},
			Tag: payload.Tag,
		})
		if err != nil {
			return nil, err
		}
	}

	return &pb.StateReply{State: true}, nil
}

func (o *Okr) GetObjective(ctx context.Context, payload *pb.ObjectiveRequest) (*pb.ObjectiveReply, error) {
	find, err := o.repo.GetObjectiveByID(ctx, payload.Objective.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.ObjectiveReply{
		Objective: &pb.Objective{
			Id:        find.Id,
			Title:     find.Title,
			CreatedAt: find.CreatedAt,
		},
	}, nil
}

func (o *Okr) GetObjectives(ctx context.Context, _ *pb.ObjectiveRequest) (*pb.ObjectivesReply, error) {
	id, _ := md.FromIncoming(ctx)
	items, err := o.repo.ListObjectives(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.ObjectivesReply{Objective: items}, nil
}

func (o *Okr) DeleteObjective(ctx context.Context, payload *pb.ObjectiveRequest) (*pb.StateReply, error) {
	err := o.repo.DeleteObjective(ctx, payload.Objective.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (o *Okr) CreateKeyResult(ctx context.Context, payload *pb.KeyResultRequest) (*pb.StateReply, error) {
	id, _ := md.FromIncoming(ctx)
	payload.KeyResult.UserId = id

	objective, err := o.repo.GetObjectiveBySequence(ctx, id, payload.ObjectiveSequence)
	if err != nil {
		return nil, err
	}

	// check
	if payload.KeyResult.TargetValue <= 0 {
		return nil, errors.New("error key result target value")
	}
	if payload.KeyResult.ValueMode != enum.ValueSumMode &&
		payload.KeyResult.ValueMode != enum.ValueLastMode &&
		payload.KeyResult.ValueMode != enum.ValueAvgMode &&
		payload.KeyResult.ValueMode != enum.ValueMaxMode {
		return nil, errors.New("error key result value mode")
	}

	// store
	if payload.KeyResult.InitialValue > 0 {
		payload.KeyResult.CurrentValue = payload.KeyResult.InitialValue
	}
	payload.KeyResult.ObjectiveId = objective.Id
	keyResultId, err := o.repo.CreateKeyResult(ctx, payload.KeyResult)
	if err != nil {
		return nil, err
	}

	// aggregate
	err = o.repo.AggregateObjectiveValue(ctx, objective.Id)
	if err != nil {
		return nil, err
	}

	if payload.KeyResult.GetTag() != "" {
		_, err = o.middle.SaveModelTag(ctx, &pb.ModelTagRequest{
			Model: &pb.ModelTag{
				Service: enum.Chatbot,
				Model:   util.ModelName(pb.KeyResult{}),
				ModelId: keyResultId,
			},
			Tag: payload.Tag,
		})
		if err != nil {
			return nil, err
		}
	}

	return &pb.StateReply{State: true}, nil
}

func (o *Okr) GetKeyResult(ctx context.Context, payload *pb.KeyResultRequest) (*pb.KeyResultReply, error) {
	find, err := o.repo.GetKeyResultByID(ctx, payload.KeyResult.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.KeyResultReply{
		KeyResult: &pb.KeyResult{
			Id:          find.Id,
			Title:       find.Title,
			ObjectiveId: find.ObjectiveId,
			CreatedAt:   find.CreatedAt,
			UpdatedAt:   find.UpdatedAt,
		},
	}, nil
}

func (o *Okr) GetKeyResults(ctx context.Context, _ *pb.KeyResultRequest) (*pb.KeyResultsReply, error) {
	id, _ := md.FromIncoming(ctx)
	items, err := o.repo.ListKeyResults(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.KeyResultsReply{Result: items}, nil
}

func (o *Okr) DeleteKeyResult(ctx context.Context, payload *pb.KeyResultRequest) (*pb.StateReply, error) {
	err := o.repo.DeleteKeyResult(ctx, payload.KeyResult.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (o *Okr) CreateKeyResultValue(ctx context.Context, payload *pb.KeyResultValueRequest) (*pb.StateReply, error) {
	id, _ := md.FromIncoming(ctx)
	keyResult, err := o.repo.GetKeyResultBySequence(ctx, id, payload.KeyResultSequence)
	if err != nil {
		return nil, err
	}
	_, err = o.repo.CreateKeyResultValue(ctx, &pb.KeyResultValue{Value: int32(payload.Value), KeyResultId: keyResult.Id})
	if err != nil {
		return nil, err
	}
	err = o.repo.AggregateKeyResultValue(ctx, keyResult.Id)
	if err != nil {
		return nil, err
	}
	err = o.repo.AggregateObjectiveValue(ctx, keyResult.ObjectiveId)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}
