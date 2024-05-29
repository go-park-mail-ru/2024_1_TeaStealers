package repo_test

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts/repo"
	"2024_1_TeaStealers/internal/pkg/middleware"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"regexp"
	"testing"
)

// func TestGetHouseAdvertById(t *testing.T) {
//	// Initialize your repo, logger, and other dependencies here
//	var db *sql.DB
//	var mock sqlmock.Sqlmock
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer db.Close()
//
//	logger := zap.Must(zap.NewDevelopment())
//	repo := repo.NewRepository(db, logger)
//
//	// Set up a mock context with a user ID
//	ctx := context.Background()
//	ctx = context.WithValue(ctx, middleware.CookieName, int64(1))
//	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())
//
//	// Define expected advert data
//	expectedAdvertData := &models.AdvertData{
//		ID:           1,
//		AdvertType:   "House",
//		TypeSale:     "Sale",
//		Title:        "Test House",
//		Description:  "Test Description",
//		CountViews:   0, // обновить это значение
//		CountLikes:   0, // обновить это значение
//		Price:        1000000,
//		Phone:        "1234567890",
//		IsLiked:      true,
//		IsAgent:      false,
//		Metro:        "Test Metro",
//		Address:      "Test Province, Test Town, Test Street, Test House Name", // обновить это значение
//		AddressPoint: "POINT(0 0)",
//		PriceChange:  nil,
//		Images:       nil,
//		HouseProperties: &models.HouseProperties{
//			CeilingHeight: 3.0,
//			SquareArea:    100.0,
//			SquareHouse:   200.0,
//			BedroomCount:  3,
//			StatusArea:    "Living",
//			Cottage:       true,
//			StatusHome:    "New",
//			Floor:         2,
//		},
//		YearCreation: 2020,
//		Material:     "Brick",
//		ComplexProperties: &models.ComplexAdvertProperties{
//			ComplexId:    "1",
//			PhotoCompany: "company_photo.jpg",
//			NameCompany:  "Test Company",
//			NameComplex:  "Test Complex",
//		},
//		DateCreation: time.Now(),
//	}
//
//	// Mock the QueryRowContext function to return expected data
//	rows := sqlmock.NewRows([]string{
//		"id",
//		"type_placement",
//		"title",
//		"description",
//		"price",
//		"phone",
//		"is_agent",
//		"metro",
//		"name",
//		"name",
//		"name",
//		"name",
//		"address_point",
//		"ceiling_height",
//		"square_area",
//		"square_house",
//		"bedroom_count",
//		"status_area_house",
//		"cottage",
//		"status_home_house",
//		"floor",
//		"year_creation",
//		"material_building",
//		"created_at",
//		"is_liked",
//		"id",
//		"photo",
//		"name",
//		"name",
//	}).AddRow(
//		expectedAdvertData.ID,
//		expectedAdvertData.TypeSale,
//		expectedAdvertData.Title,
//		expectedAdvertData.Description,
//		expectedAdvertData.Price,
//		expectedAdvertData.Phone,
//		expectedAdvertData.IsAgent,
//		expectedAdvertData.Metro,
//		"Test House Name",
//		"Test Street",
//		"Test Town",
//		"Test Province",
//		expectedAdvertData.AddressPoint,
//		expectedAdvertData.HouseProperties.CeilingHeight,
//		expectedAdvertData.HouseProperties.SquareArea,
//		expectedAdvertData.HouseProperties.SquareHouse,
//		expectedAdvertData.HouseProperties.BedroomCount,
//		expectedAdvertData.HouseProperties.StatusArea,
//		expectedAdvertData.HouseProperties.Cottage,
//		expectedAdvertData.HouseProperties.StatusHome,
//		expectedAdvertData.HouseProperties.Floor,
//		expectedAdvertData.YearCreation,
//		string(expectedAdvertData.Material),
//		expectedAdvertData.DateCreation,
//		true,
//		expectedAdvertData.ComplexProperties.ComplexId,
//		expectedAdvertData.ComplexProperties.PhotoCompany,
//		expectedAdvertData.ComplexProperties.NameCompany,
//		expectedAdvertData.ComplexProperties.NameComplex,
//	)
//
//	query := `
//	SELECT
//        a.id,
//        a.type_placement,
//        a.title,
//        a.description,
//        pc.price,
//        a.phone,
//        a.is_agent,
//		ad.metro,
//		hn.name,
//		s.name,
//		t.name,
//		p.name,
//		ST_AsText(ad.address_point::geometry),
//        h.ceiling_height,
//        h.square_area,
//        h.square_house,
//        h.bedroom_count,
//        h.status_area_house,
//        h.cottage,
//        h.status_home_house,
//        b.floor,
//        b.year_creation,
//        COALESCE(b.material_building, 'Brick') as material,
//        a.created_at,
//		CASE
//			WHEN fa.advert_id IS NOT NULL AND fa.is_deleted=false THEN true
//			ELSE false
//		END AS is_liked,
//        cx.id AS complexid,
//        c.photo AS companyphoto,
//        c.name AS companyname,
//        cx.name AS complexname
//    FROM
//        advert AS a
//    JOIN
//        advert_type_house AS at ON a.id = at.advert_id
//    JOIN
//        house AS h ON h.id = at.house_id
//    JOIN
//        building AS b ON h.building_id = b.id
//		JOIN address AS ad ON b.address_id=ad.id
//		JOIN house_name AS hn ON hn.id=ad.house_name_id
//		JOIN street AS s ON s.id=hn.street_id
//		JOIN town AS t ON t.id=s.town_id
//		JOIN province AS p ON p.id=t.province_id
//	LEFT JOIN
//		favourite_advert AS fa ON fa.advert_id=a.id AND fa.user_id=$2
//    LEFT JOIN
//        complex AS cx ON b.complex_id = cx.id
//    LEFT JOIN
//        company AS c ON cx.company_id = c.id
//    JOIN
//        LATERAL (
//            SELECT *
//            FROM price_change AS pc
//            WHERE pc.advert_id = a.id
//            ORDER BY pc.created_at DESC
//            LIMIT 1
//        ) AS pc ON TRUE
//    WHERE
//        a.id = $1 AND a.is_deleted = FALSE;`
//	escapedQuery := regexp.QuoteMeta(query)
//
//	mock.ExpectQuery(escapedQuery).
//		WithArgs(expectedAdvertData.ID, int64(1)).
//		WillReturnRows(rows)
//
//	// Call the GetHouseAdvertById function
//	advertData, err := repo.GetHouseAdvertById(ctx, expectedAdvertData.ID)
//
//	// Assert that there were no errors
//	require.NoError(t, err)
//
//	// Assert that the returned advert data matches the expected data
//	assert.Equal(t, expectedAdvertData, advertData)
//
//	// Assert that all expectations were met
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Fatal(err)
//	}
//}

func TestCheckExistsFlat(t *testing.T) {
	// Initialize your repo, logger, and other dependencies here
	var db *sql.DB
	var mock sqlmock.Sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	logger := zap.Must(zap.NewDevelopment())
	repo := repo.NewRepository(db, logger)

	// Set up a mock context with a user ID
	ctx := context.Background()
	ctx = context.WithValue(ctx, middleware.CookieName, int64(1))
	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())

	// Define expected flat data
	expectedFlat := &models.Flat{
		ID: 1,
	}

	// Mock the QueryRowContext function to return expected data
	rows := sqlmock.NewRows([]string{"id"}).AddRow(expectedFlat.ID)

	query := `SELECT f.id FROM advert AS a JOIN advert_type_flat AS at ON a.id=at.advert_id JOIN flat AS f ON f.id=at.flat_id WHERE a.id = $1`
	escapedQuery := regexp.QuoteMeta(query)

	mock.ExpectQuery(escapedQuery).
		WithArgs(expectedFlat.ID).
		WillReturnRows(rows)

	// Call the CheckExistsFlat function
	flat, err := repo.CheckExistsFlat(ctx, expectedFlat.ID)

	// Assert that there were no errors
	require.NoError(t, err)

	// Assert that the returned flat data matches the expected data
	assert.Equal(t, expectedFlat, flat)

	// Assert that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestCheckExistsHouse(t *testing.T) {
	// Initialize your repo, logger, and other dependencies here
	var db *sql.DB
	var mock sqlmock.Sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	logger := zap.Must(zap.NewDevelopment())
	repo := repo.NewRepository(db, logger)

	// Set up a mock context with a user ID
	ctx := context.Background()
	ctx = context.WithValue(ctx, middleware.CookieName, int64(1))
	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())

	// Define expected house data
	expectedHouse := &models.House{
		ID: 1,
	}

	// Mock the QueryRowContext function to return expected data
	rows := sqlmock.NewRows([]string{"id"}).AddRow(expectedHouse.ID)

	query := `SELECT h.id FROM advert AS a JOIN advert_type_house AS at ON a.id=at.advert_id JOIN house AS h ON h.id=at.house_id WHERE a.id = $1;`
	escapedQuery := regexp.QuoteMeta(query)

	mock.ExpectQuery(escapedQuery).
		WithArgs(expectedHouse.ID).
		WillReturnRows(rows)

	// Call the CheckExistsHouse function
	house, err := repo.CheckExistsHouse(ctx, expectedHouse.ID)

	// Assert that there were no errors
	require.NoError(t, err)

	// Assert that the returned house data matches the expected data
	assert.Equal(t, expectedHouse, house)

	// Assert that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteFlatAdvertById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := zap.Must(zap.NewDevelopment())
	rep := repo.NewRepository(db, logger)

	mock.ExpectBegin()

	queryGetIdTables := regexp.QuoteMeta(`
        SELECT
            f.id as flatid
        FROM
            advert AS a
        JOIN
            advert_type_flat AS at ON a.id = at.advert_id
        JOIN
            flat AS f ON f.id = at.flat_id
        WHERE a.id=$1;`)

	mock.ExpectQuery(queryGetIdTables).
		WithArgs(1).WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{"flatid"}).AddRow(1))

	queryDeleteAdvertById := regexp.QuoteMeta(`UPDATE advert SET is_deleted=true WHERE id=$1;`)
	queryDeleteAdvertTypeById := regexp.QuoteMeta(`UPDATE advert_type_flat SET is_deleted=true WHERE advert_id=$1 AND flat_id=$2;`)
	queryDeleteFlatById := regexp.QuoteMeta(`UPDATE flat SET is_deleted=true WHERE id=$1;`)
	queryDeletePriceChanges := regexp.QuoteMeta(`UPDATE price_change SET is_deleted=true WHERE advert_id=$1;`)
	queryDeleteImages := regexp.QuoteMeta(`UPDATE image SET is_deleted=true WHERE advert_id=$1;`)

	mock.ExpectExec(queryDeleteAdvertById).
		WithArgs(1).WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(queryDeleteAdvertTypeById).
		WithArgs(1, 1).WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(queryDeleteFlatById).
		WithArgs(1).WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(queryDeletePriceChanges).
		WithArgs(1).WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(queryDeleteImages).
		WithArgs(1).WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(1, 1))

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}

	err = rep.DeleteFlatAdvertById(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tx, 1)

	if err != nil {
		t.Errorf("Error deleting flat advert by ID: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestDeleteFlatAdvertById_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := zap.Must(zap.NewDevelopment())
	rep := repo.NewRepository(db, logger)

	mock.ExpectBegin()

	queryGetIdTables := regexp.QuoteMeta(`
        SELECT
            f.id as flatid
        FROM
            advert AS a
        JOIN
            advert_type_flat AS at ON a.id = at.advert_id
        JOIN
            flat AS f ON f.id = at.flat_id
        WHERE a.id=$1;`)

	mock.ExpectQuery(queryGetIdTables).
		WithArgs(1).WillReturnError(errors.New("test error"))

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}

	err = rep.DeleteFlatAdvertById(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tx, 1)

	if err == nil {
		t.Error("Expected an error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestDeleteFlatAdvertById_ScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := zap.Must(zap.NewDevelopment())
	rep := repo.NewRepository(db, logger)

	mock.ExpectBegin()

	queryGetIdTables := regexp.QuoteMeta(`
        SELECT
            f.id as flatid
        FROM
            advert AS a
        JOIN
            advert_type_flat AS at ON a.id = at.advert_id
        JOIN
            flat AS f ON f.id = at.flat_id
        WHERE a.id=$1;`)

	mock.ExpectQuery(queryGetIdTables).
		WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"flatid"}).AddRow(nil))

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}

	err = rep.DeleteFlatAdvertById(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tx, 1)

	if err == nil {
		t.Error("Expected an error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestDeleteFlatAdvertById_ExecError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := zap.Must(zap.NewDevelopment())
	rep := repo.NewRepository(db, logger)

	mock.ExpectBegin()

	queryGetIdTables := regexp.QuoteMeta(`
        SELECT
            f.id as flatid
        FROM
            advert AS a
        JOIN
            advert_type_flat AS at ON a.id = at.advert_id
        JOIN
            flat AS f ON f.id = at.flat_id
        WHERE a.id=$1;`)

	mock.ExpectQuery(queryGetIdTables).
		WithArgs(1).WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{"flatid"}).AddRow(1))

	queryDeleteAdvertById := regexp.QuoteMeta(`UPDATE advert SET is_deleted=true WHERE id=$1;`)

	mock.ExpectExec(queryDeleteAdvertById).
		WithArgs(1).WillReturnError(errors.New("test error"))

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}

	err = rep.DeleteFlatAdvertById(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tx, 1)

	if err == nil {
		t.Error("Expected an error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestDeleteHouseAdvertById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := zap.Must(zap.NewDevelopment())
	rep := repo.NewRepository(db, logger)

	mock.ExpectBegin()

	queryGetIdTables := regexp.QuoteMeta(`
        SELECT
            h.id as houseid
        FROM
            advert AS a
        JOIN
            advert_type_house AS at ON a.id = at.advert_id
        JOIN
            house AS h ON h.id = at.house_id
        WHERE a.id=$1;`)

	mock.ExpectQuery(queryGetIdTables).
		WithArgs(1).WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{"houseid"}).AddRow(1))

	queryDeleteAdvertById := regexp.QuoteMeta(`UPDATE advert SET is_deleted=true WHERE id=$1;`)
	queryDeleteAdvertTypeById := regexp.QuoteMeta(`UPDATE advert_type_house SET is_deleted=true WHERE advert_id=$1 AND house_id=$2;`)
	queryDeleteHouseById := regexp.QuoteMeta(`UPDATE house SET is_deleted=true WHERE id=$1;`)
	queryDeletePriceChanges := regexp.QuoteMeta(`UPDATE price_change SET is_deleted=true WHERE advert_id=$1;`)
	queryDeleteImages := regexp.QuoteMeta(`UPDATE image SET is_deleted=true WHERE advert_id=$1;`)

	mock.ExpectExec(queryDeleteAdvertById).
		WithArgs(1).WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(queryDeleteAdvertTypeById).
		WithArgs(1, 1).WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(queryDeleteHouseById).
		WithArgs(1).WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(queryDeletePriceChanges).
		WithArgs(1).WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(queryDeleteImages).
		WithArgs(1).WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(1, 1))

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}

	err = rep.DeleteHouseAdvertById(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tx, 1)

	if err != nil {
		t.Errorf("Error deleting house advert by ID: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestDeleteHouseAdvertById_DatabaseError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := zap.Must(zap.NewDevelopment())
	rep := repo.NewRepository(db, logger)

	mock.ExpectBegin()

	queryGetIdTables := regexp.QuoteMeta(`
        SELECT
            h.id as houseid
        FROM
            advert AS a
        JOIN
            advert_type_house AS at ON a.id = at.advert_id
        JOIN
            house AS h ON h.id = at.house_id
        WHERE a.id=$1;`)

	mock.ExpectQuery(queryGetIdTables).
		WithArgs(1).WillReturnError(fmt.Errorf("some error"))

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}

	err = rep.DeleteHouseAdvertById(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tx, 1)

	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestDeleteHouseAdvertById_NoHouseID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	logger := zap.Must(zap.NewDevelopment())
	rep := repo.NewRepository(db, logger)

	mock.ExpectBegin()

	queryGetIdTables := regexp.QuoteMeta(`
        SELECT
            h.id as houseid
        FROM
            advert AS a
        JOIN
            advert_type_house AS at ON a.id = at.advert_id
        JOIN
            house AS h ON h.id = at.house_id
        WHERE a.id=$1;`)

	mock.ExpectQuery(queryGetIdTables).
		WithArgs(1).WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{"houseid"}))

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}

	err = rep.DeleteHouseAdvertById(context.WithValue(context.Background(), "requestId", uuid.NewV4().String()), tx, 1)

	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
