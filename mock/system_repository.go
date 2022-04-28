// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/app/chatbot/bot/system/repository/system.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pb "github.com/tsundata/assistant/api/pb"
)

// MockSystemRepository is a mock of SystemRepository interface.
type MockSystemRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSystemRepositoryMockRecorder
}

// MockSystemRepositoryMockRecorder is the mock recorder for MockSystemRepository.
type MockSystemRepositoryMockRecorder struct {
	mock *MockSystemRepository
}

// NewMockSystemRepository creates a new mock instance.
func NewMockSystemRepository(ctrl *gomock.Controller) *MockSystemRepository {
	mock := &MockSystemRepository{ctrl: ctrl}
	mock.recorder = &MockSystemRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSystemRepository) EXPECT() *MockSystemRepositoryMockRecorder {
	return m.recorder
}

// CreateCounter mocks base method.
func (m *MockSystemRepository) CreateCounter(ctx context.Context, counter *pb.Counter) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCounter", ctx, counter)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCounter indicates an expected call of CreateCounter.
func (mr *MockSystemRepositoryMockRecorder) CreateCounter(ctx, counter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCounter", reflect.TypeOf((*MockSystemRepository)(nil).CreateCounter), ctx, counter)
}

// DecreaseCounter mocks base method.
func (m *MockSystemRepository) DecreaseCounter(ctx context.Context, id, amount int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecreaseCounter", ctx, id, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// DecreaseCounter indicates an expected call of DecreaseCounter.
func (mr *MockSystemRepositoryMockRecorder) DecreaseCounter(ctx, id, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecreaseCounter", reflect.TypeOf((*MockSystemRepository)(nil).DecreaseCounter), ctx, id, amount)
}

// GetCounter mocks base method.
func (m *MockSystemRepository) GetCounter(ctx context.Context, id int64) (pb.Counter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCounter", ctx, id)
	ret0, _ := ret[0].(pb.Counter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCounter indicates an expected call of GetCounter.
func (mr *MockSystemRepositoryMockRecorder) GetCounter(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCounter", reflect.TypeOf((*MockSystemRepository)(nil).GetCounter), ctx, id)
}

// GetCounterByFlag mocks base method.
func (m *MockSystemRepository) GetCounterByFlag(ctx context.Context, userId int64, flag string) (pb.Counter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCounterByFlag", ctx, userId, flag)
	ret0, _ := ret[0].(pb.Counter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCounterByFlag indicates an expected call of GetCounterByFlag.
func (mr *MockSystemRepositoryMockRecorder) GetCounterByFlag(ctx, userId, flag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCounterByFlag", reflect.TypeOf((*MockSystemRepository)(nil).GetCounterByFlag), ctx, userId, flag)
}

// IncreaseCounter mocks base method.
func (m *MockSystemRepository) IncreaseCounter(ctx context.Context, id, amount int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncreaseCounter", ctx, id, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncreaseCounter indicates an expected call of IncreaseCounter.
func (mr *MockSystemRepositoryMockRecorder) IncreaseCounter(ctx, id, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseCounter", reflect.TypeOf((*MockSystemRepository)(nil).IncreaseCounter), ctx, id, amount)
}

// ListCounter mocks base method.
func (m *MockSystemRepository) ListCounter(ctx context.Context, userId int64) ([]*pb.Counter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCounter", ctx, userId)
	ret0, _ := ret[0].([]*pb.Counter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCounter indicates an expected call of ListCounter.
func (mr *MockSystemRepositoryMockRecorder) ListCounter(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCounter", reflect.TypeOf((*MockSystemRepository)(nil).ListCounter), ctx, userId)
}