// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/pb/workflow.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pb "github.com/tsundata/assistant/api/pb"
	grpc "google.golang.org/grpc"
)

// MockWorkflowSvcClient is a mock of WorkflowSvcClient interface.
type MockWorkflowSvcClient struct {
	ctrl     *gomock.Controller
	recorder *MockWorkflowSvcClientMockRecorder
}

// MockWorkflowSvcClientMockRecorder is the mock recorder for MockWorkflowSvcClient.
type MockWorkflowSvcClientMockRecorder struct {
	mock *MockWorkflowSvcClient
}

// NewMockWorkflowSvcClient creates a new mock instance.
func NewMockWorkflowSvcClient(ctrl *gomock.Controller) *MockWorkflowSvcClient {
	mock := &MockWorkflowSvcClient{ctrl: ctrl}
	mock.recorder = &MockWorkflowSvcClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkflowSvcClient) EXPECT() *MockWorkflowSvcClientMockRecorder {
	return m.recorder
}

// ActionDoc mocks base method.
func (m *MockWorkflowSvcClient) ActionDoc(ctx context.Context, in *pb.WorkflowRequest, opts ...grpc.CallOption) (*pb.WorkflowReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ActionDoc", varargs...)
	ret0, _ := ret[0].(*pb.WorkflowReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActionDoc indicates an expected call of ActionDoc.
func (mr *MockWorkflowSvcClientMockRecorder) ActionDoc(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionDoc", reflect.TypeOf((*MockWorkflowSvcClient)(nil).ActionDoc), varargs...)
}

// CreateTrigger mocks base method.
func (m *MockWorkflowSvcClient) CreateTrigger(ctx context.Context, in *pb.TriggerRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTrigger", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTrigger indicates an expected call of CreateTrigger.
func (mr *MockWorkflowSvcClientMockRecorder) CreateTrigger(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrigger", reflect.TypeOf((*MockWorkflowSvcClient)(nil).CreateTrigger), varargs...)
}

// CronTrigger mocks base method.
func (m *MockWorkflowSvcClient) CronTrigger(ctx context.Context, in *pb.TriggerRequest, opts ...grpc.CallOption) (*pb.WorkflowReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CronTrigger", varargs...)
	ret0, _ := ret[0].(*pb.WorkflowReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CronTrigger indicates an expected call of CronTrigger.
func (mr *MockWorkflowSvcClientMockRecorder) CronTrigger(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CronTrigger", reflect.TypeOf((*MockWorkflowSvcClient)(nil).CronTrigger), varargs...)
}

// DeleteTrigger mocks base method.
func (m *MockWorkflowSvcClient) DeleteTrigger(ctx context.Context, in *pb.TriggerRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteTrigger", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTrigger indicates an expected call of DeleteTrigger.
func (mr *MockWorkflowSvcClientMockRecorder) DeleteTrigger(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrigger", reflect.TypeOf((*MockWorkflowSvcClient)(nil).DeleteTrigger), varargs...)
}

// RunAction mocks base method.
func (m *MockWorkflowSvcClient) RunAction(ctx context.Context, in *pb.WorkflowRequest, opts ...grpc.CallOption) (*pb.WorkflowReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunAction", varargs...)
	ret0, _ := ret[0].(*pb.WorkflowReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunAction indicates an expected call of RunAction.
func (mr *MockWorkflowSvcClientMockRecorder) RunAction(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunAction", reflect.TypeOf((*MockWorkflowSvcClient)(nil).RunAction), varargs...)
}

// SyntaxCheck mocks base method.
func (m *MockWorkflowSvcClient) SyntaxCheck(ctx context.Context, in *pb.WorkflowRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SyntaxCheck", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyntaxCheck indicates an expected call of SyntaxCheck.
func (mr *MockWorkflowSvcClientMockRecorder) SyntaxCheck(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyntaxCheck", reflect.TypeOf((*MockWorkflowSvcClient)(nil).SyntaxCheck), varargs...)
}

// WebhookTrigger mocks base method.
func (m *MockWorkflowSvcClient) WebhookTrigger(ctx context.Context, in *pb.TriggerRequest, opts ...grpc.CallOption) (*pb.WorkflowReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WebhookTrigger", varargs...)
	ret0, _ := ret[0].(*pb.WorkflowReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WebhookTrigger indicates an expected call of WebhookTrigger.
func (mr *MockWorkflowSvcClientMockRecorder) WebhookTrigger(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WebhookTrigger", reflect.TypeOf((*MockWorkflowSvcClient)(nil).WebhookTrigger), varargs...)
}

// MockWorkflowSvcServer is a mock of WorkflowSvcServer interface.
type MockWorkflowSvcServer struct {
	ctrl     *gomock.Controller
	recorder *MockWorkflowSvcServerMockRecorder
}

// MockWorkflowSvcServerMockRecorder is the mock recorder for MockWorkflowSvcServer.
type MockWorkflowSvcServerMockRecorder struct {
	mock *MockWorkflowSvcServer
}

// NewMockWorkflowSvcServer creates a new mock instance.
func NewMockWorkflowSvcServer(ctrl *gomock.Controller) *MockWorkflowSvcServer {
	mock := &MockWorkflowSvcServer{ctrl: ctrl}
	mock.recorder = &MockWorkflowSvcServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkflowSvcServer) EXPECT() *MockWorkflowSvcServerMockRecorder {
	return m.recorder
}

// ActionDoc mocks base method.
func (m *MockWorkflowSvcServer) ActionDoc(arg0 context.Context, arg1 *pb.WorkflowRequest) (*pb.WorkflowReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionDoc", arg0, arg1)
	ret0, _ := ret[0].(*pb.WorkflowReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActionDoc indicates an expected call of ActionDoc.
func (mr *MockWorkflowSvcServerMockRecorder) ActionDoc(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionDoc", reflect.TypeOf((*MockWorkflowSvcServer)(nil).ActionDoc), arg0, arg1)
}

// CreateTrigger mocks base method.
func (m *MockWorkflowSvcServer) CreateTrigger(arg0 context.Context, arg1 *pb.TriggerRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTrigger", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTrigger indicates an expected call of CreateTrigger.
func (mr *MockWorkflowSvcServerMockRecorder) CreateTrigger(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrigger", reflect.TypeOf((*MockWorkflowSvcServer)(nil).CreateTrigger), arg0, arg1)
}

// CronTrigger mocks base method.
func (m *MockWorkflowSvcServer) CronTrigger(arg0 context.Context, arg1 *pb.TriggerRequest) (*pb.WorkflowReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CronTrigger", arg0, arg1)
	ret0, _ := ret[0].(*pb.WorkflowReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CronTrigger indicates an expected call of CronTrigger.
func (mr *MockWorkflowSvcServerMockRecorder) CronTrigger(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CronTrigger", reflect.TypeOf((*MockWorkflowSvcServer)(nil).CronTrigger), arg0, arg1)
}

// DeleteTrigger mocks base method.
func (m *MockWorkflowSvcServer) DeleteTrigger(arg0 context.Context, arg1 *pb.TriggerRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrigger", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTrigger indicates an expected call of DeleteTrigger.
func (mr *MockWorkflowSvcServerMockRecorder) DeleteTrigger(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrigger", reflect.TypeOf((*MockWorkflowSvcServer)(nil).DeleteTrigger), arg0, arg1)
}

// RunAction mocks base method.
func (m *MockWorkflowSvcServer) RunAction(arg0 context.Context, arg1 *pb.WorkflowRequest) (*pb.WorkflowReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunAction", arg0, arg1)
	ret0, _ := ret[0].(*pb.WorkflowReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunAction indicates an expected call of RunAction.
func (mr *MockWorkflowSvcServerMockRecorder) RunAction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunAction", reflect.TypeOf((*MockWorkflowSvcServer)(nil).RunAction), arg0, arg1)
}

// SyntaxCheck mocks base method.
func (m *MockWorkflowSvcServer) SyntaxCheck(arg0 context.Context, arg1 *pb.WorkflowRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyntaxCheck", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyntaxCheck indicates an expected call of SyntaxCheck.
func (mr *MockWorkflowSvcServerMockRecorder) SyntaxCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyntaxCheck", reflect.TypeOf((*MockWorkflowSvcServer)(nil).SyntaxCheck), arg0, arg1)
}

// WebhookTrigger mocks base method.
func (m *MockWorkflowSvcServer) WebhookTrigger(arg0 context.Context, arg1 *pb.TriggerRequest) (*pb.WorkflowReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WebhookTrigger", arg0, arg1)
	ret0, _ := ret[0].(*pb.WorkflowReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WebhookTrigger indicates an expected call of WebhookTrigger.
func (mr *MockWorkflowSvcServerMockRecorder) WebhookTrigger(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WebhookTrigger", reflect.TypeOf((*MockWorkflowSvcServer)(nil).WebhookTrigger), arg0, arg1)
}