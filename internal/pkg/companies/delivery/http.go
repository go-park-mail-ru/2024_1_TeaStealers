package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/companies"
	"2024_1_TeaStealers/internal/pkg/utils"
	"net/http"

	"github.com/satori/uuid"
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

// GetCompanyById handles the request for retrieving a company by its Id.
func (h *CompanyHandler) GetCompanyById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	companyId, err := uuid.FromString(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	company, err := h.uc.GetCompanyById(r.Context(), companyId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, company); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// GetCompaniesList handles the request for retrieving a companies.
func (h *CompanyHandler) GetCompaniesList(w http.ResponseWriter, r *http.Request) {

	companies, err := h.uc.GetCompaniesList(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, companies); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// DeleteCompanyById handles the request for deleting a company by its Id.
func (h *CompanyHandler) DeleteCompanyById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	companyId, err := uuid.FromString(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	err = h.uc.DeleteCompanyById(r.Context(), companyId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, "DELETED company by id: "+id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}
