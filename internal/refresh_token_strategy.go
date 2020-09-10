// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite/handler/oauth2 (interfaces: RefreshTokenStrategy)

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory/fosite"
)

// MockRefreshTokenStrategy is a mock of RefreshTokenStrategy interface
type MockRefreshTokenStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockRefreshTokenStrategyMockRecorder
}

// MockRefreshTokenStrategyMockRecorder is the mock recorder for MockRefreshTokenStrategy
type MockRefreshTokenStrategyMockRecorder struct {
	mock *MockRefreshTokenStrategy
}

// NewMockRefreshTokenStrategy creates a new mock instance
func NewMockRefreshTokenStrategy(ctrl *gomock.Controller) *MockRefreshTokenStrategy {
	mock := &MockRefreshTokenStrategy{ctrl: ctrl}
	mock.recorder = &MockRefreshTokenStrategyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRefreshTokenStrategy) EXPECT() *MockRefreshTokenStrategyMockRecorder {
	return m.recorder
}

// GenerateRefreshToken mocks base method
func (m *MockRefreshTokenStrategy) GenerateRefreshToken(arg0 context.Context, arg1 fosite.Requester) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRefreshToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GenerateRefreshToken indicates an expected call of GenerateRefreshToken
func (mr *MockRefreshTokenStrategyMockRecorder) GenerateRefreshToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRefreshToken", reflect.TypeOf((*MockRefreshTokenStrategy)(nil).GenerateRefreshToken), arg0, arg1)
}

// RefreshTokenSignature mocks base method
func (m *MockRefreshTokenStrategy) RefreshTokenSignature(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshTokenSignature", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// RefreshTokenSignature indicates an expected call of RefreshTokenSignature
func (mr *MockRefreshTokenStrategyMockRecorder) RefreshTokenSignature(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshTokenSignature", reflect.TypeOf((*MockRefreshTokenStrategy)(nil).RefreshTokenSignature), arg0)
}

// ValidateRefreshToken mocks base method
func (m *MockRefreshTokenStrategy) ValidateRefreshToken(arg0 context.Context, arg1 fosite.Requester, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateRefreshToken", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateRefreshToken indicates an expected call of ValidateRefreshToken
func (mr *MockRefreshTokenStrategyMockRecorder) ValidateRefreshToken(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateRefreshToken", reflect.TypeOf((*MockRefreshTokenStrategy)(nil).ValidateRefreshToken), arg0, arg1, arg2)
}
