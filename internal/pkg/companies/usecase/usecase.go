package usecase

import (
	"2024_1_TeaStealers/internal/pkg/companies"
)

// CompanyUsecase represents the usecase for company using.
type CompanyUsecase struct {
	repo companies.CompanyRepo
}

// NewCompanyUsecase creates a new instance of CompanyUsecase.
func NewCompanyUsecase(repo companies.CompanyRepo) *CompanyUsecase {
	return &CompanyUsecase{repo: repo}
}
