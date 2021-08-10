package repository

import (
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"testing"
)

func TestOrgRepository_CreateObjective(t *testing.T) {
	sto, err := CreateOrgRepository(enum.Org)
	if err != nil {
		t.Fatalf("create org Repository err %+v", err)
	}

	type args struct {
		objective pb.Objective
	}
	tests := []struct {
		name    string
		r       OrgRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{objective: pb.Objective{
			Name: "obj1",
			TagId:  1,
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.CreateObjective(tt.args.objective)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrgRepository.CreateObjective() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestOrgRepository_GetObjectiveByID(t *testing.T) {
	sto, err := CreateOrgRepository(enum.Org)
	if err != nil {
		t.Fatalf("create org Repository err %+v", err)
	}

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		r       OrgRepository
		args    args
		wantErr bool
	}{
		{"id=1", sto, args{id: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetObjectiveByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrgRepository.GetObjectiveByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestOrgRepository_ListObjectives(t *testing.T) {
	sto, err := CreateOrgRepository(enum.Org)
	if err != nil {
		t.Fatalf("create org Repository err %+v", err)
	}

	tests := []struct {
		name    string
		r       OrgRepository
		wantErr bool
	}{
		{"list", sto, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.ListObjectives()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrgRepository.ListObjectives() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestOrgRepository_DeleteObjective(t *testing.T) {
	sto, err := CreateOrgRepository(enum.Org)
	if err != nil {
		t.Fatalf("create org Repository err %+v", err)
	}

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		r       OrgRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{id: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.DeleteObjective(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("OrgRepository.DeleteObjective() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrgRepository_CreateKeyResult(t *testing.T) {
	sto, err := CreateOrgRepository(enum.Org)
	if err != nil {
		t.Fatalf("create org Repository err %+v", err)
	}

	type args struct {
		keyResult pb.KeyResult
	}
	tests := []struct {
		name    string
		r       OrgRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{keyResult: pb.KeyResult{ObjectiveId: 1, Name: "kr1", TagId: 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.CreateKeyResult(tt.args.keyResult)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrgRepository.CreateKeyResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestOrgRepository_GetKeyResultByID(t *testing.T) {
	sto, err := CreateOrgRepository(enum.Org)
	if err != nil {
		t.Fatalf("create org Repository err %+v", err)
	}

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		r       OrgRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{id: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetKeyResultByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrgRepository.GetKeyResultByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestOrgRepository_ListKeyResults(t *testing.T) {
	sto, err := CreateOrgRepository(enum.Org)
	if err != nil {
		t.Fatalf("create org Repository err %+v", err)
	}

	tests := []struct {
		name    string
		r       OrgRepository
		wantErr bool
	}{
		{"case1", sto, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.ListKeyResults()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrgRepository.ListKeyResults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestOrgRepository_DeleteKeyResult(t *testing.T) {
	sto, err := CreateOrgRepository(enum.Org)
	if err != nil {
		t.Fatalf("create org Repository err %+v", err)
	}

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		r       OrgRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{id: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.DeleteKeyResult(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("OrgRepository.DeleteKeyResult() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
