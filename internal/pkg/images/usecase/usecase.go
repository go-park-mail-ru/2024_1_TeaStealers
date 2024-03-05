package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/images"
	"context"

	"time"

	"github.com/satori/uuid"
)

// ImageUsecase represents the usecase for manage images.
type ImageUsecase struct {
	repo images.ImageRepo
}

// NewImageUsecase creates a new instance of ImageUsecase.
func NewImageUsecase(repo images.ImageRepo) *ImageUsecase {
	return &ImageUsecase{repo: repo}
}

// CreateImage handles the image creation process.
func (u *ImageUsecase) CreateImage(ctx context.Context, data *models.ImageCreateData, id uuid.UUID) (*models.Image, error) {
	newImage := &models.Image{
		ID:           id,
		Filename:     data.AdvertId.String() + "/" + id.String(),
		AdvertId:     data.AdvertId,
		Priority:     data.Priority,
		DataCreation: time.Now(),
		IsDeleted:    false,
	}

	if err := u.repo.CreateImage(ctx, newImage); err != nil {
		return nil, err
	}

	return newImage, nil
}

// GetImagesByAdvertId handles the images getting process.
func (u *ImageUsecase) GetImagesByAdvertId(ctx context.Context, advertId uuid.UUID) (findImages []*models.Image, err error) {
	if findImages, err = u.repo.GetImagesByAdvertId(ctx, advertId); err != nil {
		return nil, err
	}

	return findImages, nil
}

// DeleteImageById handles the deleting image process.
func (u *ImageUsecase) DeleteImageById(ctx context.Context, id uuid.UUID) (err error) {
	if err = u.repo.DeleteImageById(ctx, id); err != nil {
		return err
	}

	return nil
}
