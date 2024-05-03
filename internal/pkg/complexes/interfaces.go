//go:generate mockgen -destination=mock/${GOFILE} -package=${GOPACKAGE}_mock -source=${GOFILE}
package complexes

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"io"
)

// ComplexUsecase represents the usecase interface for complexes.
type ComplexUsecase interface {
	CreateComplex(ctx context.Context, data *models.ComplexCreateData) (*models.Complex, error)
	CreateBuilding(ctx context.Context, data *models.BuildingCreateData) (*models.Building, error)
	UpdateComplexPhoto(file io.Reader, fileType string, id int64) (string, error)
	GetComplexById(ctx context.Context, id int64) (foundComplex *models.ComplexData, err error)
	CreateFlatAdvert(ctx context.Context, data *models.ComplexAdvertFlatCreateData) (*models.Advert, error)
	CreateHouseAdvert(ctx context.Context, data *models.ComplexAdvertHouseCreateData) (*models.Advert, error)
}

// ComplexRepo represents the repository interface for complexes.
type ComplexRepo interface {
	CreateComplex(ctx context.Context, company *models.Complex) (*models.Complex, error)
	CreateBuilding(ctx context.Context, complex *models.Building) (*models.Building, error)
	UpdateComplexPhoto(id int64, fileName string) (string, error)
	GetComplexById(ctx context.Context, complexId int64) (*models.ComplexData, error)
	BeginTx(ctx context.Context) (models.Transaction, error)
	CreateAdvertTypeHouse(ctx context.Context, tx models.Transaction, newAdvertType *models.HouseTypeAdvert) error
	CreateAdvertTypeFlat(ctx context.Context, tx models.Transaction, newAdvertType *models.FlatTypeAdvert) error
	CreateAdvert(ctx context.Context, tx models.Transaction, newAdvert *models.Advert) (int64, error)
	CreatePriceChange(ctx context.Context, tx models.Transaction, newPriceChange *models.PriceChange) error
	CreateHouse(ctx context.Context, tx models.Transaction, newHouse *models.House) (int64, error)
	CreateFlat(ctx context.Context, tx models.Transaction, newFlat *models.Flat) (int64, error)
}
