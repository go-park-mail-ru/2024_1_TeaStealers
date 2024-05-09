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

func (r *ComplexRepo) BeginTx(ctx context.Context) (models.Transaction, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return tx, nil
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

// CreateCompany creates a new company in the database.
func (r *ComplexRepo) CreateCompany(ctx context.Context, company *models.Company) (*models.Company, error) {
	insert := `INSERT INTO company (name, photo, creation_year, phone, description) VALUES ($1, '', $2, $3, $4) RETURNING id`
	if err := r.db.QueryRowContext(ctx, insert, company.Name, company.YearFounded, company.Phone, company.Description).Scan(&company.ID); err != nil {
		log.Println(err)
		return nil, err
	}
	query := `SELECT id, name, creation_year, phone, description FROM company WHERE id = $1`

	res := r.db.QueryRow(query, company.ID)

	newCompany := &models.Company{}
	if err := res.Scan(&newCompany.ID, &newCompany.Name, &newCompany.YearFounded, &newCompany.Phone, &newCompany.Description); err != nil {
		log.Println(err)
		return nil, err
	}

	return newCompany, nil
}

func (r *ComplexRepo) UpdateCompanyPhoto(id int64, fileName string) (string, error) {
	query := `UPDATE company SET photo = $1 WHERE id = $2`
	if _, err := r.db.Exec(query, fileName, id); err != nil {
		log.Println(err)
		return "", err
	}
	return fileName, nil
}

// GetCompanyById ...
func (r *ComplexRepo) GetCompanyById(ctx context.Context, companyId int64) (*models.CompanyData, error) {
	query := `SELECT id, photo, name, creation_year, phone, description FROM company WHERE id = $1`

	companyData := &models.CompanyData{}

	res := r.db.QueryRowContext(ctx, query, companyId)

	if err := res.Scan(&companyData.ID, &companyData.Photo, &companyData.Name, &companyData.YearFounded, &companyData.Phone, &companyData.Description); err != nil {
		return nil, err
	}

	return companyData, nil
}
