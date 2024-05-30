//go:generate mockgen -destination=mock/${GOFILE} -package=${GOPACKAGE}_mock -source=${GOFILE}
package complex

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"github.com/jackc/pgx/v4"
)

// ComplexUsecase represents the usecase interface for complexes.
type ComplexUsecase interface {
	CreateComplex(ctx context.Context, data *models.ComplexCreateData) (*models.Complex, error)
	UpdateComplexPhoto(ctx context.Context, id int64, filename string) (string, error)
	GetComplexById(ctx context.Context, id int64) (foundComplex *models.ComplexData, err error)
	CreateCompany(ctx context.Context, data *models.CompanyCreateData) (*models.Company, error)
	UpdateCompanyPhoto(ctx context.Context, id int64, filename string) (string, error)
	GetCompanyById(ctx context.Context, id int64) (foundCompanyData *models.CompanyData, err error)
}

// ComplexRepo represents the repository interface for complexes.
type ComplexRepo interface {
	CreateComplex(ctx context.Context, company *models.Complex) (*models.Complex, error)
	UpdateComplexPhoto(ctx context.Context, id int64, fileName string) (string, error)
	GetComplexById(ctx context.Context, complexId int64) (*models.ComplexData, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	CreateCompany(ctx context.Context, company *models.Company) (*models.Company, error)
	UpdateCompanyPhoto(ctx context.Context, id int64, fileName string) (string, error)
	GetCompanyById(ctx context.Context, companyId int64) (*models.CompanyData, error)
}
