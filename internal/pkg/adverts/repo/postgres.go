package repo

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"
	"fmt"
	"strings"

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
	query := `SELECT id, user_id, phone, description, building_id, company_id, price, location, data_creation, is_deleted FROM adverts WHERE id = $1`

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
	query := `SELECT id, user_id, phone, description, building_id, company_id, price, location, data_creation, is_deleted FROM adverts`
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

func (r *AdvertRepo) GetAdvertsWithImages(ctx context.Context) ([]*models.AdvertWithImages, error) {
	adverts, err := r.GetAdvertsList(ctx)
	if err != nil {
		return nil, err
	}

	advertsWithImages := []*models.AdvertWithImages{}
	for _, advert := range adverts {
		images, err := r.GetImagesByAdvertId(ctx, advert.ID)
		if err != nil {
			return nil, err
		}

		advertWithImages := &models.AdvertWithImages{
			Advert: advert,
			Images: images,
		}
		advertsWithImages = append(advertsWithImages, advertWithImages)
	}

	return advertsWithImages, nil
}

func (r *AdvertRepo) GetImagesByAdvertId(ctx context.Context, advertId uuid.UUID) ([]*models.Image, error) {
	query := `SELECT id, advert_id, filename, priority, data_creation, is_deleted FROM images WHERE advert_id = $1`

	rows, err := r.db.QueryContext(ctx, query, advertId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	images := []*models.Image{}
	for rows.Next() {
		image := &models.Image{}
		err := rows.Scan(&image.ID, &image.AdvertId, &image.Filename, &image.Priority, &image.DataCreation, &image.IsDeleted)
		if err != nil {
			return nil, err
		}
		images = append(images, image)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return images, nil
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

// UpdateAdvertById updates fields advert from the database by their id.
func (r *AdvertRepo) UpdateAdvertById(ctx context.Context, body map[string]interface{}, id uuid.UUID) (err error) {
	var updates []string
	var values []interface{}
	i := 1
	for key, value := range body {
		if key == "id" {
			return fmt.Errorf("ID is not changeable")
		}
		updates = append(updates, fmt.Sprintf("%s = $%d", key, i))
		values = append(values, value)
		i++
	}
	values = append(values, id)

	query := fmt.Sprintf("UPDATE adverts SET %s WHERE id = $%d", strings.Join(updates, ", "), len(values))
	_, err = r.db.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}
