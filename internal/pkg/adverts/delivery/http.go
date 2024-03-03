package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts"
	"2024_1_TeaStealers/internal/pkg/utils"
	"net/http"

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

// CreateAdvert handles the request for creating a new advert.
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

// GetAdverById handles the request for retrieving a advert by its Id.
func (h *AdvertHandler) GetAdvertById(w http.ResponseWriter, r *http.Request) {
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

	advert, err := h.uc.GetAdvertById(r.Context(), advertId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, advert); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// GetAdvertsList handles the request for retrieving an adverts.
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
