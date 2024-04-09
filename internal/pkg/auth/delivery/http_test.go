package delivery

import (
	"2024_1_TeaStealers/internal/models"
	mock "2024_1_TeaStealers/internal/pkg/auth/mock"
	"2024_1_TeaStealers/internal/pkg/middleware"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthHandler_SignUp(t *testing.T) {
	type fields struct {
		usecase *mock.MockAuthUsecase
	}
	type args struct {
		data *models.UserSignUpData
	}
	type want struct {
		user    *models.User
		status  int
		err     error
		message string
		token   string
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
			name: "success signup",
			fields: fields{
				usecase: mock.NewMockAuthUsecase(ctrl),
			},
			args: args{
				data: &models.UserSignUpData{
					Email:    "my@mail.ru",
					Phone:    "+123456",
					Password: "pass",
				},
			},
			want: want{
				user: &models.User{
					ID: id1,
					//FirstName:    "Maksim",
					//SecondName:   "Shagaev",
					Phone:        "+79003249325",
					Email:        "my@mail.ru",
					PasswordHash: "hash",
					//Photo:        "/path/to/photo/test.jpg",
					//DateBirthday: time.Now(),
				},
				message: "",
				token:   "token",
				status:  http.StatusCreated,
				err:     nil,
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().SignUp(gomock.Any(), gomock.Eq(a.data)).Return(w.user, w.token, time.Now(), w.err)
				handler := NewAuthHandler(f.usecase)
				reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
				// req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, a.cookieId))
				rec := httptest.NewRecorder()
				handler.SignUp(rec, req)
				return rec
			},
		},
		{
			name: "incorrect data format",
			fields: fields{
				usecase: mock.NewMockAuthUsecase(ctrl),
			},
			args: args{
				data: &models.UserSignUpData{
					Email:    "my@mail.ru",
					Phone:    "+123456",
					Password: "pass",
				},
			},
			want: want{
				user: &models.User{
					ID: id1,
					// FirstName:    "Maksim",
					// SecondName:   "Shagaev",
					Phone:        "+79003249325",
					Email:        "my@mail.ru",
					PasswordHash: "hash",
					// Photo:        "/path/to/photo/test.jpg",
					// DateBirthday: time.Now(),
				},
				message: "incorrect data format",
				token:   "token",
				status:  http.StatusBadRequest,
				err:     nil,
			},
			IsMessage: true,
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				// f.usecase.EXPECT().SignUp(gomock.Any(), gomock.Eq(a.data)).Return(w.user, w.token, time.Now(), w.err)
				handler := NewAuthHandler(f.usecase)
				// reqBody, _ := json.Marshal("lolol")
				req := httptest.NewRequest(http.MethodPost, "/signup", nil)
				// req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, a.cookieId))
				rec := httptest.NewRecorder()
				handler.SignUp(rec, req)
				return rec
			},
		},
		{
			name: "data already used",
			fields: fields{
				usecase: mock.NewMockAuthUsecase(ctrl),
			},
			args: args{
				data: &models.UserSignUpData{
					Email:    "my@mail.ru",
					Phone:    "+123456",
					Password: "pass",
				},
			},
			want: want{
				user: &models.User{
					ID: id1,
					// FirstName:    "Maksim",
					// SecondName:   "Shagaev",
					Phone:        "+79003249325",
					Email:        "my@mail.ru",
					PasswordHash: "hash",
					// Photo:        "/path/to/photo/test.jpg",
					// DateBirthday: time.Now(),
				},
				message: "data already is used",
				token:   "token",
				status:  http.StatusBadRequest,
				err:     errors.New("error"),
			},
			IsMessage: true,
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().SignUp(gomock.Any(), gomock.Eq(a.data)).Return(w.user, w.token, time.Now(), w.err)
				handler := NewAuthHandler(f.usecase)
				reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
				// req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, a.cookieId))
				rec := httptest.NewRecorder()
				handler.SignUp(rec, req)
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
				assert.Equal(t, tt.want.user.Phone, payload["phone"])
				assert.Equal(t, tt.want.user.Email, payload["email"])
				code, _ := responseMap["statusCode"].(float64)
				assert.Equal(t, tt.want.status, int(code))
			}
		})
	}
}

func TestAuthHandler_Login(t *testing.T) {
	type fields struct {
		usecase *mock.MockAuthUsecase
	}
	type args struct {
		data *models.UserLoginData
	}
	type want struct {
		user    *models.User
		status  int
		err     error
		message string
		token   string
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
			name: "success login",
			fields: fields{
				usecase: mock.NewMockAuthUsecase(ctrl),
			},
			args: args{
				data: &models.UserLoginData{
					Login:    "my@mail.ru",
					Password: "pass",
				},
			},
			want: want{
				user: &models.User{
					ID:           id1,
					Phone:        "+79003249325",
					Email:        "my@mail.ru",
					PasswordHash: "hash",
				},
				message: "",
				token:   "token",
				status:  http.StatusOK,
				err:     nil,
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().Login(gomock.Any(), gomock.Eq(a.data)).Return(w.user, w.token, time.Now(), w.err)
				handler := NewAuthHandler(f.usecase)
				reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodGet, "/login", bytes.NewBuffer(reqBody))
				// req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, a.cookieId))
				rec := httptest.NewRecorder()
				handler.Login(rec, req)
				return rec
			},
		},
		{
			name: "cant read request data login",
			fields: fields{
				usecase: mock.NewMockAuthUsecase(ctrl),
			},
			args: args{
				data: &models.UserLoginData{
					Login:    "my@mail.ru",
					Password: "pass",
				},
			},
			want: want{
				user: &models.User{
					ID:           id1,
					Phone:        "+79003249325",
					Email:        "my@mail.ru",
					PasswordHash: "hash",
				},
				message: "unexpected end of JSON input",
				token:   "token",
				status:  http.StatusBadRequest,
				err:     nil,
			},
			IsMessage: true,
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				// f.usecase.EXPECT().Login(gomock.Any(), gomock.Eq(a.data)).Return(w.user, w.token, time.Now(), w.err)
				handler := NewAuthHandler(f.usecase)
				// reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodGet, "/login", nil)
				// req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, a.cookieId))
				rec := httptest.NewRecorder()
				handler.Login(rec, req)
				return rec
			},
		},
		{
			name: "incorrect password or login",
			fields: fields{
				usecase: mock.NewMockAuthUsecase(ctrl),
			},
			args: args{
				data: &models.UserLoginData{
					Login:    "my@mail.ru",
					Password: "pass",
				},
			},
			want: want{
				user: &models.User{
					ID:           id1,
					Phone:        "+79003249325",
					Email:        "my@mail.ru",
					PasswordHash: "hash",
				},
				message: "incorrect password or login",
				token:   "token",
				status:  http.StatusBadRequest,
				err:     errors.New("wrong password or login"),
			},
			IsMessage: true,
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().Login(gomock.Any(), gomock.Eq(a.data)).Return(w.user, w.token, time.Now(), w.err)
				handler := NewAuthHandler(f.usecase)
				reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodGet, "/login", bytes.NewBuffer(reqBody))
				// req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, a.cookieId))
				rec := httptest.NewRecorder()
				handler.Login(rec, req)
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
				assert.Equal(t, tt.want.user.Phone, payload["phone"])
				assert.Equal(t, tt.want.user.ID.String(), payload["id"])
				assert.Equal(t, tt.want.user.Email, payload["email"])
				code, _ := responseMap["statusCode"].(float64)
				assert.Equal(t, tt.want.status, int(code))
			}
		})
	}
}

func TestAuthHandler_Logout(t *testing.T) {
	type fields struct {
		usecase *mock.MockAuthUsecase
	}
	type want struct {
		status  int
		err     error
		message string
		payload string
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name      string
		fields    fields
		want      want
		IsMessage bool
		prepare   func(f *fields, w *want) *httptest.ResponseRecorder
	}{
		{
			name: "success logout",
			fields: fields{
				usecase: mock.NewMockAuthUsecase(ctrl),
			},
			want: want{
				message: "",
				status:  http.StatusOK,
				err:     nil,
				payload: "success logout",
			},
			prepare: func(f *fields, w *want) *httptest.ResponseRecorder {
				// f.usecase.EXPECT().Login(gomock.Any(), gomock.Eq(a.data)).Return(w.user, w.token, time.Now(), w.err)
				handler := NewAuthHandler(f.usecase)
				// reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodGet, "/auth/logout", nil)
				// req = req.WithContext(context.WithValue(req.Context(), middleware.CookieName, a.cookieId))
				rec := httptest.NewRecorder()
				handler.Logout(rec, req)
				return rec
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := tt.prepare(&tt.fields, &tt.want)
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

func TestAuthHandler_CheckAuth(t *testing.T) {
	type fields struct {
		usecase *mock.MockAuthUsecase
	}
	type args struct {
		token    string
		checkErr error
		id       uuid.UUID
	}
	type want struct {
		status  int
		message string
		payload string
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
			name: "success check",
			fields: fields{
				usecase: mock.NewMockAuthUsecase(ctrl),
			},
			args: args{
				token:    "token",
				checkErr: nil,
				id:       id1,
			},
			want: want{
				message: "",
				payload: id1.String(),
				status:  http.StatusOK,
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().CheckAuth(gomock.Any(), gomock.Eq(a.id)).Return(a.checkErr)
				handler := NewAuthHandler(f.usecase)
				req := httptest.NewRequest(http.MethodGet, "/auth/logout", nil)
				ctx := context.WithValue(req.Context(), middleware.CookieName, a.id)
				req = req.WithContext(ctx)
				rec := httptest.NewRecorder()
				handler.CheckAuth(rec, req)
				return rec
			},
		},
		{
			name: "no cookie",
			fields: fields{
				usecase: mock.NewMockAuthUsecase(ctrl),
			},
			args: args{
				token:    "token",
				checkErr: nil,
				id:       id1,
			},
			want: want{
				message: "token not found",
				payload: id1.String(),
				status:  http.StatusUnauthorized,
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				// f.usecase.EXPECT().CheckAuth(gomock.Any(), gomock.Eq(a.id)).Return(a.checkErr)
				handler := NewAuthHandler(f.usecase)
				req := httptest.NewRequest(http.MethodGet, "/auth/logout", nil)
				// ctx := context.WithValue(req.Context(), middleware.CookieName, a.id.String())
				// req = req.WithContext(ctx)
				rec := httptest.NewRecorder()
				handler.CheckAuth(rec, req)
				return rec
			},
			IsMessage: true,
		},
		{
			name: "incorrect user id",
			fields: fields{
				usecase: mock.NewMockAuthUsecase(ctrl),
			},
			args: args{
				token:    "token",
				checkErr: errors.New("error"),
				id:       id1,
			},
			want: want{
				message: "incorrect user id",
				payload: id1.String(),
				status:  http.StatusUnauthorized,
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				// f.usecase.EXPECT().CheckAuth(gomock.Any(), gomock.Eq(a.id)).Return(a.checkErr)
				handler := NewAuthHandler(f.usecase)
				req := httptest.NewRequest(http.MethodGet, "/auth/logout", nil)
				ctx := context.WithValue(req.Context(), middleware.CookieName, "notid")
				req = req.WithContext(ctx)
				rec := httptest.NewRecorder()
				handler.CheckAuth(rec, req)
				return rec
			},
			IsMessage: true,
		},
		{
			name: "user not exists",
			fields: fields{
				usecase: mock.NewMockAuthUsecase(ctrl),
			},
			args: args{
				token:    "token",
				checkErr: errors.New("error"),
				id:       id1,
			},
			want: want{
				message: "user not exists",
				payload: id1.String(),
				status:  http.StatusUnauthorized,
			},
			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().CheckAuth(gomock.Any(), gomock.Eq(a.id)).Return(a.checkErr)
				handler := NewAuthHandler(f.usecase)
				req := httptest.NewRequest(http.MethodGet, "/auth/logout", nil)
				ctx := context.WithValue(req.Context(), middleware.CookieName, a.id)
				req = req.WithContext(ctx)
				rec := httptest.NewRecorder()
				handler.CheckAuth(rec, req)
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
