// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/pb/nlp.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pb "github.com/tsundata/assistant/api/pb"
	grpc "google.golang.org/grpc"
)

// MockNLPSvcClient is a mock of NLPSvcClient interface.
type MockNLPSvcClient struct {
	ctrl     *gomock.Controller
	recorder *MockNLPSvcClientMockRecorder
}

// MockNLPSvcClientMockRecorder is the mock recorder for MockNLPSvcClient.
type MockNLPSvcClientMockRecorder struct {
	mock *MockNLPSvcClient
}

// NewMockNLPSvcClient creates a new mock instance.
func NewMockNLPSvcClient(ctrl *gomock.Controller) *MockNLPSvcClient {
	mock := &MockNLPSvcClient{ctrl: ctrl}
	mock.recorder = &MockNLPSvcClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNLPSvcClient) EXPECT() *MockNLPSvcClientMockRecorder {
	return m.recorder
}

// Classifier mocks base method.
func (m *MockNLPSvcClient) Classifier(ctx context.Context, in *pb.TextRequest, opts ...grpc.CallOption) (*pb.TextReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Classifier", varargs...)
	ret0, _ := ret[0].(*pb.TextReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Classifier indicates an expected call of Classifier.
func (mr *MockNLPSvcClientMockRecorder) Classifier(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Classifier", reflect.TypeOf((*MockNLPSvcClient)(nil).Classifier), varargs...)
}

// Pinyin mocks base method.
func (m *MockNLPSvcClient) Pinyin(ctx context.Context, in *pb.TextRequest, opts ...grpc.CallOption) (*pb.WordsReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Pinyin", varargs...)
	ret0, _ := ret[0].(*pb.WordsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pinyin indicates an expected call of Pinyin.
func (mr *MockNLPSvcClientMockRecorder) Pinyin(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pinyin", reflect.TypeOf((*MockNLPSvcClient)(nil).Pinyin), varargs...)
}

// Segmentation mocks base method.
func (m *MockNLPSvcClient) Segmentation(ctx context.Context, in *pb.TextRequest, opts ...grpc.CallOption) (*pb.WordsReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Segmentation", varargs...)
	ret0, _ := ret[0].(*pb.WordsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Segmentation indicates an expected call of Segmentation.
func (mr *MockNLPSvcClientMockRecorder) Segmentation(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Segmentation", reflect.TypeOf((*MockNLPSvcClient)(nil).Segmentation), varargs...)
}

// MockNLPSvcServer is a mock of NLPSvcServer interface.
type MockNLPSvcServer struct {
	ctrl     *gomock.Controller
	recorder *MockNLPSvcServerMockRecorder
}

// MockNLPSvcServerMockRecorder is the mock recorder for MockNLPSvcServer.
type MockNLPSvcServerMockRecorder struct {
	mock *MockNLPSvcServer
}

// NewMockNLPSvcServer creates a new mock instance.
func NewMockNLPSvcServer(ctrl *gomock.Controller) *MockNLPSvcServer {
	mock := &MockNLPSvcServer{ctrl: ctrl}
	mock.recorder = &MockNLPSvcServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNLPSvcServer) EXPECT() *MockNLPSvcServerMockRecorder {
	return m.recorder
}

// Classifier mocks base method.
func (m *MockNLPSvcServer) Classifier(arg0 context.Context, arg1 *pb.TextRequest) (*pb.TextReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Classifier", arg0, arg1)
	ret0, _ := ret[0].(*pb.TextReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Classifier indicates an expected call of Classifier.
func (mr *MockNLPSvcServerMockRecorder) Classifier(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Classifier", reflect.TypeOf((*MockNLPSvcServer)(nil).Classifier), arg0, arg1)
}

// Pinyin mocks base method.
func (m *MockNLPSvcServer) Pinyin(arg0 context.Context, arg1 *pb.TextRequest) (*pb.WordsReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pinyin", arg0, arg1)
	ret0, _ := ret[0].(*pb.WordsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pinyin indicates an expected call of Pinyin.
func (mr *MockNLPSvcServerMockRecorder) Pinyin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pinyin", reflect.TypeOf((*MockNLPSvcServer)(nil).Pinyin), arg0, arg1)
}

// Segmentation mocks base method.
func (m *MockNLPSvcServer) Segmentation(arg0 context.Context, arg1 *pb.TextRequest) (*pb.WordsReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Segmentation", arg0, arg1)
	ret0, _ := ret[0].(*pb.WordsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Segmentation indicates an expected call of Segmentation.
func (mr *MockNLPSvcServerMockRecorder) Segmentation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Segmentation", reflect.TypeOf((*MockNLPSvcServer)(nil).Segmentation), arg0, arg1)
}