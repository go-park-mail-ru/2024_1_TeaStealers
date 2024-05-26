package usecase_test

import (
	"2024_1_TeaStealers/internal/models"
	trans_mock "2024_1_TeaStealers/internal/models/mock"
	adverts_mock "2024_1_TeaStealers/internal/pkg/adverts/mock"
	"2024_1_TeaStealers/internal/pkg/adverts/usecase"
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetAdvertById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := adverts_mock.NewMockAdvertRepo(ctrl)
	logger := zap.Must(zap.NewDevelopment())
	usecase := usecase.NewAdvertUsecase(mockRepo, logger)

	advData := &models.AdvertData{
		ID:           123,
		AdvertType:   "House",
		TypeSale:     "Sale",
		Title:        "Beautiful House for Sale",
		Description:  "Spacious house with a large garden",
		Price:        100000,
		Phone:        "123-456-7890",
		IsAgent:      true,
		Address:      "123 Main St, Cityville",
		AddressPoint: "Coordinates",
		//Images:       []*models.ImageResp{},
		HouseProperties: &models.HouseProperties{
			CeilingHeight: 2.7,
			SquareArea:    200.5,
			SquareHouse:   180.0,
			BedroomCount:  4,
			StatusArea:    "Living room, kitchen, bedroom",
			Cottage:       false,
			StatusHome:    "New",
			Floor:         2,
		},
		ComplexProperties: &models.ComplexAdvertProperties{
			ComplexId:    "1234",
			NameComplex:  "Luxury Estates",
			PhotoCompany: "luxury_estates.jpg",
			NameCompany:  "Elite Realty",
		},
		//YearCreation: time.Now().Year(),
		Material: "Brick",
		//DateCreation: time.Now(),
	}

	type args struct {
		id int64
	}
	type want struct {
		adv *models.AdvertData
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful GetAdvertById",
			args: args{
				id: advData.ID,
			},
			want: want{
				adv: advData,
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := models.AdvertTypeHouse
			mockRepo.EXPECT().GetTypeAdvertById(gomock.Any(), gomock.Any()).Return(&a, nil)
			mockRepo.EXPECT().GetHouseAdvertById(gomock.Any(), gomock.Any()).Return(tt.want.adv, nil)
			mockRepo.EXPECT().SelectImages(gomock.Any(), gomock.Any()).Return(nil, nil)
			mockRepo.EXPECT().SelectPriceChanges(gomock.Any(), gomock.Any()).Return(nil, nil)
			mockRepo.EXPECT().SelectCountLikes(gomock.Any(), gomock.Any()).Return(int64(123), nil)
			mockRepo.EXPECT().SelectCountViews(gomock.Any(), gomock.Any()).Return(int64(123), nil)
			_, goterr := usecase.GetAdvertById(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tt.args.id)

			assert.Equal(t, tt.want.err, goterr)

		})
	}
}

func TestLUpdateAdvertById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := adverts_mock.NewMockAdvertRepo(ctrl)
	logger := zap.Must(zap.NewDevelopment())
	usecase := usecase.NewAdvertUsecase(mockRepo, logger)

	advData := &models.AdvertData{
		ID:           123,
		AdvertType:   "House",
		TypeSale:     "Sale",
		Title:        "Beautiful House for Sale",
		Description:  "Spacious house with a large garden",
		Price:        100000,
		Phone:        "123-456-7890",
		IsAgent:      true,
		Address:      "123 Main St, Cityville",
		AddressPoint: "Coordinates",
		//Images:       []*models.ImageResp{},
		HouseProperties: &models.HouseProperties{
			CeilingHeight: 2.7,
			SquareArea:    200.5,
			SquareHouse:   180.0,
			BedroomCount:  4,
			StatusArea:    "Living room, kitchen, bedroom",
			Cottage:       false,
			StatusHome:    "New",
			Floor:         2,
		},
		ComplexProperties: &models.ComplexAdvertProperties{
			ComplexId:    "1234",
			NameComplex:  "Luxury Estates",
			PhotoCompany: "luxury_estates.jpg",
			NameCompany:  "Elite Realty",
		},
		//YearCreation: time.Now().Year(),
		Material: "Brick",
		//DateCreation: time.Now(),
	}

	type args struct {
		id int64
	}
	type want struct {
		adv *models.AdvertData
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful GetAdvertById",
			args: args{
				id: advData.ID,
			},
			want: want{
				adv: advData,
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl2 := gomock.NewController(t)
			mockTr := trans_mock.NewMockTransaction(ctrl2)
			mockRepo.EXPECT().BeginTx(gomock.Any()).Return(mockTr, nil)
			a := models.AdvertTypeHouse
			mockRepo.EXPECT().GetTypeAdvertById(gomock.Any(), gomock.Any()).Return(&a, nil)
			mockRepo.EXPECT().ChangeTypeAdvert(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			// mockRepo.EXPECT().UpdateHouseAdvertById(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockTr.EXPECT().Commit().Return(nil)
			mockTr.EXPECT().Rollback().Return(nil)

			goterr := usecase.UpdateAdvertById(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), &models.AdvertUpdateData{})
			// assert.Equal(t, tt.want.adv, gotAdv)
			assert.Equal(t, tt.want.err, goterr)
			ctrl2.Finish()
		})
	}
}

func TestDeleteAdvertById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := adverts_mock.NewMockAdvertRepo(ctrl)
	logger := zap.Must(zap.NewDevelopment())
	usecase := usecase.NewAdvertUsecase(mockRepo, logger)

	advData := &models.AdvertData{
		ID:           123,
		AdvertType:   "House",
		TypeSale:     "Sale",
		Title:        "Beautiful House for Sale",
		Description:  "Spacious house with a large garden",
		Price:        100000,
		Phone:        "123-456-7890",
		IsAgent:      true,
		Address:      "123 Main St, Cityville",
		AddressPoint: "Coordinates",
		//Images:       []*models.ImageResp{},
		HouseProperties: &models.HouseProperties{
			CeilingHeight: 2.7,
			SquareArea:    200.5,
			SquareHouse:   180.0,
			BedroomCount:  4,
			StatusArea:    "Living room, kitchen, bedroom",
			Cottage:       false,
			StatusHome:    "New",
			Floor:         2,
		},
		ComplexProperties: &models.ComplexAdvertProperties{
			ComplexId:    "1234",
			NameComplex:  "Luxury Estates",
			PhotoCompany: "luxury_estates.jpg",
			NameCompany:  "Elite Realty",
		},
		//YearCreation: time.Now().Year(),
		Material: "Brick",
		//DateCreation: time.Now(),
	}

	type args struct {
		idi int64
	}
	type want struct {
		adv *models.AdvertData
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful DeleteAdvertByID",
			args: args{
				idi: advData.ID,
			},
			want: want{
				adv: advData,
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl2 := gomock.NewController(t)
			mockTr := trans_mock.NewMockTransaction(ctrl2)
			mockRepo.EXPECT().BeginTx(gomock.Any()).Return(mockTr, nil)
			a := models.AdvertTypeHouse
			mockRepo.EXPECT().GetTypeAdvertById(gomock.Any(), gomock.Any()).Return(&a, nil)
			// mockRepo.EXPECT().ChangeTypeAdvert(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockRepo.EXPECT().DeleteHouseAdvertById(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			// mockRepo.EXPECT().UpdateHouseAdvertById(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockTr.EXPECT().Commit().Return(nil)
			mockTr.EXPECT().Rollback().Return(nil)

			goterr := usecase.DeleteAdvertById(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), 123)
			// assert.Equal(t, tt.want.adv, gotAdv)
			assert.Equal(t, tt.want.err, goterr)
			ctrl2.Finish()
		})
	}
}

func TestGetSquareAdvertsList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := adverts_mock.NewMockAdvertRepo(ctrl)
	logger := zap.Must(zap.NewDevelopment())
	usecase := usecase.NewAdvertUsecase(mockRepo, logger)

	// Prepare test data
	ctx := context.WithValue(context.Background(), "requestId", uuid.NewV4().String())
	// filter := AdvertFilter{ /* populate filter if needed */ /*}
	mockResult := &models.AdvertSquareData{
		ID:           123,
		TypeAdvert:   "House",
		Photo:        "example.jpg",
		TypeSale:     "Sale",
		Address:      "123 Main Street",
		Price:        1000,
		DateCreation: time.Now(),
	}
	adverts := []*models.AdvertSquareData{mockResult, mockResult}
	// Set expectations
	mockRepo.EXPECT().GetSquareAdverts(gomock.Any(), gomock.Any(), gomock.Any()).Return(adverts, nil)

	// Call the method under test
	result, err := usecase.GetSquareAdvertsList(ctx, 2, 4)

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, adverts, result)
}

func TestGetRectangleAdvertsList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := adverts_mock.NewMockAdvertRepo(ctrl)
	logger := zap.Must(zap.NewDevelopment())
	usecase := usecase.NewAdvertUsecase(mockRepo, logger)

	// Prepare test data
	ctx := context.WithValue(context.Background(), "requestId", uuid.NewV4().String())
	// filter := AdvertFilter{ /* populate filter if needed */ /*}
	// mockResult := &models.AdvertDataPage{}
	// Set expectations
	mockRepo.EXPECT().GetRectangleAdverts(gomock.Any(), gomock.Any()).Return(&models.AdvertDataPage{}, nil)

	// Call the method under test
	_, err := usecase.GetRectangleAdvertsList(ctx, models.AdvertFilter{})

	// Assert the results
	assert.NoError(t, err)
	// assert.Equal(t, adverts, result)
}

/*
func TestGetExistBuildingsByAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := adverts_mock.NewMockAdvertRepo(ctrl)
	logger := zap.Must(zap.NewDevelopment())
	usecase := usecase.NewAdvertUsecase(mockRepo, logger)

	// Prepare test data
	ctx := context.WithValue(context.Background(), "requestId", uuid.NewV4().String())
	// filter := AdvertFilter{ /* populate filter if needed */ /*}
	// mockResult := &models.AdvertDataPage{}
	// Set expectations
	mockRepo.EXPECT().CheckExistsBuildings(gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.BuildingData{}, nil)

	// Call the method under test
	_, err := usecase.GetExistBuildingsByAddress(ctx, "address", 2)

	// Assert the results
	assert.NoError(t, err)
	// assert.Equal(t, adverts, result)
}
*/

func TestGetRectangleAdvertsByUserId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := adverts_mock.NewMockAdvertRepo(ctrl)
	logger := zap.Must(zap.NewDevelopment())
	usecase := usecase.NewAdvertUsecase(mockRepo, logger)

	// Prepare test data
	ctx := context.WithValue(context.Background(), "requestId", uuid.NewV4().String())
	// filter := AdvertFilter{ /* populate filter if needed */ /*}
	// mockResult := &models.AdvertDataPage{}
	// Set expectations
	mockRepo.EXPECT().GetRectangleAdvertsByUserId(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.AdvertRectangleData{}, nil)

	// Call the method under test
	_, err := usecase.GetRectangleAdvertsByUserId(ctx, 2, 3, 123)

	// Assert the results
	assert.NoError(t, err)
	// assert.Equal(t, adverts, result)
}

func TestGetRectangleAdvertsByComplexId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := adverts_mock.NewMockAdvertRepo(ctrl)
	logger := zap.Must(zap.NewDevelopment())
	usecase := usecase.NewAdvertUsecase(mockRepo, logger)

	// Prepare test data
	ctx := context.WithValue(context.Background(), "requestId", uuid.NewV4().String())
	// filter := AdvertFilter{ /* populate filter if needed */ /*}
	// mockResult := &models.AdvertDataPage{}
	// Set expectations
	mockRepo.EXPECT().GetRectangleAdvertsByComplexId(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]*models.AdvertRectangleData{}, nil)

	// Call the method under test
	_, err := usecase.GetRectangleAdvertsByComplexId(ctx, 2, 3, 123)

	// Assert the results
	assert.NoError(t, err)
	// assert.Equal(t, adverts, result)
}
