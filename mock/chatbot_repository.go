// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/app/chatbot/repository/chatbot.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pb "github.com/tsundata/assistant/api/pb"
)

// MockChatbotRepository is a mock of ChatbotRepository interface.
type MockChatbotRepository struct {
	ctrl     *gomock.Controller
	recorder *MockChatbotRepositoryMockRecorder
}

// MockChatbotRepositoryMockRecorder is the mock recorder for MockChatbotRepository.
type MockChatbotRepositoryMockRecorder struct {
	mock *MockChatbotRepository
}

// NewMockChatbotRepository creates a new mock instance.
func NewMockChatbotRepository(ctrl *gomock.Controller) *MockChatbotRepository {
	mock := &MockChatbotRepository{ctrl: ctrl}
	mock.recorder = &MockChatbotRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatbotRepository) EXPECT() *MockChatbotRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockChatbotRepository) Create(ctx context.Context, message *pb.Bot) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, message)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockChatbotRepositoryMockRecorder) Create(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockChatbotRepository)(nil).Create), ctx, message)
}

// CreateGroup mocks base method.
func (m *MockChatbotRepository) CreateGroup(ctx context.Context, group *pb.Group) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroup", ctx, group)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroup indicates an expected call of CreateGroup.
func (mr *MockChatbotRepositoryMockRecorder) CreateGroup(ctx, group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroup", reflect.TypeOf((*MockChatbotRepository)(nil).CreateGroup), ctx, group)
}

// Delete mocks base method.
func (m *MockChatbotRepository) Delete(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockChatbotRepositoryMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockChatbotRepository)(nil).Delete), ctx, id)
}

// DeleteGroup mocks base method.
func (m *MockChatbotRepository) DeleteGroup(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroup", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGroup indicates an expected call of DeleteGroup.
func (mr *MockChatbotRepositoryMockRecorder) DeleteGroup(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroup", reflect.TypeOf((*MockChatbotRepository)(nil).DeleteGroup), ctx, id)
}

// GetByID mocks base method.
func (m *MockChatbotRepository) GetByID(ctx context.Context, id int64) (*pb.Bot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*pb.Bot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockChatbotRepositoryMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockChatbotRepository)(nil).GetByID), ctx, id)
}

// GetByIdentifier mocks base method.
func (m *MockChatbotRepository) GetByIdentifier(ctx context.Context, uuid string) (*pb.Bot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIdentifier", ctx, uuid)
	ret0, _ := ret[0].(*pb.Bot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIdentifier indicates an expected call of GetByIdentifier.
func (mr *MockChatbotRepositoryMockRecorder) GetByIdentifier(ctx, uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIdentifier", reflect.TypeOf((*MockChatbotRepository)(nil).GetByIdentifier), ctx, uuid)
}

// GetByUUID mocks base method.
func (m *MockChatbotRepository) GetByUUID(ctx context.Context, uuid string) (*pb.Bot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUUID", ctx, uuid)
	ret0, _ := ret[0].(*pb.Bot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUUID indicates an expected call of GetByUUID.
func (mr *MockChatbotRepositoryMockRecorder) GetByUUID(ctx, uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUUID", reflect.TypeOf((*MockChatbotRepository)(nil).GetByUUID), ctx, uuid)
}

// GetGroup mocks base method.
func (m *MockChatbotRepository) GetGroup(ctx context.Context, id int64) (*pb.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroup", ctx, id)
	ret0, _ := ret[0].(*pb.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroup indicates an expected call of GetGroup.
func (mr *MockChatbotRepositoryMockRecorder) GetGroup(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*MockChatbotRepository)(nil).GetGroup), ctx, id)
}

// GetGroupBySequence mocks base method.
func (m *MockChatbotRepository) GetGroupBySequence(ctx context.Context, userId, sequence int64) (*pb.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupBySequence", ctx, userId, sequence)
	ret0, _ := ret[0].(*pb.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupBySequence indicates an expected call of GetGroupBySequence.
func (mr *MockChatbotRepositoryMockRecorder) GetGroupBySequence(ctx, userId, sequence interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupBySequence", reflect.TypeOf((*MockChatbotRepository)(nil).GetGroupBySequence), ctx, userId, sequence)
}

// GetGroupByUUID mocks base method.
func (m *MockChatbotRepository) GetGroupByUUID(ctx context.Context, uuid string) (*pb.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupByUUID", ctx, uuid)
	ret0, _ := ret[0].(*pb.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupByUUID indicates an expected call of GetGroupByUUID.
func (mr *MockChatbotRepositoryMockRecorder) GetGroupByUUID(ctx, uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupByUUID", reflect.TypeOf((*MockChatbotRepository)(nil).GetGroupByUUID), ctx, uuid)
}

// List mocks base method.
func (m *MockChatbotRepository) List(ctx context.Context) ([]*pb.Bot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]*pb.Bot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockChatbotRepositoryMockRecorder) List(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockChatbotRepository)(nil).List), ctx)
}

// ListGroup mocks base method.
func (m *MockChatbotRepository) ListGroup(ctx context.Context, userId int64) ([]*pb.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroup", ctx, userId)
	ret0, _ := ret[0].([]*pb.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroup indicates an expected call of ListGroup.
func (mr *MockChatbotRepositoryMockRecorder) ListGroup(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroup", reflect.TypeOf((*MockChatbotRepository)(nil).ListGroup), ctx, userId)
}
