package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts"
	genAdverts "2024_1_TeaStealers/internal/pkg/adverts/delivery/grpc/gen"
	genComplex "2024_1_TeaStealers/internal/pkg/complexes/delivery/grpc/gen"
	"2024_1_TeaStealers/internal/pkg/middleware"
	"2024_1_TeaStealers/internal/pkg/utils"
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

// AdvertsClientHandler handles HTTP requests for advert changes.
type AdvertsClientHandler struct {
	// uc represents the usecase interface for advert changes.
	client        genAdverts.AdvertsClient
	clientComplex genComplex.ComplexClient
	uc            adverts.AdvertUsecase
	logger        *zap.Logger
}

// NewAdvertsClientHandler creates a new instance of AdvertHandler.
func NewAdvertsClientHandler(grpcConn *grpc.ClientConn, grpcConn2 *grpc.ClientConn, uc adverts.AdvertUsecase, logger *zap.Logger) *AdvertsClientHandler {
	return &AdvertsClientHandler{client: genAdverts.NewAdvertsClient(grpcConn), clientComplex: genComplex.NewComplexClient(grpcConn2), uc: uc, logger: logger}
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
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateFlatAdvertMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}

	var mater genAdverts.MaterialBuilding
	switch data.Material {
	case models.MaterialBrick:
		mater = genAdverts.MaterialBuilding_MATERIAL_BRICK
	case models.MaterialMonolithic:
		mater = genAdverts.MaterialBuilding_MATERIAL_MONOLITHIC
	case models.MaterialWood:
		mater = genAdverts.MaterialBuilding_MATERIAL_WOOD
	case models.MaterialPanel:
		mater = genAdverts.MaterialBuilding_MATERIAL_PANEL
	case models.MaterialStalinsky:
		mater = genAdverts.MaterialBuilding_MATERIAL_STALINSKY
	case models.MaterialBlock:
		mater = genAdverts.MaterialBuilding_MATERIAL_BLOCK
	case models.MaterialMonolithicBlock:
		mater = genAdverts.MaterialBuilding_MATERIAL_MONOLITHIC_BLOCK
	case models.MaterialFrame:
		mater = genAdverts.MaterialBuilding_MATERIAL_FRAME
	case models.MaterialAeratedConcreteBlock:
		mater = genAdverts.MaterialBuilding_MATERIAL_AERATED_CONCRETE_BLOCK
	case models.MaterialGasSilicateBlock:
		mater = genAdverts.MaterialBuilding_MATERIAL_GAS_SILICATE_BLOCK
	case models.MaterialFoamConcreteBlock:
		mater = genAdverts.MaterialBuilding_MATERIAL_FOAM_CONCRETE_BLOCK
	}

	req := &genAdverts.CreateFlatAdvertRequest{
		UserId:      data.UserID,
		TypeSale:    string(data.AdvertTypeSale),
		Title:       data.Title,
		Description: data.Description,
		Phone:       data.Phone,
		IsAgent:     data.IsAgent,
		CreateFlatProp: &genAdverts.FlatProperties{
			Floor:             int32(data.Floor),
			CeilingHeight:     data.CeilingHeight,
			SquareGeneral:     data.SquareGeneral,
			RoomCount:         int32(data.RoomCount),
			SquareResidential: data.SquareResidential,
			Apartment:         data.Apartment,
			FloorGeneral:      int32(data.FloorGeneral),
		},
		Price:    data.Price,
		Material: mater,
		Address: &genAdverts.AddressData{Province: data.Address.Province,
			Town: data.Address.Town, Street: data.Address.Street, House: data.Address.House,
			Metro: data.Address.Metro, AddressPoint: data.Address.AddressPoint},
		YearCreation: int32(data.YearCreation),
	}

	newAdvertResp, err := h.client.CreateFlatAdvert(r.Context(), req)

	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateFlatAdvertMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	date := newAdvertResp.DateCreation[:19]
	dCr, err := utils.StringToTime("2006-01-02 15:04:05", date)

	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateFlatAdvertMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	newAdvert := models.Advert{
		ID:             newAdvertResp.Id,
		UserID:         newAdvertResp.UserId,
		AdvertTypeSale: models.TypePlacementAdvert(newAdvertResp.TypeSale),
		Title:          newAdvertResp.Title,
		Description:    newAdvertResp.Description,
		Phone:          newAdvertResp.Phone,
		IsAgent:        newAdvertResp.IsAgent,
		Priority:       int(newAdvertResp.Priority),
		DateCreation:   dCr,
		IsDeleted:      false,
	}
	newAdvert.Sanitize()

	if err = utils.WriteResponse(w, http.StatusCreated, newAdvert); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateFlatAdvertMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateFlatAdvertMethod)
		return
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

	var statusArea genAdverts.StatusAreaHouse
	switch data.StatusArea {
	case "IHC":
		statusArea = genAdverts.StatusAreaHouse_STATUS_AREA_IHC
	case "DNP":
		statusArea = genAdverts.StatusAreaHouse_STATUS_AREA_DNP
	case "G":
		statusArea = genAdverts.StatusAreaHouse_STATUS_AREA_G
	case "F":
		statusArea = genAdverts.StatusAreaHouse_STATUS_AREA_F
	case "PSP":
		statusArea = genAdverts.StatusAreaHouse_STATUS_AREA_PSP
	}

	var statusHome genAdverts.StatusHomeHouse
	switch data.StatusHome {
	case "Live":
		statusHome = genAdverts.StatusHomeHouse_STATUS_HOME_LIVE
	case "RepairNeed":
		statusHome = genAdverts.StatusHomeHouse_STATUS_HOME_REPAIR_NEED
	case "CompleteNeed":
		statusHome = genAdverts.StatusHomeHouse_STATUS_HOME_COMPLETE_NEED
	case "Renovation":
		statusHome = genAdverts.StatusHomeHouse_STATUS_HOME_RENOVATION
	}

	var mater genAdverts.MaterialBuilding
	switch data.Material {
	case models.MaterialBrick:
		mater = genAdverts.MaterialBuilding_MATERIAL_BRICK
	case models.MaterialMonolithic:
		mater = genAdverts.MaterialBuilding_MATERIAL_MONOLITHIC
	case models.MaterialWood:
		mater = genAdverts.MaterialBuilding_MATERIAL_WOOD
	case models.MaterialPanel:
		mater = genAdverts.MaterialBuilding_MATERIAL_PANEL
	case models.MaterialStalinsky:
		mater = genAdverts.MaterialBuilding_MATERIAL_STALINSKY
	case models.MaterialBlock:
		mater = genAdverts.MaterialBuilding_MATERIAL_BLOCK
	case models.MaterialMonolithicBlock:
		mater = genAdverts.MaterialBuilding_MATERIAL_MONOLITHIC_BLOCK
	case models.MaterialFrame:
		mater = genAdverts.MaterialBuilding_MATERIAL_FRAME
	case models.MaterialAeratedConcreteBlock:
		mater = genAdverts.MaterialBuilding_MATERIAL_AERATED_CONCRETE_BLOCK
	case models.MaterialGasSilicateBlock:
		mater = genAdverts.MaterialBuilding_MATERIAL_GAS_SILICATE_BLOCK
	case models.MaterialFoamConcreteBlock:
		mater = genAdverts.MaterialBuilding_MATERIAL_FOAM_CONCRETE_BLOCK
	}

	resp, err := h.client.CreateHouseAdvert(r.Context(), &genAdverts.CreateHouseAdvertRequest{
		UserId:      id,
		TypeSale:    string(data.AdvertTypeSale),
		Title:       data.Title,
		Description: data.Description,
		Phone:       data.Phone,
		IsAgent:     data.IsAgent,
		CreateHouseProp: &genAdverts.HouseProperties{
			CeilingHeight: data.CeilingHeight,
			SquareArea:    data.SquareArea,
			SquareHouse:   data.SquareHouse,
			BedroomCount:  int32(data.BedroomCount),
			StatusArea:    statusArea,
			Cottage:       data.Cottage,
			StatusHome:    statusHome,
			Floor:         int32(data.FloorGeneral),
		},
		Address: &genAdverts.AddressData{Province: data.Address.Province,
			Town: data.Address.Town, Street: data.Address.Street, House: data.Address.House,
			Metro: data.Address.Metro, AddressPoint: data.Address.AddressPoint},

		Price:        data.Price,
		YearCreation: int32(data.YearCreation),
		Material:     mater})

	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateHouseAdvertMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	date := resp.DateCreation[:19]
	tCreate, err := utils.StringToTime("2006-01-02 15:04:05", date)

	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateHouseAdvertMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	newAdvert := models.Advert{
		ID:             resp.Id,
		UserID:         resp.UserId,
		AdvertTypeSale: models.TypePlacementAdvert(resp.TypeSale),
		Title:          resp.Title,
		Description:    resp.Description,
		Phone:          resp.Phone,
		IsAgent:        resp.IsAgent,
		Priority:       int(resp.Priority),
		DateCreation:   tCreate,
	}

	newAdvert.Sanitize()

	if err = utils.WriteResponse(w, http.StatusCreated, newAdvert); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateHouseAdvertMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateHouseAdvertMethod)
		return
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

	md := metadata.New(map[string]string{"userid": strconv.FormatInt(r.Context().Value(middleware.CookieName).(int64), 10)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	advertDataResponse, err := h.client.GetAdvertById(ctx, &genAdverts.GetAdvertByIdRequest{Id: advertId})
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetAdvertByIdMethod, err, int(advertDataResponse.RespCode))
		utils.WriteError(w, int(advertDataResponse.RespCode), err.Error())
		return
	}

	var priceHistory []*models.PriceChangeData
	for _, pcd := range advertDataResponse.PriceHistory {
		date := pcd.DateCreation[:19]
		dateTime, err := utils.StringToTime("2006-01-02 15:04:05", date)
		if err != nil {
			utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetAdvertByIdMethod, err, http.StatusInternalServerError)
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
	case genAdverts.MaterialBuilding_MATERIAL_BRICK:
		material = models.MaterialBrick
	case genAdverts.MaterialBuilding_MATERIAL_MONOLITHIC:
		material = models.MaterialMonolithic
	case genAdverts.MaterialBuilding_MATERIAL_WOOD:
		material = models.MaterialWood
	case genAdverts.MaterialBuilding_MATERIAL_PANEL:
		material = models.MaterialPanel
	case genAdverts.MaterialBuilding_MATERIAL_STALINSKY:
		material = models.MaterialStalinsky
	case genAdverts.MaterialBuilding_MATERIAL_BLOCK:
		material = models.MaterialBlock
	case genAdverts.MaterialBuilding_MATERIAL_MONOLITHIC_BLOCK:
		material = models.MaterialMonolithicBlock
	case genAdverts.MaterialBuilding_MATERIAL_FRAME:
		material = models.MaterialFrame
	case genAdverts.MaterialBuilding_MATERIAL_AERATED_CONCRETE_BLOCK:
		material = models.MaterialAeratedConcreteBlock
	case genAdverts.MaterialBuilding_MATERIAL_GAS_SILICATE_BLOCK:
		material = models.MaterialGasSilicateBlock
	case genAdverts.MaterialBuilding_MATERIAL_FOAM_CONCRETE_BLOCK:
		material = models.MaterialFoamConcreteBlock
	}

	var statusArea models.StatusAreaHouse
	var statusHome models.StatusHomeHouse
	var houseProperties *models.HouseProperties = nil
	var flatProperties *models.FlatProperties = nil
	var complexProperties *models.ComplexAdvertProperties = nil

	if advertDataResponse.HouseProp != nil {
		switch advertDataResponse.HouseProp.StatusArea {
		case genAdverts.StatusAreaHouse_STATUS_AREA_IHC:
			statusArea = "IHC"
		case genAdverts.StatusAreaHouse_STATUS_AREA_DNP:
			statusArea = "DNP"
		case genAdverts.StatusAreaHouse_STATUS_AREA_G:
			statusArea = "G"
		case genAdverts.StatusAreaHouse_STATUS_AREA_F:
			statusArea = "F"
		case genAdverts.StatusAreaHouse_STATUS_AREA_PSP:
			statusArea = "PSP"
		}

		switch advertDataResponse.HouseProp.StatusHome {
		case genAdverts.StatusHomeHouse_STATUS_HOME_LIVE:
			statusHome = "Live"
		case genAdverts.StatusHomeHouse_STATUS_HOME_REPAIR_NEED:
			statusHome = "RepairNeed"
		case genAdverts.StatusHomeHouse_STATUS_HOME_COMPLETE_NEED:
			statusHome = "CompleteNeed"
		case genAdverts.StatusHomeHouse_STATUS_HOME_RENOVATION:
			statusHome = "Renovation"
		}

		houseProperties = &models.HouseProperties{CeilingHeight: advertDataResponse.HouseProp.CeilingHeight,
			SquareArea: advertDataResponse.HouseProp.SquareArea, SquareHouse: advertDataResponse.HouseProp.SquareHouse,
			BedroomCount: int(advertDataResponse.HouseProp.BedroomCount), StatusArea: statusArea,
			Cottage: advertDataResponse.HouseProp.Cottage, StatusHome: statusHome, Floor: int(advertDataResponse.HouseProp.Floor)}
	}

	if advertDataResponse.FlatProperties != nil {
		log.Println(advertDataResponse.FlatProperties)
		flatProperties = &models.FlatProperties{CeilingHeight: advertDataResponse.FlatProperties.CeilingHeight,
			RoomCount: int(advertDataResponse.FlatProperties.RoomCount), FloorGeneral: int(advertDataResponse.FlatProperties.FloorGeneral),
			Apartment: advertDataResponse.FlatProperties.Apartment, SquareGeneral: advertDataResponse.FlatProperties.SquareGeneral,
			Floor: int(advertDataResponse.FlatProperties.Floor), SquareResidential: advertDataResponse.FlatProperties.SquareResidential}
	}

	if advertDataResponse.ComplexProperties != nil {
		complexProperties = &models.ComplexAdvertProperties{ComplexId: advertDataResponse.ComplexProperties.ComplexId,
			NameComplex: advertDataResponse.ComplexProperties.NameComplex, PhotoCompany: advertDataResponse.ComplexProperties.PhotoCompany,
			NameCompany: advertDataResponse.ComplexProperties.NameCompany}
	}

	date := advertDataResponse.DateCreation[:19]
	dateTime, err := utils.StringToTime("2006-01-02 15:04:05", date)
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetAdvertByIdMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	advert := &models.AdvertData{ID: advertDataResponse.Id, AdvertType: advertDataResponse.AdvertType, TypeSale: advertDataResponse.TypeSale,
		Title: advertDataResponse.Title, Description: advertDataResponse.Description, CountViews: advertDataResponse.CountViews,
		CountLikes: advertDataResponse.CountLikes, Price: advertDataResponse.Price, Phone: advertDataResponse.Phone, IsLiked: advertDataResponse.IsLiked,
		IsAgent: advertDataResponse.IsAgent, Metro: advertDataResponse.Metro, Address: advertDataResponse.Address, AddressPoint: advertDataResponse.AddressPoint,
		YearCreation: int(advertDataResponse.YearCreation), PriceChange: priceHistory, Images: images, FlatProperties: flatProperties,
		HouseProperties: houseProperties, Material: material, ComplexProperties: complexProperties, DateCreation: dateTime}

	if err = utils.WriteResponse(w, http.StatusOK, advert); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetAdvertByIdMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetAdvertByIdMethod)
		return
	}
}

// UpdatePriority handles the request for updatin advert priority
func (h *AdvertsClientHandler) UpdatePriority(w http.ResponseWriter, r *http.Request) {
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

	data := models.UpdatePriorityRequest{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateHouseAdvertMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}

	amount, err := strconv.Atoi(data.DonationAmount)
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetAdvertByIdMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	priorityResponse, err := h.client.IncreasePriority(r.Context(), &genAdverts.IncreasePriorityRequest{AdvertId: advertId, Amount: int64(amount)})
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetAdvertByIdMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, &models.Rating{Rating: priorityResponse.Amount}); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetAdvertByIdMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetAdvertByIdMethod)
	}
}

// GetPriority handles the request for priority advert priority
func (h *AdvertsClientHandler) GetPriority(w http.ResponseWriter, r *http.Request) {
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

	priorityResponse, err := h.client.GetPriority(r.Context(), &genAdverts.GetPriorityRequest{AdvertId: advertId})
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetAdvertByIdMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, &models.Rating{Rating: priorityResponse.Amount}); err != nil {
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
	data.Sanitize()

	data.ID = advertId

	var statusArea genAdverts.StatusAreaHouse
	switch data.HouseProperties.StatusArea {
	case "IHC":
		statusArea = genAdverts.StatusAreaHouse_STATUS_AREA_IHC
	case "DNP":
		statusArea = genAdverts.StatusAreaHouse_STATUS_AREA_DNP
	case "G":
		statusArea = genAdverts.StatusAreaHouse_STATUS_AREA_G
	case "F":
		statusArea = genAdverts.StatusAreaHouse_STATUS_AREA_F
	case "PSP":
		statusArea = genAdverts.StatusAreaHouse_STATUS_AREA_PSP
	}

	var statusHome genAdverts.StatusHomeHouse
	switch data.HouseProperties.StatusHome {
	case "Live":
		statusHome = genAdverts.StatusHomeHouse_STATUS_HOME_LIVE
	case "RepairNeed":
		statusHome = genAdverts.StatusHomeHouse_STATUS_HOME_REPAIR_NEED
	case "CompleteNeed":
		statusHome = genAdverts.StatusHomeHouse_STATUS_HOME_COMPLETE_NEED
	case "Renovation":
		statusHome = genAdverts.StatusHomeHouse_STATUS_HOME_RENOVATION
	}

	var mater genAdverts.MaterialBuilding
	switch data.Material {
	case models.MaterialBrick:
		mater = genAdverts.MaterialBuilding_MATERIAL_BRICK
	case models.MaterialMonolithic:
		mater = genAdverts.MaterialBuilding_MATERIAL_MONOLITHIC
	case models.MaterialWood:
		mater = genAdverts.MaterialBuilding_MATERIAL_WOOD
	case models.MaterialPanel:
		mater = genAdverts.MaterialBuilding_MATERIAL_PANEL
	case models.MaterialStalinsky:
		mater = genAdverts.MaterialBuilding_MATERIAL_STALINSKY
	case models.MaterialBlock:
		mater = genAdverts.MaterialBuilding_MATERIAL_BLOCK
	case models.MaterialMonolithicBlock:
		mater = genAdverts.MaterialBuilding_MATERIAL_MONOLITHIC_BLOCK
	case models.MaterialFrame:
		mater = genAdverts.MaterialBuilding_MATERIAL_FRAME
	case models.MaterialAeratedConcreteBlock:
		mater = genAdverts.MaterialBuilding_MATERIAL_AERATED_CONCRETE_BLOCK
	case models.MaterialGasSilicateBlock:
		mater = genAdverts.MaterialBuilding_MATERIAL_GAS_SILICATE_BLOCK
	case models.MaterialFoamConcreteBlock:
		mater = genAdverts.MaterialBuilding_MATERIAL_FOAM_CONCRETE_BLOCK
	}

	resp, err := h.client.UpdateAdvertById(r.Context(), &genAdverts.UpdateAdvertByIdRequest{
		Id:          advertId,
		AdvertType:  data.TypeAdvert,
		TypeSale:    data.TypeSale,
		Title:       data.Title,
		Description: data.Description,
		Price:       data.Price,
		IsAgent:     data.IsAgent,
		Address: &genAdverts.AddressData{Province: data.Address.Province,
			Town: data.Address.Town, Street: data.Address.Street, House: data.Address.House,
			Metro: data.Address.Metro, AddressPoint: data.Address.AddressPoint},
		HouseProp: &genAdverts.HouseProperties{
			CeilingHeight: data.HouseProperties.CeilingHeight,
			SquareArea:    data.HouseProperties.SquareArea,
			SquareHouse:   data.HouseProperties.SquareHouse,
			BedroomCount:  int32(data.HouseProperties.BedroomCount),
			StatusArea:    statusArea,
			Cottage:       data.HouseProperties.Cottage,
			StatusHome:    statusHome,
			Floor:         int32(data.HouseProperties.Floor),
		},
		FlatProperties: &genAdverts.FlatProperties{
			Floor:             int32(data.FlatProperties.Floor),
			CeilingHeight:     data.FlatProperties.CeilingHeight,
			SquareGeneral:     data.FlatProperties.SquareGeneral,
			RoomCount:         int32(data.FlatProperties.RoomCount),
			SquareResidential: data.FlatProperties.SquareResidential,
			Apartment:         data.FlatProperties.Apartment,
			FloorGeneral:      int32(data.FlatProperties.FloorGeneral),
		},
		YearCreation: int32(data.YearCreation),
		Material:     mater,
	})

	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, UpdateAdvertByIdMethod, err, int(resp.RespCode))
		utils.WriteError(w, int(resp.RespCode), err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, "advert successfully updated"); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, UpdateAdvertByIdMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, UpdateAdvertByIdMethod)
		return
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

	_, err = h.client.DeleteAdvertById(r.Context(), &genAdverts.DeleteAdvertByIdRequest{Id: advertId})
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, DeleteAdvertByIdMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, "Advert successfully deleted"); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, DeleteAdvertByIdMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, DeleteAdvertByIdMethod)
		return
	}
}

// GetSquareAdvertsList handles the request for retrieving a square adverts.
func (h *AdvertsClientHandler) GetSquareAdvertsList(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("pageNum")
	size := r.URL.Query().Get("pageSize")

	pageNum, err := strconv.Atoi(page)
	if err != nil {
		pageNum = 1
	}
	pageSize, err := strconv.Atoi(size)
	if err != nil {
		pageSize = 10000000
	}
	utils.LogSuccesResponse(h.logger, "HERE ONE", utils.DeliveryLayer, GetSquareAdvertsListMethod)

	offset := (pageNum - 1) * pageSize
	utils.LogSuccesResponse(h.logger, "HERE TWO", utils.DeliveryLayer, GetSquareAdvertsListMethod)

	advResp, err := h.client.GetSquareAdvertsList(r.Context(), &genAdverts.GetSquareAdvertsListRequest{Offset: int64(offset), PageSize: int64(pageSize)}) // todo здесь ломается???
	if err != nil {
		utils.LogSuccesResponse(h.logger, "HERE THREE ERROR", utils.DeliveryLayer, GetSquareAdvertsListMethod)
		// utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetSquareAdvertsListMethod, err, int(advResp.RespCode))
		// utils.WriteError(w, int(advResp.RespCode), err.Error())
		return
	}
	utils.LogSuccesResponse(h.logger, "HERE THREE", utils.DeliveryLayer, GetSquareAdvertsListMethod)

	foundAdverts := make([]*models.AdvertSquareData, 0) // todo

	for _, adv := range advResp.SquareData {
		date := adv.DateCreation[:19]
		var dateTime time.Time
		if date != "" {
			dateTime, err = utils.StringToTime("2006-01-02 15:04:05", date)
			if err != nil {
				utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetSquareAdvertsListMethod, err, http.StatusInternalServerError)
				utils.WriteError(w, int(advResp.RespCode), err.Error())
				return
			}
		}

		newadv := &models.AdvertSquareData{
			ID:           adv.Id,
			TypeAdvert:   adv.TypeAdvert,
			Photo:        adv.Photo,
			TypeSale:     adv.TypeSale,
			Address:      adv.Address,
			Metro:        adv.Metro,
			Price:        int(adv.Price),
			DateCreation: dateTime,
		}

		if adv.HouseProp != nil {
			newadv.HouseProperties = &models.HouseSquareProperties{Cottage: adv.HouseProp.Cottage,
				SquareArea: adv.HouseProp.SquareArea, SquareHouse: adv.HouseProp.SquareHouse, BedroomCount: int(adv.HouseProp.BedroomCount),
				Floor: int(adv.HouseProp.Floor)}
		}

		if adv.FlatProp != nil {
			newadv.FlatProperties = &models.FlatSquareProperties{
				Floor:         int(adv.FlatProp.Floor),
				FloorGeneral:  int(adv.FlatProp.FloorGeneral),
				RoomCount:     int(adv.FlatProp.RoomCount),
				SquareGeneral: adv.FlatProp.SquareGeneral,
			}
		}

		foundAdverts = append(foundAdverts, newadv)
	}

	if err = utils.WriteResponse(w, http.StatusOK, foundAdverts); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetSquareAdvertsListMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetSquareAdvertsListMethod)
		return
	}
}

// GetExistBuildingByAddress handles the request for retrieving an existing buildings by address.
func (h *AdvertsClientHandler) GetExistBuildingByAddress(w http.ResponseWriter, r *http.Request) {
	data := models.AddressData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateHouseAdvertMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}

	respBuilding, err := h.client.GetExistBuildingByAddress(r.Context(), &genAdverts.GetExistBuildingByAddressRequest{
		AdrData: &genAdverts.AddressData{Province: data.Province,
			Town: data.Town, Street: data.Street, House: data.House,
			Metro: data.Metro, AddressPoint: data.AddressPoint},
	})

	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetExistBuildingsByAddressMethod, err, int(respBuilding.RespCode))
		utils.WriteError(w, int(respBuilding.RespCode), err.Error())
		return
	}

	building := models.BuildingData{
		ComplexName:  respBuilding.ComplexName,
		Floor:        int(respBuilding.Floor),
		Material:     models.MaterialBuilding(respBuilding.Material),
		YearCreation: int(respBuilding.YearCreation),
	}

	building.Sanitize()

	if err = utils.WriteResponse(w, http.StatusOK, building); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetExistBuildingsByAddressMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetExistBuildingsByAddressMethod)
		return
	}
}

// GetRectangleAdvertsList handles the request for retrieving a rectangle adverts with search.
func (h *AdvertsClientHandler) GetRectangleAdvertsList(w http.ResponseWriter, r *http.Request) {
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
		page = 0
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 10000000000
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

	// offset := (page - 1) * size

	advList, err := h.client.GetRectangleAdvertsList(r.Context(), &genAdverts.GetRectangleAdvertsListRequest{
		MinPrice:   minPrice,
		MaxPrice:   maxPrice,
		Page:       int64(page),
		Size:       int64(size),
		RoomCount:  int32(roomCount),
		Address:    adress,
		DealType:   dealType,
		AdvertType: advertType,
	})

	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetRectangeAdvertsListMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	foundAdverts := make([]*models.AdvertRectangleData, 0)

	for _, adv := range advList.Adverts {
		date := adv.DateCreation[:19]
		tCr, err := utils.StringToTime("2006-01-02 15:04:05", date)
		if err != nil {
			utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod, err, http.StatusInternalServerError)
			utils.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}

		newadv := &models.AdvertRectangleData{
			ID:           adv.Id,
			Title:        adv.Title,
			Description:  adv.Description,
			TypeAdvert:   adv.AdvertType,
			Photo:        adv.Photo,
			Phone:        adv.Phone,
			TypeSale:     adv.TypeSale,
			Address:      adv.Address,
			IsLiked:      adv.IsLiked,
			Price:        int(adv.Price),
			DateCreation: tCr,
			Rating:       adv.Rating,
		}

		if adv.FlatProperties != nil {
			newadv.FlatProperties = &models.FlatRectangleProperties{
				Floor:         int(adv.FlatProperties.Floor),
				FloorGeneral:  int(adv.FlatProperties.FloorGeneral),
				RoomCount:     int(adv.FlatProperties.RoomCount),
				SquareGeneral: adv.FlatProperties.SquareGeneral,
			}
		}

		if adv.HouseProp != nil {
			newadv.HouseProperties = &models.HouseRectangleProperties{Cottage: adv.HouseProp.Cottage,
				SquareArea: adv.HouseProp.SquareArea, SquareHouse: adv.HouseProp.SquareHouse, BedroomCount: int(adv.HouseProp.BedroomCount),
				Floor: int(adv.HouseProp.Floor)}
		}

		foundAdverts = append(foundAdverts, newadv)
	}

	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetRectangeAdvertsListMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, foundAdverts); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetRectangeAdvertsListMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetRectangeAdvertsListMethod)
		return
	}
}

func (h *AdvertsClientHandler) GetUserAdverts(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middleware.CookieName)
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 0
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 100000000
	}

	UId, ok := userId.(int64)
	if !ok {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod, errors.New("error with id user"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}

	userAdverts, err := h.client.GetRectangleAdvertsByUser(r.Context(), &genAdverts.GetUserAdvertsRequest{
		Page: int64(page), Size: int64(size), UserId: UId,
	})
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod, err, int(userAdverts.RespCode))
		utils.WriteError(w, http.StatusBadRequest, "error getting user adverts")
		return
	}

	foundAdverts := make([]*models.AdvertRectangleData, 0)

	for _, adv := range userAdverts.RectDataSlice {
		date := adv.DateCreation[:19]
		tCr, err := utils.StringToTime("2006-01-02 15:04:05", date)
		if err != nil {
			utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod, err, http.StatusInternalServerError)
			utils.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}
		newadv := &models.AdvertRectangleData{
			ID:           adv.Id,
			Title:        adv.Title,
			Description:  adv.Description,
			TypeAdvert:   adv.AdvertType,
			Photo:        adv.Photo,
			Phone:        adv.Phone,
			TypeSale:     adv.TypeSale,
			Address:      adv.Address,
			IsLiked:      adv.IsLiked,
			Price:        int(adv.Price),
			Rating:       adv.Rating,
			DateCreation: tCr,
		}

		if adv.FlatProperties != nil {
			newadv.FlatProperties = &models.FlatRectangleProperties{
				Floor:         int(adv.FlatProperties.Floor),
				FloorGeneral:  int(adv.FlatProperties.FloorGeneral),
				RoomCount:     int(adv.FlatProperties.RoomCount),
				SquareGeneral: adv.FlatProperties.SquareGeneral,
			}
		}

		if adv.HouseProp != nil {
			newadv.HouseProperties = &models.HouseRectangleProperties{Cottage: adv.HouseProp.Cottage,
				SquareArea: adv.HouseProp.SquareArea, SquareHouse: adv.HouseProp.SquareHouse, BedroomCount: int(adv.HouseProp.BedroomCount),
				Floor: int(adv.HouseProp.Floor)}
		}

		foundAdverts = append(foundAdverts, newadv)
	}

	if err := utils.WriteResponse(w, http.StatusOK, foundAdverts); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod)
		return
	}
}

func (h *AdvertsClientHandler) GetComplexAdverts(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 0
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 1000000
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

	var foundAdverts []*models.AdvertRectangleData
	complexAdverts, err := h.client.GetRectangleAdvertsByComplex(r.Context(), &genAdverts.GetComplexAdvertsRequest{ComplexId: complexId,
		PageSize: int64(size), Offset: int64(page)})
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetComplexAdvertsMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "error getting complex adverts")
		return
	}
	for _, adv := range complexAdverts.RectDataSlice {
		date := adv.DateCreation[:19]
		tCr, err := utils.StringToTime("2006-01-02 15:04:05", date)
		if err != nil {
			utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod, err, http.StatusInternalServerError)
			utils.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}
		newadv := &models.AdvertRectangleData{
			ID:          adv.Id,
			Title:       adv.Title,
			Description: adv.Description,
			TypeAdvert:  adv.AdvertType,
			Photo:       adv.Photo,
			Phone:       adv.Phone,
			TypeSale:    adv.TypeSale,
			Address:     adv.Address,
			IsLiked:     adv.IsLiked,
			FlatProperties: &models.FlatRectangleProperties{
				Floor:         int(adv.FlatProperties.Floor),
				FloorGeneral:  int(adv.FlatProperties.FloorGeneral),
				RoomCount:     int(adv.FlatProperties.RoomCount),
				SquareGeneral: adv.FlatProperties.SquareGeneral,
			},
			HouseProperties: &models.HouseRectangleProperties{Cottage: adv.HouseProp.Cottage,
				SquareArea: adv.HouseProp.SquareArea, SquareHouse: adv.HouseProp.SquareHouse, BedroomCount: int(adv.HouseProp.BedroomCount),
				Floor: int(adv.HouseProp.Floor)},
			Price:        int(adv.Price),
			DateCreation: tCr,
		}

		foundAdverts = append(foundAdverts, newadv)
	}

	if err := utils.WriteResponse(w, http.StatusOK, foundAdverts); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetComplexAdvertsMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetComplexAdvertsMethod)
		return
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

	_, err = h.client.LikeAdvert(r.Context(), &genAdverts.LikeAdvertRequest{AdvId: advertIdInt, UserId: id})

	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateFlatAdvertMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusCreated, "success liked"); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateFlatAdvertMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateFlatAdvertMethod)
		return
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

	_, err = h.client.DislikeAdvert(r.Context(), &genAdverts.DislikeAdvertRequest{AdvId: advertIdInt, UserId: id})

	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateFlatAdvertMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusCreated, "success disliked"); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateFlatAdvertMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CreateFlatAdvertMethod)
		return
	}
}

func (h *AdvertsClientHandler) GetLikedUserAdverts(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.CookieName)
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 0
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 1000000
	}

	userID, ok := id.(int64)
	if !ok {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod, errors.New("error with id user"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}

	var foundAdverts []*models.AdvertRectangleData
	userAdvertsResp, err := h.client.GetLikedUserAdverts(r.Context(), &genAdverts.GetLikedUserAdvertsRequest{UserId: userID, PageSize: int64(size), Offset: int64(page)})
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod, err, int(userAdvertsResp.RespCode))
		utils.WriteError(w, http.StatusBadRequest, "error getting user adverts")
		return
	}

	for _, adv := range userAdvertsResp.RectDataSlice {
		date := adv.DateCreation[:19]
		tCr, err := utils.StringToTime("2006-01-02 15:04:05", date)
		if err != nil {
			utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod, err, http.StatusInternalServerError)
			utils.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}
		newadv := &models.AdvertRectangleData{
			ID:           adv.Id,
			Title:        adv.Title,
			Description:  adv.Description,
			TypeAdvert:   adv.AdvertType,
			Photo:        adv.Photo,
			Phone:        adv.Phone,
			TypeSale:     adv.TypeSale,
			Address:      adv.Address,
			IsLiked:      adv.IsLiked,
			Rating:       adv.Rating,
			Price:        int(adv.Price),
			DateCreation: tCr,
		}

		if adv.FlatProperties != nil {
			newadv.FlatProperties = &models.FlatRectangleProperties{
				Floor:         int(adv.FlatProperties.Floor),
				FloorGeneral:  int(adv.FlatProperties.FloorGeneral),
				RoomCount:     int(adv.FlatProperties.RoomCount),
				SquareGeneral: adv.FlatProperties.SquareGeneral,
			}
		}

		if adv.HouseProp != nil {
			newadv.HouseProperties = &models.HouseRectangleProperties{Cottage: adv.HouseProp.Cottage,
				SquareArea: adv.HouseProp.SquareArea, SquareHouse: adv.HouseProp.SquareHouse, BedroomCount: int(adv.HouseProp.BedroomCount),
				Floor: int(adv.HouseProp.Floor)}
		}

		foundAdverts = append(foundAdverts, newadv)
	}

	if err := utils.WriteResponse(w, http.StatusOK, foundAdverts); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, GetUserAdvertsMethod)
		return
	}
}
