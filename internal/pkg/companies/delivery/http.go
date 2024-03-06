package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/companies"
	"2024_1_TeaStealers/internal/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

// @Summary Create a new company
// @Description Create a new company
// @Tags companies
// @Accept json
// @Produce json
// @Param input body models.CompanyCreateData true "Company data"
// @Success 201 {object} models.Company
// @Failure 400 {string} string "Incorrect data format"
// @Failure 500 {string} string "Internal server error"
// @Router /companies [post]
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

// @Summary Get company by ID
// @Description Get company by ID
// @Tags companies
// @Accept json
// @Produce json
// @Param id path string true "Company ID"
// @Success 200 {object} models.Company
// @Failure 400 {string} string "Invalid ID parameter"
// @Failure 404 {string} string "Company not found"
// @Failure 500 {string} string "Internal server error"
// @Router /companies/{id} [get]
func (h *CompanyHandler) GetCompanyById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
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

// @Summary Get list of companies
// @Description Get list of companies
// @Tags companies
// @Accept json
// @Produce json
// @Success 200 {array} models.Company
// @Failure 500 {string} string "Internal server error"
// @Router /companies/list/ [get]
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

// @Summary Delete company by ID
// @Description Delete company by ID
// @Tags companies
// @Accept json
// @Produce json
// @Param id path string true "Company ID"
// @Success 200 {string} string "DELETED company"
// @Failure 400 {string} string "Invalid ID parameter"
// @Failure 500 {string} string "Internal server error"
// @Router /companies/{id} [delete]
func (h *CompanyHandler) DeleteCompanyById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
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

// @Summary Update company by ID
// @Description Update company by ID
// @Tags companies
// @Accept json
// @Produce json
// @Param id path string true "Company ID"
// @Param body body map[string]interface{} true "Company data"
// @Success 200 {string} string "UPDATED company"
// @Failure 400 {string} string "Invalid ID parameter or incorrect data format"
// @Failure 500 {string} string "Internal server error"
// @Router /companies/{id} [post]
func (h *CompanyHandler) UpdateCompanyById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	companyId, err := uuid.FromString(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var body map[string]interface{}

	if err = json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.uc.UpdateCompanyById(r.Context(), body, companyId); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, "UPDATED company by id: "+id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}
