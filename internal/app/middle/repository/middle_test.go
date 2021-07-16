package repository

import (
	"github.com/tsundata/assistant/api/model"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/util"
	"testing"
)

func TestMiddleRepository_CreatePage(t *testing.T) {
	sto, err := CreateMiddleRepository(app.Middle)
	if err != nil {
		t.Fatalf("create middle Preposiory error, %+v", err)
	}
	uuid, _ := util.GenerateUUID()
	type args struct {
		page pb.Page
	}
	tests := []struct {
		name    string
		r       MiddleRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{page: pb.Page{Uuid: uuid}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.CreatePage(tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.CreatePage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddleRepository_GetPageByUUID(t *testing.T) {
	sto, err := CreateMiddleRepository(app.Middle)
	if err != nil {
		t.Fatalf("create middle Preposiory error, %+v", err)
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
		{"case1", sto, args{uuid: "1"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetPageByUUID(tt.args.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.GetPageByUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddleRepository_ListApps(t *testing.T) {
	sto, err := CreateMiddleRepository(app.Middle)
	if err != nil {
		t.Fatalf("create middle Preposiory error, %+v", err)
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
			_, err := tt.r.ListApps()
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.ListApps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddleRepository_GetAvailableAppByType(t *testing.T) {
	sto, err := CreateMiddleRepository(app.Middle)
	if err != nil {
		t.Fatalf("create middle Preposiory error, %+v", err)
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
		{"case1", sto, args{t: "1"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetAvailableAppByType(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.GetAvailableAppByType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddleRepository_GetAppByType(t *testing.T) {
	sto, err := CreateMiddleRepository(app.Middle)
	if err != nil {
		t.Fatalf("create middle Preposiory error, %+v", err)
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
		{"case1", sto, args{t: "1"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetAppByType(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.GetAppByType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddleRepository_UpdateAppByID(t *testing.T) {
	sto, err := CreateMiddleRepository(app.Middle)
	if err != nil {
		t.Fatalf("create middle Preposiory error, %+v", err)
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
			if err := tt.r.UpdateAppByID(tt.args.id, tt.args.token, tt.args.extra); (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.UpdateAppByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMiddleRepository_CreateApp(t *testing.T) {
	sto, err := CreateMiddleRepository(app.Middle)
	if err != nil {
		t.Fatalf("create middle Preposiory error, %+v", err)
	}
	type args struct {
		app model.App
	}
	tests := []struct {
		name    string
		r       MiddleRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{app: model.App{Name: "test"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.CreateApp(tt.args.app)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.CreateApp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddleRepository_GetCredentialByName(t *testing.T) {
	sto, err := CreateMiddleRepository(app.Middle)
	if err != nil {
		t.Fatalf("create middle Preposiory error, %+v", err)
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
		{"case1", sto, args{name: "1"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetCredentialByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.GetCredentialByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddleRepository_GetCredentialByType(t *testing.T) {
	sto, err := CreateMiddleRepository(app.Middle)
	if err != nil {
		t.Fatalf("create middle Preposiory error, %+v", err)
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
		{"case1", sto, args{t: "1"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetCredentialByType(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.GetCredentialByType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddleRepository_ListCredentials(t *testing.T) {
	sto, err := CreateMiddleRepository(app.Middle)
	if err != nil {
		t.Fatalf("create middle Preposiory error, %+v", err)
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
			_, err := tt.r.ListCredentials()
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.ListCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddleRepository_CreateCredential(t *testing.T) {
	sto, err := CreateMiddleRepository(app.Middle)
	if err != nil {
		t.Fatalf("create middle Preposiory error, %+v", err)
	}
	name := util.GeneratePassword(10, "lowercase")
	type args struct {
		credential model.Credential
	}
	tests := []struct {
		name    string
		r       MiddleRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{credential: model.Credential{Name: name}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.CreateCredential(tt.args.credential)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlMiddleRepository.CreateCredential() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
