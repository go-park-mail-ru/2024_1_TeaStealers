package usecase_test

import (
	"2024_1_TeaStealers/internal/models"
	transaction "2024_1_TeaStealers/internal/models/mock"
	complexes_mock "2024_1_TeaStealers/internal/pkg/complexes/mock"
	"2024_1_TeaStealers/internal/pkg/complexes/usecase"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCreateComplex(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	newComplex := &models.Complex{
		ID:                     uuid.NewV4(),
		CompanyId:              uuid.NewV4(),
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
		CompanyId:              uuid.NewV4(),
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
		Address: "adr",
	}

	createBuild := &models.Building{
		Address: "adr",
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
	foundCData := &models.ComplexData{}
	t.Run(name, func(t *testing.T) {
		mockRepo.EXPECT().GetComplexById(gomock.Any(), gomock.Any()).Return(foundCData, nil)
		_, goterr := usecase.GetComplexById(context.Background(), uuid.NewV4())
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
	foundCData := &models.ComplexData{}
	t.Run(name, func(t *testing.T) {
		mockRepo.EXPECT().GetComplexById(gomock.Any(), gomock.Any()).Return(foundCData, errors.New("error"))
		_, goterr := usecase.GetComplexById(context.Background(), uuid.NewV4())
		// assert.Equal(t, tt.want.user, gotUser)
		assert.Equal(t, errors.New("error"), goterr)
	})
}

func TestCreateFlatAdvert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctrl2 := gomock.NewController(t)
	defer ctrl2.Finish()
	mockTrans := transaction.NewMockTransaction(ctrl2)
	mockRepo := complexes_mock.NewMockComplexRepo(ctrl)
	usecase := usecase.NewComplexUsecase(mockRepo, &zap.Logger{})
	name := "create flat advert ok"
	AdvData := &models.ComplexAdvertFlatCreateData{}
	t.Run(name, func(t *testing.T) {
		mockRepo.EXPECT().BeginTx(gomock.Any()).Return(mockTrans, nil)
		mockTrans.EXPECT().Commit().Return(nil)
		mockTrans.EXPECT().Rollback().Return(nil)

		mockRepo.EXPECT().CreateAdvertType(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		mockRepo.EXPECT().CreateFlat(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		mockRepo.EXPECT().CreateAdvert(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		mockRepo.EXPECT().CreatePriceChange(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		_, goterr := usecase.CreateFlatAdvert(context.Background(), AdvData)
		// assert.Equal(t, tt.want.user, gotUser)
		assert.Equal(t, nil, goterr)
	})
}

func TestCreateHouseAdvert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctrl2 := gomock.NewController(t)
	defer ctrl2.Finish()
	mockTrans := transaction.NewMockTransaction(ctrl2)
	mockRepo := complexes_mock.NewMockComplexRepo(ctrl)
	usecase := usecase.NewComplexUsecase(mockRepo, &zap.Logger{})
	name := "create flat advert ok"
	AdvData := &models.ComplexAdvertHouseCreateData{}
	t.Run(name, func(t *testing.T) {
		mockRepo.EXPECT().BeginTx(gomock.Any()).Return(mockTrans, nil)
		mockTrans.EXPECT().Commit().Return(nil)
		mockTrans.EXPECT().Rollback().Return(nil)

		mockRepo.EXPECT().CreateAdvertType(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		mockRepo.EXPECT().CreateHouse(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		mockRepo.EXPECT().CreateAdvert(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		mockRepo.EXPECT().CreatePriceChange(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		_, goterr := usecase.CreateHouseAdvert(context.Background(), AdvData)
		// assert.Equal(t, tt.want.user, gotUser)
		assert.Equal(t, nil, goterr)
	})
}

/*
func TestUpdateUserInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := users_mock.NewMockUserRepo(ctrl)
	usecase := usecase.NewUserUsecase(mockRepo)
	id := uuid.NewV4()
	type args struct {
		userUUID uuid.UUID
		data     *models.UserUpdateData
	}
	type want struct {
		user *models.User
		err  error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful update user",
			args: args{
				userUUID: id,
				data: &models.UserUpdateData{
					FirstName:    "newname1",
					SecondName:   "newname2",
					DateBirthday: time.Now(),
					Phone:        "+712345678",
					Email:        "new@mail.ru",
				},
			},
			want: want{
				user: &models.User{
					ID:           id,
					PasswordHash: "hhhhash",
					LevelUpdate:  1,
					FirstName:    "newname1",
					SecondName:   "newname2",
					DateBirthday: time.Now(),
					Email:        "new@mail.ru",
					Phone:        "+712345678",
					//Photo:        "/url/to/photo",
				},
				err: nil,
			},
		},
		{
			name: "empty phone user",
			args: args{
				userUUID: id,
				data: &models.UserUpdateData{
					FirstName:    "newname1",
					SecondName:   "newname2",
					DateBirthday: time.Now(),
					Phone:        "",
					Email:        "new@mail.ru",
				},
			},
			want: want{
				user: nil,
				err:  errors.New("phone cannot be empty"),
			},
		},
		{
			name: "empty email user",
			args: args{
				userUUID: id,
				data: &models.UserUpdateData{
					FirstName:    "newname1",
					SecondName:   "newname2",
					DateBirthday: time.Now(),
					Phone:        "+712345678",
					Email:        "",
				},
			},
			want: want{
				user: nil,
				err:  errors.New("email cannot be empty"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "successful update user" {
				mockRepo.EXPECT().UpdateUserInfo(gomock.Eq(tt.args.userUUID), gomock.Eq(tt.args.data)).Return(tt.want.user, tt.want.err)
			}
			gotUser, goterr := usecase.UpdateUserInfo(tt.args.userUUID, tt.args.data)
			assert.Equal(t, tt.want.user, gotUser)
			assert.Equal(t, tt.want.err, goterr)
		})
	}
}

func TestUpdateUserPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := users_mock.NewMockUserRepo(ctrl)
	usecase := usecase.NewUserUsecase(mockRepo)
	id := uuid.NewV4()
	type args struct {
		update            *models.UserUpdatePassword
		errCheckPassword  error
		errUpdatePassword error
		UpdatePassword    bool
		CheckPassword     bool
	}
	type want struct {
		token   string
		expTime time.Time
		err     error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "pass must not match",
			args: args{
				update: &models.UserUpdatePassword{
					ID:          id,
					OldPassword: "oldpassword",
					NewPassword: "oldpassword",
				},
				errCheckPassword:  errors.New("error"),
				errUpdatePassword: errors.New("error"),
				UpdatePassword:    false,
				CheckPassword:     false,
			},
			want: want{
				token:   "",
				expTime: time.Now(),
				err:     errors.New("passwords must not match"),
			},
		},
		{
			name: "incorrect id or passwordhash",
			args: args{
				update: &models.UserUpdatePassword{
					ID:          id,
					OldPassword: "oldpassword",
					NewPassword: "newpassword",
				},
				errCheckPassword:  nil,
				errUpdatePassword: errors.New("error"),
				UpdatePassword:    true,
				CheckPassword:     true,
			},
			want: want{
				token:   "",
				expTime: time.Now(),
				err:     errors.New("incorrect id or passwordhash"),
			},
		},

		{
			name: "ok changePassword",
			args: args{
				update: &models.UserUpdatePassword{
					ID:          id,
					OldPassword: "oldpassword",
					NewPassword: "newpassword",
				},
				errCheckPassword:  nil,
				errUpdatePassword: nil,
				UpdatePassword:    true,
				CheckPassword:     true,
			},
			want: want{
				token:   "",
				expTime: time.Now(),
				err:     nil,
			},
		},

		{
			name: "invalid old password",
			args: args{
				update: &models.UserUpdatePassword{
					ID:          id,
					OldPassword: "oldpassword",
					NewPassword: "newpassword",
				},
				errCheckPassword:  nil,
				errUpdatePassword: errors.New("error"),
				UpdatePassword:    true,
				CheckPassword:     true,
			},
			want: want{
				token:   "",
				expTime: time.Now(),
				err:     errors.New("incorrect id or passwordhash"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.CheckPassword {
				mockRepo.EXPECT().CheckUserPassword(gomock.Eq(tt.args.update.ID), gomock.Eq(utils.GenerateHashString(tt.args.update.OldPassword))).Return(tt.args.errCheckPassword)
			}
			if tt.args.UpdatePassword {
				mockRepo.EXPECT().UpdateUserPassword(gomock.Eq(tt.args.update.ID), gomock.Eq(utils.GenerateHashString(tt.args.update.NewPassword))).Return(1, tt.args.errUpdatePassword)
			}

			_, _, goterr := usecase.UpdateUserPassword(tt.args.update)
			assert.Equal(t, tt.want.err, goterr)
		})
	}
}
*/
