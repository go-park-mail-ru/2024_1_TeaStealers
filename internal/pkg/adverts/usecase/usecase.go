package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts"
	"context"

	"github.com/satori/uuid"
)

// AdvertUsecase represents the usecase for authentication.
type AdvertUsecase struct {
	repo adverts.AdvertRepo
}

// NewAdvertUsecase creates a new instance of AdvertUsecase.
func NewAdvertUsecase(repo adverts.AdvertRepo) *AdvertUsecase {
	return &AdvertUsecase{repo: repo}
}

// CreateFlatAdvert handles the creation advert process.
func (u *AdvertUsecase) CreateFlatAdvert(ctx context.Context, data *models.AdvertFlatCreateData) (*models.Advert, error) {
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
		if err := u.repo.CreateBuilding(ctx, building); err != nil {
			return nil, err
		}
	}

	newFlat := &models.Flat{
		ID:                uuid.NewV4(),
		BuildingID:        building.ID,
		AdvertTypeID:      newAdvertType.ID,
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

	if err := u.repo.CreateAdvertType(ctx, newAdvertType); err != nil {
		return nil, err
	}

	if err := u.repo.CreateFlat(ctx, newFlat); err != nil {
		return nil, err
	}

	if err := u.repo.CreateAdvert(ctx, newAdvert); err != nil {
		return nil, err
	}

	if err := u.repo.CreatePriceChange(ctx, newPriceChange); err != nil {
		return nil, err
	}

	return newAdvert, nil
}

// CreateFlatAdvert handles the creation advert process.
func (u *AdvertUsecase) CreateHouseAdvert(ctx context.Context, data *models.AdvertHouseCreateData) (*models.Advert, error) {
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

	newBuilding := &models.Building{
		ID:           uuid.NewV4(),
		Floor:        data.FloorGeneral,
		Material:     data.Material,
		Address:      data.Address,
		AddressPoint: data.AddressPoint,
		YearCreation: data.YearCreation,
	}

	newHouse := &models.House{
		ID:            uuid.NewV4(),
		BuildingID:    newBuilding.ID,
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

	if err := u.repo.CreateAdvertType(ctx, newAdvertType); err != nil {
		return nil, err
	}

	if err := u.repo.CreateBuilding(ctx, newBuilding); err != nil {
		return nil, err
	}

	if err := u.repo.CreateHouse(ctx, newHouse); err != nil {
		return nil, err
	}

	if err := u.repo.CreateAdvert(ctx, newAdvert); err != nil {
		return nil, err
	}

	if err := u.repo.CreatePriceChange(ctx, newPriceChange); err != nil {
		return nil, err
	}

	return newAdvert, nil
}
