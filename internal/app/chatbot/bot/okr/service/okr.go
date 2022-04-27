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

func (o *Okr) UpdateObjective(ctx context.Context, payload *pb.ObjectiveRequest) (*pb.StateReply, error) {
	id, _ := md.FromIncoming(ctx)
	payload.Objective.UserId = id
	err := o.repo.UpdateObjective(ctx, payload.Objective)
	if err != nil {
		return nil, err
	}
	return &pb.StateReply{State: true}, nil
}

func (o *Okr) GetObjective(ctx context.Context, payload *pb.ObjectiveRequest) (*pb.ObjectiveReply, error) {
	id, _ := md.FromIncoming(ctx)
	find, err := o.repo.GetObjectiveBySequence(ctx, id, payload.Objective.GetSequence())
	if err != nil {
		return nil, err
	}

	return &pb.ObjectiveReply{Objective: find}, nil
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
	id, _ := md.FromIncoming(ctx)
	err := o.repo.DeleteObjectiveBySequence(ctx, id, payload.Objective.GetSequence())
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

func (o *Okr) UpdateKeyResult(ctx context.Context, payload *pb.KeyResultRequest) (*pb.StateReply, error) {
	id, _ := md.FromIncoming(ctx)

	if payload.KeyResult.ValueMode != enum.ValueSumMode &&
		payload.KeyResult.ValueMode != enum.ValueLastMode &&
		payload.KeyResult.ValueMode != enum.ValueAvgMode &&
		payload.KeyResult.ValueMode != enum.ValueMaxMode {
		return nil, errors.New("error key result value mode")
	}

	payload.KeyResult.UserId = id
	err := o.repo.UpdateKeyResult(ctx, payload.KeyResult)
	if err != nil {
		return nil, err
	}

	// update value
	reply, err := o.repo.GetKeyResultBySequence(ctx, id, payload.KeyResult.Sequence)
	if err != nil {
		return nil, err
	}
	err = o.repo.AggregateKeyResultValue(ctx, reply.Id)
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (o *Okr) GetKeyResult(ctx context.Context, payload *pb.KeyResultRequest) (*pb.KeyResultReply, error) {
	id, _ := md.FromIncoming(ctx)
	find, err := o.repo.GetKeyResultBySequence(ctx, id, payload.KeyResult.GetSequence())
	if err != nil {
		return nil, err
	}

	return &pb.KeyResultReply{KeyResult: find}, nil
}

func (o *Okr) GetKeyResults(ctx context.Context, _ *pb.KeyResultRequest) (*pb.KeyResultsReply, error) {
	id, _ := md.FromIncoming(ctx)
	items, err := o.repo.ListKeyResults(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.KeyResultsReply{Result: items}, nil
}

func (o *Okr) GetKeyResultsByTag(ctx context.Context, payload *pb.KeyResultRequest) (*pb.KeyResultsReply, error) {
	reply, err := o.middle.GetModelTags(ctx, &pb.ModelTagRequest{Model: &pb.ModelTag{
		Service: enum.Chatbot,
		Model:   util.ModelName(pb.KeyResult{}),
		Name:    payload.Tag,
	}})
	if err != nil {
		return nil, err
	}

	var krId []int64
	for _, item := range reply.Tags {
		krId = append(krId, item.Id)
	}

	items, err := o.repo.ListKeyResultsById(ctx, krId)
	if err != nil {
		return nil, err
	}

	return &pb.KeyResultsReply{Result: items}, nil
}

func (o *Okr) DeleteKeyResult(ctx context.Context, payload *pb.KeyResultRequest) (*pb.StateReply, error) {
	id, _ := md.FromIncoming(ctx)
	err := o.repo.DeleteKeyResultBySequence(ctx, id, payload.KeyResult.GetSequence())
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
