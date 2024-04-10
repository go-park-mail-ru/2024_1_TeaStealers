package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/images"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/satori/uuid"
	"go.uber.org/zap"
)

// ImageUsecase represents the usecase for images for advert.
type ImageUsecase struct {
	repo   images.ImageRepo
	logger *zap.Logger
}

// NewImageUsecase creates a new instance of ImageUsecase.
func NewImageUsecase(repo images.ImageRepo, logger *zap.Logger) *ImageUsecase {
	return &ImageUsecase{repo: repo, logger: logger}
}

// UploadImage upload image for advert
func (u *ImageUsecase) UploadImage(file io.Reader, fileType string, advertUUID uuid.UUID) (*models.ImageResp, error) {
	newId := uuid.NewV4()
	fileName := newId.String() + fileType
	subDirectory := filepath.Join("adverts", advertUUID.String())
	directory := filepath.Join(os.Getenv("DOCKER_DIR"), subDirectory)
	if err := os.MkdirAll(directory, 0755); err != nil {
		return nil, err
	}
	destination, err := os.Create(directory + "/" + fileName)
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
	newImage := &models.Image{
		ID:       newId,
		AdvertID: advertUUID,
		Photo:    subDirectory + "/" + fileName,
		Priority: 1,
	}
	image, err := u.repo.StoreImage(newImage)
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
