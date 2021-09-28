package repository

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/util"
	"os"
	"testing"
)

var (
	uuid1, uuid2, uuid3 string
)

func TestMain(m *testing.M) {
	uuid1 = util.UUID()
	uuid2 = util.UUID()
	uuid3 = util.UUID()
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
		{"case1", sto, args{message: &pb.Message{Uuid: uuid1, Text: "test1"}}, false},
		{"case2", sto, args{message: &pb.Message{Uuid: uuid2, Text: "test2"}}, false},
		{"case3", sto, args{message: &pb.Message{Uuid: uuid3, Text: "test3"}}, false},
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
		{"id=1", sto, args{id: 1}, false},
		{"id=2", sto, args{id: 2}, false},
		{"id=3", sto, args{id: 3}, false},
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

func TestMessageRepository_GetByUUID(t *testing.T) {
	sto, err := CreateMessageRepository(enum.Message)
	if err != nil {
		t.Fatalf("create message Repository error, %+v", err)
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
		{"uuid=1", sto, args{uuid: uuid1}, false},
		{"uuid=2", sto, args{uuid: uuid2}, false},
		{"uuid=3", sto, args{uuid: uuid3}, false},
		{"uuid=ff4103db-d554-4f22-b6c7-57a3f708d5eb", sto, args{uuid: "ff4103db-d554-4f22-b6c7-57a3f708d5eb"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetByUUID(context.Background(), tt.args.uuid)
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

func TestMessageRepository_CreateGroup(t *testing.T) {
	sto, err := CreateMessageRepository(enum.Message)
	if err != nil {
		t.Fatalf("create message Repository error, %+v", err)
	}

	type args struct {
		group *pb.Group
	}
	tests := []struct {
		name    string
		r       MessageRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{group: &pb.Group{Uuid: uuid1, Name: "test1", UserId: 1}}, false},
		{"case2", sto, args{group: &pb.Group{Uuid: uuid2, Name: "test2", UserId: 1}}, false},
		{"case3", sto, args{group: &pb.Group{Uuid: uuid3, Name: "test3", UserId: 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.CreateGroup(context.Background(), tt.args.group)
			if (err != nil) != tt.wantErr {
				t.Errorf("MessageRepository.CreateGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMessageRepository_GetGroup(t *testing.T) {
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
		{"id=1", sto, args{id: 1}, false},
		{"id=2", sto, args{id: 2}, false},
		{"id=3", sto, args{id: 3}, false},
		{"id=99999", sto, args{id: 99999}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetGroup(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("MessageRepository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMessageRepository_GetGroupByUUID(t *testing.T) {
	sto, err := CreateMessageRepository(enum.Message)
	if err != nil {
		t.Fatalf("create message Repository error, %+v", err)
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
		{"uuid=1", sto, args{uuid: uuid1}, false},
		{"uuid=2", sto, args{uuid: uuid2}, false},
		{"uuid=3", sto, args{uuid: uuid3}, false},
		{"uuid=99999", sto, args{uuid: "no_uuid"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetGroupByUUID(context.Background(), tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("MessageRepository.GetGroupByUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMessageRepository_ListGroup(t *testing.T) {
	sto, err := CreateMessageRepository(enum.Message)
	if err != nil {
		t.Fatalf("create message Repository error, %+v", err)
	}
	type args struct {
		userId int64
	}
	tests := []struct {
		name    string
		r       MessageRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{userId: 1}, false},
		{"case1", sto, args{userId: 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.ListGroup(context.Background(), tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("MessageRepository.ListGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMessageRepository_DeleteGroup(t *testing.T) {
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
			if err := tt.r.DeleteGroup(context.Background(), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("MessageRepository.DeleteGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
