package repo_test

/*
import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts/repo"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/satori/uuid"
	"go.uber.org/zap"
	"regexp"
)

func (suite *AdvertRepoTestSuite) TestChangeTypeAdvert1() {
	type args struct {
		typeAdvert models.AdvertTypeAdvert
		buildingId int64
		advertId   int64
		houseId    int64
		flatId     int64
	}
	type want struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		want    want
		prepare func(a *args, w *want)
	}{
		{
			name: "successful TestChangeTypeAdvert1",
			args: args{
				typeAdvert: models.AdvertTypeFlat,
				advertId:   124,
				buildingId: 111,
				houseId:    33,
			},
			want: want{
				err: errors.New("some error"),
			},
			prepare: func(a *args, w *want) {
				query := `SELECT 			CASE
	WHEN ath.house_id IS NOT NULL THEN 'House'
	WHEN atf.flat_id IS NOT NULL THEN 'Flat'
	ELSE 'None'
END AS type_advert FROM advert AS a LEFT JOIN advert_type_flat AS atf ON a.id=atf.advert_id LEFT JOIN advert_type_house AS ath ON a.id=ath.advert_id WHERE a.id = $1;`
				escapedQuery := regexp.QuoteMeta(query)
				rows := sqlmock.NewRows([]string{"id"})
				rows = rows.AddRow(a.typeAdvert)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)
				querySelectBuildingIdByFlat := `SELECT b.id AS buildingid, f.id AS flatid  FROM advert AS a JOIN advert_type_flat AS at ON at.advert_id=a.id JOIN flat AS f ON f.id=at.flat_id JOIN building AS b ON f.building_id=b.id WHERE a.id=$1`

				escapedQuery = regexp.QuoteMeta(querySelectBuildingIdByFlat)
				rows = sqlmock.NewRows([]string{"bid", "fid"})
				rows = rows.AddRow(a.buildingId, a.flatId)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				queryDeleteFlatById := `UPDATE flat SET is_deleted=true WHERE id=$1;`
				escapedQuery = regexp.QuoteMeta(queryDeleteFlatById)
				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))

				queryDeleteAdvertTypeFlat := `UPDATE advert_type_flat SET is_deleted=true WHERE advert_id=$1 AND flat_id=$2;`
				escapedQuery = regexp.QuoteMeta(queryDeleteAdvertTypeFlat)

				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))

				query = `SELECT h.id FROM advert AS a JOIN advert_type_house AS at ON a.id=at.advert_id JOIN house AS h ON h.id=at.house_id WHERE a.id = $1;`
				escapedQuery = regexp.QuoteMeta(query)
				rows = sqlmock.NewRows([]string{"id"})
				rows = rows.AddRow(a.houseId)

				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				queryRestoreHouseById := `UPDATE house SET is_deleted=false WHERE id=$1;`

				escapedQuery = regexp.QuoteMeta(queryRestoreHouseById)
				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(errors.New("some error")).WithArgs(sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))

			},
		},

		{
			name: "successful TestChangeTypeAdvert1",
			args: args{
				typeAdvert: models.AdvertTypeFlat,
				advertId:   124,
				buildingId: 111,
				houseId:    33,
			},
			want: want{
				err: errors.New("some error"),
			},
			prepare: func(a *args, w *want) {
				query := `SELECT 			CASE
	WHEN ath.house_id IS NOT NULL THEN 'House'
	WHEN atf.flat_id IS NOT NULL THEN 'Flat'
	ELSE 'None'
END AS type_advert FROM advert AS a LEFT JOIN advert_type_flat AS atf ON a.id=atf.advert_id LEFT JOIN advert_type_house AS ath ON a.id=ath.advert_id WHERE a.id = $1;`
				escapedQuery := regexp.QuoteMeta(query)
				rows := sqlmock.NewRows([]string{"id"})
				rows = rows.AddRow(a.typeAdvert)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)
				querySelectBuildingIdByFlat := `SELECT b.id AS buildingid, f.id AS flatid  FROM advert AS a JOIN advert_type_flat AS at ON at.advert_id=a.id JOIN flat AS f ON f.id=at.flat_id JOIN building AS b ON f.building_id=b.id WHERE a.id=$1`

				escapedQuery = regexp.QuoteMeta(querySelectBuildingIdByFlat)
				rows = sqlmock.NewRows([]string{"bid", "fid"})
				rows = rows.AddRow(a.buildingId, a.flatId)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				queryDeleteFlatById := `UPDATE flat SET is_deleted=true WHERE id=$1;`
				escapedQuery = regexp.QuoteMeta(queryDeleteFlatById)
				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))

				queryDeleteAdvertTypeFlat := `UPDATE advert_type_flat SET is_deleted=true WHERE advert_id=$1 AND flat_id=$2;`
				escapedQuery = regexp.QuoteMeta(queryDeleteAdvertTypeFlat)

				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))

				query = `SELECT h.id FROM advert AS a JOIN advert_type_house AS at ON a.id=at.advert_id JOIN house AS h ON h.id=at.house_id WHERE a.id = $1;`
				escapedQuery = regexp.QuoteMeta(query)
				rows = sqlmock.NewRows([]string{"id"})
				rows = rows.AddRow(a.houseId)

				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				queryRestoreHouseById := `UPDATE house SET is_deleted=false WHERE id=$1;`

				escapedQuery = regexp.QuoteMeta(queryRestoreHouseById)
				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))

				queryRestoreAdvertTypeHouse := `UPDATE advert_type_house SET is_deleted=false WHERE advert_id=$1 AND house_id=$2;`
				escapedQuery = regexp.QuoteMeta(queryRestoreAdvertTypeHouse)
				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(errors.New("some error")).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},

		{
			name: "successful TestChangeTypeAdvert1",
			args: args{
				typeAdvert: models.AdvertTypeHouse,
				advertId:   124,
				buildingId: 111,
				houseId:    33,
			},
			want: want{
				err: errors.New("some error"),
			},
			prepare: func(a *args, w *want) {
				query := `SELECT 			CASE
	WHEN ath.house_id IS NOT NULL THEN 'House'
	WHEN atf.flat_id IS NOT NULL THEN 'Flat'
	ELSE 'None'
END AS type_advert FROM advert AS a LEFT JOIN advert_type_flat AS atf ON a.id=atf.advert_id LEFT JOIN advert_type_house AS ath ON a.id=ath.advert_id WHERE a.id = $1;`
				escapedQuery := regexp.QuoteMeta(query)
				rows := sqlmock.NewRows([]string{"id"})
				rows = rows.AddRow(a.typeAdvert)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				querySelectBuildingIdByHouse := `SELECT b.id AS buildingid, h.id AS houseid  FROM advert AS a JOIN advert_type_house AS at ON at.advert_id=a.id JOIN house AS h ON h.id=at.house_id JOIN building AS b ON h.building_id=b.id WHERE a.id=$1`

				escapedQuery = regexp.QuoteMeta(querySelectBuildingIdByHouse)
				rows = sqlmock.NewRows([]string{"1"})
				rows = rows.AddRow(a.buildingId)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(errors.New("some error")).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)
			},
		},

		{
			name: "successful TestChangeTypeAdvert1",
			args: args{
				typeAdvert: models.AdvertTypeHouse,
				advertId:   124,
				buildingId: 111,
				houseId:    33,
			},
			want: want{
				err: errors.New("some error"),
			},
			prepare: func(a *args, w *want) {
				query := `SELECT 			CASE
	WHEN ath.house_id IS NOT NULL THEN 'House'
	WHEN atf.flat_id IS NOT NULL THEN 'Flat'
	ELSE 'None'
END AS type_advert FROM advert AS a LEFT JOIN advert_type_flat AS atf ON a.id=atf.advert_id LEFT JOIN advert_type_house AS ath ON a.id=ath.advert_id WHERE a.id = $1;`
				escapedQuery := regexp.QuoteMeta(query)
				rows := sqlmock.NewRows([]string{"id"})
				rows = rows.AddRow(a.typeAdvert)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				querySelectBuildingIdByHouse := `SELECT b.id AS buildingid, h.id AS houseid  FROM advert AS a JOIN advert_type_house AS at ON at.advert_id=a.id JOIN house AS h ON h.id=at.house_id JOIN building AS b ON h.building_id=b.id WHERE a.id=$1`

				escapedQuery = regexp.QuoteMeta(querySelectBuildingIdByHouse)
				rows = sqlmock.NewRows([]string{"1", "2"})
				rows = rows.AddRow(a.buildingId, a.houseId)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				queryDeleteHouseById := `UPDATE house SET is_deleted=true WHERE id=$1;`
				escapedQuery = regexp.QuoteMeta(queryDeleteHouseById)
				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(errors.New("some error")).WithArgs(sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))

			},
		},

		{
			name: "successful TestChangeTypeAdvert1",
			args: args{
				typeAdvert: models.AdvertTypeHouse,
				advertId:   124,
				buildingId: 111,
				houseId:    33,
			},
			want: want{
				err: errors.New("some error"),
			},
			prepare: func(a *args, w *want) {
				query := `SELECT 			CASE
	WHEN ath.house_id IS NOT NULL THEN 'House'
	WHEN atf.flat_id IS NOT NULL THEN 'Flat'
	ELSE 'None'
END AS type_advert FROM advert AS a LEFT JOIN advert_type_flat AS atf ON a.id=atf.advert_id LEFT JOIN advert_type_house AS ath ON a.id=ath.advert_id WHERE a.id = $1;`
				escapedQuery := regexp.QuoteMeta(query)
				rows := sqlmock.NewRows([]string{"id"})
				rows = rows.AddRow(a.typeAdvert)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				querySelectBuildingIdByHouse := `SELECT b.id AS buildingid, h.id AS houseid  FROM advert AS a JOIN advert_type_house AS at ON at.advert_id=a.id JOIN house AS h ON h.id=at.house_id JOIN building AS b ON h.building_id=b.id WHERE a.id=$1`

				escapedQuery = regexp.QuoteMeta(querySelectBuildingIdByHouse)
				rows = sqlmock.NewRows([]string{"1", "2"})
				rows = rows.AddRow(a.buildingId, a.houseId)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				queryDeleteHouseById := `UPDATE house SET is_deleted=true WHERE id=$1;`
				escapedQuery = regexp.QuoteMeta(queryDeleteHouseById)
				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))

				queryDeleteAdvertTypeHouse := `UPDATE advert_type_house SET is_deleted=true WHERE advert_id=$1 AND house_id=$2;`
				escapedQuery = regexp.QuoteMeta(queryDeleteAdvertTypeHouse)
				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(errors.New("some error")).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},

		{
			name: "successful TestChangeTypeAdvert1",
			args: args{
				typeAdvert: models.AdvertTypeHouse,
				advertId:   124,
				buildingId: 111,
				houseId:    33,
			},
			want: want{
				err: errors.New("some error"),
			},
			prepare: func(a *args, w *want) {
				query := `SELECT 			CASE
	WHEN ath.house_id IS NOT NULL THEN 'House'
	WHEN atf.flat_id IS NOT NULL THEN 'Flat'
	ELSE 'None'
END AS type_advert FROM advert AS a LEFT JOIN advert_type_flat AS atf ON a.id=atf.advert_id LEFT JOIN advert_type_house AS ath ON a.id=ath.advert_id WHERE a.id = $1;`
				escapedQuery := regexp.QuoteMeta(query)
				rows := sqlmock.NewRows([]string{"id"})
				rows = rows.AddRow(a.typeAdvert)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				querySelectBuildingIdByHouse := `SELECT b.id AS buildingid, h.id AS houseid  FROM advert AS a JOIN advert_type_house AS at ON at.advert_id=a.id JOIN house AS h ON h.id=at.house_id JOIN building AS b ON h.building_id=b.id WHERE a.id=$1`
				escapedQuery = regexp.QuoteMeta(querySelectBuildingIdByHouse)
				rows = sqlmock.NewRows([]string{"1", "2"})
				rows = rows.AddRow(a.buildingId, a.houseId)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				queryDeleteHouseById := `UPDATE house SET is_deleted=true WHERE id=$1;`
				escapedQuery = regexp.QuoteMeta(queryDeleteHouseById)
				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))

				queryDeleteAdvertTypeHouse := `UPDATE advert_type_house SET is_deleted=true WHERE advert_id=$1 AND house_id=$2;`
				escapedQuery = regexp.QuoteMeta(queryDeleteAdvertTypeHouse)
				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))

				// check exosts flat query!!
				query = `SELECT f.id FROM advert AS a JOIN advert_type_flat AS at ON a.id=at.advert_id JOIN flat AS f ON f.id=at.flat_id WHERE a.id = $1`
				escapedQuery = regexp.QuoteMeta(query)
				rows = sqlmock.NewRows([]string{"1"})
				rows = rows.AddRow(a.advertId)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(errors.New("some error")).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				queryInsertFlat := `INSERT INTO flat (building_id, floor, ceiling_height, square_general, bedroom_count, square_residential, apartament)
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
				escapedQuery = regexp.QuoteMeta(queryInsertFlat)
				rows = sqlmock.NewRows([]string{"1"})
				rows = rows.AddRow(a.flatId)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(errors.New("some error")).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(rows)

			},
		},

		{
			name: "successful TestChangeTypeAdvert1",
			args: args{
				typeAdvert: models.AdvertTypeHouse,
				advertId:   124,
				buildingId: 111,
				houseId:    33,
			},
			want: want{
				err: errors.New("some error"),
			},
			prepare: func(a *args, w *want) {
				query := `SELECT 			CASE
	WHEN ath.house_id IS NOT NULL THEN 'House'
	WHEN atf.flat_id IS NOT NULL THEN 'Flat'
	ELSE 'None'
END AS type_advert FROM advert AS a LEFT JOIN advert_type_flat AS atf ON a.id=atf.advert_id LEFT JOIN advert_type_house AS ath ON a.id=ath.advert_id WHERE a.id = $1;`
				escapedQuery := regexp.QuoteMeta(query)
				rows := sqlmock.NewRows([]string{"id"})
				rows = rows.AddRow(a.typeAdvert)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				querySelectBuildingIdByHouse := `SELECT b.id AS buildingid, h.id AS houseid  FROM advert AS a JOIN advert_type_house AS at ON at.advert_id=a.id JOIN house AS h ON h.id=at.house_id JOIN building AS b ON h.building_id=b.id WHERE a.id=$1`
				escapedQuery = regexp.QuoteMeta(querySelectBuildingIdByHouse)
				rows = sqlmock.NewRows([]string{"1", "2"})
				rows = rows.AddRow(a.buildingId, a.houseId)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				queryDeleteHouseById := `UPDATE house SET is_deleted=true WHERE id=$1;`
				escapedQuery = regexp.QuoteMeta(queryDeleteHouseById)
				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))

				queryDeleteAdvertTypeHouse := `UPDATE advert_type_house SET is_deleted=true WHERE advert_id=$1 AND house_id=$2;`
				escapedQuery = regexp.QuoteMeta(queryDeleteAdvertTypeHouse)
				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))

				// check exosts flat query!!
				query = `SELECT f.id FROM advert AS a JOIN advert_type_flat AS at ON a.id=at.advert_id JOIN flat AS f ON f.id=at.flat_id WHERE a.id = $1`
				escapedQuery = regexp.QuoteMeta(query)
				rows = sqlmock.NewRows([]string{"1"})
				rows = rows.AddRow(a.advertId)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(errors.New("some error")).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				queryInsertFlat := `INSERT INTO flat (building_id, floor, ceiling_height, square_general, bedroom_count, square_residential, apartament)
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
				escapedQuery = regexp.QuoteMeta(queryInsertFlat)
				rows = sqlmock.NewRows([]string{"1"})
				rows = rows.AddRow(a.flatId)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
					sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(rows)

				queryInsertTypeFlat := `INSERT INTO advert_type_flat (advert_id, flat_id) VALUES ($1, $2);`
				escapedQuery = regexp.QuoteMeta(queryInsertTypeFlat)
				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(errors.New("some error")).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))

			},
		},

		{
			name: "successful TestChangeTypeAdvert1",
			args: args{
				typeAdvert: models.AdvertTypeHouse,
				advertId:   124,
				buildingId: 111,
				houseId:    33,
			},
			want: want{
				err: errors.New("some error"),
			},
			prepare: func(a *args, w *want) {
				query := `SELECT 			CASE
	WHEN ath.house_id IS NOT NULL THEN 'House'
	WHEN atf.flat_id IS NOT NULL THEN 'Flat'
	ELSE 'None'
END AS type_advert FROM advert AS a LEFT JOIN advert_type_flat AS atf ON a.id=atf.advert_id LEFT JOIN advert_type_house AS ath ON a.id=ath.advert_id WHERE a.id = $1;`
				escapedQuery := regexp.QuoteMeta(query)
				rows := sqlmock.NewRows([]string{"id"})
				rows = rows.AddRow(a.typeAdvert)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				querySelectBuildingIdByHouse := `SELECT b.id AS buildingid, h.id AS houseid  FROM advert AS a JOIN advert_type_house AS at ON at.advert_id=a.id JOIN house AS h ON h.id=at.house_id JOIN building AS b ON h.building_id=b.id WHERE a.id=$1`
				escapedQuery = regexp.QuoteMeta(querySelectBuildingIdByHouse)
				rows = sqlmock.NewRows([]string{"1", "2"})
				rows = rows.AddRow(a.buildingId, a.houseId)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				queryDeleteHouseById := `UPDATE house SET is_deleted=true WHERE id=$1;`
				escapedQuery = regexp.QuoteMeta(queryDeleteHouseById)
				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))

				queryDeleteAdvertTypeHouse := `UPDATE advert_type_house SET is_deleted=true WHERE advert_id=$1 AND house_id=$2;`
				escapedQuery = regexp.QuoteMeta(queryDeleteAdvertTypeHouse)
				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))

				// check exosts flat query!!
				query = `SELECT f.id FROM advert AS a JOIN advert_type_flat AS at ON a.id=at.advert_id JOIN flat AS f ON f.id=at.flat_id WHERE a.id = $1`
				escapedQuery = regexp.QuoteMeta(query)
				rows = sqlmock.NewRows([]string{"1"})
				rows = rows.AddRow(a.advertId)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				queryRestoreFlatById := `UPDATE flat SET is_deleted=false WHERE id=$1;`
				escapedQuery = regexp.QuoteMeta(queryRestoreFlatById)

				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(errors.New("some error")).WithArgs(sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))

			},
		},

		{
			name: "successful TestChangeTypeAdvert1",
			args: args{
				typeAdvert: models.AdvertTypeHouse,
				advertId:   124,
				buildingId: 111,
				houseId:    33,
			},
			want: want{
				err: errors.New("some error"),
			},
			prepare: func(a *args, w *want) {
				query := `SELECT 			CASE
	WHEN ath.house_id IS NOT NULL THEN 'House'
	WHEN atf.flat_id IS NOT NULL THEN 'Flat'
	ELSE 'None'
END AS type_advert FROM advert AS a LEFT JOIN advert_type_flat AS atf ON a.id=atf.advert_id LEFT JOIN advert_type_house AS ath ON a.id=ath.advert_id WHERE a.id = $1;`
				escapedQuery := regexp.QuoteMeta(query)
				rows := sqlmock.NewRows([]string{"id"})
				rows = rows.AddRow(a.typeAdvert)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				querySelectBuildingIdByHouse := `SELECT b.id AS buildingid, h.id AS houseid  FROM advert AS a JOIN advert_type_house AS at ON at.advert_id=a.id JOIN house AS h ON h.id=at.house_id JOIN building AS b ON h.building_id=b.id WHERE a.id=$1`
				escapedQuery = regexp.QuoteMeta(querySelectBuildingIdByHouse)
				rows = sqlmock.NewRows([]string{"1", "2"})
				rows = rows.AddRow(a.buildingId, a.houseId)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				queryDeleteHouseById := `UPDATE house SET is_deleted=true WHERE id=$1;`
				escapedQuery = regexp.QuoteMeta(queryDeleteHouseById)
				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))

				queryDeleteAdvertTypeHouse := `UPDATE advert_type_house SET is_deleted=true WHERE advert_id=$1 AND house_id=$2;`
				escapedQuery = regexp.QuoteMeta(queryDeleteAdvertTypeHouse)
				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))

				// check exosts flat query!!
				query = `SELECT f.id FROM advert AS a JOIN advert_type_flat AS at ON a.id=at.advert_id JOIN flat AS f ON f.id=at.flat_id WHERE a.id = $1`
				escapedQuery = regexp.QuoteMeta(query)
				rows = sqlmock.NewRows([]string{"1"})
				rows = rows.AddRow(a.advertId)
				suite.mock.ExpectQuery(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)

				queryRestoreFlatById := `UPDATE flat SET is_deleted=false WHERE id=$1;`
				escapedQuery = regexp.QuoteMeta(queryRestoreFlatById)

				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(nil).WithArgs(sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))

				queryRestoreAdvertTypeFlat := `UPDATE advert_type_flat SET is_deleted=false WHERE advert_id=$1 AND flat_id=$2;`
				escapedQuery = regexp.QuoteMeta(queryRestoreAdvertTypeFlat)

				suite.mock.ExpectExec(escapedQuery).
					WillReturnError(errors.New("some error")).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
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
			tt.prepare(&tt.args, &tt.want)
			logger := zap.Must(zap.NewDevelopment())
			rep := repo.NewRepository(suite.db, logger)
			gotErr := rep.ChangeTypeAdvert(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()),
				tx, tt.args.advertId)

			suite.Assert().Equal(tt.want.err, gotErr)
			suite.Assert().NoError(suite.mock.ExpectationsWereMet())
		})
	}
}

*/
