package grpc_test

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts/delivery/grpc"
	genAdverts "2024_1_TeaStealers/internal/pkg/adverts/delivery/grpc/gen"
	adverts_mock "2024_1_TeaStealers/internal/pkg/adverts/mock"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestGetAdvertById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := adverts_mock.NewMockAdvertUsecase(ctrl)
	logger := zap.Must(zap.NewDevelopment())
	serverHandler := grpc.NewServerAdvertsHandler(mockUsecase, logger)

	id1 := int64(102)
	advData := &genAdverts.GetAdvertByIdRequest{ //это входной request
		Id: id1,
	}

	advResp := &genAdverts.GetAdvertByIdResponse{ //это response
		Id:           id1,
		DateCreation: "0001-01-01 00:00:00 +0000 UTC", // тут мини костыль со временем
	}

	ucRes := &models.AdvertData{ // это что вернёт usecase
		ID: id1,
	}

	type args struct {
		req *genAdverts.GetAdvertByIdRequest
	}
	type want struct {
		resp   *genAdverts.GetAdvertByIdResponse
		ucWant *models.AdvertData
		ucErr  error
		err    error
	}

	tests := []struct {
		name    string
		args    args
		want    want
		prepare func(a *args, w *want)
	}{
		{
			name: "successful GetAdvertById",
			args: args{
				req: advData,
			},
			want: want{
				ucWant: ucRes,
				resp:   advResp,
				ucErr:  nil,
				err:    nil,
			},
			/*
				 AdvertData struct {
				    ID                int64                    `json:"advertId"`
				    AdvertType        string                   `json:"advertType"`
				    TypeSale          string                   `json:"typeSale"`
				    Title             string                   `json:"title"`
				    Description       string                   `json:"description"`
				    CountViews        int64                    `json:"countViews"`
				    CountLikes        int64                    `json:"countLikes"`
				    Price             int64                    `json:"price"`
				    Phone             string                   `json:"phone"`
				    IsLiked           bool                     `json:"isLiked"`
				    IsAgent           bool                     `json:"isAgent"`
				    Metro             string                   `json:"metro"`
				    Address           string                   `json:"adress"`
				    AddressPoint      string                   `json:"adressPoint"`
				    PriceChange       []*PriceChangeData       `json:"priceHistory"`
				    Images            []*ImageResp             `json:"images"`
				    HouseProperties   *HouseProperties         `json:"houseProperties,omitempty"`
				    FlatProperties    *FlatProperties          `json:"flatProperties,omitempty"`
				    YearCreation      int                      `json:"yearCreation"`
				    Material          MaterialBuilding         `json:"material"`
				    ComplexProperties *ComplexAdvertProperties `json:"complexProperties,omitempty"`
				    DateCreation      time.Time                `json:"dateCreation"`
				}
			*/
			prepare: func(a *args, w *want) { // здесь пишем моки функций
				// их порядок мокирования должен совпадать с порядком их вызова
				// внутри тестируемой функции иначе будет ошибка
				mockUsecase.EXPECT().GetAdvertById(gomock.Any(), id1).Return(w.ucWant, w.ucErr) // вместо контекста gomock.Any() - иначе не совпадут из-за requestId
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(&tt.args, &tt.want)                                                    // иницилизация моков
			gotResp, gotErr := serverHandler.GetAdvertById(context.Background(), tt.args.req) // вызов функции с request из arg

			assert.Equal(t, tt.want.resp, gotResp)
			assert.Equal(t, tt.want.err, gotErr)

		})
	}
}
