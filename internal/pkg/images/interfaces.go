package images

import (
	"2024_1_TeaStealers/internal/models"
	"github.com/satori/uuid"
	"io"
)

// ImageUsecase represents the usecase interface for images for advert.
type ImageUsecase interface {
	UploadImage(io.Reader, string, uuid.UUID) (*models.ImageResp, error)
	GetAdvertImages(uuid.UUID) ([]*models.ImageResp, error)
	DeleteImage(uuid.UUID) ([]*models.ImageResp, error)
}

// ImagesRepo represents the repository interface for images for advert.
type ImageRepo interface {
	StoreImage(io.Reader, *models.Image, string) (*models.ImageResp, error)
	SelectImages(uuid.UUID) ([]*models.ImageResp, error)
	DeleteImage(uuid.UUID) ([]*models.ImageResp, error)
}
