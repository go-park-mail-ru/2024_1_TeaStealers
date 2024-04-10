//go:generate mockgen -destination=mock/${GOFILE} -package=${GOPACKAGE}_mock -source=${GOFILE}
package companies

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"io"

	"github.com/satori/uuid"
)

// CompanyUsecase represents the usecase interface for companies.
type CompanyUsecase interface {
	CreateCompany(ctx context.Context, data *models.CompanyCreateData) (*models.Company, error)
	UpdateCompanyPhoto(file io.Reader, fileType string, id uuid.UUID) (string, error)
	GetCompanyById(ctx context.Context, id uuid.UUID) (foundCompanyData *models.CompanyData, err error)
}

// CompanyRepo represents the repository interface for companies.
type CompanyRepo interface {
	CreateCompany(ctx context.Context, company *models.Company) (*models.Company, error)
	UpdateCompanyPhoto(id uuid.UUID, fileName string) (string, error)
	GetCompanyById(ctx context.Context, companyId uuid.UUID) (*models.CompanyData, error)
}
