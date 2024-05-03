// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock_users is a generated GoMock package.
package mock_users

import (
	models "2024_1_TeaStealers/internal/models"
	context "context"
	io "io"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockUserUsecase is a mock of UserUsecase interface.
type MockUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUsecaseMockRecorder
}

// MockUserUsecaseMockRecorder is the mock recorder for MockUserUsecase.
type MockUserUsecaseMockRecorder struct {
	mock *MockUserUsecase
}

// NewMockUserUsecase creates a new mock instance.
func NewMockUserUsecase(ctrl *gomock.Controller) *MockUserUsecase {
	mock := &MockUserUsecase{ctrl: ctrl}
	mock.recorder = &MockUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUsecase) EXPECT() *MockUserUsecaseMockRecorder {
	return m.recorder
}

// DeleteUserPhoto mocks base method.
func (m *MockUserUsecase) DeleteUserPhoto(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserPhoto", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserPhoto indicates an expected call of DeleteUserPhoto.
func (mr *MockUserUsecaseMockRecorder) DeleteUserPhoto(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserPhoto", reflect.TypeOf((*MockUserUsecase)(nil).DeleteUserPhoto), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockUserUsecase) GetUser(arg0 context.Context, arg1 int64) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserUsecaseMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserUsecase)(nil).GetUser), arg0, arg1)
}

// UpdateUserInfo mocks base method.
func (m *MockUserUsecase) UpdateUserInfo(arg0 context.Context, arg1 int64, arg2 *models.UserUpdateData) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserInfo indicates an expected call of UpdateUserInfo.
func (mr *MockUserUsecaseMockRecorder) UpdateUserInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserInfo", reflect.TypeOf((*MockUserUsecase)(nil).UpdateUserInfo), arg0, arg1, arg2)
}

// UpdateUserPassword mocks base method.
func (m *MockUserUsecase) UpdateUserPassword(arg0 context.Context, arg1 *models.UserUpdatePassword) (string, time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPassword", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(time.Time)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateUserPassword indicates an expected call of UpdateUserPassword.
func (mr *MockUserUsecaseMockRecorder) UpdateUserPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPassword", reflect.TypeOf((*MockUserUsecase)(nil).UpdateUserPassword), arg0, arg1)
}

// UpdateUserPhoto mocks base method.
func (m *MockUserUsecase) UpdateUserPhoto(arg0 context.Context, arg1 io.Reader, arg2 string, arg3 int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPhoto", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserPhoto indicates an expected call of UpdateUserPhoto.
func (mr *MockUserUsecaseMockRecorder) UpdateUserPhoto(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPhoto", reflect.TypeOf((*MockUserUsecase)(nil).UpdateUserPhoto), arg0, arg1, arg2, arg3)
}

// MockUserRepo is a mock of UserRepo interface.
type MockUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoMockRecorder
}

// MockUserRepoMockRecorder is the mock recorder for MockUserRepo.
type MockUserRepoMockRecorder struct {
	mock *MockUserRepo
}

// NewMockUserRepo creates a new mock instance.
func NewMockUserRepo(ctrl *gomock.Controller) *MockUserRepo {
	mock := &MockUserRepo{ctrl: ctrl}
	mock.recorder = &MockUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepo) EXPECT() *MockUserRepoMockRecorder {
	return m.recorder
}

// CheckUserPassword mocks base method.
func (m *MockUserRepo) CheckUserPassword(arg0 context.Context, arg1 int64, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserPassword", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckUserPassword indicates an expected call of CheckUserPassword.
func (mr *MockUserRepoMockRecorder) CheckUserPassword(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserPassword", reflect.TypeOf((*MockUserRepo)(nil).CheckUserPassword), arg0, arg1, arg2)
}

// DeleteUserPhoto mocks base method.
func (m *MockUserRepo) DeleteUserPhoto(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserPhoto", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserPhoto indicates an expected call of DeleteUserPhoto.
func (mr *MockUserRepoMockRecorder) DeleteUserPhoto(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserPhoto", reflect.TypeOf((*MockUserRepo)(nil).DeleteUserPhoto), arg0, arg1)
}

// GetUserById mocks base method.
func (m *MockUserRepo) GetUserById(arg0 context.Context, arg1 int64) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserRepoMockRecorder) GetUserById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserRepo)(nil).GetUserById), arg0, arg1)
}

// UpdateUserInfo mocks base method.
func (m *MockUserRepo) UpdateUserInfo(arg0 context.Context, arg1 int64, arg2 *models.UserUpdateData) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserInfo indicates an expected call of UpdateUserInfo.
func (mr *MockUserRepoMockRecorder) UpdateUserInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserInfo", reflect.TypeOf((*MockUserRepo)(nil).UpdateUserInfo), arg0, arg1, arg2)
}

// UpdateUserPassword mocks base method.
func (m *MockUserRepo) UpdateUserPassword(arg0 context.Context, arg1 int64, arg2 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPassword", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserPassword indicates an expected call of UpdateUserPassword.
func (mr *MockUserRepoMockRecorder) UpdateUserPassword(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPassword", reflect.TypeOf((*MockUserRepo)(nil).UpdateUserPassword), arg0, arg1, arg2)
}

// UpdateUserPhoto mocks base method.
func (m *MockUserRepo) UpdateUserPhoto(arg0 context.Context, arg1 int64, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPhoto", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserPhoto indicates an expected call of UpdateUserPhoto.
func (mr *MockUserRepoMockRecorder) UpdateUserPhoto(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPhoto", reflect.TypeOf((*MockUserRepo)(nil).UpdateUserPhoto), arg0, arg1, arg2)
}