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
	uuid1, uuid2, uuid3                   string
	identifier1, identifier2, identifier3 string
)

func TestMain(m *testing.M) {
	uuid1 = "f5ca59e9-25c0-4cf7-a370-b4ca3a173c5b"
	uuid2 = "c75345a7-1035-40bd-a11b-04af13b1a0fc"
	uuid3 = "68ee68b3-91a2-4d2c-80b5-64962d2efdaf"
	identifier1 = util.RandString(8, "lowercase") + "_bot"
	identifier2 = util.RandString(8, "lowercase") + "_bot"
	identifier3 = util.RandString(8, "lowercase") + "_bot"
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

func TestChatbotRepository_CreateGroup(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create chatbot Repository error, %+v", err)
	}

	type args struct {
		group *pb.Group
	}
	tests := []struct {
		name    string
		r       ChatbotRepository
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
				t.Errorf("ChatbotRepository.CreateGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestChatbotRepository_GetGroup(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create chatbot Repository error, %+v", err)
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
			_, err := tt.r.GetGroup(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatbotRepository.GetGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestChatbotRepository_GetGroupByUUID(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create chatbot Repository error, %+v", err)
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
		{"uuid=99999", sto, args{uuid: "no_uuid"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetGroupByUUID(context.Background(), tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatbotRepository.GetGroupByUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestChatbotRepository_GetGroupBySequence(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create chatbot Repository error, %+v", err)
	}

	type args struct {
		userId   int64
		sequence int64
	}
	tests := []struct {
		name    string
		r       ChatbotRepository
		args    args
		wantErr bool
	}{
		{"id=1", sto, args{userId: 1, sequence: 1}, false},
		{"id=2", sto, args{userId: 1, sequence: 2}, false},
		{"id=3", sto, args{userId: 1, sequence: 3}, false},
		{"id=99999", sto, args{userId: 1, sequence: 99999}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetGroupBySequence(context.Background(), tt.args.userId, tt.args.sequence)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatbotRepository.GetGroupBySequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestChatbotRepository_ListGroup(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create chatbot Repository error, %+v", err)
	}
	type args struct {
		userId int64
	}
	tests := []struct {
		name    string
		r       ChatbotRepository
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
				t.Errorf("ChatbotRepository.ListGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestChatbotRepository_DeleteGroup(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create chatbot Repository error, %+v", err)
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
			if err := tt.r.DeleteGroup(context.Background(), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("ChatbotRepository.DeleteGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChatbotRepository_ListGroupBot(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create chatbot Repository error, %+v", err)
	}

	type args struct {
		groupId int64
	}
	tests := []struct {
		name    string
		r       ChatbotRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.ListGroupBot(context.Background(), tt.args.groupId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatbotRepository.ListGroupBot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestChatbotRepository_CreateGroupBot(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create chatbot Repository error, %+v", err)
	}

	type args struct {
		groupId int64
		bot     *pb.Bot
	}
	tests := []struct {
		name    string
		r       ChatbotRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{1, &pb.Bot{Id: 1, Name: "a_bot"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.CreateGroupBot(context.Background(), tt.args.groupId, tt.args.bot); (err != nil) != tt.wantErr {
				t.Errorf("ChatbotRepository.CreateGroupBot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChatbotRepository_DeleteGroupBot(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create chatbot Repository error, %+v", err)
	}

	type args struct {
		groupId int64
		botId   int64
	}
	tests := []struct {
		name    string
		r       ChatbotRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{1, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.DeleteGroupBot(context.Background(), tt.args.groupId, tt.args.botId); (err != nil) != tt.wantErr {
				t.Errorf("ChatbotRepository.DeleteGroupBot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChatbotRepository_UpdateGroup(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create chatbot Repository error, %+v", err)
	}

	type args struct {
		group *pb.Group
	}
	tests := []struct {
		name    string
		r       ChatbotRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{&pb.Group{Id: 1, Name: "update"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.UpdateGroup(context.Background(), tt.args.group); (err != nil) != tt.wantErr {
				t.Errorf("ChatbotRepository.UpdateGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChatbotRepository_UpdateGroupSetting(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create chatbot Repository error, %+v", err)
	}

	type args struct {
		groupId int64
		kvs     []*pb.KV
	}
	tests := []struct {
		name    string
		r       ChatbotRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{1, []*pb.KV{{Key: "k", Value: "v"}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.UpdateGroupSetting(context.Background(), tt.args.groupId, tt.args.kvs); (err != nil) != tt.wantErr {
				t.Errorf("ChatbotRepository.UpdateGroupSetting() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChatbotRepository_UpdateGroupBotSetting(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create chatbot Repository error, %+v", err)
	}

	type args struct {
		groupId int64
		botId   int64
		kvs     []*pb.KV
	}
	tests := []struct {
		name    string
		r       ChatbotRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{1, 1, []*pb.KV{{Key: "k", Value: "u"}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.UpdateGroupBotSetting(context.Background(), tt.args.groupId, tt.args.botId, tt.args.kvs); (err != nil) != tt.wantErr {
				t.Errorf("ChatbotRepository.UpdateGroupBotSetting() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChatbotRepository_ListGroupTag(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create chatbot Repository error, %+v", err)
	}

	type args struct {
		groupId int64
	}
	tests := []struct {
		name    string
		r       ChatbotRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.ListGroupTag(context.Background(), tt.args.groupId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatbotRepository.ListGroupTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestChatbotRepository_CreateGroupTag(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create chatbot Repository error, %+v", err)
	}

	type args struct {
		tag *pb.GroupTag
	}
	tests := []struct {
		name    string
		r       ChatbotRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{&pb.GroupTag{GroupId: 1, Tag: "t"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.CreateGroupTag(context.Background(), tt.args.tag)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatbotRepository.CreateGroupTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestChatbotRepository_DeleteGroupTag(t *testing.T) {
	sto, err := CreateChatbotRepository(enum.Chatbot)
	if err != nil {
		t.Fatalf("create chatbot Repository error, %+v", err)
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
		{"case1", sto, args{1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.DeleteGroupTag(context.Background(), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("ChatbotRepository.DeleteGroupTag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
