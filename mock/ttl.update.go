// Code generated by MockGen. DO NOT EDIT.
// Source: ttl.update.go

// Package mock is a generated GoMock package.
package mock

import (
	aws "github.com/aws/aws-sdk-go/aws"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUpdateTTL is a mock of UpdateTTL interface
type MockUpdateTTL struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateTTLMockRecorder
}

// MockUpdateTTLMockRecorder is the mock recorder for MockUpdateTTL
type MockUpdateTTLMockRecorder struct {
	mock *MockUpdateTTL
}

// NewMockUpdateTTL creates a new mock instance
func NewMockUpdateTTL(ctrl *gomock.Controller) *MockUpdateTTL {
	mock := &MockUpdateTTL{ctrl: ctrl}
	mock.recorder = &MockUpdateTTLMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpdateTTL) EXPECT() *MockUpdateTTLMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockUpdateTTL) Run() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockUpdateTTLMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockUpdateTTL)(nil).Run))
}

// RunWithContext mocks base method
func (m *MockUpdateTTL) RunWithContext(ctx aws.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunWithContext", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunWithContext indicates an expected call of RunWithContext
func (mr *MockUpdateTTLMockRecorder) RunWithContext(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunWithContext", reflect.TypeOf((*MockUpdateTTL)(nil).RunWithContext), ctx)
}