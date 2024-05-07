package usecase_test

import (
	"2024_1_TeaStealers/internal/models"
	companies_mock "2024_1_TeaStealers/internal/pkg/companies/mock"
	"2024_1_TeaStealers/internal/pkg/companies/usecase"
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewCompany(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := companies_mock.NewMockCompanyRepo(ctrl)
	logger := zap.Must(zap.NewDevelopment())
	usecase := usecase.NewCompanyUsecase(mockRepo, logger)
	id := rand.Int63()
	year := 2022
	type args struct {
		companyUUID int64
		data        *models.CompanyCreateData
	}
	type want struct {
		company *models.Company
		err     error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful create company",
			args: args{
				companyUUID: id,
				data: &models.CompanyCreateData{
					Name:        "name1",
					YearFounded: year,
					Phone:       "+7115251523",
					Description: "description",
				},
			},
			want: want{
				company: &models.Company{
					ID:           id,
					Photo:        "/url/to/photo",
					Name:         "name1",
					YearFounded:  year,
					Description:  "description",
					DateCreation: time.Now(),
					Phone:        "+7115251523",
					IsDeleted:    false,
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().CreateCompany(gomock.Any(), gomock.Any()).
				Return(tt.want.company, tt.want.err)
			gotUser, goterr := usecase.CreateCompany(
				context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tt.args.data)
			assert.Equal(t, tt.want.company.ID, gotUser.ID)
			assert.Equal(t, tt.want.err, goterr)
		})
	}
}

func TestGetCompanyById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := companies_mock.NewMockCompanyRepo(ctrl)
	logger := zap.Must(zap.NewDevelopment())
	usecase := usecase.NewCompanyUsecase(mockRepo, logger)
	id := rand.Int63()
	year := 2022
	type args struct {
		companyUUID int64
	}
	type want struct {
		company *models.CompanyData
		err     error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful get company",
			args: args{
				companyUUID: id,
			},
			want: want{
				company: &models.CompanyData{
					ID:          id,
					Photo:       "/url/to/photo",
					Name:        "name1",
					YearFounded: year,
					Description: "description",
					Phone:       "+7115251523",
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().GetCompanyById(gomock.Any(), gomock.Any()).
				Return(tt.want.company, tt.want.err)
			gotUser, goterr := usecase.GetCompanyById(
				context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tt.args.companyUUID)
			assert.Equal(t, tt.want.company.ID, gotUser.ID)
			assert.Equal(t, tt.want.err, goterr)
		})
	}
}
