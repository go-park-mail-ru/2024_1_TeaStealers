package repo

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"
)

// CompanyRepo represents a repository for companies.
type CompanyRepo struct {
	db *sql.DB
}

// NewRepository creates a new instance of CompanyRepo.
func NewRepository(db *sql.DB) *CompanyRepo {
	return &CompanyRepo{db: db}
}

// CreateCompany creates a new company in the database.
func (r *CompanyRepo) CreateCompany(ctx context.Context, company *models.Company) error {
	insert := `INSERT INTO companies (id, name, phone, description, data_creation, is_deleted) VALUES ($1, $2, $3, $4, $5, $6)`

	if _, err := r.db.ExecContext(ctx, insert, company.ID, company.Name, company.Phone, company.Descpription, company.DataCreation, company.IsDeleted); err != nil {
		return err
	}
	return nil
}
