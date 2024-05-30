package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts"
	"2024_1_TeaStealers/internal/pkg/utils"
	"context"
	"github.com/jackc/pgx/v4"

	"go.uber.org/zap"
)

// AdvertUsecase represents the usecase for adverts using.
type AdvertUsecase struct {
	repo   adverts.AdvertRepo
	logger *zap.Logger
}

// NewAdvertUsecase creates a new instance of AdvertUsecase.
func NewAdvertUsecase(repo adverts.AdvertRepo, logger *zap.Logger) *AdvertUsecase {
	return &AdvertUsecase{repo: repo, logger: logger}
}

// CreateFlatAdvert handles the creation advert process.
func (u *AdvertUsecase) CreateFlatAdvert(ctx context.Context, data *models.AdvertFlatCreateData) (*models.Advert, error) {
	tx, err := u.repo.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return nil, err
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		}
	}()

	building, err := u.repo.CheckExistsBuilding(ctx, &data.Address)
	if err != nil {

		id, err := u.repo.CreateProvince(ctx, tx, data.Address.Province)
		if err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
			return nil, err
		}

		id, err = u.repo.CreateTown(ctx, tx, id, data.Address.Town)
		if err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
			return nil, err
		}

		id, err = u.repo.CreateStreet(ctx, tx, id, data.Address.Street)
		if err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
			return nil, err
		}

		id, err = u.repo.CreateHouseAddress(ctx, tx, id, data.Address.House)
		if err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
			return nil, err
		}

		id, err = u.repo.CreateAddress(ctx, tx, id, data.Address.Metro, data.Address.AddressPoint)
		if err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
			return nil, err
		}

		building = &models.Building{
			Floor:        data.FloorGeneral,
			Material:     data.Material,
			AddressID:    id,
			YearCreation: data.YearCreation,
		}
		if building.ID, err = u.repo.CreateBuilding(ctx, tx, building); err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
			return nil, err
		}
	}

	newFlat := &models.Flat{
		BuildingID:        building.ID,
		RoomCount:         data.RoomCount,
		Floor:             data.Floor,
		CeilingHeight:     data.CeilingHeight,
		SquareGeneral:     data.SquareGeneral,
		SquareResidential: data.SquareResidential,
		Apartment:         data.Apartment,
	}

	idFlat, err := u.repo.CreateFlat(ctx, tx, newFlat)
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return nil, err
	}

	newAdvert := &models.Advert{
		UserID:         data.UserID,
		AdvertTypeSale: data.AdvertTypeSale,
		Title:          data.Title,
		Description:    data.Description,
		Phone:          data.Phone,
		IsAgent:        data.IsAgent,
		Priority:       1, // Разобраться в будущем, как это менять за деньги(money)
	}

	idAdvert, err := u.repo.CreateAdvert(ctx, tx, newAdvert)
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return nil, err
	}
	newAdvert.ID = idAdvert

	newAdvertTypeFlat := &models.FlatTypeAdvert{
		FlatID:   idFlat,
		AdvertID: idAdvert,
	}

	if err := u.repo.CreateAdvertTypeFlat(ctx, tx, newAdvertTypeFlat); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return nil, err
	}

	newPriceChange := &models.PriceChange{
		AdvertID: idAdvert,
		Price:    data.Price,
	}

	if err := u.repo.CreatePriceChange(ctx, tx, newPriceChange); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return nil, err
	}

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod)
	return newAdvert, nil
}

// CreateHouseAdvert handles the creation advert process.
func (u *AdvertUsecase) CreateHouseAdvert(ctx context.Context, data *models.AdvertHouseCreateData) (*models.Advert, error) {
	tx, err := u.repo.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return nil, err
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		}
	}()

	building, err := u.repo.CheckExistsBuilding(ctx, &data.Address)
	if err != nil {

		id, err := u.repo.CreateProvince(ctx, tx, data.Address.Province)
		if err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
			return nil, err
		}

		id, err = u.repo.CreateTown(ctx, tx, id, data.Address.Town)
		if err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
			return nil, err
		}

		id, err = u.repo.CreateStreet(ctx, tx, id, data.Address.Street)
		if err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
			return nil, err
		}

		id, err = u.repo.CreateHouseAddress(ctx, tx, id, data.Address.House)
		if err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
			return nil, err
		}

		id, err = u.repo.CreateAddress(ctx, tx, id, data.Address.Metro, data.Address.AddressPoint)
		if err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
			return nil, err
		}

		building = &models.Building{
			Floor:        data.FloorGeneral,
			Material:     data.Material,
			AddressID:    id,
			YearCreation: data.YearCreation,
		}
		if building.ID, err = u.repo.CreateBuilding(ctx, tx, building); err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
			return nil, err
		}
	}

	newHouse := &models.House{
		BuildingID:    building.ID,
		CeilingHeight: data.CeilingHeight,
		SquareArea:    data.SquareArea,
		SquareHouse:   data.SquareHouse,
		BedroomCount:  data.BedroomCount,
		StatusArea:    data.StatusArea,
		Cottage:       data.Cottage,
		StatusHome:    data.StatusHome,
	}

	idHouse, err := u.repo.CreateHouse(ctx, tx, newHouse)
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return nil, err
	}

	newAdvert := &models.Advert{
		UserID:         data.UserID,
		AdvertTypeSale: data.AdvertTypeSale,
		Title:          data.Title,
		Description:    data.Description,
		Phone:          data.Phone,
		IsAgent:        data.IsAgent,
		Priority:       1, // Разобраться в будущем, как это менять за деньги(money)
	}

	idAdvert, err := u.repo.CreateAdvert(ctx, tx, newAdvert)
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return nil, err
	}
	newAdvert.ID = idAdvert

	newAdvertTypeHouse := &models.HouseTypeAdvert{
		HouseID:  idHouse,
		AdvertID: idAdvert,
	}

	if err := u.repo.CreateAdvertTypeHouse(ctx, tx, newAdvertTypeHouse); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return nil, err
	}

	newPriceChange := &models.PriceChange{
		AdvertID: idAdvert,
		Price:    data.Price,
	}

	if err := u.repo.CreatePriceChange(ctx, tx, newPriceChange); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return nil, err
	}

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod)
	return newAdvert, nil
}

// GetAdvertById handles the getting house advert process.
func (u *AdvertUsecase) GetAdvertById(ctx context.Context, id int64) (foundAdvert *models.AdvertData, err error) {
	var typeAdvert *models.AdvertTypeAdvert
	typeAdvert, err = u.repo.GetTypeAdvertById(ctx, id)
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetAdvertByIdMethod, err)
		return nil, err
	}

	switch *typeAdvert {
	case models.AdvertTypeFlat:
		if foundAdvert, err = u.repo.GetFlatAdvertById(ctx, id); err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetAdvertByIdMethod, err)
			return nil, err
		}
	case models.AdvertTypeHouse:
		if foundAdvert, err = u.repo.GetHouseAdvertById(ctx, id); err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetAdvertByIdMethod, err)
			return nil, err
		}
	}

	var foundImages []*models.ImageResp
	if foundImages, err = u.repo.SelectImages(ctx, foundAdvert.ID); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetAdvertByIdMethod, err)
		return nil, err
	}

	foundAdvert.Images = foundImages

	var priceChanges []*models.PriceChangeData
	if priceChanges, err = u.repo.SelectPriceChanges(ctx, foundAdvert.ID); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetAdvertByIdMethod, err)
		return nil, err
	}

	foundAdvert.PriceChange = priceChanges

	if foundAdvert.CountLikes, err = u.repo.SelectCountLikes(ctx, foundAdvert.ID); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetAdvertByIdMethod, err)
		return nil, err
	}

	if foundAdvert.CountViews, err = u.repo.SelectCountViews(ctx, foundAdvert.ID); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetAdvertByIdMethod, err)
		return nil, err
	}

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetAdvertByIdMethod)
	return foundAdvert, nil
}

// UpdateAdvertById handles the updating advert process.
func (u *AdvertUsecase) UpdateAdvertById(ctx context.Context, advertUpdateData *models.AdvertUpdateData) (err error) {
	typeAdvert := advertUpdateData.TypeAdvert
	tx, err := u.repo.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.UpdateAdvertByIdMethod, err)
		return err
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.UpdateAdvertByIdMethod, err)
		}
	}()

	if typeAdvertOld, err := u.repo.GetTypeAdvertById(ctx, advertUpdateData.ID); err == nil {
		if *typeAdvertOld != models.AdvertTypeAdvert(typeAdvert) {

			if err = u.repo.ChangeTypeAdvert(ctx, tx, advertUpdateData.ID); err != nil {
				utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.UpdateAdvertByIdMethod, err)
				return err
			}
		}
	} else {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.UpdateAdvertByIdMethod, err)
		return err
	}

	switch models.AdvertTypeAdvert(typeAdvert) {
	case models.AdvertTypeFlat:
		if err = u.repo.UpdateFlatAdvertById(ctx, tx, advertUpdateData); err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.UpdateAdvertByIdMethod, err)
			return err
		}
	case models.AdvertTypeHouse:
		if err = u.repo.UpdateHouseAdvertById(ctx, tx, advertUpdateData); err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.UpdateAdvertByIdMethod, err)
			return err
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.UpdateAdvertByIdMethod, err)
		return err
	}

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.UpdateAdvertByIdMethod)
	return nil
}

// DeleteAdvertById handles the deleting advert process.
func (u *AdvertUsecase) DeleteAdvertById(ctx context.Context, advertId int64) (err error) {
	tx, err := u.repo.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.DeleteAdvertByIdMethod, err)
		return err
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.DeleteAdvertByIdMethod, err)
		}
	}()

	typeAdvert, err := u.repo.GetTypeAdvertById(ctx, advertId)
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.DeleteAdvertByIdMethod, err)
		return err
	}

	switch *typeAdvert {
	case models.AdvertTypeFlat:
		if err = u.repo.DeleteFlatAdvertById(ctx, tx, advertId); err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.DeleteAdvertByIdMethod, err)
			return err
		}
	case models.AdvertTypeHouse:
		if err = u.repo.DeleteHouseAdvertById(ctx, tx, advertId); err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.DeleteAdvertByIdMethod, err)
			return err
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.DeleteAdvertByIdMethod, err)
		return err
	}

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.DeleteAdvertByIdMethod)
	return nil
}

// GetSquareAdvertsList handles the square adverts getting process with paggination.
func (u *AdvertUsecase) GetSquareAdvertsList(ctx context.Context, pageSize, offset int) (foundAdverts []*models.AdvertSquareData, err error) {
	if foundAdverts, err = u.repo.GetSquareAdverts(ctx, pageSize, offset); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetSquareAdvertsListMethod, err)
		return nil, err
	}

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetSquareAdvertsListMethod)
	return foundAdverts, nil
}

// GetRectangleAdvertsList handles the rectangle adverts getting process with paggination and search.
func (u *AdvertUsecase) GetRectangleAdvertsList(ctx context.Context, advertFilter models.AdvertFilter) (foundAdverts *models.AdvertDataPage, err error) {
	if foundAdverts, err = u.repo.GetRectangleAdverts(ctx, advertFilter); err != nil {
		// utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsListMethod, err)
		return nil, err
	}

	// utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsListMethod)
	return foundAdverts, nil
}

// UpdatePriority handles the updating advert priority.
func (u *AdvertUsecase) UpdatePriority(ctx context.Context, advertId int64, priority int64) (newPriority int64, err error) {
	tx, err := u.repo.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.DeleteAdvertByIdMethod, err)
		return 0, err
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.DeleteAdvertByIdMethod, err)
		}
	}()

	if newPriority, err = u.repo.UpdatePriority(ctx, tx, advertId, priority); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsListMethod, err)
		return 0, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsListMethod, err)
		return 0, err
	}

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsListMethod)
	return newPriority, nil
}

// GetPriority handles the getting advert priority.
func (u *AdvertUsecase) GetPriority(ctx context.Context, advertId int64) (priority int64, err error) {
	tx, err := u.repo.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.DeleteAdvertByIdMethod, err)
		return 0, err
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.DeleteAdvertByIdMethod, err)
		}
	}()

	if priority, err = u.repo.GetPriority(ctx, tx, advertId); err != nil {
		// utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsListMethod, err)
		return 0, err
	}
	if err = tx.Commit(ctx); err != nil {
		return 0, nil
	}

	// utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsListMethod)
	return priority, nil
}

// GetExistBuildingByAddress handles the buildings getting process by address with paggination.
func (u *AdvertUsecase) GetExistBuildingByAddress(ctx context.Context, address *models.AddressData) (foundBuilding *models.BuildingData, err error) {
	if foundBuilding, err = u.repo.CheckExistsBuildingData(ctx, address); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetExistBuildingsByAddressMethod, err)
		return nil, err
	}

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetExistBuildingsByAddressMethod)
	return foundBuilding, nil
}

// GetRectangleAdvertsByUserId handles the rectangle adverts getting process with paggination by userId.
func (u *AdvertUsecase) GetRectangleAdvertsByUserId(ctx context.Context, pageSize, offset int, userId int64) (foundAdverts []*models.AdvertRectangleData, err error) {
	if foundAdverts, err = u.repo.GetRectangleAdvertsByUserId(ctx, pageSize, offset, userId); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
		return nil, err
	}

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsByUserIdMethod)
	return foundAdverts, nil
}

// GetRectangleAdvertsLikedByUserId handles the rectangle adverts getting process with paggination by userId.
func (u *AdvertUsecase) GetRectangleAdvertsLikedByUserId(ctx context.Context, pageSize, offset int, userId int64) (foundAdverts []*models.AdvertRectangleData, err error) {
	if foundAdverts, err = u.repo.GetRectangleAdvertsLikedByUserId(ctx, pageSize, offset, userId); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
		return nil, err
	}

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsByUserIdMethod)
	return foundAdverts, nil
}

// GetRectangleAdvertsByComplexId handles the rectangle adverts getting process with paggination by complexId.
func (u *AdvertUsecase) GetRectangleAdvertsByComplexId(ctx context.Context, pageSize, offset int, comlexId int64) (foundAdverts []*models.AdvertRectangleData, err error) {
	if foundAdverts, err = u.repo.GetRectangleAdvertsByComplexId(ctx, pageSize, offset, comlexId); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsByComplexIdMethod, err)
		return nil, err
	}

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsByComplexIdMethod)
	return foundAdverts, nil
}

func (u *AdvertUsecase) LikeAdvert(ctx context.Context, advertId int64, userId int64) error {
	if err := u.repo.LikeAdvert(ctx, advertId, userId); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsByComplexIdMethod, err)
		return err
	}

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsByComplexIdMethod)
	return nil
}

func (u *AdvertUsecase) DislikeAdvert(ctx context.Context, advertId int64, userId int64) error {
	if err := u.repo.DislikeAdvert(ctx, advertId, userId); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsByComplexIdMethod, err)
		return err
	}

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsByComplexIdMethod)
	return nil
}
