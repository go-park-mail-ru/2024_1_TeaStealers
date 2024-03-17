package adverts

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"
)

// AdvertUsecase represents the usecase interface for adverts.
type AdvertUsecase interface {
	CreateFlatAdvert(context.Context, *models.AdvertFlatCreateData) (*models.Advert, error)
	CreateHouseAdvert(context.Context, *models.AdvertHouseCreateData) (*models.Advert, error)
	GetHouseSquareAdvertsList(ctx context.Context) (foundAdverts []*models.AdvertSquareData, err error)
	GetFlatSquareAdvertsList(ctx context.Context) (foundAdverts []*models.AdvertSquareData, err error)
}

// AdvertRepo represents the repository interface for adverts.
type AdvertRepo interface {
	BeginTx(ctx context.Context) (*sql.Tx, error)
	CreateAdvertType(ctx context.Context, tx *sql.Tx, newAdvertType *models.AdvertType) error
	CreateAdvert(ctx context.Context, tx *sql.Tx, newAdvert *models.Advert) error
	CreatePriceChange(ctx context.Context, tx *sql.Tx, newPriceChange *models.PriceChange) error
	CreateBuilding(ctx context.Context, tx *sql.Tx, newBuilding *models.Building) error
	CreateHouse(ctx context.Context, tx *sql.Tx, newHouse *models.House) error
	CreateFlat(ctx context.Context, tx *sql.Tx, newFlat *models.Flat) error
	CheckExistsBuilding(ctx context.Context, adress string) (*models.Building, error)
	GetHouseSquareAdvertsList(ctx context.Context) ([]*models.AdvertSquareData, error)
	GetFlatSquareAdvertsList(ctx context.Context) ([]*models.AdvertSquareData, error)
}
