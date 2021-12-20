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
	identifier1, identifier2, identifier3 string
)

func TestMain(m *testing.M) {
	uuid1 = util.UUID()
	uuid2 = util.UUID()
	uuid3 = util.UUID()
	identifier1 = "a_bot"
	identifier2 = "b_bot"
	identifier3 = "c_bot"
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
		{"case1", sto, args{bot: &pb.Bot{Uuid: uuid1, Name: "test1", Identifier: identifier1}}, false},
		{"case2", sto, args{bot: &pb.Bot{Uuid: uuid2, Name: "test2", Identifier: identifier2}}, false},
		{"case3", sto, args{bot: &pb.Bot{Uuid: uuid3, Name: "test3", Identifier: identifier3}}, false},
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

func TestChatbotRepository_GetByIdentifier(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create bot Repository error, %+v", err)
	}

	type args struct {
		identifier string
	}
	tests := []struct {
		name    string
		r       ChatbotRepository
		args    args
		wantErr bool
	}{
		{"identifier=1", sto, args{identifier: identifier1}, false},
		{"identifier=2", sto, args{identifier: identifier2}, false},
		{"identifier=3", sto, args{identifier: identifier3}, false},
		{"identifier=404", sto, args{identifier: "404"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetByIdentifier(context.Background(), tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatbotRepository.GetByIdentifier() error = %v, wantErr %v", err, tt.wantErr)
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
