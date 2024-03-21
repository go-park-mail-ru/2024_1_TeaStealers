package repo

import (
	"database/sql"
)

// ComplexRepo represents a repository for complex changes.
type ComplexRepo struct {
	db *sql.DB
}

// NewRepository creates a new instance of ComplexRepo.
func NewRepository(db *sql.DB) *ComplexRepo {
	return &ComplexRepo{db: db}
}
