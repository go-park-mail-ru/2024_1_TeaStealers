package usecase_test

import (
	"2024_1_TeaStealers/internal/models"
	complexes_mock "2024_1_TeaStealers/internal/pkg/complexes/mock"
	"2024_1_TeaStealers/internal/pkg/complexes/usecase"
	"context"
	"errors"
	"math/rand"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCreateComplex(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	newComplex := &models.Complex{
		ID:                     123,
		CompanyId:              321,
		Name:                   "name",
		Description:            "descr",
		Address:                "adr",
		DateBeginBuild:         time.Date(2004, 12, 12, 0, 0, 0, 0, time.Local),
		DateEndBuild:           time.Date(2004, 12, 12, 0, 0, 0, 0, time.Local),
		WithoutFinishingOption: true,
		FinishingOption:        true,
		PreFinishingOption:     true,
		ClassHousing:           models.ClassHouseBusiness,
		Parking:                true,
		Security:               true,
	}

	createcompl := &models.ComplexCreateData{
		CompanyId:              321,
		Name:                   "name",
		Description:            "descr",
		Address:                "adr",
		DateBeginBuild:         time.Date(2004, 12, 12, 0, 0, 0, 0, time.Local),
		DateEndBuild:           time.Date(2004, 12, 12, 0, 0, 0, 0, time.Local),
		WithoutFinishingOption: true,
		FinishingOption:        true,
		PreFinishingOption:     true,
		ClassHousing:           models.ClassHouseBusiness,
		Parking:                true,
		Security:               true,
	}
	mockRepo := complexes_mock.NewMockComplexRepo(ctrl)
	usecase := usecase.NewComplexUsecase(mockRepo, &zap.Logger{})
	type args struct {
		crcompl *models.ComplexCreateData
	}
	type want struct {
		compl *models.Complex
		err   error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful create complex",
			args: args{
				crcompl: createcompl,
			},
			want: want{
				compl: newComplex,
				err:   nil,
			},
		},
		{
			name: "fail create complex",
			args: args{
				crcompl: createcompl,
			},
			want: want{
				compl: newComplex,
				err:   errors.New("fail create complex"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().CreateComplex(gomock.Any(), gomock.Any()).Return(tt.want.compl, tt.want.err)
			_, goterr := usecase.CreateComplex(context.Background(), tt.args.crcompl)
			// assert.Equal(t, tt.want.user, gotUser)
			assert.Equal(t, tt.want.err, goterr)
		})
	}
}

func TestCreateBuilding(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	newBuild := &models.BuildingCreateData{
		ComplexID:    321,
		Floor:        2,
		Material:     "steel",
		AddressID:    111,
		YearCreation: 2000,
	}

	createBuild := &models.Building{
		ID:           123,
		ComplexID:    321,
		Floor:        2,
		Material:     "steel",
		AddressID:    111,
		YearCreation: 2000,
		DateCreation: time.Now(),
		IsDeleted:    false,
	}
	mockRepo := complexes_mock.NewMockComplexRepo(ctrl)
	usecase := usecase.NewComplexUsecase(mockRepo, &zap.Logger{})
	type args struct {
		crbuild *models.BuildingCreateData
	}
	type want struct {
		building *models.Building
		err      error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful create building",
			args: args{
				crbuild: newBuild,
			},
			want: want{
				building: createBuild,
				err:      nil,
			},
		},
		{
			name: "fail create building",
			args: args{
				crbuild: newBuild,
			},
			want: want{
				building: createBuild,
				err:      errors.New("errror create"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().CreateBuilding(gomock.Any(), gomock.Any()).Return(tt.want.building, tt.want.err)
			_, goterr := usecase.CreateBuilding(context.Background(), tt.args.crbuild)
			// assert.Equal(t, tt.want.user, gotUser)
			assert.Equal(t, tt.want.err, goterr)
		})
	}
}

func TestGetComplexById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := complexes_mock.NewMockComplexRepo(ctrl)
	usecase := usecase.NewComplexUsecase(mockRepo, &zap.Logger{})
	name := "get complex by id ok"
	id := rand.Int63()
	foundCData := &models.ComplexData{}
	t.Run(name, func(t *testing.T) {
		mockRepo.EXPECT().GetComplexById(gomock.Any(), gomock.Any()).Return(foundCData, nil)
		_, goterr := usecase.GetComplexById(context.Background(), id)
		// assert.Equal(t, tt.want.user, gotUser)
		assert.Equal(t, nil, goterr)
	})
}

func TestGetComplexById2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := complexes_mock.NewMockComplexRepo(ctrl)
	usecase := usecase.NewComplexUsecase(mockRepo, &zap.Logger{})
	name := "get complex by id error"
	id := rand.Int63()
	foundCData := &models.ComplexData{}
	t.Run(name, func(t *testing.T) {
		mockRepo.EXPECT().GetComplexById(gomock.Any(), gomock.Any()).Return(foundCData, errors.New("error"))
		_, goterr := usecase.GetComplexById(context.Background(), id)
		// assert.Equal(t, tt.want.user, gotUser)
		assert.Equal(t, errors.New("error"), goterr)
	})
}

/*
func TestCreateFlatAdvert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctrl2 := gomock.NewController(t)
	defer ctrl2.Finish()
	mockTrans := models_mock.NewMockTransaction(ctrl2)
	mockRepo := complexes_mock.NewMockComplexRepo(ctrl)
	usecase := usecase.NewComplexUsecase(mockRepo, &zap.Logger{})
	name := "create flat advert ok"
	AdvData := &models.ComplexAdvertFlatCreateData{}
	t.Run(name, func(t *testing.T) {
		mockRepo.EXPECT().BeginTx(gomock.Any()).Return(mockTrans, nil)
		mockTrans.EXPECT().Commit().Return(nil)
		mockTrans.EXPECT().Rollback().Return(nil)

		mockRepo.EXPECT().CreateFlat(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(123), nil)
		mockRepo.EXPECT().CreateAdvert(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(123), nil)
		mockRepo.EXPECT().CreatePriceChange(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		_, goterr := usecase.CreateFlatAdvert(context.Background(), AdvData)
		// assert.Equal(t, tt.want.user, gotUser)
		assert.Equal(t, nil, goterr)
	})
}

*/
/*
func TestCreateHouseAdvert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctrl2 := gomock.NewController(t)
	defer ctrl2.Finish()
	mockTrans := models_mock.NewMockTransaction(ctrl2)
	mockRepo := complexes_mock.NewMockComplexRepo(ctrl)
	usecase := usecase.NewComplexUsecase(mockRepo, &zap.Logger{})
	name := "create flat advert ok"
	AdvData := &models.ComplexAdvertHouseCreateData{}
	t.Run(name, func(t *testing.T) {
		mockRepo.EXPECT().BeginTx(gomock.Any()).Return(mockTrans, nil)
		mockTrans.EXPECT().Commit().Return(nil)
		mockTrans.EXPECT().Rollback().Return(nil)

		mockRepo.EXPECT().CreateHouse(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(123), nil)
		mockRepo.EXPECT().CreateAdvert(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(123), nil)
		mockRepo.EXPECT().CreatePriceChange(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		_, goterr := usecase.CreateHouseAdvert(context.Background(), AdvData)
		// assert.Equal(t, tt.want.user, gotUser)
		assert.Equal(t, nil, goterr)
	})
}
*/
