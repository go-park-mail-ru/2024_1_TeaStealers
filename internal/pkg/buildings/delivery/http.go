package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/buildings"
	"2024_1_TeaStealers/internal/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/satori/uuid"
	_ "github.com/swaggo/http-swagger"
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

// @Summary Create a building
// @Description Create a new building based on the provided data
// @Accept json
// @Produce json
// @Param buildingData body models.BuildingCreateData true "Data for creating a building"
// @Success 201 {object} models.Building "Returns the created building"
// @Failure 400
// @Router /building/create [post]
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

// @Summary Get building by ID
// @Description Get a building by its ID
// @Tags buildings
// @Accept json
// @Produce json
// @Param id query string true "Building ID"
// @Success 200 {object} models.Building
// @Router /building/get/by/id [get]
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

// @Summary Get list of buildings
// @Description Get a list of buildings
// @Tags buildings
// @Accept json
// @Produce json
// @Success 200 {array} models.Building
// @Router /building/get/list [get]
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

// @Summary Delete building by ID
// @Description Delete a building by its ID
// @Tags buildings
// @Accept json
// @Produce json
// @Param id query string true "Building ID"
// @Success 200 {string} string "Building deleted successfully"
// @Router /building/delete/by/id [delete]
func (h *BuildingHandler) DeleteBuildingById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	advertId, err := uuid.FromString(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	err = h.uc.DeleteBuildingById(r.Context(), advertId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, "DELETED building by id: "+id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// @Summary Update building by ID
// @Description Update a building by its ID
// @Tags buildings
// @Accept json
// @Produce json
// @Param id query string true "Building ID"
// @Param body body map[string]interface{} true "Updated building data"
// @Success 200 {string} string "Building updated successfully"
// @Router /building/update/by/id [put]
// @Router /building/update/by/id [post]
func (h *BuildingHandler) UpdateBuildingById(w http.ResponseWriter, r *http.Request) {
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

	var body map[string]interface{}

	if err = json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.uc.UpdateBuildingById(r.Context(), body, buildingId); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, "UPDATED building by id: "+id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}
