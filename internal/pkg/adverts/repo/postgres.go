package repo

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"

	"github.com/satori/uuid"
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

// GetAdvertById retrieves a advert from the database by their id.
func (r *AdvertRepo) GetAdvertById(ctx context.Context, id uuid.UUID) (*models.Advert, error) {
	query := `SELECT * FROM adverts WHERE id = $1`

	res := r.db.QueryRowContext(ctx, query, id)

	advert := &models.Advert{
		ID: id,
	}

	if err := res.Scan(&advert.ID, &advert.UserId, &advert.Phone, &advert.Descpription, &advert.BuildingId, &advert.CompanyId, &advert.Price, &advert.Location, &advert.DataCreation, &advert.IsDeleted); err != nil {
		return nil, err
	}

	return advert, nil
}

// GetAdvertsList retrieves a adverts from the database.
func (r *AdvertRepo) GetAdvertsList(ctx context.Context) ([]*models.Advert, error) {
	query := `SELECT * FROM adverts`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	adverts := []*models.Advert{}
	for rows.Next() {
		advert := &models.Advert{}
		err := rows.Scan(&advert.ID, &advert.UserId, &advert.Phone, &advert.Descpription, &advert.BuildingId, &advert.CompanyId, &advert.Price, &advert.Location, &advert.DataCreation, &advert.IsDeleted)
		if err != nil {
			return nil, err
		}
		adverts = append(adverts, advert)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return adverts, nil
}

// DeleteAdvertById set is_deleted on true on an advert from the database by their id.
func (r *AdvertRepo) DeleteAdvertById(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE adverts SET is_deleted = true WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
