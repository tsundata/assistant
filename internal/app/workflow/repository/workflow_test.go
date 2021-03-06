package repository

import (
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/model"
	"testing"
)

func TestWorkflowRepository_GetTriggerByFlag(t *testing.T) {
	sto, err := CreateWorkflowRepository(app.Workflow)
	if err != nil {
		t.Fatalf("create workflow Preposiory error, %+v", err)
	}
	type args struct {
		t    string
		flag string
	}
	tests := []struct {
		name    string
		r       WorkflowRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{t: "1", flag: ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetTriggerByFlag(tt.args.t, tt.args.flag)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlWorkflowRepository.GetTriggerByFlag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestWorkflowRepository_ListTriggersByType(t *testing.T) {
	sto, err := CreateWorkflowRepository(app.Workflow)
	if err != nil {
		t.Fatalf("create workflow Preposiory error, %+v", err)
	}
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		r       WorkflowRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{t: "1"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.ListTriggersByType(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlWorkflowRepository.ListTriggersByType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestWorkflowRepository_CreateTrigger(t *testing.T) {
	sto, err := CreateWorkflowRepository(app.Workflow)
	if err != nil {
		t.Fatalf("create workflow Preposiory error, %+v", err)
	}
	type args struct {
		trigger model.Trigger
	}
	tests := []struct {
		name    string
		r       WorkflowRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{trigger: model.Trigger{Type: "test"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.CreateTrigger(tt.args.trigger)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlWorkflowRepository.CreateTrigger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestWorkflowRepository_DeleteTriggerByMessageID(t *testing.T) {
	sto, err := CreateWorkflowRepository(app.Workflow)
	if err != nil {
		t.Fatalf("create workflow Preposiory error, %+v", err)
	}
	type args struct {
		messageID int64
	}
	tests := []struct {
		name    string
		r       WorkflowRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{messageID: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.DeleteTriggerByMessageID(tt.args.messageID); (err != nil) != tt.wantErr {
				t.Errorf("MysqlWorkflowRepository.DeleteTriggerByMessageID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
