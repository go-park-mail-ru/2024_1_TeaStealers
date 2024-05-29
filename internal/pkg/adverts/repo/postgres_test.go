package repo_test

/*
import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts/repo"
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"math/rand"
	"regexp"
	"testing"
	"time"
)

type AdvertRepoTestSuite struct {
	suite.Suite
	db   *sql.DB
	mock sqlmock.Sqlmock
}

func (suite *AdvertRepoTestSuite) SetupTest() {
	var err error
	suite.db, suite.mock, err = sqlmock.New()
	suite.Require().NoError(err)
}

func (suite *AdvertRepoTestSuite) TearDownTest() {
	suite.mock.ExpectClose()
	suite.Require().NoError(suite.db.Close())
}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(AdvertRepoTestSuite))
}

func TestBeginTx(t *testing.T) {
	fakeDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer fakeDB.Close()
	logger := zap.Must(zap.NewDevelopment())
	rep := repo.NewRepository(fakeDB, logger)
	// ctx := context.Background()
	// tx := new(sql.Tx)
	mock.ExpectBegin().WillReturnError(nil)

	tx, err := rep.BeginTx(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()))
	assert.NoError(t, err)
	assert.NotEmpty(t, tx)
}

func TestBeginTxFail(t *testing.T) {
	fakeDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer fakeDB.Close()
	logger := zap.Must(zap.NewDevelopment())
	rep := repo.NewRepository(fakeDB, logger) // ctx := context.Background()
	// tx := new(sql.Tx)
	mock.ExpectBegin().WillReturnError(errors.New("error"))
	tx, err := rep.BeginTx(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()))
	assert.Error(t, err)
	assert.Empty(t, tx)
}

func (suite *AdvertRepoTestSuite) TestCreateAdvertTypeHouse() {
	type args struct {
		adv     *models.HouseTypeAdvert
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
			name: "successful create advert type",
			args: args{
				adv: &models.HouseTypeAdvert{
					HouseID:   rand.Int63(),
					AdvertID:  rand.Int63(),
					IsDeleted: false,
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
				adv: &models.HouseTypeAdvert{
					HouseID:   rand.Int63(),
					AdvertID:  rand.Int63(),
					IsDeleted: false,
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
			suite.setupMockCreateAdvertTypeHouse(tt.args.adv, tt.args.errExec, tt.args.expExec)
			logger := zap.Must(zap.NewDevelopment())
			rep := repo.NewRepository(suite.db, logger)
			gotErr := rep.CreateAdvertTypeHouse(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tx, tt.args.adv)
			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *AdvertRepoTestSuite) setupMockCreateAdvertTypeHouse(advType *models.HouseTypeAdvert, errExec error, expExec bool) {
	if expExec {
		query := `INSERT INTO advert_type_house (house_id, advert_id) VALUES ($1, $2)`
		escapedQuery := regexp.QuoteMeta(query)
		suite.mock.ExpectExec(escapedQuery).
			WillReturnError(errExec).WillReturnResult(sqlmock.NewResult(1, 1)).WithArgs(
			advType.HouseID, advType.AdvertID)
	}
}

func (suite *AdvertRepoTestSuite) TestCreateAdvertTypeFlat() {
	type args struct {
		adv     *models.FlatTypeAdvert
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
			name: "successful create advert type",
			args: args{
				adv: &models.FlatTypeAdvert{
					FlatID:    rand.Int63(),
					AdvertID:  rand.Int63(),
					IsDeleted: false,
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
				adv: &models.FlatTypeAdvert{
					FlatID:    rand.Int63(),
					AdvertID:  rand.Int63(),
					IsDeleted: false,
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
			suite.setupMockCreateAdvertTypeFlat(tt.args.adv, tt.args.errExec, tt.args.expExec)
			logger := zap.Must(zap.NewDevelopment())
			rep := repo.NewRepository(suite.db, logger)
			gotErr := rep.CreateAdvertTypeFlat(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tx, tt.args.adv)
			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *AdvertRepoTestSuite) setupMockCreateAdvertTypeFlat(advType *models.FlatTypeAdvert, errExec error, expExec bool) {
	if expExec {
		query := `INSERT INTO advert_type_flat (flat_id, advert_id) VALUES ($1, $2)`
		escapedQuery := regexp.QuoteMeta(query)
		suite.mock.ExpectExec(escapedQuery).
			WillReturnError(errExec).WillReturnResult(sqlmock.NewResult(1, 1)).WithArgs(
			advType.FlatID, advType.AdvertID)
	}
}

func (suite *AdvertRepoTestSuite) TestCreateAdvert() {
	type args struct {
		adv     *models.Advert
		errExec error
		expExec bool
	}
	type want struct {
		newId int64
		err   error
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
					ID:             11,
					UserID:         22,
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
				newId: 11,
				err:   nil,
			},
		},
		{
			name: "fail create advert",
			args: args{
				adv: &models.Advert{
					ID:             11,
					UserID:         22,
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
				newId: 0,
				err:   errors.New("some error"),
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
			suite.setupMockCreateAdvert(tt.args.adv, tt.want.newId, tt.args.errExec, tt.args.expExec)
			logger := zap.Must(zap.NewDevelopment())
			rep := repo.NewRepository(suite.db, logger)
			advId, gotErr := rep.CreateAdvert(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tx, tt.args.adv)
			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().Equal(tt.want.newId, advId)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *AdvertRepoTestSuite) setupMockCreateAdvert(newAdvert *models.Advert, advId int64, errExec error, expExec bool) {
	if expExec {
		query := `INSERT INTO advert (user_id, type_placement, title, description, phone, is_agent, priority) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
		escapedQuery := regexp.QuoteMeta(query)
		// suite.mock.ExpectExec(escapedQuery).
		// 		WillReturnError(errExec).WillReturnResult(sqlmock.NewResult(advId, 1)).WithArgs(
		// 			newAdvert.UserID, newAdvert.AdvertTypeSale, newAdvert.Title, newAdvert.Description,
		// 			newAdvert.Phone, newAdvert.IsAgent, newAdvert.Priority)
			suite.mock.ExpectExec(query).
				WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(nil).
				WithArgs(newAdvert.UserID, newAdvert.AdvertTypeSale, newAdvert.Title,
					newAdvert.Description, newAdvert.Phone, newAdvert.IsAgent, newAdvert.Priority)

		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(newAdvert.ID)

		if expExec {
			suite.mock.ExpectQuery(escapedQuery).
				WillReturnError(errExec).WithArgs(
				newAdvert.UserID, newAdvert.AdvertTypeSale, newAdvert.Title, newAdvert.Description,
				newAdvert.Phone, newAdvert.IsAgent, newAdvert.Priority).WillReturnRows(rows)
		}

	}
}

func (suite *AdvertRepoTestSuite) TestCreateProvince() {
	type args struct {
		adv       string
		errQuery1 error
		errQuery2 error
		expQuery1 bool
		expQuery2 bool
	}
	type want struct {
		newId int64
		err   error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful create province",
			args: args{
				adv:       "PROVINCE",
				errQuery1: nil,
				errQuery2: nil,
				expQuery1: true,
				expQuery2: false,
			},
			want: want{
				newId: 11,
				err:   nil,
			},
		},
		{
			name: "fail create province",
			args: args{
				adv:       "PROVINCE",
				errQuery1: errors.New("some error"),
				errQuery2: errors.New("some error"),
				expQuery1: true,
				expQuery2: true,
			},
			want: want{
				newId: 0,
				err:   errors.New("some error"),
			},
		},
		{
			name: "create province",
			args: args{
				adv:       "PROVINCE",
				errQuery1: errors.New("some error"),
				errQuery2: nil,
				expQuery1: true,
				expQuery2: true,
			},
			want: want{
				newId: 12,
				err:   nil,
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
			suite.setupMockCreateProvince(tt.args.adv, tt.want.newId, tt.args.errQuery1, tt.args.errQuery2,
				tt.args.expQuery1, tt.args.expQuery2)
			logger := zap.Must(zap.NewDevelopment())
			rep := repo.NewRepository(suite.db, logger)
			advId, gotErr := rep.CreateProvince(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tx, tt.args.adv)
			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().Equal(tt.want.newId, advId)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *AdvertRepoTestSuite) setupMockCreateProvince(newProvince string, provId int64, errExec1 error, errExec2 error,
	expExec1 bool, expExec2 bool) {
	if expExec1 {
		query := `SELECT id FROM province WHERE name=$1`
		escapedQuery := regexp.QuoteMeta(query)
		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(provId)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(errExec1).WithArgs(
			newProvince).WillReturnRows(rows)
	}

	if expExec2 {
		insert := `INSERT INTO province (name) VALUES ($1) RETURNING id`
		escapedQuery := regexp.QuoteMeta(insert)
		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(provId)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(errExec2).WithArgs(
			newProvince).WillReturnRows(rows)
	}
}

func (suite *AdvertRepoTestSuite) TestCreateTown() {
	type args struct {
		prId      int64
		name      string
		errQuery1 error
		errQuery2 error
		expQuery1 bool
		expQuery2 bool
	}
	type want struct {
		newId int64
		err   error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful found town",
			args: args{
				prId:      100,
				name:      "Town",
				errQuery1: nil,
				errQuery2: nil,
				expQuery1: true,
				expQuery2: false,
			},
			want: want{
				newId: 11,
				err:   nil,
			},
		},
		{
			name: "fail create town",
			args: args{
				prId:      100,
				name:      "Town",
				errQuery1: errors.New("some error"),
				errQuery2: errors.New("some error"),
				expQuery1: true,
				expQuery2: true,
			},
			want: want{
				newId: 0,
				err:   errors.New("some error"),
			},
		},
		{
			name: "create town",
			args: args{
				prId:      100,
				name:      "Town",
				errQuery1: errors.New("some error"),
				errQuery2: nil,
				expQuery1: true,
				expQuery2: true,
			},
			want: want{
				newId: 100,
				err:   nil,
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
			suite.setupMockCreateTown(tt.args.prId, tt.args.name, tt.want.newId, tt.args.errQuery1, tt.args.errQuery2,
				tt.args.expQuery1, tt.args.expQuery2)
			logger := zap.Must(zap.NewDevelopment())
			rep := repo.NewRepository(suite.db, logger)
			advId, gotErr := rep.CreateTown(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tx, tt.args.prId, tt.args.name)
			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().Equal(tt.want.newId, advId)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *AdvertRepoTestSuite) setupMockCreateTown(prId int64, townName string, townId int64, errExec1 error, errExec2 error,
	expExec1 bool, expExec2 bool) {
	if expExec1 {
		query := `SELECT id FROM town WHERE name=$1 AND province_id=$2`
		escapedQuery := regexp.QuoteMeta(query)
		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(townId)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(errExec1).WithArgs(
			townName, prId).WillReturnRows(rows)
	}

	if expExec2 {
		insert := `INSERT INTO town (name, province_id) VALUES ($1, $2) RETURNING id`
		escapedQuery := regexp.QuoteMeta(insert)
		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(townId)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(errExec2).WithArgs(
			townName, prId).WillReturnRows(rows)
	}
}

func (suite *AdvertRepoTestSuite) TestCreateStreet() {
	type args struct {
		idTown    int64
		name      string
		errQuery1 error
		errQuery2 error
		expQuery1 bool
		expQuery2 bool
	}
	type want struct {
		newId int64
		err   error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful found street",
			args: args{
				idTown:    100,
				name:      "Town",
				errQuery1: nil,
				errQuery2: nil,
				expQuery1: true,
				expQuery2: false,
			},
			want: want{
				newId: 11,
				err:   nil,
			},
		},
		{
			name: "fail create street",
			args: args{
				idTown:    100,
				name:      "Town",
				errQuery1: errors.New("some error"),
				errQuery2: errors.New("some error"),
				expQuery1: true,
				expQuery2: true,
			},
			want: want{
				newId: 0,
				err:   errors.New("some error"),
			},
		},
		{
			name: "create street",
			args: args{
				idTown:    100,
				name:      "Town",
				errQuery1: errors.New("some error"),
				errQuery2: nil,
				expQuery1: true,
				expQuery2: true,
			},
			want: want{
				newId: 100,
				err:   nil,
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
			suite.setupMockCreateStreet(tt.args.idTown, tt.args.name, tt.want.newId, tt.args.errQuery1, tt.args.errQuery2,
				tt.args.expQuery1, tt.args.expQuery2)
			logger := zap.Must(zap.NewDevelopment())
			rep := repo.NewRepository(suite.db, logger)
			strId, gotErr := rep.CreateStreet(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tx, tt.args.idTown, tt.args.name)
			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().Equal(tt.want.newId, strId)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *AdvertRepoTestSuite) setupMockCreateStreet(prId int64, townName string, strId int64, errExec1 error, errExec2 error,
	expExec1 bool, expExec2 bool) {
	if expExec1 {
		query := `SELECT id FROM street WHERE name=$1 AND town_id=$2`
		escapedQuery := regexp.QuoteMeta(query)
		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(strId)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(errExec1).WithArgs(
			townName, prId).WillReturnRows(rows)
	}

	if expExec2 {
		insert := `INSERT INTO street (name, town_id) VALUES ($1, $2) RETURNING id`
		escapedQuery := regexp.QuoteMeta(insert)
		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(strId)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(errExec2).WithArgs(
			townName, prId).WillReturnRows(rows)
	}
}

func (suite *AdvertRepoTestSuite) TestCreateHouseAddress() {
	type args struct {
		strId     int64
		houseadr  string
		errQuery1 error
		errQuery2 error
		expQuery1 bool
		expQuery2 bool
	}
	type want struct {
		newId int64
		err   error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful found house address",
			args: args{
				strId:     100,
				houseadr:  "adr",
				errQuery1: nil,
				errQuery2: nil,
				expQuery1: true,
				expQuery2: false,
			},
			want: want{
				newId: 11,
				err:   nil,
			},
		},
		{
			name: "fail create house address",
			args: args{
				strId:     100,
				houseadr:  "adr",
				errQuery1: errors.New("some error"),
				errQuery2: errors.New("some error"),
				expQuery1: true,
				expQuery2: true,
			},
			want: want{
				newId: 0,
				err:   errors.New("some error"),
			},
		},
		{
			name: "create house address",
			args: args{
				strId:     100,
				houseadr:  "adr",
				errQuery1: errors.New("some error"),
				errQuery2: nil,
				expQuery1: true,
				expQuery2: true,
			},
			want: want{
				newId: 100,
				err:   nil,
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
			suite.setupMockCreateHouseAddress(tt.args.strId, tt.args.houseadr, tt.want.newId, tt.args.errQuery1, tt.args.errQuery2,
				tt.args.expQuery1, tt.args.expQuery2)
			logger := zap.Must(zap.NewDevelopment())
			rep := repo.NewRepository(suite.db, logger)
			strId, gotErr := rep.CreateHouseAddress(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tx, tt.args.strId, tt.args.houseadr)
			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().Equal(tt.want.newId, strId)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *AdvertRepoTestSuite) setupMockCreateHouseAddress(prId int64, townName string, strId int64, errExec1 error, errExec2 error,
	expExec1 bool, expExec2 bool) {
	if expExec1 {
		query := `SELECT id FROM house_name WHERE name=$1 AND street_id=$2`
		escapedQuery := regexp.QuoteMeta(query)
		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(strId)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(errExec1).WithArgs(
			townName, prId).WillReturnRows(rows)
	}

	if expExec2 {
		insert := `INSERT INTO house_name (name, street_id) VALUES ($1, $2) RETURNING id`
		escapedQuery := regexp.QuoteMeta(insert)
		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(strId)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(errExec2).WithArgs(
			townName, prId).WillReturnRows(rows)
	}
}

func (suite *AdvertRepoTestSuite) TestCreateAddress() {
	type args struct {
		hId       int64
		metro     string
		adrPoint  string
		errQuery1 error
		errQuery2 error
		expQuery1 bool
		expQuery2 bool
	}
	type want struct {
		newId int64
		err   error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful found house address",
			args: args{
				hId:       100,
				metro:     "metro",
				adrPoint:  "adr",
				errQuery1: nil,
				errQuery2: nil,
				expQuery1: true,
				expQuery2: false,
			},
			want: want{
				newId: 11,
				err:   nil,
			},
		},
		{
			name: "fail create house address",
			args: args{
				hId:       100,
				metro:     "metro",
				adrPoint:  "adr",
				errQuery1: errors.New("some error"),
				errQuery2: errors.New("some error"),
				expQuery1: true,
				expQuery2: true,
			},
			want: want{
				newId: 0,
				err:   errors.New("some error"),
			},
		},
		{
			name: "create house address",
			args: args{
				hId:       100,
				metro:     "metro",
				adrPoint:  "adr",
				errQuery1: errors.New("some error"),
				errQuery2: nil,
				expQuery1: true,
				expQuery2: true,
			},
			want: want{
				newId: 10,
				err:   nil,
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
			suite.setupMockCreateAddress(tt.args.hId, tt.args.adrPoint, tt.args.metro, tt.want.newId, tt.args.errQuery1, tt.args.errQuery2,
				tt.args.expQuery1, tt.args.expQuery2)
			logger := zap.Must(zap.NewDevelopment())
			rep := repo.NewRepository(suite.db, logger)
			strId, gotErr := rep.CreateAddress(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tx, tt.args.hId, tt.args.metro, tt.args.adrPoint)
			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().Equal(tt.want.newId, strId)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *AdvertRepoTestSuite) setupMockCreateAddress(hId int64, point string, metro string, adrId int64, errExec1 error, errExec2 error,
	expExec1 bool, expExec2 bool) {
	if expExec1 {
		query := `SELECT id FROM address WHERE house_name_id=$1`
		escapedQuery := regexp.QuoteMeta(query)
		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(adrId)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(errExec1).WithArgs(hId).WillReturnRows(rows)
	}

	if expExec2 {
		insert := `INSERT INTO address (metro, house_name_id, address_point) VALUES ($1, $2, $3) RETURNING id`
		escapedQuery := regexp.QuoteMeta(insert)
		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(adrId)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(errExec2).WithArgs(metro,
			hId, point).WillReturnRows(rows)
	}
}

func (suite *AdvertRepoTestSuite) TestCreatePriceChange() {
	type args struct {
		newPrice  models.PriceChange
		errQuery1 error
		expQuery1 bool
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
			name: "successful cr price change",
			args: args{
				newPrice: models.PriceChange{Price: 10, ID: 124,
					AdvertID: 234, IsDeleted: false},
				errQuery1: nil,
				expQuery1: true,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "fail cr price change",
			args: args{
				newPrice: models.PriceChange{Price: 10, ID: 124,
					AdvertID: 234, IsDeleted: false},
				errQuery1: errors.New("some error"),
				expQuery1: true,
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
			suite.setupMockCreatePriceChange(tt.args.newPrice,
				tt.args.errQuery1, tt.args.expQuery1)
			logger := zap.Must(zap.NewDevelopment())
			rep := repo.NewRepository(suite.db, logger)
			gotErr := rep.CreatePriceChange(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()),
				tx, &tt.args.newPrice)
			suite.Assert().Equal(tt.want.err, gotErr)
			// suite.Assert().Equal(tt.want.newId, strId)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *AdvertRepoTestSuite) setupMockCreatePriceChange(newPrice models.PriceChange, errExec1 error,
	expExec1 bool) {
	if expExec1 {
		insert := `INSERT INTO price_change (advert_id, price) VALUES ($1, $2)`
		escapedQuery := regexp.QuoteMeta(insert)
		// rows := sqlmock.NewRows([]string{"id"})
		// rows = rows.AddRow(adrId)

		suite.mock.ExpectExec(escapedQuery).
			WillReturnError(errExec1).WithArgs(newPrice.AdvertID, newPrice.Price).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}
}

func (suite *AdvertRepoTestSuite) TestCreateBuilding() {
	type args struct {
		newPrice  models.Building
		errQuery1 error
		expQuery1 bool
	}
	type want struct {
		wId int64
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful cr price change",
			args: args{
				newPrice: models.Building{ID: 124,
					ComplexID: 234, Floor: 2, IsDeleted: false},
				errQuery1: nil,
				expQuery1: true,
			},
			want: want{
				wId: 124,
				err: nil,
			},
		},
		{
			name: "fail cr price change",
			args: args{
				newPrice: models.Building{ID: 124,
					ComplexID: 234, Floor: 2, IsDeleted: false},
				errQuery1: errors.New("some error"),
				expQuery1: true,
			},
			want: want{
				wId: 0,
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
			suite.setupMockCreateBuilding(tt.args.newPrice,
				tt.args.errQuery1, tt.args.expQuery1)
			logger := zap.Must(zap.NewDevelopment())
			rep := repo.NewRepository(suite.db, logger)
			bId, gotErr := rep.CreateBuilding(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()),
				tx, &tt.args.newPrice)
			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().Equal(tt.want.wId, bId)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *AdvertRepoTestSuite) setupMockCreateBuilding(newPrice models.Building, errExec1 error,
	expExec1 bool) {
	if expExec1 {
		insert := `INSERT INTO building (floor, material_building, address_id, year_creation) VALUES ($1, $2, $3, $4) RETURNING id`
		escapedQuery := regexp.QuoteMeta(insert)
		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(newPrice.ID)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(errExec1).WithArgs(newPrice.Floor, newPrice.Material,
			newPrice.AddressID, newPrice.YearCreation).WillReturnRows(rows)
	}
}

func (suite *AdvertRepoTestSuite) TestCheckExistsBuilding() {
	type args struct {
		adr       models.AddressData
		build     models.Building
		errQuery1 error
		expQuery1 bool
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
			name: "successful check ex",
			args: args{
				adr: models.AddressData{
					Province: "pr",
					Town:     "town",
					Street:   "street",
				},
				build: models.Building{ID: 124,
					ComplexID: 0, Floor: 2, IsDeleted: false,
				},
				errQuery1: nil,
				expQuery1: true,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "fail check ex",
			args: args{
				adr: models.AddressData{
					Province: "pr",
					Town:     "town",
					Street:   "street",
				},
				build: models.Building{ID: 124,
					ComplexID: 0, Floor: 2, IsDeleted: false,
				},
				errQuery1: errors.New("some error"),
				expQuery1: true,
			},
			want: want{
				err: errors.New("some error"),
			},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.mock.ExpectBegin()
			_, err := suite.db.Begin()
			if err != nil {
				suite.T().Fatal("Error beginning transaction:", err)
			}
			suite.setupMockCheckExistsBuilding(tt.args.adr, tt.args.build,
				tt.args.errQuery1, tt.args.expQuery1)
			logger := zap.Must(zap.NewDevelopment())
			rep := repo.NewRepository(suite.db, logger)
			gotBuild, gotErr := rep.CheckExistsBuilding(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()),
				&tt.args.adr)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil {
				suite.Assert().Nil(gotBuild)
			} else {
				suite.Assert().Equal(&tt.args.build, gotBuild)
			}
			// suite.Assert().Equal(tt.want.wId, bId)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *AdvertRepoTestSuite) setupMockCheckExistsBuilding(adr models.AddressData, build models.Building,
	errExec1 error, expExec1 bool) {
	if expExec1 {
		query := `SELECT b.id, b.address_id, b.floor, b.material_building, b.year_creation FROM building AS b JOIN address AS a ON b.address_id=a.id JOIN house_name AS h ON a.house_name_id=h.id JOIN street AS s ON h.street_id=s.id JOIN town AS t ON s.town_id=t.id JOIN province AS p ON t.province_id=p.id WHERE p.name=$1 AND t.name=$2 AND s.name=$3 AND h.name=$4;`
		escapedQuery := regexp.QuoteMeta(query)
		rows := sqlmock.NewRows([]string{"id", "address", "floor", "material", "yearCreation"})
		rows = rows.AddRow(build.ID, build.AddressID, build.Floor, build.Material,
			build.YearCreation)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(errExec1).WithArgs(adr.Province, adr.Town, adr.Street, adr.House).
			WillReturnRows(rows)
	}
}

func (suite *AdvertRepoTestSuite) TestCheckExistsBuildingData() {
	type args struct {
		adr       models.AddressData
		build     models.BuildingData
		errQuery1 error
		expQuery1 bool
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
			name: "successful check ex",
			args: args{
				adr: models.AddressData{
					Province: "pr",
					Town:     "town",
					Street:   "street",
				},
				build: models.BuildingData{ComplexName: "name",
					Material: "material", YearCreation: 2020, Floor: 2,
				},
				errQuery1: nil,
				expQuery1: true,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "fail check ex",
			args: args{
				adr: models.AddressData{
					Province: "pr",
					Town:     "town",
					Street:   "street",
				},
				build: models.BuildingData{ComplexName: "name",
					Material: "material", YearCreation: 2020, Floor: 2,
				},
				errQuery1: errors.New("not found"),
				expQuery1: true,
			},
			want: want{
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.mock.ExpectBegin()
			_, err := suite.db.Begin()
			if err != nil {
				suite.T().Fatal("Error beginning transaction:", err)
			}
			suite.setupMockCheckExistsBuildingData(tt.args.adr, tt.args.build,
				tt.args.errQuery1, tt.args.expQuery1)
			logger := zap.Must(zap.NewDevelopment())
			rep := repo.NewRepository(suite.db, logger)
			gotBuild, gotErr := rep.CheckExistsBuildingData(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()),
				&tt.args.adr)
			suite.Assert().Equal(tt.want.err, gotErr)
			if tt.want.err != nil || tt.args.errQuery1 != nil {
				suite.Assert().Nil(gotBuild)
			} else {
				suite.Assert().Equal(&tt.args.build, gotBuild)
			}
			// suite.Assert().Equal(tt.want.wId, bId)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *AdvertRepoTestSuite) setupMockCheckExistsBuildingData(adr models.AddressData, build models.BuildingData,
	errExec1 error, expExec1 bool) {
	if expExec1 {
		query := `SELECT b.floor, b.material_building, b.year_creation, COALESCE(c.name, '')
		FROM building AS b JOIN address AS a ON b.address_id=a.id JOIN house_name AS h ON a.house_name_id=h.id
		    JOIN street AS s ON h.street_id=s.id JOIN town AS t ON s.town_id=t.id JOIN province AS p ON
		        t.province_id=p.id LEFT JOIN complex AS c ON c.id=b.complex_id WHERE p.name=$1 AND t.name=$2 AND
		                                                                             s.name=$3 AND h.name=$4;`
		escapedQuery := regexp.QuoteMeta(query)
		rows := sqlmock.NewRows([]string{"id", "address", "floor", "material"})
		rows = rows.AddRow(build.Floor, build.Material, build.YearCreation, build.ComplexName)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(errExec1).WithArgs(adr.Province, adr.Town, adr.Street, adr.House).
			WillReturnRows(rows)
	}
}

func (suite *AdvertRepoTestSuite) TestCreateHouse() {
	type args struct {
		house     models.House
		errQuery1 error
		expQuery1 bool
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
			name: "successful check ex",
			args: args{
				house:     models.House{ID: 124, BuildingID: 122, CeilingHeight: 124.214},
				errQuery1: nil,
				expQuery1: true,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "fail check ex",
			args: args{
				house:     models.House{ID: 0, BuildingID: 122, CeilingHeight: 124.214},
				errQuery1: errors.New("some error"),
				expQuery1: true,
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
			suite.setupMockCreateHouse(tt.args.house,
				tt.args.errQuery1, tt.args.expQuery1)
			logger := zap.Must(zap.NewDevelopment())
			rep := repo.NewRepository(suite.db, logger)
			gotHouseId, gotErr := rep.CreateHouse(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()),
				tx, &tt.args.house)
			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().Equal(tt.args.house.ID, gotHouseId)
			// suite.Assert().Equal(tt.want.wId, bId)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *AdvertRepoTestSuite) setupMockCreateHouse(newHouse models.House,
	errExec1 error, expExec1 bool) {
	if expExec1 {
		insert := `INSERT INTO house (building_id, ceiling_height, square_area, square_house, bedroom_count, status_area_house, cottage, status_home_house) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

		escapedQuery := regexp.QuoteMeta(insert)
		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(newHouse.ID)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(errExec1).WithArgs(newHouse.BuildingID, newHouse.CeilingHeight, newHouse.SquareArea,
			newHouse.SquareHouse, newHouse.BedroomCount, newHouse.StatusArea, newHouse.Cottage, newHouse.StatusHome).
			WillReturnRows(rows)
	}
}

func (suite *AdvertRepoTestSuite) TestCreateFlat() {
	type args struct {
		flat      models.Flat
		errQuery1 error
		expQuery1 bool
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
				flat: models.Flat{ID: 124, BuildingID: 122, CeilingHeight: 124.214,
					RoomCount: 2},
				errQuery1: nil,
				expQuery1: true,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "fail create flate",
			args: args{
				flat: models.Flat{ID: 0, BuildingID: 122, CeilingHeight: 124.214,
					RoomCount: 2},
				errQuery1: errors.New("some error"),
				expQuery1: true,
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
			suite.setupMockCreateFlat(tt.args.flat,
				tt.args.errQuery1, tt.args.expQuery1)
			logger := zap.Must(zap.NewDevelopment())
			rep := repo.NewRepository(suite.db, logger)
			gotFlatId, gotErr := rep.CreateFlat(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()),
				tx, &tt.args.flat)
			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().Equal(tt.args.flat.ID, gotFlatId)
			// suite.Assert().Equal(tt.want.wId, bId)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *AdvertRepoTestSuite) setupMockCreateFlat(newFlat models.Flat,
	errExec1 error, expExec1 bool) {
	if expExec1 {
		insert := `INSERT INTO flat (building_id, floor, ceiling_height, square_general, bedroom_count, square_residential, apartament) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

		escapedQuery := regexp.QuoteMeta(insert)
		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(newFlat.ID)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(errExec1).WithArgs(
			newFlat.BuildingID, newFlat.Floor, newFlat.CeilingHeight,
			newFlat.SquareGeneral, newFlat.RoomCount, newFlat.SquareResidential, newFlat.Apartment).
			WillReturnRows(rows)
	}
}

func (suite *AdvertRepoTestSuite) TestSelectImages() {
	type args struct {
		advertId  int64
		errQuery1 error
		expQuery1 bool
	}
	type want struct {
		err  error
		resp []*models.ImageResp
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful select images",
			args: args{
				advertId:  124,
				errQuery1: nil,
				expQuery1: true,
			},
			want: want{
				err: nil,
				resp: []*models.ImageResp{
					{
						ID:       1,
						Photo:    "/path1",
						Priority: 1,
					},
					{
						ID:       2,
						Photo:    "/path2",
						Priority: 2,
					},
				},
			},
		},
		{
			name: "fail select images",
			args: args{
				advertId:  124,
				errQuery1: errors.New("some error"),
				expQuery1: true,
			},
			want: want{
				err:  errors.New("some error"),
				resp: nil,
			},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.mock.ExpectBegin()
			_, err := suite.db.Begin()
			if err != nil {
				suite.T().Fatal("Error beginning transaction:", err)
			}
			suite.setupMockSelectImages(tt.args.advertId, tt.want.resp,
				tt.args.errQuery1, tt.args.expQuery1)
			logger := zap.Must(zap.NewDevelopment())
			rep := repo.NewRepository(suite.db, logger)
			gotImages, gotErr := rep.SelectImages(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()),
				tt.args.advertId)
			suite.Assert().Equal(tt.want.resp, gotImages)
			suite.Assert().Equal(tt.want.err, gotErr)
			// suite.Assert().Equal(tt.want.wId, bId)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *AdvertRepoTestSuite) setupMockSelectImages(advId int64, imresp []*models.ImageResp,
	errExec1 error, expExec1 bool) {
	if expExec1 {
		selectQuery := `SELECT id, photo, priority FROM image WHERE advert_id = $1 AND is_deleted = false`

		escapedQuery := regexp.QuoteMeta(selectQuery)
		rows := sqlmock.NewRows([]string{"id", "photo", "priority"})
		if imresp != nil {

			rows = rows.AddRow(imresp[0].ID, imresp[0].Photo, imresp[0].Priority)
			rows = rows.AddRow(imresp[1].ID, imresp[1].Photo, imresp[1].Priority)

		}
		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(errExec1).WithArgs(advId).
			WillReturnRows(rows)
	}
}

func (suite *AdvertRepoTestSuite) TestSelectPriceChanges() {
	type args struct {
		advertId  int64
		errQuery1 error
		expQuery1 bool
	}
	type want struct {
		err  error
		resp []*models.PriceChangeData
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful select price change",
			args: args{
				advertId:  124,
				errQuery1: nil,
				expQuery1: true,
			},
			want: want{
				err: nil,
				resp: []*models.PriceChangeData{
					{
						Price:        1224,
						DateCreation: time.Now(),
					},
					{
						Price:        12224,
						DateCreation: time.Now(),
					},
				},
			},
		},
		{
			name: "fail select price change",
			args: args{
				advertId:  124,
				errQuery1: errors.New("some error"),
				expQuery1: true,
			},
			want: want{
				err:  errors.New("some error"),
				resp: nil,
			},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.mock.ExpectBegin()
			_, err := suite.db.Begin()
			if err != nil {
				suite.T().Fatal("Error beginning transaction:", err)
			}
			suite.setupMockSelectPriceChanges(tt.args.advertId, tt.want.resp,
				tt.args.errQuery1, tt.args.expQuery1)
			logger := zap.Must(zap.NewDevelopment())
			rep := repo.NewRepository(suite.db, logger)
			gotImages, gotErr := rep.SelectPriceChanges(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()),
				tt.args.advertId)
			suite.Assert().Equal(tt.want.resp, gotImages)
			suite.Assert().Equal(tt.want.err, gotErr)
			// suite.Assert().Equal(tt.want.wId, bId)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *AdvertRepoTestSuite) setupMockSelectPriceChanges(advId int64, imresp []*models.PriceChangeData,
	errExec1 error, expExec1 bool) {
	if expExec1 {
		selectQuery := `SELECT price, created_at FROM price_change WHERE advert_id = $1 AND is_deleted = false`

		escapedQuery := regexp.QuoteMeta(selectQuery)
		rows := sqlmock.NewRows([]string{"price", "created_at"})
		if imresp != nil {
			rows = rows.AddRow(imresp[0].Price, imresp[0].DateCreation)
			rows = rows.AddRow(imresp[1].Price, imresp[1].DateCreation)
		}
		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(errExec1).WithArgs(advId).
			WillReturnRows(rows)
	}
}

func (suite *AdvertRepoTestSuite) TestGetTypeAdvertById() {
	type args struct {
		advertId  int64
		errQuery1 error
		expQuery1 bool
	}
	type want struct {
		err  error
		resp models.AdvertTypeAdvert
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful get type advert",
			args: args{
				advertId:  124,
				errQuery1: nil,
				expQuery1: true,
			},
			want: want{
				err:  nil,
				resp: models.AdvertTypeHouse,
			},
		},
		{
			name: "fail get type advert",
			args: args{
				advertId:  124,
				errQuery1: errors.New("some error"),
				expQuery1: true,
			},
			want: want{
				err:  errors.New("some error"),
				resp: "",
			},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.mock.ExpectBegin()
			_, err := suite.db.Begin()
			if err != nil {
				suite.T().Fatal("Error beginning transaction:", err)
			}
			suite.setupMockGetTypeAdvertById(tt.args.advertId, tt.want.resp,
				tt.args.errQuery1, tt.args.expQuery1)
			logger := zap.Must(zap.NewDevelopment())
			rep := repo.NewRepository(suite.db, logger)
			gotType, gotErr := rep.GetTypeAdvertById(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()),
				tt.args.advertId)

			if tt.want.resp == "" {
				suite.Assert().Nil(gotType)
			} else {
				suite.Assert().Equal(&tt.want.resp, gotType)
			}
			suite.Assert().Equal(tt.want.err, gotErr)
			// suite.Assert().Equal(tt.want.wId, bId)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *AdvertRepoTestSuite) setupMockGetTypeAdvertById(advId int64, advType models.AdvertTypeAdvert,
	errExec1 error, expExec1 bool) {
	if expExec1 {
		query := `SELECT                   CASE
	WHEN ath.house_id IS NOT NULL THEN 'House'
	WHEN atf.flat_id IS NOT NULL THEN 'Flat'
	ELSE 'None'
END AS type_advert FROM advert AS a LEFT JOIN advert_type_house AS ath ON a.id=ath.advert_id LEFT JOIN advert_type_flat AS atf ON a.id=atf.advert_id WHERE a.id=$1`

		escapedQuery := regexp.QuoteMeta(query)
		rows := sqlmock.NewRows([]string{"adv_type"})
		rows = rows.AddRow(advType)
		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(errExec1).WithArgs(advId).
			WillReturnRows(rows)
	}
}

func (suite *AdvertRepoTestSuite) TestChangeTypeAdvert() {
	type args struct {
		typeAdvert models.AdvertTypeAdvert
		expBool    []bool  // 15
		expError   []error // 15

		expBoolCheck  []bool  // 2
		expErrorCheck []error // 2
		buildId       int64
		advertId      int64
		houseId       int64
		flatId        int64
	}
	type want struct {
		err error
	}
	errTest := errors.New("some error")
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "fail ChangeTypeAdvert read advertType",
			args: args{
				typeAdvert:    models.AdvertTypeFlat,
				advertId:      124,
				houseId:       122,
				buildId:       555,
				flatId:        1221,
				expBool:       []bool{true, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
				expError:      []error{errTest, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
				expBoolCheck:  []bool{false, false},
				expErrorCheck: []error{nil, nil},
			},
			want: want{
				err: errTest,
			},
		},
		{
			name: "fail ChangeTypeAdvert querySelectBuildingIdByFlat",
			args: args{
				typeAdvert:    models.AdvertTypeFlat,
				advertId:      124,
				houseId:       122,
				buildId:       555,
				flatId:        1221,
				expBool:       []bool{true, true, false, false, false, false, false, false, false, false, false, false, false, false, false},
				expError:      []error{nil, errTest, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
				expBoolCheck:  []bool{false, false},
				expErrorCheck: []error{nil, nil},
			},
			want: want{
				err: errTest,
			},
		},

		{
			name: "fail ChangeTypeAdvert queryDeleteFlatById",
			args: args{
				typeAdvert:    models.AdvertTypeFlat,
				advertId:      124,
				houseId:       122,
				buildId:       555,
				flatId:        1221,
				expBool:       []bool{true, true, false, false, false, false, false, false, false, true, false, false, false, false, false},
				expError:      []error{nil, nil, nil, nil, nil, nil, nil, nil, nil, errTest, nil, nil, nil, nil, nil},
				expBoolCheck:  []bool{false, false},
				expErrorCheck: []error{nil, nil},
			},
			want: want{
				err: errTest,
			},
		},
		{
			name: "fail ChangeTypeAdvert queryDeleteAdvertTypeFlat",
			args: args{
				typeAdvert: models.AdvertTypeFlat,
				advertId:   124,
				houseId:    122,
				buildId:    555,
				flatId:     1221,
				expBool: []bool{true, true, false, false, false,
					false, false, false, false, true,
					false, true, false, false, false},

				expError: []error{nil, nil, nil, nil, nil,
					nil, nil, nil, nil, nil,
					nil, errTest, nil, nil, nil},

				expBoolCheck:  []bool{false, false},
				expErrorCheck: []error{nil, nil},
			},
			want: want{
				err: errTest,
			},
		},

		{
			name: "fail ChangeTypeAdvert queryInsertHouse",
			args: args{
				typeAdvert: models.AdvertTypeFlat,
				advertId:   124,
				houseId:    122,
				buildId:    555,
				flatId:     1221,
				// query querySelectBuildingIdByFlat querySelectBuildingIdByHouse queryInsertFlat queryInsertHouse
				// queryInsertTypeFlat queryInsertTypeHouse queryRestoreFlatById queryRestoreHouseById queryDeleteFlatById
				// queryDeleteHouseById queryDeleteAdvertTypeFlat queryDeleteAdvertTypeHouse queryRestoreAdvertTypeFlat queryRestoreAdvertTypeHouse
				expBool: []bool{true, true, false, false, true,
					false, false, false, false, true,
					false, true, false, false, false},

				expError: []error{nil, nil, nil, nil, errTest,
					nil, nil, nil, nil, nil,
					nil, nil, nil, nil, nil},

				expBoolCheck:  []bool{true, false},
				expErrorCheck: []error{errTest, nil},
			},
			want: want{
				err: errTest,
			},
		},

		{
			name: "fail ChangeTypeAdvert queryInsertTypeHouse",
			args: args{
				typeAdvert: models.AdvertTypeFlat,
				advertId:   124,
				houseId:    122,
				buildId:    555,
				flatId:     1221,
				// query querySelectBuildingIdByFlat querySelectBuildingIdByHouse queryInsertFlat queryInsertHouse
				// queryInsertTypeFlat queryInsertTypeHouse queryRestoreFlatById queryRestoreHouseById queryDeleteFlatById
				// queryDeleteHouseById queryDeleteAdvertTypeFlat queryDeleteAdvertTypeHouse queryRestoreAdvertTypeFlat queryRestoreAdvertTypeHouse
				expBool: []bool{true, true, false, false, true,
					false, true, false, false, true,
					false, true, false, false, false},

				expError: []error{nil, nil, nil, nil, nil,
					nil, errTest, nil, nil, nil,
					nil, nil, nil, nil, nil},

				expBoolCheck:  []bool{true, false},
				expErrorCheck: []error{errTest, nil},
			},
			want: want{
				err: errTest,
			},
		},
			{
				name: "fail ChangeTypeAdvert queryRestoreHouseById",
				args: args{
					typeAdvert: models.AdvertTypeFlat,
					advertId:   124,
					houseId:    122,
					buildId:    555,
					flatId:     1221,
					// query querySelectBuildingIdByFlat querySelectBuildingIdByHouse queryInsertFlat queryInsertHouse
					// queryInsertTypeFlat queryInsertTypeHouse queryRestoreFlatById queryRestoreHouseById queryDeleteFlatById
					// queryDeleteHouseById queryDeleteAdvertTypeFlat queryDeleteAdvertTypeHouse queryRestoreAdvertTypeFlat queryRestoreAdvertTypeHouse
					expBool: []bool{true, true, false, false, false,
						false, false, false, true, true,
						false, true, false, false, false},

					expError: []error{nil, nil, nil, nil, nil,
						nil, nil, nil, errTest, nil,
						nil, nil, nil, nil, nil},

					expBoolCheck:  []bool{true, false},
					expErrorCheck: []error{nil, nil},
				},
				want: want{
					err: errTest,
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
			suite.setupMockChangeTypeAdvert(tt.args.typeAdvert, tt.args.advertId, tt.args.buildId, tt.args.flatId, tt.args.houseId,
				tt.args.expBool, tt.args.expError, tt.args.expBoolCheck, tt.args.expErrorCheck)
			logger := zap.Must(zap.NewDevelopment())
			rep := repo.NewRepository(suite.db, logger)
			gotErr := rep.ChangeTypeAdvert(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()),
				tx, tt.args.advertId)

			suite.Assert().Equal(tt.want.err, gotErr)
			// suite.Assert().Equal(tt.want.wId, bId)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *AdvertRepoTestSuite) setupMockChangeTypeAdvert(advertType models.AdvertTypeAdvert,
	advertId, buildingId, flatId, houseId int64, whatExp []bool,
	expError []error, whatExpCheck []bool, expErrorCheck []error) {
	if whatExp[0] {
		query := `SELECT 			CASE
	WHEN ath.house_id IS NOT NULL THEN 'House'
	WHEN atf.flat_id IS NOT NULL THEN 'Flat'
	ELSE 'None'
END AS type_advert FROM advert AS a LEFT JOIN advert_type_flat AS atf ON a.id=atf.advert_id LEFT JOIN advert_type_house AS ath ON a.id=ath.advert_id WHERE a.id = $1;`

		escapedQuery := regexp.QuoteMeta(query)
		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(advertType)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(expError[0]).WithArgs(sqlmock.AnyArg()).
			WillReturnRows(rows)
	}

	if whatExp[1] {
		querySelectBuildingIdByFlat := `SELECT b.id AS buildingid, f.id AS flatid  FROM advert AS a JOIN advert_type_flat AS at ON at.advert_id=a.id JOIN flat AS f ON f.id=at.flat_id JOIN building AS b ON f.building_id=b.id WHERE a.id=$1`

		escapedQuery := regexp.QuoteMeta(querySelectBuildingIdByFlat)
		rows := sqlmock.NewRows([]string{"bid", "fid"})
		rows = rows.AddRow(buildingId, flatId)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(expError[1]).WithArgs(sqlmock.AnyArg()).
			WillReturnRows(rows)
	}

	if whatExp[9] {
		queryDeleteFlatById := `UPDATE flat SET is_deleted=true WHERE id=$1;`

		escapedQuery := regexp.QuoteMeta(queryDeleteFlatById)
		// rows := sqlmock.NewRows([]string{"id"})
		// rows = rows.AddRow(houseId)

		suite.mock.ExpectExec(escapedQuery).
			WillReturnError(expError[9]).WithArgs(sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	if whatExp[11] {
		queryDeleteAdvertTypeFlat := `UPDATE advert_type_flat SET is_deleted=true WHERE advert_id=$1 AND flat_id=$2;`

		escapedQuery := regexp.QuoteMeta(queryDeleteAdvertTypeFlat)
		// rows := sqlmock.NewRows([]string{"id"})
		// rows = rows.AddRow(houseId)

		suite.mock.ExpectExec(escapedQuery).
			WillReturnError(expError[11]).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	if whatExpCheck[0] {
		query := `SELECT h.id FROM advert AS a JOIN advert_type_house AS at ON a.id=at.advert_id JOIN house AS h ON h.id=at.house_id WHERE a.id = $1;`

		escapedQuery := regexp.QuoteMeta(query)
		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(houseId)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(expErrorCheck[0]).WithArgs(sqlmock.AnyArg()).
			WillReturnRows(rows)
	}

	if whatExp[4] {
		queryInsertHouse := `INSERT INTO house (building_id, ceiling_height, square_area, square_house, bedroom_count, status_area_house, cottage, status_home_house)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
		escapedQuery := regexp.QuoteMeta(queryInsertHouse)
		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(houseId)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(expError[4]).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
			sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(rows)
	}

	if whatExp[6] {
		queryInsertTypeHouse := `INSERT INTO advert_type_house (advert_id, house_id) VALUES ($1, $2);`

		escapedQuery := regexp.QuoteMeta(queryInsertTypeHouse)
		// rows := sqlmock.NewRows([]string{"id"})
		// rows = rows.AddRow(houseId)

		suite.mock.ExpectExec(escapedQuery).
			WillReturnError(expError[6]).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	if whatExp[10] {
		queryDeleteHouseById := `UPDATE house SET is_deleted=true WHERE id=$1;`

		escapedQuery := regexp.QuoteMeta(queryDeleteHouseById)
		// rows := sqlmock.NewRows([]string{"id"})
		// rows = rows.AddRow(houseId)

		suite.mock.ExpectExec(escapedQuery).
			WillReturnError(expError[10]).WithArgs(sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	if whatExp[14] {
		queryRestoreAdvertTypeHouse := `UPDATE advert_type_house SET is_deleted=false WHERE advert_id=$1 AND house_id=$2;`

		escapedQuery := regexp.QuoteMeta(queryRestoreAdvertTypeHouse)
		// rows := sqlmock.NewRows([]string{"id"})
		// rows = rows.AddRow(houseId)

		suite.mock.ExpectExec(escapedQuery).
			WillReturnError(expError[14]).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	if whatExp[2] {
		querySelectBuildingIdByHouse := `SELECT b.id AS buildingid, h.id AS houseId  FROM advert AS a JOIN advert_type_house AS at ON at.advert_id=a.id JOIN house AS h ON h.id=at.house_id JOIN building AS b ON h.building_id=b.id WHERE a.id=$1`

		escapedQuery := regexp.QuoteMeta(querySelectBuildingIdByHouse)
		rows := sqlmock.NewRows([]string{"bid", "hid"})
		rows = rows.AddRow(buildingId, houseId)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(expError[2]).WithArgs(sqlmock.AnyArg()).
			WillReturnRows(rows)
	}
	if whatExp[8] {
		escapedQuery := `UPDATE house SET is_deleted=false WHERE id=$1;`

		escapedQuery = regexp.QuoteMeta(escapedQuery)
		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(houseId)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(expError[8]).WithArgs(sqlmock.AnyArg()).
			WillReturnRows(rows)
	}

	if whatExp[12] {
		queryDeleteAdvertTypeHouse := `UPDATE advert_type_house SET is_deleted=true WHERE advert_id=$1 AND house_id=$2;`

		escapedQuery := regexp.QuoteMeta(queryDeleteAdvertTypeHouse)
		// rows := sqlmock.NewRows([]string{"id"})
		// rows = rows.AddRow(houseId)

		suite.mock.ExpectExec(escapedQuery).
			WillReturnError(expError[12]).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	if whatExpCheck[1] {
		query := `SELECT f.id FROM advert AS a JOIN advert_type_flat AS at ON a.id=at.advert_id JOIN flat AS f ON f.id=at.flat_id WHERE a.id = $1`

		escapedQuery := regexp.QuoteMeta(query)
		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(flatId)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(expErrorCheck[1]).WithArgs(sqlmock.AnyArg()).
			WillReturnRows(rows)
	}

	if whatExp[3] {
		queryInsertFlat := `INSERT INTO flat (building_id, floor, ceiling_height, square_general, bedroom_count, square_residential, apartament)
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

		escapedQuery := regexp.QuoteMeta(queryInsertFlat)
		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(flatId)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(expError[3]).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
			sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(rows)
	}

	if whatExp[5] {
		queryInsertTypeFlat := `INSERT INTO advert_type_flat (advert_id, flat_id) VALUES ($1, $2);`

		escapedQuery := regexp.QuoteMeta(queryInsertTypeFlat)
		// rows := sqlmock.NewRows([]string{"id"})
		// rows = rows.AddRow(houseId)

		suite.mock.ExpectExec(escapedQuery).
			WillReturnError(expError[5]).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	if whatExp[7] {
		queryRestoreFlatById := `UPDATE flat SET is_deleted=false WHERE id=$1;`

		escapedQuery := regexp.QuoteMeta(queryRestoreFlatById)
		rows := sqlmock.NewRows([]string{"id"})
		rows = rows.AddRow(flatId)

		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(expError[7]).WithArgs(sqlmock.AnyArg()).
			WillReturnRows(rows)
	}

	if whatExp[13] {
		queryRestoreAdvertTypeFlat := `UPDATE advert_type_flat SET is_deleted=false WHERE advert_id=$1 AND flat_id=$2;`

		escapedQuery := regexp.QuoteMeta(queryRestoreAdvertTypeFlat)
		// rows := sqlmock.NewRows([]string{"id"})
		// rows = rows.AddRow(houseId)

		suite.mock.ExpectExec(escapedQuery).
			WillReturnError(expError[13]).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

}

// 		prepare   func(f *fields, a *args, w *want) *httptest.ResponseRecorder
type args struct {
		typeAdvert models.AdvertTypeAdvert
		expBool    []bool  // 15
		expError   []error // 15

		expBoolCheck  []bool  // 2
		expErrorCheck []error // 2
		buildId       int64
		advertId      int64
		houseId       int64
		flatId        int64
	}
	type want struct {
		err error
	}
	errTest := errors.New("some error")
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "fail ChangeTypeAdvert read advertType",
			args: args{
				typeAdvert:    models.AdvertTypeFlat,
				advertId:      124,
				houseId:       122,
				buildId:       555,
				flatId:        1221,
				expBool:       []bool{true, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
				expError:      []error{errTest, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
				expBoolCheck:  []bool{false, false},
				expErrorCheck: []error{nil, nil},
			},
			want: want{
				err: errTest,
			},
		},

func (suite *AdvertRepoTestSuite) TestGetHouseAdvertById() {
	type args struct {
		advertId  int64
		errQuery1 error
		expQuery1 bool
	}
	type want struct {
		err  error
		resp *models.AdvertData
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful get type advert",
			args: args{
				advertId:  124,
				errQuery1: nil,
				expQuery1: true,
			},
			want: want{
				err: nil,
				resp: &models.AdvertData{
					ID:         123,
					AdvertType: "house",
					Title:      "title",
				},
			},
		},
		{
			name: "fail get type advert",
			args: args{
				advertId:  124,
				errQuery1: errors.New("some error"),
				expQuery1: true,
			},
			want: want{
				err:  errors.New("some error"),
				resp: nil,
			},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.mock.ExpectBegin()
			_, err := suite.db.Begin()
			if err != nil {
				suite.T().Fatal("Error beginning transaction:", err)
			}
			suite.setupMockGetHouseAdvertById(tt.args.advertId, *tt.want.resp,
				tt.args.errQuery1, tt.args.expQuery1)
			logger := zap.Must(zap.NewDevelopment())
			rep := repo.NewRepository(suite.db, logger)
			ctx := context.WithValue(context.Background(), "requestId", uuid.NewV4().String())
			ctx = context.WithValue(ctx, middleware.CookieName, 11112)
			gotAdv, gotErr := rep.GetHouseAdvertById(ctx, tt.args.advertId)

			if tt.want.resp == nil {
				suite.Assert().Nil(gotAdv)
			} else {
				suite.Assert().Equal(tt.want.resp, gotAdv)
			}

			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

func (suite *AdvertRepoTestSuite) setupMockGetHouseAdvertById(advId int64, advertData models.AdvertData,
	errExec1 error, expExec1 bool) {
	if expExec1 {
		query := `
	SELECT
        a.id,
        a.type_placement,
        a.title,
        a.description,
        pc.price,
        a.phone,
        a.is_agent,
		ad.metro,
		hn.name,
		s.name,
		t.name,
		p.name,
		ST_AsText(ad.address_point::geometry),
        h.ceiling_height,
        h.square_area,
        h.square_house,
        h.bedroom_count,
        h.status_area_house,
        h.cottage,
        h.status_home_house,
        b.floor,
        b.year_creation,
        COALESCE(b.material_building, 'Brick') as material,
        a.created_at,
		CASE
			WHEN fa.advert_id IS NOT NULL AND fa.is_deleted=false THEN true
			ELSE false
		END AS is_liked,
        cx.id AS complexid,
        c.photo AS companyphoto,
        c.name AS companyname,
        cx.name AS complexname
    FROM
        advert AS a
    JOIN
        advert_type_house AS at ON a.id = at.advert_id
    JOIN
        house AS h ON h.id = at.house_id
    JOIN
        building AS b ON h.building_id = b.id
		JOIN address AS ad ON b.address_id=ad.id
		JOIN house_name AS hn ON hn.id=ad.house_name_id
		JOIN street AS s ON s.id=hn.street_id
		JOIN town AS t ON t.id=s.town_id
		JOIN province AS p ON p.id=t.province_id
	LEFT JOIN
		favourite_advert AS fa ON fa.advert_id=a.id AND fa.user_id=$2
    LEFT JOIN
        complex AS cx ON b.complex_id = cx.id
    LEFT JOIN
        company AS c ON cx.company_id = c.id
    JOIN
        LATERAL (
            SELECT *
            FROM price_change AS pc
            WHERE pc.advert_id = a.id
            ORDER BY pc.created_at DESC
            LIMIT 1
        ) AS pc ON TRUE
    WHERE
        a.id = $1 AND a.is_deleted = FALSE;`

		escapedQuery := regexp.QuoteMeta(query)
		rows := sqlmock.NewRows([]string{"adv_type"})
		rows = rows.AddRow(advType)
		suite.mock.ExpectQuery(escapedQuery).
			WillReturnError(errExec1).WithArgs(advId, 11112).
			WillReturnRows(rows)
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
				logger := zap.Must(zap.NewDevelopment())
				rep := repo.NewRepository(suite.db, logger)
				gotErr := rep.CreatePriceChange(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tx, tt.args.adv)
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
				logger := zap.Must(zap.NewDevelopment())
				rep := repo.NewRepository(suite.db, logger)
				gotErr := rep.CreateHouse(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tx, tt.args.adv)
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
				logger := zap.Must(zap.NewDevelopment())
				rep := repo.NewRepository(suite.db, logger)
				gotErr := rep.CreateFlat(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tx, tt.args.adv)
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
						Floor:        5,
						Material:     models.MaterialStalinsky,
						Address:      "123 Main Street",
						AddressPoint: "40.7128° N, 74.0060° W",
						YearCreation: 2000,
						// DateCreation: time.Now(),
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
						Floor:        5,
						Material:     models.MaterialStalinsky,
						Address:      "123 Main Street",
						AddressPoint: "40.7128° N, 74.0060° W",
						YearCreation: 2000,
						// DateCreation: time.Now(),
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
				logger := zap.Must(zap.NewDevelopment())
				rep := repo.NewRepository(suite.db, logger)
				gotErr := rep.CreateBuilding(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tx, tt.args.building)
				suite.Assert().Equal(tt.want.err, gotErr)
				suite.Assert().NoError(suite.mock.ExpectationsWereMet())
			})
		}
	}

	func (suite *UserRepoTestSuite) setupMockCreateBuilding(newBuilding *models.Building, errExec, errQuery error, expExec, epxQuery bool) {
		rows := sqlmock.NewRows([]string{"id", "floor", "material", "adress", "adressPoint", "yearCreation"})
		rows = rows.AddRow(newBuilding.ID, newBuilding.Floor, newBuilding.Material, newBuilding.Address, newBuilding.AddressPoint, newBuilding.YearCreation)
		if expExec {
			suite.mock.ExpectExec(`INSERT INTO buildings \(id, floor, material, adress, adresspoint, yearcreation\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6\)`).
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
				logger := zap.Must(zap.NewDevelopment())
				rep := repo.NewRepository(suite.db, logger)
				gotBuild, gotErr := rep.CheckExistsBuilding(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tt.want.build.Address)
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
			build *models.BuildingData
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
					build: &models.BuildingData{
						ID: uuid.NewV4(),
						// ComplexID:    uuid.NewV4(),
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
					build: &models.BuildingData{
						ID: uuid.NewV4(),
						// ComplexID:    uuid.NewV4(),
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
				logger := zap.Must(zap.NewDevelopment())
				rep := repo.NewRepository(suite.db, logger)
				gotBuild, gotErr := rep.CheckExistsBuildings(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tt.args.pageS, tt.want.build.Address)
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

	func (suite *UserRepoTestSuite) setupMockCheckExistsBuilding2(building *models.BuildingData, pageSize int, errQuery error, epxQuery bool) {
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
				logger := zap.Must(zap.NewDevelopment())
				rep := repo.NewRepository(suite.db, logger)
				gotImages, gotErr := rep.SelectImages(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tt.args.advertID)
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
				logger := zap.Must(zap.NewDevelopment())
				rep := repo.NewRepository(suite.db, logger)
				gotType, gotErr := rep.GetTypeAdvertById(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tt.args.id)
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
						AdvertType:   "House",
						TypeSale:     "Sale",
						Title:        "Beautiful House for Sale",
						Description:  "Spacious house with a large garden",
						Price:        100000,
						Phone:        "123-456-7890",
						IsAgent:      true,
						Address:      "123 Main St, Cityville",
						AddressPoint: "Coordinates",
						// Images:       []*models.ImageResp{},
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
						// YearCreation: time.Now().Year(),
						Material: "Brick",
						// DateCreation: time.Now(),
					},
					err: nil,
				},
			},
		}
		for _, tt := range tests {
			suite.Run(tt.name, func() {
				suite.setupMockGetHouseAdvertById(tt.want.advertData, tt.want.advertData.ID, tt.args.errQuery, tt.args.expQuery)
				logger := zap.Must(zap.NewDevelopment())
				rep := repo.NewRepository(suite.db, logger)
				gotAdvertData, gotErr := rep.GetHouseAdvertById(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tt.want.advertData.ID)
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
			advertData.AdvertType,
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
			WillReturnRows(rows).WillReturnError(nil)
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
				logger := zap.Must(zap.NewDevelopment())
				rep := repo.NewRepository(suite.db, logger)
				gotFlat, gotErr := rep.CheckExistsFlat(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tt.want.flat.ID)
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
				logger := zap.Must(zap.NewDevelopment())
				rep := repo.NewRepository(suite.db, logger)
				gotHouse, gotErr := rep.CheckExistsHouse(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tt.args.advId)
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

	func (suite *UserRepoTestSuite) TestUserRepo_DeleteFlatAdvertById() {
		advertId := uuid.NewV4()
		advertTypeId := uuid.NewV4()
		flatId := uuid.NewV4()
		ctx := context.WithValue(context.Background(), "requestId", uuid.NewV4().String())

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

		logger := zap.Must(zap.NewDevelopment())
		rep := repo.NewRepository(suite.db, logger)
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
		ctx := context.WithValue(context.Background(), "requestId", uuid.NewV4().String())

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

		logger := zap.Must(zap.NewDevelopment())
		rep := repo.NewRepository(suite.db, logger)
		err = rep.DeleteHouseAdvertById(ctx, tx, advertId)
		suite.Assert().NoError(err)

		suite.mock.ExpectCommit()
		err = tx.Commit()
		suite.Assert().NoError(err)

		err = suite.mock.ExpectationsWereMet()
		suite.Assert().NoError(err)
	}

	func (suite *UserRepoTestSuite) TestAdvertRepo_ChangeTypeAdvert() {
		advertId := uuid.NewV4()
		advertTypeId := uuid.NewV4()
		houseId := uuid.NewV4()
		buildingId := uuid.NewV4()
		ctx := context.WithValue(context.Background(), "requestId", uuid.NewV4().String())
		advertType := models.AdvertTypeHouse
		suite.mock.ExpectBegin()
		tx, err := suite.db.Begin()
		suite.NoError(err)

		query := regexp.QuoteMeta(`SELECT at.id, at.adverttype FROM adverts AS a JOIN adverttypes AS at ON a.adverttypeid=at.id WHERE a.id = $1;`)
		// querySelectBuildingIdByFlat := regexp.QuoteMeta(`SELECT b.id AS buildingid, f.id AS flatid  FROM adverts AS a JOIN adverttypes AS at ON at.id=a.adverttypeid JOIN flats AS f ON f.adverttypeid=at.id JOIN buildings AS b ON f.buildingid=b.id WHERE a.id=$1`)
		querySelectBuildingIdByHouse := regexp.QuoteMeta(`SELECT b.id AS buildingid, h.id AS houseid  FROM adverts AS a JOIN adverttypes AS at ON at.id=a.adverttypeid JOIN houses AS h ON h.adverttypeid=at.id JOIN buildings AS b ON h.buildingid=b.id WHERE a.id=$1`)
		// queryInsertFlat := regexp.QuoteMeta(`INSERT INTO flats (id, buildingId, advertTypeId, floor, ceilingHeight, squareGeneral, roomCount, squareResidential, apartament)
		// VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`)
		// queryInsertHouse := regexp.QuoteMeta(`INSERT INTO houses (id, buildingId, advertTypeId, ceilingHeight, squareArea, squareHouse, bedroomCount, statusArea, cottage, statusHome)
		// VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`)
		// queryRestoreFlatById := regexp.QuoteMeta(`UPDATE flats SET isdeleted=false WHERE id=$1;`)
		// queryRestoreHouseById := regexp.QuoteMeta(`UPDATE houses SET isdeleted=false WHERE id=$1;`)
		// queryDeleteFlatById := regexp.QuoteMeta(`UPDATE flats SET isdeleted=true WHERE id=$1;`)
		queryDeleteHouseById := regexp.QuoteMeta(`UPDATE houses SET isdeleted=true WHERE id=$1;`)

		suite.mock.ExpectQuery(query).WithArgs(advertId).
			WillReturnRows(sqlmock.NewRows([]string{"adverttypeid", "advType"}).AddRow(advertTypeId, advertType))

		suite.mock.ExpectQuery(querySelectBuildingIdByHouse).WithArgs(advertId).
			WillReturnRows(sqlmock.NewRows([]string{"adverttypeid", "advType"}).AddRow(buildingId, houseId))

		suite.mock.ExpectExec(queryDeleteHouseById).WithArgs(houseId).WillReturnError(errors.New("error")).
			WillReturnResult(sqlmock.NewResult(1, 1))

		logger := zap.Must(zap.NewDevelopment())
		rep := repo.NewRepository(suite.db, logger)
		err = rep.ChangeTypeAdvert(ctx, tx, advertId)
		suite.Assert().Equal(errors.New("error"), err)

		// err = tx.Commit()

		// suite.Assert().NoError(err)

		err = suite.mock.ExpectationsWereMet()
		suite.Assert().NoError(err)
		suite.db.Close()
	}

	func (suite *UserRepoTestSuite) TestAdvertRepo_UpdateHouseAdvertById() {
		// advertId := uuid.NewV4()
		advertTypeId := uuid.NewV4()
		houseId := uuid.NewV4()
		buildingId := uuid.NewV4()
		ctx := context.WithValue(context.Background(), "requestId", uuid.NewV4().String())
		// advertType := models.AdvertTypeHouse
		suite.mock.ExpectBegin()
		tx, err := suite.db.Begin()
		suite.NoError(err)
		advertUpdateData := &models.AdvertUpdateData{
			ID:              uuid.NewV4(), // Генерируем новый UUID
			TypeAdvert:      "Flat",
			TypeSale:        "Sale",
			Title:           "Beautiful Apartment for Sale",
			Description:     "Spacious apartment with great view",
			Price:           100000,
			Phone:           "+1234567890",
			IsAgent:         true,
			Address:         "123 Main Street",
			AddressPoint:    "Latitude: 40.7128, Longitude: -74.0060",
			HouseProperties: &models.HouseProperties{
				// Заполняем данные для HouseProperties
				// Например: BedroomCount, BathroomCount, SquareHouse и т.д.
			},
			FlatProperties: &models.FlatProperties{
				// Заполняем данные для FlatProperties
				// Например: Floor, SquareGeneral, RoomCount и т.д.
			},
			YearCreation: 2000,
			Material:     "Brick",
		}
		queryGetIdTables := regexp.QuoteMeta(`
	        SELECT
	            at.id as adverttypeid,
	            b.id as buildingid,
	            h.id as houseid,
	            pc.price
	        FROM
	            adverts AS a
	        JOIN
	            adverttypes AS at ON a.adverttypeid = at.id
	        JOIN
	            houses AS h ON h.adverttypeid = at.id
	        JOIN
	            buildings AS b ON h.buildingid = b.id
	        LEFT JOIN
	            LATERAL (
	                SELECT *
	                FROM pricechanges AS pc
	                WHERE pc.advertid = a.id
	                ORDER BY pc.datecreation DESC
	                LIMIT 1
	            ) AS pc ON TRUE
	        WHERE a.id=$1;`)
		queryUpdateAdvertById := regexp.QuoteMeta(`UPDATE adverts SET adverttypeplacement=$1, title=$2, description=$3, phone=$4, isagent=$5 WHERE id=$6;`)
		queryUpdateAdvertTypeById := regexp.QuoteMeta(`UPDATE adverttypes SET adverttype=$1 WHERE id=$2;`)
		queryUpdateBuildingById := regexp.QuoteMeta(`UPDATE buildings SET floor=$1, material=$2, adress=$3, adresspoint=$4, yearcreation=$5 WHERE id=$6;`)
		queryUpdateHouseById := regexp.QuoteMeta(`UPDATE houses SET ceilingheight=$1, squarearea=$2, squarehouse=$3, bedroomcount=$4, statusarea=$5, cottage=$6, statushome=$7 WHERE id=$8;`)
		price := 124.124
		suite.mock.ExpectQuery(queryGetIdTables).WithArgs(advertUpdateData.ID).
			WillReturnRows(sqlmock.NewRows([]string{"adverttypeid", "advType", "awd", "price"}).AddRow(advertTypeId, buildingId, houseId, price))

		suite.mock.ExpectExec(queryUpdateAdvertById).WithArgs(advertUpdateData.TypeSale, advertUpdateData.Title, advertUpdateData.Description, advertUpdateData.Phone, advertUpdateData.IsAgent, advertUpdateData.ID).WillReturnError(nil).
			WillReturnResult(sqlmock.NewResult(1, 1))

		suite.mock.ExpectExec(queryUpdateAdvertTypeById).WithArgs(advertUpdateData.TypeAdvert, advertTypeId).WillReturnError(nil).
			WillReturnResult(sqlmock.NewResult(1, 1))

		suite.mock.ExpectExec(queryUpdateBuildingById).WithArgs(advertUpdateData.HouseProperties.Floor, advertUpdateData.Material, advertUpdateData.Address, advertUpdateData.AddressPoint, advertUpdateData.YearCreation, buildingId).WillReturnError(nil).
			WillReturnResult(sqlmock.NewResult(1, 1))

		suite.mock.ExpectExec(queryUpdateHouseById).WithArgs(advertUpdateData.HouseProperties.CeilingHeight,
			advertUpdateData.HouseProperties.SquareArea, advertUpdateData.HouseProperties.SquareHouse,
			advertUpdateData.HouseProperties.BedroomCount, advertUpdateData.HouseProperties.StatusArea,
			advertUpdateData.HouseProperties.Cottage, advertUpdateData.HouseProperties.StatusHome, houseId).WillReturnError(nil).
			WillReturnResult(sqlmock.NewResult(1, 1))

		queryInsertPriceChange := regexp.QuoteMeta(`INSERT INTO pricechanges (id, advertId, price)
	            VALUES ($1, $2, $3)`)

		suite.mock.ExpectExec(queryInsertPriceChange).WithArgs(sqlmock.AnyArg(), advertUpdateData.ID, advertUpdateData.Price).WillReturnError(nil).
			WillReturnResult(sqlmock.NewResult(1, 1))

		logger := zap.Must(zap.NewDevelopment())
		rep := repo.NewRepository(suite.db, logger)
		err = rep.UpdateHouseAdvertById(ctx, tx, advertUpdateData)
		suite.Assert().NoError(err)

		// err = tx.Commit()

		// suite.Assert().NoError(err)

		err = suite.mock.ExpectationsWereMet()
		suite.Assert().NoError(err)
		suite.db.Close()
	}

	func (suite *UserRepoTestSuite) TestAdvertRepo_UpdateFlatAdvertById() {
		// advertId := uuid.NewV4()
		advertTypeId := uuid.NewV4()
		houseId := uuid.NewV4()
		buildingId := uuid.NewV4()
		ctx := context.WithValue(context.Background(), "requestId", uuid.NewV4().String())
		// advertType := models.AdvertTypeHouse
		suite.mock.ExpectBegin()
		tx, err := suite.db.Begin()
		suite.NoError(err)
		advertUpdateData := &models.AdvertUpdateData{
			ID:              uuid.NewV4(), // Генерируем новый UUID
			TypeAdvert:      "Flat",
			TypeSale:        "Sale",
			Title:           "Beautiful Apartment for Sale",
			Description:     "Spacious apartment with great view",
			Price:           100000,
			Phone:           "+1234567890",
			IsAgent:         true,
			Address:         "123 Main Street",
			AddressPoint:    "Latitude: 40.7128, Longitude: -74.0060",
			HouseProperties: &models.HouseProperties{
				// Заполняем данные для HouseProperties
				// Например: BedroomCount, BathroomCount, SquareHouse и т.д.
			},
			FlatProperties: &models.FlatProperties{
				// Заполняем данные для FlatProperties
				// Например: Floor, SquareGeneral, RoomCount и т.д.
			},
			YearCreation: 2000,
			Material:     "Brick",
		}
		queryGetIdTables := regexp.QuoteMeta(`
	        SELECT
	            at.id as adverttypeid,
	            b.id as buildingid,
	            f.id as flatid,
	            pc.price
	        FROM
	            adverts AS a
	        JOIN
	            adverttypes AS at ON a.adverttypeid = at.id
	        JOIN
	            flats AS f ON f.adverttypeid = at.id
	        JOIN
	            buildings AS b ON f.buildingid = b.id
	        LEFT JOIN
	            LATERAL (
	                SELECT *
	                FROM pricechanges AS pc
	                WHERE pc.advertid = a.id
	                ORDER BY pc.datecreation DESC
	                LIMIT 1
	            ) AS pc ON TRUE
	        WHERE a.id=$1;`)
		queryUpdateAdvertById := regexp.QuoteMeta(`UPDATE adverts SET adverttypeplacement=$1, title=$2, description=$3, phone=$4, isagent=$5 WHERE id=$6;`)
		queryUpdateAdvertTypeById := regexp.QuoteMeta(`UPDATE adverttypes SET adverttype=$1 WHERE id=$2;`)
		queryUpdateBuildingById := regexp.QuoteMeta(`UPDATE buildings SET floor=$1, material=$2, adress=$3, adresspoint=$4, yearcreation=$5 WHERE id=$6;`)
		queryUpdateFlatById := regexp.QuoteMeta(`UPDATE flats SET floor=$1, ceilingheight=$2, squaregeneral=$3, roomcount=$4, squareresidential=$5, apartament=$6 WHERE id=$7;`)
		price := 124.124
		suite.mock.ExpectQuery(queryGetIdTables).WithArgs(advertUpdateData.ID).
			WillReturnRows(sqlmock.NewRows([]string{"adverttypeid", "advType", "awd", "price"}).AddRow(advertTypeId, buildingId, houseId, price))

		suite.mock.ExpectExec(queryUpdateAdvertById).WithArgs(advertUpdateData.TypeSale, advertUpdateData.Title, advertUpdateData.Description, advertUpdateData.Phone, advertUpdateData.IsAgent, advertUpdateData.ID).WillReturnError(nil).
			WillReturnResult(sqlmock.NewResult(1, 1))

		suite.mock.ExpectExec(queryUpdateAdvertTypeById).WithArgs(advertUpdateData.TypeAdvert, advertTypeId).WillReturnError(nil).
			WillReturnResult(sqlmock.NewResult(1, 1))

		suite.mock.ExpectExec(queryUpdateBuildingById).WithArgs(advertUpdateData.HouseProperties.Floor, advertUpdateData.Material, advertUpdateData.Address, advertUpdateData.AddressPoint, advertUpdateData.YearCreation, buildingId).WillReturnError(nil).
			WillReturnResult(sqlmock.NewResult(1, 1))

		suite.mock.ExpectExec(queryUpdateFlatById).WithArgs(advertUpdateData.FlatProperties.Floor, advertUpdateData.FlatProperties.CeilingHeight, advertUpdateData.FlatProperties.SquareGeneral, advertUpdateData.FlatProperties.RoomCount, advertUpdateData.FlatProperties.SquareResidential, advertUpdateData.FlatProperties.Apartment, sqlmock.AnyArg()).WillReturnError(nil).
			WillReturnResult(sqlmock.NewResult(1, 1))

		queryInsertPriceChange := regexp.QuoteMeta(`INSERT INTO pricechanges (id, advertId, price)
	            VALUES ($1, $2, $3)`)

		suite.mock.ExpectExec(queryInsertPriceChange).WithArgs(sqlmock.AnyArg(), advertUpdateData.ID, advertUpdateData.Price).WillReturnError(nil).
			WillReturnResult(sqlmock.NewResult(1, 1))

		logger := zap.Must(zap.NewDevelopment())
		rep := repo.NewRepository(suite.db, logger)
		err = rep.UpdateFlatAdvertById(ctx, tx, advertUpdateData)
		suite.Assert().NoError(err)

		// err = tx.Commit()

		// suite.Assert().NoError(err)

		err = suite.mock.ExpectationsWereMet()
		suite.Assert().NoError(err)
		suite.db.Close()
	}
func (suite *UserRepoTestSuite) TestAdvertRepo_GetFlatAdvertById() {
	// advertId := uuid.NewV4()
	ctx := context.WithValue(context.Background(), "requestId", uuid.NewV4().String())
	// advertType := models.AdvertTypeHouse
	// suite.mock.ExpectBegin()

	id := uuid.NewV4()
	advertData := &models.AdvertData{
		ID:           uuid.NewV4(),
		AdvertType:   "House",
		TypeSale:     "Sale",
		Title:        "Beautiful House for Sale",
		Description:  "Spacious house with a large garden",
		Price:        100000,
		Phone:        "123-456-7890",
		IsAgent:      true,
		Address:      "123 Main St, Cityville",
		AddressPoint: "Coordinates",
		// Images:       []*models.ImageResp{},
		FlatProperties: &models.FlatProperties{
			CeilingHeight:     2.7,
			FloorGeneral:      3,
			RoomCount:         2,
			SquareResidential: 2222.22,
			SquareGeneral:     2333.3,
			Apartment:         true,
			Floor:             2,
		},
		ComplexProperties: &models.ComplexAdvertProperties{
			ComplexId:    "1234",
			NameComplex:  "Luxury Estates",
			PhotoCompany: "luxury_estates.jpg",
			NameCompany:  "Elite Realty",
		},
		// YearCreation: time.Now().Year(),
		Material: "Brick",
		// DateCreation: time.Now(),
	}
	query := regexp.QuoteMeta(`
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
        f.floor,
        f.ceilingheight,
        f.squaregeneral,
        f.roomcount,
        f.squareresidential,
        f.apartament,
        b.floor AS floorGeneral,
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
        flats AS f ON f.adverttypeid = at.id
    JOIN
        buildings AS b ON f.buildingid = b.id
    LEFT JOIN
        complexes AS cx ON b.complexid = cx.id
    LEFT JOIN
        companies AS c ON cx.companyid = c.id
    LEFT JOIN
        LATERAL (
            SELECT *
            FROM pricechanges AS pc
            WHERE pc.advertid = a.id
            ORDER BY pc.datecreation DESC
            LIMIT 1
        ) AS pc ON TRUE
    WHERE
        a.id = $1 AND a.isdeleted = FALSE;`)
	advertData.FlatProperties.Floor = 2
	suite.mock.ExpectQuery(query).WithArgs(id).WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15",
			"16", "17", "18", "19", "20", "21", "22", "23", "24"}).AddRow(
			advertData.ID,
			advertData.AdvertType,
			advertData.TypeSale,
			advertData.Title,
			advertData.Description,
			advertData.Price,
			advertData.Phone,
			advertData.IsAgent,
			advertData.Address,
			advertData.AddressPoint,
			advertData.FlatProperties.Floor,
			advertData.FlatProperties.CeilingHeight,
			advertData.FlatProperties.SquareGeneral,
			advertData.FlatProperties.RoomCount,
			advertData.FlatProperties.SquareResidential,
			advertData.FlatProperties.Apartment,
			advertData.FlatProperties.FloorGeneral,
			&advertData.YearCreation,
			&advertData.Material,
			&advertData.DateCreation,
			advertData.ComplexProperties.ComplexId,
			advertData.ComplexProperties.PhotoCompany,
			advertData.ComplexProperties.NameCompany,
			advertData.ComplexProperties.NameComplex))

	logger := zap.Must(zap.NewDevelopment())
	rep := repo.NewRepository(suite.db, logger)
	_, err := rep.GetFlatAdvertById(ctx, id)
	suite.Assert().NoError(err)

	// err = tx.Commit()

	// suite.Assert().NoError(err)

	err = suite.mock.ExpectationsWereMet()
	suite.Assert().NoError(err)
	suite.db.Close()
}

func (suite *UserRepoTestSuite) TestAdvertRepo_GetRectangleAdvertsByUserId() {
	// advertId := uuid.NewV4()
	ctx := context.WithValue(context.Background(), "requestId", uuid.NewV4().String())
	// advertType := models.AdvertTypeHouse
	// suite.mock.ExpectBegin()

	rectangleAdvert := &models.AdvertRectangleData{
		ID:          uuid.NewV4(),
		TypeAdvert:  "House",
		TypeSale:    "Sale",
		Title:       "Beautiful House for Sale",
		Description: "Spacious house with a large garden",
		Price:       100000,
		Phone:       "123-456-7890",
		Address:     "123 Main St, Cityville",
		// Images:       []*models.ImageResp{},
		FlatProperties: &models.FlatRectangleProperties{
			FloorGeneral:  3,
			RoomCount:     2,
			SquareGeneral: 2333.3,
			Floor:         2,
		},

		// DateCreation: time.Now(),
	}
	queryBaseAdvert := regexp.QuoteMeta(`
        SELECT
            a.id,
            a.title,
            a.description,
            at.adverttype,
            CASE
                WHEN at.adverttype = 'Flat' THEN f.roomcount
                WHEN at.adverttype = 'House' THEN h.bedroomcount
                ELSE NULL
            END AS rcount,
            a.phone,
            a.adverttypeplacement,
            b.adress,
            pc.price,
            i.photo,
            a.datecreation
        FROM
            adverts AS a
            JOIN adverttypes AS at ON a.adverttypeid = at.id
            LEFT JOIN flats AS f ON f.adverttypeid = at.id
            LEFT JOIN houses AS h ON h.adverttypeid = at.id
            LEFT JOIN buildings AS b ON (f.buildingid = b.id OR h.buildingid = b.id)
            LEFT JOIN LATERAL (
                SELECT *
                FROM pricechanges AS pc
                WHERE pc.advertid = a.id
                ORDER BY pc.datecreation DESC
                LIMIT 1
            ) AS pc ON TRUE
            JOIN images AS i ON i.advertid = a.id
        WHERE i.priority = (
                SELECT MIN(priority)
                FROM images
                WHERE advertid = a.id
                    AND isdeleted = FALSE
            )
            AND i.isdeleted = FALSE
            AND a.isdeleted = FALSE
            AND userid = $1
        ORDER BY datecreation DESC
        LIMIT $2
        OFFSET $3;`)
	queryHouse := regexp.QuoteMeta(`
        SELECT
            b.adress,
            h.cottage,
            h.squarehouse,
            h.squarearea,
            b.floor
        FROM
            adverts AS a
            JOIN adverttypes AS at ON a.adverttypeid = at.id
            JOIN houses AS h ON h.adverttypeid = at.id
            JOIN buildings AS b ON h.buildingid = b.id
        WHERE a.id = $1
        ORDER BY
            a.datecreation DESC;`)

	userId := uuid.NewV4()
	pageSize := 3
	offset := 2
	roomCount := 4
	suite.mock.ExpectQuery(queryBaseAdvert).WithArgs(userId, pageSize, offset).WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}).AddRow(
			rectangleAdvert.ID, rectangleAdvert.Title, rectangleAdvert.Description, rectangleAdvert.TypeAdvert,
			roomCount, rectangleAdvert.Phone, rectangleAdvert.TypeSale, rectangleAdvert.Address, rectangleAdvert.Price,
			rectangleAdvert.Photo, rectangleAdvert.DateCreation))

	suite.mock.ExpectQuery(queryHouse).WithArgs(rectangleAdvert.ID).WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{"1", "2", "3", "4", "5"}).AddRow(
			rectangleAdvert.Address,
			true, 124.44, 444.444, 4))
	logger := zap.Must(zap.NewDevelopment())
	rep := repo.NewRepository(suite.db, logger)
	_, err := rep.GetRectangleAdvertsByUserId(ctx, pageSize, offset, userId)
	suite.Assert().NoError(err)

	// err = tx.Commit()

	// suite.Assert().NoError(err)

	err = suite.mock.ExpectationsWereMet()
	suite.Assert().NoError(err)
	suite.db.Close()
}

func (suite *UserRepoTestSuite) TestAdvertRepo_GetRectangleAdvertsByComplexId() {
	// advertId := uuid.NewV4()
	ctx := context.WithValue(context.Background(), "requestId", uuid.NewV4().String())
	// advertType := models.AdvertTypeHouse
	// suite.mock.ExpectBegin()

	rectangleAdvert := &models.AdvertRectangleData{
		ID:          uuid.NewV4(),
		TypeAdvert:  "House",
		TypeSale:    "Sale",
		Title:       "Beautiful House for Sale",
		Description: "Spacious house with a large garden",
		Price:       100000,
		Phone:       "123-456-7890",
		Address:     "123 Main St, Cityville",
		// Images:       []*models.ImageResp{},
		FlatProperties: &models.FlatRectangleProperties{
			FloorGeneral:  3,
			RoomCount:     2,
			SquareGeneral: 2333.3,
			Floor:         2,
		},

		// DateCreation: time.Now(),
	}
	queryBaseAdvert := regexp.QuoteMeta(`
        SELECT
            a.id,
            a.title,
            a.description,
            at.adverttype,
            CASE
                WHEN at.adverttype = 'Flat' THEN f.roomcount
                WHEN at.adverttype = 'House' THEN h.bedroomcount
                ELSE 0
            END AS rcount,
            a.phone,
            a.adverttypeplacement,
            b.adress,
            pc.price,
            i.photo,
            a.datecreation
        FROM
            adverts AS a
            JOIN adverttypes AS at ON a.adverttypeid = at.id
            LEFT JOIN flats AS f ON f.adverttypeid = at.id
            LEFT JOIN houses AS h ON h.adverttypeid = at.id
            LEFT JOIN buildings AS b ON (f.buildingid = b.id OR h.buildingid = b.id)
            LEFT JOIN LATERAL (
                SELECT *
                FROM pricechanges AS pc
                WHERE pc.advertid = a.id
                ORDER BY pc.datecreation DESC
                LIMIT 1
            ) AS pc ON TRUE
            JOIN images AS i ON i.advertid = a.id
        WHERE i.priority = (
                SELECT MIN(priority)
                FROM images
                WHERE advertid = a.id
                    AND isdeleted = FALSE
            )
            AND i.isdeleted = FALSE
            AND a.isdeleted = FALSE
            AND b.complexid = $1
        ORDER BY datecreation DESC
        LIMIT $2
        OFFSET $3;`)
	queryHouse := regexp.QuoteMeta(`
        SELECT
            b.adress,
            h.cottage,
            h.squarehouse,
            h.squarearea,
            b.floor
        FROM
            adverts AS a
            JOIN adverttypes AS at ON a.adverttypeid = at.id
            JOIN houses AS h ON h.adverttypeid = at.id
            JOIN buildings AS b ON h.buildingid = b.id
        WHERE a.id = $1
        ORDER BY
            a.datecreation DESC;`)

	userId := uuid.NewV4()
	pageSize := 3
	offset := 2
	roomCount := 4
	suite.mock.ExpectQuery(queryBaseAdvert).WithArgs(userId, pageSize, offset).WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}).AddRow(
			rectangleAdvert.ID, rectangleAdvert.Title, rectangleAdvert.Description, rectangleAdvert.TypeAdvert,
			roomCount, rectangleAdvert.Phone, rectangleAdvert.TypeSale, rectangleAdvert.Address, rectangleAdvert.Price,
			rectangleAdvert.Photo, rectangleAdvert.DateCreation))

	suite.mock.ExpectQuery(queryHouse).WithArgs(rectangleAdvert.ID).WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{"1", "2", "3", "4", "5"}).AddRow(
			rectangleAdvert.Address,
			true, 124.44, 444.444, 4))
	logger := zap.Must(zap.NewDevelopment())
	rep := repo.NewRepository(suite.db, logger)
	_, err := rep.GetRectangleAdvertsByComplexId(ctx, pageSize, offset, userId)
	suite.Assert().NoError(err)

	// err = tx.Commit()

	// suite.Assert().NoError(err)

	err = suite.mock.ExpectationsWereMet()
	suite.Assert().NoError(err)
	suite.db.Close()
}
*/
