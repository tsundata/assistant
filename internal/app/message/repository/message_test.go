package repository

import (
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/util"
	"testing"
)

func TestMessageRepository_GetByID(t *testing.T) {
	sto, err := CreateMessageRepository(enum.Message)
	if err != nil {
		t.Fatalf("create message Preposiory error, %+v", err)
	}

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		r       MessageRepository
		args    args
		wantErr bool
	}{
		{"id=1", sto, args{id: 1}, false},
		{"id=2", sto, args{id: 2}, false},
		{"id=3", sto, args{id: 3}, false},
		{"id=99999", sto, args{id: 99999}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("MessageRepository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMessageRepository_GetByUUID(t *testing.T) {
	sto, err := CreateMessageRepository(enum.Message)
	if err != nil {
		t.Fatalf("create message Preposiory error, %+v", err)
	}

	type args struct {
		uuid string
	}
	tests := []struct {
		name    string
		r       MessageRepository
		args    args
		wantErr bool
	}{
		{"uuid=1", sto, args{uuid: "1"}, false},
		{"uuid=2", sto, args{uuid: "2"}, false},
		{"uuid=3", sto, args{uuid: "3"}, false},
		{"uuid=ff4103db-d554-4f22-b6c7-57a3f708d5eb", sto, args{uuid: "ff4103db-d554-4f22-b6c7-57a3f708d5eb"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetByUUID(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("MessageRepository.GetByUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMessageRepository_ListByType(t *testing.T) {
	sto, err := CreateMessageRepository(enum.Message)
	if err != nil {
		t.Fatalf("create message Preposiory error, %+v", err)
	}

	type args struct {
		t string
	}
	tests := []struct {
		name    string
		r       MessageRepository
		args    args
		wantErr bool
	}{
		{"t=text", sto, args{t: "text"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.ListByType(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("MessageRepository.ListByType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMessageRepository_List(t *testing.T) {
	sto, err := CreateMessageRepository(enum.Message)
	if err != nil {
		t.Fatalf("create message Preposiory error, %+v", err)
	}

	tests := []struct {
		name    string
		r       MessageRepository
		wantErr bool
	}{
		{"case1", sto, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("MessageRepository.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMessageRepository_Create(t *testing.T) {
	sto, err := CreateMessageRepository(enum.Message)
	if err != nil {
		t.Fatalf("create message Preposiory error, %+v", err)
	}

	uuid, err := util.GenerateUUID()
	if err != nil {
		t.Fatalf("generate uuid error, %+v", err)
	}

	type args struct {
		message pb.Message
	}
	tests := []struct {
		name    string
		r       MessageRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{message: pb.Message{Uuid: uuid, Text: "test"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.Create(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("MessageRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMessageRepository_Delete(t *testing.T) {
	sto, err := CreateMessageRepository(enum.Message)
	if err != nil {
		t.Fatalf("create message Preposiory error, %+v", err)
	}

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		r       MessageRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{id: 100}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("MessageRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
