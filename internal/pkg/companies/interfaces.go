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
}

// CompanyRepo represents the repository interface for companies.
type CompanyRepo interface {
	CreateCompany(ctx context.Context, company *models.Company) (*models.Company, error)
	UpdateCompanyPhoto(id uuid.UUID, fileName string) (string, error)
}
