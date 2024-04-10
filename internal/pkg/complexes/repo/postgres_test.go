package repo_test

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/complexes/repo"
	"context"
	"database/sql"
	"errors"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
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

func (suite *UserRepoTestSuite) TestCreateComplex() {
	type args struct {
		complex           *models.Complex
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
			name: "successful create complex",
			args: args{
				complex: &models.Complex{
					ID:          uuid.NewV4(),
					CompanyId:   uuid.NewV4(),
					Name:        "Complex Name",
					Address:     "Complex Address",
					Description: "Complex Description",
					//DateBeginBuild:         time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
					//DateEndBuild:           time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
					WithoutFinishingOption: true,
					FinishingOption:        true,
					PreFinishingOption:     true,
					ClassHousing:           models.ClassHouseBusiness,
					Parking:                true,
					Security:               true,
					//DateCreation:           time.Now(),
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
			name: "fail create complex",
			args: args{
				complex: &models.Complex{
					ID:          uuid.NewV4(),
					CompanyId:   uuid.NewV4(),
					Name:        "Complex Name",
					Address:     "Complex Address",
					Description: "Complex Description",
					//DateBeginBuild:         time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
					//DateEndBuild:           time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
					WithoutFinishingOption: true,
					FinishingOption:        true,
					PreFinishingOption:     true,
					ClassHousing:           models.ClassHouseBusiness,
					Parking:                true,
					Security:               true,
					//DateCreation:           time.Now(),
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
		{
			name: "fail create complex2",
			args: args{
				complex: &models.Complex{
					ID:          uuid.NewV4(),
					CompanyId:   uuid.NewV4(),
					Name:        "Complex Name",
					Address:     "Complex Address",
					Description: "Complex Description",
					//DateBeginBuild:         time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
					//DateEndBuild:           time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
					WithoutFinishingOption: true,
					FinishingOption:        true,
					PreFinishingOption:     true,
					ClassHousing:           models.ClassHouseBusiness,
					Parking:                true,
					Security:               true,
					//DateCreation:           time.Now(),
					IsDeleted: false,
				},
				errExec:  nil,
				errQuery: errors.New("error"),
				expExec:  true,
				expQuery: true,
			},
			want: want{
				err: errors.New("error"),
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.setupMockCreateCompany(tt.args.complex, tt.args.errExec, tt.args.errQuery, tt.args.expExec, tt.args.expQuery)
			rep := repo.NewRepository(suite.db, &zap.Logger{})
			newComplex, gotErr := rep.CreateComplex(context.Background(), tt.args.complex)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Empty(newComplex)
			} else {
				suite.Assert().Equal(tt.args.complex, newComplex)
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}
func (suite *UserRepoTestSuite) setupMockCreateCompany(complex *models.Complex, errExec, errQuery error, expExec, epxQuery bool) {
	rows := sqlmock.NewRows([]string{"id", "companyId", "name", "address", "photo", "description", "dateBeginBuild", "dateEndBuild", "withoutFinishingOption", "finishingOption", "preFinishingOption", "classHousing", "parking", "security"})
	rows = rows.AddRow(complex.ID, complex.CompanyId, complex.Name, complex.Address, "", complex.Description, complex.DateBeginBuild,
		complex.DateEndBuild, complex.WithoutFinishingOption, complex.FinishingOption, complex.PreFinishingOption, complex.ClassHousing, complex.Parking, complex.Security)
	if expExec {
		suite.mock.ExpectExec(`INSERT INTO complexes \(id, companyid, name, adress, photo, description, datebeginbuild, dateendbuild, withoutfinishingoption, finishingoption, prefinishingoption, classhousing, parking, security\) VALUES \(\$1, \$2, \$3, \$4, '', \$5, \$6, \$7, \$8, \$9, \$10, \$11, \$12, \$13\);`).
			WithArgs(complex.ID, complex.CompanyId, complex.Name, complex.Address, complex.Description, complex.DateBeginBuild,
				complex.DateEndBuild, complex.WithoutFinishingOption, complex.FinishingOption, complex.PreFinishingOption, complex.ClassHousing, complex.Parking, complex.Security).
			WillReturnError(errExec).WillReturnResult(sqlmock.NewResult(1, 1))
	}
	if epxQuery {
		suite.mock.ExpectQuery(`SELECT id, companyid, name, adress, photo, description, datebeginbuild, dateendbuild, withoutfinishingoption, finishingoption, prefinishingoption, classhousing, parking, security FROM complexes WHERE id = \$1`).
			WithArgs(complex.ID).WillReturnRows(rows).WillReturnError(errQuery)
	}
}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
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
		{
			name: "fail create building2",
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
				errQuery: errors.New("error"),
				expExec:  true,
				expQuery: true,
			},
			want: want{
				err: errors.New("error"),
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.setupMockCreateBuilding(tt.args.building, tt.args.errExec, tt.args.errQuery, tt.args.expExec, tt.args.expQuery)
			rep := repo.NewRepository(suite.db, &zap.Logger{})
			newComplex, gotErr := rep.CreateBuilding(context.Background(), tt.args.building)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Empty(newComplex)
			} else {
				suite.Assert().Equal(tt.args.building, newComplex)
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}
func (suite *UserRepoTestSuite) setupMockCreateBuilding(building *models.Building, errExec, errQuery error, expExec, epxQuery bool) {
	rows := sqlmock.NewRows([]string{"id", "complexId", "floor", "material", "adress", "adressPoint", "yearCreation"})
	rows = rows.AddRow(building.ID, building.ComplexID, building.Floor, building.Material, building.Address, building.AddressPoint, building.YearCreation)
	if expExec {
		suite.mock.ExpectExec(`INSERT INTO buildings \(id, complexId, floor, material, adress, adressPoint, yearCreation\)
	VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7\);`).
			WithArgs(building.ID, building.ComplexID, building.Floor, building.Material, building.Address, building.AddressPoint, building.YearCreation).
			WillReturnError(errExec).WillReturnResult(sqlmock.NewResult(1, 1))
	}
	if epxQuery {
		suite.mock.ExpectQuery(`SELECT id, complexId, floor, material, adress, adressPoint, yearCreation FROM buildings WHERE id = \$1`).
			WithArgs(building.ID).WillReturnRows(rows).WillReturnError(errQuery)
	}
}

func (suite *UserRepoTestSuite) TestUpdateComplexPhoto() {
	type args struct {
		id       uuid.UUID
		filename string
		errExec  error
		expExec  bool
	}
	type want struct {
		err error
	}
	id1 := uuid.NewV4()
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful update complex photo",
			args: args{
				id:       id1,
				filename: "file/name",
				errExec:  nil,
				expExec:  true,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "fail update photo complex",
			args: args{
				id:       id1,
				filename: "file/name",
				errExec:  errors.New("some error"),
				expExec:  true,
			},
			want: want{
				err: errors.New("some error"),
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.setupMockUpdateComplexPhoto(tt.args.id, tt.args.filename, tt.args.errExec, tt.args.expExec)
			rep := repo.NewRepository(suite.db, &zap.Logger{})
			newFilename, gotErr := rep.UpdateComplexPhoto(tt.args.id, tt.args.filename)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Empty(newFilename)
			} else {
				suite.Assert().Equal(tt.args.filename, newFilename)
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *UserRepoTestSuite) setupMockUpdateComplexPhoto(idi uuid.UUID, filen string, errExec error, expExec bool) {
	if expExec {
		suite.mock.ExpectExec(`UPDATE complexes SET photo = \$1 WHERE id = \$2`).
			WillReturnError(errExec).WillReturnResult(sqlmock.NewResult(1, 1)).WithArgs(filen, idi)
	}
}

func (suite *UserRepoTestSuite) TestGetComplexById() {
	type args struct {
		// ctx context.Context
		errExec, errQuery error
		expExec, expQuery bool
	}
	type want struct {
		// user *models.User
		err      error
		compData *models.ComplexData
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful get company",
			args: args{

				errExec:  nil,
				errQuery: nil,
				expExec:  true,
				expQuery: true,
			},
			want: want{
				compData: &models.ComplexData{ // дата не возвращается!
					ID:          uuid.NewV4(),
					CompanyId:   uuid.NewV4(),
					Name:        "Name",
					Address:     "address",
					Photo:       "Photo",
					Description: "descr",
				},
				err: nil,
			},
		},
		{
			name: "fail get company",
			args: args{
				errExec:  nil,
				errQuery: errors.New("error"),
				expExec:  true,
				expQuery: true,
			},
			want: want{
				compData: &models.ComplexData{ // дата не возвращается!
					ID:          uuid.NewV4(),
					CompanyId:   uuid.NewV4(),
					Name:        "Name",
					Address:     "address",
					Photo:       "Photo",
					Description: "descr",
				},
				err: errors.New("error"),
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.setupMockGetComplexById(tt.want.compData, tt.args.errQuery, tt.args.expQuery)
			rep := repo.NewRepository(suite.db, &zap.Logger{})
			updCompany, gotErr := rep.GetComplexById(context.Background(), tt.want.compData.ID)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Empty(updCompany)
			} else {
				suite.Assert().Equal(tt.want.compData, updCompany)
			}
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}
func (suite *UserRepoTestSuite) setupMockGetComplexById(complexData *models.ComplexData, errQuery error, epxQuery bool) {
	rows := sqlmock.NewRows([]string{"id", "companyId", "name", "address", "photo", "description", "dateBeginBuild", "dateEndBuild",
		"withoutFinishingOption", "finishingOption", "preFinishingOption", "classHousing", "parking", "security"})
	rows = rows.AddRow(complexData.ID, complexData.CompanyId, complexData.Name, complexData.Address, complexData.Photo,
		complexData.Description, complexData.DateBeginBuild, complexData.DateEndBuild, complexData.WithoutFinishingOption,
		complexData.FinishingOption, complexData.PreFinishingOption, complexData.ClassHousing, complexData.Parking,
		complexData.Security)

	if epxQuery {
		suite.mock.ExpectQuery(`SELECT id, companyid, name, adress, photo, description, datebeginbuild, dateendbuild, withoutfinishingoption, finishingoption, prefinishingoption, classhousing, parking, security FROM complexes WHERE id = \$1`).
			WithArgs(complexData.ID).
			WillReturnRows(rows).WillReturnError(errQuery)
	}
}
func TestBeginTx(t *testing.T) {
	fakeDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer fakeDB.Close()
	rep := repo.NewRepository(fakeDB, &zap.Logger{})
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
	rep := repo.NewRepository(fakeDB, &zap.Logger{})
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
			rep := repo.NewRepository(suite.db, &zap.Logger{})
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
			rep := repo.NewRepository(suite.db, &zap.Logger{})
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
			rep := repo.NewRepository(suite.db, &zap.Logger{})
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
			rep := repo.NewRepository(suite.db, &zap.Logger{})
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
			rep := repo.NewRepository(suite.db, &zap.Logger{})
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
