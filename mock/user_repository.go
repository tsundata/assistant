// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/app/user/repository/user.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/tsundata/assistant/internal/pkg/model"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// ChangeRoleAttr mocks base method.
func (m *MockUserRepository) ChangeRoleAttr(userID int, attr string, val int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeRoleAttr", userID, attr, val)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeRoleAttr indicates an expected call of ChangeRoleAttr.
func (mr *MockUserRepositoryMockRecorder) ChangeRoleAttr(userID, attr, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeRoleAttr", reflect.TypeOf((*MockUserRepository)(nil).ChangeRoleAttr), userID, attr, val)
}

// ChangeRoleExp mocks base method.
func (m *MockUserRepository) ChangeRoleExp(userID, exp int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeRoleExp", userID, exp)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeRoleExp indicates an expected call of ChangeRoleExp.
func (mr *MockUserRepositoryMockRecorder) ChangeRoleExp(userID, exp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeRoleExp", reflect.TypeOf((*MockUserRepository)(nil).ChangeRoleExp), userID, exp)
}

// Create mocks base method.
func (m *MockUserRepository) Create(user model.User) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoryMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), user)
}

// GetByID mocks base method.
func (m *MockUserRepository) GetByID(id int64) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUserRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUserRepository)(nil).GetByID), id)
}

// GetByName mocks base method.
func (m *MockUserRepository) GetByName(name string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", name)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockUserRepositoryMockRecorder) GetByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockUserRepository)(nil).GetByName), name)
}

// GetRole mocks base method.
func (m *MockUserRepository) GetRole(userID int) (model.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", userID)
	ret0, _ := ret[0].(model.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRole indicates an expected call of GetRole.
func (mr *MockUserRepositoryMockRecorder) GetRole(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockUserRepository)(nil).GetRole), userID)
}

// List mocks base method.
func (m *MockUserRepository) List() ([]model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUserRepositoryMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserRepository)(nil).List))
}

// Update mocks base method.
func (m *MockUserRepository) Update(user model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserRepositoryMockRecorder) Update(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserRepository)(nil).Update), user)
}
