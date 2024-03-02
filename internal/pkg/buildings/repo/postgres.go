package repo

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"

	"github.com/satori/uuid"
)

// BuildingRepo represents a repository for buildings.
type BuildingRepo struct {
	db *sql.DB
}

// NewRepository creates a new instance of BuildingRepo.
func NewRepository(db *sql.DB) *BuildingRepo {
	return &BuildingRepo{db: db}
}

// CreateBuilding creates a new building in the database.
func (r *BuildingRepo) CreateBuilding(ctx context.Context, building *models.Building) error {
	insert := `INSERT INTO buildings (id, location, description, data_creation, is_deleted) VALUES ($1, $2, $3, $4, $5)`

	if _, err := r.db.ExecContext(ctx, insert, building.ID, building.Location, building.Descpription, building.DataCreation, building.IsDeleted); err != nil {
		return err
	}
	return nil
}

// GetBuildingById retrieves a building from the database by their id.
func (r *BuildingRepo) GetBuildingById(ctx context.Context, id uuid.UUID) (*models.Building, error) {
	query := `SELECT * FROM buildings WHERE id = $1`

	res := r.db.QueryRowContext(ctx, query, id)

	building := &models.Building{
		ID: id,
	}
	if err := res.Scan(&building.ID, &building.Location, &building.Descpription, &building.DataCreation, &building.IsDeleted); err != nil {
		return nil, err
	}

	return building, nil
}
