// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package images_mock is a generated GoMock package.
package images_mock

import (
	models "2024_1_TeaStealers/internal/models"
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockImageUsecase is a mock of ImageUsecase interface.
type MockImageUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockImageUsecaseMockRecorder
}

// MockImageUsecaseMockRecorder is the mock recorder for MockImageUsecase.
type MockImageUsecaseMockRecorder struct {
	mock *MockImageUsecase
}

// NewMockImageUsecase creates a new mock instance.
func NewMockImageUsecase(ctrl *gomock.Controller) *MockImageUsecase {
	mock := &MockImageUsecase{ctrl: ctrl}
	mock.recorder = &MockImageUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageUsecase) EXPECT() *MockImageUsecaseMockRecorder {
	return m.recorder
}

// DeleteImage mocks base method.
func (m *MockImageUsecase) DeleteImage(arg0 context.Context, arg1 int64) ([]*models.ImageResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteImage", arg0, arg1)
	ret0, _ := ret[0].([]*models.ImageResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteImage indicates an expected call of DeleteImage.
func (mr *MockImageUsecaseMockRecorder) DeleteImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteImage", reflect.TypeOf((*MockImageUsecase)(nil).DeleteImage), arg0, arg1)
}

// GetAdvertImages mocks base method.
func (m *MockImageUsecase) GetAdvertImages(arg0 context.Context, arg1 int64) ([]*models.ImageResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdvertImages", arg0, arg1)
	ret0, _ := ret[0].([]*models.ImageResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdvertImages indicates an expected call of GetAdvertImages.
func (mr *MockImageUsecaseMockRecorder) GetAdvertImages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdvertImages", reflect.TypeOf((*MockImageUsecase)(nil).GetAdvertImages), arg0, arg1)
}

// UploadImage mocks base method.
func (m *MockImageUsecase) UploadImage(arg0 context.Context, arg1 io.Reader, arg2 string, arg3 int64) (*models.ImageResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadImage", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*models.ImageResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadImage indicates an expected call of UploadImage.
func (mr *MockImageUsecaseMockRecorder) UploadImage(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadImage", reflect.TypeOf((*MockImageUsecase)(nil).UploadImage), arg0, arg1, arg2, arg3)
}

// MockImageRepo is a mock of ImageRepo interface.
type MockImageRepo struct {
	ctrl     *gomock.Controller
	recorder *MockImageRepoMockRecorder
}

// MockImageRepoMockRecorder is the mock recorder for MockImageRepo.
type MockImageRepoMockRecorder struct {
	mock *MockImageRepo
}

// NewMockImageRepo creates a new mock instance.
func NewMockImageRepo(ctrl *gomock.Controller) *MockImageRepo {
	mock := &MockImageRepo{ctrl: ctrl}
	mock.recorder = &MockImageRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageRepo) EXPECT() *MockImageRepoMockRecorder {
	return m.recorder
}

// DeleteImage mocks base method.
func (m *MockImageRepo) DeleteImage(arg0 context.Context, arg1 int64) ([]*models.ImageResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteImage", arg0, arg1)
	ret0, _ := ret[0].([]*models.ImageResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteImage indicates an expected call of DeleteImage.
func (mr *MockImageRepoMockRecorder) DeleteImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteImage", reflect.TypeOf((*MockImageRepo)(nil).DeleteImage), arg0, arg1)
}

// SelectImages mocks base method.
func (m *MockImageRepo) SelectImages(arg0 context.Context, arg1 int64) ([]*models.ImageResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectImages", arg0, arg1)
	ret0, _ := ret[0].([]*models.ImageResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectImages indicates an expected call of SelectImages.
func (mr *MockImageRepoMockRecorder) SelectImages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectImages", reflect.TypeOf((*MockImageRepo)(nil).SelectImages), arg0, arg1)
}

// StoreImage mocks base method.
func (m *MockImageRepo) StoreImage(arg0 context.Context, arg1 *models.Image) (*models.ImageResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreImage", arg0, arg1)
	ret0, _ := ret[0].(*models.ImageResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreImage indicates an expected call of StoreImage.
func (mr *MockImageRepoMockRecorder) StoreImage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreImage", reflect.TypeOf((*MockImageRepo)(nil).StoreImage), arg0, arg1)
}
