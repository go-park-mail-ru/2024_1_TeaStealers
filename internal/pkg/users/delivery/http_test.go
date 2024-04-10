package delivery

/*import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/middleware"
	mock "2024_1_TeaStealers/internal/pkg/users/mock"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUsersHandler_GetCurUser(t *testing.T) {
	type fields struct {
		usecase *mock.MockUserUsecase
	}
	type args struct {
		cookieId uuid.UUID
		//	userinfoErr error
	}
	type want struct {
		user    *models.User
		status  int
		err     error
		message string
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id1 := uuid.NewV4()
	id2 := uuid.NewV4()
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      want
		IsMessage bool
		prepare   func(f *fields, a *args, w *want) *httptest.ResponseRecorder
	}{
		{
			name: "success got user",
			fields: fields{
				usecase: mock.NewMockUserUsecase(ctrl),
			},
			args: args{
				cookieId: id1,
			},
			want: want{
				user: &models.User{
					ID:           id2,
					FirstName:    "Maksim",
					SecondName:   "Shagaev",
					Phone:        "+79003249325",
					Email:        "my@mail.ru",
					Photo:        "/path/to/photo/test.jpg",
					DateBirthday: time.Now(),
				},
				status: http.StatusOK,
				err:    nil,
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().GetUser(gomock.Any(), gomock.Eq(a.cookieId)).Return(w.user, w.err)
				handler := NewUserHandler(f.usecase)

				req := httptest.NewRequest(http.MethodGet, "/me", nil)
				req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, a.cookieId))
				rec := httptest.NewRecorder()
				handler.GetCurUser(rec, req)
				return rec

			},
		},
		{
			name:      "user does not exists",
			IsMessage: true,
			fields: fields{
				usecase: mock.NewMockUserUsecase(ctrl),
			},
			args: args{
				cookieId: id1,
			},
			want: want{
				user: &models.User{
					ID:           id2,
					FirstName:    "Maksim",
					SecondName:   "Shagaev",
					Phone:        "+79003249325",
					Email:        "my@mail.ru",
					Photo:        "/path/to/photo/test.jpg",
					DateBirthday: time.Now(),
				},
				message: "user is not exists",
				status:  http.StatusBadRequest,
				err:     errors.New("error"),
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().GetUser(gomock.Any(), gomock.Eq(a.cookieId)).Return(w.user, w.err)
				handler := NewUserHandler(f.usecase)

				req := httptest.NewRequest(http.MethodGet, "/me", nil)
				req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, a.cookieId))
				rec := httptest.NewRecorder()
				handler.GetCurUser(rec, req)
				return rec

			},
		},
		{
			name:      "incorrect id",
			IsMessage: true,
			fields: fields{
				usecase: mock.NewMockUserUsecase(ctrl),
			},
			args: args{
				cookieId: id1,
			},
			want: want{
				user: &models.User{
					ID:           id2,
					FirstName:    "Maksim",
					SecondName:   "Shagaev",
					Phone:        "+79003249325",
					Email:        "my@mail.ru",
					Photo:        "/path/to/photo/test.jpg",
					DateBirthday: time.Now(),
				},
				message: "incorrect id",
				status:  http.StatusBadRequest,
				err:     errors.New("error"),
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				// f.usecase.EXPECT().GetUser(gomock.Eq(a.cookieId)).Return(w.user, w.err)
				handler := NewUserHandler(f.usecase)

				req := httptest.NewRequest(http.MethodGet, "/me", nil)
				req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, "notid"))
				rec := httptest.NewRecorder()
				handler.GetCurUser(rec, req)
				return rec

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := tt.prepare(&tt.fields, &tt.args, &tt.want)
			response := rec.Result()
			assert.Equal(t, tt.want.status, response.StatusCode)
			bodyBytes, _ := io.ReadAll(response.Body)
			var responseMap map[string]interface{}
			err := json.Unmarshal(bodyBytes, &responseMap)
			if tt.IsMessage {
				assert.Equal(t, tt.want.message, responseMap["message"])
			} else {
				payload := responseMap["payload"].(map[string]interface{})
				if err != nil {
					fmt.Println("Error unmarshalling response:", err)
				}
				assert.Equal(t, tt.want.user.ID.String(), payload["id"])
				assert.Equal(t, tt.want.user.FirstName, payload["firstName"])
				assert.Equal(t, tt.want.user.SecondName, payload["secondName"])

				wantDate := tt.want.user.DateBirthday.Truncate(time.Second)
				gotDate, err := time.Parse(time.RFC3339, payload["dateBirthday"].(string))
				if err != nil {
					t.Fatalf("failed to parse date: %v", err)
				}
				gotDate = gotDate.Truncate(time.Second)

				assert.True(t, wantDate == gotDate)
				assert.Equal(t, tt.want.user.Phone, payload["phone"])
				assert.Equal(t, tt.want.user.Email, payload["email"])
				assert.Equal(t, tt.want.user.Photo, payload["photo"])
				code, _ := responseMap["statusCode"].(float64)
				assert.Equal(t, tt.want.status, int(code))
			}
		})
	}
}

func TestUsersHandler_UpdateUserPhoto(t *testing.T) {
	type fields struct {
		usecase *mock.MockUserUsecase
	}
	type args struct {
		cookieId      uuid.UUID
		updateUseCerr error
		fileBytes     []byte
		fileType      string
		fileName      string
	}
	type want struct {
		filepath string
		status   int
		// err      error
		message string
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id1 := uuid.NewV4()
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      want
		IsMessage bool
		prepare   func(f *fields, a *args, w *want) *httptest.ResponseRecorder
	}{
		{
			name: "success update user photo",
			fields: fields{
				usecase: mock.NewMockUserUsecase(ctrl),
			},
			args: args{
				cookieId:      id1,
				fileBytes:     []byte("test file content"),
				fileType:      ".jpg",
				fileName:      "test_file",
				updateUseCerr: nil,
			},
			want: want{
				filepath: "test_file.jpg",
				status:   http.StatusOK,
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().UpdateUserPhoto(gomock.Any(), gomock.Any(), gomock.Eq(a.fileType), gomock.Eq(a.cookieId)).Return(w.filepath, a.updateUseCerr)
				handler := NewUserHandler(f.usecase)
				fakeFormData := new(bytes.Buffer)
				fakeWriter := multipart.NewWriter(fakeFormData)
				fakePart, _ := fakeWriter.CreateFormFile("file", a.fileName+a.fileType)
				if _, err := fakePart.Write(a.fileBytes); err != nil {
					return nil
				}
				fakeWriter.Close()

				request := httptest.NewRequest(http.MethodPost, "/me", fakeFormData)
				request.Header.Set("Content-Type", fakeWriter.FormDataContentType())
				request = request.WithContext(context.WithValue(request.Context(), middleware.CookieName, a.cookieId))
				recorder := httptest.NewRecorder()
				handler.UpdateUserPhoto(recorder, request)
				return recorder
			},
			IsMessage: false,
		},
		{
			name: "not allowed extension photo",
			fields: fields{
				usecase: mock.NewMockUserUsecase(ctrl),
			},
			args: args{
				cookieId:      id1,
				fileBytes:     []byte("test file content"),
				fileType:      ".txt",
				fileName:      "test_file",
				updateUseCerr: nil,
			},
			want: want{
				filepath: "test_file.txt",
				status:   http.StatusBadRequest,
				message:  "jpg, jpeg, png only",
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				// f.usecase.EXPECT().UpdateUserPhoto(gomock.Any(), gomock.Eq(a.fileType), gomock.Eq(a.cookieId)).Return(w.filepath, a.updateUseCerr)
				handler := NewUserHandler(f.usecase)
				fakeFormData := new(bytes.Buffer)
				fakeWriter := multipart.NewWriter(fakeFormData)
				fakePart, _ := fakeWriter.CreateFormFile("file", a.fileName+a.fileType)
				if _, err := fakePart.Write(a.fileBytes); err != nil {
					return nil
				}
				fakeWriter.Close()

				request := httptest.NewRequest(http.MethodPost, "/avatar", fakeFormData)
				request.Header.Set("Content-Type", fakeWriter.FormDataContentType())
				request = request.WithContext(context.WithValue(request.Context(), middleware.CookieName, a.cookieId))
				recorder := httptest.NewRecorder()
				handler.UpdateUserPhoto(recorder, request)
				return recorder
			},
			IsMessage: true,
		},
		{
			name: "bad data request",
			fields: fields{
				usecase: mock.NewMockUserUsecase(ctrl),
			},
			args: args{
				cookieId:      id1,
				fileBytes:     []byte("test file content"),
				fileType:      ".jpg",
				fileName:      "test_file",
				updateUseCerr: errors.New("error"),
			},
			want: want{
				filepath: "test_file.jpg",
				status:   http.StatusBadRequest,
				message:  "bad data request",
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				// f.usecase.EXPECT().UpdateUserPhoto(gomock.Any(), gomock.Eq(a.fileType), gomock.Eq(a.cookieId)).Return(w.filepath, a.updateUseCerr)
				handler := NewUserHandler(f.usecase)
				fakeFormData := new(bytes.Buffer)
				fakeWriter := multipart.NewWriter(fakeFormData)
				fakePart, _ := fakeWriter.CreateFormFile("notfile", a.fileName+a.fileType)
				if _, err := fakePart.Write(a.fileBytes); err != nil {
					return nil
				}
				fakeWriter.Close()

				request := httptest.NewRequest(http.MethodPost, "/avatar", fakeFormData)
				request.Header.Set("Content-Type", fakeWriter.FormDataContentType())
				request = request.WithContext(context.WithValue(request.Context(), middleware.CookieName, a.cookieId))
				recorder := httptest.NewRecorder()
				handler.UpdateUserPhoto(recorder, request)
				return recorder
			},
			IsMessage: true,
		},
		{
			name: "failed update photo",
			fields: fields{
				usecase: mock.NewMockUserUsecase(ctrl),
			},
			args: args{
				cookieId:      id1,
				fileBytes:     []byte("test file content"),
				fileType:      ".jpg",
				fileName:      "test_file",
				updateUseCerr: errors.New("error"),
			},
			want: want{
				filepath: "test_file.jpg",
				status:   http.StatusBadRequest,
				message:  "failed upload file",
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().UpdateUserPhoto(gomock.Any(), gomock.Any(), gomock.Eq(a.fileType), gomock.Eq(a.cookieId)).Return(w.filepath, a.updateUseCerr)
				handler := NewUserHandler(f.usecase)
				fakeFormData := new(bytes.Buffer)
				fakeWriter := multipart.NewWriter(fakeFormData)
				fakePart, _ := fakeWriter.CreateFormFile("file", a.fileName+a.fileType)
				if _, err := fakePart.Write(a.fileBytes); err != nil {
					return nil
				}
				fakeWriter.Close()

				request := httptest.NewRequest(http.MethodPost, "/avatar", fakeFormData)
				request.Header.Set("Content-Type", fakeWriter.FormDataContentType())
				request = request.WithContext(context.WithValue(request.Context(), middleware.CookieName, a.cookieId))
				recorder := httptest.NewRecorder()
				handler.UpdateUserPhoto(recorder, request)
				return recorder
			},
			IsMessage: true,
		},
		{
			name: "incorrect id photo",
			fields: fields{
				usecase: mock.NewMockUserUsecase(ctrl),
			},
			args: args{
				cookieId:      id1,
				fileBytes:     []byte("test file content"),
				fileType:      ".jpg",
				fileName:      "test_file",
				updateUseCerr: nil,
			},
			want: want{
				filepath: "test_file.jpg",
				status:   http.StatusBadRequest,
				message:  "incorrect id",
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				// f.usecase.EXPECT().UpdateUserPhoto(gomock.Any(), gomock.Eq(a.fileType), gomock.Eq(a.cookieId)).Return(w.filepath, a.updateUseCerr)
				handler := NewUserHandler(f.usecase)
				fakeFormData := new(bytes.Buffer)
				fakeWriter := multipart.NewWriter(fakeFormData)
				fakePart, _ := fakeWriter.CreateFormFile("file", a.fileName+a.fileType)
				if _, err := fakePart.Write(a.fileBytes); err != nil {
					return nil
				}
				fakeWriter.Close()

				request := httptest.NewRequest(http.MethodPost, "/avatar", fakeFormData)
				request.Header.Set("Content-Type", fakeWriter.FormDataContentType())
				request = request.WithContext(context.WithValue(request.Context(), middleware.CookieName, "not id"))
				recorder := httptest.NewRecorder()
				handler.UpdateUserPhoto(recorder, request)
				return recorder
			},
			IsMessage: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := tt.prepare(&tt.fields, &tt.args, &tt.want)
			response := rec.Result()
			assert.Equal(t, tt.want.status, response.StatusCode)
			bodyBytes, _ := io.ReadAll(response.Body)
			var responseMap map[string]interface{}
			err := json.Unmarshal(bodyBytes, &responseMap)
			if err != nil {
				fmt.Println("Error unmarshalling response:", err)
			}
			if tt.IsMessage {
				assert.Equal(t, tt.want.message, responseMap["message"])
			} else {
				filename := responseMap["payload"].(string)

				assert.Equal(t, tt.want.filepath, filename)
				code := responseMap["statusCode"].(float64)
				assert.Equal(t, tt.want.status, int(code))
			}
		})
	}
}

func TestUsersHandler_DeleteUserPhoto(t *testing.T) {
	type fields struct {
		usecase *mock.MockUserUsecase
	}
	type args struct {
		cookieId      uuid.UUID
		userDeleteErr error
	}
	type want struct {
		payload string
		status  int
		// err     error
		message string
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id1 := uuid.NewV4()
	// id2 := uuid.NewV4()
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      want
		IsMessage bool
		prepare   func(f *fields, a *args, w *want) *httptest.ResponseRecorder
	}{
		{
			name: "success delete avatar",
			fields: fields{
				usecase: mock.NewMockUserUsecase(ctrl),
			},
			args: args{
				cookieId:      id1,
				userDeleteErr: nil,
			},
			want: want{
				payload: "success delete avatar",
				message: "",
				status:  http.StatusOK,
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().DeleteUserPhoto(gomock.Any(), gomock.Eq(a.cookieId)).Return(a.userDeleteErr)
				handler := NewUserHandler(f.usecase)
				req := httptest.NewRequest(http.MethodDelete, "/avatar", nil)
				req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, a.cookieId))
				rec := httptest.NewRecorder()
				handler.DeleteUserPhoto(rec, req)
				return rec
			},
			IsMessage: false,
		},
		{
			name: "error delete photo",
			fields: fields{
				usecase: mock.NewMockUserUsecase(ctrl),
			},
			args: args{
				cookieId:      id1,
				userDeleteErr: errors.New("error"),
			},
			want: want{
				payload: "",
				message: "error delete avatar",
				status:  http.StatusBadRequest,
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().DeleteUserPhoto(gomock.Any(), gomock.Eq(a.cookieId)).Return(a.userDeleteErr)
				handler := NewUserHandler(f.usecase)
				req := httptest.NewRequest(http.MethodDelete, "/avatar", nil)
				req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, a.cookieId))
				rec := httptest.NewRecorder()
				handler.DeleteUserPhoto(rec, req)
				return rec
			},
			IsMessage: true,
		},
		{
			name: "incorrect id delete photo",
			fields: fields{
				usecase: mock.NewMockUserUsecase(ctrl),
			},
			args: args{
				cookieId:      id1,
				userDeleteErr: nil,
			},
			want: want{
				payload: "",
				message: "incorrect id",
				status:  http.StatusBadRequest,
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				// f.usecase.EXPECT().DeleteUserPhoto(gomock.Eq(a.cookieId)).Return(w.err)
				handler := NewUserHandler(f.usecase)
				req := httptest.NewRequest(http.MethodDelete, "/avatar", nil)
				req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, "not id"))
				rec := httptest.NewRecorder()
				handler.DeleteUserPhoto(rec, req)
				return rec
			},
			IsMessage: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := tt.prepare(&tt.fields, &tt.args, &tt.want)
			response := rec.Result()
			assert.Equal(t, tt.want.status, response.StatusCode)
			bodyBytes, _ := io.ReadAll(response.Body)
			var responseMap map[string]interface{}
			err := json.Unmarshal(bodyBytes, &responseMap)
			if tt.IsMessage {
				assert.Equal(t, tt.want.message, responseMap["message"])
			} else {
				payload := responseMap["payload"].(string)
				if err != nil {
					fmt.Println("Error unmarshalling response:", err)
				}
				assert.Equal(t, tt.want.payload, payload)
				code, _ := responseMap["statusCode"].(float64)
				assert.Equal(t, tt.want.status, int(code))
			}
		})
	}
}

func TestUsersHandler_UpdateUserInfo(t *testing.T) {
	type fields struct {
		usecase *mock.MockUserUsecase
	}
	type args struct {
		cookieId      uuid.UUID
		userUpdData   *models.UserUpdateData
		userUpdateErr error
	}
	type want struct {
		wantUser *models.User
		payload  string
		status   int
		message  string
	}
	timeB := time.Now().Truncate(time.Hour)
	updData := &models.UserUpdateData{
		FirstName:    "Maksim",
		SecondName:   "Shagaev",
		DateBirthday: timeB,
		Phone:        "+123456",
		Email:        "my@mail.ru",
	}
	user := &models.User{
		ID:           uuid.NewV4(),
		PasswordHash: "hash",
		LevelUpdate:  1,
		FirstName:    "Maksim",
		SecondName:   "Shagaev",
		DateBirthday: timeB,
		Phone:        "+123456",
		Email:        "my@mail.ru",
		Photo:        "path/to/photo",
		DateCreation: time.Now(),
		IsDeleted:    false,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id1 := uuid.NewV4()
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      want
		IsMessage bool
		prepare   func(f *fields, a *args, w *want) *httptest.ResponseRecorder
	}{
		{
			name: "success update info",
			fields: fields{
				usecase: mock.NewMockUserUsecase(ctrl),
			},
			args: args{
				cookieId:      id1,
				userUpdData:   updData,
				userUpdateErr: nil,
			},
			want: want{
				wantUser: user,
				payload:  "",
				message:  "",
				status:   http.StatusOK,
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().UpdateUserInfo(gomock.Any(), gomock.Eq(a.cookieId), gomock.Eq(a.userUpdData)).Return(w.wantUser, a.userUpdateErr)
				handler := NewUserHandler(f.usecase)
				reqBody, _ := json.Marshal(a.userUpdData)
				req := httptest.NewRequest(http.MethodPost, "/info", bytes.NewBuffer(reqBody))
				req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, a.cookieId))
				rec := httptest.NewRecorder()
				handler.UpdateUserInfo(rec, req)
				return rec
			},
			IsMessage: false,
		},
		{
			name: "incorrect id info",
			fields: fields{
				usecase: mock.NewMockUserUsecase(ctrl),
			},
			args: args{
				cookieId:      id1,
				userUpdData:   updData,
				userUpdateErr: nil,
			},
			want: want{
				wantUser: user,
				payload:  "",
				message:  "incorrect id",
				status:   http.StatusBadRequest,
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				// f.usecase.EXPECT().UpdateUserInfo(gomock.Eq(a.cookieId), gomock.Eq(a.userUpdData)).Return(w.wantUser, a.userUpdateErr)
				handler := NewUserHandler(f.usecase)
				reqBody, _ := json.Marshal(a.userUpdData)
				req := httptest.NewRequest(http.MethodPost, "/info", bytes.NewBuffer(reqBody))
				req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, "notid"))
				rec := httptest.NewRecorder()
				handler.UpdateUserInfo(rec, req)
				return rec
			},
			IsMessage: true,
		},
		{
			name: "incorrect data format info",
			fields: fields{
				usecase: mock.NewMockUserUsecase(ctrl),
			},
			args: args{
				cookieId:      id1,
				userUpdData:   updData,
				userUpdateErr: nil,
			},
			want: want{
				wantUser: user,
				payload:  "",
				message:  "incorrect data format",
				status:   http.StatusBadRequest,
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				// f.usecase.EXPECT().UpdateUserInfo(gomock.Eq(a.cookieId), gomock.Eq(a.userUpdData)).Return(w.wantUser, a.userUpdateErr)
				handler := NewUserHandler(f.usecase)
				req := httptest.NewRequest(http.MethodPost, "/info", nil)
				req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, a.cookieId))
				rec := httptest.NewRecorder()
				handler.UpdateUserInfo(rec, req)
				return rec
			},
			IsMessage: true,
		},
		{
			name: "error update info",
			fields: fields{
				usecase: mock.NewMockUserUsecase(ctrl),
			},
			args: args{
				cookieId:      id1,
				userUpdData:   updData,
				userUpdateErr: errors.New("email cannot be empty"),
			},
			want: want{
				wantUser: user,
				payload:  "",
				message:  "email cannot be empty",
				status:   http.StatusBadRequest,
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().UpdateUserInfo(gomock.Any(), gomock.Eq(a.cookieId), gomock.Eq(a.userUpdData)).Return(w.wantUser, a.userUpdateErr)
				handler := NewUserHandler(f.usecase)
				reqBody, _ := json.Marshal(a.userUpdData)
				req := httptest.NewRequest(http.MethodPost, "/info", bytes.NewBuffer(reqBody))
				req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, a.cookieId))
				rec := httptest.NewRecorder()
				handler.UpdateUserInfo(rec, req)
				return rec
			},
			IsMessage: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := tt.prepare(&tt.fields, &tt.args, &tt.want)
			response := rec.Result()
			assert.Equal(t, tt.want.status, response.StatusCode)
			bodyBytes, _ := io.ReadAll(response.Body)
			var responseMap map[string]interface{}
			err := json.Unmarshal(bodyBytes, &responseMap)
			if tt.IsMessage {
				assert.Equal(t, tt.want.message, responseMap["message"])
			} else {
				payload := responseMap["payload"].(map[string]interface{})
				if err != nil {
					fmt.Println("Error unmarshalling response:", err)
				}
				assert.Equal(t, tt.want.wantUser.ID.String(), payload["id"])
				assert.Equal(t, tt.want.wantUser.FirstName, payload["firstName"])
				assert.Equal(t, tt.want.wantUser.SecondName, payload["secondName"])
				assert.Equal(t, tt.want.wantUser.Email, payload["email"])
				assert.Equal(t, tt.want.wantUser.Phone, payload["phone"])
				assert.Equal(t, tt.want.wantUser.Photo, payload["photo"])
				code, _ := responseMap["statusCode"].(float64)
				assert.Equal(t, tt.want.status, int(code))
			}
		})
	}
}

func TestUsersHandler_UpdateUserPassword(t *testing.T) {
	type fields struct {
		usecase *mock.MockUserUsecase
	}
	type args struct {
		cookieId      uuid.UUID
		userUpdData   *models.UserUpdatePassword
		token         string
		userUpdateErr error
	}
	type want struct {
		payload string
		status  int
		message string
	}
	updData := &models.UserUpdatePassword{
		OldPassword: "oldpass",
		NewPassword: "newpass",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id1 := uuid.NewV4()
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      want
		IsMessage bool
		prepare   func(f *fields, a *args, w *want) *httptest.ResponseRecorder
	}{
		{
			name: "success update password",
			fields: fields{
				usecase: mock.NewMockUserUsecase(ctrl),
			},
			args: args{
				cookieId:      id1,
				userUpdData:   updData,
				userUpdateErr: nil,
				token:         "token",
			},
			want: want{
				payload: "success update password",
				message: "",
				status:  http.StatusOK,
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().UpdateUserPassword(gomock.Any(), gomock.Eq(a.userUpdData)).Return(a.token, time.Now(), a.userUpdateErr)
				handler := NewUserHandler(f.usecase)
				reqBody, _ := json.Marshal(a.userUpdData)
				req := httptest.NewRequest(http.MethodPost, "/password", bytes.NewBuffer(reqBody))
				req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, a.cookieId))
				rec := httptest.NewRecorder()
				handler.UpdateUserPassword(rec, req)
				return rec
			},
			IsMessage: false,
		},
		{
			name: "incorrect id password",
			fields: fields{
				usecase: mock.NewMockUserUsecase(ctrl),
			},
			args: args{
				cookieId:      id1,
				userUpdData:   updData,
				userUpdateErr: nil,
				token:         "token",
			},
			want: want{
				payload: "",
				message: "incorrect id",
				status:  http.StatusBadRequest,
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				// f.usecase.EXPECT().UpdateUserPassword(gomock.Eq(a.userUpdData)).Return(a.token, time.Now(), a.userUpdateErr)
				handler := NewUserHandler(f.usecase)
				reqBody, _ := json.Marshal(a.userUpdData)
				req := httptest.NewRequest(http.MethodPost, "/password", bytes.NewBuffer(reqBody))
				req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, "notid"))
				rec := httptest.NewRecorder()
				handler.UpdateUserPassword(rec, req)
				return rec
			},
			IsMessage: true,
		},
		{
			name: "Incorrect data format password",
			fields: fields{
				usecase: mock.NewMockUserUsecase(ctrl),
			},
			args: args{
				cookieId:      id1,
				userUpdData:   updData,
				userUpdateErr: nil,
				token:         "token",
			},
			want: want{
				payload: "",
				message: "incorrect data format",
				status:  http.StatusBadRequest,
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				// f.usecase.EXPECT().UpdateUserPassword(gomock.Eq(a.userUpdData)).Return(a.token, time.Now(), a.userUpdateErr)
				handler := NewUserHandler(f.usecase)
				req := httptest.NewRequest(http.MethodPost, "/password", nil)
				req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, a.cookieId))
				rec := httptest.NewRecorder()
				handler.UpdateUserPassword(rec, req)
				return rec
			},
			IsMessage: true,
		},
		{
			name: "error update password",
			fields: fields{
				usecase: mock.NewMockUserUsecase(ctrl),
			},
			args: args{
				cookieId:      id1,
				userUpdData:   updData,
				userUpdateErr: errors.New("some error"),
				token:         "token",
			},
			want: want{
				payload: "",
				message: "some error",
				status:  http.StatusBadRequest,
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().UpdateUserPassword(gomock.Any(), gomock.Eq(a.userUpdData)).Return(a.token, time.Now(), a.userUpdateErr)
				handler := NewUserHandler(f.usecase)
				reqBody, _ := json.Marshal(a.userUpdData)
				req := httptest.NewRequest(http.MethodPost, "/password", bytes.NewBuffer(reqBody))
				req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, a.cookieId))
				rec := httptest.NewRecorder()
				handler.UpdateUserPassword(rec, req)
				return rec
			},
			IsMessage: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := tt.prepare(&tt.fields, &tt.args, &tt.want)
			response := rec.Result()
			assert.Equal(t, tt.want.status, response.StatusCode)
			bodyBytes, _ := io.ReadAll(response.Body)
			var responseMap map[string]interface{}
			err := json.Unmarshal(bodyBytes, &responseMap)
			if tt.IsMessage {
				assert.Equal(t, tt.want.message, responseMap["message"])
			} else {
				payload := responseMap["payload"].(string)
				if err != nil {
					fmt.Println("Error unmarshalling response:", err)
				}
				assert.Equal(t, tt.want.payload, payload)
				code, _ := responseMap["statusCode"].(float64)
				assert.Equal(t, tt.want.status, int(code))
			}
		})
	}
}*/
