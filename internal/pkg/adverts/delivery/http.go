package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts"
	"2024_1_TeaStealers/internal/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satori/uuid"
)

// AdvertHandler handles HTTP requests for advert changes.
type AdvertHandler struct {
	// uc represents the usecase interface for advert changes.
	uc adverts.AdvertUsecase
}

// NewAdvertHandler creates a new instance of AdvertHandler.
func NewAdvertHandler(uc adverts.AdvertUsecase) *AdvertHandler {
	return &AdvertHandler{uc: uc}
}

func (h *AdvertHandler) CreateFlatAdvert(w http.ResponseWriter, r *http.Request) {
	data := models.AdvertFlatCreateData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Incorrect data format")
		return
	}

	newAdvert, err := h.uc.CreateFlatAdvert(r.Context(), &data)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusCreated, newAdvert); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *AdvertHandler) CreateHouseAdvert(w http.ResponseWriter, r *http.Request) {
	data := models.AdvertHouseCreateData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Incorrect data format")
		return
	}

	newAdvert, err := h.uc.CreateHouseAdvert(r.Context(), &data)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusCreated, newAdvert); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// GetHouseSquareAdvertsList handles the request for retrieving a square house adverts.
func (h *AdvertHandler) GetHouseSquareAdvertsList(w http.ResponseWriter, r *http.Request) {
	adverts, err := h.uc.GetHouseSquareAdvertsList(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, adverts); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// GetFlatSquareAdvertsList handles the request for retrieving a square flat adverts.
func (h *AdvertHandler) GetFlatSquareAdvertsList(w http.ResponseWriter, r *http.Request) {
	adverts, err := h.uc.GetFlatSquareAdvertsList(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, adverts); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// GetFlatRectangleAdvertsList handles the request for retrieving a rectangle flat adverts.
func (h *AdvertHandler) GetFlatRectangleAdvertsList(w http.ResponseWriter, r *http.Request) {
	adverts, err := h.uc.GetFlatRectangleAdvertsList(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, adverts); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// GetHouseRectangleAdvertsList handles the request for retrieving a rectangle house adverts.
func (h *AdvertHandler) GetHouseRectangleAdvertsList(w http.ResponseWriter, r *http.Request) {
	adverts, err := h.uc.GetHouseRectangleAdvertsList(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, adverts); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// GetAdvertById handles the request for getting advert by id
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

	advertData, err := h.uc.GetAdvertById(r.Context(), advertId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, advertData); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}
