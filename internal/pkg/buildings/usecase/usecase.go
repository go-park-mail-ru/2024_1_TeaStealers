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
