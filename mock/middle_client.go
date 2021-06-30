// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/pb/middle.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pb "github.com/tsundata/assistant/api/pb"
	grpc "google.golang.org/grpc"
)

// MockMiddleClient is a mock of MiddleClient interface.
type MockMiddleClient struct {
	ctrl     *gomock.Controller
	recorder *MockMiddleClientMockRecorder
}

// MockMiddleClientMockRecorder is the mock recorder for MockMiddleClient.
type MockMiddleClientMockRecorder struct {
	mock *MockMiddleClient
}

// NewMockMiddleClient creates a new mock instance.
func NewMockMiddleClient(ctrl *gomock.Controller) *MockMiddleClient {
	mock := &MockMiddleClient{ctrl: ctrl}
	mock.recorder = &MockMiddleClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMiddleClient) EXPECT() *MockMiddleClientMockRecorder {
	return m.recorder
}

// CreateCredential mocks base method.
func (m *MockMiddleClient) CreateCredential(ctx context.Context, in *pb.KVsRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateCredential", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCredential indicates an expected call of CreateCredential.
func (mr *MockMiddleClientMockRecorder) CreateCredential(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCredential", reflect.TypeOf((*MockMiddleClient)(nil).CreateCredential), varargs...)
}

// CreatePage mocks base method.
func (m *MockMiddleClient) CreatePage(ctx context.Context, in *pb.PageRequest, opts ...grpc.CallOption) (*pb.TextReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreatePage", varargs...)
	ret0, _ := ret[0].(*pb.TextReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePage indicates an expected call of CreatePage.
func (mr *MockMiddleClientMockRecorder) CreatePage(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePage", reflect.TypeOf((*MockMiddleClient)(nil).CreatePage), varargs...)
}

// CreateSetting mocks base method.
func (m *MockMiddleClient) CreateSetting(ctx context.Context, in *pb.KVRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateSetting", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSetting indicates an expected call of CreateSetting.
func (mr *MockMiddleClientMockRecorder) CreateSetting(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSetting", reflect.TypeOf((*MockMiddleClient)(nil).CreateSetting), varargs...)
}

// GetApps mocks base method.
func (m *MockMiddleClient) GetApps(ctx context.Context, in *pb.TextRequest, opts ...grpc.CallOption) (*pb.AppsReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetApps", varargs...)
	ret0, _ := ret[0].(*pb.AppsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApps indicates an expected call of GetApps.
func (mr *MockMiddleClientMockRecorder) GetApps(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApps", reflect.TypeOf((*MockMiddleClient)(nil).GetApps), varargs...)
}

// GetAvailableApp mocks base method.
func (m *MockMiddleClient) GetAvailableApp(ctx context.Context, in *pb.TextRequest, opts ...grpc.CallOption) (*pb.AppReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAvailableApp", varargs...)
	ret0, _ := ret[0].(*pb.AppReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableApp indicates an expected call of GetAvailableApp.
func (mr *MockMiddleClientMockRecorder) GetAvailableApp(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableApp", reflect.TypeOf((*MockMiddleClient)(nil).GetAvailableApp), varargs...)
}

// GetCredential mocks base method.
func (m *MockMiddleClient) GetCredential(ctx context.Context, in *pb.CredentialRequest, opts ...grpc.CallOption) (*pb.CredentialReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCredential", varargs...)
	ret0, _ := ret[0].(*pb.CredentialReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredential indicates an expected call of GetCredential.
func (mr *MockMiddleClientMockRecorder) GetCredential(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredential", reflect.TypeOf((*MockMiddleClient)(nil).GetCredential), varargs...)
}

// GetCredentials mocks base method.
func (m *MockMiddleClient) GetCredentials(ctx context.Context, in *pb.TextRequest, opts ...grpc.CallOption) (*pb.CredentialsReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCredentials", varargs...)
	ret0, _ := ret[0].(*pb.CredentialsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentials indicates an expected call of GetCredentials.
func (mr *MockMiddleClientMockRecorder) GetCredentials(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentials", reflect.TypeOf((*MockMiddleClient)(nil).GetCredentials), varargs...)
}

// GetMaskingCredentials mocks base method.
func (m *MockMiddleClient) GetMaskingCredentials(ctx context.Context, in *pb.TextRequest, opts ...grpc.CallOption) (*pb.MaskingReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMaskingCredentials", varargs...)
	ret0, _ := ret[0].(*pb.MaskingReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaskingCredentials indicates an expected call of GetMaskingCredentials.
func (mr *MockMiddleClientMockRecorder) GetMaskingCredentials(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaskingCredentials", reflect.TypeOf((*MockMiddleClient)(nil).GetMaskingCredentials), varargs...)
}

// GetMenu mocks base method.
func (m *MockMiddleClient) GetMenu(ctx context.Context, in *pb.TextRequest, opts ...grpc.CallOption) (*pb.TextReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMenu", varargs...)
	ret0, _ := ret[0].(*pb.TextReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMenu indicates an expected call of GetMenu.
func (mr *MockMiddleClientMockRecorder) GetMenu(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenu", reflect.TypeOf((*MockMiddleClient)(nil).GetMenu), varargs...)
}

// GetPage mocks base method.
func (m *MockMiddleClient) GetPage(ctx context.Context, in *pb.PageRequest, opts ...grpc.CallOption) (*pb.PageReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPage", varargs...)
	ret0, _ := ret[0].(*pb.PageReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPage indicates an expected call of GetPage.
func (mr *MockMiddleClientMockRecorder) GetPage(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPage", reflect.TypeOf((*MockMiddleClient)(nil).GetPage), varargs...)
}

// GetQrUrl mocks base method.
func (m *MockMiddleClient) GetQrUrl(ctx context.Context, in *pb.TextRequest, opts ...grpc.CallOption) (*pb.TextReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQrUrl", varargs...)
	ret0, _ := ret[0].(*pb.TextReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQrUrl indicates an expected call of GetQrUrl.
func (mr *MockMiddleClientMockRecorder) GetQrUrl(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQrUrl", reflect.TypeOf((*MockMiddleClient)(nil).GetQrUrl), varargs...)
}

// GetRoleImageUrl mocks base method.
func (m *MockMiddleClient) GetRoleImageUrl(ctx context.Context, in *pb.TextRequest, opts ...grpc.CallOption) (*pb.TextReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRoleImageUrl", varargs...)
	ret0, _ := ret[0].(*pb.TextReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleImageUrl indicates an expected call of GetRoleImageUrl.
func (mr *MockMiddleClientMockRecorder) GetRoleImageUrl(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleImageUrl", reflect.TypeOf((*MockMiddleClient)(nil).GetRoleImageUrl), varargs...)
}

// GetSetting mocks base method.
func (m *MockMiddleClient) GetSetting(ctx context.Context, in *pb.TextRequest, opts ...grpc.CallOption) (*pb.SettingReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSetting", varargs...)
	ret0, _ := ret[0].(*pb.SettingReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSetting indicates an expected call of GetSetting.
func (mr *MockMiddleClientMockRecorder) GetSetting(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSetting", reflect.TypeOf((*MockMiddleClient)(nil).GetSetting), varargs...)
}

// GetSettings mocks base method.
func (m *MockMiddleClient) GetSettings(ctx context.Context, in *pb.TextRequest, opts ...grpc.CallOption) (*pb.SettingsReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSettings", varargs...)
	ret0, _ := ret[0].(*pb.SettingsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSettings indicates an expected call of GetSettings.
func (mr *MockMiddleClientMockRecorder) GetSettings(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettings", reflect.TypeOf((*MockMiddleClient)(nil).GetSettings), varargs...)
}

// GetStats mocks base method.
func (m *MockMiddleClient) GetStats(ctx context.Context, in *pb.TextRequest, opts ...grpc.CallOption) (*pb.TextReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStats", varargs...)
	ret0, _ := ret[0].(*pb.TextReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStats indicates an expected call of GetStats.
func (mr *MockMiddleClientMockRecorder) GetStats(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockMiddleClient)(nil).GetStats), varargs...)
}

// StoreAppOAuth mocks base method.
func (m *MockMiddleClient) StoreAppOAuth(ctx context.Context, in *pb.AppRequest, opts ...grpc.CallOption) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StoreAppOAuth", varargs...)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreAppOAuth indicates an expected call of StoreAppOAuth.
func (mr *MockMiddleClientMockRecorder) StoreAppOAuth(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreAppOAuth", reflect.TypeOf((*MockMiddleClient)(nil).StoreAppOAuth), varargs...)
}

// MockMiddleServer is a mock of MiddleServer interface.
type MockMiddleServer struct {
	ctrl     *gomock.Controller
	recorder *MockMiddleServerMockRecorder
}

// MockMiddleServerMockRecorder is the mock recorder for MockMiddleServer.
type MockMiddleServerMockRecorder struct {
	mock *MockMiddleServer
}

// NewMockMiddleServer creates a new mock instance.
func NewMockMiddleServer(ctrl *gomock.Controller) *MockMiddleServer {
	mock := &MockMiddleServer{ctrl: ctrl}
	mock.recorder = &MockMiddleServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMiddleServer) EXPECT() *MockMiddleServerMockRecorder {
	return m.recorder
}

// CreateCredential mocks base method.
func (m *MockMiddleServer) CreateCredential(arg0 context.Context, arg1 *pb.KVsRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCredential", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCredential indicates an expected call of CreateCredential.
func (mr *MockMiddleServerMockRecorder) CreateCredential(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCredential", reflect.TypeOf((*MockMiddleServer)(nil).CreateCredential), arg0, arg1)
}

// CreatePage mocks base method.
func (m *MockMiddleServer) CreatePage(arg0 context.Context, arg1 *pb.PageRequest) (*pb.TextReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePage", arg0, arg1)
	ret0, _ := ret[0].(*pb.TextReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePage indicates an expected call of CreatePage.
func (mr *MockMiddleServerMockRecorder) CreatePage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePage", reflect.TypeOf((*MockMiddleServer)(nil).CreatePage), arg0, arg1)
}

// CreateSetting mocks base method.
func (m *MockMiddleServer) CreateSetting(arg0 context.Context, arg1 *pb.KVRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSetting", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSetting indicates an expected call of CreateSetting.
func (mr *MockMiddleServerMockRecorder) CreateSetting(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSetting", reflect.TypeOf((*MockMiddleServer)(nil).CreateSetting), arg0, arg1)
}

// GetApps mocks base method.
func (m *MockMiddleServer) GetApps(arg0 context.Context, arg1 *pb.TextRequest) (*pb.AppsReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApps", arg0, arg1)
	ret0, _ := ret[0].(*pb.AppsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApps indicates an expected call of GetApps.
func (mr *MockMiddleServerMockRecorder) GetApps(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApps", reflect.TypeOf((*MockMiddleServer)(nil).GetApps), arg0, arg1)
}

// GetAvailableApp mocks base method.
func (m *MockMiddleServer) GetAvailableApp(arg0 context.Context, arg1 *pb.TextRequest) (*pb.AppReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableApp", arg0, arg1)
	ret0, _ := ret[0].(*pb.AppReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableApp indicates an expected call of GetAvailableApp.
func (mr *MockMiddleServerMockRecorder) GetAvailableApp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableApp", reflect.TypeOf((*MockMiddleServer)(nil).GetAvailableApp), arg0, arg1)
}

// GetCredential mocks base method.
func (m *MockMiddleServer) GetCredential(arg0 context.Context, arg1 *pb.CredentialRequest) (*pb.CredentialReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredential", arg0, arg1)
	ret0, _ := ret[0].(*pb.CredentialReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredential indicates an expected call of GetCredential.
func (mr *MockMiddleServerMockRecorder) GetCredential(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredential", reflect.TypeOf((*MockMiddleServer)(nil).GetCredential), arg0, arg1)
}

// GetCredentials mocks base method.
func (m *MockMiddleServer) GetCredentials(arg0 context.Context, arg1 *pb.TextRequest) (*pb.CredentialsReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentials", arg0, arg1)
	ret0, _ := ret[0].(*pb.CredentialsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentials indicates an expected call of GetCredentials.
func (mr *MockMiddleServerMockRecorder) GetCredentials(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentials", reflect.TypeOf((*MockMiddleServer)(nil).GetCredentials), arg0, arg1)
}

// GetMaskingCredentials mocks base method.
func (m *MockMiddleServer) GetMaskingCredentials(arg0 context.Context, arg1 *pb.TextRequest) (*pb.MaskingReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaskingCredentials", arg0, arg1)
	ret0, _ := ret[0].(*pb.MaskingReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaskingCredentials indicates an expected call of GetMaskingCredentials.
func (mr *MockMiddleServerMockRecorder) GetMaskingCredentials(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaskingCredentials", reflect.TypeOf((*MockMiddleServer)(nil).GetMaskingCredentials), arg0, arg1)
}

// GetMenu mocks base method.
func (m *MockMiddleServer) GetMenu(arg0 context.Context, arg1 *pb.TextRequest) (*pb.TextReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMenu", arg0, arg1)
	ret0, _ := ret[0].(*pb.TextReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMenu indicates an expected call of GetMenu.
func (mr *MockMiddleServerMockRecorder) GetMenu(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenu", reflect.TypeOf((*MockMiddleServer)(nil).GetMenu), arg0, arg1)
}

// GetPage mocks base method.
func (m *MockMiddleServer) GetPage(arg0 context.Context, arg1 *pb.PageRequest) (*pb.PageReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPage", arg0, arg1)
	ret0, _ := ret[0].(*pb.PageReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPage indicates an expected call of GetPage.
func (mr *MockMiddleServerMockRecorder) GetPage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPage", reflect.TypeOf((*MockMiddleServer)(nil).GetPage), arg0, arg1)
}

// GetQrUrl mocks base method.
func (m *MockMiddleServer) GetQrUrl(arg0 context.Context, arg1 *pb.TextRequest) (*pb.TextReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQrUrl", arg0, arg1)
	ret0, _ := ret[0].(*pb.TextReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQrUrl indicates an expected call of GetQrUrl.
func (mr *MockMiddleServerMockRecorder) GetQrUrl(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQrUrl", reflect.TypeOf((*MockMiddleServer)(nil).GetQrUrl), arg0, arg1)
}

// GetRoleImageUrl mocks base method.
func (m *MockMiddleServer) GetRoleImageUrl(arg0 context.Context, arg1 *pb.TextRequest) (*pb.TextReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleImageUrl", arg0, arg1)
	ret0, _ := ret[0].(*pb.TextReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleImageUrl indicates an expected call of GetRoleImageUrl.
func (mr *MockMiddleServerMockRecorder) GetRoleImageUrl(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleImageUrl", reflect.TypeOf((*MockMiddleServer)(nil).GetRoleImageUrl), arg0, arg1)
}

// GetSetting mocks base method.
func (m *MockMiddleServer) GetSetting(arg0 context.Context, arg1 *pb.TextRequest) (*pb.SettingReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSetting", arg0, arg1)
	ret0, _ := ret[0].(*pb.SettingReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSetting indicates an expected call of GetSetting.
func (mr *MockMiddleServerMockRecorder) GetSetting(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSetting", reflect.TypeOf((*MockMiddleServer)(nil).GetSetting), arg0, arg1)
}

// GetSettings mocks base method.
func (m *MockMiddleServer) GetSettings(arg0 context.Context, arg1 *pb.TextRequest) (*pb.SettingsReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSettings", arg0, arg1)
	ret0, _ := ret[0].(*pb.SettingsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSettings indicates an expected call of GetSettings.
func (mr *MockMiddleServerMockRecorder) GetSettings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettings", reflect.TypeOf((*MockMiddleServer)(nil).GetSettings), arg0, arg1)
}

// GetStats mocks base method.
func (m *MockMiddleServer) GetStats(arg0 context.Context, arg1 *pb.TextRequest) (*pb.TextReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStats", arg0, arg1)
	ret0, _ := ret[0].(*pb.TextReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStats indicates an expected call of GetStats.
func (mr *MockMiddleServerMockRecorder) GetStats(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockMiddleServer)(nil).GetStats), arg0, arg1)
}

// StoreAppOAuth mocks base method.
func (m *MockMiddleServer) StoreAppOAuth(arg0 context.Context, arg1 *pb.AppRequest) (*pb.StateReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreAppOAuth", arg0, arg1)
	ret0, _ := ret[0].(*pb.StateReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreAppOAuth indicates an expected call of StoreAppOAuth.
func (mr *MockMiddleServerMockRecorder) StoreAppOAuth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreAppOAuth", reflect.TypeOf((*MockMiddleServer)(nil).StoreAppOAuth), arg0, arg1)
}
