package repo

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"
)

// AdvertRepo represents a repository for adverts changes.
type AdvertRepo struct {
	db *sql.DB
}

// NewRepository creates a new instance of AdvertRepo.
func NewRepository(db *sql.DB) *AdvertRepo {
	return &AdvertRepo{db: db}
}

// CreateAdvertType creates a new advertType in the database.
func (r *AdvertRepo) CreateAdvertType(ctx context.Context, newAdvertType *models.AdvertType) error {
	insert := `INSERT INTO adverttypes (id, adverttype) VALUES ($1, $2)`
	if _, err := r.db.ExecContext(ctx, insert, newAdvertType.ID, newAdvertType.AdvertType); err != nil {
		return err
	}
	return nil
}

// CreateAdvert creates a new advert in the database.
func (r *AdvertRepo) CreateAdvert(ctx context.Context, newAdvert *models.Advert) error {
	insert := `INSERT INTO adverts (id, userid, adverttypeid, adverttypeplacement, title, description, phone, isagent, priority) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	if _, err := r.db.ExecContext(ctx, insert, newAdvert.ID, newAdvert.UserID, newAdvert.AdvertTypeID, newAdvert.AdvertTypeSale, newAdvert.Title, newAdvert.Description, newAdvert.Phone, newAdvert.IsAgent, newAdvert.Priority); err != nil {
		return err
	}
	return nil
}

// CreatePriceChange creates a new price change in the database.
func (r *AdvertRepo) CreatePriceChange(ctx context.Context, newPriceChange *models.PriceChange) error {
	insert := `INSERT INTO pricechanges (id, advertid, price) VALUES ($1, $2, $3)`
	if _, err := r.db.ExecContext(ctx, insert, newPriceChange.ID, newPriceChange.AdvertID, newPriceChange.Price); err != nil {
		return err
	}
	return nil
}

// CreateBuilding creates a new building in the database.
func (r *AdvertRepo) CreateBuilding(ctx context.Context, newBuilding *models.Building) error {
	insert := `INSERT INTO buildings (id, floor, material, adress, adresspoint, yearcreation) VALUES ($1, $2, $3, $4, $5, $6)`
	if _, err := r.db.ExecContext(ctx, insert, newBuilding.ID, newBuilding.Floor, newBuilding.Material, newBuilding.Address, newBuilding.AddressPoint, newBuilding.YearCreation); err != nil {
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
func (r *AdvertRepo) CreateHouse(ctx context.Context, newHouse *models.House) error {
	insert := `INSERT INTO houses (id, buildingid, adverttypeid, ceilingheight, squarearea, squarehouse, bedroomcount, statusarea, cottage, statushome) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	if _, err := r.db.ExecContext(ctx, insert, newHouse.ID, newHouse.BuildingID, newHouse.AdvertTypeID, newHouse.CeilingHeight, newHouse.SquareArea, newHouse.SquareHouse, newHouse.BedroomCount, newHouse.StatusArea, newHouse.Cottage, newHouse.StatusHome); err != nil {
		return err
	}
	return nil
}

// CreateFlat creates a new flat in the database.
func (r *AdvertRepo) CreateFlat(ctx context.Context, newFlat *models.Flat) error {
	insert := `INSERT INTO flats (id, buildingid, adverttypeid, floor, ceilingheight, squaregeneral, roomcount, squareresidential, apartament) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	if _, err := r.db.ExecContext(ctx, insert, newFlat.ID, newFlat.BuildingID, newFlat.AdvertTypeID, newFlat.Floor, newFlat.CeilingHeight, newFlat.SquareGeneral, newFlat.RoomCount, newFlat.SquareResidential, newFlat.Apartment); err != nil {
		return err
	}
	return nil
}

// GetHouseSquareAdvertsList retrieves a square house adverts from the database.
func (r *AdvertRepo) GetHouseSquareAdvertsList(ctx context.Context) ([]*models.AdvertSquareData, error) {
	query := `SELECT
    a.id,
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
FROM adverts AS a LEFT JOIN
    LATERAL(
        SELECT *
        FROM pricechanges AS pc
        WHERE pc.advertid = a.id
        ORDER BY pc.datecreation DESC
        LIMIT 1
    ) AS pc
    ON true
INNER JOIN adverttypes AS at ON a.adverttypeid = at.id
INNER JOIN houses AS h ON at.id = h.adverttypeid
INNER JOIN buildings AS b ON h.buildingid = b.id
LEFT JOIN images AS i ON a.id = i.advertid
ORDER BY a.datecreation DESC;`
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

// GetFlatSquareAdvertsList retrieves a square Flat adverts from the database.
func (r *AdvertRepo) GetFlatSquareAdvertsList(ctx context.Context) ([]*models.AdvertSquareData, error) {
	query := `SELECT
    a.id,
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
FROM adverts AS a LEFT JOIN
    LATERAL(
        SELECT *
        FROM pricechanges AS pc
        WHERE pc.advertid = a.id
        ORDER BY pc.datecreation DESC
        LIMIT 1
    ) AS pc
    ON true
INNER JOIN adverttypes AS at ON a.adverttypeid = at.id
INNER JOIN flats AS f ON at.id = f.adverttypeid
INNER JOIN buildings AS b ON f.buildingid = b.id
LEFT JOIN images AS i ON a.id = i.advertid
ORDER BY a.datecreation DESC;`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	squareAdverts := []*models.AdvertSquareData{}
	for rows.Next() {
		squareAdvert := &models.AdvertSquareData{}
		var floor, floorGeneral, roomCount int
		var squareGenereal float64
		err := rows.Scan(&squareAdvert.ID, &squareAdvert.TypeAdvert, &squareAdvert.Photo, &squareAdvert.TypeSale, &squareAdvert.Address, &floor, &squareGenereal, &roomCount, &floorGeneral, &squareAdvert.Price, &squareAdvert.DateCreation)
		if err != nil {
			return nil, err
		}
		squareAdvert.Properties = make(map[string]interface{})
		squareAdvert.Properties["floor"] = floor
		squareAdvert.Properties["floorGeneral"] = floorGeneral
		squareAdvert.Properties["squareGeneral"] = squareGenereal
		squareAdvert.Properties["roomCount"] = roomCount
		squareAdverts = append(squareAdverts, squareAdvert)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return squareAdverts, nil
}
