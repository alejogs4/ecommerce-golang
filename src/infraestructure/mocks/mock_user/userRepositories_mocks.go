// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/user/domain/user/userRepository.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	reflect "reflect"

	user "github.com/alejogs4/hn-website/src/user/domain/user"
	gomock "github.com/golang/mock/gomock"
)

// MockQueries is a mock of Queries interface.
type MockQueries struct {
	ctrl     *gomock.Controller
	recorder *MockQueriesMockRecorder
}

// MockQueriesMockRecorder is the mock recorder for MockQueries.
type MockQueriesMockRecorder struct {
	mock *MockQueries
}

// NewMockQueries creates a new mock instance.
func NewMockQueries(ctrl *gomock.Controller) *MockQueries {
	mock := &MockQueries{ctrl: ctrl}
	mock.recorder = &MockQueriesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueries) EXPECT() *MockQueriesMockRecorder {
	return m.recorder
}

// GetByID mocks base method.
func (m *MockQueries) GetByID(id string) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockQueriesMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockQueries)(nil).GetByID), id)
}

// MockCommandsRepository is a mock of CommandsRepository interface.
type MockCommandsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCommandsRepositoryMockRecorder
}

// MockCommandsRepositoryMockRecorder is the mock recorder for MockCommandsRepository.
type MockCommandsRepositoryMockRecorder struct {
	mock *MockCommandsRepository
}

// NewMockCommandsRepository creates a new mock instance.
func NewMockCommandsRepository(ctrl *gomock.Controller) *MockCommandsRepository {
	mock := &MockCommandsRepository{ctrl: ctrl}
	mock.recorder = &MockCommandsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommandsRepository) EXPECT() *MockCommandsRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockCommandsRepository) CreateUser(user user.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockCommandsRepositoryMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockCommandsRepository)(nil).CreateUser), user)
}

// LoginUser mocks base method.
func (m *MockCommandsRepository) LoginUser(email, password string) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", email, password)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockCommandsRepositoryMockRecorder) LoginUser(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockCommandsRepository)(nil).LoginUser), email, password)
}

// VerifyEmail mocks base method.
func (m *MockCommandsRepository) VerifyEmail(userEmail string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyEmail", userEmail)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyEmail indicates an expected call of VerifyEmail.
func (mr *MockCommandsRepositoryMockRecorder) VerifyEmail(userEmail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyEmail", reflect.TypeOf((*MockCommandsRepository)(nil).VerifyEmail), userEmail)
}
