// Code generated by MockGen. DO NOT EDIT.
// Source: tables.list.go

// Package mock is a generated GoMock package.
package mock

import (
	aws "github.com/aws/aws-sdk-go/aws"
	gomock "github.com/golang/mock/gomock"
	dynamo "github.com/guregu/dynamo"
	reflect "reflect"
)

// MockListTables is a mock of ListTables interface
type MockListTables struct {
	ctrl     *gomock.Controller
	recorder *MockListTablesMockRecorder
}

// MockListTablesMockRecorder is the mock recorder for MockListTables
type MockListTablesMockRecorder struct {
	mock *MockListTables
}

// NewMockListTables creates a new mock instance
func NewMockListTables(ctrl *gomock.Controller) *MockListTables {
	mock := &MockListTables{ctrl: ctrl}
	mock.recorder = &MockListTablesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockListTables) EXPECT() *MockListTablesMockRecorder {
	return m.recorder
}

// All mocks base method
func (m *MockListTables) All() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All
func (mr *MockListTablesMockRecorder) All() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockListTables)(nil).All))
}

// AllWithContext mocks base method
func (m *MockListTables) AllWithContext(ctx aws.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllWithContext", ctx)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllWithContext indicates an expected call of AllWithContext
func (mr *MockListTablesMockRecorder) AllWithContext(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllWithContext", reflect.TypeOf((*MockListTables)(nil).AllWithContext), ctx)
}

// Iter mocks base method
func (m *MockListTables) Iter() dynamo.Iter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iter")
	ret0, _ := ret[0].(dynamo.Iter)
	return ret0
}

// Iter indicates an expected call of Iter
func (mr *MockListTablesMockRecorder) Iter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iter", reflect.TypeOf((*MockListTables)(nil).Iter))
}
