package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts"
	"context"

	"github.com/satori/uuid"
)

// AdvertUsecase represents the usecase for adverts using.
type AdvertUsecase struct {
	repo adverts.AdvertRepo
}

// NewAdvertUsecase creates a new instance of AdvertUsecase.
func NewAdvertUsecase(repo adverts.AdvertRepo) *AdvertUsecase {
	return &AdvertUsecase{repo: repo}
}

// CreateFlatAdvert handles the creation advert process.
func (u *AdvertUsecase) CreateFlatAdvert(ctx context.Context, data *models.AdvertFlatCreateData) (*models.Advert, error) {
	tx, err := u.repo.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
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
		return nil, err
	}

	if err := u.repo.CreateFlat(ctx, tx, newFlat); err != nil {
		return nil, err
	}

	if err := u.repo.CreateAdvert(ctx, tx, newAdvert); err != nil {
		return nil, err
	}

	if err := u.repo.CreatePriceChange(ctx, tx, newPriceChange); err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return newAdvert, nil
}

// CreateFlatAdvert handles the creation advert process.
func (u *AdvertUsecase) CreateHouseAdvert(ctx context.Context, data *models.AdvertHouseCreateData) (*models.Advert, error) {
	tx, err := u.repo.BeginTx(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
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
		return nil, err
	}

	if err := u.repo.CreateHouse(ctx, tx, newHouse); err != nil {
		return nil, err
	}

	if err := u.repo.CreateAdvert(ctx, tx, newAdvert); err != nil {
		return nil, err
	}

	if err := u.repo.CreatePriceChange(ctx, tx, newPriceChange); err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return newAdvert, nil
}

// GetAdvertById handles the getting house advert process.
func (u *AdvertUsecase) GetAdvertById(ctx context.Context, id uuid.UUID) (foundAdvert *models.AdvertData, err error) {
	var typeAdvert *models.AdvertTypeAdvert
	typeAdvert, err = u.repo.GetTypeAdvertById(ctx, id)
	if err != nil {
		return nil, err
	}

	switch *typeAdvert {
	case models.AdvertTypeFlat:
		if foundAdvert, err = u.repo.GetFlatAdvertById(ctx, id); err != nil {
			return nil, err
		}
	case models.AdvertTypeHouse:
		if foundAdvert, err = u.repo.GetHouseAdvertById(ctx, id); err != nil {
			return nil, err
		}
	}

	var foundImages []*models.ImageResp
	if foundImages, err = u.repo.SelectImages(foundAdvert.ID); err != nil {
		return nil, err
	}

	foundAdvert.Images = foundImages

	return foundAdvert, nil
}

// UpdateAdvertById handles the updating advert process.
func (u *AdvertUsecase) UpdateAdvertById(ctx context.Context, advertUpdateData *models.AdvertUpdateData) (err error) {
	typeAdvert := advertUpdateData.TypeAdvert
	tx, err := u.repo.BeginTx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
		}
	}()

	if typeAdvertOld, err := u.repo.GetTypeAdvertById(ctx, advertUpdateData.ID); err == nil {
		if *typeAdvertOld != models.AdvertTypeAdvert(typeAdvert) {

			if err = u.repo.ChangeTypeAdvert(ctx, tx, advertUpdateData.ID); err != nil {
				return err
			}
		}
	} else {
		return err
	}

	switch models.AdvertTypeAdvert(typeAdvert) {
	case models.AdvertTypeFlat:
		if err = u.repo.UpdateFlatAdvertById(ctx, tx, advertUpdateData); err != nil {
			return err
		}
	case models.AdvertTypeHouse:
		if err = u.repo.UpdateHouseAdvertById(ctx, tx, advertUpdateData); err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// DeleteAdvertById handles the deleting advert process.
func (u *AdvertUsecase) DeleteAdvertById(ctx context.Context, advertId uuid.UUID) (err error) {
	tx, err := u.repo.BeginTx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
		}
	}()

	typeAdvert, err := u.repo.GetTypeAdvertById(ctx, advertId)
	if err != nil {
		return err
	}

	switch *typeAdvert {
	case models.AdvertTypeFlat:
		if err = u.repo.DeleteFlatAdvertById(ctx, tx, advertId); err != nil {
			return err
		}
	case models.AdvertTypeHouse:
		if err = u.repo.DeleteHouseAdvertById(ctx, tx, advertId); err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// GetSquareAdvertsList handles the square adverts getting process with paggination.
func (u *AdvertUsecase) GetSquareAdvertsList(ctx context.Context, pageSize, offset int) (foundAdverts []*models.AdvertSquareData, err error) {
	if foundAdverts, err = u.repo.GetSquareAdverts(ctx, pageSize, offset); err != nil {
		return nil, err
	}

	return foundAdverts, nil
}

// GetRectangleAdvertsList handles the rectangle adverts getting process with paggination and search.
func (u *AdvertUsecase) GetRectangleAdvertsList(ctx context.Context, advertFilter models.AdvertFilter) (foundAdverts *models.AdvertDataPage, err error) {
	if foundAdverts, err = u.repo.GetRectangleAdverts(ctx, advertFilter); err != nil {
		return nil, err
	}

	return foundAdverts, nil
}

// GetExistBuildingsByAddress handles the buildings getting process by address with paggination.
func (u *AdvertUsecase) GetExistBuildingsByAddress(ctx context.Context, address string, pageSize int) (foundBuildings []*models.BuildingsExistData, err error) {
	if foundBuildings, err = u.repo.CheckExistsBuildings(ctx, pageSize, address); err != nil {
		return nil, err
	}

	return foundBuildings, nil
}

// GetRectangleAdvertsByUserId handles the rectangle adverts getting process with paggination by userId.
func (u *AdvertUsecase) GetRectangleAdvertsByUserId(ctx context.Context, pageSize, offset int, userId uuid.UUID) (foundAdverts []*models.AdvertRectangleData, err error) {
	if foundAdverts, err = u.repo.GetRectangleAdvertsByUserId(ctx, pageSize, offset, userId); err != nil {
		return nil, err
	}

	return foundAdverts, nil
}

// GetRectangleAdvertsByComplexId handles the rectangle adverts getting process with paggination by complexId.
func (u *AdvertUsecase) GetRectangleAdvertsByComplexId(ctx context.Context, pageSize, offset int, comlexId uuid.UUID) (foundAdverts []*models.AdvertRectangleData, err error) {
	if foundAdverts, err = u.repo.GetRectangleAdvertsByComplexId(ctx, pageSize, offset, comlexId); err != nil {
		return nil, err
	}

	return foundAdverts, nil
}
