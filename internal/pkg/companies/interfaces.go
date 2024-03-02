package companies

import (
	"2024_1_TeaStealers/internal/models"
	"context"
)

// CompanyUsecase represents the usecase interface for manage companies.
type CompanyUsecase interface {
	CreateCompany(ctx context.Context, data *models.CompanyCreateData) (*models.Company, error)
}

// CompanyRepo represents the repository interface for manage companies.
type CompanyRepo interface {
	CreateCompany(ctx context.Context, company *models.Company) error
}
