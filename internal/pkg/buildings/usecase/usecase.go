package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/buildings"
	"context"

	"time"

	"github.com/satori/uuid"
)

// BuildingUsecase represents the usecase for manage buildings.
type BuildingUsecase struct {
	repo buildings.BuildingRepo
}

// NewBuildingUsecase creates a new instance of BuildingUsecase.
func NewBuildingUsecase(repo buildings.BuildingRepo) *BuildingUsecase {
	return &BuildingUsecase{repo: repo}
}

// CreateBuilding handles the building creation process.
func (u *BuildingUsecase) CreateBuilding(ctx context.Context, data *models.BuildingCreateData) (*models.Building, error) {
	newBuilding := &models.Building{
		ID:           uuid.NewV4(),
		Location:     data.Location,
		Descpription: data.Descpription,
		DataCreation: time.Now(),
		IsDeleted:    false,
	}

	if err := u.repo.CreateBuilding(ctx, newBuilding); err != nil {
		return nil, err
	}

	return newBuilding, nil
}

// GetBuildingById handles the building getting process.
func (u *BuildingUsecase) GetBuildingById(ctx context.Context, id uuid.UUID) (findBuilding *models.Building, err error) {
	if findBuilding, err = u.repo.GetBuildingById(ctx, id); err != nil {
		return nil, err
	}

	return findBuilding, nil
}

// GetBuildingsList handles the buildings getting process.
func (u *BuildingUsecase) GetBuildingsList(ctx context.Context) (findBuildings []*models.Building, err error) {
	if findBuildings, err = u.repo.GetBuildingsList(ctx); err != nil {
		return nil, err
	}

	return findBuildings, nil
}

// DeleteBuildingById handles the deleting building process.
func (u *BuildingUsecase) DeleteBuildingById(ctx context.Context, id uuid.UUID) (err error) {
	if err = u.repo.DeleteBuildingById(ctx, id); err != nil {
		return err
	}

	return nil
}

// UpdateBuildingById handles the updating building process.
func (u *BuildingUsecase) UpdateBuildingById(ctx context.Context, body map[string]interface{}, id uuid.UUID) (err error) {
	if err = u.repo.UpdateBuildingById(ctx, body, id); err != nil {
		return err
	}

	return nil
}
