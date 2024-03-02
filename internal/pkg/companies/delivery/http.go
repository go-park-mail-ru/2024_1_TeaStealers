package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/companies"
	"2024_1_TeaStealers/internal/pkg/utils"
	"net/http"
)

// CompanyHandler handles HTTP requests for manage company.
type CompanyHandler struct {
	// uc represents the usecase interface for manage company.
	uc companies.CompanyUsecase
}

// NewCompanyHandler creates a new instance of CompanyHandler.
func NewCompanyHandler(uc companies.CompanyUsecase) *CompanyHandler {
	return &CompanyHandler{uc: uc}
}

// CreateCompany handles the request for creating a new company.
func (h *CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	data := models.CompanyCreateData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}

	newCompany, err := h.uc.CreateCompany(r.Context(), &data)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusCreated, newCompany); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}
