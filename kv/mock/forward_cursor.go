// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influxdata/influxdb/v2/kv (interfaces: ForwardCursor)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockForwardCursor is a mock of ForwardCursor interface
type MockForwardCursor struct {
	ctrl     *gomock.Controller
	recorder *MockForwardCursorMockRecorder
}

// MockForwardCursorMockRecorder is the mock recorder for MockForwardCursor
type MockForwardCursorMockRecorder struct {
	mock *MockForwardCursor
}

// NewMockForwardCursor creates a new mock instance
func NewMockForwardCursor(ctrl *gomock.Controller) *MockForwardCursor {
	mock := &MockForwardCursor{ctrl: ctrl}
	mock.recorder = &MockForwardCursorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockForwardCursor) EXPECT() *MockForwardCursorMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockForwardCursor) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockForwardCursorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockForwardCursor)(nil).Close))
}

// Err mocks base method
func (m *MockForwardCursor) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (mr *MockForwardCursorMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockForwardCursor)(nil).Err))
}

// Next mocks base method
func (m *MockForwardCursor) Next() ([]byte, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// Next indicates an expected call of Next
func (mr *MockForwardCursorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockForwardCursor)(nil).Next))
}
