package repo

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"database/sql"
	"fmt"
	"strings"

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
	query := `SELECT id, name, phone, description, data_creation, is_deleted FROM companies WHERE id = $1`

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
	query := `SELECT id, name, phone, description, data_creation, is_deleted FROM companies`
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

// DeleteCompanyById set is_deleted on true on a company from the database by their id.
func (r *CompanyRepo) DeleteCompanyById(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE companies SET is_deleted = true WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

// UpdateCompanyById updates fields company from the database by their id.
func (r *CompanyRepo) UpdateCompanyById(ctx context.Context, body map[string]interface{}, id uuid.UUID) (err error) {
	var updates []string
	var values []interface{}
	i := 1
	for key, value := range body {
		if key == "id" {
			return fmt.Errorf("ID is not changeable")
		}
		updates = append(updates, fmt.Sprintf("%s = $%d", key, i))
		values = append(values, value)
		i++
	}
	values = append(values, id)

	query := fmt.Sprintf("UPDATE companies SET %s WHERE id = $%d", strings.Join(updates, ", "), len(values))
	_, err = r.db.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}
