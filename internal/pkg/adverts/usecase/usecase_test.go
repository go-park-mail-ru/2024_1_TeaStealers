package usecase_test

/*
import (
	"2024_1_TeaStealers/internal/models"
	trans_mock "2024_1_TeaStealers/internal/models/mock"
	adverts_mock "2024_1_TeaStealers/internal/pkg/adverts/mock"
	"2024_1_TeaStealers/internal/pkg/adverts/usecase"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCreateFlatAdvert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := adverts_mock.NewMockAdvertRepo(ctrl)
	logger := zap.Must(zap.NewDevelopment())
	usecase := usecase.NewAdvertUsecase(mockRepo, logger)
	id1 := uuid.NewV4()
	type args struct {
		data *models.AdvertFlatCreateData
		// newAdvertType  *models.AdvertType
		// newAdvert      *models.Advert
		// building       *models.Building
		// newFlat        *models.Flat
		// newPriceChange *models.PriceChange
	}
	type want struct {
		adv *models.Advert
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful get user",
			args: args{
				data: &models.AdvertFlatCreateData{
					UserID:              id1,
					AdvertTypeSale:      models.TypePlacementSale,
					AdvertTypePlacement: models.AdvertTypeFlat,
					Title:               "Test Flat Advert",
					Description:         "This is a test flat advert description.",
					Phone:               "123-456-789",
					IsAgent:             false,
					Floor:               2,
					CeilingHeight:       2.5,
					SquareGeneral:       80.5,
					RoomCount:           3,
					SquareResidential:   65.0,
					Apartment:           false,
					Price:               150000,
					FloorGeneral:        5,
					Material:            models.MaterialStalinsky,
					Address:             "123 Test Street",
					AddressPoint:        "51.5074° N, 0.1278° W",
					YearCreation:        2020,
				},
			},
			want: want{
				adv: &models.Advert{
					ID:             uuid.NewV4(),
					UserID:         id1,
					AdvertTypeID:   uuid.NewV4(),
					AdvertTypeSale: models.TypePlacementSale,
					Title:          "Test Flat Advert",
					Description:    "This is a test flat advert description.",
					Phone:          "123-456-789",
					IsAgent:        false,
					Priority:       0,
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl2 := gomock.NewController(t)
			mockTr := trans_mock.NewMockTransaction(ctrl2)
			mockRepo.EXPECT().BeginTx(gomock.Any()).Return(mockTr, nil)
			mockRepo.EXPECT().CheckExistsBuilding(gomock.Any(), gomock.Any()).Return(&models.Building{}, errors.New("no"))
			mockRepo.EXPECT().CreateBuilding(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockRepo.EXPECT().CreateAdvertType(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockRepo.EXPECT().CreateFlat(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockRepo.EXPECT().CreateAdvert(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockRepo.EXPECT().CreatePriceChange(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockTr.EXPECT().Commit().Return(nil)
			mockTr.EXPECT().Rollback().Return(nil)

			_, goterr := usecase.CreateFlatAdvert(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tt.args.data)
			// assert.Equal(t, tt.want.adv, gotAdv)
			assert.Equal(t, tt.want.err, goterr)
			ctrl2.Finish()

		})
	}
}

func TestCreateHouseAdvert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := adverts_mock.NewMockAdvertRepo(ctrl)
	logger := zap.Must(zap.NewDevelopment())
	usecase := usecase.NewAdvertUsecase(mockRepo, logger)
	id1 := uuid.NewV4()
	type args struct {
		data *models.AdvertHouseCreateData
		// newAdvertType  *models.AdvertType
		// newAdvert      *models.Advert
		// building       *models.Building
		// newFlat        *models.Flat
		// newPriceChange *models.PriceChange
	}
	type want struct {
		adv *models.Advert
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful get user",
			args: args{
				data: &models.AdvertHouseCreateData{
					UserID:              id1,
					AdvertTypeSale:      models.TypePlacementSale,
					AdvertTypePlacement: models.AdvertTypeFlat,
					Title:               "Test Flat Advert",
					Description:         "This is a test flat advert description.",
					Phone:               "123-456-789",
					IsAgent:             false,
					CeilingHeight:       2.5,
					Price:               150000,
					FloorGeneral:        5,
					Material:            models.MaterialStalinsky,
					Address:             "123 Test Street",
					AddressPoint:        "51.5074° N, 0.1278° W",
					YearCreation:        2020,
				},
			},
			want: want{
				adv: &models.Advert{
					ID:             uuid.NewV4(),
					UserID:         id1,
					AdvertTypeID:   uuid.NewV4(),
					AdvertTypeSale: models.TypePlacementSale,
					Title:          "Test Flat Advert",
					Description:    "This is a test flat advert description.",
					Phone:          "123-456-789",
					IsAgent:        false,
					Priority:       0,
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl2 := gomock.NewController(t)
			mockTr := trans_mock.NewMockTransaction(ctrl2)
			mockRepo.EXPECT().BeginTx(gomock.Any()).Return(mockTr, nil)
			mockRepo.EXPECT().CheckExistsBuilding(gomock.Any(), gomock.Any()).Return(&models.Building{}, errors.New("no"))
			mockRepo.EXPECT().CreateBuilding(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockRepo.EXPECT().CreateAdvertType(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockRepo.EXPECT().CreateHouse(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockRepo.EXPECT().CreateAdvert(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockRepo.EXPECT().CreatePriceChange(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			mockTr.EXPECT().Commit().Return(nil)
			mockTr.EXPECT().Rollback().Return(nil)

			_, goterr := usecase.CreateHouseAdvert(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tt.args.data)
			// assert.Equal(t, tt.want.adv, gotAdv)
			assert.Equal(t, tt.want.err, goterr)
			ctrl2.Finish()

		})
	}
}

func TestGetAdvertById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := adverts_mock.NewMockAdvertRepo(ctrl)
	logger := zap.Must(zap.NewDevelopment())
	usecase := usecase.NewAdvertUsecase(mockRepo, logger)

	advData := &models.AdvertData{
		ID:           uuid.NewV4(),
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
		idi uuid.UUID
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
			a := models.AdvertTypeHouse
			mockRepo.EXPECT().GetTypeAdvertById(gomock.Any(), gomock.Any()).Return(&a, nil)
			mockRepo.EXPECT().GetHouseAdvertById(gomock.Any(), gomock.Any()).Return(tt.want.adv, nil)
			mockRepo.EXPECT().SelectImages(gomock.Any(), gomock.Any()).Return(nil, nil)
			_, goterr := usecase.GetAdvertById(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tt.args.idi)

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
		ID:           uuid.NewV4(),
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
		idi uuid.UUID
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
		ID:           uuid.NewV4(),
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
		idi uuid.UUID
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

			goterr := usecase.DeleteAdvertById(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), uuid.NewV4())
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
		ID:           uuid.NewV4(),
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
	_, err := usecase.GetRectangleAdvertsByUserId(ctx, 2, 3, uuid.NewV4())

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
	_, err := usecase.GetRectangleAdvertsByComplexId(ctx, 2, 3, uuid.NewV4())

	// Assert the results
	assert.NoError(t, err)
	// assert.Equal(t, adverts, result)
}*/
