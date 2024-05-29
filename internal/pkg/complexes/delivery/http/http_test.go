package delivery

/*
import (
	"2024_1_TeaStealers/internal/models"
	mocks "2024_1_TeaStealers/internal/pkg/complexes/mock"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

	func TestImagesHandler_CreateComplex(t *testing.T) {
		type fields struct {
			usecase *mocks.MockComplexUsecase
		}
		type args struct {
			data *models.ComplexCreateData
		}
		type want struct {
			complexResp *models.Complex
			status      int
			err         error
		}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		id1 := uuid.NewV4()
		id2 := uuid.NewV4()
		id3 := uuid.NewV4()
		tests := []struct {
			name    string
			fields  fields
			args    args
			want    want
			prepare func(f *fields, a *args, w *want) *http.Response
			check   func(response *http.Response, a *args, w *want)
		}{
			{
				name: "success",
				fields: fields{
					usecase: mocks.NewMockComplexUsecase(ctrl),
				},
				args: args{
					advertId: id1,
				},
				want: want{
					imagesResp: []*models.ImageResp{
						{
							ID:       id2,
							Photo:    "/images/image1.jpg",
							Priority: 1,
						},
						{
							ID:       id3,
							Photo:    "/images/image2.jpg",
							Priority: 2,
						},
					},
					status: http.StatusOK,
					err:    nil,
				},
				prepare: func(f *fields, a *args, w *want) *http.Response {
					f.usecase.EXPECT().GetAdvertImages(gomock.Eq(a.advertId)).Return(w.imagesResp, nil)
					handler := NewImageHandler(f.usecase)
					vars := map[string]string{
						"id": a.advertId.String(),
					}

					body := &bytes.Buffer{}
					// writer := multipart.NewWriter(body)
					// writer.Close()
					req, _ := http.NewRequest(http.MethodGet, "/{id}/image", body)
					req = mux.SetURLVars(req, vars)
					recorder := httptest.NewRecorder()
					handler.GetAdvertImages(recorder, req)
					resp := recorder.Result()
					return resp
				},
				check: func(response *http.Response, args *args, want *want) {
					assert.Equal(t, http.StatusOK, response.StatusCode)
					bodyBytes, _ := io.ReadAll(response.Body)
					var responseMap map[string]interface{}
					err := json.Unmarshal(bodyBytes, &responseMap)
					if err != nil {
						fmt.Println("Error unmarshalling response:", err)
					}
					payload := responseMap["payload"].([]interface{})
					images := make([]*models.ImageResp, len(payload))
					for i, p := range payload {
						image := p.(map[string]interface{})
						images[i] = &models.ImageResp{
							ID:       uuid.FromStringOrNil(image["id"].(string)),
							Photo:    image["photo"].(string),
							Priority: int(image["priority"].(float64)),
						}
					}
					assert.Equal(t, want.imagesResp, images)
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				resp := tt.prepare(&tt.fields, &tt.args, &tt.want)
				tt.check(resp, &tt.args, &tt.want)
			})
		}

}
*/
/*
prepare: func(f *fields, a *args, w *want) *http.Response {
					f.usecase.EXPECT().GetAdvertImages(gomock.Eq(a.advertId)).Return(w.imagesResp, nil)
					handler := NewImageHandler(f.usecase)
					vars := map[string]string{
						"id": a.advertId.String(),
					}

					body := &bytes.Buffer{}
					// writer := multipart.NewWriter(body)
					// writer.Close()
					req, _ := http.NewRequest(http.MethodGet, "/{id}/image", body)
					req = mux.SetURLVars(req, vars)
					recorder := httptest.NewRecorder()
					handler.GetAdvertImages(recorder, req)
					resp := recorder.Result()
					return resp
				},
*/
/*func TestComplexHandler_CreateComplex(t *testing.T) {
	type fields struct {
		usecase *mocks.MockComplexUsecase
	}
	type args struct {
		data *models.ComplexCreateData
	}
	type want struct {
		complexResp *models.Complex
		status      int
		err         error
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id1 := uuid.NewV4()
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        want
		IsMessage   bool
		message     string
		prepare     func(f *fields, a *args, w *want) *httptest.ResponseRecorder
		needExpect1 bool
	}{
		{
			needExpect1: true,
			name:        "success create",
			fields: fields{
				usecase: mocks.NewMockComplexUsecase(ctrl),
			},
			args: args{
				data: &models.ComplexCreateData{
					CompanyId: id1,
					Name:      "name",
					Address:   "adress",
				},
			},
			want: want{
				complexResp: &models.Complex{
					CompanyId: id1,
					Name:      "name",
					Address:   "adress",
				},
				status: http.StatusCreated,
				err:    nil,
			},
			IsMessage: false,
			message:   "",

			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().CreateComplex(gomock.Any(), gomock.Eq(a.data)).Return(w.complexResp, w.err)
				handler := NewComplexHandler(f.usecase, &zap.Logger{})
				reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodPost, "/create/complex", bytes.NewBuffer(reqBody))
				rec := httptest.NewRecorder()
				handler.CreateComplex(rec, req)
				return rec
			},
		},
		{
			needExpect1: true,
			name:        "data already is used crComplex",
			fields: fields{
				usecase: mocks.NewMockComplexUsecase(ctrl),
			},
			args: args{
				data: &models.ComplexCreateData{
					CompanyId: id1,
					Name:      "name",
					Address:   "adress",
				},
			},
			want: want{
				complexResp: &models.Complex{
					CompanyId: id1,
					Name:      "name",
					Address:   "adress",
				},
				status: http.StatusBadRequest,
				err:    errors.New("error"),
			},
			IsMessage: true,
			message:   "data already is used",

			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().CreateComplex(gomock.Any(), gomock.Eq(a.data)).Return(w.complexResp, w.err)
				handler := NewComplexHandler(f.usecase, &zap.Logger{})
				reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodPost, "/create/complex", bytes.NewBuffer(reqBody))
				rec := httptest.NewRecorder()
				handler.CreateComplex(rec, req)
				return rec
			},
		},
		{
			needExpect1: true,
			name:        "incorrect data format crComplex",
			fields: fields{
				usecase: mocks.NewMockComplexUsecase(ctrl),
			},
			args: args{
				data: &models.ComplexCreateData{
					CompanyId: id1,
					Name:      "name",
					Address:   "address",
				},
			},
			want: want{
				complexResp: &models.Complex{
					CompanyId: id1,
					Name:      "name",
					Address:   "address",
				},
				status: http.StatusBadRequest,
				err:    errors.New("error"),
			},
			IsMessage: true,
			message:   "incorrect data format",

			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				// f.usecase.EXPECT().CreateComplex(gomock.Any(), gomock.Eq(a.data)).Return(w.complexResp, w.err)
				handler := NewComplexHandler(f.usecase, &zap.Logger{})
				// reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodPost, "/create/complex", nil)
				rec := httptest.NewRecorder()
				handler.CreateComplex(rec, req)
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
				assert.Equal(t, tt.message, responseMap["message"])
			} else {
				// payload := responseMap["payload"].(map[string]interface{})
				if err != nil {
					fmt.Println("Error unmarshalling response:", err)
				}
				// assert.Equal(t, tt.want.complexResp.Address, payload["address"])
				// assert.Equal(t, tt.want.complexResp.ID, payload["id"].(uuid.UUID))
				// assert.Equal(t, tt.want.user.Email, payload["email"])
				code, _ := responseMap["statusCode"].(float64)
				assert.Equal(t, tt.want.status, int(code))
			}
		})
	}
}
*/ /*
func TestComplexHandler_CreateBuilding(t *testing.T) {
	type fields struct {
		usecase *mocks.MockComplexUsecase
	}
	type args struct {
		data *models.BuildingCreateData
	}
	type want struct {
		complexResp *models.Building
		status      int
		err         error
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id1 := uuid.NewV4()
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        want
		IsMessage   bool
		message     string
		prepare     func(f *fields, a *args, w *want) *httptest.ResponseRecorder
		needExpect1 bool
	}{
		{
			needExpect1: true,
			name:        "success create building",
			fields: fields{
				usecase: mocks.NewMockComplexUsecase(ctrl),
			},
			args: args{
				data: &models.BuildingCreateData{
					ComplexID: id1,
					Address:   "adress",
				},
			},
			want: want{
				complexResp: &models.Building{
					ComplexID: id1,
					Address:   "adress",
				},
				status: http.StatusCreated,
				err:    nil,
			},
			IsMessage: false,
			message:   "",

			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().CreateBuilding(gomock.Any(), gomock.Eq(a.data)).Return(w.complexResp, w.err)
				handler := NewComplexHandler(f.usecase, &zap.Logger{})
				reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodPost, "/create/complex", bytes.NewBuffer(reqBody))
				rec := httptest.NewRecorder()
				handler.CreateBuilding(rec, req)
				return rec
			},
		},
		{
			needExpect1: true,
			name:        "data already is used crComplex",
			fields: fields{
				usecase: mocks.NewMockComplexUsecase(ctrl),
			},
			args: args{
				data: &models.BuildingCreateData{
					ComplexID: id1,
					Address:   "adress",
				},
			},
			want: want{
				complexResp: &models.Building{
					ComplexID: id1,
					Address:   "adress",
				},
				status: http.StatusBadRequest,
				err:    errors.New("error"),
			},
			IsMessage: true,
			message:   "data already is used",

			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().CreateBuilding(gomock.Any(), gomock.Eq(a.data)).Return(w.complexResp, w.err)
				handler := NewComplexHandler(f.usecase, &zap.Logger{})
				reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodPost, "/create/complex", bytes.NewBuffer(reqBody))
				rec := httptest.NewRecorder()
				handler.CreateBuilding(rec, req)
				return rec
			},
		},
		{
			needExpect1: true,
			name:        "incorrect data format crComplex",
			fields: fields{
				usecase: mocks.NewMockComplexUsecase(ctrl),
			},
			args: args{
				data: &models.BuildingCreateData{
					ComplexID: id1,
					Address:   "adress",
				},
			},
			want: want{
				complexResp: &models.Building{
					ComplexID: id1,
					Address:   "adress",
				},
				status: http.StatusBadRequest,
				err:    errors.New("error"),
			},
			IsMessage: true,
			message:   "incorrect data format",

			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				// f.usecase.EXPECT().CreateBuilding(gomock.Any(), gomock.Eq(a.data)).Return(w.complexResp, w.err)
				handler := NewComplexHandler(f.usecase, &zap.Logger{})
				// reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodPost, "/create/complex", nil)
				rec := httptest.NewRecorder()
				handler.CreateBuilding(rec, req)
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
				assert.Equal(t, tt.message, responseMap["message"])
			} else {
				// payload := responseMap["payload"].(map[string]interface{})
				if err != nil {
					fmt.Println("Error unmarshalling response:", err)
				}
				// assert.Equal(t, tt.want.complexResp.Address, payload["address"])
				// assert.Equal(t, tt.want.complexResp.ID, payload["id"].(uuid.UUID))
				// assert.Equal(t, tt.want.user.Email, payload["email"])
				code, _ := responseMap["statusCode"].(float64)
				assert.Equal(t, tt.want.status, int(code))
			}
		})
	}
}

func TestComplexHandler_CreateHouseAdvert(t *testing.T) {
	type fields struct {
		usecase *mocks.MockComplexUsecase
	}
	type args struct {
		data *models.ComplexAdvertHouseCreateData
	}
	type want struct {
		complexResp *models.Advert
		status      int
		err         error
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id1 := uuid.NewV4()
	id2 := uuid.NewV4()
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        want
		IsMessage   bool
		message     string
		prepare     func(f *fields, a *args, w *want) *httptest.ResponseRecorder
		needExpect1 bool
	}{
		{
			needExpect1: true,
			name:        "success create building",
			fields: fields{
				usecase: mocks.NewMockComplexUsecase(ctrl),
			},
			args: args{
				data: &models.ComplexAdvertHouseCreateData{
					UserID:  id2,
					Address: "adress",
				},
			},
			want: want{
				complexResp: &models.Advert{
					ID: id1,
				},
				status: http.StatusCreated,
				err:    nil,
			},
			IsMessage: false,
			message:   "",

			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().CreateHouseAdvert(gomock.Any(), gomock.Eq(a.data)).Return(w.complexResp, w.err)
				handler := NewComplexHandler(f.usecase, &zap.Logger{})
				reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodPost, "/create/complex", bytes.NewBuffer(reqBody))
				rec := httptest.NewRecorder()
				handler.CreateHouseAdvert(rec, req)
				return rec
			},
		},
		{
			needExpect1: true,
			name:        "data already is used crComplex",
			fields: fields{
				usecase: mocks.NewMockComplexUsecase(ctrl),
			},
			args: args{
				data: &models.ComplexAdvertHouseCreateData{
					UserID:  id2,
					Address: "adress",
				},
			},
			want: want{
				complexResp: &models.Advert{
					ID: id1,
				},
				status: http.StatusBadRequest,
				err:    errors.New("error create house advert"),
			},
			IsMessage: true,
			message:   "error create house advert",

			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().CreateHouseAdvert(gomock.Any(), gomock.Eq(a.data)).Return(w.complexResp, w.err)
				handler := NewComplexHandler(f.usecase, &zap.Logger{})
				reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodPost, "/create/complex", bytes.NewBuffer(reqBody))
				rec := httptest.NewRecorder()
				handler.CreateHouseAdvert(rec, req)
				return rec
			},
		},
		{
			needExpect1: true,
			name:        "incorrect data format crComplex",
			fields: fields{
				usecase: mocks.NewMockComplexUsecase(ctrl),
			},
			args: args{
				data: &models.ComplexAdvertHouseCreateData{
					UserID:  id2,
					Address: "adress",
				},
			},
			want: want{
				complexResp: &models.Advert{
					ID: id1,
				},
				status: http.StatusBadRequest,
				err:    errors.New("error"),
			},
			IsMessage: true,
			message:   "incorrect data format",

			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				// f.usecase.EXPECT().CreateHouseAdvert(gomock.Any(), gomock.Eq(a.data)).Return(w.complexResp, w.err)
				handler := NewComplexHandler(f.usecase, &zap.Logger{})
				// reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodPost, "/create/complex", nil)
				rec := httptest.NewRecorder()
				handler.CreateHouseAdvert(rec, req)
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
				assert.Equal(t, tt.message, responseMap["message"])
			} else {
				// payload := responseMap["payload"].(map[string]interface{})
				if err != nil {
					fmt.Println("Error unmarshalling response:", err)
				}
				// assert.Equal(t, tt.want.complexResp.Address, payload["address"])
				// assert.Equal(t, tt.want.complexResp.ID, payload["id"].(uuid.UUID))
				// assert.Equal(t, tt.want.user.Email, payload["email"])
				code, _ := responseMap["statusCode"].(float64)
				assert.Equal(t, tt.want.status, int(code))
			}
		})
	}
}

/*
func TestComplexHandler_CreateFlatAdvert(t *testing.T) {
	type fields struct {
		usecase *mocks.MockComplexUsecase
	}
	type args struct {
		data *models.ComplexAdvertFlatCreateData
	}
	type want struct {
		complexResp *models.Advert
		status      int
		err         error
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id1 := uuid.NewV4()
	id2 := uuid.NewV4()
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        want
		IsMessage   bool
		message     string
		prepare     func(f *fields, a *args, w *want) *httptest.ResponseRecorder
		needExpect1 bool
	}{
		{
			needExpect1: true,
			name:        "success create building",
			fields: fields{
				usecase: mocks.NewMockComplexUsecase(ctrl),
			},
			args: args{
				data: &models.ComplexAdvertFlatCreateData{
					UserID:  id2,
					Address: "adress",
				},
			},
			want: want{
				complexResp: &models.Advert{
					ID: id1,
				},
				status: http.StatusCreated,
				err:    nil,
			},
			IsMessage: false,
			message:   "",

			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().CreateFlatAdvert(gomock.Any(), gomock.Eq(a.data)).Return(w.complexResp, w.err)
				handler := NewComplexHandler(f.usecase, &zap.Logger{})
				reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodPost, "/create/complex", bytes.NewBuffer(reqBody))
				rec := httptest.NewRecorder()
				handler.CreateFlatAdvert(rec, req)
				return rec
			},
		},
		{
			needExpect1: true,
			name:        "data already is used crComplex",
			fields: fields{
				usecase: mocks.NewMockComplexUsecase(ctrl),
			},
			args: args{
				data: &models.ComplexAdvertFlatCreateData{
					UserID:  id2,
					Address: "adress",
				},
			},
			want: want{
				complexResp: &models.Advert{
					ID: id1,
				},
				status: http.StatusBadRequest,
				err:    errors.New("error create house advert"),
			},
			IsMessage: true,
			message:   "error create house advert",

			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().CreateFlatAdvert(gomock.Any(), gomock.Eq(a.data)).Return(w.complexResp, w.err)
				handler := NewComplexHandler(f.usecase, &zap.Logger{})
				reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodPost, "/create/complex", bytes.NewBuffer(reqBody))
				rec := httptest.NewRecorder()
				handler.CreateFlatAdvert(rec, req)
				return rec
			},
		},
		{
			needExpect1: true,
			name:        "incorrect data format crComplex",
			fields: fields{
				usecase: mocks.NewMockComplexUsecase(ctrl),
			},
			args: args{
				data: &models.ComplexAdvertFlatCreateData{
					UserID:  id2,
					Address: "adress",
				},
			},
			want: want{
				complexResp: &models.Advert{
					ID: id1,
				},
				status: http.StatusBadRequest,
				err:    errors.New("error"),
			},
			IsMessage: true,
			message:   "incorrect data format",

			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				// f.usecase.EXPECT().CreateFlatAdvert(gomock.Any(), gomock.Eq(a.data)).Return(w.complexResp, w.err)
				handler := NewComplexHandler(f.usecase, &zap.Logger{})
				// reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodPost, "/create/complex", nil)
				rec := httptest.NewRecorder()
				handler.CreateFlatAdvert(rec, req)
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
				assert.Equal(t, tt.message, responseMap["message"])
			} else {
				// payload := responseMap["payload"].(map[string]interface{})
				if err != nil {
					fmt.Println("Error unmarshalling response:", err)
				}
				// assert.Equal(t, tt.want.complexResp.Address, payload["address"])
				// assert.Equal(t, tt.want.complexResp.ID, payload["id"].(uuid.UUID))
				// assert.Equal(t, tt.want.user.Email, payload["email"])
				code, _ := responseMap["statusCode"].(float64)
				assert.Equal(t, tt.want.status, int(code))
			}
		})
	}
}

func TestComplexHandler_GetComplexById(t *testing.T) {
	type fields struct {
		usecase *mocks.MockComplexUsecase
	}
	type args struct {
		data uuid.UUID
	}
	type want struct {
		complexResp *models.ComplexData
		status      int
		err         error
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id1 := uuid.NewV4()
	// id2 := uuid.NewV4()
	tests := []struct {
		name        string
		fields      fields
		args        args
		want        want
		IsMessage   bool
		message     string
		prepare     func(f *fields, a *args, w *want) *httptest.ResponseRecorder
		needExpect1 bool
	}{
		{
			needExpect1: true,
			name:        "success create building",
			fields: fields{
				usecase: mocks.NewMockComplexUsecase(ctrl),
			},
			args: args{
				data: id1,
			},
			want: want{
				complexResp: &models.ComplexData{
					ID:   id1,
					Name: "name complex",
				},
				status: http.StatusOK,
				err:    nil,
			},
			IsMessage: false,
			message:   "",

			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().GetComplexById(gomock.Any(), gomock.Eq(a.data)).Return(w.complexResp, w.err)
				handler := NewComplexHandler(f.usecase, &zap.Logger{})
				vars := map[string]string{
					"id": id1.String(),
				}

				// writer := multipart.NewWriter(body)
				// writer.Close()
				reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodPost, "/create/complex", bytes.NewBuffer(reqBody))
				req = mux.SetURLVars(req, vars)

				rec := httptest.NewRecorder()
				handler.GetComplexById(rec, req)
				return rec
			},
		},
		{
			needExpect1: true,
			name:        "id parameter is required",
			fields: fields{
				usecase: mocks.NewMockComplexUsecase(ctrl),
			},
			args: args{
				data: id1,
			},
			want: want{
				complexResp: &models.ComplexData{
					ID:   id1,
					Name: "name complex",
				},
				status: http.StatusBadRequest,
				err:    nil,
			},
			IsMessage: true,
			message:   "id parameter is required",

			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				// f.usecase.EXPECT().GetComplexById(gomock.Any(), gomock.Eq(a.data)).Return(w.complexResp, w.err)
				handler := NewComplexHandler(f.usecase, &zap.Logger{})
				// vars := map[string]string{
				// 	"id": id1.String(),
				// }
				reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodPost, "/create/complex", bytes.NewBuffer(reqBody))
				// req = mux.SetURLVars(req, vars)

				rec := httptest.NewRecorder()
				handler.GetComplexById(rec, req)
				return rec
			},
		},

		{
			needExpect1: true,
			name:        "error get complex by id",
			fields: fields{
				usecase: mocks.NewMockComplexUsecase(ctrl),
			},
			args: args{
				data: id1,
			},
			want: want{
				complexResp: &models.ComplexData{
					ID:   id1,
					Name: "name complex",
				},
				status: http.StatusBadRequest,
				err:    errors.New("some error"),
			},
			IsMessage: true,
			message:   "some error",

			prepare: func(f *fields, a *args, w *want) *httptest.ResponseRecorder {
				f.usecase.EXPECT().GetComplexById(gomock.Any(), gomock.Eq(a.data)).Return(w.complexResp, w.err)
				handler := NewComplexHandler(f.usecase, &zap.Logger{})
				vars := map[string]string{
					"id": id1.String(),
				}

				// writer := multipart.NewWriter(body)
				// writer.Close()
				reqBody, _ := json.Marshal(a.data)
				req := httptest.NewRequest(http.MethodPost, "/create/complex", bytes.NewBuffer(reqBody))
				req = mux.SetURLVars(req, vars)

				rec := httptest.NewRecorder()
				handler.GetComplexById(rec, req)
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
				assert.Equal(t, tt.message, responseMap["message"])
			} else {
				// payload := responseMap["payload"].(map[string]interface{})
				if err != nil {
					fmt.Println("Error unmarshalling response:", err)
				}
				// assert.Equal(t, tt.want.complexResp.Address, payload["address"])
				// assert.Equal(t, tt.want.complexResp.ID, payload["id"].(uuid.UUID))
				// assert.Equal(t, tt.want.user.Email, payload["email"])
				code, _ := responseMap["statusCode"].(float64)
				assert.Equal(t, tt.want.status, int(code))
			}
		})
	}
}
*/
