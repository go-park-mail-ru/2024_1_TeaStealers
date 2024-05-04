package repo

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts"
	"2024_1_TeaStealers/internal/pkg/utils"
	"context"
	"database/sql"
	"fmt"

	"go.uber.org/zap"
)

// AdvertRepo represents a repository for adverts changes.
type AdvertRepo struct {
	db     *sql.DB
	logger *zap.Logger
}

// NewRepository creates a new instance of AdvertRepo.
func NewRepository(db *sql.DB, logger *zap.Logger) *AdvertRepo {
	return &AdvertRepo{db: db, logger: logger}
}

func (r *AdvertRepo) BeginTx(ctx context.Context) (models.Transaction, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.BeginTxMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.BeginTxMethod)
	return tx, nil
}

// CreateAdvertTypeHouse creates a new advertTypeHouse in the database.
func (r *AdvertRepo) CreateAdvertTypeHouse(ctx context.Context, tx models.Transaction, newAdvertType *models.HouseTypeAdvert) error {
	insert := `INSERT INTO advert_type_house (house_id, advert_id) VALUES ($1, $2)`
	if _, err := tx.ExecContext(ctx, insert, newAdvertType.HouseID, newAdvertType.AdvertID); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertTypeMethod, err)
		return err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertTypeMethod)
	return nil
}

// CreateAdvertTypeFlat creates a new advertTypeFlat in the database.
func (r *AdvertRepo) CreateAdvertTypeFlat(ctx context.Context, tx models.Transaction, newAdvertType *models.FlatTypeAdvert) error {
	insert := `INSERT INTO advert_type_flat (flat_id, advert_id) VALUES ($1, $2)`
	if _, err := tx.ExecContext(ctx, insert, newAdvertType.FlatID, newAdvertType.AdvertID); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertTypeMethod, err)
		return err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertTypeMethod)
	return nil
}

// CreateAdvert creates a new advert in the database.
func (r *AdvertRepo) CreateAdvert(ctx context.Context, tx models.Transaction, newAdvert *models.Advert) (int64, error) {
	insert := `INSERT INTO advert (user_id, type_placement, title, description, phone, is_agent, priority) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	var idAdvert int64
	if err := tx.QueryRowContext(ctx, insert, newAdvert.UserID, newAdvert.AdvertTypeSale, newAdvert.Title, newAdvert.Description, newAdvert.Phone, newAdvert.IsAgent, newAdvert.Priority).Scan(&idAdvert); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return 0, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod)
	return idAdvert, nil
}

// CreateProvince creates a new province in the database.
func (r *AdvertRepo) CreateProvince(ctx context.Context, tx models.Transaction, name string) (int64, error) {
	query := `SELECT id FROM province WHERE name=$1`

	res := r.db.QueryRow(query, name)

	var provinceId int64
	if err := res.Scan(&provinceId); err == nil {
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return provinceId, nil
	}

	insert := `INSERT INTO province (name) VALUES ($1) RETURNING id`
	if err := tx.QueryRowContext(ctx, insert, name).Scan(&provinceId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return 0, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod)
	return provinceId, nil
}

// CreateTown creates a new town in the database.
func (r *AdvertRepo) CreateTown(ctx context.Context, tx models.Transaction, idProvince int64, name string) (int64, error) {
	query := `SELECT id FROM town WHERE name=$1 AND province_id=$2`

	res := r.db.QueryRow(query, name, idProvince)

	var townId int64
	if err := res.Scan(&townId); err == nil {
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return townId, nil
	}

	insert := `INSERT INTO town (name, province_id) VALUES ($1, $2) RETURNING id`
	if err := tx.QueryRowContext(ctx, insert, name, idProvince).Scan(&townId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return 0, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod)
	return townId, nil
}

// CreateStreet creates a new street in the database.
func (r *AdvertRepo) CreateStreet(ctx context.Context, tx models.Transaction, idTown int64, name string) (int64, error) {
	query := `SELECT id FROM street WHERE name=$1 AND town_id=$2`

	res := r.db.QueryRow(query, name, idTown)

	var streetId int64
	if err := res.Scan(&streetId); err == nil {
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return streetId, nil
	}

	insert := `INSERT INTO street (name, town_id) VALUES ($1, $2) RETURNING id`
	if err := tx.QueryRowContext(ctx, insert, name, idTown).Scan(&streetId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return 0, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod)
	return streetId, nil
}

// CreateHouse creates a new house in the database.
func (r *AdvertRepo) CreateHouseAddress(ctx context.Context, tx models.Transaction, idStreet int64, name string) (int64, error) {
	query := `SELECT id FROM house_name WHERE name=$1 AND street_id=$2`

	res := r.db.QueryRow(query, name, idStreet)

	var houseId int64
	if err := res.Scan(&houseId); err == nil {
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return houseId, nil
	}

	insert := `INSERT INTO house_name (name, street_id) VALUES ($1, $2) RETURNING id`
	if err := tx.QueryRowContext(ctx, insert, name, idStreet).Scan(&houseId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return 0, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod)
	return houseId, nil
}

// CreateAddress creates a new address in the database.
func (r *AdvertRepo) CreateAddress(ctx context.Context, tx models.Transaction, idHouse int64, metro string, address_point string) (int64, error) {
	query := `SELECT id FROM address WHERE house_name_id=$1`

	res := r.db.QueryRow(query, idHouse)

	var addressId int64
	if err := res.Scan(&addressId); err == nil {
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return addressId, nil
	}

	insert := `INSERT INTO address (metro, house_name_id, address_point) VALUES ($1, $2, $3) RETURNING id`
	if err := tx.QueryRowContext(ctx, insert, metro, idHouse, address_point).Scan(&addressId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return 0, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod)
	return addressId, nil
}

// CreatePriceChange creates a new price change in the database.
func (r *AdvertRepo) CreatePriceChange(ctx context.Context, tx models.Transaction, newPriceChange *models.PriceChange) error {
	insert := `INSERT INTO price_change (advert_id, price) VALUES ($1, $2)`
	if _, err := tx.ExecContext(ctx, insert, newPriceChange.AdvertID, newPriceChange.Price); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreatePriceChangeMethod, err)
		return err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreatePriceChangeMethod)
	return nil
}

// CreateBuilding creates a new building in the database.
func (r *AdvertRepo) CreateBuilding(ctx context.Context, tx models.Transaction, newBuilding *models.Building) (int64, error) {
	insert := `INSERT INTO building (floor, material_building, address_id, year_creation) VALUES ($1, $2, $3, $4) RETURNING id`
	var idBuilding int64
	if err := tx.QueryRowContext(ctx, insert, newBuilding.Floor, newBuilding.Material, newBuilding.AddressID, newBuilding.YearCreation).Scan(&idBuilding); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateBuildingMethod, err)
		return 0, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateBuildingMethod)
	return idBuilding, nil
}

// CheckExistsBuilding check exists building.
func (r *AdvertRepo) CheckExistsBuilding(ctx context.Context, adress *models.AddressData) (*models.Building, error) {
	query := `SELECT b.id, b.address_id, b.floor, b.material_building, b.year_creation FROM building AS b JOIN address AS a ON b.address_id=a.id JOIN house_name AS h ON a.house_name_id=h.id JOIN street AS s ON h.street_id=s.id JOIN town AS t ON s.town_id=t.id JOIN province AS p ON t.province_id=p.id WHERE p.name=$1 AND t.name=$2 AND s.name=$3 AND h.name=$4;`

	building := &models.Building{}

	res := r.db.QueryRowContext(ctx, query, adress.Province, adress.Town, adress.Street, adress.House)

	if err := res.Scan(&building.ID, &building.AddressID, &building.Floor, &building.Material, &building.YearCreation); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsBuildingMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsBuildingMethod)
	return building, nil
}

// CheckExistsBuildings check exists buildings. Нужна для выпадающего списка существующих зданий по адресу(Для создания объявления)
func (r *AdvertRepo) CheckExistsBuildingData(ctx context.Context, adress *models.AddressData) (*models.BuildingData, error) {
	query := `SELECT b.floor, b.material_building, b.year_creation, COALESCE(c.name, '') FROM building AS b JOIN address AS a ON b.address_id=a.id JOIN house_name AS h ON a.house_name_id=h.id JOIN street AS s ON h.street_id=s.id JOIN town AS t ON s.town_id=t.id JOIN province AS p ON t.province_id=p.id LEFT JOIN complex AS c ON c.id=b.complex_id WHERE p.name=$1 AND t.name=$2 AND s.name=$3 AND h.name=$4;`

	building := &models.BuildingData{}

	res := r.db.QueryRowContext(ctx, query, adress.Province, adress.Town, adress.Street, adress.House)

	if err := res.Scan(&building.Floor, &building.Material, &building.YearCreation, &building.ComplexName); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsBuildingMethod, err)
		return nil, nil
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsBuildingMethod)
	return building, nil
}

// CreateHouse creates a new house in the database.
func (r *AdvertRepo) CreateHouse(ctx context.Context, tx models.Transaction, newHouse *models.House) (int64, error) {
	insert := `INSERT INTO house (building_id, ceiling_height, square_area, square_house, bedroom_count, status_area_house, cottage, status_home_house) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	var lastInsertID int64
	if err := tx.QueryRowContext(ctx, insert, newHouse.BuildingID, newHouse.CeilingHeight, newHouse.SquareArea, newHouse.SquareHouse, newHouse.BedroomCount, newHouse.StatusArea, newHouse.Cottage, newHouse.StatusHome).Scan(&lastInsertID); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateHouseMethod, err)
		return 0, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateHouseMethod)
	return lastInsertID, nil
}

// CreateFlat creates a new flat in the database.
func (r *AdvertRepo) CreateFlat(ctx context.Context, tx models.Transaction, newFlat *models.Flat) (int64, error) {
	insert := `INSERT INTO flat (building_id, floor, ceiling_height, square_general, bedroom_count, square_residential, apartament) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	var idFlat int64
	if err := tx.QueryRowContext(ctx, insert, newFlat.BuildingID, newFlat.Floor, newFlat.CeilingHeight, newFlat.SquareGeneral, newFlat.RoomCount, newFlat.SquareResidential, newFlat.Apartment).Scan(&idFlat); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateFlatMethod, err)
		return 0, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateFlatMethod)
	return idFlat, nil
}

// SelectImages select list images for advert
func (r *AdvertRepo) SelectImages(ctx context.Context, advertId int64) ([]*models.ImageResp, error) {
	selectQuery := `SELECT id, photo, priority FROM image WHERE advert_id = $1 AND is_deleted = false`
	rows, err := r.db.Query(selectQuery, advertId)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.SelectImagesMethod, err)
		return nil, err
	}
	defer rows.Close()

	images := []*models.ImageResp{}

	for rows.Next() {
		var id int64
		var photo string
		var priority int
		if err := rows.Scan(&id, &photo, &priority); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.SelectImagesMethod, err)
			return nil, err
		}
		image := &models.ImageResp{
			ID:       id,
			Photo:    photo,
			Priority: priority,
		}
		images = append(images, image)
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.SelectImagesMethod)
	return images, nil
}

// GetTypeAdvertById return type of advert
func (r *AdvertRepo) GetTypeAdvertById(ctx context.Context, id int64) (*models.AdvertTypeAdvert, error) {
	query := `SELECT                   CASE
	WHEN ath.house_id IS NOT NULL THEN 'House'
	WHEN atf.flat_id IS NOT NULL THEN 'Flat'
	ELSE 'None'
END AS type_advert FROM advert AS a LEFT JOIN advert_type_house AS ath ON a.id=ath.advert_id LEFT JOIN advert_type_flat AS atf ON a.id=atf.advert_id WHERE a.id=$1`

	res := r.db.QueryRowContext(ctx, query, id)

	var advertType *models.AdvertTypeAdvert

	if err := res.Scan(&advertType); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetTypeAdvertByIdMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetTypeAdvertByIdMethod)
	return advertType, nil
}

// GetHouseAdvertById retrieves full information about house advert from the database.
func (r *AdvertRepo) GetHouseAdvertById(ctx context.Context, id int64) (*models.AdvertData, error) {
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
		ad.address_point,
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
	res := r.db.QueryRowContext(ctx, query, id)

	advertData := &models.AdvertData{}
	var cottage bool
	var squareHouse, squareArea, ceilingHeight float64
	var bedroomCount, floor int
	var statusArea models.StatusAreaHouse
	var statusHome models.StatusHomeHouse
	var complexId, companyPhoto, companyName, complexName sql.NullString
	var metro, houseName, street, town, province string

	if err := res.Scan(
		&advertData.ID,
		&advertData.TypeSale,
		&advertData.Title,
		&advertData.Description,
		&advertData.Price,
		&advertData.Phone,
		&advertData.IsAgent,
		&metro,
		&houseName,
		&street,
		&town,
		&province,
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
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetHouseAdvertByIdMethod, err)
		return nil, err
	}

	advertData.AdvertType = "House"

	advertData.HouseProperties = &models.HouseProperties{}
	advertData.HouseProperties.CeilingHeight = ceilingHeight
	advertData.HouseProperties.SquareArea = squareArea
	advertData.HouseProperties.SquareHouse = squareHouse
	advertData.HouseProperties.BedroomCount = bedroomCount
	advertData.HouseProperties.StatusArea = statusArea
	advertData.HouseProperties.Cottage = cottage
	advertData.HouseProperties.StatusHome = statusHome
	advertData.HouseProperties.Floor = floor

	advertData.Address = province + ", " + town + ", " + street + ", " + houseName
	advertData.Metro = metro

	if complexId.String != "" {
		advertData.ComplexProperties = &models.ComplexAdvertProperties{}
		advertData.ComplexProperties.ComplexId = complexId.String
		advertData.ComplexProperties.PhotoCompany = companyPhoto.String
		advertData.ComplexProperties.NameCompany = companyName.String
		advertData.ComplexProperties.NameComplex = complexName.String
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetHouseAdvertByIdMethod)
	return advertData, nil
}

// CheckExistsFlat check exists flat.
func (r *AdvertRepo) CheckExistsFlat(ctx context.Context, advertId int64) (*models.Flat, error) {
	query := `SELECT f.id FROM advert AS a JOIN advert_type_flat AS at ON a.id=at.advert_id JOIN flat AS f ON f.id=at.flat_id WHERE a.id = $1`

	flat := &models.Flat{}

	res := r.db.QueryRowContext(ctx, query, advertId)

	if err := res.Scan(&flat.ID); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsFlatMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsFlatMethod)
	return flat, nil
}

// CheckExistsHouse check exists flat.
func (r *AdvertRepo) CheckExistsHouse(ctx context.Context, advertId int64) (*models.House, error) {
	query := `SELECT h.id FROM advert AS a JOIN advert_type_house AS at ON a.id=at.advert_id JOIN house AS h ON h.id=at.house_id WHERE a.id = $1;`

	house := &models.House{}

	res := r.db.QueryRowContext(ctx, query, advertId)

	if err := res.Scan(&house.ID); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsHouseMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsHouseMethod)
	return house, nil
}

// DeleteFlatAdvertById deletes a flat advert by ID.
func (r *AdvertRepo) DeleteFlatAdvertById(ctx context.Context, tx models.Transaction, advertId int64) error {
	queryGetIdTables := `
        SELECT
            f.id as flatid
        FROM
            advert AS a
        JOIN
            advert_type_flat AS at ON a.id = at.advert_id
        JOIN
            flat AS f ON f.id = at.flat_id
        WHERE a.id=$1;`

	res := tx.QueryRowContext(ctx, queryGetIdTables, advertId)

	var flatId int64
	if err := res.Scan(&flatId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteFlatAdvertByIdMethod, err)
		return err
	}

	queryDeleteAdvertById := `UPDATE advert SET is_deleted=true WHERE id=$1;`
	queryDeleteAdvertTypeById := `UPDATE advert_type_flat SET is_deleted=true WHERE advert_id=$1 AND flat_id=$2;`
	queryDeleteFlatById := `UPDATE flat SET is_deleted=true WHERE id=$1;`
	queryDeletePriceChanges := `UPDATE price_change SET is_deleted=true WHERE advert_id=$1;`
	queryDeleteImages := `UPDATE image SET is_deleted=true WHERE advert_id=$1;`

	if _, err := tx.Exec(queryDeleteAdvertById, advertId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteFlatAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(queryDeleteAdvertTypeById, advertId, flatId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteFlatAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(queryDeleteFlatById, flatId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteFlatAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(queryDeletePriceChanges, advertId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteFlatAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(queryDeleteImages, advertId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteFlatAdvertByIdMethod, err)
		return err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteFlatAdvertByIdMethod)
	return nil
}

// DeleteHouseAdvertById deletes a house advert by ID.
func (r *AdvertRepo) DeleteHouseAdvertById(ctx context.Context, tx models.Transaction, advertId int64) error {
	queryGetIdTables := `
        SELECT
            h.id as houseid
        FROM
            advert AS a
        JOIN
            advert_type_house AS at ON a.id = at.advert_id
        JOIN
            house AS h ON h.id = at.house_id
        WHERE a.id=$1;`

	res := tx.QueryRowContext(ctx, queryGetIdTables, advertId)

	var houseId int64
	if err := res.Scan(&houseId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteHouseAdvertByIdMethod, err)
		return err
	}

	queryDeleteAdvertById := `UPDATE advert SET is_deleted=true WHERE id=$1;`
	queryDeleteAdvertTypeById := `UPDATE advert_type_house SET is_deleted=true WHERE advert_id=$1 AND house_id=$2;`
	queryDeleteHouseById := `UPDATE house SET is_deleted=true WHERE id=$1;`
	queryDeletePriceChanges := `UPDATE price_change SET is_deleted=true WHERE advert_id=$1;`
	queryDeleteImages := `UPDATE image SET is_deleted=true WHERE advert_id=$1;`

	if _, err := tx.Exec(queryDeleteAdvertById, advertId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteHouseAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(queryDeleteAdvertTypeById, advertId, houseId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteHouseAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(queryDeleteHouseById, houseId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteHouseAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(queryDeletePriceChanges, advertId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteHouseAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(queryDeleteImages, advertId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteHouseAdvertByIdMethod, err)
		return err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteHouseAdvertByIdMethod)
	return nil
}

// ChangeTypeAdvert. Когда мы захотели поменять тип объявления(Дом, Квартира), Меняем сущности в бд
func (r *AdvertRepo) ChangeTypeAdvert(ctx context.Context, tx models.Transaction, advertId int64) (err error) {
	query := `SELECT 			CASE
	WHEN ath.house_id IS NOT NULL THEN 'House'
	WHEN atf.flat_id IS NOT NULL THEN 'Flat'
	ELSE 'None'
END AS type_advert FROM advert AS a LEFT JOIN advert_type_flat AS atf ON a.id=atf.advert_id LEFT JOIN advert_type_house AS ath ON a.id=ath.advert_id WHERE a.id = $1;`
	querySelectBuildingIdByFlat := `SELECT b.id AS buildingid, f.id AS flatid  FROM advert AS a JOIN advert_type_flat AS at ON at.advert_id=a.id JOIN flat AS f ON f.id=at.flat_id JOIN building AS b ON f.building_id=b.id WHERE a.id=$1`
	querySelectBuildingIdByHouse := `SELECT b.id AS buildingid, h.id AS houseid  FROM advert AS a JOIN advert_type_house AS at ON at.advert_id=a.id JOIN house AS h ON h.id=at.house_id JOIN building AS b ON h.building_id=b.id WHERE a.id=$1`
	queryInsertFlat := `INSERT INTO flat (building_id, floor, ceiling_height, square_general, bedroom_count, square_residential, apartament)
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	queryInsertHouse := `INSERT INTO house (building_id, ceiling_height, square_area, square_house, bedroom_count, status_area_house, cottage, status_home_house)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	queryInsertTypeFlat := `INSERT INTO advert_type_flat (advert_id, flat_id) VALUES ($1, $2);`
	queryInsertTypeHouse := `INSERT INTO advert_type_house (advert_id, house_id) VALUES ($1, $2);`
	queryRestoreFlatById := `UPDATE flat SET is_deleted=false WHERE id=$1;`
	queryRestoreHouseById := `UPDATE house SET is_deleted=false WHERE id=$1;`
	queryDeleteFlatById := `UPDATE flat SET is_deleted=true WHERE id=$1;`
	queryDeleteHouseById := `UPDATE house SET is_deleted=true WHERE id=$1;`
	queryDeleteAdvertTypeFlat := `UPDATE advert_type_flat SET is_deleted=true WHERE advert_id=$1 AND flat_id=$2;`
	queryDeleteAdvertTypeHouse := `UPDATE advert_type_house SET is_deleted=true WHERE advert_id=$1 AND house_id=$2;`
	queryRestoreAdvertTypeFlat := `UPDATE advert_type_flat SET is_deleted=false WHERE advert_id=$1 AND flat_id=$2;`
	queryRestoreAdvertTypeHouse := `UPDATE advert_type_house SET is_deleted=false WHERE advert_id=$1 AND house_id=$2;`

	var advertType models.AdvertTypeAdvert
	res := r.db.QueryRowContext(ctx, query, advertId)

	if err := res.Scan(&advertType); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
		return err
	}
	var buildingId int64
	switch advertType {
	case models.AdvertTypeFlat:
		res := r.db.QueryRowContext(ctx, querySelectBuildingIdByFlat, advertId)

		var flatId int64

		if err := res.Scan(&buildingId, &flatId); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
			return err
		}

		if _, err := tx.Exec(queryDeleteFlatById, flatId); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
			return err
		}

		if _, err := tx.Exec(queryDeleteAdvertTypeFlat, advertId, flatId); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
			return err
		}

		house, err := r.CheckExistsHouse(ctx, advertId)
		if err != nil {
			var id int64
			house := &models.House{}
			err := tx.QueryRowContext(ctx, queryInsertHouse, buildingId, house.CeilingHeight, house.SquareArea, house.SquareHouse, house.BedroomCount, models.StatusAreaDNP, house.Cottage, models.StatusHomeCompleteNeed).Scan(&id)
			if err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
				return err
			}
			if _, err := tx.Exec(queryInsertTypeHouse, advertId, id); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
				return err
			}
		} else {
			if _, err := tx.Exec(queryRestoreHouseById, house.ID); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
				return err
			}

			if _, err := tx.Exec(queryRestoreAdvertTypeHouse, advertId, house.ID); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
				return err
			}
		}
	case models.AdvertTypeHouse:
		res := r.db.QueryRowContext(ctx, querySelectBuildingIdByHouse, advertId)

		var houseId int64

		if err := res.Scan(&buildingId, &houseId); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
			return err
		}

		if _, err := tx.Exec(queryDeleteHouseById, houseId); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
			return err
		}

		if _, err := tx.Exec(queryDeleteAdvertTypeHouse, advertId, houseId); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
			return err
		}

		flat, err := r.CheckExistsFlat(ctx, advertId)
		if err != nil {
			var id int64
			flat = &models.Flat{}
			err := tx.QueryRowContext(ctx, queryInsertFlat, buildingId, flat.Floor, flat.CeilingHeight, flat.SquareGeneral, flat.RoomCount, flat.SquareResidential, flat.Apartment).Scan(&id)
			if err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
				return err
			}
			if _, err := tx.Exec(queryInsertTypeFlat, advertId, id); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
				return err
			}
		} else {
			if _, err := tx.Exec(queryRestoreFlatById, flat.ID); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
				return err
			}

			if _, err := tx.Exec(queryRestoreAdvertTypeFlat, advertId, flat.ID); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
				return err
			}
		}
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod)
	return nil
}

// UpdateHouseAdvertById updates a house advert in the database.
func (r *AdvertRepo) UpdateHouseAdvertById(ctx context.Context, tx models.Transaction, advertUpdateData *models.AdvertUpdateData) error {
	queryGetIdTables := `
        SELECT
            b.id as buildingid,
            h.id as houseid,
            pc.price
        FROM
            advert AS a
        JOIN
            advert_type_house AS at ON a.id = at.advert_id
        JOIN
            house AS h ON h.id = at.house_id
        JOIN
            building AS b ON h.building_id = b.id
        LEFT JOIN
            LATERAL (
                SELECT *
                FROM price_change AS pc
                WHERE pc.advert_id = a.id
                ORDER BY pc.created_at DESC
                LIMIT 1
            ) AS pc ON TRUE
        WHERE a.id=$1;`

	res := tx.QueryRowContext(ctx, queryGetIdTables, advertUpdateData.ID)

	var buildingId, houseId int64
	var price float64
	if err := res.Scan(&buildingId, &houseId, &price); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateHouseAdvertByIdMethod, err)
		return err
	}

	id, err := r.CreateProvince(ctx, tx, advertUpdateData.Address.Province)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return err
	}

	id, err = r.CreateTown(ctx, tx, id, advertUpdateData.Address.Town)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return err
	}

	id, err = r.CreateStreet(ctx, tx, id, advertUpdateData.Address.Street)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return err
	}

	id, err = r.CreateHouseAddress(ctx, tx, id, advertUpdateData.Address.House)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return err
	}

	id, err = r.CreateAddress(ctx, tx, id, advertUpdateData.Address.Metro, advertUpdateData.Address.AddressPoint)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return err
	}

	queryUpdateAdvertById := `UPDATE advert SET type_placement=$1, title=$2, description=$3, phone=$4, is_agent=$5 WHERE id=$6;`
	queryUpdateBuildingById := `UPDATE building SET floor=$1, material_building=$2, address_id=$3, year_creation=$4 WHERE id=$5;`
	queryUpdateHouseById := `UPDATE house SET ceiling_height=$1, square_area=$2, square_house=$3, bedroom_count=$4, status_area_house=$5, cottage=$6, status_home_house=$7 WHERE id=$8;`

	if _, err := tx.Exec(queryUpdateAdvertById, advertUpdateData.TypeSale, advertUpdateData.Title, advertUpdateData.Description, advertUpdateData.Phone, advertUpdateData.IsAgent, advertUpdateData.ID); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateHouseAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(queryUpdateBuildingById, advertUpdateData.HouseProperties.Floor, advertUpdateData.Material, id, advertUpdateData.YearCreation, buildingId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateHouseAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(queryUpdateHouseById, advertUpdateData.HouseProperties.CeilingHeight, advertUpdateData.HouseProperties.SquareArea, advertUpdateData.HouseProperties.SquareHouse, advertUpdateData.HouseProperties.BedroomCount, advertUpdateData.HouseProperties.StatusArea, advertUpdateData.HouseProperties.Cottage, advertUpdateData.HouseProperties.StatusHome, houseId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateHouseAdvertByIdMethod, err)
		return err
	}
	if advertUpdateData.Price != price {
		queryInsertPriceChange := `INSERT INTO price_change (advertId, price)
            VALUES ($1, $2)`
		if _, err := tx.Exec(queryInsertPriceChange, advertUpdateData.ID, advertUpdateData.Price); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateHouseAdvertByIdMethod, err)
			return err
		}
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateHouseAdvertByIdMethod)
	return nil
}

// UpdateFlatAdvertById updates a flat advert in the database.
func (r *AdvertRepo) UpdateFlatAdvertById(ctx context.Context, tx models.Transaction, advertUpdateData *models.AdvertUpdateData) error {
	queryGetIdTables := `
        SELECT
            b.id as buildingid,
            f.id as flatid,
            pc.price
        FROM
            advert AS a
        JOIN
            advert_type_flat AS at ON a.id = at.advert_id
        JOIN
            flat AS f ON f.id = at.flat_id
        JOIN
            building AS b ON f.building_id = b.id
        LEFT JOIN
            LATERAL (
                SELECT *
                FROM price_change AS pc
                WHERE pc.advert_id = a.id
                ORDER BY pc.created_at DESC
                LIMIT 1
            ) AS pc ON TRUE
        WHERE a.id=$1;`

	res := tx.QueryRowContext(ctx, queryGetIdTables, advertUpdateData.ID)

	var buildingId, flatId int64
	var price float64
	if err := res.Scan(&buildingId, &flatId, &price); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateFlatAdvertByIdMethod, err)
		return err
	}

	id, err := r.CreateProvince(ctx, tx, advertUpdateData.Address.Province)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return err
	}

	id, err = r.CreateTown(ctx, tx, id, advertUpdateData.Address.Town)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return err
	}

	id, err = r.CreateStreet(ctx, tx, id, advertUpdateData.Address.Street)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return err
	}

	id, err = r.CreateHouseAddress(ctx, tx, id, advertUpdateData.Address.House)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return err
	}

	id, err = r.CreateAddress(ctx, tx, id, advertUpdateData.Address.Metro, advertUpdateData.Address.AddressPoint)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.UsecaseLayer, adverts.CreateHouseAdvertMethod, err)
		return err
	}

	queryUpdateAdvertById := `UPDATE advert SET type_placement=$1, title=$2, description=$3, phone=$4, is_agent=$5 WHERE id=$6;`
	queryUpdateBuildingById := `UPDATE building SET floor=$1, material_building=$2, address_id=$3, year_creation=$4 WHERE id=$5;`
	queryUpdateFlatById := `UPDATE flat SET floor=$1, ceiling_height=$2, square_general=$3, bedroom_count=$4, square_residential=$5, apartament=$6 WHERE id=$7;`

	if _, err := tx.Exec(queryUpdateAdvertById, advertUpdateData.TypeSale, advertUpdateData.Title, advertUpdateData.Description, advertUpdateData.Phone, advertUpdateData.IsAgent, advertUpdateData.ID); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateFlatAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(queryUpdateBuildingById, advertUpdateData.FlatProperties.FloorGeneral, advertUpdateData.Material, id, advertUpdateData.YearCreation, buildingId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateFlatAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(queryUpdateFlatById, advertUpdateData.FlatProperties.Floor, advertUpdateData.FlatProperties.CeilingHeight, advertUpdateData.FlatProperties.SquareGeneral, advertUpdateData.FlatProperties.RoomCount, advertUpdateData.FlatProperties.SquareResidential, advertUpdateData.FlatProperties.Apartment, flatId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateFlatAdvertByIdMethod, err)
		return err
	}

	if advertUpdateData.Price != price {
		queryInsertPriceChange := `INSERT INTO price_change (advert_id, price)
            VALUES ($1, $2)`
		if _, err := tx.Exec(queryInsertPriceChange, advertUpdateData.ID, advertUpdateData.Price); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateFlatAdvertByIdMethod, err)
			return err
		}
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateFlatAdvertByIdMethod)
	return nil
}

// GetFlatAdvertById retrieves full information about flat advert from the database.
func (r *AdvertRepo) GetFlatAdvertById(ctx context.Context, id int64) (*models.AdvertData, error) {
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
		ad.address_point,
        f.floor,
        f.ceiling_height,
        f.square_general,
        f.bedroom_count,
        f.square_residential,
        f.apartament,
        b.floor AS floorGeneral,
        b.year_creation,
        COALESCE(b.material_building, 'Brick') as material,
        a.created_at,
        cx.id AS complexid,
        c.photo AS companyphoto,
        c.name AS companyname,
        cx.name AS complexname
    FROM
        advert AS a
    JOIN
        advert_type_flat AS at ON a.id = at.advert_id
    JOIN
        flat AS f ON f.id = at.flat_id
    JOIN
        building AS b ON f.building_id = b.id
		JOIN address AS ad ON b.address_id=ad.id
		JOIN house_name AS hn ON hn.id=ad.house_name_id
		JOIN street AS s ON s.id=hn.street_id
		JOIN town AS t ON t.id=s.town_id
		JOIN province AS p ON p.id=t.province_id
    LEFT JOIN
        complex AS cx ON b.complex_id = cx.id
    LEFT JOIN
        company AS c ON cx.company_id = c.id
    LEFT JOIN
        LATERAL (
            SELECT *
            FROM price_change AS pc
            WHERE pc.advert_id = a.id
            ORDER BY pc.created_at DESC
            LIMIT 1
        ) AS pc ON TRUE
    WHERE
        a.id = $1 AND a.is_deleted = FALSE;`
	res := r.db.QueryRowContext(ctx, query, id)

	advertData := &models.AdvertData{}
	var floor, floorGeneral, roomCount int
	var squareGenereal, squareResidential, ceilingHeight float64
	var apartament sql.NullBool
	var complexId, companyPhoto, companyName, complexName sql.NullString
	var metro, houseName, street, town, province string

	if err := res.Scan(
		&advertData.ID,
		&advertData.TypeSale,
		&advertData.Title,
		&advertData.Description,
		&advertData.Price,
		&advertData.Phone,
		&advertData.IsAgent,
		&metro,
		&houseName,
		&street,
		&town,
		&province,
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
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetFlatAdvertByIdMethod, err)
		return nil, err
	}

	advertData.AdvertType = "Flat"
	advertData.FlatProperties = &models.FlatProperties{}
	advertData.FlatProperties.CeilingHeight = ceilingHeight
	advertData.FlatProperties.Apartment = apartament.Bool
	advertData.FlatProperties.SquareResidential = squareResidential
	advertData.FlatProperties.RoomCount = roomCount
	advertData.FlatProperties.SquareGeneral = squareGenereal
	advertData.FlatProperties.FloorGeneral = floorGeneral
	advertData.FlatProperties.Floor = floor

	advertData.Address = province + ", " + town + ", " + street + ", " + houseName
	advertData.Metro = metro

	if complexId.String != "" {
		advertData.ComplexProperties = &models.ComplexAdvertProperties{}
		advertData.ComplexProperties.ComplexId = complexId.String
		advertData.ComplexProperties.PhotoCompany = companyPhoto.String
		advertData.ComplexProperties.NameCompany = companyName.String
		advertData.ComplexProperties.NameComplex = complexName.String
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetFlatAdvertByIdMethod)
	return advertData, nil
}

// GetSquareAdverts retrieves square adverts from the database.
func (r *AdvertRepo) GetSquareAdverts(ctx context.Context, pageSize, offset int) ([]*models.AdvertSquareData, error) {
	queryBaseAdvert := `
        SELECT
            a.id,
            a.type_placement,
			CASE
           		WHEN ath.house_id IS NOT NULL THEN 'House'
           		WHEN atf.flat_id IS NOT NULL THEN 'Flat'
           		ELSE 'None'
       		END AS type_advert,
            i.photo,
            pc.price,
            a.created_at
        FROM
            advert AS a
			LEFT JOIN advert_type_house AS ath ON ath.advert_id=a.id
			LEFT JOIN advert_type_flat AS atf ON atf.advert_id=a.id
            LEFT JOIN LATERAL (
                SELECT *
                FROM price_change AS pc
                WHERE pc.advert_id = a.id
                ORDER BY pc.created_at DESC
                LIMIT 1
            ) AS pc ON TRUE
            JOIN image AS i ON i.advert_id = a.id
        WHERE i.priority = (
                SELECT MIN(priority)
                FROM image
                WHERE advert_id = a.id
                    AND is_deleted = FALSE
            )
            AND i.is_deleted = FALSE
        ORDER BY
            a.created_at DESC
        LIMIT $1
        OFFSET $2;`
	queryFlat := `
        SELECT 
            f.square_general,
            f.floor,
            ad.metro,
			hn.name,
			s.name,
			t.name,
			p.name,
            b.floor AS floorgeneral,
            f.bedroom_count
        FROM
            advert AS a
            JOIN advert_type_flat AS at ON a.id = at.advert_id
            JOIN flat AS f ON f.id=at.flat_id
            JOIN building AS b ON f.building_id=b.id
			JOIN address AS ad ON b.address_id=ad.id
			JOIN house_name AS hn ON hn.id=ad.house_name_id
			JOIN street AS s ON s.id=hn.street_id
			JOIN town AS t ON t.id=s.town_id
			JOIN province AS p ON p.id=t.province_id
        WHERE a.id=$1 AND a.is_deleted = FALSE
        ORDER BY
            a.created_at DESC;`
	queryHouse := `
        	SELECT 
			ad.metro,
			hn.name,
			s.name,
			t.name,
			p.name,
            h.cottage,
            h.square_house,
            h.square_area,
            h.bedroom_count,
            b.floor
        FROM
            advert AS a
            JOIN advert_type_house AS at ON a.id = at.advert_id
            JOIN house AS h ON h.id=at.house_id
            JOIN building AS b ON h.building_id=b.id
			JOIN address AS ad ON b.address_id=ad.id
			JOIN house_name AS hn ON hn.id=ad.house_name_id
			JOIN street AS s ON s.id=hn.street_id
			JOIN town AS t ON t.id=s.town_id
			JOIN province AS p ON p.id=t.province_id
        WHERE a.id=$1
        ORDER BY
            a.created_at DESC;`

	rows, err := r.db.Query(queryBaseAdvert, pageSize, offset)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetSquareAdvertsMethod, err)
		return nil, err
	}
	defer rows.Close()

	squareAdverts := []*models.AdvertSquareData{}
	for rows.Next() {
		squareAdvert := &models.AdvertSquareData{}
		err := rows.Scan(&squareAdvert.ID, &squareAdvert.TypeSale, &squareAdvert.TypeAdvert, &squareAdvert.Photo, &squareAdvert.Price, &squareAdvert.DateCreation)
		if err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetSquareAdvertsMethod, err)
			return nil, err
		}
		var metro, province, town, street, houseName string
		switch squareAdvert.TypeAdvert {
		case string(models.AdvertTypeFlat):
			var squareGeneral float64
			var floor, floorGeneral, roomCount int
			row := r.db.QueryRowContext(ctx, queryFlat, squareAdvert.ID)
			if err := row.Scan(&squareGeneral, &floor, &metro, &houseName, &street, &town, &province, &floorGeneral, &roomCount); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetSquareAdvertsMethod, err)
				return nil, err
			}
			squareAdvert.FlatProperties = &models.FlatSquareProperties{}
			squareAdvert.FlatProperties.Floor = floor
			squareAdvert.FlatProperties.FloorGeneral = floorGeneral
			squareAdvert.FlatProperties.RoomCount = roomCount
			squareAdvert.FlatProperties.SquareGeneral = squareGeneral
		case string(models.AdvertTypeHouse):
			var cottage bool
			var squareHouse, squareArea float64
			var bedroomCount, floor int
			row := r.db.QueryRowContext(ctx, queryHouse, squareAdvert.ID)
			if err := row.Scan(&metro, &houseName, &street, &town, &province, &cottage, &squareHouse, &squareArea, &bedroomCount, &floor); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetSquareAdvertsMethod, err)
				return nil, err
			}
			squareAdvert.HouseProperties = &models.HouseSquareProperties{}
			squareAdvert.HouseProperties.Cottage = cottage
			squareAdvert.HouseProperties.SquareHouse = squareHouse
			squareAdvert.HouseProperties.SquareArea = squareArea
			squareAdvert.HouseProperties.BedroomCount = bedroomCount
			squareAdvert.HouseProperties.Floor = floor
		}

		squareAdvert.Address = province + ", " + town + ", " + street + ", " + houseName
		squareAdvert.Metro = metro

		squareAdverts = append(squareAdverts, squareAdvert)
	}
	if err := rows.Err(); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetSquareAdvertsMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetSquareAdvertsMethod)
	return squareAdverts, nil
}

// GetRectangleAdverts retrieves rectangle adverts from the database with search.
func (r *AdvertRepo) GetRectangleAdverts(ctx context.Context, advertFilter models.AdvertFilter) (*models.AdvertDataPage, error) {
	queryBaseAdvert := `
        SELECT
			a.id,
			a.title,
			a.description,
			CASE
			   WHEN ath.house_id IS NOT NULL THEN 'House'
			   WHEN atf.flat_id IS NOT NULL THEN 'Flat'
			   ELSE 'None'
		    END AS type_advert, 
            CASE
                WHEN atf.flat_id IS NOT NULL THEN f.bedroom_count
                WHEN ath.house_id IS NOT NULL THEN h.bedroom_count
                ELSE 0
            END AS rcount,
            a.phone,
            a.type_placement,
            pc.price,
            i.photo,
            a.created_at
        FROM
            advert AS a
            LEFT JOIN advert_type_house AS ath ON a.id = ath.advert_id
			LEFT JOIN advert_type_flat AS atf ON a.id = atf.advert_id
            LEFT JOIN flat AS f ON f.id = atf.flat_id
            LEFT JOIN house AS h ON h.id = ath.house_id
            JOIN building AS b ON (f.building_id = b.id OR h.building_id = b.id)
			JOIN address AS ad ON b.address_id=ad.id
			JOIN house_name AS hn ON hn.id=ad.house_name_id
			JOIN street AS s ON s.id=hn.street_id
			JOIN town AS t ON t.id=s.town_id
			JOIN province AS p ON p.id=t.province_id
            LEFT JOIN LATERAL (
                SELECT *
                FROM price_change AS pc
                WHERE pc.advert_id = a.id
                ORDER BY pc.created_at DESC
                LIMIT 1
            ) AS pc ON TRUE
            JOIN image AS i ON i.advert_id = a.id
        WHERE i.priority = (
                SELECT MIN(priority)
                FROM image
                WHERE advert_id = a.id
                    AND is_deleted = FALSE
            )
            AND i.is_deleted = FALSE
            AND a.is_deleted = FALSE
            AND pc.price >= $1
            AND pc.price <= $2
            AND CONCAT_WS(', ', COALESCE(p.name, ''), COALESCE(t.name, ''), COALESCE(s.name, ''), COALESCE(hn.name, '')) ILIKE $3`
	queryFlat := `
        SELECT
            f.square_general,
            f.floor,
			ad.address_point,
			ad.metro,
			hn.name,
			s.name,
			t.name,
			p.name,
            b.floor AS floorgeneral
        FROM
            advert AS a
            JOIN advert_type_flat AS at ON a.id = at.advert_id
            JOIN flat AS f ON f.id = at.flat_id
            JOIN building AS b ON f.building_id = b.id
			JOIN address AS ad ON b.address_id=ad.id
			JOIN house_name AS hn ON hn.id=ad.house_name_id
			JOIN street AS s ON s.id=hn.street_id
			JOIN town AS t ON t.id=s.town_id
			JOIN province AS p ON p.id=t.province_id
        WHERE a.id = $1
        ORDER BY
            a.created_at DESC;`
	queryHouse := `
        SELECT
			ad.address_point,
			ad.metro,
			hn.name,
			s.name,
			t.name,
			p.name,
            h.cottage,
            h.square_house,
            h.square_area,
            b.floor
        FROM
			advert AS a
			JOIN advert_type_house AS at ON a.id = at.advert_id
			JOIN house AS h ON h.id = at.house_id
			JOIN building AS b ON h.building_id = b.id
			JOIN address AS ad ON b.address_id=ad.id
			JOIN house_name AS hn ON hn.id=ad.house_name_id
			JOIN street AS s ON s.id=hn.street_id
			JOIN town AS t ON t.id=s.town_id
			JOIN province AS p ON p.id=t.province_id
        WHERE a.id = $1
        ORDER BY
            a.created_at DESC;`

	pageInfo := &models.PageInfo{}

	var argsForQuery []interface{}
	i := 4 // Изначально в запросе проставлены minPrice, maxPrice и address, поэтому начинаем с 4 для формирования поиска

	advertFilter.Address = "%" + advertFilter.Address + "%"

	if advertFilter.DealType != "" {
		queryBaseAdvert += " AND a.type_placement = $" + fmt.Sprint(i) + " "
		argsForQuery = append(argsForQuery, advertFilter.DealType)
		i++
	}

	if advertFilter.AdvertType != "" {
		queryBaseAdvert = "SELECT * FROM (" + queryBaseAdvert + ") AS subqueryforadverttypecalculate WHERE type_advert = $" + fmt.Sprint(i) + " "
		argsForQuery = append(argsForQuery, advertFilter.AdvertType)
		i++
	}

	if advertFilter.RoomCount != 0 {
		queryBaseAdvert = "SELECT * FROM (" + queryBaseAdvert + ") AS subqueryforroomcountcalculate WHERE rcount = $" + fmt.Sprint(i) + " "
		argsForQuery = append(argsForQuery, advertFilter.RoomCount)
		i++
	}

	queryCount := "SELECT COUNT(*) FROM (" + queryBaseAdvert + ") AS subqueryforpaginate"
	queryBaseAdvert += " ORDER BY created_at DESC LIMIT $" + fmt.Sprint(i) + " OFFSET $" + fmt.Sprint(i+1) + ";"
	rowCountQuery := r.db.QueryRowContext(ctx, queryCount, append([]interface{}{advertFilter.MinPrice, advertFilter.MaxPrice, advertFilter.Address}, argsForQuery...)...)

	if err := rowCountQuery.Scan(&pageInfo.TotalElements); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsMethod, err)
		return nil, err
	}

	argsForQuery = append(argsForQuery, advertFilter.Page, advertFilter.Offset)
	rows, err := r.db.Query(queryBaseAdvert, append([]interface{}{advertFilter.MinPrice, advertFilter.MaxPrice, advertFilter.Address}, argsForQuery...)...)

	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsMethod, err)
		return nil, err
	}

	defer rows.Close()

	rectangleAdverts := []*models.AdvertRectangleData{}

	for rows.Next() {
		var roomCount int
		rectangleAdvert := &models.AdvertRectangleData{}
		err := rows.Scan(&rectangleAdvert.ID, &rectangleAdvert.Title, &rectangleAdvert.Description, &rectangleAdvert.TypeAdvert, &roomCount, &rectangleAdvert.Phone, &rectangleAdvert.TypeSale, &rectangleAdvert.Price, &rectangleAdvert.Photo, &rectangleAdvert.DateCreation)

		if err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsMethod, err)
			return nil, err
		}

		var metro, houseName, street, town, province string

		switch rectangleAdvert.TypeAdvert {
		case string(models.AdvertTypeFlat):
			var squareGeneral float64
			var floor, floorGeneral int
			row := r.db.QueryRowContext(ctx, queryFlat, rectangleAdvert.ID)

			if err := row.Scan(&squareGeneral, &floor, &rectangleAdvert.AddressPoint, &metro, &houseName, &street, &town, &province, &floorGeneral); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsMethod, err)
				return nil, err
			}

			rectangleAdvert.FlatProperties = &models.FlatRectangleProperties{}
			rectangleAdvert.FlatProperties.Floor = floor
			rectangleAdvert.FlatProperties.FloorGeneral = floorGeneral
			rectangleAdvert.FlatProperties.SquareGeneral = squareGeneral
			rectangleAdvert.FlatProperties.RoomCount = roomCount
		case string(models.AdvertTypeHouse):
			var cottage bool
			var squareHouse, squareArea float64
			var floor int
			row := r.db.QueryRowContext(ctx, queryHouse, rectangleAdvert.ID)

			if err := row.Scan(&rectangleAdvert.AddressPoint, &metro, &houseName, &street, &town, &province, &cottage, &squareHouse, &squareArea, &floor); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsMethod, err)
				return nil, err
			}

			rectangleAdvert.HouseProperties = &models.HouseRectangleProperties{}
			rectangleAdvert.HouseProperties.Cottage = cottage
			rectangleAdvert.HouseProperties.SquareHouse = squareHouse
			rectangleAdvert.HouseProperties.SquareArea = squareArea
			rectangleAdvert.HouseProperties.BedroomCount = roomCount
			rectangleAdvert.HouseProperties.Floor = floor
		}

		rectangleAdvert.Address = province + ", " + town + ", " + street + ", " + houseName
		rectangleAdvert.Metro = metro

		rectangleAdverts = append(rectangleAdverts, rectangleAdvert)
	}

	if err := rows.Err(); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsMethod, err)
		return nil, err
	}

	pageInfo.PageSize = advertFilter.Page
	pageInfo.TotalPages = pageInfo.TotalElements / pageInfo.PageSize

	if pageInfo.TotalElements%pageInfo.PageSize != 0 {
		pageInfo.TotalPages++
	}

	pageInfo.CurrentPage = (advertFilter.Offset / pageInfo.PageSize) + 1

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsMethod)

	return &models.AdvertDataPage{
		Adverts:  rectangleAdverts,
		PageInfo: pageInfo,
	}, nil
}

// GetRectangleAdvertsByUserId retrieves rectangle adverts from the database by user id.
func (r *AdvertRepo) GetRectangleAdvertsByUserId(ctx context.Context, pageSize, offset int, userId int64) ([]*models.AdvertRectangleData, error) {
	queryBaseAdvert := `
        SELECT
			a.id,
			a.title,
			a.description,
			CASE
			   WHEN ath.house_id IS NOT NULL THEN 'House'
			   WHEN atf.flat_id IS NOT NULL THEN 'Flat'
			   ELSE 'None'
		    END AS type_advert, 
            CASE
                WHEN atf.flat_id IS NOT NULL THEN f.bedroom_count
                WHEN ath.house_id IS NOT NULL THEN h.bedroom_count
                ELSE 0
            END AS rcount,
            a.phone,
            a.type_placement,
            pc.price,
            i.photo,
            a.created_at
        FROM
            advert AS a
            LEFT JOIN advert_type_house AS ath ON a.id = ath.advert_id
			LEFT JOIN advert_type_flat AS atf ON a.id = atf.advert_id
            LEFT JOIN flat AS f ON f.id = atf.flat_id
            LEFT JOIN house AS h ON h.id = ath.house_id
            JOIN building AS b ON (f.building_id = b.id OR h.building_id = b.id)
			JOIN address AS ad ON b.address_id=ad.id
			JOIN house_name AS hn ON hn.id=ad.house_name_id
			JOIN street AS s ON s.id=hn.street_id
			JOIN town AS t ON t.id=s.town_id
			JOIN province AS p ON p.id=t.province_id
            LEFT JOIN LATERAL (
                SELECT *
                FROM price_change AS pc
                WHERE pc.advert_id = a.id
                ORDER BY pc.created_at DESC
                LIMIT 1
            ) AS pc ON TRUE
            JOIN image AS i ON i.advert_id = a.id
        WHERE i.priority = (
                SELECT MIN(priority)
                FROM image
                WHERE advert_id = a.id
                    AND is_deleted = FALSE
            )
            AND i.is_deleted = FALSE
            AND a.is_deleted = FALSE
            AND a.user_id = $1
			ORDER BY a.created_at DESC
			LIMIT $2
			OFFSET $3`
	queryFlat := `
        SELECT
            f.square_general,
            f.floor,
			ad.metro,
			hn.name,
			s.name,
			t.name,
			p.name,
            b.floor AS floorgeneral
        FROM
            advert AS a
            JOIN advert_type_flat AS at ON a.id = at.advert_id
            JOIN flat AS f ON f.id = at.flat_id
            JOIN building AS b ON f.building_id = b.id
			JOIN address AS ad ON b.address_id=ad.id
			JOIN house_name AS hn ON hn.id=ad.house_name_id
			JOIN street AS s ON s.id=hn.street_id
			JOIN town AS t ON t.id=s.town_id
			JOIN province AS p ON p.id=t.province_id
        WHERE a.id = $1
        ORDER BY
            a.created_at DESC;`
	queryHouse := `
        SELECT
			ad.metro,
			hn.name,
			s.name,
			t.name,
			p.name,
            h.cottage,
            h.square_house,
            h.square_area,
            b.floor
        FROM
			advert AS a
			JOIN advert_type_house AS at ON a.id = at.advert_id
			JOIN house AS h ON h.id = at.house_id
			JOIN building AS b ON h.building_id = b.id
			JOIN address AS ad ON b.address_id=ad.id
			JOIN house_name AS hn ON hn.id=ad.house_name_id
			JOIN street AS s ON s.id=hn.street_id
			JOIN town AS t ON t.id=s.town_id
			JOIN province AS p ON p.id=t.province_id
        WHERE a.id = $1
        ORDER BY
            a.created_at DESC;`

	rows, err := r.db.Query(queryBaseAdvert, userId, pageSize, offset)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
		return nil, err
	}
	defer rows.Close()

	rectangleAdverts := []*models.AdvertRectangleData{}

	for rows.Next() {
		var metro, houseName, street, town, province string
		var roomCount int
		rectangleAdvert := &models.AdvertRectangleData{}
		err := rows.Scan(&rectangleAdvert.ID, &rectangleAdvert.Title, &rectangleAdvert.Description, &rectangleAdvert.TypeAdvert,
			&roomCount, &rectangleAdvert.Phone, &rectangleAdvert.TypeSale, &rectangleAdvert.Price,
			&rectangleAdvert.Photo, &rectangleAdvert.DateCreation)

		if err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
			return nil, err
		}

		switch rectangleAdvert.TypeAdvert {
		case string(models.AdvertTypeFlat):
			var squareGeneral float64
			var floor, floorGeneral int
			row := r.db.QueryRowContext(ctx, queryFlat, rectangleAdvert.ID)

			if err := row.Scan(&squareGeneral, &floor, &metro, &houseName, &street, &town, &province, &floorGeneral); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
				return nil, err
			}

			rectangleAdvert.FlatProperties = &models.FlatRectangleProperties{}
			rectangleAdvert.FlatProperties.Floor = floor
			rectangleAdvert.FlatProperties.FloorGeneral = floorGeneral
			rectangleAdvert.FlatProperties.SquareGeneral = squareGeneral
			rectangleAdvert.FlatProperties.RoomCount = roomCount
		case string(models.AdvertTypeHouse):
			var cottage bool
			var squareHouse, squareArea float64
			var floor int
			row := r.db.QueryRowContext(ctx, queryHouse, rectangleAdvert.ID)

			if err := row.Scan(&metro, &houseName, &street, &town, &province, &cottage, &squareHouse, &squareArea, &floor); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
				return nil, err
			}

			rectangleAdvert.HouseProperties = &models.HouseRectangleProperties{}
			rectangleAdvert.HouseProperties.Cottage = cottage
			rectangleAdvert.HouseProperties.SquareHouse = squareHouse
			rectangleAdvert.HouseProperties.SquareArea = squareArea
			rectangleAdvert.HouseProperties.BedroomCount = roomCount
			rectangleAdvert.HouseProperties.Floor = floor
		}

		rectangleAdvert.Address = province + ", " + town + ", " + street + ", " + houseName
		rectangleAdvert.Metro = metro

		rectangleAdverts = append(rectangleAdverts, rectangleAdvert)
	}
	if err := rows.Err(); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod)
	return rectangleAdverts, nil
}

// GetRectangleAdvertsByComplexId retrieves rectangle adverts from the database by complex id.
func (r *AdvertRepo) GetRectangleAdvertsByComplexId(ctx context.Context, pageSize, offset int, complexId int64) ([]*models.AdvertRectangleData, error) {
	queryBaseAdvert := `
        SELECT
			a.id,
			a.title,
			a.description,
			CASE
			   WHEN ath.house_id IS NOT NULL THEN 'House'
			   WHEN atf.flat_id IS NOT NULL THEN 'Flat'
			   ELSE 'None'
		    END AS type_advert, 
            CASE
                WHEN atf.flat_id IS NOT NULL THEN f.bedroom_count
                WHEN ath.house_id IS NOT NULL THEN h.bedroom_count
                ELSE 0
            END AS rcount,
            a.phone,
            a.type_placement,
            pc.price,
            i.photo,
            a.created_at
        FROM
            advert AS a
            LEFT JOIN advert_type_house AS ath ON a.id = ath.advert_id
			LEFT JOIN advert_type_flat AS atf ON a.id = atf.advert_id
            LEFT JOIN flat AS f ON f.id = atf.flat_id
            LEFT JOIN house AS h ON h.id = ath.house_id
            JOIN building AS b ON (f.building_id = b.id OR h.building_id = b.id)
			JOIN address AS ad ON b.address_id=ad.id
			JOIN house_name AS hn ON hn.id=ad.house_name_id
			JOIN street AS s ON s.id=hn.street_id
			JOIN town AS t ON t.id=s.town_id
			JOIN province AS p ON p.id=t.province_id
            LEFT JOIN LATERAL (
                SELECT *
                FROM price_change AS pc
                WHERE pc.advert_id = a.id
                ORDER BY pc.created_at DESC
                LIMIT 1
            ) AS pc ON TRUE
            JOIN image AS i ON i.advert_id = a.id
        WHERE i.priority = (
                SELECT MIN(priority)
                FROM image
                WHERE advert_id = a.id
                    AND is_deleted = FALSE
            )
            AND i.is_deleted = FALSE
            AND a.is_deleted = FALSE
            AND b.complex_id = $1
			ORDER BY a.created_at DESC
			LIMIT $2
			OFFSET $3`
	queryFlat := `
        SELECT
            f.square_general,
            f.floor,
			ad.metro,
			hn.name,
			s.name,
			t.name,
			p.name,
            b.floor AS floorgeneral
        FROM
            advert AS a
            JOIN advert_type_flat AS at ON a.id = at.advert_id
            JOIN flat AS f ON f.id = at.flat_id
            JOIN building AS b ON f.building_id = b.id
			JOIN address AS ad ON b.address_id=ad.id
			JOIN house_name AS hn ON hn.id=ad.house_name_id
			JOIN street AS s ON s.id=hn.street_id
			JOIN town AS t ON t.id=s.town_id
			JOIN province AS p ON p.id=t.province_id
        WHERE a.id = $1
        ORDER BY
            a.created_at DESC;`
	queryHouse := `
        SELECT
			ad.metro,
			hn.name,
			s.name,
			t.name,
			p.name,
            h.cottage,
            h.square_house,
            h.square_area,
            b.floor
        FROM
			advert AS a
			JOIN advert_type_house AS at ON a.id = at.advert_id
			JOIN house AS h ON h.id = at.house_id
			JOIN building AS b ON h.building_id = b.id
			JOIN address AS ad ON b.address_id=ad.id
			JOIN house_name AS hn ON hn.id=ad.house_name_id
			JOIN street AS s ON s.id=hn.street_id
			JOIN town AS t ON t.id=s.town_id
			JOIN province AS p ON p.id=t.province_id
        WHERE a.id = $1
        ORDER BY
            a.created_at DESC;`

	rows, err := r.db.Query(queryBaseAdvert, complexId, pageSize, offset)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
		return nil, err
	}
	defer rows.Close()

	rectangleAdverts := []*models.AdvertRectangleData{}

	for rows.Next() {
		var metro, houseName, street, town, province string
		var roomCount int
		rectangleAdvert := &models.AdvertRectangleData{}
		err := rows.Scan(&rectangleAdvert.ID, &rectangleAdvert.Title, &rectangleAdvert.Description, &rectangleAdvert.TypeAdvert,
			&roomCount, &rectangleAdvert.Phone, &rectangleAdvert.TypeSale, &rectangleAdvert.Price,
			&rectangleAdvert.Photo, &rectangleAdvert.DateCreation)

		if err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
			return nil, err
		}

		switch rectangleAdvert.TypeAdvert {
		case string(models.AdvertTypeFlat):
			var squareGeneral float64
			var floor, floorGeneral int
			row := r.db.QueryRowContext(ctx, queryFlat, rectangleAdvert.ID)

			if err := row.Scan(&squareGeneral, &floor, &metro, &houseName, &street, &town, &province, &floorGeneral); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
				return nil, err
			}

			rectangleAdvert.FlatProperties = &models.FlatRectangleProperties{}
			rectangleAdvert.FlatProperties.Floor = floor
			rectangleAdvert.FlatProperties.FloorGeneral = floorGeneral
			rectangleAdvert.FlatProperties.SquareGeneral = squareGeneral
			rectangleAdvert.FlatProperties.RoomCount = roomCount
		case string(models.AdvertTypeHouse):
			var cottage bool
			var squareHouse, squareArea float64
			var floor int
			row := r.db.QueryRowContext(ctx, queryHouse, rectangleAdvert.ID)

			if err := row.Scan(&metro, &houseName, &street, &town, &province, &cottage, &squareHouse, &squareArea, &floor); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
				return nil, err
			}

			rectangleAdvert.HouseProperties = &models.HouseRectangleProperties{}
			rectangleAdvert.HouseProperties.Cottage = cottage
			rectangleAdvert.HouseProperties.SquareHouse = squareHouse
			rectangleAdvert.HouseProperties.SquareArea = squareArea
			rectangleAdvert.HouseProperties.BedroomCount = roomCount
			rectangleAdvert.HouseProperties.Floor = floor
		}

		rectangleAdvert.Address = province + ", " + town + ", " + street + ", " + houseName
		rectangleAdvert.Metro = metro

		rectangleAdverts = append(rectangleAdverts, rectangleAdvert)
	}
	if err := rows.Err(); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod)
	return rectangleAdverts, nil
}
