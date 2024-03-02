package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/buildings"
	"2024_1_TeaStealers/internal/pkg/utils"
	"net/http"
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
