// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/app/middle/repository/middle.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/tsundata/assistant/internal/pkg/model"
)

// MockMiddleRepository is a mock of MiddleRepository interface.
type MockMiddleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMiddleRepositoryMockRecorder
}

// MockMiddleRepositoryMockRecorder is the mock recorder for MockMiddleRepository.
type MockMiddleRepositoryMockRecorder struct {
	mock *MockMiddleRepository
}

// NewMockMiddleRepository creates a new mock instance.
func NewMockMiddleRepository(ctrl *gomock.Controller) *MockMiddleRepository {
	mock := &MockMiddleRepository{ctrl: ctrl}
	mock.recorder = &MockMiddleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMiddleRepository) EXPECT() *MockMiddleRepositoryMockRecorder {
	return m.recorder
}

// CreateApp mocks base method.
func (m *MockMiddleRepository) CreateApp(app model.App) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApp", app)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateApp indicates an expected call of CreateApp.
func (mr *MockMiddleRepositoryMockRecorder) CreateApp(app interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApp", reflect.TypeOf((*MockMiddleRepository)(nil).CreateApp), app)
}

// CreateCredential mocks base method.
func (m *MockMiddleRepository) CreateCredential(credential model.Credential) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCredential", credential)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCredential indicates an expected call of CreateCredential.
func (mr *MockMiddleRepositoryMockRecorder) CreateCredential(credential interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCredential", reflect.TypeOf((*MockMiddleRepository)(nil).CreateCredential), credential)
}

// CreatePage mocks base method.
func (m *MockMiddleRepository) CreatePage(page model.Page) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePage", page)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePage indicates an expected call of CreatePage.
func (mr *MockMiddleRepositoryMockRecorder) CreatePage(page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePage", reflect.TypeOf((*MockMiddleRepository)(nil).CreatePage), page)
}

// GetAppByType mocks base method.
func (m *MockMiddleRepository) GetAppByType(t string) (model.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppByType", t)
	ret0, _ := ret[0].(model.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAppByType indicates an expected call of GetAppByType.
func (mr *MockMiddleRepositoryMockRecorder) GetAppByType(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppByType", reflect.TypeOf((*MockMiddleRepository)(nil).GetAppByType), t)
}

// GetAvailableAppByType mocks base method.
func (m *MockMiddleRepository) GetAvailableAppByType(t string) (model.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableAppByType", t)
	ret0, _ := ret[0].(model.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableAppByType indicates an expected call of GetAvailableAppByType.
func (mr *MockMiddleRepositoryMockRecorder) GetAvailableAppByType(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableAppByType", reflect.TypeOf((*MockMiddleRepository)(nil).GetAvailableAppByType), t)
}

// GetCredentialByName mocks base method.
func (m *MockMiddleRepository) GetCredentialByName(name string) (model.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentialByName", name)
	ret0, _ := ret[0].(model.Credential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentialByName indicates an expected call of GetCredentialByName.
func (mr *MockMiddleRepositoryMockRecorder) GetCredentialByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentialByName", reflect.TypeOf((*MockMiddleRepository)(nil).GetCredentialByName), name)
}

// GetCredentialByType mocks base method.
func (m *MockMiddleRepository) GetCredentialByType(t string) (model.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentialByType", t)
	ret0, _ := ret[0].(model.Credential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentialByType indicates an expected call of GetCredentialByType.
func (mr *MockMiddleRepositoryMockRecorder) GetCredentialByType(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentialByType", reflect.TypeOf((*MockMiddleRepository)(nil).GetCredentialByType), t)
}

// GetPageByUUID mocks base method.
func (m *MockMiddleRepository) GetPageByUUID(uuid string) (model.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPageByUUID", uuid)
	ret0, _ := ret[0].(model.Page)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPageByUUID indicates an expected call of GetPageByUUID.
func (mr *MockMiddleRepositoryMockRecorder) GetPageByUUID(uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPageByUUID", reflect.TypeOf((*MockMiddleRepository)(nil).GetPageByUUID), uuid)
}

// ListApps mocks base method.
func (m *MockMiddleRepository) ListApps() ([]model.App, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListApps")
	ret0, _ := ret[0].([]model.App)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApps indicates an expected call of ListApps.
func (mr *MockMiddleRepositoryMockRecorder) ListApps() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApps", reflect.TypeOf((*MockMiddleRepository)(nil).ListApps))
}

// ListCredentials mocks base method.
func (m *MockMiddleRepository) ListCredentials() ([]model.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCredentials")
	ret0, _ := ret[0].([]model.Credential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCredentials indicates an expected call of ListCredentials.
func (mr *MockMiddleRepositoryMockRecorder) ListCredentials() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCredentials", reflect.TypeOf((*MockMiddleRepository)(nil).ListCredentials))
}

// UpdateAppByID mocks base method.
func (m *MockMiddleRepository) UpdateAppByID(id int64, token, extra string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAppByID", id, token, extra)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAppByID indicates an expected call of UpdateAppByID.
func (mr *MockMiddleRepositoryMockRecorder) UpdateAppByID(id, token, extra interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAppByID", reflect.TypeOf((*MockMiddleRepository)(nil).UpdateAppByID), id, token, extra)
}
