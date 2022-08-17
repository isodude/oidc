// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/isodude/oidc/pkg/op (interfaces: DiscoverStorage)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	jose "gopkg.in/square/go-jose.v2"
)

// MockDiscoverStorage is a mock of DiscoverStorage interface.
type MockDiscoverStorage struct {
	ctrl     *gomock.Controller
	recorder *MockDiscoverStorageMockRecorder
}

// MockDiscoverStorageMockRecorder is the mock recorder for MockDiscoverStorage.
type MockDiscoverStorageMockRecorder struct {
	mock *MockDiscoverStorage
}

// NewMockDiscoverStorage creates a new mock instance.
func NewMockDiscoverStorage(ctrl *gomock.Controller) *MockDiscoverStorage {
	mock := &MockDiscoverStorage{ctrl: ctrl}
	mock.recorder = &MockDiscoverStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiscoverStorage) EXPECT() *MockDiscoverStorageMockRecorder {
	return m.recorder
}

// SignatureAlgorithms mocks base method.
func (m *MockDiscoverStorage) SignatureAlgorithms(arg0 context.Context) ([]jose.SignatureAlgorithm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignatureAlgorithms", arg0)
	ret0, _ := ret[0].([]jose.SignatureAlgorithm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignatureAlgorithms indicates an expected call of SignatureAlgorithms.
func (mr *MockDiscoverStorageMockRecorder) SignatureAlgorithms(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignatureAlgorithms", reflect.TypeOf((*MockDiscoverStorage)(nil).SignatureAlgorithms), arg0)
}
