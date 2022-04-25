package repository

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"testing"
)

func TestOkrRepository_CreateObjective(t *testing.T) {
	sto, err := CreateOkrRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create okr Repository err %+v", err)
	}

	type args struct {
		objective *pb.Objective
	}
	tests := []struct {
		name    string
		r       OkrRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{objective: &pb.Objective{
			Title: "obj1",
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.CreateObjective(context.Background(), tt.args.objective)
			if (err != nil) != tt.wantErr {
				t.Errorf("OkrRepository.CreateObjective() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestOkrRepository_GetObjectiveByID(t *testing.T) {
	sto, err := CreateOkrRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create okr Repository err %+v", err)
	}

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		r       OkrRepository
		args    args
		wantErr bool
	}{
		{"id=1", sto, args{id: 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetObjectiveByID(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("OkrRepository.GetObjectiveByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestOkrRepository_ListObjectives(t *testing.T) {
	sto, err := CreateOkrRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create okr Repository err %+v", err)
	}

	tests := []struct {
		name    string
		r       OkrRepository
		wantErr bool
	}{
		{"list", sto, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.ListObjectives(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("OkrRepository.ListObjectives() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestOkrRepository_DeleteObjective(t *testing.T) {
	sto, err := CreateOkrRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create okr Repository err %+v", err)
	}

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		r       OkrRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{id: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.DeleteObjective(context.Background(), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("OkrRepository.DeleteObjective() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOkrRepository_CreateKeyResult(t *testing.T) {
	sto, err := CreateOkrRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create okr Repository err %+v", err)
	}

	type args struct {
		keyResult *pb.KeyResult
	}
	tests := []struct {
		name    string
		r       OkrRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{keyResult: &pb.KeyResult{ObjectiveId: 1, Title: "kr1"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.CreateKeyResult(context.Background(), tt.args.keyResult)
			if (err != nil) != tt.wantErr {
				t.Errorf("OkrRepository.CreateKeyResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestOkrRepository_GetKeyResultByID(t *testing.T) {
	sto, err := CreateOkrRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create okr Repository err %+v", err)
	}

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		r       OkrRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{id: 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetKeyResultByID(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("OkrRepository.GetKeyResultByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestOkrRepository_ListKeyResults(t *testing.T) {
	sto, err := CreateOkrRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create okr Repository err %+v", err)
	}

	tests := []struct {
		name    string
		r       OkrRepository
		wantErr bool
	}{
		{"case1", sto, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.ListKeyResults(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("OkrRepository.ListKeyResults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestOkrRepository_DeleteKeyResult(t *testing.T) {
	sto, err := CreateOkrRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create okr Repository err %+v", err)
	}

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		r       OkrRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{id: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.DeleteKeyResult(context.Background(), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("OkrRepository.DeleteKeyResult() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
