// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/app/message/repository/message.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/tsundata/assistant/internal/pkg/model"
)

// MockMessageRepository is a mock of MessageRepository interface.
type MockMessageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMessageRepositoryMockRecorder
}

// MockMessageRepositoryMockRecorder is the mock recorder for MockMessageRepository.
type MockMessageRepositoryMockRecorder struct {
	mock *MockMessageRepository
}

// NewMockMessageRepository creates a new mock instance.
func NewMockMessageRepository(ctrl *gomock.Controller) *MockMessageRepository {
	mock := &MockMessageRepository{ctrl: ctrl}
	mock.recorder = &MockMessageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageRepository) EXPECT() *MockMessageRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMessageRepository) Create(message model.Message) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", message)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMessageRepositoryMockRecorder) Create(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMessageRepository)(nil).Create), message)
}

// Delete mocks base method.
func (m *MockMessageRepository) Delete(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMessageRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMessageRepository)(nil).Delete), id)
}

// GetByID mocks base method.
func (m *MockMessageRepository) GetByID(id int64) (model.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockMessageRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockMessageRepository)(nil).GetByID), id)
}

// GetByUUID mocks base method.
func (m *MockMessageRepository) GetByUUID(uuid string) (model.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUUID", uuid)
	ret0, _ := ret[0].(model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUUID indicates an expected call of GetByUUID.
func (mr *MockMessageRepositoryMockRecorder) GetByUUID(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUUID", reflect.TypeOf((*MockMessageRepository)(nil).GetByUUID), uuid)
}

// List mocks base method.
func (m *MockMessageRepository) List() ([]model.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockMessageRepositoryMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMessageRepository)(nil).List))
}

// ListByType mocks base method.
func (m *MockMessageRepository) ListByType(t string) ([]model.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByType", t)
	ret0, _ := ret[0].([]model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByType indicates an expected call of ListByType.
func (mr *MockMessageRepositoryMockRecorder) ListByType(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByType", reflect.TypeOf((*MockMessageRepository)(nil).ListByType), t)
}
