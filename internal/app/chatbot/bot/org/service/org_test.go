package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/tsundata/assistant/mock"
	"reflect"
	"testing"
	"time"

	"github.com/tsundata/assistant/api/pb"
)

func TestOrg_CreateObjective(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockOrgRepository(ctl)
	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		repo.EXPECT().CreateObjective(gomock.Any(), gomock.Any()).Return(int64(1), nil),
		middle.EXPECT().SaveModelTag(gomock.Any(), gomock.Any()).Return(&pb.ModelTagReply{Model: &pb.ModelTag{TagId: 1}}, nil),
	)

	s := NewOrg(repo, middle)

	type args struct {
		ctx     context.Context
		payload *pb.ObjectiveRequest
	}
	tests := []struct {
		name    string
		o       pb.OrgSvcServer
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{"case1", s, args{context.Background(), &pb.ObjectiveRequest{Objective: &pb.Objective{
			Name: "obj1",
			Tag:  "test",
		}}}, &pb.StateReply{State: true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.CreateObjective(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Org.CreateObjective() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Org.CreateObjective() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrg_GetObjective(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	now := time.Now().Unix()
	repo := mock.NewMockOrgRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().GetObjectiveByID(gomock.Any(), gomock.Any()).Return(&pb.Objective{
			Id:   1,
			Name: "obj1",
			//Tag:       "test",
			CreatedAt: now,
		}, nil),
	)

	s := NewOrg(repo, nil)

	type args struct {
		ctx     context.Context
		payload *pb.ObjectiveRequest
	}
	tests := []struct {
		name    string
		o       pb.OrgSvcServer
		args    args
		want    *pb.ObjectiveReply
		wantErr bool
	}{
		{"case1", s, args{context.Background(), &pb.ObjectiveRequest{Objective: &pb.Objective{Id: 1}}},
			&pb.ObjectiveReply{Objective: &pb.Objective{
				Id:   1,
				Name: "obj1",
				//Tag:       "test",
				CreatedAt: now,
			}}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.GetObjective(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Org.GetObjective() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Org.GetObjective() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrg_GetObjectives(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	now := time.Now().Unix()
	repo := mock.NewMockOrgRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().ListObjectives(gomock.Any()).Return([]*pb.Objective{
			{
				Id:   1,
				Name: "obj1",
				//Tag:       "test",
				CreatedAt: now,
			},
		}, nil),
	)

	s := NewOrg(repo, nil)

	type args struct {
		ctx     context.Context
		payload *pb.ObjectiveRequest
	}
	tests := []struct {
		name    string
		o       pb.OrgSvcServer
		args    args
		want    *pb.ObjectivesReply
		wantErr bool
	}{
		{"case1", s, args{context.Background(), &pb.ObjectiveRequest{}}, &pb.ObjectivesReply{Objective: []*pb.Objective{
			{
				Id:   1,
				Name: "obj1",
				//Tag:       "test",
				CreatedAt: now,
			},
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.GetObjectives(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Org.GetObjectives() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Org.GetObjectives() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrg_DeleteObjective(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockOrgRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().DeleteObjective(gomock.Any(), gomock.Any()).Return(nil),
	)

	s := NewOrg(repo, nil)

	type args struct {
		ctx     context.Context
		payload *pb.ObjectiveRequest
	}
	tests := []struct {
		name    string
		o       pb.OrgSvcServer
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{"case1", s, args{context.Background(),
			&pb.ObjectiveRequest{Objective: &pb.Objective{Id: 1}}}, &pb.StateReply{State: true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.DeleteObjective(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Org.DeleteObjective() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Org.DeleteObjective() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrg_CreateKeyResult(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockOrgRepository(ctl)
	middle := mock.NewMockMiddleSvcClient(ctl)
	gomock.InOrder(
		repo.EXPECT().CreateKeyResult(gomock.Any(), gomock.Any()).Return(int64(1), nil),
		middle.EXPECT().SaveModelTag(gomock.Any(), gomock.Any()).Return(&pb.ModelTagReply{Model: &pb.ModelTag{TagId: 1}}, nil),
	)

	s := NewOrg(repo, middle)

	type args struct {
		ctx     context.Context
		payload *pb.KeyResultRequest
	}
	tests := []struct {
		name    string
		o       pb.OrgSvcServer
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{"case1", s, args{context.Background(), &pb.KeyResultRequest{KeyResult: &pb.KeyResult{
			ObjectiveId: 1,
			Name:        "obj1",
			Tag:         "test",
		}}}, &pb.StateReply{State: true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.CreateKeyResult(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Org.CreateKeyResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Org.CreateKeyResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrg_GetKeyResult(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	now := time.Now().Unix()
	repo := mock.NewMockOrgRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().GetKeyResultByID(gomock.Any(), gomock.Any()).Return(&pb.KeyResult{
			Id:          1,
			ObjectiveId: 1,
			Name:        "obj1",
			//Tag:         "test",
			CreatedAt: now,
		}, nil),
	)

	s := NewOrg(repo, nil)

	type args struct {
		ctx     context.Context
		payload *pb.KeyResultRequest
	}
	tests := []struct {
		name    string
		o       pb.OrgSvcServer
		args    args
		want    *pb.KeyResultReply
		wantErr bool
	}{
		{"case1", s, args{context.Background(), &pb.KeyResultRequest{KeyResult: &pb.KeyResult{Id: 1}}},
			&pb.KeyResultReply{KeyResult: &pb.KeyResult{
				Id:          1,
				ObjectiveId: 1,
				Name:        "obj1",
				//Tag:         "test",
				CreatedAt: now,
			}}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.GetKeyResult(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Org.GetKeyResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Org.GetKeyResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrg_GetKeyResults(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	now := time.Now().Unix()
	repo := mock.NewMockOrgRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().ListKeyResults(gomock.Any()).Return([]*pb.KeyResult{
			{
				Id:          1,
				ObjectiveId: 1,
				Name:        "obj1",
				//Tag:         "test",
				CreatedAt: now,
			},
		}, nil),
	)

	s := NewOrg(repo, nil)

	type args struct {
		ctx     context.Context
		payload *pb.KeyResultRequest
	}
	tests := []struct {
		name    string
		o       pb.OrgSvcServer
		args    args
		want    *pb.KeyResultsReply
		wantErr bool
	}{
		{"case1", s, args{context.Background(), &pb.KeyResultRequest{}}, &pb.KeyResultsReply{Result: []*pb.KeyResult{
			{
				Id:          1,
				ObjectiveId: 1,
				Name:        "obj1",
				//Tag:         "test",
				CreatedAt: now,
			},
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.GetKeyResults(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Org.GetKeyResults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Org.GetKeyResults() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrg_DeleteKeyResult(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockOrgRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().DeleteKeyResult(gomock.Any(), gomock.Any()).Return(nil),
	)

	s := NewOrg(repo, nil)

	type args struct {
		ctx     context.Context
		payload *pb.KeyResultRequest
	}
	tests := []struct {
		name    string
		o       pb.OrgSvcServer
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{"case1", s, args{context.Background(),
			&pb.KeyResultRequest{KeyResult: &pb.KeyResult{Id: 1}}}, &pb.StateReply{State: true}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.DeleteKeyResult(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Org.DeleteKeyResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Org.DeleteKeyResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
