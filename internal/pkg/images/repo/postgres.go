package repo

import (
	"2024_1_TeaStealers/internal/models"
	"database/sql"
	"fmt"

	"github.com/satori/uuid"
	"go.uber.org/zap"
)

// ImageRepo represents a repository for adverts images changes.
type ImageRepo struct {
	db     *sql.DB
	logger *zap.Logger
}

// NewRepository creates a new instance of ImageRepo.
func NewRepository(db *sql.DB, logger *zap.Logger) *ImageRepo {
	return &ImageRepo{db: db, logger: logger}
}

// StoreImage insert new images and create file in directory
func (repo *ImageRepo) StoreImage(image *models.Image) (*models.ImageResp, error) {
	maxPriority := 0
	query := `SELECT MAX(priority) FROM images WHERE advertid = $1`

	_ = repo.db.QueryRow(query, image.AdvertID).Scan(&maxPriority)
	maxPriority++

	insert := `INSERT INTO images (id, advertid, photo, priority) VALUES ($1, $2, $3, $4)`
	if _, err := repo.db.Exec(insert, image.ID, image.AdvertID, image.Photo, maxPriority); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	newImage := &models.ImageResp{}
	selectQuery := `SELECT photo, priority FROM images WHERE id = $1`
	err := repo.db.QueryRow(selectQuery, image.ID).Scan(&newImage.Photo, &newImage.Priority)
	if err != nil {
		return nil, err
	}
	return newImage, nil
}

// SelectImages select list images for advert
func (repo *ImageRepo) SelectImages(advertId uuid.UUID) ([]*models.ImageResp, error) {
	selectQuery := `SELECT id, photo, priority FROM images WHERE advertid = $1 AND isdeleted = false`
	rows, err := repo.db.Query(selectQuery, advertId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	images := []*models.ImageResp{}

	for rows.Next() {
		var id uuid.UUID
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
func (repo *ImageRepo) DeleteImage(idImage uuid.UUID) ([]*models.ImageResp, error) {
	query := `UPDATE images SET isdeleted = true WHERE id = $1`
	if _, err := repo.db.Exec(query, idImage); err != nil {
		return nil, err
	}
	var idAdvert string
	querySelectId := `SELECT advertid FROM images WHERE id = $1`
	if err := repo.db.QueryRow(querySelectId, idImage).Scan(&idAdvert); err != nil {
		return nil, err
	}
	idAdvertUUID, err := uuid.FromString(idAdvert)

	if err != nil {
		return nil, err
	}
	list, err := repo.SelectImages(idAdvertUUID)
	if err != nil {
		return nil, err
	}

	return list, nil
}
