package repository

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/util"
	"testing"
)

func TestMiddleRepository_CreatePage(t *testing.T) {
	sto, err := CreateMiddleRepository(enum.Middle)
	if err != nil {
		t.Fatalf("create middle Repository error, %+v", err)
	}
	uuid := util.UUID()
	type args struct {
		page *pb.Page
	}
	tests := []struct {
		name    string
		r       MiddleRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{page: &pb.Page{Uuid: uuid}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.CreatePage(context.Background(), tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.CreatePage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddleRepository_GetPageByUUID(t *testing.T) {
	sto, err := CreateMiddleRepository(enum.Middle)
	if err != nil {
		t.Fatalf("create middle Repository error, %+v", err)
	}
	type args struct {
		uuid string
	}
	tests := []struct {
		name    string
		r       MiddleRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{uuid: "1"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetPageByUUID(context.Background(), tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.GetPageByUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddleRepository_ListApps(t *testing.T) {
	sto, err := CreateMiddleRepository(enum.Middle)
	if err != nil {
		t.Fatalf("create middle Repository error, %+v", err)
	}
	type args struct {
		userId int64
	}
	tests := []struct {
		name    string
		r       MiddleRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{userId: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.ListApps(context.Background(), tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.ListApps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddleRepository_GetAvailableAppByType(t *testing.T) {
	sto, err := CreateMiddleRepository(enum.Middle)
	if err != nil {
		t.Fatalf("create middle Repository error, %+v", err)
	}
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		r       MiddleRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{t: "1"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetAvailableAppByType(context.Background(), tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.GetAvailableAppByType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddleRepository_GetAppByType(t *testing.T) {
	sto, err := CreateMiddleRepository(enum.Middle)
	if err != nil {
		t.Fatalf("create middle Repository error, %+v", err)
	}
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		r       MiddleRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{t: "1"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetAppByType(context.Background(), tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.GetAppByType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddleRepository_UpdateAppByID(t *testing.T) {
	sto, err := CreateMiddleRepository(enum.Middle)
	if err != nil {
		t.Fatalf("create middle Repository error, %+v", err)
	}
	type args struct {
		id    int64
		token string
		extra string
	}
	tests := []struct {
		name    string
		r       MiddleRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{id: 1, token: "test", extra: "{}"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.UpdateAppByID(context.Background(), tt.args.id, tt.args.token, tt.args.extra); (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.UpdateAppByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMiddleRepository_CreateApp(t *testing.T) {
	sto, err := CreateMiddleRepository(enum.Middle)
	if err != nil {
		t.Fatalf("create middle Repository error, %+v", err)
	}
	type args struct {
		app *pb.App
	}
	tests := []struct {
		name    string
		r       MiddleRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{app: &pb.App{UserId: 1, Name: "test"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.CreateApp(context.Background(), tt.args.app)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.CreateApp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddleRepository_GetCredentialByName(t *testing.T) {
	sto, err := CreateMiddleRepository(enum.Middle)
	if err != nil {
		t.Fatalf("create middle Repository error, %+v", err)
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		r       MiddleRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{name: "1"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetCredentialByName(context.Background(), 1, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.GetCredentialByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddleRepository_GetCredentialByType(t *testing.T) {
	sto, err := CreateMiddleRepository(enum.Middle)
	if err != nil {
		t.Fatalf("create middle Repository error, %+v", err)
	}
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		r       MiddleRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{t: "1"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetCredentialByType(context.Background(), 1, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.GetCredentialByType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddleRepository_ListCredentials(t *testing.T) {
	sto, err := CreateMiddleRepository(enum.Middle)
	if err != nil {
		t.Fatalf("create middle Repository error, %+v", err)
	}
	tests := []struct {
		name    string
		r       MiddleRepository
		wantErr bool
	}{
		{"case1", sto, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.ListCredentials(context.Background(), 1)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.ListCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddleRepository_CreateCredential(t *testing.T) {
	sto, err := CreateMiddleRepository(enum.Middle)
	if err != nil {
		t.Fatalf("create middle Repository error, %+v", err)
	}
	name := util.RandString(10, "lowercase")
	type args struct {
		credential *pb.Credential
	}
	tests := []struct {
		name    string
		r       MiddleRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{credential: &pb.Credential{Name: name}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.CreateCredential(context.Background(), tt.args.credential)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.CreateCredential() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddleRepository_ListTags(t *testing.T) {
	sto, err := CreateMiddleRepository(enum.Middle)
	if err != nil {
		t.Fatalf("create middle Repository error, %+v", err)
	}

	tests := []struct {
		name    string
		r       MiddleRepository
		wantErr bool
	}{
		{"case1", sto, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.ListTags(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.ListTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddleRepository_GetOrCreateTag(t *testing.T) {
	sto, err := CreateMiddleRepository(enum.Middle)
	if err != nil {
		t.Fatalf("create middle Repository error, %+v", err)
	}

	type args struct {
		tag *pb.Tag
	}
	tests := []struct {
		name    string
		r       MiddleRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{tag: &pb.Tag{Name: "test1"}}, false},
		{"case2", sto, args{tag: &pb.Tag{Name: "test2"}}, false},
		{"case3", sto, args{tag: &pb.Tag{Name: "test2"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetOrCreateTag(context.Background(), tt.args.tag)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.GetOrCreateTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
