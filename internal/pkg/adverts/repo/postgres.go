package repo

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts"
	"2024_1_TeaStealers/internal/pkg/utils"
	"context"
	"database/sql"
	"fmt"

	"github.com/satori/uuid"
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

// CreateAdvertType creates a new advertType in the database.
func (r *AdvertRepo) CreateAdvertType(ctx context.Context, tx models.Transaction, newAdvertType *models.AdvertType) error {
	insert := `INSERT INTO adverttypes (id, adverttype) VALUES ($1, $2)`
	if _, err := tx.ExecContext(ctx, insert, newAdvertType.ID, newAdvertType.AdvertType); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertTypeMethod, err)
		return err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertTypeMethod)
	return nil
}

// CreateAdvert creates a new advert in the database.
func (r *AdvertRepo) CreateAdvert(ctx context.Context, tx models.Transaction, newAdvert *models.Advert) error {
	insert := `INSERT INTO adverts (id, userid, adverttypeid, adverttypeplacement, title, description, phone, isagent, priority) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	if _, err := tx.ExecContext(ctx, insert, newAdvert.ID, newAdvert.UserID, newAdvert.AdvertTypeID, newAdvert.AdvertTypeSale, newAdvert.Title, newAdvert.Description, newAdvert.Phone, newAdvert.IsAgent, newAdvert.Priority); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod, err)
		return err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertMethod)
	return nil
}

// CreatePriceChange creates a new price change in the database.
func (r *AdvertRepo) CreatePriceChange(ctx context.Context, tx models.Transaction, newPriceChange *models.PriceChange) error {
	insert := `INSERT INTO pricechanges (id, advertid, price) VALUES ($1, $2, $3)`
	if _, err := tx.ExecContext(ctx, insert, newPriceChange.ID, newPriceChange.AdvertID, newPriceChange.Price); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreatePriceChangeMethod, err)
		return err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreatePriceChangeMethod)
	return nil
}

// CreateBuilding creates a new building in the database.
func (r *AdvertRepo) CreateBuilding(ctx context.Context, tx models.Transaction, newBuilding *models.Building) error {
	insert := `INSERT INTO buildings (id, floor, material, adress, adresspoint, yearcreation) VALUES ($1, $2, $3, $4, $5, $6)`
	if _, err := tx.ExecContext(ctx, insert, newBuilding.ID, newBuilding.Floor, newBuilding.Material, newBuilding.Address, newBuilding.AddressPoint, newBuilding.YearCreation); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateBuildingMethod, err)
		return err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateBuildingMethod)
	return nil
}

// CheckExistsBuilding check exists building.
func (r *AdvertRepo) CheckExistsBuilding(ctx context.Context, adress string) (*models.Building, error) {
	query := `SELECT id FROM buildings WHERE adress = $1`

	building := &models.Building{}

	res := r.db.QueryRowContext(ctx, query, adress)

	if err := res.Scan(&building.ID); err != nil { // Сканируем только id, потому что используем только id, если здание нашлось
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsBuildingMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsBuildingMethod)
	return building, nil
}

// CheckExistsBuildings check exists buildings. Нужна для выпадающего списка существующих зданий по адресу(Для создания объявления)
func (r *AdvertRepo) CheckExistsBuildings(ctx context.Context, pageSize int, adress string) ([]*models.BuildingData, error) {
	query := `SELECT b.id, b.floor, COALESCE(b.material, 'Brick'), b.adress, b.adresspoint, b.yearcreation, COALESCE(cx.name, '') FROM buildings AS b LEFT JOIN complexes AS cx ON b.complexid=cx.id WHERE b.adress ILIKE $1 LIMIT $2`

	rows, err := r.db.Query(query, "%"+adress+"%", pageSize)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsBuildingsMethod, err)
		return nil, err
	}
	defer rows.Close()

	buildings := []*models.BuildingData{}
	for rows.Next() {
		building := &models.BuildingData{}
		err := rows.Scan(&building.ID, &building.Floor, &building.Material, &building.Address, &building.AddressPoint, &building.YearCreation, &building.ComplexName)
		if err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsBuildingsMethod, err)
			return nil, err
		}

		buildings = append(buildings, building)
	}
	if err := rows.Err(); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsBuildingsMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CheckExistsBuildingsMethod)
	return buildings, nil
}

// CreateHouse creates a new house in the database.
func (r *AdvertRepo) CreateHouse(ctx context.Context, tx models.Transaction, newHouse *models.House) error {
	insert := `INSERT INTO houses (id, buildingid, adverttypeid, ceilingheight, squarearea, squarehouse, bedroomcount, statusarea, cottage, statushome) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	if _, err := tx.ExecContext(ctx, insert, newHouse.ID, newHouse.BuildingID, newHouse.AdvertTypeID, newHouse.CeilingHeight, newHouse.SquareArea, newHouse.SquareHouse, newHouse.BedroomCount, newHouse.StatusArea, newHouse.Cottage, newHouse.StatusHome); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateHouseMethod, err)
		return err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateHouseMethod)
	return nil
}

// CreateFlat creates a new flat in the database.
func (r *AdvertRepo) CreateFlat(ctx context.Context, tx models.Transaction, newFlat *models.Flat) error {
	insert := `INSERT INTO flats (id, buildingid, adverttypeid, floor, ceilingheight, squaregeneral, roomcount, squareresidential, apartament) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	if _, err := tx.ExecContext(ctx, insert, newFlat.ID, newFlat.BuildingID, newFlat.AdvertTypeID, newFlat.Floor, newFlat.CeilingHeight, newFlat.SquareGeneral, newFlat.RoomCount, newFlat.SquareResidential, newFlat.Apartment); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateFlatMethod, err)
		return err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateFlatMethod)
	return nil
}

// SelectImages select list images for advert
func (r *AdvertRepo) SelectImages(ctx context.Context, advertId uuid.UUID) ([]*models.ImageResp, error) {
	selectQuery := `SELECT id, photo, priority FROM images WHERE advertid = $1 AND isdeleted = false`
	rows, err := r.db.Query(selectQuery, advertId)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.SelectImagesMethod, err)
		return nil, err
	}
	defer rows.Close()

	images := []*models.ImageResp{}

	for rows.Next() {
		var id uuid.UUID
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
func (r *AdvertRepo) GetTypeAdvertById(ctx context.Context, id uuid.UUID) (*models.AdvertTypeAdvert, error) {
	query := `SELECT at.adverttype FROM adverts AS a JOIN adverttypes AS at ON a.adverttypeid=at.id WHERE a.id = $1`

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
		&advertData.AdvertType,
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
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetHouseAdvertByIdMethod, err)
		return nil, err
	}

	advertData.HouseProperties = &models.HouseProperties{}
	advertData.HouseProperties.CeilingHeight = ceilingHeight
	advertData.HouseProperties.SquareArea = squareArea
	advertData.HouseProperties.SquareHouse = squareHouse
	advertData.HouseProperties.BedroomCount = bedroomCount
	advertData.HouseProperties.StatusArea = statusArea
	advertData.HouseProperties.Cottage = cottage
	advertData.HouseProperties.StatusHome = statusHome
	advertData.HouseProperties.Floor = floor

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
func (r *AdvertRepo) CheckExistsFlat(ctx context.Context, advertId uuid.UUID) (*models.Flat, error) {
	query := `SELECT f.id FROM adverts AS a JOIN adverttypes AS at ON a.adverttypeid=at.id JOIN flats AS f ON f.adverttypeid=at.id WHERE a.id = $1`

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
func (r *AdvertRepo) CheckExistsHouse(ctx context.Context, advertId uuid.UUID) (*models.House, error) {
	query := `SELECT h.id FROM adverts AS a JOIN adverttypes AS at ON a.adverttypeid=at.id JOIN houses AS h ON h.adverttypeid=at.id WHERE a.id = $1;`

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
func (r *AdvertRepo) DeleteFlatAdvertById(ctx context.Context, tx models.Transaction, advertId uuid.UUID) error {
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
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteFlatAdvertByIdMethod, err)
		return err
	}

	queryDeleteAdvertById := `UPDATE adverts SET isdeleted=true WHERE id=$1;`
	queryDeleteAdvertTypeById := `UPDATE adverttypes SET isdeleted=true WHERE id=$1;`
	queryDeleteFlatById := `UPDATE flats SET isdeleted=true WHERE id=$1;`
	queryDeletePriceChanges := `UPDATE pricechanges SET isdeleted=true WHERE advertid=$1;`
	queryDeleteImages := `UPDATE images SET isdeleted=true WHERE advertid=$1;`

	if _, err := tx.Exec(queryDeleteAdvertById, advertId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteFlatAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(queryDeleteAdvertTypeById, advertTypeId); err != nil {
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
func (r *AdvertRepo) DeleteHouseAdvertById(ctx context.Context, tx models.Transaction, advertId uuid.UUID) error {
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
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteHouseAdvertByIdMethod, err)
		return err
	}

	queryDeleteAdvertById := `UPDATE adverts SET isdeleted=true WHERE id=$1;`
	queryDeleteAdvertTypeById := `UPDATE adverttypes SET isdeleted=true WHERE id=$1;`
	queryDeleteHouseById := `UPDATE houses SET isdeleted=true WHERE id=$1;`
	queryDeletePriceChanges := `UPDATE pricechanges SET isdeleted=true WHERE advertid=$1;`
	queryDeleteImages := `UPDATE images SET isdeleted=true WHERE advertid=$1;`

	if _, err := tx.Exec(queryDeleteAdvertById, advertId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.DeleteHouseAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(queryDeleteAdvertTypeById, advertTypeId); err != nil {
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
func (r *AdvertRepo) ChangeTypeAdvert(ctx context.Context, tx models.Transaction, advertId uuid.UUID) (err error) {
	query := `SELECT at.id, at.adverttype FROM adverts AS a JOIN adverttypes AS at ON a.adverttypeid=at.id WHERE a.id = $1;`
	querySelectBuildingIdByFlat := `SELECT b.id AS buildingid, f.id AS flatid  FROM adverts AS a JOIN adverttypes AS at ON at.id=a.adverttypeid JOIN flats AS f ON f.adverttypeid=at.id JOIN buildings AS b ON f.buildingid=b.id WHERE a.id=$1`
	querySelectBuildingIdByHouse := `SELECT b.id AS buildingid, h.id AS houseid  FROM adverts AS a JOIN adverttypes AS at ON at.id=a.adverttypeid JOIN houses AS h ON h.adverttypeid=at.id JOIN buildings AS b ON h.buildingid=b.id WHERE a.id=$1`
	queryInsertFlat := `INSERT INTO flats (id, buildingId, advertTypeId, floor, ceilingHeight, squareGeneral, roomCount, squareResidential, apartament)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`
	queryInsertHouse := `INSERT INTO houses (id, buildingId, advertTypeId, ceilingHeight, squareArea, squareHouse, bedroomCount, statusArea, cottage, statusHome)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`
	queryRestoreFlatById := `UPDATE flats SET isdeleted=false WHERE id=$1;`
	queryRestoreHouseById := `UPDATE houses SET isdeleted=false WHERE id=$1;`
	queryDeleteFlatById := `UPDATE flats SET isdeleted=true WHERE id=$1;`
	queryDeleteHouseById := `UPDATE houses SET isdeleted=true WHERE id=$1;`

	var advertType models.AdvertTypeAdvert
	var advertTypeId uuid.UUID
	res := r.db.QueryRowContext(ctx, query, advertId)

	if err := res.Scan(&advertTypeId, &advertType); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
		return err
	}
	var buildingId uuid.UUID
	switch advertType {
	case models.AdvertTypeFlat:
		res := r.db.QueryRowContext(ctx, querySelectBuildingIdByFlat, advertId)

		var flatId uuid.UUID

		if err := res.Scan(&buildingId, &flatId); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
			return err
		}

		if _, err := tx.Exec(queryDeleteFlatById, flatId); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
			return err
		}

		house, err := r.CheckExistsHouse(ctx, advertId)
		if err != nil {
			house := &models.House{}
			if _, err := tx.Exec(queryInsertHouse, uuid.NewV4(), buildingId, advertTypeId, house.CeilingHeight, house.SquareArea, house.SquareHouse, house.BedroomCount, models.StatusAreaDNP, house.Cottage, models.StatusHomeCompleteNeed); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
				return err
			}
		} else {
			if _, err := tx.Exec(queryRestoreHouseById, house.ID); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
				return err
			}
		}
	case models.AdvertTypeHouse:
		res := r.db.QueryRowContext(ctx, querySelectBuildingIdByHouse, advertId)

		var houseId uuid.UUID

		if err := res.Scan(&buildingId, &houseId); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
			return err
		}

		if _, err := tx.Exec(queryDeleteHouseById, houseId); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
			return err
		}

		flat, err := r.CheckExistsFlat(ctx, advertId)
		if err != nil {
			flat = &models.Flat{}
			if _, err := tx.Exec(queryInsertFlat, uuid.NewV4(), buildingId, advertTypeId, flat.Floor, flat.CeilingHeight, flat.SquareGeneral, flat.RoomCount, flat.SquareResidential, flat.Apartment); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.ChangeTypeAdvertMethod, err)
				return err
			}
		} else {
			if _, err := tx.Exec(queryRestoreFlatById, flat.ID); err != nil {
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
	var price float64
	if err := res.Scan(&advertTypeId, &buildingId, &houseId, &price); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateHouseAdvertByIdMethod, err)
		return err
	}

	queryUpdateAdvertById := `UPDATE adverts SET adverttypeplacement=$1, title=$2, description=$3, phone=$4, isagent=$5 WHERE id=$6;`
	queryUpdateAdvertTypeById := `UPDATE adverttypes SET adverttype=$1 WHERE id=$2;`
	queryUpdateBuildingById := `UPDATE buildings SET floor=$1, material=$2, adress=$3, adresspoint=$4, yearcreation=$5 WHERE id=$6;`
	queryUpdateHouseById := `UPDATE houses SET ceilingheight=$1, squarearea=$2, squarehouse=$3, bedroomcount=$4, statusarea=$5, cottage=$6, statushome=$7 WHERE id=$8;`

	if _, err := tx.Exec(queryUpdateAdvertById, advertUpdateData.TypeSale, advertUpdateData.Title, advertUpdateData.Description, advertUpdateData.Phone, advertUpdateData.IsAgent, advertUpdateData.ID); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateHouseAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(queryUpdateAdvertTypeById, advertUpdateData.TypeAdvert, advertTypeId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateHouseAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(queryUpdateBuildingById, advertUpdateData.HouseProperties.Floor, advertUpdateData.Material, advertUpdateData.Address, advertUpdateData.AddressPoint, advertUpdateData.YearCreation, buildingId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateHouseAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(queryUpdateHouseById, advertUpdateData.HouseProperties.CeilingHeight, advertUpdateData.HouseProperties.SquareArea, advertUpdateData.HouseProperties.SquareHouse, advertUpdateData.HouseProperties.BedroomCount, advertUpdateData.HouseProperties.StatusArea, advertUpdateData.HouseProperties.Cottage, advertUpdateData.HouseProperties.StatusHome, houseId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateHouseAdvertByIdMethod, err)
		return err
	}
	if advertUpdateData.Price != price {
		queryInsertPriceChange := `INSERT INTO pricechanges (id, advertId, price)
            VALUES ($1, $2, $3)`
		if _, err := tx.Exec(queryInsertPriceChange, uuid.NewV4(), advertUpdateData.ID, advertUpdateData.Price); err != nil {
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
	var price float64
	if err := res.Scan(&advertTypeId, &buildingId, &flatId, &price); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateFlatAdvertByIdMethod, err)
		return err
	}

	queryUpdateAdvertById := `UPDATE adverts SET adverttypeplacement=$1, title=$2, description=$3, phone=$4, isagent=$5 WHERE id=$6;`
	queryUpdateAdvertTypeById := `UPDATE adverttypes SET adverttype=$1 WHERE id=$2;`
	queryUpdateBuildingById := `UPDATE buildings SET floor=$1, material=$2, adress=$3, adresspoint=$4, yearcreation=$5 WHERE id=$6;`
	queryUpdateFlatById := `UPDATE flats SET floor=$1, ceilingheight=$2, squaregeneral=$3, roomcount=$4, squareresidential=$5, apartament=$6 WHERE id=$7;`

	if _, err := tx.Exec(queryUpdateAdvertById, advertUpdateData.TypeSale, advertUpdateData.Title, advertUpdateData.Description, advertUpdateData.Phone, advertUpdateData.IsAgent, advertUpdateData.ID); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateFlatAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(queryUpdateAdvertTypeById, advertUpdateData.TypeAdvert, advertTypeId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateFlatAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(queryUpdateBuildingById, advertUpdateData.FlatProperties.FloorGeneral, advertUpdateData.Material, advertUpdateData.Address, advertUpdateData.AddressPoint, advertUpdateData.YearCreation, buildingId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateFlatAdvertByIdMethod, err)
		return err
	}
	if _, err := tx.Exec(queryUpdateFlatById, advertUpdateData.FlatProperties.Floor, advertUpdateData.FlatProperties.CeilingHeight, advertUpdateData.FlatProperties.SquareGeneral, advertUpdateData.FlatProperties.RoomCount, advertUpdateData.FlatProperties.SquareResidential, advertUpdateData.FlatProperties.Apartment, flatId); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateFlatAdvertByIdMethod, err)
		return err
	}

	if advertUpdateData.Price != price {
		queryInsertPriceChange := `INSERT INTO pricechanges (id, advertId, price)
            VALUES ($1, $2, $3)`
		if _, err := tx.Exec(queryInsertPriceChange, uuid.NewV4(), advertUpdateData.ID, advertUpdateData.Price); err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateFlatAdvertByIdMethod, err)
			return err
		}
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.UpdateFlatAdvertByIdMethod)
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
		&advertData.AdvertType,
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
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetFlatAdvertByIdMethod, err)
		return nil, err
	}

	advertData.FlatProperties = &models.FlatProperties{}
	advertData.FlatProperties.CeilingHeight = ceilingHeight
	advertData.FlatProperties.Apartment = apartament.Bool
	advertData.FlatProperties.SquareResidential = squareResidential
	advertData.FlatProperties.RoomCount = roomCount
	advertData.FlatProperties.SquareGeneral = squareGenereal
	advertData.FlatProperties.FloorGeneral = floorGeneral
	advertData.FlatProperties.Floor = floor

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

	rows, err := r.db.Query(queryBaseAdvert, pageSize, offset)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetSquareAdvertsMethod, err)
		return nil, err
	}
	defer rows.Close()

	squareAdverts := []*models.AdvertSquareData{}
	for rows.Next() {
		squareAdvert := &models.AdvertSquareData{}
		err := rows.Scan(&squareAdvert.ID, &squareAdvert.TypeAdvert, &squareAdvert.TypeSale, &squareAdvert.Photo, &squareAdvert.Price, &squareAdvert.DateCreation)
		if err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetSquareAdvertsMethod, err)
			return nil, err
		}
		switch squareAdvert.TypeAdvert {
		case string(models.AdvertTypeFlat):
			var squareGeneral float64
			var floor, floorGeneral, roomCount int
			row := r.db.QueryRowContext(ctx, queryFlat, squareAdvert.ID)
			if err := row.Scan(&squareGeneral, &floor, &squareAdvert.Address, &floorGeneral, &roomCount); err != nil {
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
			if err := row.Scan(&squareAdvert.Address, &cottage, &squareHouse, &squareArea, &bedroomCount, &floor); err != nil {
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
            AND pc.price >= $1
            AND pc.price <= $2
            AND b.adress ILIKE $3`
	queryFlat := `
        SELECT
            f.squaregeneral,
            f.floor,
            b.adress,
            b.floor AS floorgeneral
        FROM
            adverts AS a
            JOIN adverttypes AS at ON a.adverttypeid = at.id
            JOIN flats AS f ON f.adverttypeid = at.id
            JOIN buildings AS b ON f.buildingid = b.id
        WHERE a.id = $1
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
            JOIN houses AS h ON h.adverttypeid = at.id
            JOIN buildings AS b ON h.buildingid = b.id
        WHERE a.id = $1
        ORDER BY
            a.datecreation DESC;`

	pageInfo := &models.PageInfo{}

	var argsForQuery []interface{}
	i := 4 // Изначально в запросе проставлены minPrice, maxPrice и address, поэтому начинаем с 4 для формирования поиска

	advertFilter.Address = "%" + advertFilter.Address + "%"

	if advertFilter.AdvertType != "" {
		queryBaseAdvert += " AND at.adverttype = $" + fmt.Sprint(i) + " "
		argsForQuery = append(argsForQuery, advertFilter.AdvertType)
		i++
	}

	if advertFilter.DealType != "" {
		queryBaseAdvert += " AND a.adverttypeplacement = $" + fmt.Sprint(i) + " "
		argsForQuery = append(argsForQuery, advertFilter.DealType)
		i++
	}

	if advertFilter.RoomCount != 0 {
		queryBaseAdvert = "SELECT * FROM (" + queryBaseAdvert + ") AS subqueryforroomcountcalculate WHERE rcount = $" + fmt.Sprint(i) + " "
		argsForQuery = append(argsForQuery, advertFilter.RoomCount)
		i++
	}

	queryCount := "SELECT COUNT(*) FROM (" + queryBaseAdvert + ") AS subqueryforpaginate"
	queryBaseAdvert += " ORDER BY datecreation DESC LIMIT $" + fmt.Sprint(i) + " OFFSET $" + fmt.Sprint(i+1) + ";"
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
		err := rows.Scan(&rectangleAdvert.ID, &rectangleAdvert.Title, &rectangleAdvert.Description, &rectangleAdvert.TypeAdvert, &roomCount, &rectangleAdvert.Phone, &rectangleAdvert.TypeSale, &rectangleAdvert.Address, &rectangleAdvert.Price, &rectangleAdvert.Photo, &rectangleAdvert.DateCreation)

		if err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsMethod, err)
			return nil, err
		}

		switch rectangleAdvert.TypeAdvert {
		case string(models.AdvertTypeFlat):
			var squareGeneral float64
			var floor, floorGeneral int
			row := r.db.QueryRowContext(ctx, queryFlat, rectangleAdvert.ID)

			if err := row.Scan(&squareGeneral, &floor, &rectangleAdvert.Address, &floorGeneral); err != nil {
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

			if err := row.Scan(&rectangleAdvert.Address, &cottage, &squareHouse, &squareArea, &floor); err != nil {
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
func (r *AdvertRepo) GetRectangleAdvertsByUserId(ctx context.Context, pageSize, offset int, userId uuid.UUID) ([]*models.AdvertRectangleData, error) {
	queryBaseAdvert := `
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
        OFFSET $3;`
	queryFlat := `
        SELECT
            f.squaregeneral,
            f.floor,
            b.adress,
            b.floor AS floorgeneral
        FROM
            adverts AS a
            JOIN adverttypes AS at ON a.adverttypeid = at.id
            JOIN flats AS f ON f.adverttypeid = at.id
            JOIN buildings AS b ON f.buildingid = b.id
        WHERE a.id = $1
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
            JOIN houses AS h ON h.adverttypeid = at.id
            JOIN buildings AS b ON h.buildingid = b.id
        WHERE a.id = $1
        ORDER BY
            a.datecreation DESC;`

	rows, err := r.db.Query(queryBaseAdvert, userId, pageSize, offset)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByUserIdMethod, err)
		return nil, err
	}
	defer rows.Close()

	rectangleAdverts := []*models.AdvertRectangleData{}

	for rows.Next() {
		var roomCount int
		rectangleAdvert := &models.AdvertRectangleData{}
		err := rows.Scan(&rectangleAdvert.ID, &rectangleAdvert.Title, &rectangleAdvert.Description, &rectangleAdvert.TypeAdvert,
			&roomCount, &rectangleAdvert.Phone, &rectangleAdvert.TypeSale, &rectangleAdvert.Address, &rectangleAdvert.Price,
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

			if err := row.Scan(&squareGeneral, &floor, &rectangleAdvert.Address, &floorGeneral); err != nil {
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

			if err := row.Scan(&rectangleAdvert.Address, &cottage, &squareHouse, &squareArea, &floor); err != nil {
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
func (r *AdvertRepo) GetRectangleAdvertsByComplexId(ctx context.Context, pageSize, offset int, complexId uuid.UUID) ([]*models.AdvertRectangleData, error) {
	queryBaseAdvert := `
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
        OFFSET $3;`
	queryFlat := `
        SELECT
            f.squaregeneral,
            f.floor,
            b.adress,
            b.floor AS floorgeneral
        FROM
            adverts AS a
            JOIN adverttypes AS at ON a.adverttypeid = at.id
            JOIN flats AS f ON f.adverttypeid = at.id
            JOIN buildings AS b ON f.buildingid = b.id
        WHERE a.id = $1
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
            JOIN houses AS h ON h.adverttypeid = at.id
            JOIN buildings AS b ON h.buildingid = b.id
        WHERE a.id = $1
        ORDER BY
            a.datecreation DESC;`

	rows, err := r.db.Query(queryBaseAdvert, complexId, pageSize, offset)
	if err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByComplexIdMethod, err)
		return nil, err
	}
	defer rows.Close()

	rectangleAdverts := []*models.AdvertRectangleData{}
	for rows.Next() {
		var roomCount int
		rectangleAdvert := &models.AdvertRectangleData{}
		err := rows.Scan(&rectangleAdvert.ID, &rectangleAdvert.Title, &rectangleAdvert.Description, &rectangleAdvert.TypeAdvert, &roomCount, &rectangleAdvert.Phone, &rectangleAdvert.TypeSale, &rectangleAdvert.Address, &rectangleAdvert.Price, &rectangleAdvert.Photo, &rectangleAdvert.DateCreation)
		if err != nil {
			utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByComplexIdMethod, err)
			return nil, err
		}
		switch rectangleAdvert.TypeAdvert {
		case string(models.AdvertTypeFlat):
			var squareGeneral float64
			var floor, floorGeneral int
			row := r.db.QueryRowContext(ctx, queryFlat, rectangleAdvert.ID)
			if err := row.Scan(&squareGeneral, &floor, &rectangleAdvert.Address, &floorGeneral); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByComplexIdMethod, err)
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
			if err := row.Scan(&rectangleAdvert.Address, &cottage, &squareHouse, &squareArea, &floor); err != nil {
				utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByComplexIdMethod, err)
				return nil, err
			}
			rectangleAdvert.HouseProperties = &models.HouseRectangleProperties{}
			rectangleAdvert.HouseProperties.Cottage = cottage
			rectangleAdvert.HouseProperties.SquareHouse = squareHouse
			rectangleAdvert.HouseProperties.SquareArea = squareArea
			rectangleAdvert.HouseProperties.BedroomCount = roomCount
			rectangleAdvert.HouseProperties.Floor = floor
		}

		rectangleAdverts = append(rectangleAdverts, rectangleAdvert)
	}
	if err := rows.Err(); err != nil {
		utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByComplexIdMethod, err)
		return nil, err
	}

	utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.GetRectangleAdvertsByComplexIdMethod)
	return rectangleAdverts, nil
}
