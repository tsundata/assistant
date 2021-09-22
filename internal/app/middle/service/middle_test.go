package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/vendors"
	"github.com/tsundata/assistant/mock"
	"reflect"
	"testing"
	"time"
)

func TestMiddle_GetMenu(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	conf, err := config.CreateAppConfig(enum.Middle)
	if err != nil {
		t.Fatal(err)
	}

	user := mock.NewMockUserSvcClient(ctl)
	gomock.InOrder(
		user.EXPECT().
			GetAuthToken(gomock.Any(), gomock.Any()).
			Return(&pb.TextReply{Text: "test"}, nil),
	)

	s := NewMiddle(conf, nil, nil, user)

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
	conf, err := config.CreateAppConfig(enum.Middle)
	if err != nil {
		t.Fatal(err)
	}

	s := NewMiddle(conf, nil, nil, nil)

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

	conf, err := config.CreateAppConfig(enum.Middle)
	if err != nil {
		t.Fatal(err)
	}

	repo := mock.NewMockMiddleRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().CreatePage(gomock.Any(), gomock.Any()).Return(int64(1), nil),
	)

	s := NewMiddle(conf, nil, repo, nil)

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
			args{context.Background(), &pb.PageRequest{Page: &pb.Page{Type: "html", Title: "title", Content: "test"}}},
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
		repo.EXPECT().GetPageByUUID(gomock.Any(), gomock.Any()).Return(pb.Page{
			Id:      1,
			Uuid:    "test",
			Type:    "html",
			Title:   "test",
			Content: "test",
		}, nil),
	)

	s := NewMiddle(nil, nil, repo, nil)

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
			args{context.Background(), &pb.PageRequest{Page: &pb.Page{Uuid: "test"}}},
			&pb.PageReply{
				Page: &pb.Page{
					Uuid:    "test",
					Title:   "test",
					Content: "test",
					Type:    "html",
				},
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
		repo.EXPECT().ListApps(gomock.Any()).Return([]pb.App{{
			Id:    1,
			Type:  "github",
			Extra: `{"name": "github", "type":"github", "key": "test"}`,
		}}, nil),
	)

	s := NewMiddle(nil, nil, repo, nil)

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
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockMiddleRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().GetAvailableAppByType(gomock.Any(), gomock.Any()).Return(pb.App{
			Id:    1,
			Name:  "github",
			Type:  "github",
			Extra: `{"name": "github", "type":"github", "key": "test"}`,
			Token: "test",
		}, nil),
	)

	s := NewMiddle(nil, nil, repo, nil)

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
		{
			"case1",
			s,
			args{context.Background(), &pb.TextRequest{Text: "github"}},
			&pb.AppReply{Name: "github", Type: "github", Token: "test", Extra: []*pb.KV{}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetAvailableApp(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetAvailableApp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && (got.Token != tt.want.Token || got.Type != tt.want.Type) {
				t.Errorf("Middle.GetAvailableApp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddle_StoreAppOAuth(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockMiddleRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().GetAppByType(gomock.Any(), gomock.Any()).Return(pb.App{
			Id:    1,
			Name:  "github",
			Type:  "github",
			Extra: `{"name": "github", "type":"github", "key": "test"}`,
			Token: "test",
		}, nil),
		repo.EXPECT().UpdateAppByID(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil),
		repo.EXPECT().GetAppByType(gomock.Any(), gomock.Any()).Return(pb.App{
			Id: 0,
		}, nil),
		repo.EXPECT().CreateApp(gomock.Any(), gomock.Any()).Return(int64(1), nil),
	)

	s := NewMiddle(nil, nil, repo, nil)

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
		{
			"case1",
			s,
			args{context.Background(), &pb.AppRequest{App: &pb.App{Type: "github", Token: "test", Extra: "{}"}}},
			&pb.StateReply{State: true},
			false,
		},
		{
			"case2",
			s,
			args{context.Background(), &pb.AppRequest{App: &pb.App{Type: "github", Token: "test", Extra: "{}"}}},
			&pb.StateReply{State: true},
			false,
		},
		{
			"case2",
			s,
			args{context.Background(), &pb.AppRequest{App: &pb.App{Type: "github", Token: "", Extra: "{}"}}},
			&pb.StateReply{State: false},
			false,
		},
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
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockMiddleRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().GetCredentialByName(gomock.Any(), gomock.Any()).Return(pb.Credential{
			Id:      1,
			Name:    "github",
			Type:    "github",
			Content: `{"name": "github", "type":"github", "key": "test"}`,
		}, nil),
		repo.EXPECT().GetCredentialByType(gomock.Any(), gomock.Any()).Return(pb.Credential{
			Id:      1,
			Name:    "github",
			Type:    "github",
			Content: `{"name": "github", "type":"github", "key": "test"}`,
		}, nil),
	)

	s := NewMiddle(nil, nil, repo, nil)

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
		{
			"case1",
			s,
			args{context.Background(), &pb.CredentialRequest{Name: "github"}},
			&pb.CredentialReply{
				Name:    "github",
				Type:    "github",
				Content: []*pb.KV{},
			},
			false,
		},
		{
			"case2",
			s,
			args{context.Background(), &pb.CredentialRequest{Type: "github"}},
			&pb.CredentialReply{
				Name:    "github",
				Type:    "github",
				Content: []*pb.KV{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetCredential(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetCredential() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && (got.Name != tt.want.Name || got.Type != tt.want.Type) {
				t.Errorf("Middle.GetCredential() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddle_GetCredentials(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockMiddleRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().ListCredentials(gomock.Any()).Return([]pb.Credential{{
			Id:      1,
			Name:    "github",
			Type:    "github",
			Content: `{"name": "github", "type":"github", "key": "test"}`,
		}}, nil),
	)

	s := NewMiddle(nil, nil, repo, nil)

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
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetCredentials(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && len(got.Credentials) != tt.want {
				t.Errorf("Middle.GetCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddle_GetMaskingCredentials(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockMiddleRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().ListCredentials(gomock.Any()).Return([]pb.Credential{{
			Id:      1,
			Name:    "github",
			Type:    "github",
			Content: `{"name": "github", "type":"github", "key": "test"}`,
		}}, nil),
	)

	s := NewMiddle(nil, nil, repo, nil)

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
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetMaskingCredentials(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetMaskingCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && len(got.Items) != tt.want {
				t.Errorf("Middle.GetMaskingCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddle_CreateCredential(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockMiddleRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().CreateCredential(gomock.Any(), gomock.Any()).Return(int64(1), nil),
	)

	s := NewMiddle(nil, nil, repo, nil)

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
		{
			"case1",
			s,
			args{context.Background(), &pb.KVsRequest{Kvs: []*pb.KV{
				{Key: "name", Value: "github"},
				{Key: "type", Value: "github"},
				{Key: "key", Value: "123456"},
			}}},
			&pb.StateReply{State: true},
			false,
		},
		{
			"case2",
			s,
			args{context.Background(), &pb.KVsRequest{Kvs: []*pb.KV{
				{Key: "name", Value: ""},
				{Key: "type", Value: "github"},
				{Key: "key", Value: "123456"},
			}}},
			nil,
			true,
		},
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
	conf, err := config.CreateAppConfig(enum.Middle)
	if err != nil {
		t.Fatal(err)
	}

	s := NewMiddle(conf, nil, nil, nil)

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
			_, err := tt.s.GetSettings(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetSettings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddle_GetSetting(t *testing.T) {
	conf, err := config.CreateAppConfig(enum.Middle)
	if err != nil {
		t.Fatal(err)
	}

	s := NewMiddle(conf, nil, nil, nil)

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
			_, err := tt.s.GetSetting(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetSetting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddle_CreateSetting(t *testing.T) {
	conf, err := config.CreateAppConfig(enum.Middle)
	if err != nil {
		t.Fatal(err)
	}

	s := NewMiddle(conf, nil, nil, nil)

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
		{
			"case1",
			s,
			args{context.Background(), &pb.KVRequest{Key: "test", Value: "test"}},
			&pb.StateReply{State: true},
			false,
		},
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
	rdb, err := vendors.CreateRedisClient(enum.Middle)
	if err != nil {
		t.Fatal(err)
	}
	rdb.MSet(context.Background(), "stats:count:test", "test")
	rdb.MSet(context.Background(), "stats:month:0000", 0)

	s := NewMiddle(nil, rdb, nil, nil)

	type args struct {
		ctx context.Context
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
			_, err := tt.s.GetStats(tt.args.ctx, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMiddle_GetRoleImageUrl(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	user := mock.NewMockUserSvcClient(ctl)
	gomock.InOrder(
		user.EXPECT().
			GetAuthToken(gomock.Any(), gomock.Any()).
			Return(&pb.TextReply{Text: "test"}, nil),
	)

	conf, err := config.CreateAppConfig(enum.Middle)
	if err != nil {
		t.Fatal(err)
	}

	s := NewMiddle(conf, nil, nil, user)

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
		{
			"case1",
			s,
			args{context.Background(), &pb.TextRequest{}},
			&pb.TextReply{Text: fmt.Sprintf("%s/role/%s", conf.Web.Url, "test")},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetRoleImageUrl(tt.args.ctx, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetRoleImageUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Middle.GetRoleImageUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscribe_List(t *testing.T) {
	rdb, err := vendors.CreateRedisClient(enum.Middle)
	if err != nil {
		t.Fatal(err)
	}
	_, err = rdb.Del(context.Background(), RuleKey).Result()
	if err != nil {
		t.Fatal(err)
	}
	_, err = rdb.HMSet(context.Background(), RuleKey, "test1", "true").Result()
	if err != nil {
		t.Fatal(err)
	}
	_, err = rdb.HMSet(context.Background(), RuleKey, "test2", "true").Result()
	if err != nil {
		t.Fatal(err)
	}

	s := NewMiddle(nil, rdb, nil, nil)

	type args struct {
		ctx context.Context
		in1 *pb.SubscribeRequest
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
			args{context.Background(), &pb.SubscribeRequest{}},
			2,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ListSubscribe(tt.args.ctx, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subscribe.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && len(got.Subscribe) != tt.want {
				t.Errorf("Subscribe.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscribe_Register(t *testing.T) {
	rdb, err := vendors.CreateRedisClient(enum.Middle)
	if err != nil {
		t.Fatal(err)
	}

	s := NewMiddle(nil, rdb, nil, nil)

	type args struct {
		ctx     context.Context
		payload *pb.SubscribeRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.SubscribeRequest{Text: "test"}},
			&pb.StateReply{State: true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.RegisterSubscribe(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subscribe.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subscribe.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscribe_Open(t *testing.T) {
	rdb, err := vendors.CreateRedisClient(enum.Middle)
	if err != nil {
		t.Fatal(err)
	}

	s := NewMiddle(nil, rdb, nil, nil)

	type args struct {
		ctx     context.Context
		payload *pb.SubscribeRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.SubscribeRequest{Text: "test"}},
			&pb.StateReply{State: true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.OpenSubscribe(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subscribe.Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subscribe.Open() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscribe_Close(t *testing.T) {
	rdb, err := vendors.CreateRedisClient(enum.Middle)
	if err != nil {
		t.Fatal(err)
	}

	s := NewMiddle(nil, rdb, nil, nil)

	type args struct {
		ctx     context.Context
		payload *pb.SubscribeRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.SubscribeRequest{Text: "test"}},
			&pb.StateReply{State: true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CloseSubscribe(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subscribe.Close() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subscribe.Close() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscribe_Status(t *testing.T) {
	rdb, err := vendors.CreateRedisClient(enum.Middle)
	if err != nil {
		t.Fatal(err)
	}
	_, err = rdb.Del(context.Background(), RuleKey).Result()
	if err != nil {
		t.Fatal(err)
	}
	_, err = rdb.HMSet(context.Background(), RuleKey, "test1", "true").Result()
	if err != nil {
		t.Fatal(err)
	}
	_, err = rdb.HMSet(context.Background(), RuleKey, "test2", "false").Result()
	if err != nil {
		t.Fatal(err)
	}

	s := NewMiddle(nil, rdb, nil, nil)

	type args struct {
		ctx     context.Context
		payload *pb.SubscribeRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.SubscribeRequest{Text: "test1"}},
			&pb.StateReply{State: true},
			false,
		},
		{
			"case2",
			s,
			args{context.Background(), &pb.SubscribeRequest{Text: "test2"}},
			&pb.StateReply{State: false},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetSubscribeStatus(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subscribe.Status() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subscribe.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddle_GetTags(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockMiddleRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().ListTags(gomock.Any()).Return([]pb.Tag{{
			Id:        1,
			Name:      "test1",
			CreatedAt: time.Now().Unix(),
		}}, nil),
	)

	s := NewMiddle(nil, nil, repo, nil)

	type args struct {
		in0 context.Context
		in1 *pb.TagRequest
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
			args{context.Background(), &pb.TagRequest{Tag: &pb.Tag{Name: "test1"}}},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetTags(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && len(got.Tags) != tt.want {
				t.Errorf("Middle.GetTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddle_GetOrCreateTag(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockMiddleRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().GetOrCreateTag(gomock.Any(), gomock.Any()).Return(pb.Tag{
			Id:        1,
			Name:      "test1",
			CreatedAt: time.Now().Unix(),
		}, nil),
	)

	s := NewMiddle(nil, nil, repo, nil)

	type args struct {
		in0 context.Context
		in1 *pb.TagRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		want    pb.Tag
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TagRequest{Tag: &pb.Tag{Name: "test1"}}},
			pb.Tag{Id: 1, Name: "test1"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetOrCreateTag(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Middle.GetOrCreateTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.Tag.Id != tt.want.Id && got.Tag.Name != tt.want.Name {
				t.Errorf("Middle.GetOrCreateTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetChartData(t *testing.T) {
	rdb, err := vendors.CreateRedisClient(enum.Middle)
	if err != nil {
		t.Fatal(err)
	}

	s := NewMiddle(nil, rdb, nil, nil)

	type args struct {
		ctx     context.Context
		payload *pb.ChartDataRequest
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
			args{context.Background(), &pb.ChartDataRequest{ChartData: &pb.ChartData{
				Title:    "chart1",
				SubTitle: "sub_title",
				XAxis:    []string{"01", "02", "03"},
				Series:   []float64{2, 3, 4},
			}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.s.SetChartData(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subscribe.Close() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetChartData(t *testing.T) {
	uuid := "c30b82f7-430a-4c2e-bccd-30e920fc8136"
	rdb, err := vendors.CreateRedisClient(enum.Middle)
	if err != nil {
		t.Fatal(err)
	}

	data := pb.ChartData{
		Uuid:     uuid,
		Title:    "chart1",
		SubTitle: "sub_title",
		XAxis:    []string{"01", "02", "03"},
		Series:   []float64{2, 3, 4},
	}
	d, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	_, err = rdb.Set(context.Background(), fmt.Sprintf("middle:chart:%s", uuid), util.ByteToString(d), time.Hour).Result()
	if err != nil {
		t.Fatal(err)
	}

	s := NewMiddle(nil, rdb, nil, nil)

	type args struct {
		ctx     context.Context
		payload *pb.ChartDataRequest
	}
	tests := []struct {
		name    string
		s       *Middle
		args    args
		want    *pb.ChartDataReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.ChartDataRequest{ChartData: &pb.ChartData{Uuid: uuid}}},
			&pb.ChartDataReply{ChartData: &data},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetChartData(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subscribe.Status() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subscribe.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListCron(t *testing.T) {
	rdb, err := vendors.CreateRedisClient(enum.Middle)
	if err != nil {
		t.Fatal(err)
	}
	_, err = rdb.Del(context.Background(), CronKey).Result()
	if err != nil {
		t.Fatal(err)
	}
	_, err = rdb.HMSet(context.Background(), CronKey, "test1", "true").Result()
	if err != nil {
		t.Fatal(err)
	}
	_, err = rdb.HMSet(context.Background(), CronKey, "test2", "true").Result()
	if err != nil {
		t.Fatal(err)
	}

	s := NewMiddle(nil, rdb, nil, nil)

	type args struct {
		ctx context.Context
		in1 *pb.CronRequest
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
			args{context.Background(), &pb.CronRequest{}},
			2,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ListCron(tt.args.ctx, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cron.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && len(got.Cron) != tt.want {
				t.Errorf("Cron.List() = %v, want %v", got, tt.want)
			}
		})
	}
}
