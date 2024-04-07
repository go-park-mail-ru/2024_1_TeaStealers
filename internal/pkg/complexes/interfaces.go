//go:generate mockgen -destination=mock/${GOFILE} -package=${GOPACKAGE}_mock -source=${GOFILE}
package complexes

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"
	"io"

	"github.com/satori/uuid"
)

// ComplexUsecase represents the usecase interface for complexes.
type ComplexUsecase interface {
	CreateComplex(ctx context.Context, data *models.ComplexCreateData) (*models.Complex, error)
	CreateBuilding(ctx context.Context, data *models.BuildingCreateData) (*models.Building, error)
	UpdateComplexPhoto(file io.Reader, fileType string, id uuid.UUID) (string, error)
	GetComplexById(ctx context.Context, id uuid.UUID) (foundComplex *models.ComplexData, err error)
	CreateFlatAdvert(ctx context.Context, data *models.ComplexAdvertFlatCreateData) (*models.Advert, error)
	CreateHouseAdvert(ctx context.Context, data *models.ComplexAdvertHouseCreateData) (*models.Advert, error)
}

// ComplexRepo represents the repository interface for complexes.
type ComplexRepo interface {
	CreateComplex(ctx context.Context, company *models.Complex) (*models.Complex, error)
	CreateBuilding(ctx context.Context, complex *models.Building) (*models.Building, error)
	UpdateComplexPhoto(id uuid.UUID, fileName string) (string, error)
	GetComplexById(ctx context.Context, complexId uuid.UUID) (*models.ComplexData, error)
	BeginTx(ctx context.Context) (*sql.Tx, error)
	CreateAdvertType(ctx context.Context, tx *sql.Tx, newAdvertType *models.AdvertType) error
	CreateAdvert(ctx context.Context, tx *sql.Tx, newAdvert *models.Advert) error
	CreatePriceChange(ctx context.Context, tx *sql.Tx, newPriceChange *models.PriceChange) error
	CreateHouse(ctx context.Context, tx *sql.Tx, newHouse *models.House) error
	CreateFlat(ctx context.Context, tx *sql.Tx, newFlat *models.Flat) error
}
