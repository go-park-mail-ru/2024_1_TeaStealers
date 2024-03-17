package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/images"
	"github.com/satori/uuid"
	"io"
	"path/filepath"
)

// ImageUsecase represents the usecase for images for advert.
type ImageUsecase struct {
	repo images.ImageRepo
}

// NewImageUsecase creates a new instance of ImageUsecase.
func NewImageUsecase(repo images.ImageRepo) *ImageUsecase {
	return &ImageUsecase{repo: repo}
}

// UploadImage upload image for advert
func (u *ImageUsecase) UploadImage(file io.Reader, fileType string, advertUUID uuid.UUID) (*models.ImageResp, error) {
	newId := uuid.NewV4()
	directory := filepath.Join("advert", advertUUID.String())
	newImage := &models.Image{
		ID:       newId,
		AdvertID: advertUUID,
		Photo:    filepath.Join(directory, newId.String()+fileType),
		Priority: 1,
	}
	image, err := u.repo.StoreImage(file, newImage, directory)
	if err != nil {
		return nil, err
	}
	return image, nil
}

// GetAdvertImages return list of images for advert
func (u *ImageUsecase) GetAdvertImages(advertId uuid.UUID) ([]*models.ImageResp, error) {
	imagesList, err := u.repo.SelectImages(advertId)
	if err != nil {
		return nil, err
	}
	return imagesList, nil
}

// DeleteImage delete image bby id and return new list images
func (u *ImageUsecase) DeleteImage(imageId uuid.UUID) ([]*models.ImageResp, error) {
	imagesList, err := u.repo.DeleteImage(imageId)
	if err != nil {
		return nil, err
	}
	return imagesList, nil
}
