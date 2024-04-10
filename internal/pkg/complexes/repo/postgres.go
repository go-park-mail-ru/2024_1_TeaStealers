package repo

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"
	"log"

	"github.com/satori/uuid"
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
	insert := `INSERT INTO complexes (id, companyid, name, adress, photo, description, datebeginbuild, dateendbuild, withoutfinishingoption, finishingoption, prefinishingoption, classhousing, parking, security)
	VALUES ($1, $2, $3, $4, '', $5, $6, $7, $8, $9, $10, $11, $12, $13);`
	if _, err := r.db.ExecContext(ctx, insert, complex.ID, complex.CompanyId, complex.Name,
		complex.Address, complex.Description, complex.DateBeginBuild, complex.DateEndBuild, complex.WithoutFinishingOption,
		complex.FinishingOption, complex.PreFinishingOption, complex.ClassHousing, complex.Parking, complex.Security); err != nil {
		return nil, err
	}
	query := `SELECT id, companyid, name, adress, photo, description, datebeginbuild, dateendbuild, withoutfinishingoption, finishingoption, prefinishingoption, classhousing, parking, security FROM complexes WHERE id = $1`

	res := r.db.QueryRow(query, complex.ID)

	newComplex := &models.Complex{}
	if err := res.Scan(&newComplex.ID, &newComplex.CompanyId, &newComplex.Name, &newComplex.Address, &newComplex.Photo, &newComplex.Description, &newComplex.DateBeginBuild, &newComplex.DateEndBuild, &newComplex.WithoutFinishingOption, &newComplex.FinishingOption, &newComplex.PreFinishingOption, &newComplex.ClassHousing, &newComplex.Parking, &newComplex.Security); err != nil {
		return nil, err
	}

	return newComplex, nil
}

// CreateBuilding creates a new building in the database.
func (r *ComplexRepo) CreateBuilding(ctx context.Context, building *models.Building) (*models.Building, error) {
	insert := `INSERT INTO buildings (id, complexId, floor, material, adress, adressPoint, yearCreation)
	VALUES ($1, $2, $3, $4, $5, $6, $7);`
	log.Println(building)
	if _, err := r.db.ExecContext(ctx, insert, building.ID, building.ComplexID, building.Floor, building.Material, building.Address, building.AddressPoint, building.YearCreation); err != nil {
		return nil, err
	}
	query := `SELECT id, complexId, floor, material, adress, adressPoint, yearCreation FROM buildings WHERE id = $1`

	res := r.db.QueryRow(query, building.ID)

	newBuilding := &models.Building{}
	if err := res.Scan(&newBuilding.ID, &newBuilding.ComplexID, &newBuilding.Floor, &newBuilding.Material, &newBuilding.Address, &newBuilding.AddressPoint, &newBuilding.YearCreation); err != nil {
		return nil, err
	}

	return newBuilding, nil
}

func (r *ComplexRepo) UpdateComplexPhoto(id uuid.UUID, fileName string) (string, error) {
	query := `UPDATE complexes SET photo = $1 WHERE id = $2`
	if _, err := r.db.Exec(query, fileName, id); err != nil {
		log.Println(err)
		return "", err
	}
	return fileName, nil
}

// GetComplexById ...
func (r *ComplexRepo) GetComplexById(ctx context.Context, complexId uuid.UUID) (*models.ComplexData, error) {
	query := `SELECT id, companyid, name, adress, photo, description, datebeginbuild, dateendbuild, withoutfinishingoption, finishingoption, prefinishingoption, classhousing, parking, security FROM complexes WHERE id = $1`

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

// CreateAdvertType creates a new advertType in the database.
func (r *ComplexRepo) CreateAdvertType(ctx context.Context, tx models.Transaction, newAdvertType *models.AdvertType) error {
	insert := `INSERT INTO adverttypes (id, adverttype) VALUES ($1, $2)`
	if _, err := tx.ExecContext(ctx, insert, newAdvertType.ID, newAdvertType.AdvertType); err != nil {
		return err
	}
	return nil
}

// CreateAdvert creates a new advert in the database.
func (r *ComplexRepo) CreateAdvert(ctx context.Context, tx models.Transaction, newAdvert *models.Advert) error {
	insert := `INSERT INTO adverts (id, userid, adverttypeid, adverttypeplacement, title, description, phone, isagent, priority) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	if _, err := tx.ExecContext(ctx, insert, newAdvert.ID, newAdvert.UserID, newAdvert.AdvertTypeID, newAdvert.AdvertTypeSale, newAdvert.Title, newAdvert.Description, newAdvert.Phone, newAdvert.IsAgent, newAdvert.Priority); err != nil {
		return err
	}
	return nil
}

// CreatePriceChange creates a new price change in the database.
func (r *ComplexRepo) CreatePriceChange(ctx context.Context, tx models.Transaction, newPriceChange *models.PriceChange) error {
	insert := `INSERT INTO pricechanges (id, advertid, price) VALUES ($1, $2, $3)`
	if _, err := tx.ExecContext(ctx, insert, newPriceChange.ID, newPriceChange.AdvertID, newPriceChange.Price); err != nil {
		return err
	}
	return nil
}

// CreateHouse creates a new house in the database.
func (r *ComplexRepo) CreateHouse(ctx context.Context, tx models.Transaction, newHouse *models.House) error {
	insert := `INSERT INTO houses (id, buildingid, adverttypeid, ceilingheight, squarearea, squarehouse, bedroomcount, statusarea, cottage, statushome) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	if _, err := tx.ExecContext(ctx, insert, newHouse.ID, newHouse.BuildingID, newHouse.AdvertTypeID, newHouse.CeilingHeight, newHouse.SquareArea, newHouse.SquareHouse, newHouse.BedroomCount, newHouse.StatusArea, newHouse.Cottage, newHouse.StatusHome); err != nil {
		return err
	}
	return nil
}

// CreateFlat creates a new flat in the database.
func (r *ComplexRepo) CreateFlat(ctx context.Context, tx models.Transaction, newFlat *models.Flat) error {
	insert := `INSERT INTO flats (id, buildingid, adverttypeid, floor, ceilingheight, squaregeneral, roomcount, squareresidential, apartament) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	if _, err := tx.ExecContext(ctx, insert, newFlat.ID, newFlat.BuildingID, newFlat.AdvertTypeID, newFlat.Floor, newFlat.CeilingHeight, newFlat.SquareGeneral, newFlat.RoomCount, newFlat.SquareResidential, newFlat.Apartment); err != nil {
		return err
	}
	return nil
}
