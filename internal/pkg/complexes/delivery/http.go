package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/complexes"
	"2024_1_TeaStealers/internal/pkg/utils"
	"log"
	"net/http"
	"path/filepath"
	"slices"
	"strings"

	"github.com/gorilla/mux"
	"github.com/satori/uuid"
)

// ComplexHandler handles HTTP requests for complex changes.
type ComplexHandler struct {
	// uc represents the usecase interface for complex changes.
	uc complexes.ComplexUsecase
}

// NewComplexHandler creates a new instance of ComplexHandler.
func NewComplexHandler(uc complexes.ComplexUsecase) *ComplexHandler {
	return &ComplexHandler{uc: uc}
}

func (h *ComplexHandler) CreateComplex(w http.ResponseWriter, r *http.Request) {
	data := models.ComplexCreateData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}
	data.Sanitize()

	newComplex, err := h.uc.CreateComplex(r.Context(), &data)
	newComplex.Sanitize()
	if err != nil {
		log.Println(err)
		utils.WriteError(w, http.StatusBadRequest, "data already is used")
		return
	}

	if err = utils.WriteResponse(w, http.StatusCreated, newComplex); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *ComplexHandler) CreateBuilding(w http.ResponseWriter, r *http.Request) {
	data := models.BuildingCreateData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}
	data.Sanitize()

	newBuilding, err := h.uc.CreateBuilding(r.Context(), &data)
	if err != nil {
		log.Println(err)
		utils.WriteError(w, http.StatusBadRequest, "data already is used")
		return
	}
	newBuilding.Sanitize()

	if err = utils.WriteResponse(w, http.StatusCreated, newBuilding); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *ComplexHandler) UpdateComplexPhoto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	complexId, err := uuid.FromString(id)
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

	fileName, err := h.uc.UpdateComplexPhoto(file, fileType, complexId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "failed upload file")
		return
	}
	if err := utils.WriteResponse(w, http.StatusOK, fileName); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	}
}

// GetComplexById handles the request for getting complex by id
func (h *ComplexHandler) GetComplexById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	complexId, err := uuid.FromString(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	complexData, err := h.uc.GetComplexById(r.Context(), complexId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	complexData.Sanitize()

	if err = utils.WriteResponse(w, http.StatusOK, complexData); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *ComplexHandler) CreateHouseAdvert(w http.ResponseWriter, r *http.Request) {
	data := models.ComplexAdvertHouseCreateData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}
	data.Sanitize()

	newAdvert, err := h.uc.CreateHouseAdvert(r.Context(), &data)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	newAdvert.Sanitize()

	if err = utils.WriteResponse(w, http.StatusCreated, newAdvert); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *ComplexHandler) CreateFlatAdvert(w http.ResponseWriter, r *http.Request) {
	data := models.ComplexAdvertFlatCreateData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}
	data.Sanitize()
	newAdvert, err := h.uc.CreateFlatAdvert(r.Context(), &data)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	newAdvert.Sanitize()

	if err = utils.WriteResponse(w, http.StatusCreated, newAdvert); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}
