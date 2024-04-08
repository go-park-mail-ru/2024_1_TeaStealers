// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package adverts_mock is a generated GoMock package.
package adverts_mock

import (
	models "2024_1_TeaStealers/internal/models"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/satori/uuid"
)

// MockAdvertUsecase is a mock of AdvertUsecase interface.
type MockAdvertUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockAdvertUsecaseMockRecorder
}

// MockAdvertUsecaseMockRecorder is the mock recorder for MockAdvertUsecase.
type MockAdvertUsecaseMockRecorder struct {
	mock *MockAdvertUsecase
}

// NewMockAdvertUsecase creates a new mock instance.
func NewMockAdvertUsecase(ctrl *gomock.Controller) *MockAdvertUsecase {
	mock := &MockAdvertUsecase{ctrl: ctrl}
	mock.recorder = &MockAdvertUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdvertUsecase) EXPECT() *MockAdvertUsecaseMockRecorder {
	return m.recorder
}

// CreateFlatAdvert mocks base method.
func (m *MockAdvertUsecase) CreateFlatAdvert(arg0 context.Context, arg1 *models.AdvertFlatCreateData) (*models.Advert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFlatAdvert", arg0, arg1)
	ret0, _ := ret[0].(*models.Advert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFlatAdvert indicates an expected call of CreateFlatAdvert.
func (mr *MockAdvertUsecaseMockRecorder) CreateFlatAdvert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFlatAdvert", reflect.TypeOf((*MockAdvertUsecase)(nil).CreateFlatAdvert), arg0, arg1)
}

// CreateHouseAdvert mocks base method.
func (m *MockAdvertUsecase) CreateHouseAdvert(arg0 context.Context, arg1 *models.AdvertHouseCreateData) (*models.Advert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHouseAdvert", arg0, arg1)
	ret0, _ := ret[0].(*models.Advert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateHouseAdvert indicates an expected call of CreateHouseAdvert.
func (mr *MockAdvertUsecaseMockRecorder) CreateHouseAdvert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHouseAdvert", reflect.TypeOf((*MockAdvertUsecase)(nil).CreateHouseAdvert), arg0, arg1)
}

// DeleteAdvertById mocks base method.
func (m *MockAdvertUsecase) DeleteAdvertById(ctx context.Context, advertId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAdvertById", ctx, advertId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAdvertById indicates an expected call of DeleteAdvertById.
func (mr *MockAdvertUsecaseMockRecorder) DeleteAdvertById(ctx, advertId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAdvertById", reflect.TypeOf((*MockAdvertUsecase)(nil).DeleteAdvertById), ctx, advertId)
}

// GetAdvertById mocks base method.
func (m *MockAdvertUsecase) GetAdvertById(ctx context.Context, id uuid.UUID) (*models.AdvertData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdvertById", ctx, id)
	ret0, _ := ret[0].(*models.AdvertData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdvertById indicates an expected call of GetAdvertById.
func (mr *MockAdvertUsecaseMockRecorder) GetAdvertById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdvertById", reflect.TypeOf((*MockAdvertUsecase)(nil).GetAdvertById), ctx, id)
}

// GetExistBuildingsByAddress mocks base method.
func (m *MockAdvertUsecase) GetExistBuildingsByAddress(ctx context.Context, address string, pageSize int) ([]*models.BuildingData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExistBuildingsByAddress", ctx, address, pageSize)
	ret0, _ := ret[0].([]*models.BuildingData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExistBuildingsByAddress indicates an expected call of GetExistBuildingsByAddress.
func (mr *MockAdvertUsecaseMockRecorder) GetExistBuildingsByAddress(ctx, address, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExistBuildingsByAddress", reflect.TypeOf((*MockAdvertUsecase)(nil).GetExistBuildingsByAddress), ctx, address, pageSize)
}

// GetRectangleAdvertsByComplexId mocks base method.
func (m *MockAdvertUsecase) GetRectangleAdvertsByComplexId(ctx context.Context, pageSize, offset int, comlexId uuid.UUID) ([]*models.AdvertRectangleData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRectangleAdvertsByComplexId", ctx, pageSize, offset, comlexId)
	ret0, _ := ret[0].([]*models.AdvertRectangleData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRectangleAdvertsByComplexId indicates an expected call of GetRectangleAdvertsByComplexId.
func (mr *MockAdvertUsecaseMockRecorder) GetRectangleAdvertsByComplexId(ctx, pageSize, offset, comlexId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRectangleAdvertsByComplexId", reflect.TypeOf((*MockAdvertUsecase)(nil).GetRectangleAdvertsByComplexId), ctx, pageSize, offset, comlexId)
}

// GetRectangleAdvertsByUserId mocks base method.
func (m *MockAdvertUsecase) GetRectangleAdvertsByUserId(ctx context.Context, pageSize, offset int, userId uuid.UUID) ([]*models.AdvertRectangleData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRectangleAdvertsByUserId", ctx, pageSize, offset, userId)
	ret0, _ := ret[0].([]*models.AdvertRectangleData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRectangleAdvertsByUserId indicates an expected call of GetRectangleAdvertsByUserId.
func (mr *MockAdvertUsecaseMockRecorder) GetRectangleAdvertsByUserId(ctx, pageSize, offset, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRectangleAdvertsByUserId", reflect.TypeOf((*MockAdvertUsecase)(nil).GetRectangleAdvertsByUserId), ctx, pageSize, offset, userId)
}

// GetRectangleAdvertsList mocks base method.
func (m *MockAdvertUsecase) GetRectangleAdvertsList(ctx context.Context, advertFilter models.AdvertFilter) (*models.AdvertDataPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRectangleAdvertsList", ctx, advertFilter)
	ret0, _ := ret[0].(*models.AdvertDataPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRectangleAdvertsList indicates an expected call of GetRectangleAdvertsList.
func (mr *MockAdvertUsecaseMockRecorder) GetRectangleAdvertsList(ctx, advertFilter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRectangleAdvertsList", reflect.TypeOf((*MockAdvertUsecase)(nil).GetRectangleAdvertsList), ctx, advertFilter)
}

// GetSquareAdvertsList mocks base method.
func (m *MockAdvertUsecase) GetSquareAdvertsList(ctx context.Context, pageSize, offset int) ([]*models.AdvertSquareData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSquareAdvertsList", ctx, pageSize, offset)
	ret0, _ := ret[0].([]*models.AdvertSquareData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSquareAdvertsList indicates an expected call of GetSquareAdvertsList.
func (mr *MockAdvertUsecaseMockRecorder) GetSquareAdvertsList(ctx, pageSize, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSquareAdvertsList", reflect.TypeOf((*MockAdvertUsecase)(nil).GetSquareAdvertsList), ctx, pageSize, offset)
}

// UpdateAdvertById mocks base method.
func (m *MockAdvertUsecase) UpdateAdvertById(ctx context.Context, advertUpdateData *models.AdvertUpdateData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdvertById", ctx, advertUpdateData)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAdvertById indicates an expected call of UpdateAdvertById.
func (mr *MockAdvertUsecaseMockRecorder) UpdateAdvertById(ctx, advertUpdateData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdvertById", reflect.TypeOf((*MockAdvertUsecase)(nil).UpdateAdvertById), ctx, advertUpdateData)
}

// MockAdvertRepo is a mock of AdvertRepo interface.
type MockAdvertRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAdvertRepoMockRecorder
}

// MockAdvertRepoMockRecorder is the mock recorder for MockAdvertRepo.
type MockAdvertRepoMockRecorder struct {
	mock *MockAdvertRepo
}

// NewMockAdvertRepo creates a new mock instance.
func NewMockAdvertRepo(ctrl *gomock.Controller) *MockAdvertRepo {
	mock := &MockAdvertRepo{ctrl: ctrl}
	mock.recorder = &MockAdvertRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdvertRepo) EXPECT() *MockAdvertRepoMockRecorder {
	return m.recorder
}

// BeginTx mocks base method.
func (m *MockAdvertRepo) BeginTx(ctx context.Context) (models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTx", ctx)
	ret0, _ := ret[0].(models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockAdvertRepoMockRecorder) BeginTx(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockAdvertRepo)(nil).BeginTx), ctx)
}

// ChangeTypeAdvert mocks base method.
func (m *MockAdvertRepo) ChangeTypeAdvert(ctx context.Context, tx models.Transaction, advertId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeTypeAdvert", ctx, tx, advertId)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeTypeAdvert indicates an expected call of ChangeTypeAdvert.
func (mr *MockAdvertRepoMockRecorder) ChangeTypeAdvert(ctx, tx, advertId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeTypeAdvert", reflect.TypeOf((*MockAdvertRepo)(nil).ChangeTypeAdvert), ctx, tx, advertId)
}

// CheckExistsBuilding mocks base method.
func (m *MockAdvertRepo) CheckExistsBuilding(ctx context.Context, adress string) (*models.Building, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExistsBuilding", ctx, adress)
	ret0, _ := ret[0].(*models.Building)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckExistsBuilding indicates an expected call of CheckExistsBuilding.
func (mr *MockAdvertRepoMockRecorder) CheckExistsBuilding(ctx, adress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExistsBuilding", reflect.TypeOf((*MockAdvertRepo)(nil).CheckExistsBuilding), ctx, adress)
}

// CheckExistsBuildings mocks base method.
func (m *MockAdvertRepo) CheckExistsBuildings(ctx context.Context, pageSize int, adress string) ([]*models.BuildingData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExistsBuildings", ctx, pageSize, adress)
	ret0, _ := ret[0].([]*models.BuildingData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckExistsBuildings indicates an expected call of CheckExistsBuildings.
func (mr *MockAdvertRepoMockRecorder) CheckExistsBuildings(ctx, pageSize, adress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExistsBuildings", reflect.TypeOf((*MockAdvertRepo)(nil).CheckExistsBuildings), ctx, pageSize, adress)
}

// CreateAdvert mocks base method.
func (m *MockAdvertRepo) CreateAdvert(ctx context.Context, tx models.Transaction, newAdvert *models.Advert) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAdvert", ctx, tx, newAdvert)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAdvert indicates an expected call of CreateAdvert.
func (mr *MockAdvertRepoMockRecorder) CreateAdvert(ctx, tx, newAdvert interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdvert", reflect.TypeOf((*MockAdvertRepo)(nil).CreateAdvert), ctx, tx, newAdvert)
}

// CreateAdvertType mocks base method.
func (m *MockAdvertRepo) CreateAdvertType(ctx context.Context, tx models.Transaction, newAdvertType *models.AdvertType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAdvertType", ctx, tx, newAdvertType)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAdvertType indicates an expected call of CreateAdvertType.
func (mr *MockAdvertRepoMockRecorder) CreateAdvertType(ctx, tx, newAdvertType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdvertType", reflect.TypeOf((*MockAdvertRepo)(nil).CreateAdvertType), ctx, tx, newAdvertType)
}

// CreateBuilding mocks base method.
func (m *MockAdvertRepo) CreateBuilding(ctx context.Context, tx models.Transaction, newBuilding *models.Building) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBuilding", ctx, tx, newBuilding)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBuilding indicates an expected call of CreateBuilding.
func (mr *MockAdvertRepoMockRecorder) CreateBuilding(ctx, tx, newBuilding interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBuilding", reflect.TypeOf((*MockAdvertRepo)(nil).CreateBuilding), ctx, tx, newBuilding)
}

// CreateFlat mocks base method.
func (m *MockAdvertRepo) CreateFlat(ctx context.Context, tx models.Transaction, newFlat *models.Flat) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFlat", ctx, tx, newFlat)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFlat indicates an expected call of CreateFlat.
func (mr *MockAdvertRepoMockRecorder) CreateFlat(ctx, tx, newFlat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFlat", reflect.TypeOf((*MockAdvertRepo)(nil).CreateFlat), ctx, tx, newFlat)
}

// CreateHouse mocks base method.
func (m *MockAdvertRepo) CreateHouse(ctx context.Context, tx models.Transaction, newHouse *models.House) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHouse", ctx, tx, newHouse)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateHouse indicates an expected call of CreateHouse.
func (mr *MockAdvertRepoMockRecorder) CreateHouse(ctx, tx, newHouse interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHouse", reflect.TypeOf((*MockAdvertRepo)(nil).CreateHouse), ctx, tx, newHouse)
}

// CreatePriceChange mocks base method.
func (m *MockAdvertRepo) CreatePriceChange(ctx context.Context, tx models.Transaction, newPriceChange *models.PriceChange) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePriceChange", ctx, tx, newPriceChange)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePriceChange indicates an expected call of CreatePriceChange.
func (mr *MockAdvertRepoMockRecorder) CreatePriceChange(ctx, tx, newPriceChange interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePriceChange", reflect.TypeOf((*MockAdvertRepo)(nil).CreatePriceChange), ctx, tx, newPriceChange)
}

// DeleteFlatAdvertById mocks base method.
func (m *MockAdvertRepo) DeleteFlatAdvertById(ctx context.Context, tx models.Transaction, advertId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFlatAdvertById", ctx, tx, advertId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFlatAdvertById indicates an expected call of DeleteFlatAdvertById.
func (mr *MockAdvertRepoMockRecorder) DeleteFlatAdvertById(ctx, tx, advertId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFlatAdvertById", reflect.TypeOf((*MockAdvertRepo)(nil).DeleteFlatAdvertById), ctx, tx, advertId)
}

// DeleteHouseAdvertById mocks base method.
func (m *MockAdvertRepo) DeleteHouseAdvertById(ctx context.Context, tx models.Transaction, advertId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHouseAdvertById", ctx, tx, advertId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHouseAdvertById indicates an expected call of DeleteHouseAdvertById.
func (mr *MockAdvertRepoMockRecorder) DeleteHouseAdvertById(ctx, tx, advertId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHouseAdvertById", reflect.TypeOf((*MockAdvertRepo)(nil).DeleteHouseAdvertById), ctx, tx, advertId)
}

// GetFlatAdvertById mocks base method.
func (m *MockAdvertRepo) GetFlatAdvertById(ctx context.Context, id uuid.UUID) (*models.AdvertData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlatAdvertById", ctx, id)
	ret0, _ := ret[0].(*models.AdvertData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFlatAdvertById indicates an expected call of GetFlatAdvertById.
func (mr *MockAdvertRepoMockRecorder) GetFlatAdvertById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlatAdvertById", reflect.TypeOf((*MockAdvertRepo)(nil).GetFlatAdvertById), ctx, id)
}

// GetHouseAdvertById mocks base method.
func (m *MockAdvertRepo) GetHouseAdvertById(ctx context.Context, id uuid.UUID) (*models.AdvertData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHouseAdvertById", ctx, id)
	ret0, _ := ret[0].(*models.AdvertData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHouseAdvertById indicates an expected call of GetHouseAdvertById.
func (mr *MockAdvertRepoMockRecorder) GetHouseAdvertById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHouseAdvertById", reflect.TypeOf((*MockAdvertRepo)(nil).GetHouseAdvertById), ctx, id)
}

// GetRectangleAdverts mocks base method.
func (m *MockAdvertRepo) GetRectangleAdverts(ctx context.Context, advertFilter models.AdvertFilter) (*models.AdvertDataPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRectangleAdverts", ctx, advertFilter)
	ret0, _ := ret[0].(*models.AdvertDataPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRectangleAdverts indicates an expected call of GetRectangleAdverts.
func (mr *MockAdvertRepoMockRecorder) GetRectangleAdverts(ctx, advertFilter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRectangleAdverts", reflect.TypeOf((*MockAdvertRepo)(nil).GetRectangleAdverts), ctx, advertFilter)
}

// GetRectangleAdvertsByComplexId mocks base method.
func (m *MockAdvertRepo) GetRectangleAdvertsByComplexId(ctx context.Context, pageSize, offset int, complexId uuid.UUID) ([]*models.AdvertRectangleData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRectangleAdvertsByComplexId", ctx, pageSize, offset, complexId)
	ret0, _ := ret[0].([]*models.AdvertRectangleData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRectangleAdvertsByComplexId indicates an expected call of GetRectangleAdvertsByComplexId.
func (mr *MockAdvertRepoMockRecorder) GetRectangleAdvertsByComplexId(ctx, pageSize, offset, complexId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRectangleAdvertsByComplexId", reflect.TypeOf((*MockAdvertRepo)(nil).GetRectangleAdvertsByComplexId), ctx, pageSize, offset, complexId)
}

// GetRectangleAdvertsByUserId mocks base method.
func (m *MockAdvertRepo) GetRectangleAdvertsByUserId(ctx context.Context, pageSize, offset int, userId uuid.UUID) ([]*models.AdvertRectangleData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRectangleAdvertsByUserId", ctx, pageSize, offset, userId)
	ret0, _ := ret[0].([]*models.AdvertRectangleData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRectangleAdvertsByUserId indicates an expected call of GetRectangleAdvertsByUserId.
func (mr *MockAdvertRepoMockRecorder) GetRectangleAdvertsByUserId(ctx, pageSize, offset, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRectangleAdvertsByUserId", reflect.TypeOf((*MockAdvertRepo)(nil).GetRectangleAdvertsByUserId), ctx, pageSize, offset, userId)
}

// GetSquareAdverts mocks base method.
func (m *MockAdvertRepo) GetSquareAdverts(ctx context.Context, pageSize, offset int) ([]*models.AdvertSquareData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSquareAdverts", ctx, pageSize, offset)
	ret0, _ := ret[0].([]*models.AdvertSquareData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSquareAdverts indicates an expected call of GetSquareAdverts.
func (mr *MockAdvertRepoMockRecorder) GetSquareAdverts(ctx, pageSize, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSquareAdverts", reflect.TypeOf((*MockAdvertRepo)(nil).GetSquareAdverts), ctx, pageSize, offset)
}

// GetTypeAdvertById mocks base method.
func (m *MockAdvertRepo) GetTypeAdvertById(ctx context.Context, id uuid.UUID) (*models.AdvertTypeAdvert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTypeAdvertById", ctx, id)
	ret0, _ := ret[0].(*models.AdvertTypeAdvert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTypeAdvertById indicates an expected call of GetTypeAdvertById.
func (mr *MockAdvertRepoMockRecorder) GetTypeAdvertById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTypeAdvertById", reflect.TypeOf((*MockAdvertRepo)(nil).GetTypeAdvertById), ctx, id)
}

// SelectImages mocks base method.
func (m *MockAdvertRepo) SelectImages(advertId uuid.UUID) ([]*models.ImageResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectImages", advertId)
	ret0, _ := ret[0].([]*models.ImageResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectImages indicates an expected call of SelectImages.
func (mr *MockAdvertRepoMockRecorder) SelectImages(advertId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectImages", reflect.TypeOf((*MockAdvertRepo)(nil).SelectImages), advertId)
}

// UpdateFlatAdvertById mocks base method.
func (m *MockAdvertRepo) UpdateFlatAdvertById(ctx context.Context, tx models.Transaction, advertUpdateData *models.AdvertUpdateData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFlatAdvertById", ctx, tx, advertUpdateData)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFlatAdvertById indicates an expected call of UpdateFlatAdvertById.
func (mr *MockAdvertRepoMockRecorder) UpdateFlatAdvertById(ctx, tx, advertUpdateData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFlatAdvertById", reflect.TypeOf((*MockAdvertRepo)(nil).UpdateFlatAdvertById), ctx, tx, advertUpdateData)
}

// UpdateHouseAdvertById mocks base method.
func (m *MockAdvertRepo) UpdateHouseAdvertById(ctx context.Context, tx models.Transaction, advertUpdateData *models.AdvertUpdateData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHouseAdvertById", ctx, tx, advertUpdateData)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHouseAdvertById indicates an expected call of UpdateHouseAdvertById.
func (mr *MockAdvertRepoMockRecorder) UpdateHouseAdvertById(ctx, tx, advertUpdateData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHouseAdvertById", reflect.TypeOf((*MockAdvertRepo)(nil).UpdateHouseAdvertById), ctx, tx, advertUpdateData)
}
