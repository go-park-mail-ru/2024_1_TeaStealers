package repo_test

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts/repo"
	"context"
	"database/sql"
	"errors"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
)

type UserRepoTestSuite struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
}

func (suite *UserRepoTestSuite) SetupTest() {
	var err error
	suite.db, suite.mock, err = sqlmock.New()
	suite.Require().NoError(err)
}

func (suite *UserRepoTestSuite) TearDownTest() {
	suite.mock.ExpectClose()
	suite.Require().NoError(suite.db.Close())
}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}

func TestBeginTx(t *testing.T) {
	fakeDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer fakeDB.Close()
	rep := repo.NewRepository(fakeDB)
	// ctx := context.Background()
	// tx := new(sql.Tx)
	mock.ExpectBegin().WillReturnError(nil)
	tx, err := rep.BeginTx(context.Background())
	assert.NoError(t, err)
	assert.NotEmpty(t, tx)
}
func TestBeginTxFail(t *testing.T) {
	fakeDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer fakeDB.Close()
	rep := repo.NewRepository(fakeDB)
	// ctx := context.Background()
	// tx := new(sql.Tx)
	mock.ExpectBegin().WillReturnError(errors.New("error"))
	tx, err := rep.BeginTx(context.Background())
	assert.Error(t, err)
	assert.Empty(t, tx)
}

func (suite *UserRepoTestSuite) TestCreateAdvertType() {
	type args struct {
		adv     *models.AdvertType
		errExec error
		expExec bool
	}
	type want struct {
		err error
	}
	// id1 := uuid.NewV4()
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful create advert type",
			args: args{
				adv: &models.AdvertType{
					ID:         uuid.NewV4(),
					AdvertType: models.AdvertTypeHouse,
					IsDeleted:  false,
				},
				errExec: nil,
				expExec: true,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "fail create advert type",
			args: args{
				adv: &models.AdvertType{
					ID:         uuid.NewV4(),
					AdvertType: models.AdvertTypeHouse,
					IsDeleted:  false,
				},
				errExec: errors.New("some error"),
				expExec: true,
			},
			want: want{
				err: errors.New("some error"),
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.mock.ExpectBegin()
			tx, err := suite.db.Begin()
			if err != nil {
				suite.T().Fatal("Error beginning transaction:", err)
			}
			suite.setupMockCreateAdvertType(tt.args.adv, tt.args.errExec, tt.args.expExec)
			rep := repo.NewRepository(suite.db)
			gotErr := rep.CreateAdvertType(context.Background(), tx, tt.args.adv)
			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *UserRepoTestSuite) setupMockCreateAdvertType(advType *models.AdvertType, errExec error, expExec bool) {
	if expExec {
		suite.mock.ExpectExec(`INSERT INTO adverttypes \(id, adverttype\) VALUES \(\$1, \$2\)`).
			WillReturnError(errExec).WillReturnResult(sqlmock.NewResult(1, 1)).WithArgs(advType.ID, advType.AdvertType)
	}
}

func (suite *UserRepoTestSuite) TestCreateAdvert() {
	type args struct {
		adv     *models.Advert
		errExec error
		expExec bool
	}
	type want struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful create advert",
			args: args{
				adv: &models.Advert{
					ID:             uuid.NewV4(),
					UserID:         uuid.NewV4(),
					AdvertTypeID:   uuid.NewV4(),
					AdvertTypeSale: models.TypePlacementRent,
					Title:          "title",
					Description:    "descr",
					Phone:          "phone",
					IsAgent:        true,
					Priority:       1,
				},
				errExec: nil,
				expExec: true,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "fail create advert",
			args: args{
				adv: &models.Advert{
					ID:             uuid.NewV4(),
					UserID:         uuid.NewV4(),
					AdvertTypeID:   uuid.NewV4(),
					AdvertTypeSale: models.TypePlacementRent,
					Title:          "title",
					Description:    "descr",
					Phone:          "phone",
					IsAgent:        true,
					Priority:       1,
				},
				errExec: errors.New("some error"),
				expExec: true,
			},
			want: want{
				err: errors.New("some error"),
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.mock.ExpectBegin()
			tx, err := suite.db.Begin()
			if err != nil {
				suite.T().Fatal("Error beginning transaction:", err)
			}
			suite.setupMockCreateAdvert(tt.args.adv, tt.args.errExec, tt.args.expExec)
			rep := repo.NewRepository(suite.db)
			gotErr := rep.CreateAdvert(context.Background(), tx, tt.args.adv)
			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *UserRepoTestSuite) setupMockCreateAdvert(newAdvert *models.Advert, errExec error, expExec bool) {
	if expExec {
		suite.mock.ExpectExec(`INSERT INTO adverts \(id, userid, adverttypeid, adverttypeplacement, title, description, phone, isagent, priority\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8, \$9\)`).
			WillReturnError(errExec).WillReturnResult(sqlmock.NewResult(1, 1)).WithArgs(newAdvert.ID,
			newAdvert.UserID, newAdvert.AdvertTypeID, newAdvert.AdvertTypeSale, newAdvert.Title, newAdvert.Description,
			newAdvert.Phone, newAdvert.IsAgent, newAdvert.Priority)
	}
}

func (suite *UserRepoTestSuite) TestCreatePriceChange() {
	type args struct {
		adv     *models.PriceChange
		errExec error
		expExec bool
	}
	type want struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful create price change",
			args: args{
				adv: &models.PriceChange{
					ID:       uuid.NewV4(),
					AdvertID: uuid.NewV4(),
					Price:    1100000,
				},
				errExec: nil,
				expExec: true,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "fail create price change",
			args: args{
				adv: &models.PriceChange{
					ID:       uuid.NewV4(),
					AdvertID: uuid.NewV4(),
					Price:    1100000,
				},
				errExec: errors.New("some error"),
				expExec: true,
			},
			want: want{
				err: errors.New("some error"),
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.mock.ExpectBegin()
			tx, err := suite.db.Begin()
			if err != nil {
				suite.T().Fatal("Error beginning transaction:", err)
			}
			suite.setupMockCreatePriceChange(tt.args.adv, tt.args.errExec, tt.args.expExec)
			rep := repo.NewRepository(suite.db)
			gotErr := rep.CreatePriceChange(context.Background(), tx, tt.args.adv)
			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *UserRepoTestSuite) setupMockCreatePriceChange(newPriceChange *models.PriceChange, errExec error, expExec bool) {
	if expExec {
		suite.mock.ExpectExec(`INSERT INTO pricechanges \(id, advertid, price\) VALUES \(\$1, \$2, \$3\)`).
			WillReturnError(errExec).WillReturnResult(sqlmock.NewResult(1, 1)).WithArgs(
			newPriceChange.ID, newPriceChange.AdvertID, newPriceChange.Price)
	}
}

func (suite *UserRepoTestSuite) TestCreateHouse() {
	type args struct {
		adv     *models.House
		errExec error
		expExec bool
	}
	type want struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful create house",
			args: args{
				adv: &models.House{
					ID:            uuid.NewV4(),
					BuildingID:    uuid.NewV4(),
					AdvertTypeID:  uuid.NewV4(),
					CeilingHeight: 10,
					SquareArea:    124.123,
					SquareHouse:   124.124,
					BedroomCount:  2,
					StatusArea:    models.StatusAreaF,
					Cottage:       true,
					StatusHome:    models.StatusHomeRenovation,
				},
				errExec: nil,
				expExec: true,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "fail  create house",
			args: args{
				adv: &models.House{
					ID:            uuid.NewV4(),
					BuildingID:    uuid.NewV4(),
					AdvertTypeID:  uuid.NewV4(),
					CeilingHeight: 10,
					SquareArea:    124.123,
					SquareHouse:   124.124,
					BedroomCount:  2,
					StatusArea:    models.StatusAreaF,
					Cottage:       true,
					StatusHome:    models.StatusHomeRenovation,
				},
				errExec: errors.New("some error"),
				expExec: true,
			},
			want: want{
				err: errors.New("some error"),
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.mock.ExpectBegin()
			tx, err := suite.db.Begin()
			if err != nil {
				suite.T().Fatal("Error beginning transaction:", err)
			}
			suite.setupMockCreateHouse(tt.args.adv, tt.args.errExec, tt.args.expExec)
			rep := repo.NewRepository(suite.db)
			gotErr := rep.CreateHouse(context.Background(), tx, tt.args.adv)
			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *UserRepoTestSuite) setupMockCreateHouse(newHouse *models.House, errExec error, expExec bool) {
	if expExec {
		suite.mock.ExpectExec(`INSERT INTO houses \(id, buildingid, adverttypeid, ceilingheight, squarearea, squarehouse, bedroomcount, statusarea, cottage, statushome\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8, \$9, \$10\)`).
			WillReturnError(errExec).WillReturnResult(sqlmock.NewResult(1, 1)).WithArgs(
			newHouse.ID, newHouse.BuildingID, newHouse.AdvertTypeID, newHouse.CeilingHeight, newHouse.SquareArea,
			newHouse.SquareHouse, newHouse.BedroomCount, newHouse.StatusArea, newHouse.Cottage, newHouse.StatusHome)
	}
}

func (suite *UserRepoTestSuite) TestCreateFlat() {
	type args struct {
		adv     *models.Flat
		errExec error
		expExec bool
	}
	type want struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful create flat",
			args: args{
				adv: &models.Flat{
					ID:                uuid.NewV4(),
					BuildingID:        uuid.NewV4(),
					AdvertTypeID:      uuid.NewV4(),
					CeilingHeight:     10,
					SquareGeneral:     124.123,
					SquareResidential: 124.12224,
					Apartment:         true,
				},
				errExec: nil,
				expExec: true,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "fail  create flat",
			args: args{
				adv: &models.Flat{
					ID:                uuid.NewV4(),
					BuildingID:        uuid.NewV4(),
					AdvertTypeID:      uuid.NewV4(),
					CeilingHeight:     10,
					SquareGeneral:     124.123,
					SquareResidential: 124.12224,
					Apartment:         true,
				},
				errExec: errors.New("some error"),
				expExec: true,
			},
			want: want{
				err: errors.New("some error"),
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.mock.ExpectBegin()
			tx, err := suite.db.Begin()
			if err != nil {
				suite.T().Fatal("Error beginning transaction:", err)
			}
			suite.setupMockCreateFlat(tt.args.adv, tt.args.errExec, tt.args.expExec)
			rep := repo.NewRepository(suite.db)
			gotErr := rep.CreateFlat(context.Background(), tx, tt.args.adv)
			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *UserRepoTestSuite) setupMockCreateFlat(newFlat *models.Flat, errExec error, expExec bool) {
	if expExec {
		suite.mock.ExpectExec(`INSERT INTO flats \(id, buildingid, adverttypeid, floor, ceilingheight, squaregeneral, roomcount, squareresidential, apartament\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8, \$9\)`).
			WillReturnError(errExec).WillReturnResult(sqlmock.NewResult(1, 1)).WithArgs(
			newFlat.ID, newFlat.BuildingID, newFlat.AdvertTypeID, newFlat.Floor, newFlat.CeilingHeight,
			newFlat.SquareGeneral, newFlat.RoomCount, newFlat.SquareResidential, newFlat.Apartment)
	}
}

func (suite *UserRepoTestSuite) TestCreateBuilding() {
	type args struct {
		building          *models.Building
		errExec, errQuery error
		expExec, expQuery bool
	}
	type want struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful create building",
			args: args{
				building: &models.Building{
					ID:           uuid.NewV4(),
					ComplexID:    uuid.NewV4(),
					Floor:        5,
					Material:     models.MaterialStalinsky,
					Address:      "123 Main Street",
					AddressPoint: "40.7128° N, 74.0060° W",
					YearCreation: 2000,
					//DateCreation: time.Now(),
					IsDeleted: false,
				},
				errExec:  nil,
				errQuery: nil,
				expExec:  true,
				expQuery: true,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "fail create building",
			args: args{
				building: &models.Building{
					ID:           uuid.NewV4(),
					ComplexID:    uuid.NewV4(),
					Floor:        5,
					Material:     models.MaterialStalinsky,
					Address:      "123 Main Street",
					AddressPoint: "40.7128° N, 74.0060° W",
					YearCreation: 2000,
					//DateCreation: time.Now(),
					IsDeleted: false,
				},
				errExec:  errors.New("error"),
				errQuery: nil,
				expExec:  true,
				expQuery: false,
			},
			want: want{
				err: errors.New("error"),
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.mock.ExpectBegin()
			tx, err := suite.db.Begin()
			if err != nil {
				suite.T().Fatal("Error beginning transaction:", err)
			}
			suite.setupMockCreateBuilding(tt.args.building, tt.args.errExec, tt.args.errQuery, tt.args.expExec, tt.args.expQuery)
			rep := repo.NewRepository(suite.db)
			gotErr := rep.CreateBuilding(context.Background(), tx, tt.args.building)
			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}
func (suite *UserRepoTestSuite) setupMockCreateBuilding(newBuilding *models.Building, errExec, errQuery error, expExec, epxQuery bool) {
	rows := sqlmock.NewRows([]string{"id", "complexId", "floor", "material", "adress", "adressPoint", "yearCreation"})
	rows = rows.AddRow(newBuilding.ID, newBuilding.ComplexID, newBuilding.Floor, newBuilding.Material, newBuilding.Address, newBuilding.AddressPoint, newBuilding.YearCreation)
	if expExec {
		suite.mock.ExpectExec(`INSERT INTO buildings \(id, floor, material, adress, adresspoint, yearcreation\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7\)`).
			WithArgs(newBuilding.ID, newBuilding.Floor, newBuilding.Material, newBuilding.Address, newBuilding.AddressPoint, newBuilding.YearCreation).
			WillReturnError(errExec).WillReturnResult(sqlmock.NewResult(1, 1))
	}
}

func (suite *UserRepoTestSuite) TestCheckExistsBuilding1() {
	type args struct {
		errExec, errQuery error
		expExec, expQuery bool
	}
	type want struct {
		err   error
		build *models.Building
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful get building",
			args: args{
				errExec:  nil,
				errQuery: nil,
				expExec:  true,
				expQuery: true,
			},
			want: want{
				build: &models.Building{
					ID:           uuid.NewV4(),
					ComplexID:    uuid.NewV4(),
					Floor:        2,
					Material:     models.MaterialStalinsky,
					Address:      "address",
					AddressPoint: "point",
				},
				err: nil,
			},
		},
		{
			name: "fail get building",
			args: args{
				errExec:  nil,
				errQuery: errors.New("error"),
				expExec:  true,
				expQuery: true,
			},
			want: want{
				build: &models.Building{
					ID:           uuid.NewV4(),
					ComplexID:    uuid.NewV4(),
					Floor:        2,
					Material:     models.MaterialStalinsky,
					Address:      "address",
					AddressPoint: "point",
				},
				err: errors.New("error"),
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.setupMockCheckExistsBuilding1(tt.want.build, tt.args.errQuery, tt.args.expQuery)
			rep := repo.NewRepository(suite.db)
			gotBuild, gotErr := rep.CheckExistsBuilding(context.Background(), tt.want.build.Address)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Empty(gotBuild)
			} else {
				suite.Assert().Equal(tt.want.build.ID, gotBuild.ID)
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}
func (suite *UserRepoTestSuite) setupMockCheckExistsBuilding1(building *models.Building, errQuery error, epxQuery bool) {
	rows := sqlmock.NewRows([]string{"id"})
	rows = rows.AddRow(building.ID)

	if epxQuery {
		suite.mock.ExpectQuery(`SELECT id FROM buildings WHERE adress = \$1`).
			WithArgs(building.Address).
			WillReturnRows(rows).WillReturnError(errQuery)
	}
}

func (suite *UserRepoTestSuite) TestCheckExistsBuilding2() {
	type args struct {
		errExec, errQuery error
		expExec, expQuery bool
		pageS             int
	}
	type want struct {
		err   error
		build *models.BuildingsExistData
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful check building",
			args: args{
				errExec:  nil,
				errQuery: nil,
				expExec:  true,
				expQuery: true,
				pageS:    2,
			},
			want: want{
				build: &models.BuildingsExistData{
					ID: uuid.NewV4(),
					//ComplexID:    uuid.NewV4(),
					Floor:        2,
					Material:     models.MaterialStalinsky,
					Address:      "address",
					AddressPoint: "point",
				},
				err: nil,
			},
		},
		{
			name: "fail check building",
			args: args{
				errExec:  nil,
				errQuery: errors.New("error"),
				expExec:  true,
				expQuery: true,
				pageS:    2,
			},
			want: want{
				build: &models.BuildingsExistData{
					ID: uuid.NewV4(),
					//ComplexID:    uuid.NewV4(),
					Floor:        2,
					Material:     models.MaterialStalinsky,
					Address:      "address",
					AddressPoint: "point",
				},
				err: errors.New("error"),
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.setupMockCheckExistsBuilding2(tt.want.build, tt.args.pageS, tt.args.errQuery, tt.args.expQuery)
			rep := repo.NewRepository(suite.db)
			gotBuild, gotErr := rep.CheckExistsBuildings(context.Background(), tt.args.pageS, tt.want.build.Address)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Empty(gotBuild)
			} else {
				suite.Assert().Equal(tt.want.build, gotBuild[0])
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *UserRepoTestSuite) setupMockCheckExistsBuilding2(building *models.BuildingsExistData, pageSize int, errQuery error, epxQuery bool) {
	rows := sqlmock.NewRows([]string{"id", "floor", "material", "adress", "adressPoint", "yearCreation", "complexName"})
	rows = rows.AddRow(building.ID, building.Floor, building.Material, building.Address, building.AddressPoint, building.YearCreation, building.ComplexName)

	if epxQuery {
		suite.mock.ExpectQuery(`SELECT b.id, b.floor, COALESCE\(b.material, 'Brick'\), b.adress, b.adresspoint, b.yearcreation, COALESCE\(cx.name, ''\) FROM buildings AS b LEFT JOIN complexes AS cx ON b.complexid\=cx.id WHERE b.adress ILIKE \$1 LIMIT \$2`).
			WithArgs("%"+building.Address+"%", pageSize).
			WillReturnRows(rows).WillReturnError(errQuery)
	}
}

func (suite *UserRepoTestSuite) TestSelectImages() {
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

func (suite *UserRepoTestSuite) setupMockSelectImage(advertID uuid.UUID, images []*models.ImageResp) {
	rows := sqlmock.NewRows([]string{"id", "photo", "priority"})
	for _, image := range images {
		rows = rows.AddRow(image.ID, image.Photo, image.Priority)
	}
	suite.mock.ExpectQuery(`SELECT id, photo, priority FROM images WHERE advertid = \$1 AND isdeleted = false`).
		WithArgs(advertID).
		WillReturnRows(rows)
}

func (suite *UserRepoTestSuite) TestCheckGetTypeAdvertById() {
	type args struct {
		errExec, errQuery error
		expExec, expQuery bool
		id                uuid.UUID
	}
	type want struct {
		err        error
		advertType models.AdvertTypeAdvert
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful get typeAdvert",
			args: args{
				id:       uuid.NewV4(),
				errExec:  nil,
				errQuery: nil,
				expExec:  true,
				expQuery: true,
			},
			want: want{
				advertType: models.AdvertTypeFlat,
				err:        nil,
			},
		},
		{
			name: "fail get typeAdvert",
			args: args{
				id:       uuid.NewV4(),
				errExec:  nil,
				errQuery: errors.New("error"),
				expExec:  true,
				expQuery: true,
			},
			want: want{
				advertType: models.AdvertTypeFlat,
				err:        errors.New("error"),
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.setupMockGetTypeAdvertById(tt.want.advertType, tt.args.id, tt.args.errQuery, tt.args.expQuery)
			rep := repo.NewRepository(suite.db)
			gotType, gotErr := rep.GetTypeAdvertById(context.Background(), tt.args.id)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Empty(gotType)
			} else {
				suite.Assert().Equal(&tt.want.advertType, gotType)
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}
func (suite *UserRepoTestSuite) setupMockGetTypeAdvertById(advertType models.AdvertTypeAdvert, idi uuid.UUID, errQuery error, epxQuery bool) {
	rows := sqlmock.NewRows([]string{"advertType"})
	rows = rows.AddRow(advertType)

	if epxQuery {
		suite.mock.ExpectQuery(`SELECT at.adverttype FROM adverts AS a JOIN adverttypes AS at ON a.adverttypeid\=at.id WHERE a.id \= \$1`).
			WithArgs(idi).
			WillReturnRows(rows).WillReturnError(errQuery)
	}
}

func (suite *UserRepoTestSuite) TestCheckGetHouseAdvertById() {
	type args struct {
		errExec, errQuery error
		expExec, expQuery bool
	}
	type want struct {
		err        error
		advertData *models.AdvertData
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful get advert data",
			args: args{
				errExec:  nil,
				errQuery: nil,
				expExec:  true,
				expQuery: true,
			},
			want: want{
				advertData: &models.AdvertData{
					ID:           uuid.NewV4(),
					TypeAdvert:   "House",
					TypeSale:     "Sale",
					Title:        "Beautiful House for Sale",
					Description:  "Spacious house with a large garden",
					Price:        100000,
					Phone:        "123-456-7890",
					IsAgent:      true,
					Address:      "123 Main St, Cityville",
					AddressPoint: "Coordinates",
					//Images:       []*models.ImageResp{},
					HouseProperties: &models.HouseProperties{
						CeilingHeight: 2.7,
						SquareArea:    200.5,
						SquareHouse:   180.0,
						BedroomCount:  4,
						StatusArea:    "Living room, kitchen, bedroom",
						Cottage:       false,
						StatusHome:    "New",
						Floor:         2,
					},
					ComplexProperties: &models.ComplexAdvertProperties{
						ComplexId:    "1234",
						NameComplex:  "Luxury Estates",
						PhotoCompany: "luxury_estates.jpg",
						NameCompany:  "Elite Realty",
					},
					//YearCreation: time.Now().Year(),
					Material: "Brick",
					//DateCreation: time.Now(),
				},
				err: nil,
			},
		},
		{
			name: "fail get advert data",
			args: args{
				errExec:  nil,
				errQuery: errors.New("error"),
				expExec:  true,
				expQuery: true,
			},
			want: want{
				advertData: &models.AdvertData{
					ID:           uuid.NewV4(),
					TypeAdvert:   "House",
					TypeSale:     "Sale",
					Title:        "Beautiful House for Sale",
					Description:  "Spacious house with a large garden",
					Price:        100000,
					Phone:        "123-456-7890",
					IsAgent:      true,
					Address:      "123 Main St, Cityville",
					AddressPoint: "Coordinates",
					//Images:       []*models.ImageResp{},
					HouseProperties: &models.HouseProperties{
						CeilingHeight: 2.7,
						SquareArea:    200.5,
						SquareHouse:   180.0,
						BedroomCount:  4,
						StatusArea:    "Living room, kitchen, bedroom",
						Cottage:       false,
						StatusHome:    "New",
						Floor:         2,
					},
					ComplexProperties: &models.ComplexAdvertProperties{
						ComplexId:    "1234",
						NameComplex:  "Luxury Estates",
						PhotoCompany: "luxury_estates.jpg",
						NameCompany:  "Elite Realty",
					},
					//YearCreation: time.Now().Year(),
					Material: "Brick",
					//DateCreation: time.Now(),
				},
				err: errors.New("error"),
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.setupMockGetHouseAdvertById(tt.want.advertData, tt.want.advertData.ID, tt.args.errQuery, tt.args.expQuery)
			rep := repo.NewRepository(suite.db)
			gotAdvertData, gotErr := rep.GetHouseAdvertById(context.Background(), tt.want.advertData.ID)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Empty(gotAdvertData)
			} else {
				suite.Assert().Equal(tt.want.advertData, gotAdvertData)
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}
func (suite *UserRepoTestSuite) setupMockGetHouseAdvertById(advertData *models.AdvertData, idi uuid.UUID, errQuery error, epxQuery bool) {
	rows := sqlmock.NewRows([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15",
		"16", "17", "18", "19", "20", "21", "22", "23", "24", "25"})
	rows = rows.AddRow(advertData.ID,
		advertData.TypeAdvert,
		advertData.TypeSale,
		advertData.Title,
		advertData.Description,
		advertData.Price,
		advertData.Phone,
		advertData.IsAgent,
		advertData.Address,
		advertData.AddressPoint,
		advertData.HouseProperties.CeilingHeight,
		advertData.HouseProperties.SquareArea,
		advertData.HouseProperties.SquareHouse,
		advertData.HouseProperties.BedroomCount,
		advertData.HouseProperties.StatusArea,
		advertData.HouseProperties.Cottage,
		advertData.HouseProperties.StatusHome,
		advertData.HouseProperties.Floor,
		advertData.YearCreation,
		advertData.Material,
		advertData.DateCreation,
		advertData.ComplexProperties.ComplexId,
		advertData.ComplexProperties.PhotoCompany,
		advertData.ComplexProperties.NameCompany,
		advertData.ComplexProperties.NameComplex)

	query := `
	SELECT
        a.id,
        at.adverttype,
        a.adverttypeplacement,
        a.title,
        a.description,
        pc.price,
        a.phone,
        a.isagent,
        b.adress,
        b.adresspoint,
        h.ceilingheight,
        h.squarearea,
        h.squarehouse,
        h.bedroomcount,
        h.statusarea,
        h.cottage,
        h.statushome,
        b.floor,
        b.yearcreation,
        COALESCE(b.material, 'Brick') as material,
        a.datecreation,
        cx.id AS complexid,
        c.photo AS companyphoto,
        c.name AS companyname,
        cx.name AS complexname
    FROM
        adverts AS a
    JOIN
        adverttypes AS at ON a.adverttypeid = at.id
    JOIN
        houses AS h ON h.adverttypeid = at.id
    JOIN
        buildings AS b ON h.buildingid = b.id
    LEFT JOIN
        complexes AS cx ON b.complexid = cx.id
    LEFT JOIN
        companies AS c ON cx.companyid = c.id
    JOIN
        LATERAL (
            SELECT *
            FROM pricechanges AS pc
            WHERE pc.advertid = a.id
            ORDER BY pc.datecreation DESC
            LIMIT 1
        ) AS pc ON TRUE
    WHERE
        a.id = $1 AND a.isdeleted = FALSE;`

	escapedQuery := regexp.QuoteMeta(query)
	suite.mock.ExpectQuery(escapedQuery).
		WithArgs(advertData.ID).
		WillReturnRows(rows).WillReturnError(errQuery)
}

func (suite *UserRepoTestSuite) TestCheckExistsFlat() {
	type args struct {
		errExec, errQuery error
		expExec, expQuery bool
	}
	type want struct {
		err  error
		flat *models.Flat
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful check flat",
			args: args{
				errExec:  nil,
				errQuery: nil,
				expExec:  true,
				expQuery: true,
			},
			want: want{
				flat: &models.Flat{
					ID:    uuid.NewV4(),
					Floor: 2,
				},
				err: nil,
			},
		},
		{
			name: "fail check flat",
			args: args{
				errExec:  nil,
				errQuery: errors.New("error"),
				expExec:  true,
				expQuery: true,
			},
			want: want{
				flat: &models.Flat{
					ID:    uuid.NewV4(),
					Floor: 2,
				},
				err: errors.New("error"),
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupMockCheckExistsFlat(tt.want.flat, tt.args.errQuery, tt.args.expQuery)
			rep := repo.NewRepository(suite.db)
			gotFlat, gotErr := rep.CheckExistsFlat(context.Background(), tt.want.flat.ID)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Empty(gotFlat)
			} else {
				suite.Assert().Equal(tt.want.flat.ID, gotFlat.ID)
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}
func (suite *UserRepoTestSuite) SetupMockCheckExistsFlat(flat *models.Flat, errQuery error, epxQuery bool) {
	rows := sqlmock.NewRows([]string{"id"})
	rows = rows.AddRow(flat.ID)
	query := `SELECT f.id FROM adverts AS a JOIN adverttypes AS at ON a.adverttypeid=at.id JOIN flats AS f ON f.adverttypeid=at.id WHERE a.id = $1`

	escapedQuery := regexp.QuoteMeta(query)
	if epxQuery {
		suite.mock.ExpectQuery(escapedQuery).
			WithArgs(flat.ID).
			WillReturnRows(rows).WillReturnError(errQuery)
	}
}

func (suite *UserRepoTestSuite) TestCheckExistsHouse() {
	type args struct {
		errExec, errQuery error
		expExec, expQuery bool
		advId             uuid.UUID
	}
	type want struct {
		err   error
		house *models.House
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful check house",
			args: args{
				errExec:  nil,
				errQuery: nil,
				expExec:  true,
				expQuery: true,
				advId:    uuid.NewV4(),
			},
			want: want{
				house: &models.House{
					ID:          uuid.NewV4(),
					SquareHouse: 241.214,
				},
				err: nil,
			},
		},
		{
			name: "fail check house",
			args: args{
				advId:    uuid.NewV4(),
				errExec:  nil,
				errQuery: errors.New("error"),
				expExec:  true,
				expQuery: true,
			},
			want: want{
				house: &models.House{
					ID:          uuid.NewV4(),
					SquareHouse: 241.214,
				},
				err: errors.New("error"),
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.SetupMockCheckExistsHouse(tt.want.house, tt.args.advId, tt.args.errQuery, tt.args.expQuery)
			rep := repo.NewRepository(suite.db)
			gotHouse, gotErr := rep.CheckExistsHouse(context.Background(), tt.args.advId)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Empty(gotHouse)
			} else {
				suite.Assert().Equal(tt.want.house.ID, gotHouse.ID)
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}
func (suite *UserRepoTestSuite) SetupMockCheckExistsHouse(house *models.House, advId uuid.UUID, errQuery error, epxQuery bool) {
	rows := sqlmock.NewRows([]string{"id"})
	rows = rows.AddRow(house.ID)
	query := `SELECT h.id FROM adverts AS a JOIN adverttypes AS at ON a.adverttypeid=at.id JOIN houses AS h ON h.adverttypeid=at.id WHERE a.id = $1;`

	escapedQuery := regexp.QuoteMeta(query)
	if epxQuery {
		suite.mock.ExpectQuery(escapedQuery).
			WithArgs(advId).
			WillReturnRows(rows).WillReturnError(errQuery)
	}
}

/*
	func (suite *UserRepoTestSuite) TestUserRepo_DeleteFlatAdvertById() {
		advertId := uuid.NewV4()
		ctx := context.Background()
		suite.mock.ExpectBegin()
		tx, err := suite.db.Begin()
		suite.NoError(err)

		queryGetIdTables := `
			SELECT
		at.id as adverttypeid,
			f.id as flatid
		FROM
		adverts AS a
		JOIN
		adverttypes AS at ON a.adverttypeid = at.id
		JOIN
		flats AS f ON f.adverttypeid = at.id
		WHERE a.id=$1;`

		quotedQueryGetIdTables := regexp.QuoteMeta(queryGetIdTables)
		suite.mock.ExpectQuery(quotedQueryGetIdTables).WithArgs(advertId).
			WillReturnRows(sqlmock.NewRows([]string{"adverttypeid", "flatid"}).AddRow(uuid.NewV4(), uuid.NewV4()))

		// Add expectations for other queries

		// Call the function under test with the mock transaction
		rep := repo.NewRepository(suite.db)
		err = rep.DeleteFlatAdvertById(ctx, tx, advertId)
		suite.Assert().NoError(err)

		// End the transaction and verify expectations were met
		err = tx.Commit()
		suite.Assert().NoError(err)

		err = suite.mock.ExpectationsWereMet()
		suite.Assert().NoError(err)
	}
*/
func (suite *UserRepoTestSuite) TestUserRepo_DeleteFlatAdvertById() {
	advertId := uuid.NewV4()
	advertTypeId := uuid.NewV4()
	flatId := uuid.NewV4()
	ctx := context.Background()

	suite.mock.ExpectBegin()
	tx, err := suite.db.Begin()
	suite.NoError(err)

	quotedQueryGetIdTables := regexp.QuoteMeta(`
        SELECT
            at.id as adverttypeid,
            f.id as flatid
        FROM
            adverts AS a
        JOIN
            adverttypes AS at ON a.adverttypeid = at.id
        JOIN
            flats AS f ON f.adverttypeid = at.id
        WHERE a.id=$1;`)
	suite.mock.ExpectQuery(quotedQueryGetIdTables).WithArgs(advertId).
		WillReturnRows(sqlmock.NewRows([]string{"adverttypeid", "flatid"}).AddRow(advertTypeId, flatId))

	queryDeleteAdvertById := regexp.QuoteMeta(`UPDATE adverts SET isdeleted=true WHERE id=$1;`)
	queryDeleteAdvertTypeById := regexp.QuoteMeta(`UPDATE adverttypes SET isdeleted=true WHERE id=$1;`)
	queryDeleteFlatById := regexp.QuoteMeta(`UPDATE flats SET isdeleted=true WHERE id=$1;`)
	queryDeletePriceChanges := regexp.QuoteMeta(`UPDATE pricechanges SET isdeleted=true WHERE advertid=$1;`)
	queryDeleteImages := regexp.QuoteMeta(`UPDATE images SET isdeleted=true WHERE advertid=$1;`)

	suite.mock.ExpectExec(queryDeleteAdvertById).WithArgs(advertId).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectExec(queryDeleteAdvertTypeById).WithArgs(advertTypeId).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectExec(queryDeleteFlatById).WithArgs(flatId).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectExec(queryDeletePriceChanges).WithArgs(advertId).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectExec(queryDeleteImages).WithArgs(advertId).WillReturnResult(sqlmock.NewResult(1, 1))

	rep := repo.NewRepository(suite.db)
	err = rep.DeleteFlatAdvertById(ctx, tx, advertId)
	suite.Assert().NoError(err)

	suite.mock.ExpectCommit()
	err = tx.Commit()
	suite.Assert().NoError(err)

	err = suite.mock.ExpectationsWereMet()
	suite.Assert().NoError(err)
}

func (suite *UserRepoTestSuite) TestUserRepo_DeleteHouseAdvertById() {
	advertId := uuid.NewV4()
	advertTypeId := uuid.NewV4()
	houseId := uuid.NewV4()
	ctx := context.Background()

	suite.mock.ExpectBegin()
	tx, err := suite.db.Begin()
	suite.NoError(err)

	quotedQueryGetIdTables := regexp.QuoteMeta(
		`SELECT
	at.id as adverttypeid,
		h.id as houseid
	FROM
	adverts AS a
	JOIN
	adverttypes AS at ON a.adverttypeid = at.id
	JOIN
	houses AS h ON h.adverttypeid = at.id
	WHERE a.id=$1;`)
	suite.mock.ExpectQuery(quotedQueryGetIdTables).WithArgs(advertId).
		WillReturnRows(sqlmock.NewRows([]string{"adverttypeid", "houseid"}).AddRow(advertTypeId, houseId))

	queryDeleteAdvertById := regexp.QuoteMeta(`UPDATE adverts SET isdeleted=true WHERE id=$1;`)
	queryDeleteAdvertTypeById := regexp.QuoteMeta(`UPDATE adverttypes SET isdeleted=true WHERE id=$1;`)
	queryDeleteHouseById := regexp.QuoteMeta(`UPDATE houses SET isdeleted=true WHERE id=$1;`)
	queryDeletePriceChanges := regexp.QuoteMeta(`UPDATE pricechanges SET isdeleted=true WHERE advertid=$1;`)
	queryDeleteImages := regexp.QuoteMeta(`UPDATE images SET isdeleted=true WHERE advertid=$1;`)

	suite.mock.ExpectExec(queryDeleteAdvertById).WithArgs(advertId).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectExec(queryDeleteAdvertTypeById).WithArgs(advertTypeId).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectExec(queryDeleteHouseById).WithArgs(houseId).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectExec(queryDeletePriceChanges).WithArgs(advertId).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectExec(queryDeleteImages).WithArgs(advertId).WillReturnResult(sqlmock.NewResult(1, 1))

	rep := repo.NewRepository(suite.db)
	err = rep.DeleteHouseAdvertById(ctx, tx, advertId)
	suite.Assert().NoError(err)

	suite.mock.ExpectCommit()
	err = tx.Commit()
	suite.Assert().NoError(err)

	err = suite.mock.ExpectationsWereMet()
	suite.Assert().NoError(err)
}
