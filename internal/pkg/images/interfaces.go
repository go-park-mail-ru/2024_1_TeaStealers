package images

import (
	"2024_1_TeaStealers/internal/models"
	"context"

	"github.com/satori/uuid"
)

// ImageUsecase represents the usecase interface for manage images.
type ImageUsecase interface {
	CreateImage(ctx context.Context, data *models.ImageCreateData, id uuid.UUID) (*models.Image, error)
	GetImagesByAdvertId(ctx context.Context, advertId uuid.UUID) (findImages []*models.Image, err error)
	DeleteImageById(ctx context.Context, id uuid.UUID) (err error)
}

// ImageRepo represents the repository interface for manage images.
type ImageRepo interface {
	CreateImage(ctx context.Context, image *models.Image) error
	GetImagesByAdvertId(ctx context.Context, advertId uuid.UUID) ([]*models.Image, error)
	DeleteImageById(ctx context.Context, id uuid.UUID) error
}
