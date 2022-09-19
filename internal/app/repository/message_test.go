package repository

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestMessageRepository_Create(t *testing.T) {
	sto, err := CreateMessageRepository(enum.Message)
	if err != nil {
		t.Fatalf("create message Repository error, %+v", err)
	}

	type args struct {
		message *pb.Message
	}
	tests := []struct {
		name    string
		r       MessageRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{message: &pb.Message{Text: "test1"}}, false},
		{"case2", sto, args{message: &pb.Message{Text: "test2"}}, false},
		{"case3", sto, args{message: &pb.Message{Text: "test3"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.Create(context.Background(), tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("MessageRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMessageRepository_GetByID(t *testing.T) {
	sto, err := CreateMessageRepository(enum.Message)
	if err != nil {
		t.Fatalf("create message Repository error, %+v", err)
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
		{"id=99999", sto, args{id: 99999}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetByID(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("MessageRepository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMessageRepository_ListByType(t *testing.T) {
	sto, err := CreateMessageRepository(enum.Message)
	if err != nil {
		t.Fatalf("create message Repository error, %+v", err)
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
			_, err := tt.r.ListByType(context.Background(), tt.args.t)
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
		t.Fatalf("create message Repository error, %+v", err)
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
			_, err := tt.r.List(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("MessageRepository.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMessageRepository_Delete(t *testing.T) {
	sto, err := CreateMessageRepository(enum.Message)
	if err != nil {
		t.Fatalf("create message Repository error, %+v", err)
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
			if err := tt.r.Delete(context.Background(), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("MessageRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
