package usecase_test

import (
	"2024_1_TeaStealers/internal/models"
	users_mock "2024_1_TeaStealers/internal/pkg/users/mock"
	"2024_1_TeaStealers/internal/pkg/users/usecase"
	"2024_1_TeaStealers/internal/pkg/utils"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := users_mock.NewMockUserRepo(ctrl)
	usecase := usecase.NewUserUsecase(mockRepo)
	id := uuid.NewV4()
	type args struct {
		userUUID uuid.UUID
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
			name: "successful get user",
			args: args{
				userUUID: id,
			},
			want: want{
				user: &models.User{
					ID:           id,
					PasswordHash: "hhhhash",
					LevelUpdate:  1,
					FirstName:    "name1",
					SecondName:   "name2",
					DateBirthday: time.Now(),
					Phone:        "+7115251523",
					Photo:        "/url/to/photo",
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().GetUserById(gomock.Eq(tt.args.userUUID)).Return(tt.want.user, tt.want.err)
			gotUser, goterr := usecase.GetUser(tt.args.userUUID)
			assert.Equal(t, tt.want.user, gotUser)
			assert.Equal(t, tt.want.err, goterr)
		})
	}
}

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
				errUpdatePassword: nil,
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
			if tt.args.errCheckPassword == nil {
				mockRepo.EXPECT().CheckUserPassword(gomock.Eq(tt.args.update.ID), gomock.Eq(utils.GenerateHashString(tt.args.update.OldPassword))).Return(nil)
			}
			if tt.args.errUpdatePassword == nil {
				mockRepo.EXPECT().UpdateUserPassword(gomock.Eq(tt.args.update.ID), gomock.Eq(utils.GenerateHashString(tt.args.update.NewPassword))).Return(1, errors.New("error"))
			}

			gotToken, gotTime, goterr := usecase.UpdateUserPassword(tt.args.update)
			assert.Equal(t, tt.want.token, gotToken)
			assert.True(t, tt.want.expTime.Truncate(time.Second).Equal(gotTime.Truncate(time.Second)))
			assert.Equal(t, tt.want.err, goterr)
		})
	}
}
