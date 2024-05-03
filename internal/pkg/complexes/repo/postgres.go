package repo

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"
	"log"

	"go.uber.org/zap"
)

// ComplexRepo represents a repository for complex changes.
type ComplexRepo struct {
	db     *sql.DB
	logger *zap.Logger
}

// NewRepository creates a new instance of ComplexRepo.
func NewRepository(db *sql.DB, logger *zap.Logger) *ComplexRepo {
	return &ComplexRepo{db: db, logger: logger}
}

// CreateComplex creates a new complex in the database.
func (r *ComplexRepo) CreateComplex(ctx context.Context, complex *models.Complex) (*models.Complex, error) {
	insert := `INSERT INTO complex (company_id, name, address, photo, description, date_begin_build, date_end_build, without_finishing_option, finishing_option, pre_finishing_option, class_housing, parking, security)
	VALUES ($1, $2, $3, '', $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`
	var id int64
	if err := r.db.QueryRowContext(ctx, insert, complex.CompanyId, complex.Name,
		complex.Address, complex.Description, complex.DateBeginBuild, complex.DateEndBuild, complex.WithoutFinishingOption,
		complex.FinishingOption, complex.PreFinishingOption, complex.ClassHousing, complex.Parking, complex.Security).Scan(&id); err != nil {
		return nil, err
	}
	query := `SELECT company_id, name, address, photo, description, date_begin_build, date_end_build, without_finishing_option, finishing_option, pre_finishing_option, class_housing, parking, security FROM complex WHERE id = $1`

	res := r.db.QueryRow(query, id)

	newComplex := &models.Complex{ID: id}
	if err := res.Scan(&newComplex.CompanyId, &newComplex.Name, &newComplex.Address, &newComplex.Photo, &newComplex.Description, &newComplex.DateBeginBuild, &newComplex.DateEndBuild, &newComplex.WithoutFinishingOption, &newComplex.FinishingOption, &newComplex.PreFinishingOption, &newComplex.ClassHousing, &newComplex.Parking, &newComplex.Security); err != nil {
		return nil, err
	}

	return newComplex, nil
}

// CreateBuilding creates a new building in the database.
func (r *ComplexRepo) CreateBuilding(ctx context.Context, building *models.Building) (*models.Building, error) {
	var id int64
	insert := `INSERT INTO building (complex_id, floor, material_building, address_id, year_creation)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`
	if err := r.db.QueryRowContext(ctx, insert, building.ComplexID, building.Floor, building.Material, building.AddressID, building.YearCreation).Scan(&id); err != nil {
		return nil, err
	}
	query := `SELECT id, complex_id, floor, material_building, address_id, year_creation FROM building WHERE id = $1`

	res := r.db.QueryRow(query, id)

	newBuilding := &models.Building{ID: id}
	if err := res.Scan(&newBuilding.ID, &newBuilding.ComplexID, &newBuilding.Floor, &newBuilding.Material, &newBuilding.AddressID, &newBuilding.YearCreation); err != nil {
		return nil, err
	}

	return newBuilding, nil
}

func (r *ComplexRepo) UpdateComplexPhoto(id int64, fileName string) (string, error) {
	query := `UPDATE complex SET photo = $1 WHERE id = $2`
	if _, err := r.db.Exec(query, fileName, id); err != nil {
		log.Println(err)
		return "", err
	}
	return fileName, nil
}

// GetComplexById ...
func (r *ComplexRepo) GetComplexById(ctx context.Context, complexId int64) (*models.ComplexData, error) {
	query := `SELECT id, company_id, name, address, photo, description, date_begin_build, date_end_build, without_finishing_option, finishing_option, pre_finishing_option, class_housing, parking, security FROM complex WHERE id = $1`

	complexData := &models.ComplexData{}

	res := r.db.QueryRowContext(ctx, query, complexId)

	if err := res.Scan(&complexData.ID, &complexData.CompanyId, &complexData.Name, &complexData.Address, &complexData.Photo,
		&complexData.Description, &complexData.DateBeginBuild, &complexData.DateEndBuild, &complexData.WithoutFinishingOption,
		&complexData.FinishingOption, &complexData.PreFinishingOption, &complexData.ClassHousing, &complexData.Parking, &complexData.Security); err != nil {
		return nil, err
	}

	return complexData, nil
}

func (r *ComplexRepo) BeginTx(ctx context.Context) (models.Transaction, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// CreateAdvertTypeHouse creates a new advertTypeHouse in the database.
func (r *ComplexRepo) CreateAdvertTypeHouse(ctx context.Context, tx models.Transaction, newAdvertType *models.HouseTypeAdvert) error {
	insert := `INSERT INTO advert_type_house (house_id, advert_id) VALUES ($1, $2)`
	if _, err := tx.ExecContext(ctx, insert, newAdvertType.HouseID, newAdvertType.AdvertID); err != nil {
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertTypeMethod, err)
		return err
	}

	// utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertTypeMethod)
	return nil
}

// CreateAdvertTypeFlat creates a new advertTypeFlat in the database.
func (r *ComplexRepo) CreateAdvertTypeFlat(ctx context.Context, tx models.Transaction, newAdvertType *models.FlatTypeAdvert) error {
	insert := `INSERT INTO advert_type_flat (flat_id, advert_id) VALUES ($1, $2)`
	if _, err := tx.ExecContext(ctx, insert, newAdvertType.FlatID, newAdvertType.AdvertID); err != nil {
		// utils.LogError(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertTypeMethod, err)
		return err
	}

	// utils.LogSucces(r.logger, ctx.Value("requestId").(string), utils.RepositoryLayer, adverts.CreateAdvertTypeMethod)
	return nil
}

// CreateAdvert creates a new advert in the database.
func (r *ComplexRepo) CreateAdvert(ctx context.Context, tx models.Transaction, newAdvert *models.Advert) (int64, error) {
	insert := `INSERT INTO advert (user_id, type_placement, title, description, phone, is_agent, priority) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	var id int64
	if err := tx.QueryRowContext(ctx, insert, newAdvert.UserID, newAdvert.AdvertTypeSale, newAdvert.Title, newAdvert.Description, newAdvert.Phone, newAdvert.IsAgent, newAdvert.Priority).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// CreatePriceChange creates a new price change in the database.
func (r *ComplexRepo) CreatePriceChange(ctx context.Context, tx models.Transaction, newPriceChange *models.PriceChange) error {
	insert := `INSERT INTO price_change (advert_id, price) VALUES ($1, $2)`
	if _, err := tx.ExecContext(ctx, insert, newPriceChange.AdvertID, newPriceChange.Price); err != nil {
		return err
	}
	return nil
}

// CreateHouse creates a new house in the database.
func (r *ComplexRepo) CreateHouse(ctx context.Context, tx models.Transaction, newHouse *models.House) (int64, error) {
	insert := `INSERT INTO house (building_id, ceiling_height, square_area, square_house, bedroom_count, status_area_house, cottage, status_home_house) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	var id int64
	if err := tx.QueryRowContext(ctx, insert, newHouse.BuildingID, newHouse.CeilingHeight, newHouse.SquareArea, newHouse.SquareHouse, newHouse.BedroomCount, newHouse.StatusArea, newHouse.Cottage, newHouse.StatusHome).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// CreateFlat creates a new flat in the database.
func (r *ComplexRepo) CreateFlat(ctx context.Context, tx models.Transaction, newFlat *models.Flat) (int64, error) {
	insert := `INSERT INTO flat (building_id, floor, ceiling_height, square_general, bedroom_count, square_residential, apartament) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	var id int64
	if err := tx.QueryRowContext(ctx, insert, newFlat.BuildingID, newFlat.Floor, newFlat.CeilingHeight, newFlat.SquareGeneral, newFlat.RoomCount, newFlat.SquareResidential, newFlat.Apartment).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
