//go:generate mockgen -destination=mock/${GOFILE} -package=${GOPACKAGE}_mock -source=${GOFILE}
package adverts

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"

	"github.com/satori/uuid"
)

// AdvertUsecase represents the usecase interface for adverts.
type AdvertUsecase interface {
	CreateFlatAdvert(context.Context, *models.AdvertFlatCreateData) (*models.Advert, error)
	CreateHouseAdvert(context.Context, *models.AdvertHouseCreateData) (*models.Advert, error)
	GetAdvertById(ctx context.Context, id uuid.UUID) (foundAdvert *models.AdvertData, err error)
	GetSquareAdvertsList(ctx context.Context, pageSize, offset int) (foundAdverts []*models.AdvertSquareData, err error)
	GetRectangleAdvertsList(ctx context.Context, advertFilter models.AdvertFilter) (foundAdverts *models.AdvertDataPage, err error)
	GetRectangleAdvertsByUserId(ctx context.Context, pageSize, offset int, userId uuid.UUID) (foundAdverts []*models.AdvertRectangleData, err error)
	UpdateAdvertById(ctx context.Context, advertUpdateData *models.AdvertUpdateData) (err error)
	DeleteAdvertById(ctx context.Context, advertId uuid.UUID) (err error)
	GetRectangleAdvertsByComplexId(ctx context.Context, pageSize, offset int, comlexId uuid.UUID) (foundAdverts []*models.AdvertRectangleData, err error)
	GetExistBuildingsByAddress(ctx context.Context, address string, pageSize int) (foundBuildings []*models.BuildingsExistData, err error)
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
	GetHouseAdvertById(ctx context.Context, id uuid.UUID) (*models.AdvertData, error)
	GetFlatAdvertById(ctx context.Context, id uuid.UUID) (*models.AdvertData, error)
	GetTypeAdvertById(ctx context.Context, id uuid.UUID) (*models.AdvertTypeAdvert, error)
	GetSquareAdverts(ctx context.Context, pageSize, offset int) ([]*models.AdvertSquareData, error)
	GetRectangleAdverts(ctx context.Context, advertFilter models.AdvertFilter) (*models.AdvertDataPage, error)
	GetRectangleAdvertsByUserId(ctx context.Context, pageSize, offset int, userId uuid.UUID) ([]*models.AdvertRectangleData, error)
	UpdateFlatAdvertById(ctx context.Context, tx *sql.Tx, advertUpdateData *models.AdvertUpdateData) error
	UpdateHouseAdvertById(ctx context.Context, tx *sql.Tx, advertUpdateData *models.AdvertUpdateData) error
	ChangeTypeAdvert(ctx context.Context, tx *sql.Tx, advertId uuid.UUID) error
	DeleteHouseAdvertById(ctx context.Context, tx *sql.Tx, advertId uuid.UUID) error
	DeleteFlatAdvertById(ctx context.Context, tx *sql.Tx, advertId uuid.UUID) error
	GetRectangleAdvertsByComplexId(ctx context.Context, pageSize, offset int, complexId uuid.UUID) ([]*models.AdvertRectangleData, error)
	CheckExistsBuildings(ctx context.Context, pageSize int, adress string) ([]*models.BuildingsExistData, error)
	SelectImages(advertId uuid.UUID) ([]*models.ImageResp, error)
}
