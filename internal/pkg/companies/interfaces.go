package companies

import (
	"2024_1_TeaStealers/internal/models"
	"context"

	"github.com/satori/uuid"
)

// CompanyUsecase represents the usecase interface for manage companies.
type CompanyUsecase interface {
	CreateCompany(ctx context.Context, data *models.CompanyCreateData) (*models.Company, error)
	GetCompanyById(ctx context.Context, id uuid.UUID) (findCompany *models.Company, err error)
	GetCompaniesList(ctx context.Context) (findCompanies []*models.Company, err error)
	DeleteCompanyById(ctx context.Context, id uuid.UUID) (err error)
}

// CompanyRepo represents the repository interface for manage companies.
type CompanyRepo interface {
	CreateCompany(ctx context.Context, company *models.Company) error
	GetCompanyById(ctx context.Context, id uuid.UUID) (*models.Company, error)
	GetCompaniesList(ctx context.Context) ([]*models.Company, error)
	DeleteCompanyById(ctx context.Context, id uuid.UUID) error
}
