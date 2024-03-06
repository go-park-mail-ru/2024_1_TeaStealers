package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/buildings"
	"2024_1_TeaStealers/internal/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

// @Summary Create a new building
// @Description Create a new building
// @Tags buildings
// @Accept json
// @Produce json
// @Param input body models.BuildingCreateData true "Building data"
// @Success 201 {object} models.Building
// @Failure 400 {string} string "Incorrect data format"
// @Failure 500 {string} string "Internal server error"
// @Router /buildings/ [post]
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
// @Description Get building by ID
// @Tags buildings
// @Accept json
// @Produce json
// @Param id path string true "Building ID"
// @Success 200 {object} models.Building
// @Failure 400 {string} string "Invalid ID parameter"
// @Failure 404 {string} string "Building not found"
// @Failure 500 {string} string "Internal server error"
// @Router /buildings/{id} [get]
func (h *BuildingHandler) GetBuildingById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
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
// @Description Get list of buildings
// @Tags buildings
// @Accept json
// @Produce json
// @Success 200 {array} models.Building
// @Failure 500 {string} string "Internal server error"
// @Router /buildings/list/ [get]
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
// @Description Delete building by ID
// @Tags buildings
// @Accept json
// @Produce json
// @Param id path string true "Building ID"
// @Success 200 {string} string "DELETED building"
// @Failure 400 {string} string "Invalid ID parameter"
// @Failure 500 {string} string "Internal server error"
// @Router /buildings/{id} [delete]
func (h *BuildingHandler) DeleteBuildingById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
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
// @Description Update building by ID
// @Tags buildings
// @Accept json
// @Produce json
// @Param id path string true "Building ID"
// @Param body body map[string]interface{} true "Building data"
// @Success 200 {string} string "UPDATED building"
// @Failure 400 {string} string "Invalid ID parameter or incorrect data format"
// @Failure 500 {string} string "Internal server error"
// @Router /buildings/{id} [post]
func (h *BuildingHandler) UpdateBuildingById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
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
