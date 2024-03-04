package repo

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"

	"github.com/satori/uuid"
)

// ImageRepo represents a repository for images.
type ImageRepo struct {
	db *sql.DB
}

// NewRepository creates a new instance of ImageRepo.
func NewRepository(db *sql.DB) *ImageRepo {
	return &ImageRepo{db: db}
}

// CreateImage creates a new image in the database.
func (r *ImageRepo) CreateImage(ctx context.Context, image *models.Image) error {
	insert := `INSERT INTO images (id, advert_id, filename, priority, data_creation, is_deleted) VALUES ($1, $2, $3, $4, $5, $6)`

	if _, err := r.db.ExecContext(ctx, insert, image.ID, image.AdvertId, image.Filename, image.Priority, image.DataCreation, image.IsDeleted); err != nil {
		return err
	}
	return nil
}

// GetImagesByAdvertId retrieves a image from the database by advert id.
func (r *ImageRepo) GetImagesByAdvertId(ctx context.Context, advertId uuid.UUID) ([]*models.Image, error) {
	query := `SELECT * FROM images WHERE advert_id = $1`

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

// DeleteImageById set is_deleted on true on a image from the database by their id.
func (r *ImageRepo) DeleteImageById(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE images SET is_deleted = true WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
