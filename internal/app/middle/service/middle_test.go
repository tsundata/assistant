package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/vendors"
	"github.com/tsundata/assistant/mock"
	"reflect"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/pb"
)

func TestMiddle_GetMenu(t *testing.T) {
	conf, err := config.CreateAppConfig(app.Middle)
	if err != nil {
		t.Fatal(err)
	}
	rdb, err := vendors.CreateRedisClient(app.Middle)
	if err != nil {
		t.Fatal(err)
	}

	s := NewMiddle(conf, rdb, nil)

	type args struct {
		in0 context.Context
		in1 *pb.TextRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TextRequest{}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.s.GetMenu(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetMenu() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddle_GetQrUrl(t *testing.T) {
	conf, err := config.CreateAppConfig(app.Middle)
	if err != nil {
		t.Fatal(err)
	}

	s := NewMiddle(conf, nil, nil)

	type args struct {
		in0     context.Context
		payload *pb.TextRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TextRequest{Text: "test"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.s.GetQrUrl(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetQrUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddle_CreatePage(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	conf, err := config.CreateAppConfig(app.Middle)
	if err != nil {
		t.Fatal(err)
	}

	repo := mock.NewMockMiddleRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().CreatePage(gomock.Any()).Return(int64(1), nil),
	)

	s := NewMiddle(conf, nil, repo)

	type args struct {
		in0     context.Context
		payload *pb.PageRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.PageRequest{Type: "html", Title: "title", Content: "test"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.s.CreatePage(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.CreatePage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddle_GetPage(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockMiddleRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().GetPageByUUID(gomock.Any()).Return(model.Page{
			ID:      1,
			UUID:    "test",
			Type:    "html",
			Title:   "test",
			Content: "test",
			Time:    time.Now(),
		}, nil),
	)

	s := NewMiddle(nil, nil, repo)

	type args struct {
		in0     context.Context
		payload *pb.PageRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		want    *pb.PageReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.PageRequest{Uuid: "test"}},
			&pb.PageReply{
				Uuid:    "test",
				Title:   "test",
				Content: "test",
				Type:    "html",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetPage(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Middle.GetPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddle_GetApps(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockMiddleRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().ListApps().Return([]model.App{{
			ID:      1,
			Type:    "github",
			Time:    time.Now(),
		}}, nil),
	)

	s := NewMiddle(nil, nil, repo)

	type args struct {
		in0 context.Context
		in1 *pb.TextRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		want    int
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TextRequest{}},
			3,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetApps(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetApps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && len(got.Apps) != tt.want {
				t.Errorf("Middle.GetApps() = %v, want %v", len(got.Apps), tt.want)
			}
		})
	}
}

func TestMiddle_GetAvailableApp(t *testing.T) {
	type args struct {
		in0     context.Context
		payload *pb.TextRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		want    *pb.AppReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetAvailableApp(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetAvailableApp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Middle.GetAvailableApp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddle_StoreAppOAuth(t *testing.T) {
	type args struct {
		in0     context.Context
		payload *pb.AppRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.StoreAppOAuth(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.StoreAppOAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Middle.StoreAppOAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddle_GetCredential(t *testing.T) {
	type args struct {
		in0     context.Context
		payload *pb.CredentialRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		want    *pb.CredentialReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetCredential(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetCredential() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Middle.GetCredential() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddle_GetCredentials(t *testing.T) {
	type args struct {
		in0 context.Context
		in1 *pb.TextRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		want    *pb.CredentialsReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetCredentials(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Middle.GetCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddle_GetMaskingCredentials(t *testing.T) {
	type args struct {
		in0 context.Context
		in1 *pb.TextRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		want    *pb.MaskingReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetMaskingCredentials(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetMaskingCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Middle.GetMaskingCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddle_CreateCredential(t *testing.T) {
	type args struct {
		in0     context.Context
		payload *pb.KVsRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateCredential(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.CreateCredential() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Middle.CreateCredential() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddle_GetSettings(t *testing.T) {
	type args struct {
		in0 context.Context
		in1 *pb.TextRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		want    *pb.SettingsReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetSettings(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetSettings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Middle.GetSettings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddle_GetSetting(t *testing.T) {
	type args struct {
		in0     context.Context
		payload *pb.TextRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		want    *pb.SettingReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetSetting(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetSetting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Middle.GetSetting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddle_CreateSetting(t *testing.T) {
	type args struct {
		in0     context.Context
		payload *pb.KVRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateSetting(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.CreateSetting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Middle.CreateSetting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddle_GetStats(t *testing.T) {
	type args struct {
		ctx context.Context
		in1 *pb.TextRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		want    *pb.TextReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetStats(tt.args.ctx, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Middle.GetStats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authUUID(t *testing.T) {
	type args struct {
		rdb *redis.Client
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := authUUID(tt.args.rdb)
			if (err != nil) != tt.wantErr {
				t.Errorf("authUUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("authUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}
