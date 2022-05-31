// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dapr/go-sdk/actor (interfaces: Server)

// Package actor is a generated GoMock package.
package mock

import (
	"context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	actor "github.com/liuxd6825/go-sdk/actor"
)

// MockServer is a mock of Server interface.
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
}

// MockServerMockRecorder is the mock recorder for MockServer.
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance.
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// ID mocks base method.
func (m *MockServer) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

func (mr *MockServerMockRecorder) Invoke(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invoke", reflect.TypeOf((*MockServer)(nil).Invoke), arg0, arg1)
}

func (m *MockServer) Invoke(ctx context.Context, input string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Invoke", ctx, input)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ID indicates an expected call of ID.
func (mr *MockServerMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockServer)(nil).ID))
}

// SaveState mocks base method.
func (m *MockServer) SaveState() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveState")
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveState indicates an expected call of SaveState.
func (mr *MockServerMockRecorder) SaveState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveState", reflect.TypeOf((*MockServer)(nil).SaveState))
}

// SetID mocks base method.
func (m *MockServer) SetID(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetID", arg0)
}

// SetID indicates an expected call of SetID.
func (mr *MockServerMockRecorder) SetID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetID", reflect.TypeOf((*MockServer)(nil).SetID), arg0)
}

// SetStateManager mocks base method.
func (m *MockServer) SetStateManager(arg0 actor.StateManager) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStateManager", arg0)
}

// SetStateManager indicates an expected call of SetStateManager.
func (mr *MockServerMockRecorder) SetStateManager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStateManager", reflect.TypeOf((*MockServer)(nil).SetStateManager), arg0)
}

// Type mocks base method.
func (m *MockServer) Type() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockServerMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockServer)(nil).Type))
}
