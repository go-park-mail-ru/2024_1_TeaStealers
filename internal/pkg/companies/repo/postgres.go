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
func (r *CompanyRepo) CreateCompany(ctx context.Context, user *models.Company) error {
	//insert := `INSERT INTO companies (id, login, password_hash) VALUES ($1, $2, $3, $4)`

	//if _, err := r.db.ExecContext(ctx, insert, user.ID, user.Login, user.PasswordHash); err != nil {
	//	return err
	//}
	return nil
}
