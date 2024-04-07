package usecase_test

import (
	"2024_1_TeaStealers/internal/models"
	trans_mock "2024_1_TeaStealers/internal/models/mock"
	adverts_mock "2024_1_TeaStealers/internal/pkg/adverts/mock"

	"2024_1_TeaStealers/internal/pkg/adverts/usecase"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateFlatAdvert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := adverts_mock.NewMockAdvertRepo(ctrl)
	usecase := usecase.NewAdvertUsecase(mockRepo)
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

			_, goterr := usecase.CreateFlatAdvert(context.Background(), tt.args.data)
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
	usecase := usecase.NewAdvertUsecase(mockRepo)
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

			_, goterr := usecase.CreateHouseAdvert(context.Background(), tt.args.data)
			// assert.Equal(t, tt.want.adv, gotAdv)
			assert.Equal(t, tt.want.err, goterr)
			ctrl2.Finish()

		})
	}
}
