package adverts

import (
	"2024_1_TeaStealers/internal/models"
	"context"
)

// AdvertUsecase represents the usecase interface for adverts.
type AdvertUsecase interface {
	CreateFlatAdvert(context.Context, *models.AdvertFlatCreateData) (*models.Advert, error)
	CreateHouseAdvert(context.Context, *models.AdvertHouseCreateData) (*models.Advert, error)
}

// AdvertRepo represents the repository interface for adverts.
type AdvertRepo interface {
	CreateAdvertType(ctx context.Context, newAdvertType *models.AdvertType) error
	CreateAdvert(ctx context.Context, newAdvert *models.Advert) error
	CreatePriceChange(ctx context.Context, newPriceChange *models.PriceChange) error
	CreateBuilding(ctx context.Context, newBuilding *models.Building) error
	CreateHouse(ctx context.Context, newHouse *models.House) error
	CreateFlat(ctx context.Context, newFlat *models.Flat) error
}
