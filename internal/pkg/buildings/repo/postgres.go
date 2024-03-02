package repo

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"
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
