package delivery

import (
	"2024_1_TeaStealers/internal/pkg/companies"
)

// CompanyHandler handles HTTP requests for company changes.
type CompanyHandler struct {
	// uc represents the usecase interface for company changes.
	uc companies.CompanyUsecase
}

// NewCompanyHandler creates a new instance of CompanyHandler.
func NewCompanyHandler(uc companies.CompanyUsecase) *CompanyHandler {
	return &CompanyHandler{uc: uc}
}

//TODO
// Ручка на создание Компании
// Ручка на добавление фото к компании
