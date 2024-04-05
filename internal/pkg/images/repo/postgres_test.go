package repo_test

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/images/repo"
	"database/sql"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/suite"
)

type ImageRepoTestSuite struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
}

func (suite *ImageRepoTestSuite) SetupTest() {
	var err error
	suite.db, suite.mock, err = sqlmock.New()
	suite.Require().NoError(err)
}

func (suite *ImageRepoTestSuite) TearDownTest() {
	suite.mock.ExpectClose()
	suite.Require().NoError(suite.db.Close())
}

func (suite *ImageRepoTestSuite) TestSelectImages() {
	type args struct {
		advertID uuid.UUID
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
			name: "successful select images",
			args: args{
				advertID: uuid.NewV4(),
			},
			want: want{
				images: []*models.ImageResp{
					{
						ID:       uuid.NewV4(),
						Photo:    "image1.jpg",
						Priority: 1,
					},
					{
						ID:       uuid.NewV4(),
						Photo:    "image2.jpg",
						Priority: 2,
					},
				},
				err: nil,
			},
		},
		{
			name: "no images found",
			args: args{
				advertID: uuid.NewV4(),
			},
			want: want{
				images: []*models.ImageResp{},
				err:    nil,
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.setupMockSelectImage(tt.args.advertID, tt.want.images)
			rep := repo.NewRepository(suite.db)
			gotImages, gotErr := rep.SelectImages(tt.args.advertID)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Empty(gotImages)
			} else {
				suite.Assert().Equal(gotImages, tt.want.images)
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *ImageRepoTestSuite) setupMockSelectImage(advertID uuid.UUID, images []*models.ImageResp) {
	rows := sqlmock.NewRows([]string{"id", "photo", "priority"})
	for _, image := range images {
		rows = rows.AddRow(image.ID, image.Photo, image.Priority)
	}
	suite.mock.ExpectQuery(`SELECT id, photo, priority FROM images WHERE advertid = \$1 AND isdeleted = false`).
		WithArgs(advertID).
		WillReturnRows(rows)
}

func TestImageRepoTestSuite(t *testing.T) {
	suite.Run(t, new(ImageRepoTestSuite))
}

func (suite *ImageRepoTestSuite) TestStoreImage() {
	type args struct {
		image *models.Image
	}
	type want struct {
		imageResp *models.ImageResp
		err       error
	}
	zeroid, _ := uuid.FromString("00000000-0000-0000-0000-000000000000")
	tests := []struct {
		name string
		args args
		want want
	}{
		{

			name: "successful store image",
			args: args{
				image: &models.Image{
					ID:           uuid.NewV4(),
					AdvertID:     uuid.NewV4(),
					Photo:        "image1.jpg",
					Priority:     1,
					DateCreation: time.Now(),
					IsDeleted:    false,
				},
			},
			want: want{
				imageResp: &models.ImageResp{
					ID:       zeroid,
					Photo:    "image1.jpg",
					Priority: 1,
				},
				err: nil,
			},
		},
		{
			name: "failed store image",
			args: args{
				image: &models.Image{
					ID:           uuid.NewV4(),
					AdvertID:     uuid.NewV4(),
					Photo:        "image1.jpg",
					Priority:     1,
					DateCreation: time.Now(),
					IsDeleted:    false,
				},
			},
			want: want{
				imageResp: nil,
				err:       sql.ErrNoRows,
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.setupMockStoreImage(tt.args.image, tt.want.imageResp, tt.want.err)
			repo := repo.NewRepository(suite.db)
			gotImageResp, gotErr := repo.StoreImage(tt.args.image)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Nil(gotImageResp)
			} else {
				suite.Assert().Equal(gotImageResp, tt.want.imageResp)
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *ImageRepoTestSuite) setupMockStoreImage(image *models.Image, imageResp *models.ImageResp, err error) {
	if err != nil {
		suite.mock.ExpectExec(`INSERT INTO images \(id, advertid, photo, priority\) VALUES \(\$1, \$2, \$3, \$4\)`).
			WithArgs(image.ID, image.AdvertID, image.Photo, image.Priority).
			WillReturnError(err)
	} else {
		suite.mock.ExpectExec(`INSERT INTO images \(id, advertid, photo, priority\) VALUES \(\$1, \$2, \$3, \$4\)`).
			WithArgs(image.ID, image.AdvertID, image.Photo, image.Priority).
			WillReturnResult(sqlmock.NewResult(1, 1))

		rows := sqlmock.NewRows([]string{"photo", "priority"}).
			AddRow(imageResp.Photo, imageResp.Priority)

		suite.mock.ExpectQuery(`SELECT photo, priority FROM images WHERE id = \$1`).
			WithArgs(image.ID).
			WillReturnRows(rows)
	}
}

func (suite *ImageRepoTestSuite) setupMockDeleteImage(idImage, advertId uuid.UUID, response []*models.ImageResp, err error) {
	// todo willreturn error hardcode
	suite.mock.ExpectExec(`UPDATE images SET isdeleted = true WHERE id = \$1`).
		WithArgs(idImage).
		WillReturnResult(sqlmock.NewResult(1, 1))

	rows := sqlmock.NewRows([]string{"advertId"}).
		AddRow(advertId)

	suite.mock.ExpectQuery(`SELECT advertid FROM images WHERE id = \$1`).
		WithArgs(idImage).
		WillReturnRows(rows).WillReturnError(nil)

	rows = sqlmock.NewRows([]string{"id", "photo", "priority"})
	for _, image := range response {
		rows = rows.AddRow(image.ID, image.Photo, image.Priority)
	}

	suite.mock.ExpectQuery(`SELECT id, photo, priority FROM images WHERE advertid = \$1 AND isdeleted = false`).
		WithArgs(advertId).
		WillReturnRows(rows).WillReturnError(nil)
}

func (suite *ImageRepoTestSuite) TestDeleteImage() {
	type args struct {
		idImage uuid.UUID
	}
	type want struct {
		imageResp []*models.ImageResp
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
				idImage: uuid.NewV4(),
			},
			want: want{
				imageResp: []*models.ImageResp{
					{
						ID:       uuid.NewV4(),
						Photo:    "image1.jpg",
						Priority: 1,
					},
					{
						ID:       uuid.NewV4(),
						Photo:    "image2.jpg",
						Priority: 2,
					},
				},
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			advertId := uuid.NewV4()
			suite.setupMockDeleteImage(tt.args.idImage, advertId, tt.want.imageResp, tt.want.err)
			repo := repo.NewRepository(suite.db)
			gotImageResp, gotErr := repo.DeleteImage(tt.args.idImage)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Nil(gotImageResp)
			} else {
				suite.Assert().Equal(gotImageResp, tt.want.imageResp)
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}
