package repo

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"
	"log"

	"go.uber.org/zap"
)

// CompanyRepo represents a repository for company changes.
type CompanyRepo struct {
	db     *sql.DB
	logger *zap.Logger
}

// NewRepository creates a new instance of CompanyRepo.
func NewRepository(db *sql.DB, logger *zap.Logger) *CompanyRepo {
	return &CompanyRepo{db: db, logger: logger}
}

// CreateCompany creates a new company in the database.
func (r *CompanyRepo) CreateCompany(ctx context.Context, company *models.Company) (*models.Company, error) {
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

func (r *CompanyRepo) UpdateCompanyPhoto(id int64, fileName string) (string, error) {
	query := `UPDATE company SET photo = $1 WHERE id = $2`
	if _, err := r.db.Exec(query, fileName, id); err != nil {
		log.Println(err)
		return "", err
	}
	return fileName, nil
}

// GetCompanyById ...
func (r *CompanyRepo) GetCompanyById(ctx context.Context, companyId int64) (*models.CompanyData, error) {
	query := `SELECT id, photo, name, creation_year, phone, description FROM company WHERE id = $1`

	companyData := &models.CompanyData{}

	res := r.db.QueryRowContext(ctx, query, companyId)

	if err := res.Scan(&companyData.ID, &companyData.Photo, &companyData.Name, &companyData.YearFounded, &companyData.Phone, &companyData.Description); err != nil {
		return nil, err
	}

	return companyData, nil
}
