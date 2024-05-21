//go:generate mockgen -destination=mock/${GOFILE} -package=${GOPACKAGE}_mock -source=${GOFILE}
package adverts

import (
	"2024_1_TeaStealers/internal/models"
	"context"
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
	GetAdvertById(ctx context.Context, id int64) (foundAdvert *models.AdvertData, err error)
	GetSquareAdvertsList(ctx context.Context, pageSize, offset int) (foundAdverts []*models.AdvertSquareData, err error)
	GetRectangleAdvertsList(ctx context.Context, advertFilter models.AdvertFilter) (foundAdverts *models.AdvertDataPage, err error)
	GetRectangleAdvertsByUserId(ctx context.Context, pageSize, offset int, userId int64) (foundAdverts []*models.AdvertRectangleData, err error)
	UpdateAdvertById(ctx context.Context, advertUpdateData *models.AdvertUpdateData) (err error)
	DeleteAdvertById(ctx context.Context, advertId int64) (err error)
	GetRectangleAdvertsByComplexId(ctx context.Context, pageSize, offset int, comlexId int64) (foundAdverts []*models.AdvertRectangleData, err error)
	GetExistBuildingByAddress(ctx context.Context, address *models.AddressData) (foundBuilding *models.BuildingData, err error)
	LikeAdvert(ctx context.Context, advertId int64, userId int64) error
	DislikeAdvert(ctx context.Context, advertId int64, userId int64) error
	GetRectangleAdvertsLikedByUserId(ctx context.Context, pageSize, offset int, userId int64) (foundAdverts []*models.AdvertRectangleData, err error)
}

// AdvertRepo represents the repository interface for adverts.
type AdvertRepo interface {
	BeginTx(ctx context.Context) (models.Transaction, error)
	CreateAdvertTypeHouse(ctx context.Context, tx models.Transaction, newAdvertType *models.HouseTypeAdvert) error
	CreateAdvertTypeFlat(ctx context.Context, tx models.Transaction, newAdvertType *models.FlatTypeAdvert) error
	CreateAdvert(ctx context.Context, tx models.Transaction, newAdvert *models.Advert) (int64, error)
	CreatePriceChange(ctx context.Context, tx models.Transaction, newPriceChange *models.PriceChange) error
	CreateBuilding(ctx context.Context, tx models.Transaction, newBuilding *models.Building) (int64, error)
	CreateHouse(ctx context.Context, tx models.Transaction, newHouse *models.House) (int64, error)
	CreateFlat(ctx context.Context, tx models.Transaction, newFlat *models.Flat) (int64, error)
	CheckExistsBuilding(ctx context.Context, address *models.AddressData) (*models.Building, error)
	GetHouseAdvertById(ctx context.Context, id int64) (*models.AdvertData, error)
	GetFlatAdvertById(ctx context.Context, id int64) (*models.AdvertData, error)
	GetTypeAdvertById(ctx context.Context, id int64) (*models.AdvertTypeAdvert, error)
	GetSquareAdverts(ctx context.Context, pageSize, offset int) ([]*models.AdvertSquareData, error)
	GetRectangleAdverts(ctx context.Context, advertFilter models.AdvertFilter) (*models.AdvertDataPage, error)
	GetRectangleAdvertsByUserId(ctx context.Context, pageSize, offset int, userId int64) ([]*models.AdvertRectangleData, error)
	UpdateFlatAdvertById(ctx context.Context, tx models.Transaction, advertUpdateData *models.AdvertUpdateData) error
	UpdateHouseAdvertById(ctx context.Context, tx models.Transaction, advertUpdateData *models.AdvertUpdateData) error
	ChangeTypeAdvert(ctx context.Context, tx models.Transaction, advertId int64) error
	DeleteHouseAdvertById(ctx context.Context, tx models.Transaction, advertId int64) error
	DeleteFlatAdvertById(ctx context.Context, tx models.Transaction, advertId int64) error
	GetRectangleAdvertsByComplexId(ctx context.Context, pageSize, offset int, complexId int64) ([]*models.AdvertRectangleData, error)
	CheckExistsBuildingData(ctx context.Context, adress *models.AddressData) (*models.BuildingData, error)
	SelectImages(ctx context.Context, advertId int64) ([]*models.ImageResp, error)
	SelectPriceChanges(ctx context.Context, advertId int64) ([]*models.PriceChangeData, error)
	CreateAddress(ctx context.Context, tx models.Transaction, idHouse int64, metro string, address_point string) (int64, error)
	CreateHouseAddress(ctx context.Context, tx models.Transaction, idStreet int64, name string) (int64, error)
	CreateStreet(ctx context.Context, tx models.Transaction, idTown int64, name string) (int64, error)
	CreateTown(ctx context.Context, tx models.Transaction, idProvince int64, name string) (int64, error)
	CreateProvince(ctx context.Context, tx models.Transaction, name string) (int64, error)
	LikeAdvert(ctx context.Context, advertId int64, userId int64) error
	DislikeAdvert(ctx context.Context, advertId int64, userId int64) error
	GetRectangleAdvertsLikedByUserId(ctx context.Context, pageSize, offset int, userId int64) ([]*models.AdvertRectangleData, error)
	SelectCountLikes(ctx context.Context, id int64) (int64, error)
	SelectCountViews(ctx context.Context, id int64) (int64, error)
}
