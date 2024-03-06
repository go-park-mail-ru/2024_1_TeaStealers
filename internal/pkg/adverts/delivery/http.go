package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts"
	"2024_1_TeaStealers/internal/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satori/uuid"
)

// AdvertHandler handles HTTP requests for manage advert.
type AdvertHandler struct {
	// uc represents the usecase interface for manage advert.
	uc adverts.AdvertUsecase
}

// NewAdvertHandler creates a new instance of AdvertHandler.
func NewAdvertHandler(uc adverts.AdvertUsecase) *AdvertHandler {
	return &AdvertHandler{uc: uc}
}

// @Summary Create a new advert
// @Description Create a new advert
// @Tags adverts
// @Accept json
// @Produce json
// @Param input body models.AdvertCreateData true "Advert data"
// @Success 201 {object} models.Advert
// @Failure 400 {string} string "Incorrect data format"
// @Failure 500 {string} string "Internal server error"
// @Router /adverts/ [post]
func (h *AdvertHandler) CreateAdvert(w http.ResponseWriter, r *http.Request) {
	data := models.AdvertCreateData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}

	newAdvert, err := h.uc.CreateAdvert(r.Context(), &data)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusCreated, newAdvert); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// @Summary Get advert by ID
// @Description Get advert by ID
// @Tags adverts
// @Accept json
// @Produce json
// @Param id path string true "Advert ID"
// @Success 200 {object} models.Advert
// @Failure 400 {string} string "Invalid ID parameter"
// @Failure 404 {string} string "Advert not found"
// @Failure 500 {string} string "Internal server error"
// @Router /adverts/{id} [get]
func (h *AdvertHandler) GetAdvertById(w http.ResponseWriter, r *http.Request) {
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

	advert, err := h.uc.GetAdvertById(r.Context(), advertId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, advert); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// @Summary Get list of adverts
// @Description Get list of adverts
// @Tags adverts
// @Accept json
// @Produce json
// @Success 200 {array} models.Advert
// @Failure 500 {string} string "Internal server error"
// @Router /adverts/list/ [get]
func (h *AdvertHandler) GetAdvertsList(w http.ResponseWriter, r *http.Request) {
	adverts, err := h.uc.GetAdvertsList(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, adverts); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// @Summary Delete advert by ID
// @Description Delete advert by ID
// @Tags adverts
// @Accept json
// @Produce json
// @Param id path string true "Advert ID"
// @Success 200 {string} string "DELETED advert"
// @Failure 400 {string} string "Invalid ID parameter"
// @Failure 500 {string} string "Internal server error"
// @Router /adverts/{id} [delete]
func (h *AdvertHandler) DeleteAdvertById(w http.ResponseWriter, r *http.Request) {
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

	err = h.uc.DeleteAdvertById(r.Context(), advertId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, "DELETED advert by id: "+id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// @Summary Update advert by ID
// @Description Update advert by ID
// @Tags adverts
// @Accept json
// @Produce json
// @Param id path string true "Advert ID"
// @Param body body map[string]interface{} true "Advert data"
// @Success 200 {string} string "UPDATED advert"
// @Failure 400 {string} string "Invalid ID parameter or incorrect data format"
// @Failure 500 {string} string "Internal server error"
// @Router /adverts/{id} [post]
func (h *AdvertHandler) UpdateAdvertById(w http.ResponseWriter, r *http.Request) {
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

	var body map[string]interface{}

	if err = json.NewDecoder(r.Body).Decode(&body); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.uc.UpdateAdvertById(r.Context(), body, advertId); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, "UPDATED advert by id: "+id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}
