package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/images/mock"
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestUploadImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := images_mock.NewMockImageRepo(ctrl)
	usecase := NewImageUsecase(mockRepo)
	type args struct {
		file          io.Reader
		fileType      string
		advertUUID    uuid.UUID
		expectedImage *models.Image
	}
	type want struct {
		imageResp *models.ImageResp
		err       error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful store image",
			args: args{
				file:       bytes.NewBufferString("test file content"),
				fileType:   "jpg",
				advertUUID: uuid.NewV4(),
				expectedImage: &models.Image{
					ID:       uuid.NewV4(),
					AdvertID: uuid.NewV4(),
					Photo:    "adverts/" + uuid.NewV4().String() + "/" + uuid.NewV4().String() + ".jpg",
					Priority: 1,
				},
			},
			want: want{
				imageResp: &models.ImageResp{
					ID:       uuid.NewV4(),
					Photo:    "adverts/" + uuid.NewV4().String() + "/" + uuid.NewV4().String() + ".jpg",
					Priority: 1,
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().StoreImage(gomock.Any()).Return(tt.want.imageResp, tt.want.err)

			gotImageResp, err := usecase.UploadImage(tt.args.file, tt.args.fileType, tt.args.advertUUID)
			assert.Equal(t, tt.want.err, err)
			assert.Equal(t, tt.want.imageResp, gotImageResp)
		})
	}
	os.RemoveAll("adverts")
}

func TestGetAdvertImages(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := images_mock.NewMockImageRepo(ctrl)
	usecase := NewImageUsecase(mockRepo)

	type args struct {
		advertUUID uuid.UUID
	}
	type want struct {
		images []*models.ImageResp
		err    error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful get images",
			args: args{
				advertUUID: uuid.NewV4(),
			},
			want: want{
				images: []*models.ImageResp{
					{
						ID:       uuid.NewV4(),
						Photo:    "adverts/" + uuid.NewV4().String() + "/image1.jpg",
						Priority: 1,
					},
					{
						ID:       uuid.NewV4(),
						Photo:    "adverts/" + uuid.NewV4().String() + "/image2.jpg",
						Priority: 2,
					},
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().SelectImages(tt.args.advertUUID).Return(tt.want.images, tt.want.err)

			gotImages, err := usecase.GetAdvertImages(tt.args.advertUUID)
			assert.Equal(t, tt.want.err, err)
			assert.Equal(t, tt.want.images, gotImages)
		})
	}
}

func TestDeleteImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := images_mock.NewMockImageRepo(ctrl)
	usecase := NewImageUsecase(mockRepo)

	type args struct {
		imageId uuid.UUID
	}
	type want struct {
		images []*models.ImageResp
		err    error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful delete image",
			args: args{
				imageId: uuid.NewV4(),
			},
			want: want{
				images: []*models.ImageResp{
					{
						ID:       uuid.NewV4(),
						Photo:    "adverts/123/image1.jpg",
						Priority: 1,
					},
					{
						ID:       uuid.NewV4(),
						Photo:    "adverts/123/image2.jpg",
						Priority: 2,
					},
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().DeleteImage(tt.args.imageId).Return(tt.want.images, tt.want.err)

			gotImages, err := usecase.DeleteImage(tt.args.imageId)
			assert.Equal(t, tt.want.err, err)
			assert.Equal(t, tt.want.images, gotImages)
		})
	}
}
