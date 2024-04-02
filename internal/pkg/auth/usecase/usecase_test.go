package usecase_test

import (
	"2024_1_TeaStealers/internal/models"
	auth_mock "2024_1_TeaStealers/internal/pkg/auth/mock"
	"2024_1_TeaStealers/internal/pkg/auth/usecase"
	"2024_1_TeaStealers/internal/pkg/utils"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := auth_mock.NewMockAuthRepo(ctrl)
	usecase := usecase.NewAuthUsecase(mockRepo)
	id := uuid.NewV4()
	dat := &models.UserSignUpData{
		Email:    "my@mail.ru",
		Phone:    "+123456",
		Password: "newpassword",
	}
	type args struct {
		data *models.UserSignUpData
		ctx  context.Context
	}
	type want struct {
		user *models.User
		err  error
	}
	//id, email, phone, passwordhash, levelupdate
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful get user",
			args: args{
				data: dat,
				ctx:  context.Background(),
			},
			want: want{
				user: &models.User{
					ID:           id,
					PasswordHash: utils.GenerateHashString(dat.Password),
					LevelUpdate:  1,
					Email:        dat.Email,
					Phone:        dat.Phone,
				},
				err: nil,
			},
		},
		{
			name: "error create user signup",
			args: args{
				data: dat,
				ctx:  context.Background(),
			},
			want: want{
				user: nil,
				err:  errors.New("error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(tt.want.user, tt.want.err)
			gotUser, gotToken, gotTime, goterr := usecase.SignUp(tt.args.ctx, tt.args.data)
			if tt.want.err != nil {
				assert.Empty(t, gotToken)
				assert.Equal(t, tt.want.err, goterr)
				assert.Nil(t, gotUser)
			} else {
				assert.NotEmpty(t, gotToken)
				assert.NotEmpty(t, gotTime)
				assert.Equal(t, tt.want.user.Email, gotUser.Email)
				assert.Equal(t, tt.want.user.Phone, gotUser.Phone)
				assert.Equal(t, tt.want.user.PasswordHash, gotUser.PasswordHash)
				assert.Equal(t, tt.want.user.LevelUpdate, gotUser.LevelUpdate)
				assert.Equal(t, tt.want.err, goterr)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := auth_mock.NewMockAuthRepo(ctrl)
	usecase := usecase.NewAuthUsecase(mockRepo)

	logData := &models.UserLoginData{
		Login:    "+12345",
		Password: "pass",
	}
	id := uuid.NewV4()

	wUser := &models.User{
		ID:           id,
		PasswordHash: utils.GenerateHashString(logData.Password),
		LevelUpdate:  1,
		Email:        "my@mail.ru",
		Phone:        "+12345",
	}
	// id, email, phone, passwordhash, levelupdate
	type args struct {
		data *models.UserLoginData
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
			name: "successful login user",
			args: args{
				data: logData,
			},
			want: want{
				user: wUser,
				err:  nil,
			},
		},
		{
			name: "fail check user",
			args: args{
				data: logData,
			},
			want: want{
				user: wUser,
				err:  errors.New("error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().CheckUser(gomock.Any(), gomock.Eq(tt.args.data.Login), gomock.Eq(utils.GenerateHashString(tt.args.data.Password))).Return(tt.want.user, tt.want.err)
			gotUser, _, _, goterr := usecase.Login(context.Background(), tt.args.data)
			if goterr != nil {
				assert.Nil(t, gotUser)
				assert.Equal(t, tt.want.err, goterr)
			} else {
				assert.Equal(t, tt.want.user, gotUser)
				assert.Equal(t, tt.want.err, goterr)
			}
		})
	}
}