// Code generated by MockGen. DO NOT EDIT.
// Source: repository/auth.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/uzimaru0000/poker/model"
	reflect "reflect"
)

// MockAuthRepository is a mock of AuthRepository interface
type MockAuthRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRepositoryMockRecorder
}

// MockAuthRepositoryMockRecorder is the mock recorder for MockAuthRepository
type MockAuthRepositoryMockRecorder struct {
	mock *MockAuthRepository
}

// NewMockAuthRepository creates a new mock instance
func NewMockAuthRepository(ctrl *gomock.Controller) *MockAuthRepository {
	mock := &MockAuthRepository{ctrl: ctrl}
	mock.recorder = &MockAuthRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthRepository) EXPECT() *MockAuthRepositoryMockRecorder {
	return m.recorder
}

// GetHash mocks base method
func (m *MockAuthRepository) GetHash(arg0 *model.Auth) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHash", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHash indicates an expected call of GetHash
func (mr *MockAuthRepositoryMockRecorder) GetHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHash", reflect.TypeOf((*MockAuthRepository)(nil).GetHash), arg0)
}

// StoreHash mocks base method
func (m *MockAuthRepository) StoreHash(arg0 *model.Auth) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreHash", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreHash indicates an expected call of StoreHash
func (mr *MockAuthRepositoryMockRecorder) StoreHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreHash", reflect.TypeOf((*MockAuthRepository)(nil).StoreHash), arg0)
}
