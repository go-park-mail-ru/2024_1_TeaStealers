package repo

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"
	"log"

	"github.com/satori/uuid"
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
	insert := `INSERT INTO companies (id, name, photo, yearFounded, phone, description) VALUES ($1, $2, '', $3, $4, $5);`
	if _, err := r.db.ExecContext(ctx, insert, company.ID, company.Name, company.YearFounded, company.Phone, company.Description); err != nil {
		return nil, err
	}
	query := `SELECT id, name, yearFounded, phone, description FROM companies WHERE id = $1`

	res := r.db.QueryRow(query, company.ID)

	newCompany := &models.Company{}
	if err := res.Scan(&newCompany.ID, &newCompany.Name, &newCompany.YearFounded, &newCompany.Phone, &newCompany.Description); err != nil {
		return nil, err
	}

	return newCompany, nil
}

func (r *CompanyRepo) UpdateCompanyPhoto(id uuid.UUID, fileName string) (string, error) {
	query := `UPDATE companies SET photo = $1 WHERE id = $2`
	if _, err := r.db.Exec(query, fileName, id); err != nil {
		log.Println(err)
		return "", err
	}
	return fileName, nil
}

// GetCompanyById ...
func (r *CompanyRepo) GetCompanyById(ctx context.Context, companyId uuid.UUID) (*models.CompanyData, error) {
	query := `SELECT id, photo, name, yearfounded, phone, description FROM companies WHERE id = $1`

	companyData := &models.CompanyData{}

	res := r.db.QueryRowContext(ctx, query, companyId)

	if err := res.Scan(&companyData.ID, &companyData.Photo, &companyData.Name, &companyData.YearFounded, &companyData.Phone, &companyData.Description); err != nil {
		return nil, err
	}

	return companyData, nil
}
