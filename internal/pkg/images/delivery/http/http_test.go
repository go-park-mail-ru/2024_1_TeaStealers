package delivery

import (
	"2024_1_TeaStealers/internal/models"
	mocks "2024_1_TeaStealers/internal/pkg/images/mock"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestImagesHandler_UploadImage(t *testing.T) {
	type fields struct {
		usecase *mocks.MockImageUsecase
	}
	type args struct {
		imageId  uuid.UUID
		advertId uuid.UUID
		fileType string
		fileName string
	}
	type want struct {
		imagesResp *models.ImageResp
		status     int
		err        error
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id1 := uuid.NewV4()
	id2 := uuid.NewV4()
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
				usecase: mocks.NewMockImageUsecase(ctrl),
			},
			args: args{
				imageId:  id1,
				advertId: id2,
				fileType: ".jpeg",
				fileName: "testdata/test_image.jpeg",
			},
			want: want{
				imagesResp: &models.ImageResp{
					ID:       id1,
					Photo:    "/images/" + id1.String() + ".jpeg",
					Priority: 1,
				},
				status: http.StatusCreated,
				err:    nil,
			},
			prepare: func(f *fields, a *args, w *want) *http.Response {
				f.usecase.EXPECT().UploadImage(gomock.Any(), gomock.Eq(a.fileType), gomock.Eq(a.advertId)).Return(w.imagesResp, nil)
				handler := NewImageHandler(f.usecase, &zap.Logger{})
				file, _ := os.Open(a.fileName)
				defer file.Close()
				body := &bytes.Buffer{}
				writer := multipart.NewWriter(body)
				part, _ := writer.CreateFormFile("file", a.fileName)
				_, err := io.Copy(part, file)
				if err != nil {
					log.Fatalf("error copy test file")
				}
				writer.Close()

				req, _ := http.NewRequest(http.MethodPost, "/image", body)
				req.Header.Set("Content-Type", writer.FormDataContentType())
				q := req.URL.Query()
				q.Add("id", a.advertId.String())
				req.URL.RawQuery = q.Encode()

				recorder := httptest.NewRecorder()
				handler.UploadImage(recorder, req)

				resp := recorder.Result()

				return resp

			},
			check: func(response *http.Response, args *args, want *want) {
				assert.Equal(t, http.StatusCreated, response.StatusCode)
				bodyBytes, _ := io.ReadAll(response.Body)
				var responseMap map[string]interface{}
				err := json.Unmarshal(bodyBytes, &responseMap)
				if err != nil {
					fmt.Println("Error unmarshalling response:", err)
				}
				assert.Equal(t, want.imagesResp.ID.String(), responseMap["payload"].(map[string]interface{})["id"])
				assert.Equal(t, "/images/"+args.imageId.String()+args.fileType, responseMap["payload"].(map[string]interface{})["photo"])
				assert.Equal(t, want.imagesResp.Priority, int(responseMap["payload"].(map[string]interface{})["priority"].(float64)))
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

func TestImagesHandler_GetAdvertImages(t *testing.T) {
	type fields struct {
		usecase *mocks.MockImageUsecase
	}
	type args struct {
		advertId uuid.UUID
	}
	type want struct {
		imagesResp []*models.ImageResp
		status     int
		err        error
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
				usecase: mocks.NewMockImageUsecase(ctrl),
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
				handler := NewImageHandler(f.usecase, &zap.Logger{})
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

// добавить проверки на возвращаемые ошибки NoError и проверить message и StatusCode payload!!!
func TestImagesHandler_DeleteImage(t *testing.T) {
	type fields struct {
		usecase *mocks.MockImageUsecase
	}
	type args struct {
		ImageId uuid.UUID
	}
	type want struct {
		imagesResp []*models.ImageResp
		status     int
		err        error
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
				usecase: mocks.NewMockImageUsecase(ctrl),
			},
			args: args{
				ImageId: id1,
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
				f.usecase.EXPECT().DeleteImage(gomock.Eq(a.ImageId)).Return(w.imagesResp, nil)
				handler := NewImageHandler(f.usecase, &zap.Logger{})
				vars := map[string]string{
					"id": a.ImageId.String(),
				}

				body := &bytes.Buffer{}
				req, _ := http.NewRequest(http.MethodDelete, "/{id}/image", body)
				req = mux.SetURLVars(req, vars)
				recorder := httptest.NewRecorder()
				handler.DeleteImage(recorder, req)
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
