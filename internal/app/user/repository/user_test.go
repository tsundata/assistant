package repository

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"testing"
)

func TestUserRepository_GetRole(t *testing.T) {
	sto, err := CreateUserRepository(enum.User)
	if err != nil {
		t.Fatalf("create user Repository error, %+v", err)
	}
	type args struct {
		userId int
	}
	tests := []struct {
		name    string
		r       UserRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{userId: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetRole(context.Background(), tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserRepository_ChangeRoleExp(t *testing.T) {
	sto, err := CreateUserRepository(enum.User)
	if err != nil {
		t.Fatalf("create user Repository error, %+v", err)
	}
	type args struct {
		userID int64
		exp    int64
	}
	tests := []struct {
		name    string
		r       UserRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{userID: 1, exp: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.ChangeRoleExp(context.Background(), tt.args.userID, tt.args.exp); (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.ChangeRoleExp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_ChangeRoleAttr(t *testing.T) {
	sto, err := CreateUserRepository(enum.User)
	if err != nil {
		t.Fatalf("create user Repository error, %+v", err)
	}
	type args struct {
		userID int64
		attr   enum.RoleAttr
		val    int64
	}
	tests := []struct {
		name    string
		r       UserRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{userID: 1, attr: enum.StrengthAttr, val: 1}, false},
		{"case2", sto, args{userID: 1, attr: enum.CultureAttr, val: 1}, false},
		{"case3", sto, args{userID: 1, attr: enum.EnvironmentAttr, val: 1}, false},
		{"case4", sto, args{userID: 1, attr: enum.CharismaAttr, val: 1}, false},
		{"case5", sto, args{userID: 1, attr: enum.TalentAttr, val: 1}, false},
		{"case5", sto, args{userID: 1, attr: enum.IntellectAttr, val: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.ChangeRoleAttr(context.Background(), tt.args.userID, string(tt.args.attr), tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.ChangeRoleAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_List(t *testing.T) {
	sto, err := CreateUserRepository(enum.User)
	if err != nil {
		t.Fatalf("create user Repository error, %+v", err)
	}
	tests := []struct {
		name    string
		r       UserRepository
		wantErr bool
	}{
		{
			"case1",
			sto,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.List(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserRepository_Create(t *testing.T) {
	sto, err := CreateUserRepository(enum.User)
	if err != nil {
		t.Fatalf("create user Repository error, %+v", err)
	}
	type args struct {
		user *pb.User
	}
	tests := []struct {
		name    string
		r       UserRepository
		args    args
		wantErr bool
	}{
		{
			"case1",
			sto,
			args{&pb.User{Nickname: "create", Mobile: "", Remark: ""}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.Create(context.Background(), tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserRepository_GetByID(t *testing.T) {
	sto, err := CreateUserRepository(enum.User)
	if err != nil {
		t.Fatalf("create user Repository error, %+v", err)
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		r       UserRepository
		args    args
		wantErr bool
	}{
		{
			"case1",
			sto,
			args{id: 1},
			false,
		},
		{
			"case2",
			sto,
			args{id: 99999999},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetByID(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserRepository_GetByName(t *testing.T) {
	sto, err := CreateUserRepository(enum.User)
	if err != nil {
		t.Fatalf("create user Repository error, %+v", err)
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		r       UserRepository
		args    args
		wantErr bool
	}{
		{
			"case1",
			sto,
			args{name: "admin"},
			false,
		},
		{
			"case2",
			sto,
			args{name: "not_found"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetByName(context.Background(), tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserRepository_Update(t *testing.T) {
	sto, err := CreateUserRepository(enum.User)
	if err != nil {
		t.Fatalf("create user Repository error, %+v", err)
	}
	type args struct {
		in0 *pb.User
	}
	tests := []struct {
		name    string
		r       UserRepository
		args    args
		wantErr bool
	}{
		{
			"case1",
			sto,
			args{&pb.User{Id: 2, Nickname: "update"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Update(context.Background(), tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_CreateDevice(t *testing.T) {
	sto, err := CreateUserRepository(enum.User)
	if err != nil {
		t.Fatalf("create user Repository error, %+v", err)
	}
	type args struct {
		user *pb.Device
	}
	tests := []struct {
		name    string
		r       UserRepository
		args    args
		wantErr bool
	}{
		{
			"case1",
			sto,
			args{&pb.Device{UserId: 1, Name: "test"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.CreateDevice(context.Background(), tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.CreateDevice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserRepository_GetDevice(t *testing.T) {
	sto, err := CreateUserRepository(enum.User)
	if err != nil {
		t.Fatalf("create user Repository error, %+v", err)
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		r       UserRepository
		args    args
		wantErr bool
	}{
		{
			"case2",
			sto,
			args{id: 99999999},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetDevice(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetDevice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserRepository_ListDevice(t *testing.T) {
	sto, err := CreateUserRepository(enum.User)
	if err != nil {
		t.Fatalf("create user Repository error, %+v", err)
	}
	type args struct {
		userID int64
	}
	tests := []struct {
		name    string
		r       UserRepository
		args    args
		wantErr bool
	}{
		{
			"case1",
			sto,
			args{1},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.ListDevice(context.Background(), tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
