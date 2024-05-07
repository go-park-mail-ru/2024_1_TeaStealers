//go:generate mockgen -destination=mock/${GOFILE} -package=${GOPACKAGE}_mock -source=${GOFILE}
package companies

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"io"
)

// CompanyUsecase represents the usecase interface for companies.
type CompanyUsecase interface {
	CreateCompany(ctx context.Context, data *models.CompanyCreateData) (*models.Company, error)
	UpdateCompanyPhoto(file io.Reader, fileType string, id int64) (string, error)
	GetCompanyById(ctx context.Context, id int64) (foundCompanyData *models.CompanyData, err error)
}

// CompanyRepo represents the repository interface for companies.
type CompanyRepo interface {
	CreateCompany(ctx context.Context, company *models.Company) (*models.Company, error)
	UpdateCompanyPhoto(id int64, fileName string) (string, error)
	GetCompanyById(ctx context.Context, companyId int64) (*models.CompanyData, error)
}
