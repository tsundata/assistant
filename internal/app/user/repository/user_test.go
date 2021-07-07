package repository

import (
	"testing"

	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/model"
)

func TestUserRepository_GetRole(t *testing.T) {
	sto, err := CreateUserRepository(app.User)
	if err != nil {
		t.Fatalf("create user Preposiory error, %+v", err)
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
			_, err := tt.r.GetRole(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlUserRepository.GetRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserRepository_ChangeRoleExp(t *testing.T) {
	sto, err := CreateUserRepository(app.User)
	if err != nil {
		t.Fatalf("create user Preposiory error, %+v", err)
	}
	type args struct {
		userID int
		exp    int
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
			if err := tt.r.ChangeRoleExp(tt.args.userID, tt.args.exp); (err != nil) != tt.wantErr {
				t.Errorf("MysqlUserRepository.ChangeRoleExp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_ChangeRoleAttr(t *testing.T) {
	sto, err := CreateUserRepository(app.User)
	if err != nil {
		t.Fatalf("create user Preposiory error, %+v", err)
	}
	type args struct {
		userID int
		attr   model.RoleAttr
		val    int
	}
	tests := []struct {
		name    string
		r       UserRepository
		args    args
		wantErr bool
	}{
		{"case1", sto, args{userID: 1, attr: model.StrengthAttr, val: 1}, false},
		{"case2", sto, args{userID: 1, attr: model.CultureAttr, val: 1}, false},
		{"case3", sto, args{userID: 1, attr: model.EnvironmentAttr, val: 1}, false},
		{"case4", sto, args{userID: 1, attr: model.CharismaAttr, val: 1}, false},
		{"case5", sto, args{userID: 1, attr: model.TalentAttr, val: 1}, false},
		{"case5", sto, args{userID: 1, attr: model.IntellectAttr, val: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.ChangeRoleAttr(tt.args.userID, string(tt.args.attr), tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("MysqlUserRepository.ChangeRoleAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMysqlUserRepository_List(t *testing.T) {
	sto, err := CreateUserRepository(app.User)
	if err != nil {
		t.Fatalf("create user Preposiory error, %+v", err)
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
			_, err := tt.r.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlUserRepository.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMysqlUserRepository_Create(t *testing.T) {
	sto, err := CreateUserRepository(app.User)
	if err != nil {
		t.Fatalf("create user Preposiory error, %+v", err)
	}
	type args struct {
		user model.User
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
			args{model.User{Name: "test", Mobile: "", Remark: ""}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.Create(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlUserRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMysqlUserRepository_GetByID(t *testing.T) {
	sto, err := CreateUserRepository(app.User)
	if err != nil {
		t.Fatalf("create user Preposiory error, %+v", err)
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
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("MysqlUserRepository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMysqlUserRepository_Update(t *testing.T) {
	sto, err := CreateUserRepository(app.User)
	if err != nil {
		t.Fatalf("create user Preposiory error, %+v", err)
	}
	type args struct {
		in0 model.User
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
			args{model.User{ID: 1, Name: "update"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Update(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("MysqlUserRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
