package repo

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"

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
	var building *models.Building
	selectResp := `SELECT id, floor, material, adress, adresspoint, yearcreation FROM buildings WHERE adress = $1`

	if err := r.db.QueryRowContext(ctx, selectResp, adress).Scan(building); err != nil {
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
        LEFT JOIN images AS i ON a.id = i.advertid
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
        LEFT JOIN images AS i ON a.id = i.advertid
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
        LEFT JOIN images AS i ON a.id = i.advertid
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
        LEFT JOIN images AS i ON a.id = i.advertid
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
    a.id = $1;`
	res := r.db.QueryRowContext(ctx, query, id)

	advertData := &models.AdvertData{}
	var cottage bool
	var squareHouse, squareArea float64
	var ceilingheight, bedroomCount, floor int
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
		&ceilingheight,
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
	advertData.Properties["ceilingHeight"] = ceilingheight
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
