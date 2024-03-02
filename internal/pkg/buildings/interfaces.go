package buildings

import (
	"2024_1_TeaStealers/internal/models"
	"context"

	"github.com/satori/uuid"
)

// BuildingUsecase represents the usecase interface for manage buildings.
type BuildingUsecase interface {
	CreateBuilding(ctx context.Context, data *models.BuildingCreateData) (*models.Building, error)
	GetBuildingById(ctx context.Context, id uuid.UUID) (findBuilding *models.Building, err error)
}

// BuildingRepo represents the repository interface for manage buildings.
type BuildingRepo interface {
	CreateBuilding(ctx context.Context, building *models.Building) error
	GetBuildingById(ctx context.Context, id uuid.UUID) (*models.Building, error)
}
