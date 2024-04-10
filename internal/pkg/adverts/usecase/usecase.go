package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts"
	"2024_1_TeaStealers/internal/pkg/utils"
	"context"

	"github.com/satori/uuid"
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
	tx, err := u.repo.BeginTx(ctx)

	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateFlatAdvertMethod, err)
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateFlatAdvertMethod, err)
		}
	}()

	newAdvertType := &models.AdvertType{
		ID:         uuid.NewV4(),
		AdvertType: data.AdvertTypePlacement,
	}

	newAdvert := &models.Advert{
		ID:             uuid.NewV4(),
		UserID:         data.UserID,
		AdvertTypeID:   newAdvertType.ID,
		AdvertTypeSale: data.AdvertTypeSale,
		Title:          data.Title,
		Description:    data.Description,
		Phone:          data.Phone,
		IsAgent:        data.IsAgent,
		Priority:       1, // Разобраться в будущем, как это менять за деньги(money)
	}

	building, err := u.repo.CheckExistsBuilding(ctx, data.Address)
	if err != nil {
		building = &models.Building{
			ID:           uuid.NewV4(),
			Floor:        data.FloorGeneral,
			Material:     data.Material,
			Address:      data.Address,
			AddressPoint: data.AddressPoint,
			YearCreation: data.YearCreation,
		}
		if err := u.repo.CreateBuilding(ctx, tx, building); err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateFlatAdvertMethod, err)
			return nil, err
		}
	}

	newFlat := &models.Flat{
		ID:                uuid.NewV4(),
		BuildingID:        building.ID,
		AdvertTypeID:      newAdvertType.ID,
		RoomCount:         data.RoomCount,
		Floor:             data.Floor,
		CeilingHeight:     data.CeilingHeight,
		SquareGeneral:     data.SquareGeneral,
		SquareResidential: data.SquareResidential,
		Apartment:         data.Apartment,
	}

	newPriceChange := &models.PriceChange{
		ID:       uuid.NewV4(),
		AdvertID: newAdvert.ID,
		Price:    data.Price,
	}

	if err := u.repo.CreateAdvertType(ctx, tx, newAdvertType); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateFlatAdvertMethod, err)
		return nil, err
	}

	if err := u.repo.CreateFlat(ctx, tx, newFlat); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateFlatAdvertMethod, err)
		return nil, err
	}

	if err := u.repo.CreateAdvert(ctx, tx, newAdvert); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateFlatAdvertMethod, err)
		return nil, err
	}

	if err := u.repo.CreatePriceChange(ctx, tx, newPriceChange); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateFlatAdvertMethod, err)
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateFlatAdvertMethod, err)
		return nil, err
	}

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateFlatAdvertMethod)
	return newAdvert, nil
}

// CreateFlatAdvert handles the creation advert process.
func (u *AdvertUsecase) CreateHouseAdvert(ctx context.Context, data *models.AdvertHouseCreateData) (*models.Advert, error) {
	tx, err := u.repo.BeginTx(ctx)
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return nil, err
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		}
	}()

	newAdvertType := &models.AdvertType{
		ID:         uuid.NewV4(),
		AdvertType: data.AdvertTypePlacement,
	}

	newAdvert := &models.Advert{
		ID:             uuid.NewV4(),
		UserID:         data.UserID,
		AdvertTypeID:   newAdvertType.ID,
		AdvertTypeSale: data.AdvertTypeSale,
		Title:          data.Title,
		Description:    data.Description,
		Phone:          data.Phone,
		IsAgent:        data.IsAgent,
		Priority:       1, // Разобраться в будущем, как это менять за деньги(money)
	}

	building, err := u.repo.CheckExistsBuilding(ctx, data.Address)
	if err != nil {
		building = &models.Building{
			ID:           uuid.NewV4(),
			Floor:        data.FloorGeneral,
			Material:     data.Material,
			Address:      data.Address,
			AddressPoint: data.AddressPoint,
			YearCreation: data.YearCreation,
		}
		if err := u.repo.CreateBuilding(ctx, tx, building); err != nil {
			utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
			return nil, err
		}
	}

	newHouse := &models.House{
		ID:            uuid.NewV4(),
		BuildingID:    building.ID,
		AdvertTypeID:  newAdvertType.ID,
		CeilingHeight: data.CeilingHeight,
		SquareArea:    data.SquareArea,
		SquareHouse:   data.SquareHouse,
		BedroomCount:  data.BedroomCount,
		StatusArea:    data.StatusArea,
		Cottage:       data.Cottage,
		StatusHome:    data.StatusHome,
	}

	newPriceChange := &models.PriceChange{
		ID:       uuid.NewV4(),
		AdvertID: newAdvert.ID,
		Price:    data.Price,
	}

	if err := u.repo.CreateAdvertType(ctx, tx, newAdvertType); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return nil, err
	}

	if err := u.repo.CreateHouse(ctx, tx, newHouse); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return nil, err
	}

	if err := u.repo.CreateAdvert(ctx, tx, newAdvert); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return nil, err
	}

	if err := u.repo.CreatePriceChange(ctx, tx, newPriceChange); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return nil, err
	}

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod)
	return newAdvert, nil
}

// GetAdvertById handles the getting house advert process.
func (u *AdvertUsecase) GetAdvertById(ctx context.Context, id uuid.UUID) (foundAdvert *models.AdvertData, err error) {
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

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetAdvertByIdMethod)
	return foundAdvert, nil
}

// UpdateAdvertById handles the updating advert process.
func (u *AdvertUsecase) UpdateAdvertById(ctx context.Context, advertUpdateData *models.AdvertUpdateData) (err error) {
	typeAdvert := advertUpdateData.TypeAdvert
	tx, err := u.repo.BeginTx(ctx)
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.UpdateAdvertByIdMethod, err)
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
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
	err = tx.Commit()
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.UpdateAdvertByIdMethod, err)
		return err
	}

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.UpdateAdvertByIdMethod)
	return nil
}

// DeleteAdvertById handles the deleting advert process.
func (u *AdvertUsecase) DeleteAdvertById(ctx context.Context, advertId uuid.UUID) (err error) {
	tx, err := u.repo.BeginTx(ctx)
	if err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.DeleteAdvertByIdMethod, err)
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
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
	err = tx.Commit()
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
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsListMethod, err)
		return nil, err
	}

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsListMethod)
	return foundAdverts, nil
}

// GetExistBuildingsByAddress handles the buildings getting process by address with paggination.
func (u *AdvertUsecase) GetExistBuildingsByAddress(ctx context.Context, address string, pageSize int) (foundBuildings []*models.BuildingData, err error) {
	if foundBuildings, err = u.repo.CheckExistsBuildings(ctx, pageSize, address); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetExistBuildingsByAddressMethod, err)
		return nil, err
	}

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetExistBuildingsByAddressMethod)
	return foundBuildings, nil
}

// GetRectangleAdvertsByUserId handles the rectangle adverts getting process with paggination by userId.
func (u *AdvertUsecase) GetRectangleAdvertsByUserId(ctx context.Context, pageSize, offset int, userId uuid.UUID) (foundAdverts []*models.AdvertRectangleData, err error) {
	if foundAdverts, err = u.repo.GetRectangleAdvertsByUserId(ctx, pageSize, offset, userId); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
		return nil, err
	}

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsByUserIdMethod)
	return foundAdverts, nil
}

// GetRectangleAdvertsByComplexId handles the rectangle adverts getting process with paggination by complexId.
func (u *AdvertUsecase) GetRectangleAdvertsByComplexId(ctx context.Context, pageSize, offset int, comlexId uuid.UUID) (foundAdverts []*models.AdvertRectangleData, err error) {
	if foundAdverts, err = u.repo.GetRectangleAdvertsByComplexId(ctx, pageSize, offset, comlexId); err != nil {
		utils.LogError(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsByComplexIdMethod, err)
		return nil, err
	}

	utils.LogSucces(u.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.GetRectangleAdvertsByComplexIdMethod)
	return foundAdverts, nil
}
