package repo

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"

	"github.com/satori/uuid"
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

// GetCompanyById retrieves a company from the database by their id.
func (r *CompanyRepo) GetCompanyById(ctx context.Context, id uuid.UUID) (*models.Company, error) {
	query := `SELECT * FROM companies WHERE id = $1`

	res := r.db.QueryRowContext(ctx, query, id)

	company := &models.Company{
		ID: id,
	}
	if err := res.Scan(&company.ID, &company.Name, &company.Phone, &company.Descpription, &company.DataCreation, &company.IsDeleted); err != nil {
		return nil, err
	}

	return company, nil
}

// GetCompaniesList retrieves a companies from the database.
func (r *CompanyRepo) GetCompaniesList(ctx context.Context) ([]*models.Company, error) {
	query := `SELECT * FROM companies`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	companies := []*models.Company{}
	for rows.Next() {
		company := &models.Company{}
		err := rows.Scan(&company.ID, &company.Name, &company.Phone, &company.Descpription, &company.DataCreation, &company.IsDeleted)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}
