// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/pb/chatbot_grpc.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pb "github.com/tsundata/assistant/api/pb"
	grpc "google.golang.org/grpc"
)

// MockChatbotSvcClient is a mock of ChatbotSvcClient interface.
type MockChatbotSvcClient struct {
	ctrl     *gomock.Controller
	recorder *MockChatbotSvcClientMockRecorder
}

// MockChatbotSvcClientMockRecorder is the mock recorder for MockChatbotSvcClient.
type MockChatbotSvcClientMockRecorder struct {
	mock *MockChatbotSvcClient
}

// NewMockChatbotSvcClient creates a new mock instance.
func NewMockChatbotSvcClient(ctrl *gomock.Controller) *MockChatbotSvcClient {
	mock := &MockChatbotSvcClient{ctrl: ctrl}
	mock.recorder = &MockChatbotSvcClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatbotSvcClient) EXPECT() *MockChatbotSvcClientMockRecorder {
	return m.recorder
}

// CreateGroup mocks base method.
func (m *MockChatbotSvcClient) CreateGroup(ctx context.Context, in *pb.GroupRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateGroup", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroup indicates an expected call of CreateGroup.
func (mr *MockChatbotSvcClientMockRecorder) CreateGroup(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroup", reflect.TypeOf((*MockChatbotSvcClient)(nil).CreateGroup), varargs...)
}

// GetBot mocks base method.
func (m *MockChatbotSvcClient) GetBot(ctx context.Context, in *pb.BotRequest, opts ...grpc.CallOption) (*pb.BotReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBot", varargs...)
	ret0, _ := ret[0].(*pb.BotReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBot indicates an expected call of GetBot.
func (mr *MockChatbotSvcClientMockRecorder) GetBot(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBot", reflect.TypeOf((*MockChatbotSvcClient)(nil).GetBot), varargs...)
}

// GetBots mocks base method.
func (m *MockChatbotSvcClient) GetBots(ctx context.Context, in *pb.BotRequest, opts ...grpc.CallOption) (*pb.BotsReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBots", varargs...)
	ret0, _ := ret[0].(*pb.BotsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBots indicates an expected call of GetBots.
func (mr *MockChatbotSvcClientMockRecorder) GetBots(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBots", reflect.TypeOf((*MockChatbotSvcClient)(nil).GetBots), varargs...)
}

// GetGroup mocks base method.
func (m *MockChatbotSvcClient) GetGroup(ctx context.Context, in *pb.GroupRequest, opts ...grpc.CallOption) (*pb.GroupReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroup", varargs...)
	ret0, _ := ret[0].(*pb.GroupReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroup indicates an expected call of GetGroup.
func (mr *MockChatbotSvcClientMockRecorder) GetGroup(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*MockChatbotSvcClient)(nil).GetGroup), varargs...)
}

// GetGroups mocks base method.
func (m *MockChatbotSvcClient) GetGroups(ctx context.Context, in *pb.GroupRequest, opts ...grpc.CallOption) (*pb.GroupsReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroups", varargs...)
	ret0, _ := ret[0].(*pb.GroupsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroups indicates an expected call of GetGroups.
func (mr *MockChatbotSvcClientMockRecorder) GetGroups(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroups", reflect.TypeOf((*MockChatbotSvcClient)(nil).GetGroups), varargs...)
}

// Handle mocks base method.
func (m *MockChatbotSvcClient) Handle(ctx context.Context, in *pb.ChatbotRequest, opts ...grpc.CallOption) (*pb.ChatbotReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Handle", varargs...)
	ret0, _ := ret[0].(*pb.ChatbotReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handle indicates an expected call of Handle.
func (mr *MockChatbotSvcClientMockRecorder) Handle(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockChatbotSvcClient)(nil).Handle), varargs...)
}

// UpdateBotSetting mocks base method.
func (m *MockChatbotSvcClient) UpdateBotSetting(ctx context.Context, in *pb.BotSettingRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateBotSetting", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBotSetting indicates an expected call of UpdateBotSetting.
func (mr *MockChatbotSvcClientMockRecorder) UpdateBotSetting(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBotSetting", reflect.TypeOf((*MockChatbotSvcClient)(nil).UpdateBotSetting), varargs...)
}

// MockChatbotSvcServer is a mock of ChatbotSvcServer interface.
type MockChatbotSvcServer struct {
	ctrl     *gomock.Controller
	recorder *MockChatbotSvcServerMockRecorder
}

// MockChatbotSvcServerMockRecorder is the mock recorder for MockChatbotSvcServer.
type MockChatbotSvcServerMockRecorder struct {
	mock *MockChatbotSvcServer
}

// NewMockChatbotSvcServer creates a new mock instance.
func NewMockChatbotSvcServer(ctrl *gomock.Controller) *MockChatbotSvcServer {
	mock := &MockChatbotSvcServer{ctrl: ctrl}
	mock.recorder = &MockChatbotSvcServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatbotSvcServer) EXPECT() *MockChatbotSvcServerMockRecorder {
	return m.recorder
}

// CreateGroup mocks base method.
func (m *MockChatbotSvcServer) CreateGroup(arg0 context.Context, arg1 *pb.GroupRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroup", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroup indicates an expected call of CreateGroup.
func (mr *MockChatbotSvcServerMockRecorder) CreateGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroup", reflect.TypeOf((*MockChatbotSvcServer)(nil).CreateGroup), arg0, arg1)
}

// GetBot mocks base method.
func (m *MockChatbotSvcServer) GetBot(arg0 context.Context, arg1 *pb.BotRequest) (*pb.BotReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBot", arg0, arg1)
	ret0, _ := ret[0].(*pb.BotReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBot indicates an expected call of GetBot.
func (mr *MockChatbotSvcServerMockRecorder) GetBot(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBot", reflect.TypeOf((*MockChatbotSvcServer)(nil).GetBot), arg0, arg1)
}

// GetBots mocks base method.
func (m *MockChatbotSvcServer) GetBots(arg0 context.Context, arg1 *pb.BotRequest) (*pb.BotsReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBots", arg0, arg1)
	ret0, _ := ret[0].(*pb.BotsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBots indicates an expected call of GetBots.
func (mr *MockChatbotSvcServerMockRecorder) GetBots(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBots", reflect.TypeOf((*MockChatbotSvcServer)(nil).GetBots), arg0, arg1)
}

// GetGroup mocks base method.
func (m *MockChatbotSvcServer) GetGroup(arg0 context.Context, arg1 *pb.GroupRequest) (*pb.GroupReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroup", arg0, arg1)
	ret0, _ := ret[0].(*pb.GroupReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroup indicates an expected call of GetGroup.
func (mr *MockChatbotSvcServerMockRecorder) GetGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*MockChatbotSvcServer)(nil).GetGroup), arg0, arg1)
}

// GetGroups mocks base method.
func (m *MockChatbotSvcServer) GetGroups(arg0 context.Context, arg1 *pb.GroupRequest) (*pb.GroupsReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroups", arg0, arg1)
	ret0, _ := ret[0].(*pb.GroupsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroups indicates an expected call of GetGroups.
func (mr *MockChatbotSvcServerMockRecorder) GetGroups(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroups", reflect.TypeOf((*MockChatbotSvcServer)(nil).GetGroups), arg0, arg1)
}

// Handle mocks base method.
func (m *MockChatbotSvcServer) Handle(arg0 context.Context, arg1 *pb.ChatbotRequest) (*pb.ChatbotReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", arg0, arg1)
	ret0, _ := ret[0].(*pb.ChatbotReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handle indicates an expected call of Handle.
func (mr *MockChatbotSvcServerMockRecorder) Handle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockChatbotSvcServer)(nil).Handle), arg0, arg1)
}

// UpdateBotSetting mocks base method.
func (m *MockChatbotSvcServer) UpdateBotSetting(arg0 context.Context, arg1 *pb.BotSettingRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBotSetting", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBotSetting indicates an expected call of UpdateBotSetting.
func (mr *MockChatbotSvcServerMockRecorder) UpdateBotSetting(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBotSetting", reflect.TypeOf((*MockChatbotSvcServer)(nil).UpdateBotSetting), arg0, arg1)
}

// MockUnsafeChatbotSvcServer is a mock of UnsafeChatbotSvcServer interface.
type MockUnsafeChatbotSvcServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeChatbotSvcServerMockRecorder
}

// MockUnsafeChatbotSvcServerMockRecorder is the mock recorder for MockUnsafeChatbotSvcServer.
type MockUnsafeChatbotSvcServerMockRecorder struct {
	mock *MockUnsafeChatbotSvcServer
}

// NewMockUnsafeChatbotSvcServer creates a new mock instance.
func NewMockUnsafeChatbotSvcServer(ctrl *gomock.Controller) *MockUnsafeChatbotSvcServer {
	mock := &MockUnsafeChatbotSvcServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeChatbotSvcServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeChatbotSvcServer) EXPECT() *MockUnsafeChatbotSvcServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedChatbotSvcServer mocks base method.
func (m *MockUnsafeChatbotSvcServer) mustEmbedUnimplementedChatbotSvcServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedChatbotSvcServer")
}

// mustEmbedUnimplementedChatbotSvcServer indicates an expected call of mustEmbedUnimplementedChatbotSvcServer.
func (mr *MockUnsafeChatbotSvcServerMockRecorder) mustEmbedUnimplementedChatbotSvcServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedChatbotSvcServer", reflect.TypeOf((*MockUnsafeChatbotSvcServer)(nil).mustEmbedUnimplementedChatbotSvcServer))
}
