package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/buildings"
	"2024_1_TeaStealers/internal/pkg/utils"
	"net/http"

	"github.com/satori/uuid"
)

// BuildingHandler handles HTTP requests for manage building.
type BuildingHandler struct {
	// uc represents the usecase interface for manage building.
	uc buildings.BuildingUsecase
}

// NewBuildingHandler creates a new instance of BuildingHandler.
func NewBuildingHandler(uc buildings.BuildingUsecase) *BuildingHandler {
	return &BuildingHandler{uc: uc}
}

// CreateBuilding handles the request for creating a new building.
func (h *BuildingHandler) CreateBuilding(w http.ResponseWriter, r *http.Request) {
	data := models.BuildingCreateData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}

	newBuilding, err := h.uc.CreateBuilding(r.Context(), &data)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusCreated, newBuilding); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// GetBuildingById handles the request for retrieving a building by its Id.
func (h *BuildingHandler) GetBuildingById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	buildingId, err := uuid.FromString(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	building, err := h.uc.GetBuildingById(r.Context(), buildingId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, building); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// GetBuildingsList handles the request for retrieving a buildings.
func (h *BuildingHandler) GetBuildingsList(w http.ResponseWriter, r *http.Request) {
	companies, err := h.uc.GetBuildingsList(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, companies); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}
