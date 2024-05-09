package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts"
	"2024_1_TeaStealers/internal/pkg/adverts/delivery/grpc/gen"
	genAdverts "2024_1_TeaStealers/internal/pkg/adverts/delivery/grpc/gen"
	"2024_1_TeaStealers/internal/pkg/middleware"
	"2024_1_TeaStealers/internal/pkg/utils"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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
type AdvertsClientHandler struct {
	// uc represents the usecase interface for advert changes.
	client genAdverts.AdvertsClient
	uc     adverts.AdvertUsecase
	logger *zap.Logger
}

// NewAdvertHandler creates a new instance of AdvertHandler.
func NewAdvertsClientHandler(grpcConn *grpc.ClientConn, uc adverts.AdvertUsecase, logger *zap.Logger) *AdvertsClientHandler {
	return &AdvertsClientHandler{client: genAdverts.NewAdvertsClient(grpcConn), uc: uc, logger: logger}
}

func (h *AdvertsClientHandler) CreateFlatAdvert(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("csrftoken")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "csrf cookie not found")
		return
	}

	id, ok := r.Context().Value(middleware.CookieName).(int64)
	if !ok {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateFlatAdvertMethod, errors.New("error with cookie"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}

	data := models.AdvertFlatCreateData{UserID: id}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), CreateFlatAdvertMethod, utils.DeliveryLayer, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}

	newAdvert, err := h.uc.CreateFlatAdvert(r.Context(), &data)
	newAdvert.Sanitize()

	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), CreateFlatAdvertMethod, utils.DeliveryLayer, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusCreated, newAdvert); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), CreateFlatAdvertMethod, utils.DeliveryLayer, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateFlatAdvertMethod)
	}
}

func (h *AdvertsClientHandler) CreateHouseAdvert(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("csrftoken")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "csrf cookie not found")
		return
	}

	id, ok := r.Context().Value(middleware.CookieName).(int64)

	if !ok {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateHouseAdvertMethod, errors.New("error with cookie"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}

	data := models.AdvertHouseCreateData{UserID: id}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateHouseAdvertMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}

	newAdvert, err := h.uc.CreateHouseAdvert(r.Context(), &data)
	newAdvert.Sanitize()

	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateHouseAdvertMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusCreated, newAdvert); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateHouseAdvertMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateHouseAdvertMethod)
	}
}

// GetAdvertById handles the request for getting advert by id
func (h *AdvertsClientHandler) GetAdvertById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetAdvertByIdMethod, errors.New("error with id advert"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	advertId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetAdvertByIdMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	advertDataResponse, err := h.client.GetAdvertById(r.Context(), &genAdverts.GetAdvertByIdRequest{Id: advertId})
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetAdvertByIdMethod, err, http.StatusConflict)
		utils.WriteError(w, http.StatusConflict, err.Error())
		return
	}

	var priceHistory []*models.PriceChangeData
	for _, pcd := range advertDataResponse.PriceHistory {
		date := pcd.DateCreation[:19]
		dateTime, err := utils.StringToTime("2006-01-02 15:04:05", date)
		if err != nil {
			// utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, LoginMethod, err, http.StatusInternalServerError)
			utils.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}

		priceHistory = append(priceHistory, &models.PriceChangeData{Price: pcd.Price, DateCreation: dateTime})
	}

	var images []*models.ImageResp
	for _, img := range advertDataResponse.Images {
		images = append(images, &models.ImageResp{ID: img.Id, Photo: img.Photo, Priority: int(img.Priority)})
	}

	var material models.MaterialBuilding

	switch advertDataResponse.Material {
	case gen.MaterialBuilding_MATERIAL_BRICK:
		material = models.MaterialBrick
	case gen.MaterialBuilding_MATERIAL_MONOLITHIC:
		material = models.MaterialMonolithic
	case gen.MaterialBuilding_MATERIAL_WOOD:
		material = models.MaterialWood
	case gen.MaterialBuilding_MATERIAL_PANEL:
		material = models.MaterialPanel
	case gen.MaterialBuilding_MATERIAL_STALINSKY:
		material = models.MaterialStalinsky
	case gen.MaterialBuilding_MATERIAL_BLOCK:
		material = models.MaterialBlock
	case gen.MaterialBuilding_MATERIAL_MONOLITHIC_BLOCK:
		material = models.MaterialMonolithicBlock
	case gen.MaterialBuilding_MATERIAL_FRAME:
		material = models.MaterialFrame
	case gen.MaterialBuilding_MATERIAL_AERATED_CONCRETE_BLOCK:
		material = models.MaterialAeratedConcreteBlock
	case gen.MaterialBuilding_MATERIAL_GAS_SILICATE_BLOCK:
		material = models.MaterialGasSilicateBlock
	case gen.MaterialBuilding_MATERIAL_FOAM_CONCRETE_BLOCK:
		material = models.MaterialFoamConcreteBlock
	}

	var statusArea models.StatusAreaHouse
	var statusHome models.StatusHomeHouse
	var houseProperties *models.HouseProperties = nil
	var flatProperties *models.FlatProperties = nil
	var complexProperties *models.ComplexAdvertProperties = nil

	if advertDataResponse.HouseProperties != nil {
		switch advertDataResponse.HouseProperties.StatusArea {
		case gen.StatusAreaHouse_STATUS_AREA_IHC:
			statusArea = "IHC"
		case gen.StatusAreaHouse_STATUS_AREA_DNP:
			statusArea = "DNP"
		case gen.StatusAreaHouse_STATUS_AREA_G:
			statusArea = "G"
		case gen.StatusAreaHouse_STATUS_AREA_F:
			statusArea = "F"
		case gen.StatusAreaHouse_STATUS_AREA_PSP:
			statusArea = "PSP"
		}

		switch advertDataResponse.HouseProperties.StatusHome {
		case gen.StatusHomeHouse_STATUS_HOME_LIVE:
			statusHome = "Live"
		case gen.StatusHomeHouse_STATUS_HOME_REPAIR_NEED:
			statusHome = "RepairNeed"
		case gen.StatusHomeHouse_STATUS_HOME_COMPLETE_NEED:
			statusHome = "CompleteNeed"
		case gen.StatusHomeHouse_STATUS_HOME_RENOVATION:
			statusHome = "Renovation"
		}

		houseProperties = &models.HouseProperties{CeilingHeight: advertDataResponse.HouseProperties.CeilingHeight, SquareArea: advertDataResponse.HouseProperties.SquareArea, SquareHouse: advertDataResponse.HouseProperties.SquareHouse, BedroomCount: int(advertDataResponse.HouseProperties.BedroomCount), StatusArea: statusArea, Cottage: advertDataResponse.HouseProperties.Cottage, StatusHome: statusHome, Floor: int(advertDataResponse.HouseProperties.Floor)}
	}

	if advertDataResponse.FlatProperties != nil {
		log.Println(advertDataResponse.FlatProperties)
		flatProperties = &models.FlatProperties{CeilingHeight: advertDataResponse.FlatProperties.CeilingHeight, RoomCount: int(advertDataResponse.FlatProperties.RoomCount), FloorGeneral: int(advertDataResponse.FlatProperties.FloorGeneral), Apartment: advertDataResponse.FlatProperties.Apartment, SquareGeneral: advertDataResponse.FlatProperties.SquareGeneral, Floor: int(advertDataResponse.FlatProperties.Floor), SquareResidential: advertDataResponse.FlatProperties.SquareResidential}
	}

	if advertDataResponse.ComplexProperties != nil {
		complexProperties = &models.ComplexAdvertProperties{ComplexId: advertDataResponse.ComplexProperties.ComplexId, NameComplex: advertDataResponse.ComplexProperties.NameComplex, PhotoCompany: advertDataResponse.ComplexProperties.PhotoCompany, NameCompany: advertDataResponse.ComplexProperties.NameCompany}
	}

	date := advertDataResponse.DateCreation[:19]
	dateTime, err := utils.StringToTime("2006-01-02 15:04:05", date)
	if err != nil {
		// utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, LoginMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	advert := &models.AdvertData{ID: advertDataResponse.Id, AdvertType: advertDataResponse.AdvertType, TypeSale: advertDataResponse.TypeSale, Title: advertDataResponse.Title, Description: advertDataResponse.Description, CountViews: advertDataResponse.CountViews, CountLikes: advertDataResponse.CountLikes, Price: advertDataResponse.Price, Phone: advertDataResponse.Phone, IsLiked: advertDataResponse.IsLiked, IsAgent: advertDataResponse.IsAgent, Metro: advertDataResponse.Metro, Address: advertDataResponse.Address, AddressPoint: advertDataResponse.AddressPoint, YearCreation: int(advertDataResponse.YearCreation), PriceChange: priceHistory, Images: images, FlatProperties: flatProperties, HouseProperties: houseProperties, Material: material, ComplexProperties: complexProperties, DateCreation: dateTime}

	if err = utils.WriteResponse(w, http.StatusOK, advert); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetAdvertByIdMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetAdvertByIdMethod)
	}
}

// UpdateAdvertById handles the request for update advert by id
func (h *AdvertsClientHandler) UpdateAdvertById(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("csrftoken")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "csrf cookie not found")
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, UpdateAdvertByIdMethod, errors.New("error with id advert"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	advertId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, UpdateAdvertByIdMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	data := models.AdvertUpdateData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, UpdateAdvertByIdMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}
	// data.Sanitize()

	data.ID = advertId

	err = h.uc.UpdateAdvertById(r.Context(), &data)
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, UpdateAdvertByIdMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, "advert successfully updated"); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, UpdateAdvertByIdMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, UpdateAdvertByIdMethod)
	}
}

// DeleteAdvertById handles the request for deleting advert by id
func (h *AdvertsClientHandler) DeleteAdvertById(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("csrftoken")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "csrf cookie not found")
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, DeleteAdvertByIdMethod, errors.New("error with id advert"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	advertId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, DeleteAdvertByIdMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	err = h.uc.DeleteAdvertById(r.Context(), advertId)
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, DeleteAdvertByIdMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, "Advert successfully deleted"); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, DeleteAdvertByIdMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, DeleteAdvertByIdMethod)
	}
}

// GetSquareAdvertsList handles the request for retrieving a square adverts.
func (h *AdvertsClientHandler) GetSquareAdvertsList(w http.ResponseWriter, r *http.Request) {
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
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetSquareAdvertsListMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	for _, adv := range adverts {
		adv.Sanitize()
	}

	if err = utils.WriteResponse(w, http.StatusOK, adverts); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetSquareAdvertsListMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetSquareAdvertsListMethod)
	}
}

// GetExistBuildingsByAddress handles the request for retrieving an existing buildings by address.
func (h *AdvertsClientHandler) GetExistBuildingByAddress(w http.ResponseWriter, r *http.Request) {
	data := models.AddressData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateHouseAdvertMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}

	building, err := h.uc.GetExistBuildingByAddress(r.Context(), &data)
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetExistBuildingsByAddressMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, building); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetExistBuildingsByAddressMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetExistBuildingsByAddressMethod)
	}
}

// GetRectangeAdvertsList handles the request for retrieving a rectangle adverts with search.
func (h *AdvertsClientHandler) GetRectangeAdvertsList(w http.ResponseWriter, r *http.Request) {
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
	adverts.Sanitize()

	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetRectangeAdvertsListMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, adverts); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetRectangeAdvertsListMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetRectangeAdvertsListMethod)
	}
}

func (h *AdvertsClientHandler) GetUserAdverts(w http.ResponseWriter, r *http.Request) {
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

	ID, ok := id.(int64)
	if !ok {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod, errors.New("error with id user"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}

	var userAdverts []*models.AdvertRectangleData
	if userAdverts, err = h.uc.GetRectangleAdvertsByUserId(r.Context(), page, size, ID); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "error getting user adverts")
		return
	}
	for _, adv := range userAdverts {
		adv.Sanitize()
	}

	if err := utils.WriteResponse(w, http.StatusOK, userAdverts); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod)
	}
}

func (h *AdvertsClientHandler) GetComplexAdverts(w http.ResponseWriter, r *http.Request) {
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
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetComplexAdvertsMethod, errors.New("error with id complex"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	complexId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetComplexAdvertsMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var complexAdverts []*models.AdvertRectangleData

	if complexAdverts, err = h.uc.GetRectangleAdvertsByComplexId(r.Context(), page, size, complexId); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetComplexAdvertsMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "error getting complex adverts")
		return
	}
	for _, adv := range complexAdverts {
		adv.Sanitize()
	}
	if err := utils.WriteResponse(w, http.StatusOK, complexAdverts); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetComplexAdvertsMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetComplexAdvertsMethod)
	}
}

func (h *AdvertsClientHandler) LikeAdvert(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("csrftoken")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "csrf cookie not found")
		return
	}

	id, ok := r.Context().Value(middleware.CookieName).(int64)
	if !ok {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateFlatAdvertMethod, errors.New("error with cookie"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}

	vars := mux.Vars(r)
	advertId := vars["id"]
	if advertId == "" {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetComplexAdvertsMethod, errors.New("error with id complex"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	advertIdInt, err := strconv.ParseInt(advertId, 10, 64)
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetComplexAdvertsMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	err = h.uc.LikeAdvert(r.Context(), advertIdInt, id)

	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), CreateFlatAdvertMethod, utils.DeliveryLayer, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusCreated, "success liked"); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), CreateFlatAdvertMethod, utils.DeliveryLayer, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateFlatAdvertMethod)
	}
}

func (h *AdvertsClientHandler) DislikeAdvert(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("csrftoken")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "csrf cookie not found")
		return
	}

	id, ok := r.Context().Value(middleware.CookieName).(int64)
	if !ok {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateFlatAdvertMethod, errors.New("error with cookie"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}

	vars := mux.Vars(r)
	advertId := vars["id"]
	if advertId == "" {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetComplexAdvertsMethod, errors.New("error with id complex"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	advertIdInt, err := strconv.ParseInt(advertId, 10, 64)
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetComplexAdvertsMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	err = h.uc.DislikeAdvert(r.Context(), advertIdInt, id)

	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), CreateFlatAdvertMethod, utils.DeliveryLayer, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusCreated, "success disliked"); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), CreateFlatAdvertMethod, utils.DeliveryLayer, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateFlatAdvertMethod)
	}
}

func (h *AdvertsClientHandler) GetLikedUserAdverts(w http.ResponseWriter, r *http.Request) {
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

	ID, ok := id.(int64)
	if !ok {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod, errors.New("error with id user"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}

	var userAdverts []*models.AdvertRectangleData
	if userAdverts, err = h.uc.GetRectangleAdvertsLikedByUserId(r.Context(), page, size, ID); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "error getting user adverts")
		return
	}
	for _, adv := range userAdverts {
		adv.Sanitize()
	}

	if err := utils.WriteResponse(w, http.StatusOK, userAdverts); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod)
	}
}
