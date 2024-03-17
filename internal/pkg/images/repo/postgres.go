package repo

import (
	"2024_1_TeaStealers/internal/models"
	"database/sql"
	"fmt"
	"github.com/satori/uuid"
	"io"
	"os"
)

// ImageRepo represents a repository for adverts images changes.
type ImageRepo struct {
	db *sql.DB
}

// NewRepository creates a new instance of ImageRepo.
func NewRepository(db *sql.DB) *ImageRepo {
	return &ImageRepo{db: db}
}

// StoreImage insert new images and create file in directory
func (repo *ImageRepo) StoreImage(file io.Reader, image *models.Image, directory string) (*models.ImageResp, error) {
	if _, err := os.Stat(os.Getenv("BASE_DIR") + directory); os.IsNotExist(err) {
		_ = os.Mkdir(os.Getenv("BASE_DIR")+directory, 0755)
	}
	destination, err := os.Create(os.Getenv("BASE_DIR") + image.Photo)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer destination.Close()

	_, err = io.Copy(destination, file)
	if err != nil {
		fmt.Println(err.Error())

		return nil, err
	}

	maxPriority := 0
	query := `SELECT MAX(priority) FROM images WHERE advertid = $1`

	_ = repo.db.QueryRow(query, image.AdvertID).Scan(&maxPriority)
	fmt.Println(maxPriority)
	maxPriority = maxPriority + 1
	fmt.Println(maxPriority)

	insert := `INSERT INTO images (id, advertid, photo, priority) VALUES ($1, $2, $3, $4)`
	if _, err := repo.db.Exec(insert, image.ID, image.AdvertID, image.Photo, maxPriority); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	newImage := &models.ImageResp{}
	selectQuery := `SELECT photo, priority FROM images WHERE id = $1`
	err = repo.db.QueryRow(selectQuery, image.ID).Scan(&newImage.Photo, &newImage.Priority)
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
	if _, err := repo.db.Query(query, idImage); err != nil {
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
