package repo

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/satori/uuid"
)

// AdvertRepo represents a repository for adverts changes.
type AdvertRepo struct {
	db *sql.DB
}

// NewRepository creates a new instance of AdvertRepo.
func NewRepository(db *sql.DB) *AdvertRepo {
	return &AdvertRepo{db: db}
}

func (r *AdvertRepo) BeginTx(ctx context.Context) (*sql.Tx, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// CreateAdvertType creates a new advertType in the database.
func (r *AdvertRepo) CreateAdvertType(ctx context.Context, tx *sql.Tx, newAdvertType *models.AdvertType) error {
	insert := `INSERT INTO adverttypes (id, adverttype) VALUES ($1, $2)`
	if _, err := tx.ExecContext(ctx, insert, newAdvertType.ID, newAdvertType.AdvertType); err != nil {
		return err
	}
	return nil
}

// CreateAdvert creates a new advert in the database.
func (r *AdvertRepo) CreateAdvert(ctx context.Context, tx *sql.Tx, newAdvert *models.Advert) error {
	insert := `INSERT INTO adverts (id, userid, adverttypeid, adverttypeplacement, title, description, phone, isagent, priority) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	if _, err := tx.ExecContext(ctx, insert, newAdvert.ID, newAdvert.UserID, newAdvert.AdvertTypeID, newAdvert.AdvertTypeSale, newAdvert.Title, newAdvert.Description, newAdvert.Phone, newAdvert.IsAgent, newAdvert.Priority); err != nil {
		return err
	}
	return nil
}

// CreatePriceChange creates a new price change in the database.
func (r *AdvertRepo) CreatePriceChange(ctx context.Context, tx *sql.Tx, newPriceChange *models.PriceChange) error {
	insert := `INSERT INTO pricechanges (id, advertid, price) VALUES ($1, $2, $3)`
	if _, err := tx.ExecContext(ctx, insert, newPriceChange.ID, newPriceChange.AdvertID, newPriceChange.Price); err != nil {
		return err
	}
	return nil
}

// CreateBuilding creates a new building in the database.
func (r *AdvertRepo) CreateBuilding(ctx context.Context, tx *sql.Tx, newBuilding *models.Building) error {
	insert := `INSERT INTO buildings (id, floor, material, adress, adresspoint, yearcreation) VALUES ($1, $2, $3, $4, $5, $6)`
	if _, err := tx.ExecContext(ctx, insert, newBuilding.ID, newBuilding.Floor, newBuilding.Material, newBuilding.Address, newBuilding.AddressPoint, newBuilding.YearCreation); err != nil {
		return err
	}
	return nil
}

// CheckExistsBuilding check exists building.
func (r *AdvertRepo) CheckExistsBuilding(ctx context.Context, adress string) (*models.Building, error) {
	query := `SELECT id FROM buildings WHERE adress = $1`

	building := &models.Building{}

	res := r.db.QueryRowContext(ctx, query, adress)

	if err := res.Scan(&building.ID); err != nil {
		return nil, err
	}

	return building, nil
}

// CreateHouse creates a new house in the database.
func (r *AdvertRepo) CreateHouse(ctx context.Context, tx *sql.Tx, newHouse *models.House) error {
	insert := `INSERT INTO houses (id, buildingid, adverttypeid, ceilingheight, squarearea, squarehouse, bedroomcount, statusarea, cottage, statushome) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	if _, err := tx.ExecContext(ctx, insert, newHouse.ID, newHouse.BuildingID, newHouse.AdvertTypeID, newHouse.CeilingHeight, newHouse.SquareArea, newHouse.SquareHouse, newHouse.BedroomCount, newHouse.StatusArea, newHouse.Cottage, newHouse.StatusHome); err != nil {
		return err
	}
	return nil
}

// CreateFlat creates a new flat in the database.
func (r *AdvertRepo) CreateFlat(ctx context.Context, tx *sql.Tx, newFlat *models.Flat) error {
	insert := `INSERT INTO flats (id, buildingid, adverttypeid, floor, ceilingheight, squaregeneral, roomcount, squareresidential, apartament) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	if _, err := tx.ExecContext(ctx, insert, newFlat.ID, newFlat.BuildingID, newFlat.AdvertTypeID, newFlat.Floor, newFlat.CeilingHeight, newFlat.SquareGeneral, newFlat.RoomCount, newFlat.SquareResidential, newFlat.Apartment); err != nil {
		return err
	}
	return nil
}

// GetHouseSquareAdvertsList retrieves square house adverts from the database.
func (r *AdvertRepo) GetHouseSquareAdvertsList(ctx context.Context) ([]*models.AdvertSquareData, error) {
	query := `
        SELECT
            a.id,
            at.adverttype,
            COALESCE(i.photo, '') AS photo,
            a.adverttypeplacement,
            b.adress,
            h.cottage,
            h.squarehouse,
            h.squarearea,
            h.bedroomcount,
            b.floor,
            pc.price,
            a.datecreation
        FROM
            adverts AS a
        LEFT JOIN
            LATERAL (
                SELECT *
                FROM pricechanges AS pc
                WHERE pc.advertid = a.id
                ORDER BY pc.datecreation DESC
                LIMIT 1
            ) AS pc ON TRUE
        INNER JOIN adverttypes AS at ON a.adverttypeid = at.id
        INNER JOIN houses AS h ON at.id = h.adverttypeid
        INNER JOIN buildings AS b ON h.buildingid = b.id
		JOIN images AS i ON i.advertid = a.id
		WHERE i.priority = (
			SELECT MIN(priority)
			FROM images
			WHERE advertid = a.id
				AND isdeleted = FALSE
			)
			AND i.isdeleted = FALSE AND a.isdeleted = TRUE
		ORDER BY a.datecreation DESC;
    `
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	squareAdverts := []*models.AdvertSquareData{}
	for rows.Next() {
		squareAdvert := &models.AdvertSquareData{}
		var cottage bool
		var squareHouse, squareArea float64
		var bedroomCount, floor int
		err := rows.Scan(&squareAdvert.ID, &squareAdvert.TypeAdvert, &squareAdvert.Photo, &squareAdvert.TypeSale, &squareAdvert.Address, &cottage, &squareHouse, &squareArea, &bedroomCount, &floor, &squareAdvert.Price, &squareAdvert.DateCreation)
		if err != nil {
			return nil, err
		}
		squareAdvert.Properties = make(map[string]interface{})
		squareAdvert.Properties["cottage"] = cottage
		squareAdvert.Properties["squareHouse"] = squareHouse
		squareAdvert.Properties["squareArea"] = squareArea
		squareAdvert.Properties["bedroomCount"] = bedroomCount
		squareAdvert.Properties["floor"] = floor
		squareAdverts = append(squareAdverts, squareAdvert)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return squareAdverts, nil
}

// GetFlatSquareAdvertsList retrieves square flat adverts from the database.
func (r *AdvertRepo) GetFlatSquareAdvertsList(ctx context.Context) ([]*models.AdvertSquareData, error) {
	query := `
        SELECT
            a.id,
            at.adverttype,
            COALESCE(i.photo, '') AS photo,
            a.adverttypeplacement,
            b.adress,
            f.floor,
            f.squaregeneral,
            f.roomcount,
            b.floor AS floorgeneral,
            pc.price,
            a.datecreation
        FROM
            adverts AS a
        LEFT JOIN
            LATERAL (
                SELECT *
                FROM pricechanges AS pc
                WHERE pc.advertid = a.id
                ORDER BY pc.datecreation DESC
                LIMIT 1
            ) AS pc ON TRUE
        INNER JOIN adverttypes AS at ON a.adverttypeid = at.id
        INNER JOIN flats AS f ON at.id = f.adverttypeid
        INNER JOIN buildings AS b ON f.buildingid = b.id
        JOIN images AS i ON i.advertid = a.id
		WHERE i.priority = (
			SELECT MIN(priority)
			FROM images
			WHERE advertid = a.id
				AND isdeleted = FALSE
			)
			AND i.isdeleted = FALSE AND a.isdeleted = FALSE
        ORDER BY a.datecreation DESC;
    `
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	squareAdverts := []*models.AdvertSquareData{}
	for rows.Next() {
		squareAdvert := &models.AdvertSquareData{}
		var floor, floorGeneral, roomCount int
		var squareGeneral float64
		err := rows.Scan(&squareAdvert.ID, &squareAdvert.TypeAdvert, &squareAdvert.Photo, &squareAdvert.TypeSale, &squareAdvert.Address, &floor, &squareGeneral, &roomCount, &floorGeneral, &squareAdvert.Price, &squareAdvert.DateCreation)
		if err != nil {
			return nil, err
		}
		squareAdvert.Properties = make(map[string]interface{})
		squareAdvert.Properties["floor"] = floor
		squareAdvert.Properties["floorGeneral"] = floorGeneral
		squareAdvert.Properties["squareGeneral"] = squareGeneral
		squareAdvert.Properties["roomCount"] = roomCount
		squareAdverts = append(squareAdverts, squareAdvert)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return squareAdverts, nil
}

// GetFlatRectangleAdvertsList retrieves a rectangle Flat adverts from the database.
func (r *AdvertRepo) GetFlatRectangleAdvertsList(ctx context.Context) ([]*models.AdvertRectangleData, error) {
	query := `
        SELECT
            a.id,
            a.title,
            a.description,
            a.phone,
            at.adverttype,
            COALESCE(i.photo, '') as photo,
            a.adverttypeplacement,
            b.adress,
            f.floor,
            f.squaregeneral,
            f.roomcount,
            b.floor AS floorgeneral,
            pc.price,
            a.datecreation
        FROM
            adverts AS a
        LEFT JOIN
            LATERAL (
                SELECT *
                FROM pricechanges AS pc
                WHERE pc.advertid = a.id
                ORDER BY pc.datecreation DESC
                LIMIT 1
            ) AS pc ON true
        INNER JOIN adverttypes AS at ON a.adverttypeid = at.id
        INNER JOIN flats AS f ON at.id = f.adverttypeid
        INNER JOIN buildings AS b ON f.buildingid = b.id
        JOIN images AS i ON i.advertid = a.id
		WHERE i.priority = (
			SELECT MIN(priority)
			FROM images
			WHERE advertid = a.id
				AND isdeleted = FALSE
			)
			AND i.isdeleted = FALSE AND a.isdeleted = FALSE
        ORDER BY a.datecreation DESC;
    `
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rectangleAdverts := []*models.AdvertRectangleData{}
	for rows.Next() {
		rectangleAdvert := &models.AdvertRectangleData{}
		var floor, floorGeneral, roomCount int
		var squareGenereal float64
		err := rows.Scan(
			&rectangleAdvert.ID,
			&rectangleAdvert.Title,
			&rectangleAdvert.Description,
			&rectangleAdvert.Phone,
			&rectangleAdvert.TypeAdvert,
			&rectangleAdvert.Photo,
			&rectangleAdvert.TypeSale,
			&rectangleAdvert.Address,
			&floor,
			&squareGenereal,
			&roomCount,
			&floorGeneral,
			&rectangleAdvert.Price,
			&rectangleAdvert.DateCreation,
		)
		if err != nil {
			return nil, err
		}
		rectangleAdvert.Properties = make(map[string]interface{})
		rectangleAdvert.Properties["floor"] = floor
		rectangleAdvert.Properties["floorGeneral"] = floorGeneral
		rectangleAdvert.Properties["squareGeneral"] = squareGenereal
		rectangleAdvert.Properties["roomCount"] = roomCount
		rectangleAdverts = append(rectangleAdverts, rectangleAdvert)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rectangleAdverts, nil
}

// GetHouseRectangleAdvertsList retrieves a rectangle House adverts from the database.
func (r *AdvertRepo) GetHouseRectangleAdvertsList(ctx context.Context) ([]*models.AdvertRectangleData, error) {
	query := `
        SELECT
            a.id,
            a.title,
            a.description,
            a.phone,
            at.adverttype,
            COALESCE(i.photo, '') as photo,
            a.adverttypeplacement,
            b.adress,
            h.cottage,
            h.squarehouse,
            h.squarearea,
            h.bedroomcount,
            b.floor,
            pc.price,
            a.datecreation
        FROM
            adverts AS a
        LEFT JOIN
            LATERAL (
                SELECT *
                FROM pricechanges AS pc
                WHERE pc.advertid = a.id
                ORDER BY pc.datecreation DESC
                LIMIT 1
            ) AS pc ON true
        INNER JOIN adverttypes AS at ON a.adverttypeid = at.id
        INNER JOIN houses AS h ON at.id = h.adverttypeid
        INNER JOIN buildings AS b ON h.buildingid = b.id
        JOIN images AS i ON i.advertid = a.id
		WHERE i.priority = (
			SELECT MIN(priority)
			FROM images
			WHERE advertid = a.id
				AND isdeleted = FALSE
			)
			AND i.isdeleted = FALSE AND a.isdeleted = FALSE
        ORDER BY a.datecreation DESC;
    `
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rectangleAdverts := []*models.AdvertRectangleData{}
	for rows.Next() {
		rectangleAdvert := &models.AdvertRectangleData{}
		var cottage bool
		var squareHouse, squareArea float64
		var bedroomCount, floor int
		err := rows.Scan(
			&rectangleAdvert.ID,
			&rectangleAdvert.Title,
			&rectangleAdvert.Description,
			&rectangleAdvert.Phone,
			&rectangleAdvert.TypeAdvert,
			&rectangleAdvert.Photo,
			&rectangleAdvert.TypeSale,
			&rectangleAdvert.Address,
			&cottage,
			&squareHouse,
			&squareArea,
			&bedroomCount,
			&floor,
			&rectangleAdvert.Price,
			&rectangleAdvert.DateCreation,
		)
		if err != nil {
			return nil, err
		}
		rectangleAdvert.Properties = make(map[string]interface{})
		rectangleAdvert.Properties["cottage"] = cottage
		rectangleAdvert.Properties["squareHouse"] = squareHouse
		rectangleAdvert.Properties["squareArea"] = squareArea
		rectangleAdvert.Properties["bedroomCount"] = bedroomCount
		rectangleAdvert.Properties["floor"] = floor
		rectangleAdverts = append(rectangleAdverts, rectangleAdvert)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rectangleAdverts, nil
}

// GetTypeAdvertById return type of advert
func (r *AdvertRepo) GetTypeAdvertById(ctx context.Context, id uuid.UUID) (*models.AdvertTypeAdvert, error) {
	query := `SELECT at.adverttype FROM adverts AS a JOIN adverttypes AS at ON a.adverttypeid=at.id WHERE a.id = $1`

	res := r.db.QueryRowContext(ctx, query, id)

	var advertType *models.AdvertTypeAdvert

	if err := res.Scan(&advertType); err != nil {
		return nil, err
	}

	return advertType, nil
}

// GetHouseAdvertById retrieves full information about house advert from the database.
func (r *AdvertRepo) GetHouseAdvertById(ctx context.Context, id uuid.UUID) (*models.AdvertData, error) {
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
    LEFT JOIN
        LATERAL (
            SELECT *
            FROM pricechanges AS pc
            WHERE pc.advertid = a.id
            ORDER BY pc.datecreation DESC
            LIMIT 1
        ) AS pc ON TRUE
    WHERE
        a.id = $1 AND a.isdeleted = FALSE;`
	res := r.db.QueryRowContext(ctx, query, id)

	advertData := &models.AdvertData{}
	var cottage bool
	var squareHouse, squareArea, ceilingHeight float64
	var bedroomCount, floor int
	var statusArea models.StatusAreaHouse
	var statusHome models.StatusHomeHouse
	var complexId, companyPhoto, companyName, complexName sql.NullString

	if err := res.Scan(
		&advertData.ID,
		&advertData.TypeAdvert,
		&advertData.TypeSale,
		&advertData.Title,
		&advertData.Description,
		&advertData.Price,
		&advertData.Phone,
		&advertData.IsAgent,
		&advertData.Address,
		&advertData.AddressPoint,
		&ceilingHeight,
		&squareArea,
		&squareHouse,
		&bedroomCount,
		&statusArea,
		&cottage,
		&statusHome,
		&floor,
		&advertData.YearCreation,
		&advertData.Material,
		&advertData.DateCreation,
		&complexId,
		&companyPhoto,
		&companyName,
		&complexName,
	); err != nil {
		return nil, err
	}

	advertData.Properties = make(map[string]interface{})
	advertData.Properties["ceilingHeight"] = ceilingHeight
	advertData.Properties["squareArea"] = squareArea
	advertData.Properties["squareHouse"] = squareHouse
	advertData.Properties["bedroomCount"] = bedroomCount
	advertData.Properties["statusArea"] = statusArea
	advertData.Properties["cottage"] = cottage
	advertData.Properties["statusHome"] = statusHome
	advertData.Properties["floor"] = floor

	advertData.Complex = make(map[string]interface{})
	advertData.Complex["complexId"] = complexId
	advertData.Complex["companyPhoto"] = companyPhoto
	advertData.Complex["companyName"] = companyName
	advertData.Complex["complexName"] = complexName

	return advertData, nil
}

// CheckExistsFlat check exists flat.
func (r *AdvertRepo) CheckExistsFlat(ctx context.Context, advertId uuid.UUID) (*models.Flat, error) {
	query := `SELECT f.id FROM adverts AS a JOIN adverttypes AS at ON a.adverttypeid=at.id JOIN flats AS f ON f.adverttypeid=at.id WHERE a.id = $1`

	flat := &models.Flat{}

	res := r.db.QueryRowContext(ctx, query, advertId)

	if err := res.Scan(&flat.ID); err != nil {
		return nil, err
	}

	return flat, nil
}

// CheckExistsHouse check exists flat.
func (r *AdvertRepo) CheckExistsHouse(ctx context.Context, advertId uuid.UUID) (*models.House, error) {
	query := `SELECT h.id FROM adverts AS a JOIN adverttypes AS at ON a.adverttypeid=at.id JOIN houses AS h ON h.adverttypeid=at.id WHERE a.id = $1;`

	house := &models.House{}

	res := r.db.QueryRowContext(ctx, query, advertId)

	if err := res.Scan(&house.ID); err != nil {
		return nil, err
	}

	return house, nil
}

// DeleteFlatAdvertById
func (r *AdvertRepo) DeleteFlatAdvertById(ctx context.Context, tx *sql.Tx, advertId uuid.UUID) error {
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
	res := tx.QueryRowContext(ctx, queryGetIdTables, advertId)

	var advertTypeId, flatId uuid.UUID
	if err := res.Scan(&advertTypeId, &flatId); err != nil {
		return err
	}

	queryDeleteAdvertById := `UPDATE adverts SET isdeleted=true WHERE id=$1;`
	queryDeleteAdvertTypeById := `UPDATE adverttypes SET isdeleted=true WHERE id=$1;`
	queryDeleteFlatById := `UPDATE flats SET isdeleted=true WHERE id=$1;`
	queryDeletePriceChanges := `UPDATE pricechanges SET isdeleted=true WHERE advertid=$1;`
	queryDeleteImages := `UPDATE images SET isdeleted=true WHERE advertid=$1;`

	if _, err := tx.Exec(queryDeleteAdvertById, advertId); err != nil {
		return err
	}
	if _, err := tx.Exec(queryDeleteAdvertTypeById, advertTypeId); err != nil {
		return err
	}
	if _, err := tx.Exec(queryDeleteFlatById, flatId); err != nil {
		return err
	}
	if _, err := tx.Exec(queryDeletePriceChanges, advertId); err != nil {
		return err
	}
	if _, err := tx.Exec(queryDeleteImages, advertId); err != nil {
		return err
	}

	return nil
}

// DeleteHouseAdvertById
func (r *AdvertRepo) DeleteHouseAdvertById(ctx context.Context, tx *sql.Tx, advertId uuid.UUID) error {
	queryGetIdTables := `
	SELECT
	at.id as adverttypeid,
	h.id as houseid
FROM
	adverts AS a
JOIN
	adverttypes AS at ON a.adverttypeid = at.id
JOIN
	houses AS h ON h.adverttypeid = at.id
	WHERE a.id=$1;`
	res := tx.QueryRowContext(ctx, queryGetIdTables, advertId)

	var advertTypeId, houseId uuid.UUID
	if err := res.Scan(&advertTypeId, &houseId); err != nil {
		return err
	}

	queryDeleteAdvertById := `UPDATE adverts SET isdeleted=true WHERE id=$1;`
	queryDeleteAdvertTypeById := `UPDATE adverttypes SET isdeleted=true WHERE id=$1;`
	queryDeleteHouseById := `UPDATE houses SET isdeleted=true WHERE id=$1;`
	queryDeletePriceChanges := `UPDATE pricechanges SET isdeleted=true WHERE advertid=$1;`
	queryDeleteImages := `UPDATE images SET isdeleted=true WHERE advertid=$1;`

	if _, err := tx.Exec(queryDeleteAdvertById, advertId); err != nil {
		return err
	}
	if _, err := tx.Exec(queryDeleteAdvertTypeById, advertTypeId); err != nil {
		return err
	}
	if _, err := tx.Exec(queryDeleteHouseById, houseId); err != nil {
		return err
	}
	if _, err := tx.Exec(queryDeletePriceChanges, advertId); err != nil {
		return err
	}
	if _, err := tx.Exec(queryDeleteImages, advertId); err != nil {
		return err
	}

	return nil
}

// ChangeTypeAdvert
func (r *AdvertRepo) ChangeTypeAdvert(ctx context.Context, tx *sql.Tx, advertId uuid.UUID) error {
	query := `SELECT at.id, at.adverttype FROM adverts AS a JOIN adverttypes AS at ON a.adverttypeid=at.id WHERE a.id = $1;`
	querySelectBuildingIdByFlat := `SELECT b.id FROM adverts AS a JOIN adverttypes AS at ON at.id=a.adverttypeid JOIN flats AS f ON f.adverttypeid=at.id JOIN buildings AS b ON f.buildingid=b.id WHERE a.id=$1`
	querySelectBuildingIdByHouse := `SELECT b.id FROM adverts AS a JOIN adverttypes AS at ON at.id=a.adverttypeid JOIN houses AS h ON h.adverttypeid=at.id JOIN buildings AS b ON h.buildingid=b.id WHERE a.id=$1`
	queryInsertFlat := `INSERT INTO flats (id, buildingId, advertTypeId, floor, ceilingHeight, squareGeneral, roomCount, squareResidential, apartament)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`
	queryInsertHouse := `INSERT INTO houses (id, buildingId, advertTypeId, ceilingHeight, squareArea, squareHouse, bedroomCount, statusArea, cottage, statusHome)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`

	var advertType models.AdvertTypeAdvert
	var advertTypeId uuid.UUID
	res := r.db.QueryRowContext(ctx, query, advertId)

	if err := res.Scan(&advertTypeId, &advertType); err != nil {
		return err
	}
	var buildingId uuid.UUID
	switch advertType {
	case models.AdvertTypeFlat:
		if _, err := r.CheckExistsHouse(ctx, advertId); err != nil {

			res := r.db.QueryRowContext(ctx, querySelectBuildingIdByFlat, advertId)

			if err := res.Scan(&buildingId); err != nil {
				return err
			}

			house := &models.House{}
			if _, err := tx.Exec(queryInsertHouse, uuid.NewV4(), buildingId, advertTypeId, house.CeilingHeight, house.SquareArea, house.SquareHouse, house.BedroomCount, house.StatusArea, house.Cottage, house.StatusHome); err != nil {
				return err
			}
		}
	case models.AdvertTypeHouse:
		if _, err := r.CheckExistsFlat(ctx, advertId); err != nil {
			res := r.db.QueryRowContext(ctx, querySelectBuildingIdByHouse, advertId)

			if err := res.Scan(&buildingId); err != nil {
				return err
			}

			flat := &models.Flat{}
			if _, err := tx.Exec(queryInsertFlat, uuid.NewV4(), buildingId, advertTypeId, flat.Floor, flat.CeilingHeight, flat.SquareGeneral, flat.RoomCount, flat.SquareResidential, flat.Apartment); err != nil {
				return err
			}
		}
	}

	return nil
}

// UpdateHouseAdvertById update house advert from the database.
func (r *AdvertRepo) UpdateHouseAdvertById(ctx context.Context, tx *sql.Tx, advertUpdateData *models.AdvertUpdateData) error {
	queryGetIdTables := `
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
	WHERE a.id=$1;`
	res := tx.QueryRowContext(ctx, queryGetIdTables, advertUpdateData.ID)

	var advertTypeId, buildingId, houseId uuid.UUID
	var price int64
	if err := res.Scan(&advertTypeId, &buildingId, &houseId, &price); err != nil {
		return err
	}

	queryUpdateAdvertById := `UPDATE adverts SET adverttypeplacement=$1, title=$2, description=$3, phone=$4, isagent=$5 WHERE id=$6;`
	queryUpdateAdvertTypeById := `UPDATE adverttypes SET adverttype=$1 WHERE id=$2;`
	queryUpdateBuildingById := `UPDATE buildings SET floor=$1, material=$2, adress=$3, adresspoint=$4, yearcreation=$5 WHERE id=$6;`
	queryUpdateHouseById := `UPDATE houses SET ceilingheight=$1, squarearea=$2, squarehouse=$3, bedroomcount=$4, statusarea=$5, cottage=$6, statushome=$7 WHERE id=$8;`

	if _, err := tx.Exec(queryUpdateAdvertById, advertUpdateData.TypeSale, advertUpdateData.Title, advertUpdateData.Description, advertUpdateData.Phone, advertUpdateData.IsAgent, advertUpdateData.ID); err != nil {
		return err
	}
	if _, err := tx.Exec(queryUpdateAdvertTypeById, advertUpdateData.TypeAdvert, advertTypeId); err != nil {
		return err
	}
	if _, err := tx.Exec(queryUpdateBuildingById, advertUpdateData.Properties["floor"], advertUpdateData.Material, advertUpdateData.Address, advertUpdateData.AddressPoint, advertUpdateData.YearCreation, buildingId); err != nil {
		return err
	}
	if _, err := tx.Exec(queryUpdateHouseById, advertUpdateData.Properties["ceilingHeight"], advertUpdateData.Properties["squareArea"], advertUpdateData.Properties["squareHouse"], advertUpdateData.Properties["bedroomCount"], advertUpdateData.Properties["statusArea"], advertUpdateData.Properties["cottage"], advertUpdateData.Properties["statusHome"], houseId); err != nil {
		return err
	}
	if advertUpdateData.Price != price {
		queryInsertPriceChange := `INSERT INTO pricechanges (id, advertId, price)
		VALUES ($1, $2, $3)`
		if _, err := tx.Exec(queryInsertPriceChange,
			uuid.NewV4(), advertUpdateData.ID, advertUpdateData.Price); err != nil {
			return err
		}
	}

	return nil
}

// UpdateFlatAdvertById update flat advert from the database.
func (r *AdvertRepo) UpdateFlatAdvertById(ctx context.Context, tx *sql.Tx, advertUpdateData *models.AdvertUpdateData) error {
	queryGetIdTables := `
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
	WHERE a.id=$1;`

	res := tx.QueryRowContext(ctx, queryGetIdTables, advertUpdateData.ID)

	var advertTypeId, buildingId, flatId uuid.UUID
	var price int64
	if err := res.Scan(&advertTypeId, &buildingId, &flatId, &price); err != nil {
		return err
	}

	queryUpdateAdvertById := `UPDATE adverts SET adverttypeplacement=$1, title=$2, description=$3, phone=$4, isagent=$5 WHERE id=$6;`
	queryUpdateAdvertTypeById := `UPDATE adverttypes SET adverttype=$1 WHERE id=$2;`
	queryUpdateBuildingById := `UPDATE buildings SET floor=$1, material=$2, adress=$3, adresspoint=$4, yearcreation=$5 WHERE id=$6;`
	queryUpdateFlatById := `UPDATE flats SET floor=$1, ceilingheight=$2, squaregeneral=$3, roomcount=$4, squareresidential=$5, apartament=$6 WHERE id=$7;`

	if _, err := tx.Exec(queryUpdateAdvertById, advertUpdateData.TypeSale, advertUpdateData.Title, advertUpdateData.Description, advertUpdateData.Phone, advertUpdateData.IsAgent, advertUpdateData.ID); err != nil {
		return err
	}
	if _, err := tx.Exec(queryUpdateAdvertTypeById, advertUpdateData.TypeAdvert, advertTypeId); err != nil {
		return err
	}
	if _, err := tx.Exec(queryUpdateBuildingById, advertUpdateData.Properties["floorGeneral"], advertUpdateData.Material, advertUpdateData.Address, advertUpdateData.AddressPoint, advertUpdateData.YearCreation, buildingId); err != nil {
		return err
	}
	if _, err := tx.Exec(queryUpdateFlatById, advertUpdateData.Properties["floor"], advertUpdateData.Properties["ceilingHeight"], advertUpdateData.Properties["squareGeneral"], advertUpdateData.Properties["roomCount"], advertUpdateData.Properties["squareResidential"], advertUpdateData.Properties["apartament"], flatId); err != nil {
		return err
	}
	if advertUpdateData.Price != price {
		queryInsertPriceChange := `INSERT INTO pricechanges (id, advertId, price)
		VALUES ($1, $2, $3)`
		if _, err := tx.Exec(queryInsertPriceChange,
			uuid.NewV4(), advertUpdateData.ID, advertUpdateData.Price); err != nil {
			return err
		}
	}

	return nil
}

// GetFlatAdvertById retrieves full information about flat advert from the database.
func (r *AdvertRepo) GetFlatAdvertById(ctx context.Context, id uuid.UUID) (*models.AdvertData, error) {
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
        a.id = $1 AND a.isdeleted = FALSE;`
	res := r.db.QueryRowContext(ctx, query, id)

	advertData := &models.AdvertData{}
	var floor, floorGeneral, roomCount int
	var squareGenereal, squareResidential, ceilingHeight float64
	var apartament sql.NullBool
	var complexId, companyPhoto, companyName, complexName sql.NullString

	if err := res.Scan(
		&advertData.ID,
		&advertData.TypeAdvert,
		&advertData.TypeSale,
		&advertData.Title,
		&advertData.Description,
		&advertData.Price,
		&advertData.Phone,
		&advertData.IsAgent,
		&advertData.Address,
		&advertData.AddressPoint,
		&floor,
		&ceilingHeight,
		&squareGenereal,
		&roomCount,
		&squareResidential,
		&apartament,
		&floorGeneral,
		&advertData.YearCreation,
		&advertData.Material,
		&advertData.DateCreation,
		&complexId,
		&companyPhoto,
		&companyName,
		&complexName,
	); err != nil {
		return nil, err
	}

	advertData.Properties = make(map[string]interface{})
	advertData.Properties["ceilingHeight"] = ceilingHeight
	advertData.Properties["apartament"] = apartament
	advertData.Properties["squareResidential"] = squareResidential
	advertData.Properties["roomCount"] = roomCount
	advertData.Properties["squareGeneral"] = squareGenereal
	advertData.Properties["floorGeneral"] = floorGeneral
	advertData.Properties["floor"] = floor

	advertData.Complex = make(map[string]interface{})
	advertData.Complex["complexId"] = complexId
	advertData.Complex["companyPhoto"] = companyPhoto
	advertData.Complex["companyName"] = companyName
	advertData.Complex["complexName"] = complexName

	return advertData, nil
}

// GetSquareAdverts retrieves square adverts from the database.
func (r *AdvertRepo) GetSquareAdverts(ctx context.Context, pageSize, offset int) ([]*models.AdvertSquareData, error) {
	query := `
	SELECT
    a.id,
    at.adverttype,
    a.adverttypeplacement,
	i.photo,
    pc.price,
    a.datecreation
FROM
    adverts AS a
    JOIN adverttypes AS at ON a.adverttypeid = at.id
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
	ORDER BY
    a.datecreation DESC
LIMIT $1
OFFSET $2;`
	queryFlat := `
	SELECT 
	f.squaregeneral,
	 f.floor,
 b.adress,
	 b.floor AS floorgeneral,
	 f.roomcount
 FROM
	 adverts AS a
	 JOIN adverttypes AS at ON a.adverttypeid = at.id
	 JOIN flats AS f ON f.adverttypeid=at.id
 JOIN buildings AS b ON f.buildingid=b.id
 WHERE a.id=$1 AND a.isdeleted = FALSE
 ORDER BY
	 a.datecreation DESC;`
	queryHouse := `
	SELECT 
        b.adress,
        h.cottage,
        h.squarehouse,
        h.squarearea,
        h.bedroomcount,
        b.floor
 FROM
         adverts AS a
         JOIN adverttypes AS at ON a.adverttypeid = at.id
         JOIN houses AS h ON h.adverttypeid=at.id
 JOIN buildings AS b ON h.buildingid=b.id
 WHERE a.id=$1
 ORDER BY
         a.datecreation DESC;`
	rows, err := r.db.Query(query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	squareAdverts := []*models.AdvertSquareData{}
	for rows.Next() {
		squareAdvert := &models.AdvertSquareData{}
		err := rows.Scan(&squareAdvert.ID, &squareAdvert.TypeAdvert, &squareAdvert.TypeSale, &squareAdvert.Photo, &squareAdvert.Price, &squareAdvert.DateCreation)
		if err != nil {
			return nil, err
		}
		squareAdvert.Properties = make(map[string]interface{})
		switch squareAdvert.TypeAdvert {
		case string(models.AdvertTypeFlat):
			var squareGeneral float64
			var floor, floorGeneral, roomCount int
			row := r.db.QueryRowContext(ctx, queryFlat, squareAdvert.ID)
			if err := row.Scan(&squareGeneral, &floor, &squareAdvert.Address, &floorGeneral, &roomCount); err != nil {
				return nil, err
			}
			squareAdvert.Properties["floor"] = floor
			squareAdvert.Properties["floorGeneral"] = floorGeneral
			squareAdvert.Properties["squareGeneral"] = squareGeneral
			squareAdvert.Properties["roomCount"] = roomCount
		case string(models.AdvertTypeHouse):
			var cottage bool
			var squareHouse, squareArea float64
			var bedroomCount, floor int
			row := r.db.QueryRowContext(ctx, queryHouse, squareAdvert.ID)
			if err := row.Scan(&squareAdvert.Address, &cottage, &squareHouse, &squareArea, &bedroomCount, &floor); err != nil {
				return nil, err
			}
			squareAdvert.Properties["cottage"] = cottage
			squareAdvert.Properties["squareHouse"] = squareHouse
			squareAdvert.Properties["squareArea"] = squareArea
			squareAdvert.Properties["bedroomCount"] = bedroomCount
			squareAdvert.Properties["floor"] = floor
		}

		squareAdverts = append(squareAdverts, squareAdvert)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return squareAdverts, nil
}

// GetRectangleAdverts retrieves rectangle adverts from the database with search.
func (r *AdvertRepo) GetRectangleAdverts(ctx context.Context, advertFilter models.AdvertFilter) (*models.AdvertDataPage, error) {
	query := `
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
JOIN
    adverttypes AS at ON a.adverttypeid = at.id
LEFT JOIN
    flats AS f ON f.adverttypeid = at.id
LEFT JOIN
    houses AS h ON h.adverttypeid = at.id
LEFT JOIN
    buildings AS b ON (f.buildingid = b.id OR h.buildingid = b.id)
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
AND i.isdeleted = FALSE AND a.isdeleted = FALSE AND pc.price>=$1 AND pc.price<=$2 AND b.adress ILIKE $3 `
	queryFlat := `
	SELECT 
	f.squaregeneral,
	 f.floor,
 b.adress,
	 b.floor AS floorgeneral
 FROM
	 adverts AS a
	 JOIN adverttypes AS at ON a.adverttypeid = at.id
	 JOIN flats AS f ON f.adverttypeid=at.id
 JOIN buildings AS b ON f.buildingid=b.id
 WHERE a.id=$1
 ORDER BY
	 a.datecreation DESC;`
	queryHouse := `
	SELECT 
        b.adress,
        h.cottage,
        h.squarehouse,
        h.squarearea,
        b.floor
 FROM
         adverts AS a
         JOIN adverttypes AS at ON a.adverttypeid = at.id
         JOIN houses AS h ON h.adverttypeid=at.id
 JOIN buildings AS b ON h.buildingid=b.id
 WHERE a.id=$1
 ORDER BY
         a.datecreation DESC;`

	pageInfo := &models.PageInfo{}
	var args []interface{}
	i := 4
	advertFilter.Address = "%" + advertFilter.Address + "%"
	if advertFilter.AdvertType != "" {
		query += "AND at.adverttype=$" + fmt.Sprint(i) + " "
		args = append(args, advertFilter.AdvertType)
		i++
	}

	if advertFilter.DealType != "" {
		query += "AND a.adverttypeplacement=$" + fmt.Sprint(i) + " "
		args = append(args, advertFilter.DealType)
		i++
	}
	if advertFilter.RoomCount != 0 {
		query = "SELECT * FROM (" + query + ") AS bobik WHERE rcount=$" + fmt.Sprint(i) + " "
		args = append(args, advertFilter.RoomCount)
		i++
	}
	queryCount := "SELECT COUNT(*) FROM (" + query + ") AS bibik;"
	query += "ORDER BY datecreation DESC LIMIT $" + fmt.Sprint(i) + "OFFSET $" + fmt.Sprint(i+1) + ";"
	rowCountQuery := r.db.QueryRowContext(ctx, queryCount, append([]interface{}{advertFilter.MinPrice, advertFilter.MaxPrice, advertFilter.Address}, args...)...)

	if err := rowCountQuery.Scan(&pageInfo.TotalElements); err != nil {
		return nil, err
	}

	args = append(args, advertFilter.Page, advertFilter.Offset)
	rows, err := r.db.Query(query, append([]interface{}{advertFilter.MinPrice, advertFilter.MaxPrice, advertFilter.Address}, args...)...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rectangleAdverts := []*models.AdvertRectangleData{}
	for rows.Next() {
		var roomCount int
		rectangleAdvert := &models.AdvertRectangleData{}
		err := rows.Scan(&rectangleAdvert.ID, &rectangleAdvert.Title, &rectangleAdvert.Description, &rectangleAdvert.TypeAdvert, &roomCount, &rectangleAdvert.Phone, &rectangleAdvert.TypeSale, &rectangleAdvert.Address, &rectangleAdvert.Price, &rectangleAdvert.Photo, &rectangleAdvert.DateCreation)
		if err != nil {
			return nil, err
		}
		rectangleAdvert.Properties = make(map[string]interface{})
		switch rectangleAdvert.TypeAdvert {
		case string(models.AdvertTypeFlat):
			var squareGeneral float64
			var floor, floorGeneral int
			row := r.db.QueryRowContext(ctx, queryFlat, rectangleAdvert.ID)
			if err := row.Scan(&squareGeneral, &floor, &rectangleAdvert.Address, &floorGeneral); err != nil {
				return nil, err
			}
			rectangleAdvert.Properties["floor"] = floor
			rectangleAdvert.Properties["floorGeneral"] = floorGeneral
			rectangleAdvert.Properties["squareGeneral"] = squareGeneral
			rectangleAdvert.Properties["roomCount"] = roomCount
		case string(models.AdvertTypeHouse):
			var cottage bool
			var squareHouse, squareArea float64
			var floor int
			row := r.db.QueryRowContext(ctx, queryHouse, rectangleAdvert.ID)
			if err := row.Scan(&rectangleAdvert.Address, &cottage, &squareHouse, &squareArea, &floor); err != nil {
				return nil, err
			}
			rectangleAdvert.Properties["cottage"] = cottage
			rectangleAdvert.Properties["squareHouse"] = squareHouse
			rectangleAdvert.Properties["squareArea"] = squareArea
			rectangleAdvert.Properties["bedroomCount"] = roomCount
			rectangleAdvert.Properties["floor"] = floor
		}

		rectangleAdverts = append(rectangleAdverts, rectangleAdvert)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	pageInfo.PageSize = advertFilter.Page
	pageInfo.TotalPages = pageInfo.TotalElements / pageInfo.PageSize
	if pageInfo.TotalElements%pageInfo.PageSize != 0 {
		pageInfo.TotalPages++
	}
	pageInfo.CurrentPage = (advertFilter.Offset / pageInfo.PageSize) + 1

	return &models.AdvertDataPage{rectangleAdverts, pageInfo}, nil
}

// GetRectangleAdvertsByUserId retrieves rectangle adverts from the database by user id.
func (r *AdvertRepo) GetRectangleAdvertsByUserId(ctx context.Context, pageSize, offset int, userId uuid.UUID) ([]*models.AdvertRectangleData, error) {
	query := `
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
JOIN
    adverttypes AS at ON a.adverttypeid = at.id
LEFT JOIN
    flats AS f ON f.adverttypeid = at.id
LEFT JOIN
    houses AS h ON h.adverttypeid = at.id
LEFT JOIN
    buildings AS b ON (f.buildingid = b.id OR h.buildingid = b.id)
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
AND i.isdeleted = FALSE AND a.isdeleted = FALSE AND userid=$1 ORDER BY datecreation DESC LIMIT $2 OFFSET $3;`
	queryFlat := `
	SELECT 
	f.squaregeneral,
	 f.floor,
 b.adress,
	 b.floor AS floorgeneral
 FROM
	 adverts AS a
	 JOIN adverttypes AS at ON a.adverttypeid = at.id
	 JOIN flats AS f ON f.adverttypeid=at.id
 JOIN buildings AS b ON f.buildingid=b.id
 WHERE a.id=$1
 ORDER BY
	 a.datecreation DESC;`
	queryHouse := `
	SELECT 
        b.adress,
        h.cottage,
        h.squarehouse,
        h.squarearea,
        b.floor
 FROM
         adverts AS a
         JOIN adverttypes AS at ON a.adverttypeid = at.id
         JOIN houses AS h ON h.adverttypeid=at.id
 JOIN buildings AS b ON h.buildingid=b.id
 WHERE a.id=$1
 ORDER BY
         a.datecreation DESC;`

	rows, err := r.db.Query(query, userId, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rectangleAdverts := []*models.AdvertRectangleData{}
	for rows.Next() {
		var roomCount int
		rectangleAdvert := &models.AdvertRectangleData{}
		err := rows.Scan(&rectangleAdvert.ID, &rectangleAdvert.Title, &rectangleAdvert.Description, &rectangleAdvert.TypeAdvert, &roomCount, &rectangleAdvert.Phone, &rectangleAdvert.TypeSale, &rectangleAdvert.Address, &rectangleAdvert.Price, &rectangleAdvert.Photo, &rectangleAdvert.DateCreation)
		if err != nil {
			return nil, err
		}
		rectangleAdvert.Properties = make(map[string]interface{})
		switch rectangleAdvert.TypeAdvert {
		case string(models.AdvertTypeFlat):
			var squareGeneral float64
			var floor, floorGeneral int
			row := r.db.QueryRowContext(ctx, queryFlat, rectangleAdvert.ID)
			if err := row.Scan(&squareGeneral, &floor, &rectangleAdvert.Address, &floorGeneral); err != nil {
				return nil, err
			}
			rectangleAdvert.Properties["floor"] = floor
			rectangleAdvert.Properties["floorGeneral"] = floorGeneral
			rectangleAdvert.Properties["squareGeneral"] = squareGeneral
			rectangleAdvert.Properties["roomCount"] = roomCount
		case string(models.AdvertTypeHouse):
			var cottage bool
			var squareHouse, squareArea float64
			var floor int
			row := r.db.QueryRowContext(ctx, queryHouse, rectangleAdvert.ID)
			if err := row.Scan(&rectangleAdvert.Address, &cottage, &squareHouse, &squareArea, &floor); err != nil {
				return nil, err
			}
			rectangleAdvert.Properties["cottage"] = cottage
			rectangleAdvert.Properties["squareHouse"] = squareHouse
			rectangleAdvert.Properties["squareArea"] = squareArea
			rectangleAdvert.Properties["bedroomCount"] = roomCount
			rectangleAdvert.Properties["floor"] = floor
		}

		rectangleAdverts = append(rectangleAdverts, rectangleAdvert)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rectangleAdverts, nil
}

// GetRectangleAdvertsByComplexId retrieves rectangle adverts from the database by complex id.
func (r *AdvertRepo) GetRectangleAdvertsByComplexId(ctx context.Context, pageSize, offset int, complexId uuid.UUID) ([]*models.AdvertRectangleData, error) {
	query := `
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
JOIN
    adverttypes AS at ON a.adverttypeid = at.id
LEFT JOIN
    flats AS f ON f.adverttypeid = at.id
LEFT JOIN
    houses AS h ON h.adverttypeid = at.id
LEFT JOIN
    buildings AS b ON (f.buildingid = b.id OR h.buildingid = b.id)
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
AND i.isdeleted = FALSE AND a.isdeleted = FALSE AND b.complexid=$1 ORDER BY datecreation DESC LIMIT $2 OFFSET $3;`
	queryFlat := `
	SELECT 
	f.squaregeneral,
	 f.floor,
 b.adress,
	 b.floor AS floorgeneral
 FROM
	 adverts AS a
	 JOIN adverttypes AS at ON a.adverttypeid = at.id
	 JOIN flats AS f ON f.adverttypeid=at.id
 JOIN buildings AS b ON f.buildingid=b.id
 WHERE a.id=$1
 ORDER BY
	 a.datecreation DESC;`
	queryHouse := `
	SELECT 
        b.adress,
        h.cottage,
        h.squarehouse,
        h.squarearea,
        b.floor
 FROM
         adverts AS a
         JOIN adverttypes AS at ON a.adverttypeid = at.id
         JOIN houses AS h ON h.adverttypeid=at.id
 JOIN buildings AS b ON h.buildingid=b.id
 WHERE a.id=$1
 ORDER BY
         a.datecreation DESC;`

	rows, err := r.db.Query(query, complexId, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rectangleAdverts := []*models.AdvertRectangleData{}
	for rows.Next() {
		var roomCount int
		rectangleAdvert := &models.AdvertRectangleData{}
		err := rows.Scan(&rectangleAdvert.ID, &rectangleAdvert.Title, &rectangleAdvert.Description, &rectangleAdvert.TypeAdvert, &roomCount, &rectangleAdvert.Phone, &rectangleAdvert.TypeSale, &rectangleAdvert.Address, &rectangleAdvert.Price, &rectangleAdvert.Photo, &rectangleAdvert.DateCreation)
		if err != nil {
			return nil, err
		}
		rectangleAdvert.Properties = make(map[string]interface{})
		switch rectangleAdvert.TypeAdvert {
		case string(models.AdvertTypeFlat):
			var squareGeneral float64
			var floor, floorGeneral int
			row := r.db.QueryRowContext(ctx, queryFlat, rectangleAdvert.ID)
			if err := row.Scan(&squareGeneral, &floor, &rectangleAdvert.Address, &floorGeneral); err != nil {
				return nil, err
			}
			rectangleAdvert.Properties["floor"] = floor
			rectangleAdvert.Properties["floorGeneral"] = floorGeneral
			rectangleAdvert.Properties["squareGeneral"] = squareGeneral
			rectangleAdvert.Properties["roomCount"] = roomCount
		case string(models.AdvertTypeHouse):
			var cottage bool
			var squareHouse, squareArea float64
			var floor int
			row := r.db.QueryRowContext(ctx, queryHouse, rectangleAdvert.ID)
			if err := row.Scan(&rectangleAdvert.Address, &cottage, &squareHouse, &squareArea, &floor); err != nil {
				return nil, err
			}
			rectangleAdvert.Properties["cottage"] = cottage
			rectangleAdvert.Properties["squareHouse"] = squareHouse
			rectangleAdvert.Properties["squareArea"] = squareArea
			rectangleAdvert.Properties["bedroomCount"] = roomCount
			rectangleAdvert.Properties["floor"] = floor
		}

		rectangleAdverts = append(rectangleAdverts, rectangleAdvert)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rectangleAdverts, nil
}
