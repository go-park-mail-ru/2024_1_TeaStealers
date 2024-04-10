package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts"
	"2024_1_TeaStealers/internal/pkg/middleware"
	"2024_1_TeaStealers/internal/pkg/utils"
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/satori/uuid"
	"go.uber.org/zap"
)

const (
	CreateFlatAdvertMethod           = "CreateFlatAdvert"
	CreateHouseAdvertMethod          = "CreateHouseAdvert"
	GetAdvertByIdMethod              = "GetAdvertById"
	UpdateAdvertByIdMethod           = "UpdateAdvertById"
	DeleteAdvertByIdMethod           = "DeleteAdvertById"
	GetSquareAdvertsListMethod       = "GetSquareAdvertsList"
	GetExistBuildingsByAddressMethod = "GetExistBuildingsByAddress"
	GetRectangeAdvertsListMethod     = "GetRectangeAdvertsList"
	GetUserAdvertsMethod             = "GetUserAdverts"
	GetComplexAdvertsMethod          = "GetComplexAdverts"
)

// AdvertHandler handles HTTP requests for advert changes.
type AdvertHandler struct {
	// uc represents the usecase interface for advert changes.
	uc     adverts.AdvertUsecase
	logger *zap.Logger
}

// NewAdvertHandler creates a new instance of AdvertHandler.
func NewAdvertHandler(uc adverts.AdvertUsecase, logger *zap.Logger) *AdvertHandler {
	return &AdvertHandler{uc: uc, logger: logger}
}

func (h *AdvertHandler) CreateFlatAdvert(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("csrftoken")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "csrf cookie not found")
		return
	}
	ctx := context.WithValue(r.Context(), "requestId", uuid.NewV4().String())

	id, ok := ctx.Value(middleware.CookieName).(uuid.UUID)
	if !ok {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, CreateFlatAdvertMethod, errors.New("error with cookie"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}

	data := models.AdvertFlatCreateData{UserID: id}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), CreateFlatAdvertMethod, utils.DeliveryLayer, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}

	newAdvert, err := h.uc.CreateFlatAdvert(ctx, &data)
	newAdvert.Sanitize()

	if err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), CreateFlatAdvertMethod, utils.DeliveryLayer, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusCreated, newAdvert); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), CreateFlatAdvertMethod, utils.DeliveryLayer, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, CreateFlatAdvertMethod)
	}
}

func (h *AdvertHandler) CreateHouseAdvert(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("csrftoken")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "csrf cookie not found")
		return
	}

	ctx := context.WithValue(r.Context(), "requestId", uuid.NewV4().String())

	id, ok := ctx.Value(middleware.CookieName).(uuid.UUID)

  if !ok {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, CreateHouseAdvertMethod, errors.New("error with cookie"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}

	data := models.AdvertHouseCreateData{UserID: id}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, CreateHouseAdvertMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}

	newAdvert, err := h.uc.CreateHouseAdvert(ctx, &data)
	newAdvert.Sanitize()

	if err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, CreateHouseAdvertMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusCreated, newAdvert); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, CreateHouseAdvertMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, CreateHouseAdvertMethod)
	}
}

// GetAdvertById handles the request for getting advert by id
func (h *AdvertHandler) GetAdvertById(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "requestId", uuid.NewV4().String())

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetAdvertByIdMethod, errors.New("error with id advert"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	advertId, err := uuid.FromString(id)
	if err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetAdvertByIdMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	advertData, err := h.uc.GetAdvertById(ctx, advertId)
	advertData.Sanitize()
  
	if err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetAdvertByIdMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, advertData); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetAdvertByIdMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetAdvertByIdMethod)
	}
}

// UpdateAdvertById handles the request for update advert by id
func (h *AdvertHandler) UpdateAdvertById(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("csrftoken")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "csrf cookie not found")
		return
	}
	ctx := context.WithValue(r.Context(), "requestId", uuid.NewV4().String())

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, UpdateAdvertByIdMethod, errors.New("error with id advert"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	advertId, err := uuid.FromString(id)
	if err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, UpdateAdvertByIdMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	data := models.AdvertUpdateData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, UpdateAdvertByIdMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}
	data.Sanitize()

	data.ID = advertId

	err = h.uc.UpdateAdvertById(ctx, &data)
	if err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, UpdateAdvertByIdMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, "advert successfully updated"); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, UpdateAdvertByIdMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, UpdateAdvertByIdMethod)
	}
}

// DeleteAdvertById handles the request for deleting advert by id
func (h *AdvertHandler) DeleteAdvertById(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("csrftoken")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "csrf cookie not found")
		return
	}

	ctx := context.WithValue(r.Context(), "requestId", uuid.NewV4().String())

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, DeleteAdvertByIdMethod, errors.New("error with id advert"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	advertId, err := uuid.FromString(id)
	if err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, DeleteAdvertByIdMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	err = h.uc.DeleteAdvertById(ctx, advertId)
	if err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, DeleteAdvertByIdMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, "Advert successfully deleted"); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, DeleteAdvertByIdMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, DeleteAdvertByIdMethod)
	}
}

// GetSquareAdvertsList handles the request for retrieving a square adverts.
func (h *AdvertHandler) GetSquareAdvertsList(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "requestId", uuid.NewV4().String())

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

	adverts, err := h.uc.GetSquareAdvertsList(ctx, size, offset)
	if err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetSquareAdvertsListMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	for _, adv := range adverts {
		adv.Sanitize()
	}

	if err = utils.WriteResponse(w, http.StatusOK, adverts); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetSquareAdvertsListMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetSquareAdvertsListMethod)
	}
}

// GetExistBuildingsByAddress handles the request for retrieving an existing buildings by address.
func (h *AdvertHandler) GetExistBuildingsByAddress(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "requestId", uuid.NewV4().String())

	pageStr := r.URL.Query().Get("page")
	address := r.URL.Query().Get("address")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 5
	}

	adverts, err := h.uc.GetExistBuildingsByAddress(ctx, address, page)
	if err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetExistBuildingsByAddressMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	for _, adv := range adverts {
		adv.Sanitize()
	}

	if err = utils.WriteResponse(w, http.StatusOK, adverts); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetExistBuildingsByAddressMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetExistBuildingsByAddressMethod)
	}
}

// GetRectangeAdvertsList handles the request for retrieving a rectangle adverts with search.
func (h *AdvertHandler) GetRectangeAdvertsList(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "requestId", uuid.NewV4().String())

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

	adverts, err := h.uc.GetRectangleAdvertsList(ctx, models.AdvertFilter{
		MinPrice:   minPrice,
		MaxPrice:   maxPrice,
		Page:       page,
		Offset:     offset,
		RoomCount:  roomCount,
		Address:    adress,
		DealType:   dealType,
		AdvertType: advertType,
	})
	adverts.Sanitize()

	if err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetRectangeAdvertsListMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, adverts); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetRectangeAdvertsListMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetRectangeAdvertsListMethod)
	}
}

func (h *AdvertHandler) GetUserAdverts(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "requestId", uuid.NewV4().String())

	id := ctx.Value(middleware.CookieName)
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
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod, errors.New("error with id user"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}

	var userAdverts []*models.AdvertRectangleData
	if userAdverts, err = h.uc.GetRectangleAdvertsByUserId(ctx, page, size, UUID); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "error getting user adverts")
		return
	}
	for _, adv := range userAdverts {
		adv.Sanitize()
	}

	if err := utils.WriteResponse(w, http.StatusOK, userAdverts); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	} else {
		utils.LogSuccesResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod)
	}
}

func (h *AdvertHandler) GetComplexAdverts(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "requestId", uuid.NewV4().String())

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
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetComplexAdvertsMethod, errors.New("error with id complex"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	complexId, err := uuid.FromString(id)
	if err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetComplexAdvertsMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var complexAdverts []*models.AdvertRectangleData

	if complexAdverts, err = h.uc.GetRectangleAdvertsByComplexId(ctx, page, size, complexId); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetComplexAdvertsMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "error getting complex adverts")
		return
	}
	for _, adv := range complexAdverts {
		adv.Sanitize()
	}
	if err := utils.WriteResponse(w, http.StatusOK, complexAdverts); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetComplexAdvertsMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	} else {
		utils.LogSuccesResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, GetComplexAdvertsMethod)
	}
}
