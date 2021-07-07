package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/vendors"
	"github.com/tsundata/assistant/mock"
)

func TestUser_Authorization(t *testing.T) {
	rdb, err := vendors.CreateRedisClient(app.User)
	if err != nil {
		t.Fatal(err)
	}
	rdb.Set(context.Background(), AuthKey, "test", time.Hour)

	s := NewUser(rdb, nil)
	type args struct {
		ctx     context.Context
		payload *pb.TextRequest
	}
	tests := []struct {
		name    string
		s       *User
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TextRequest{Text: "test"}},
			&pb.StateReply{State: true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Authorization(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.Authorization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.Authorization() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetRole(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockUserRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().GetRole(gomock.Any()).Return(model.Role{Profession: "super"}, nil),
	)

	s := NewUser(nil, repo)

	type args struct {
		in0     context.Context
		payload *pb.RoleRequest
	}
	tests := []struct {
		name    string
		s       *User
		args    args
		want    *pb.RoleReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.RoleRequest{Id: 1}}, &pb.RoleReply{
			Role: &pb.Role{
				Profession: "super",
			},
		},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetRole(tt.args.in0, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.GetRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil || got.Role.Profession != tt.want.Role.Profession {
				t.Errorf("User.GetRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetRoleImage(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockUserRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().GetRole(gomock.Any()).Return(model.Role{
			ID:          1,
			Profession:  "super",
			Level:       60,
			Exp:         1592481,
			Strength:    120,
			Culture:     150,
			Environment: 160,
			Charisma:    180,
			Talent:      190,
			Intellect:   120,
		}, nil),
	)

	s := NewUser(nil, repo)

	type args struct {
		ctx     context.Context
		payload *pb.RoleRequest
	}
	tests := []struct {
		name    string
		s       *User
		args    args
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.RoleRequest{Id: 1}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.s.GetRoleImage(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.GetRoleImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUser_GetAuthToken(t *testing.T) {
	rdb, err := vendors.CreateRedisClient(app.User)
	if err != nil {
		t.Fatal(err)
	}
	rdb.Set(context.Background(), AuthKey, "test", time.Hour)

	s := NewUser(rdb, nil)

	type args struct {
		ctx context.Context
		in1 *pb.TextRequest
	}
	tests := []struct {
		name    string
		s       *User
		args    args
		want    *pb.TextReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.TextRequest{}},
			&pb.TextReply{Text: "test"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetAuthToken(tt.args.ctx, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.GetAuthToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.GetAuthToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_CreateUser(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockUserRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().Create(gomock.Any()).Return(int64(1), nil),
	)

	s := NewUser(nil, repo)

	type args struct {
		ctx     context.Context
		payload *pb.UserRequest
	}
	tests := []struct {
		name    string
		s       *User
		args    args
		want    *pb.UserReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.UserRequest{Name: "test"}},
			&pb.UserReply{Id: 1},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateUser(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetUser(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockUserRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().GetByID(gomock.Any()).Return(model.User{ID: 1, Name: "test"}, nil),
	)

	s := NewUser(nil, repo)

	type args struct {
		ctx     context.Context
		payload *pb.UserRequest
	}
	tests := []struct {
		name    string
		s       *User
		args    args
		want    *pb.UserReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.UserRequest{Name: "test"}},
			&pb.UserReply{Id: 1, Name: "test"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetUser(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetUsers(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockUserRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().List().Return([]model.User{{ID: 1, Name: "test"}}, nil),
	)

	s := NewUser(nil, repo)

	type args struct {
		ctx     context.Context
		payload *pb.UserRequest
	}
	tests := []struct {
		name    string
		s       *User
		args    args
		want    *pb.UsersReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.UserRequest{Name: "test"}},
			&pb.UsersReply{Users: []*pb.UserItem{{Id: 1, Name: "test"}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetUsers(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_UpdateUser(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock.NewMockUserRepository(ctl)
	gomock.InOrder(
		repo.EXPECT().Update(gomock.Any()).Return(nil),
	)

	s := NewUser(nil, repo)

	type args struct {
		ctx     context.Context
		payload *pb.UserRequest
	}
	tests := []struct {
		name    string
		s       *User
		args    args
		want    *pb.StateReply
		wantErr bool
	}{
		{
			"case1",
			s,
			args{context.Background(), &pb.UserRequest{Id: 1, Name: "update"}},
			&pb.StateReply{State: true},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.UpdateUser(tt.args.ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
