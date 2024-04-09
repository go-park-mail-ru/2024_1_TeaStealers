package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/companies"
	"2024_1_TeaStealers/internal/pkg/utils"
	"net/http"
	"path/filepath"
	"slices"
	"strings"

	"github.com/gorilla/mux"
	"github.com/satori/uuid"
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

func (h *CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("csrftoken")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "csrf cookie not found")
		return
	}
	data := models.CompanyCreateData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}
	data.Sanitize()

	newCompany, err := h.uc.CreateCompany(r.Context(), &data)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "data already is used")
		return
	}
	newCompany.Sanitize()

	if err = utils.WriteResponse(w, http.StatusCreated, newCompany); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *CompanyHandler) UpdateCompanyPhoto(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("csrftoken")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "csrf cookie not found")
		return
	}
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
	if err := r.ParseMultipartForm(5 << 20); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "max size file 5 mb")
		return
	}

	file, head, err := r.FormFile("file")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "bad data request")
		return
	}
	defer file.Close()

	allowedExtensions := []string{".jpg", ".jpeg", ".png"}
	fileType := strings.ToLower(filepath.Ext(head.Filename))
	if !slices.Contains(allowedExtensions, fileType) {
		utils.WriteError(w, http.StatusBadRequest, "jpg, jpeg, png only")
		return
	}

	fileName, err := h.uc.UpdateCompanyPhoto(file, fileType, companyId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "failed upload file")
		return
	}
	if err := utils.WriteResponse(w, http.StatusOK, fileName); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	}
}

// GetCompanyById handles the request for getting company by id
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

	companyData, err := h.uc.GetCompanyById(r.Context(), companyId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	companyData.Sanitize()

	if err = utils.WriteResponse(w, http.StatusOK, companyData); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}
