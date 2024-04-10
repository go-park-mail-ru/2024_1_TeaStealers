//go:generate mockgen -destination=mock/${GOFILE} -package=${GOPACKAGE}_mock -source=${GOFILE}
package adverts

import (
	"2024_1_TeaStealers/internal/models"
	"context"

	"github.com/satori/uuid"
)

const (
	CreateFlatAdvertMethod               = "CreateFlatAdvert"
	CreateHouseAdvertMethod              = "CreateHouseAdvert"
	GetAdvertByIdMethod                  = "GetAdvertById"
	GetSquareAdvertsListMethod           = "GetSquareAdvertsList"
	GetRectangleAdvertsListMethod        = "GetRectangleAdvertsList"
	GetRectangleAdvertsByUserIdMethod    = "GetRectangleAdvertsByUserId"
	UpdateAdvertByIdMethod               = "UpdateAdvertById"
	DeleteAdvertByIdMethod               = "DeleteAdvertById"
	GetRectangleAdvertsByComplexIdMethod = "GetRectangleAdvertsByComplexId"
	GetExistBuildingsByAddressMethod     = "GetExistBuildingsByAddress"
	BeginTxMethod                        = "BeginTx"
	CreateAdvertTypeMethod               = "CreateAdvertType"
	CreateAdvertMethod                   = "CreateAdvert"
	CreatePriceChangeMethod              = "CreatePriceChange"
	CreateBuildingMethod                 = "CreateBuilding"
	CreateHouseMethod                    = "CreateHouse"
	CreateFlatMethod                     = "CreateFlat"
	CheckExistsBuildingMethod            = "CheckExistsBuilding"
	GetHouseAdvertByIdMethod             = "GetHouseAdvertById"
	GetFlatAdvertByIdMethod              = "GetFlatAdvertById"
	GetTypeAdvertByIdMethod              = "GetTypeAdvertById"
	GetSquareAdvertsMethod               = "GetSquareAdverts"
	GetRectangleAdvertsMethod            = "GetRectangleAdverts"
	UpdateFlatAdvertByIdMethod           = "UpdateFlatAdvertById"
	UpdateHouseAdvertByIdMethod          = "UpdateHouseAdvertById"
	ChangeTypeAdvertMethod               = "ChangeTypeAdvert"
	DeleteHouseAdvertByIdMethod          = "DeleteHouseAdvertById"
	DeleteFlatAdvertByIdMethod           = "DeleteFlatAdvertById"
	CheckExistsBuildingsMethod           = "CheckExistsBuildings"
	SelectImagesMethod                   = "SelectImages"
	CheckExistsFlatMethod                = "CheckExistsFlat"
	CheckExistsHouseMethod               = "CheckExistsHouse"
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
	GetExistBuildingsByAddress(ctx context.Context, address string, pageSize int) (foundBuildings []*models.BuildingData, err error)
}

// AdvertRepo represents the repository interface for adverts.
type AdvertRepo interface {
	BeginTx(ctx context.Context) (models.Transaction, error)
	CreateAdvertType(ctx context.Context, tx models.Transaction, newAdvertType *models.AdvertType) error
	CreateAdvert(ctx context.Context, tx models.Transaction, newAdvert *models.Advert) error
	CreatePriceChange(ctx context.Context, tx models.Transaction, newPriceChange *models.PriceChange) error
	CreateBuilding(ctx context.Context, tx models.Transaction, newBuilding *models.Building) error
	CreateHouse(ctx context.Context, tx models.Transaction, newHouse *models.House) error
	CreateFlat(ctx context.Context, tx models.Transaction, newFlat *models.Flat) error
	CheckExistsBuilding(ctx context.Context, adress string) (*models.Building, error)
	GetHouseAdvertById(ctx context.Context, id uuid.UUID) (*models.AdvertData, error)
	GetFlatAdvertById(ctx context.Context, id uuid.UUID) (*models.AdvertData, error)
	GetTypeAdvertById(ctx context.Context, id uuid.UUID) (*models.AdvertTypeAdvert, error)
	GetSquareAdverts(ctx context.Context, pageSize, offset int) ([]*models.AdvertSquareData, error)
	GetRectangleAdverts(ctx context.Context, advertFilter models.AdvertFilter) (*models.AdvertDataPage, error)
	GetRectangleAdvertsByUserId(ctx context.Context, pageSize, offset int, userId uuid.UUID) ([]*models.AdvertRectangleData, error)
	UpdateFlatAdvertById(ctx context.Context, tx models.Transaction, advertUpdateData *models.AdvertUpdateData) error
	UpdateHouseAdvertById(ctx context.Context, tx models.Transaction, advertUpdateData *models.AdvertUpdateData) error
	ChangeTypeAdvert(ctx context.Context, tx models.Transaction, advertId uuid.UUID) error
	DeleteHouseAdvertById(ctx context.Context, tx models.Transaction, advertId uuid.UUID) error
	DeleteFlatAdvertById(ctx context.Context, tx models.Transaction, advertId uuid.UUID) error
	GetRectangleAdvertsByComplexId(ctx context.Context, pageSize, offset int, complexId uuid.UUID) ([]*models.AdvertRectangleData, error)
	CheckExistsBuildings(ctx context.Context, pageSize int, adress string) ([]*models.BuildingData, error)
	SelectImages(ctx context.Context, advertId uuid.UUID) ([]*models.ImageResp, error)
}
