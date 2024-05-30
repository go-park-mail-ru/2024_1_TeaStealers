package repo

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/config/dbPool"
	"2024_1_TeaStealers/internal/pkg/metrics"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"

	"go.uber.org/zap"
)

// ComplexRepo represents a repository for complex changes.
type ComplexRepo struct {
	db       *pgxpool.Pool
	logger   *zap.Logger
	metricsC metrics.MetricsHTTP
}

// NewRepository creates a new instance of ComplexRepo.
func NewRepository(logger *zap.Logger, metrics metrics.MetricsHTTP) *ComplexRepo {
	return &ComplexRepo{db: dbPool.GetDBPool(), logger: logger, metricsC: metrics}
}

func (r *ComplexRepo) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	tx, err := r.db.BeginTx(ctx, txOptions)
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

	var dur time.Duration
	start := time.Now()
	if err := r.db.QueryRow(ctx, insert, complex.CompanyId, complex.Name,
		complex.Address, complex.Description, complex.DateBeginBuild, complex.DateEndBuild, complex.WithoutFinishingOption,
		complex.FinishingOption, complex.PreFinishingOption, complex.ClassHousing, complex.Parking, complex.Security).Scan(&id); err != nil {
		dur = time.Since(start)
		r.metricsC.AddDurationToQueryTimings("CreateComplex", "insert complex", dur)
		r.metricsC.IncreaseExtSystemErr("database", "insert")
		return nil, err
	}
	dur = time.Since(start)
	r.metricsC.AddDurationToQueryTimings("CreateComplex", "insert complex", dur)

	query := `SELECT company_id, name, address, photo, description, date_begin_build, date_end_build, without_finishing_option, finishing_option, pre_finishing_option, class_housing, parking, security FROM complex WHERE id = $1`

	start = time.Now()
	res := r.db.QueryRow(ctx, query, id)
	dur = time.Since(start)
	r.metricsC.AddDurationToQueryTimings("CreateComplex", "select complex", dur)

	newComplex := &models.Complex{ID: id}
	if err := res.Scan(&newComplex.CompanyId, &newComplex.Name, &newComplex.Address, &newComplex.Photo, &newComplex.Description, &newComplex.DateBeginBuild, &newComplex.DateEndBuild, &newComplex.WithoutFinishingOption, &newComplex.FinishingOption, &newComplex.PreFinishingOption, &newComplex.ClassHousing, &newComplex.Parking, &newComplex.Security); err != nil {
		r.metricsC.IncreaseExtSystemErr("database", "select")
		return nil, err
	}

	return newComplex, nil
}

func (r *ComplexRepo) UpdateComplexPhoto(ctx context.Context, id int64, fileName string) (string, error) {
	query := `UPDATE complex SET photo = $1 WHERE id = $2`
	var dur time.Duration
	start := time.Now()
	if _, err := r.db.Exec(ctx, query, fileName, id); err != nil {
		dur = time.Since(start)
		r.metricsC.AddDurationToQueryTimings("UpdateComplexPhoto", "update complex", dur)
		r.metricsC.IncreaseExtSystemErr("database", "update")
		log.Println(err)
		return "", err
	}
	dur = time.Since(start)
	r.metricsC.AddDurationToQueryTimings("UpdateComplexPhoto", "update complex", dur)
	return fileName, nil
}

// GetComplexById ...
func (r *ComplexRepo) GetComplexById(ctx context.Context, complexId int64) (*models.ComplexData, error) {
	query := `SELECT id, company_id, name, address, photo, description, date_begin_build, date_end_build, without_finishing_option, finishing_option, pre_finishing_option, class_housing, parking, security FROM complex WHERE id = $1`

	complexData := &models.ComplexData{}

	var dur time.Duration
	start := time.Now()
	res := r.db.QueryRow(ctx, query, complexId)
	dur = time.Since(start)
	r.metricsC.AddDurationToQueryTimings("GetComplexById", "select complex", dur)

	if err := res.Scan(&complexData.ID, &complexData.CompanyId, &complexData.Name, &complexData.Address, &complexData.Photo,
		&complexData.Description, &complexData.DateBeginBuild, &complexData.DateEndBuild, &complexData.WithoutFinishingOption,
		&complexData.FinishingOption, &complexData.PreFinishingOption, &complexData.ClassHousing, &complexData.Parking, &complexData.Security); err != nil {
		r.metricsC.IncreaseExtSystemErr("database", "select")
		return nil, err
	}

	return complexData, nil
}

// CreateCompany creates a new company in the database.
func (r *ComplexRepo) CreateCompany(ctx context.Context, company *models.Company) (*models.Company, error) {
	insert := `INSERT INTO company (name, photo, creation_year, phone, description) VALUES ($1, '', $2, $3, $4) RETURNING id`
	var dur time.Duration
	start := time.Now()
	if err := r.db.QueryRow(ctx, insert, company.Name, company.YearFounded, company.Phone, company.Description).Scan(&company.ID); err != nil {
		dur = time.Since(start)
		r.metricsC.AddDurationToQueryTimings("CreateCompany", "insert company", dur)
		r.metricsC.IncreaseExtSystemErr("database", "inser")
		log.Println(err)
		return nil, err
	}
	dur = time.Since(start)
	r.metricsC.AddDurationToQueryTimings("CreateCompany", "insert company", dur)

	query := `SELECT id, name, creation_year, phone, description FROM company WHERE id = $1`

	start = time.Now()
	res := r.db.QueryRow(ctx, query, company.ID)
	dur = time.Since(start)
	r.metricsC.AddDurationToQueryTimings("CreateCompany", "select company", dur)

	newCompany := &models.Company{}
	if err := res.Scan(&newCompany.ID, &newCompany.Name, &newCompany.YearFounded, &newCompany.Phone, &newCompany.Description); err != nil {
		r.metricsC.IncreaseExtSystemErr("database", "select")
		log.Println(err)
		return nil, err
	}

	return newCompany, nil
}

func (r *ComplexRepo) UpdateCompanyPhoto(ctx context.Context, id int64, fileName string) (string, error) {
	query := `UPDATE company SET photo = $1 WHERE id = $2`

	var dur time.Duration
	start := time.Now()
	if _, err := r.db.Exec(ctx, query, fileName, id); err != nil {
		dur = time.Since(start)
		r.metricsC.AddDurationToQueryTimings("UpdateCompanyPhoto", "update company", dur)
		r.metricsC.IncreaseExtSystemErr("database", "update")
		log.Println(err)
		return "", err
	}
	dur = time.Since(start)
	r.metricsC.AddDurationToQueryTimings("UpdateCompanyPhoto", "update company", dur)

	return fileName, nil
}

// GetCompanyById ...
func (r *ComplexRepo) GetCompanyById(ctx context.Context, companyId int64) (*models.CompanyData, error) {
	query := `SELECT id, photo, name, creation_year, phone, description FROM company WHERE id = $1`

	companyData := &models.CompanyData{}

	var dur time.Duration
	start := time.Now()
	res := r.db.QueryRow(ctx, query, companyId)
	dur = time.Since(start)
	r.metricsC.AddDurationToQueryTimings("GetCompanyById", "select company", dur)

	if err := res.Scan(&companyData.ID, &companyData.Photo, &companyData.Name, &companyData.YearFounded, &companyData.Phone, &companyData.Description); err != nil {
		r.metricsC.IncreaseExtSystemErr("database", "select")
		return nil, err
	}

	return companyData, nil
}
