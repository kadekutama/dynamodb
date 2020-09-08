// Code generated by MockGen. DO NOT EDIT.
// Source: batch.get.go

// Package mock is a generated GoMock package.
package mock

import (
	aws "github.com/aws/aws-sdk-go/aws"
	gomock "github.com/golang/mock/gomock"
	dynamo "github.com/guregu/dynamo"
	dynamodb "github.com/kadekutama/dynamodb"
	reflect "reflect"
)

// MockBatchGet is a mock of BatchGet interface
type MockBatchGet struct {
	ctrl     *gomock.Controller
	recorder *MockBatchGetMockRecorder
}

// MockBatchGetMockRecorder is the mock recorder for MockBatchGet
type MockBatchGetMockRecorder struct {
	mock *MockBatchGet
}

// NewMockBatchGet creates a new mock instance
func NewMockBatchGet(ctrl *gomock.Controller) *MockBatchGet {
	mock := &MockBatchGet{ctrl: ctrl}
	mock.recorder = &MockBatchGetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBatchGet) EXPECT() *MockBatchGetMockRecorder {
	return m.recorder
}

// And mocks base method
func (m *MockBatchGet) And(keys ...dynamo.Keyed) dynamodb.BatchGet {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "And", varargs...)
	ret0, _ := ret[0].(dynamodb.BatchGet)
	return ret0
}

// And indicates an expected call of And
func (mr *MockBatchGetMockRecorder) And(keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "And", reflect.TypeOf((*MockBatchGet)(nil).And), keys...)
}

// Consistent mocks base method
func (m *MockBatchGet) Consistent(on bool) dynamodb.BatchGet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consistent", on)
	ret0, _ := ret[0].(dynamodb.BatchGet)
	return ret0
}

// Consistent indicates an expected call of Consistent
func (mr *MockBatchGetMockRecorder) Consistent(on interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consistent", reflect.TypeOf((*MockBatchGet)(nil).Consistent), on)
}

// ConsumedCapacity mocks base method
func (m *MockBatchGet) ConsumedCapacity(cc *dynamo.ConsumedCapacity) dynamodb.BatchGet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsumedCapacity", cc)
	ret0, _ := ret[0].(dynamodb.BatchGet)
	return ret0
}

// ConsumedCapacity indicates an expected call of ConsumedCapacity
func (mr *MockBatchGetMockRecorder) ConsumedCapacity(cc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumedCapacity", reflect.TypeOf((*MockBatchGet)(nil).ConsumedCapacity), cc)
}

// All mocks base method
func (m *MockBatchGet) All(out interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All", out)
	ret0, _ := ret[0].(error)
	return ret0
}

// All indicates an expected call of All
func (mr *MockBatchGetMockRecorder) All(out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockBatchGet)(nil).All), out)
}

// AllWithContext mocks base method
func (m *MockBatchGet) AllWithContext(ctx aws.Context, out interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllWithContext", ctx, out)
	ret0, _ := ret[0].(error)
	return ret0
}

// AllWithContext indicates an expected call of AllWithContext
func (mr *MockBatchGetMockRecorder) AllWithContext(ctx, out interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllWithContext", reflect.TypeOf((*MockBatchGet)(nil).AllWithContext), ctx, out)
}

// Iter mocks base method
func (m *MockBatchGet) Iter() dynamo.Iter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iter")
	ret0, _ := ret[0].(dynamo.Iter)
	return ret0
}

// Iter indicates an expected call of Iter
func (mr *MockBatchGetMockRecorder) Iter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iter", reflect.TypeOf((*MockBatchGet)(nil).Iter))
}
