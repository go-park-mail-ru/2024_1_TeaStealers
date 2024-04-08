package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts"
	"2024_1_TeaStealers/internal/pkg/middleware"
	"2024_1_TeaStealers/internal/pkg/utils"
	"log"
	"net/http"
	"strconv"

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
	id, ok := r.Context().Value(middleware.CookieName).(uuid.UUID)
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}

	data := models.AdvertFlatCreateData{UserID: id}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
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
	id, ok := r.Context().Value(middleware.CookieName).(uuid.UUID)
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}

	data := models.AdvertHouseCreateData{UserID: id}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
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
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, advertData); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// UpdateAdvertById handles the request for update advert by id
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

	data := models.AdvertUpdateData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}

	data.ID = advertId

	err = h.uc.UpdateAdvertById(r.Context(), &data)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, "advert successfully updated"); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// DeleteAdvertById handles the request for deleting advert by id
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
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, "Advert successfully deleted"); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// GetSquareAdvertsList handles the request for retrieving a square adverts.
func (h *AdvertHandler) GetSquareAdvertsList(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 10
	}

	offset := (page - 1) * size

	adverts, err := h.uc.GetSquareAdvertsList(r.Context(), size, offset)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, adverts); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// GetExistBuildingsByAddress handles the request for retrieving an existing buildings by address.
func (h *AdvertHandler) GetExistBuildingsByAddress(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	address := r.URL.Query().Get("address")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 5
	}

	adverts, err := h.uc.GetExistBuildingsByAddress(r.Context(), address, page)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, adverts); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// GetRectangeAdvertsList handles the request for retrieving a rectangle adverts with search.
func (h *AdvertHandler) GetRectangeAdvertsList(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")
	advertType := r.URL.Query().Get("adverttype") // House/Advert
	minPriceStr := r.URL.Query().Get("minprice")
	maxPriceStr := r.URL.Query().Get("maxprice")
	dealType := r.URL.Query().Get("dealtype") // Sale/Rent
	roomCountStr := r.URL.Query().Get("roomcount")
	adress := r.URL.Query().Get("adress")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1000000
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 0
	}
	roomCount, err := strconv.Atoi(roomCountStr)
	if err != nil {
		roomCount = 0
	}
	minPrice, err := strconv.ParseInt(minPriceStr, 10, 64)
	if err != nil {
		minPrice = 0
	}
	maxPrice, err := strconv.ParseInt(maxPriceStr, 10, 64)
	if err != nil {
		maxPrice = 1000000000
	}

	if advertType != "House" && advertType != "Flat" {
		advertType = ""
	}

	if dealType != "Sale" && dealType != "Rent" {
		dealType = ""
	}

	offset := (page - 1) * size

	adverts, err := h.uc.GetRectangleAdvertsList(r.Context(), models.AdvertFilter{
		MinPrice:   minPrice,
		MaxPrice:   maxPrice,
		Page:       page,
		Offset:     offset,
		RoomCount:  roomCount,
		Address:    adress,
		DealType:   dealType,
		AdvertType: advertType,
	})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, adverts); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *AdvertHandler) GetUserAdverts(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.CookieName)
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1000000
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 0
	}

	UUID, ok := id.(uuid.UUID)
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}

	var userAdverts []*models.AdvertRectangleData
	if userAdverts, err = h.uc.GetRectangleAdvertsByUserId(r.Context(), page, size, UUID); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "error getting user adverts")
		return
	}

	if err := utils.WriteResponse(w, http.StatusOK, userAdverts); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	}
}

func (h *AdvertHandler) GetComplexAdverts(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1000000
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 0
	}

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

	var complexAdverts []*models.AdvertRectangleData

	if complexAdverts, err = h.uc.GetRectangleAdvertsByComplexId(r.Context(), page, size, complexId); err != nil {
		log.Println(err)
		utils.WriteError(w, http.StatusBadRequest, "error getting complex adverts")
		return
	}

	if err := utils.WriteResponse(w, http.StatusOK, complexAdverts); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	}
}
