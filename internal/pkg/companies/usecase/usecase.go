package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/companies"
	"context"

	"time"

	"github.com/satori/uuid"
)

// CompanyUsecase represents the usecase for manage companies.
type CompanyUsecase struct {
	repo companies.CompanyRepo
}

// NewCompanyUsecase creates a new instance of CompanyUsecase.
func NewCompanyUsecase(repo companies.CompanyRepo) *CompanyUsecase {
	return &CompanyUsecase{repo: repo}
}

// CreateCompany handles the company creation process.
func (u *CompanyUsecase) CreateCompany(ctx context.Context, data *models.CompanyCreateData) (*models.Company, error) {
	newCompany := &models.Company{
		ID:           uuid.NewV4(),
		Name:         data.Name,
		Phone:        data.Phone,
		Descpription: data.Descpription,
		DataCreation: time.Now(),
		IsDeleted:    false,
	}

	if err := u.repo.CreateCompany(ctx, newCompany); err != nil {
		return nil, err
	}

	return newCompany, nil
}
