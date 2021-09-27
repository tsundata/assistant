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

func TestChatbotRepository_Create(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create bot Repository error, %+v", err)
	}

	type args struct {
		bot *pb.Bot
	}
	tests := []struct {
		name    string
		r       ChatbotRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{bot: &pb.Bot{Uuid: uuid1, Name: "test1"}}, false},
		{"case2", sto, args{bot: &pb.Bot{Uuid: uuid2, Name: "test2"}}, false},
		{"case3", sto, args{bot: &pb.Bot{Uuid: uuid3, Name: "test3"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.Create(context.Background(), tt.args.bot)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatbotRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestChatbotRepository_GetByID(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create bot Repository error, %+v", err)
	}

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		r       ChatbotRepository
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
				t.Errorf("ChatbotRepository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestChatbotRepository_GetByUUID(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create bot Repository error, %+v", err)
	}

	type args struct {
		uuid string
	}
	tests := []struct {
		name    string
		r       ChatbotRepository
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
				t.Errorf("ChatbotRepository.GetByUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestChatbotRepository_List(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create bot Repository error, %+v", err)
	}

	tests := []struct {
		name    string
		r       ChatbotRepository
		wantErr bool
	}{
		{"case1", sto, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.List(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatbotRepository.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestChatbotRepository_Delete(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create bot Repository error, %+v", err)
	}

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		r       ChatbotRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{id: 100}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Delete(context.Background(), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("ChatbotRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
