//go:generate mockgen -destination=mock/${GOFILE} -package=${GOPACKAGE}_mock -source=${GOFILE}

package images

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"io"
)

// ImageUsecase represents the usecase interface for images for advert.
type ImageUsecase interface {
	UploadImage(context.Context, io.Reader, string, int64) (*models.ImageResp, error)
	GetAdvertImages(context.Context, int64) ([]*models.ImageResp, error)
	DeleteImage(context.Context, int64) ([]*models.ImageResp, error)
}

// ImagesRepo represents the repository interface for images for advert.
type ImageRepo interface {
	StoreImage(context.Context, *models.Image) (*models.ImageResp, error)
	SelectImages(context.Context, int64) ([]*models.ImageResp, error)
	DeleteImage(context.Context, int64) ([]*models.ImageResp, error)
}
