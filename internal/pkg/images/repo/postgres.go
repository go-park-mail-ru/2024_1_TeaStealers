package repo

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/config/dbPool"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"strconv"

	"go.uber.org/zap"
)

// ImageRepo represents a repository for adverts images changes.
type ImageRepo struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

// NewRepository creates a new instance of ImageRepo.
func NewRepository(logger *zap.Logger) *ImageRepo {
	return &ImageRepo{db: dbPool.GetDBPool(), logger: logger}
}

// StoreImage insert new images and create file in directory
func (repo *ImageRepo) StoreImage(ctx context.Context, image *models.Image) (*models.ImageResp, error) {
	maxPriority := 0
	query := `SELECT MAX(priority) FROM image WHERE advert_id = $1`

	_ = repo.db.QueryRow(ctx, query, image.AdvertID).Scan(&maxPriority)
	maxPriority++

	insert := `INSERT INTO image (advert_id, photo, priority) VALUES ($1, $2, $3) RETURNING id`
	if err := repo.db.QueryRow(ctx, insert, image.AdvertID, image.Photo, maxPriority).Scan(&image.ID); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	newImage := &models.ImageResp{}
	selectQuery := `SELECT photo, priority FROM image WHERE id = $1`
	err := repo.db.QueryRow(ctx, selectQuery, image.ID).Scan(&newImage.Photo, &newImage.Priority)
	if err != nil {
		return nil, err
	}
	return newImage, nil
}

// SelectImages select list images for advert
func (repo *ImageRepo) SelectImages(ctx context.Context, advertId int64) ([]*models.ImageResp, error) {
	selectQuery := `SELECT id, photo, priority FROM image WHERE advert_id = $1 AND is_deleted = false`
	rows, err := repo.db.Query(ctx, selectQuery, advertId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	images := []*models.ImageResp{}

	for rows.Next() {
		var id int64
		var photo string
		var priority int
		if err := rows.Scan(&id, &photo, &priority); err != nil {
			return nil, err
		}
		image := &models.ImageResp{
			ID:       id,
			Photo:    photo,
			Priority: priority,
		}
		images = append(images, image)
	}

	return images, nil
}

// DeleteImage delete image by id and return new list images for advert
func (repo *ImageRepo) DeleteImage(ctx context.Context, idImage int64) ([]*models.ImageResp, error) {
	query := `UPDATE image SET is_deleted = true WHERE id = $1`
	if _, err := repo.db.Exec(ctx, query, idImage); err != nil {
		return nil, err
	}

	var idAdvert string
	querySelectId := `SELECT advert_id FROM image WHERE id = $1`
	if err := repo.db.QueryRow(ctx, querySelectId, idImage).Scan(&idAdvert); err != nil {
		return nil, err
	}
	idAdvertID, err := strconv.ParseInt(idAdvert, 10, 64)

	if err != nil {
		return nil, err
	}
	list, err := repo.SelectImages(ctx, idAdvertID)
	if err != nil {
		return nil, err
	}

	return list, nil
}
