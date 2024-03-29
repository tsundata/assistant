// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/pb/todo_grpc.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pb "github.com/tsundata/assistant/api/pb"
	grpc "google.golang.org/grpc"
)

// MockTodoSvcClient is a mock of TodoSvcClient interface.
type MockTodoSvcClient struct {
	ctrl     *gomock.Controller
	recorder *MockTodoSvcClientMockRecorder
}

// MockTodoSvcClientMockRecorder is the mock recorder for MockTodoSvcClient.
type MockTodoSvcClientMockRecorder struct {
	mock *MockTodoSvcClient
}

// NewMockTodoSvcClient creates a new mock instance.
func NewMockTodoSvcClient(ctrl *gomock.Controller) *MockTodoSvcClient {
	mock := &MockTodoSvcClient{ctrl: ctrl}
	mock.recorder = &MockTodoSvcClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoSvcClient) EXPECT() *MockTodoSvcClientMockRecorder {
	return m.recorder
}

// CompleteTodo mocks base method.
func (m *MockTodoSvcClient) CompleteTodo(ctx context.Context, in *pb.TodoRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CompleteTodo", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompleteTodo indicates an expected call of CompleteTodo.
func (mr *MockTodoSvcClientMockRecorder) CompleteTodo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteTodo", reflect.TypeOf((*MockTodoSvcClient)(nil).CompleteTodo), varargs...)
}

// CreateTodo mocks base method.
func (m *MockTodoSvcClient) CreateTodo(ctx context.Context, in *pb.TodoRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTodo", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTodo indicates an expected call of CreateTodo.
func (mr *MockTodoSvcClientMockRecorder) CreateTodo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodo", reflect.TypeOf((*MockTodoSvcClient)(nil).CreateTodo), varargs...)
}

// DeleteTodo mocks base method.
func (m *MockTodoSvcClient) DeleteTodo(ctx context.Context, in *pb.TodoRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteTodo", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTodo indicates an expected call of DeleteTodo.
func (mr *MockTodoSvcClientMockRecorder) DeleteTodo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTodo", reflect.TypeOf((*MockTodoSvcClient)(nil).DeleteTodo), varargs...)
}

// GetRemindTodos mocks base method.
func (m *MockTodoSvcClient) GetRemindTodos(ctx context.Context, in *pb.TodoRequest, opts ...grpc.CallOption) (*pb.TodosReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRemindTodos", varargs...)
	ret0, _ := ret[0].(*pb.TodosReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRemindTodos indicates an expected call of GetRemindTodos.
func (mr *MockTodoSvcClientMockRecorder) GetRemindTodos(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemindTodos", reflect.TypeOf((*MockTodoSvcClient)(nil).GetRemindTodos), varargs...)
}

// GetTodo mocks base method.
func (m *MockTodoSvcClient) GetTodo(ctx context.Context, in *pb.TodoRequest, opts ...grpc.CallOption) (*pb.TodoReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTodo", varargs...)
	ret0, _ := ret[0].(*pb.TodoReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodo indicates an expected call of GetTodo.
func (mr *MockTodoSvcClientMockRecorder) GetTodo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodo", reflect.TypeOf((*MockTodoSvcClient)(nil).GetTodo), varargs...)
}

// GetTodos mocks base method.
func (m *MockTodoSvcClient) GetTodos(ctx context.Context, in *pb.TodoRequest, opts ...grpc.CallOption) (*pb.TodosReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTodos", varargs...)
	ret0, _ := ret[0].(*pb.TodosReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodos indicates an expected call of GetTodos.
func (mr *MockTodoSvcClientMockRecorder) GetTodos(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodos", reflect.TypeOf((*MockTodoSvcClient)(nil).GetTodos), varargs...)
}

// UpdateTodo mocks base method.
func (m *MockTodoSvcClient) UpdateTodo(ctx context.Context, in *pb.TodoRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTodo", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTodo indicates an expected call of UpdateTodo.
func (mr *MockTodoSvcClientMockRecorder) UpdateTodo(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTodo", reflect.TypeOf((*MockTodoSvcClient)(nil).UpdateTodo), varargs...)
}

// MockTodoSvcServer is a mock of TodoSvcServer interface.
type MockTodoSvcServer struct {
	ctrl     *gomock.Controller
	recorder *MockTodoSvcServerMockRecorder
}

// MockTodoSvcServerMockRecorder is the mock recorder for MockTodoSvcServer.
type MockTodoSvcServerMockRecorder struct {
	mock *MockTodoSvcServer
}

// NewMockTodoSvcServer creates a new mock instance.
func NewMockTodoSvcServer(ctrl *gomock.Controller) *MockTodoSvcServer {
	mock := &MockTodoSvcServer{ctrl: ctrl}
	mock.recorder = &MockTodoSvcServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoSvcServer) EXPECT() *MockTodoSvcServerMockRecorder {
	return m.recorder
}

// CompleteTodo mocks base method.
func (m *MockTodoSvcServer) CompleteTodo(arg0 context.Context, arg1 *pb.TodoRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteTodo", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompleteTodo indicates an expected call of CompleteTodo.
func (mr *MockTodoSvcServerMockRecorder) CompleteTodo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteTodo", reflect.TypeOf((*MockTodoSvcServer)(nil).CompleteTodo), arg0, arg1)
}

// CreateTodo mocks base method.
func (m *MockTodoSvcServer) CreateTodo(arg0 context.Context, arg1 *pb.TodoRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTodo", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTodo indicates an expected call of CreateTodo.
func (mr *MockTodoSvcServerMockRecorder) CreateTodo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodo", reflect.TypeOf((*MockTodoSvcServer)(nil).CreateTodo), arg0, arg1)
}

// DeleteTodo mocks base method.
func (m *MockTodoSvcServer) DeleteTodo(arg0 context.Context, arg1 *pb.TodoRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTodo", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTodo indicates an expected call of DeleteTodo.
func (mr *MockTodoSvcServerMockRecorder) DeleteTodo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTodo", reflect.TypeOf((*MockTodoSvcServer)(nil).DeleteTodo), arg0, arg1)
}

// GetRemindTodos mocks base method.
func (m *MockTodoSvcServer) GetRemindTodos(arg0 context.Context, arg1 *pb.TodoRequest) (*pb.TodosReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemindTodos", arg0, arg1)
	ret0, _ := ret[0].(*pb.TodosReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRemindTodos indicates an expected call of GetRemindTodos.
func (mr *MockTodoSvcServerMockRecorder) GetRemindTodos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemindTodos", reflect.TypeOf((*MockTodoSvcServer)(nil).GetRemindTodos), arg0, arg1)
}

// GetTodo mocks base method.
func (m *MockTodoSvcServer) GetTodo(arg0 context.Context, arg1 *pb.TodoRequest) (*pb.TodoReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodo", arg0, arg1)
	ret0, _ := ret[0].(*pb.TodoReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodo indicates an expected call of GetTodo.
func (mr *MockTodoSvcServerMockRecorder) GetTodo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodo", reflect.TypeOf((*MockTodoSvcServer)(nil).GetTodo), arg0, arg1)
}

// GetTodos mocks base method.
func (m *MockTodoSvcServer) GetTodos(arg0 context.Context, arg1 *pb.TodoRequest) (*pb.TodosReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodos", arg0, arg1)
	ret0, _ := ret[0].(*pb.TodosReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodos indicates an expected call of GetTodos.
func (mr *MockTodoSvcServerMockRecorder) GetTodos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodos", reflect.TypeOf((*MockTodoSvcServer)(nil).GetTodos), arg0, arg1)
}

// UpdateTodo mocks base method.
func (m *MockTodoSvcServer) UpdateTodo(arg0 context.Context, arg1 *pb.TodoRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTodo", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTodo indicates an expected call of UpdateTodo.
func (mr *MockTodoSvcServerMockRecorder) UpdateTodo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTodo", reflect.TypeOf((*MockTodoSvcServer)(nil).UpdateTodo), arg0, arg1)
}

// MockUnsafeTodoSvcServer is a mock of UnsafeTodoSvcServer interface.
type MockUnsafeTodoSvcServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeTodoSvcServerMockRecorder
}

// MockUnsafeTodoSvcServerMockRecorder is the mock recorder for MockUnsafeTodoSvcServer.
type MockUnsafeTodoSvcServerMockRecorder struct {
	mock *MockUnsafeTodoSvcServer
}

// NewMockUnsafeTodoSvcServer creates a new mock instance.
func NewMockUnsafeTodoSvcServer(ctrl *gomock.Controller) *MockUnsafeTodoSvcServer {
	mock := &MockUnsafeTodoSvcServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeTodoSvcServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeTodoSvcServer) EXPECT() *MockUnsafeTodoSvcServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedTodoSvcServer mocks base method.
func (m *MockUnsafeTodoSvcServer) mustEmbedUnimplementedTodoSvcServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedTodoSvcServer")
}

// mustEmbedUnimplementedTodoSvcServer indicates an expected call of mustEmbedUnimplementedTodoSvcServer.
func (mr *MockUnsafeTodoSvcServerMockRecorder) mustEmbedUnimplementedTodoSvcServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedTodoSvcServer", reflect.TypeOf((*MockUnsafeTodoSvcServer)(nil).mustEmbedUnimplementedTodoSvcServer))
}
