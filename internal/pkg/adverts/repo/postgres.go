package repo

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"
)

// AdvertRepo represents a repository for adverts.
type AdvertRepo struct {
	db *sql.DB
}

// NewRepository creates a new instance of AdvertRepo.
func NewRepository(db *sql.DB) *AdvertRepo {
	return &AdvertRepo{db: db}
}

// CreateAdvert creates a new advert in the database.
func (r *AdvertRepo) CreateAdvert(ctx context.Context, advert *models.Advert) error {
	insert := `INSERT INTO adverts (id, user_id, phone, description, building_id, company_id, price, location, data_creation, is_deleted) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	if _, err := r.db.ExecContext(ctx, insert, advert.ID, advert.UserId, advert.Phone, advert.Descpription, advert.BuildingId, advert.CompanyId, advert.Price, advert.Location, advert.DataCreation, advert.IsDeleted); err != nil {
		return err
	}
	return nil
}
