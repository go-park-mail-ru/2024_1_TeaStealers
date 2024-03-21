package repo

import (
	"database/sql"
)

// CompanyRepo represents a repository for company changes.
type CompanyRepo struct {
	db *sql.DB
}

// NewRepository creates a new instance of CompanyRepo.
func NewRepository(db *sql.DB) *CompanyRepo {
	return &CompanyRepo{db: db}
}
