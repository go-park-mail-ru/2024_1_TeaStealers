package grpc

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts"
	"2024_1_TeaStealers/internal/pkg/adverts/delivery/grpc/gen"
	genAdverts "2024_1_TeaStealers/internal/pkg/adverts/delivery/grpc/gen"
	"2024_1_TeaStealers/internal/pkg/utils"
	"context"
	"log"

	"github.com/satori/uuid"
	"go.uber.org/zap"
)

const StatusBadRequest = 400
const StatusOk = 200

type AdvertsServerHandler struct {
	genAdverts.AdvertsServer
	// uc represents the usecase interface for authentication.
	uc     adverts.AdvertUsecase
	logger *zap.Logger
}

// NewServerAdvertsHandler creates a new instance of AdvertsServerHandler.
func NewServerAdvertsHandler(uc adverts.AdvertUsecase, logger *zap.Logger) *AdvertsServerHandler {
	return &AdvertsServerHandler{uc: uc, logger: logger}
}

func (h *AdvertsServerHandler) GetAdvertById(ctx context.Context, reqAdv *genAdverts.GetAdvertByIdRequest) (*genAdverts.GetAdvertByIdResponse, error) {
	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())

	advert, err := h.uc.GetAdvertById(ctx, reqAdv.Id)

	if err != nil {

		// h.logger.Error(ctx.Value("requestId").(string) + " " + err.Error())
		// utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, SignUpMethod, err, http.StatusBadRequest)
		return &genAdverts.GetAdvertByIdResponse{RespCode: StatusBadRequest}, err
	}

	var priceHistory []*gen.PriceChangeData
	for _, pcd := range advert.PriceChange {
		priceHistory = append(priceHistory, &gen.PriceChangeData{Price: pcd.Price, DateCreation: pcd.DateCreation.String()})
	}

	var images []*gen.ImageResp
	for _, img := range advert.Images {
		images = append(images, &gen.ImageResp{Id: img.ID, Photo: img.Photo, Priority: int32(img.Priority)})
	}

	var material gen.MaterialBuilding

	switch advert.Material {
	case models.MaterialBrick:
		material = gen.MaterialBuilding_MATERIAL_BRICK
	case models.MaterialMonolithic:
		material = gen.MaterialBuilding_MATERIAL_MONOLITHIC
	case models.MaterialWood:
		material = gen.MaterialBuilding_MATERIAL_WOOD
	case models.MaterialPanel:
		material = gen.MaterialBuilding_MATERIAL_PANEL
	case models.MaterialStalinsky:
		material = gen.MaterialBuilding_MATERIAL_STALINSKY
	case models.MaterialBlock:
		material = gen.MaterialBuilding_MATERIAL_BLOCK
	case models.MaterialMonolithicBlock:
		material = gen.MaterialBuilding_MATERIAL_MONOLITHIC_BLOCK
	case models.MaterialFrame:
		material = gen.MaterialBuilding_MATERIAL_FRAME
	case models.MaterialAeratedConcreteBlock:
		material = gen.MaterialBuilding_MATERIAL_AERATED_CONCRETE_BLOCK
	case models.MaterialGasSilicateBlock:
		material = gen.MaterialBuilding_MATERIAL_GAS_SILICATE_BLOCK
	case models.MaterialFoamConcreteBlock:
		material = gen.MaterialBuilding_MATERIAL_FOAM_CONCRETE_BLOCK

	}

	var statusArea gen.StatusAreaHouse
	var statusHome gen.StatusHomeHouse
	var houseProperties *gen.HouseProperties = nil
	var flatProperties *gen.FlatProperties = nil
	var complexProperties *gen.ComplexAdvertProperties = nil
	if advert.HouseProperties != nil {
		switch advert.HouseProperties.StatusArea {
		case "IHC":
			statusArea = gen.StatusAreaHouse_STATUS_AREA_IHC
		case "DNP":
			statusArea = gen.StatusAreaHouse_STATUS_AREA_DNP
		case "G":
			statusArea = gen.StatusAreaHouse_STATUS_AREA_G
		case "F":
			statusArea = gen.StatusAreaHouse_STATUS_AREA_F
		case "PSP":
			statusArea = gen.StatusAreaHouse_STATUS_AREA_PSP
		}

		switch advert.HouseProperties.StatusHome {
		case "Live":
			statusHome = gen.StatusHomeHouse_STATUS_HOME_LIVE
		case "RepairNeed":
			statusHome = gen.StatusHomeHouse_STATUS_HOME_REPAIR_NEED
		case "CompleteNeed":
			statusHome = gen.StatusHomeHouse_STATUS_HOME_COMPLETE_NEED
		case "Renovation":
			statusHome = gen.StatusHomeHouse_STATUS_HOME_RENOVATION
		}

		houseProperties = &gen.HouseProperties{CeilingHeight: advert.HouseProperties.CeilingHeight, SquareArea: advert.HouseProperties.SquareArea, SquareHouse: advert.HouseProperties.SquareHouse, BedroomCount: int32(advert.HouseProperties.BedroomCount), StatusArea: statusArea, Cottage: advert.HouseProperties.Cottage, StatusHome: statusHome, Floor: int32(advert.HouseProperties.Floor)}
	}

	if advert.FlatProperties != nil {
		flatProperties = &gen.FlatProperties{CeilingHeight: advert.FlatProperties.CeilingHeight, RoomCount: int32(advert.FlatProperties.RoomCount), FloorGeneral: int32(advert.FlatProperties.FloorGeneral), Apartment: advert.FlatProperties.Apartment, SquareGeneral: advert.FlatProperties.SquareGeneral, Floor: int32(advert.FlatProperties.Floor), SquareResidential: advert.FlatProperties.SquareResidential}
	}

	if advert.ComplexProperties != nil {
		complexProperties = &gen.ComplexAdvertProperties{ComplexId: advert.ComplexProperties.ComplexId, NameComplex: advert.ComplexProperties.NameComplex, PhotoCompany: advert.ComplexProperties.PhotoCompany, NameCompany: advert.ComplexProperties.NameCompany}
	}

	h.logger.Info("success getting advert")
	// utils.LogSuccesResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, SignUpMethod)
	return &genAdverts.GetAdvertByIdResponse{Id: advert.ID, AdvertType: advert.AdvertType, TypeSale: advert.TypeSale,
		Title: advert.Title, Description: advert.Description, CountViews: advert.CountViews,
		CountLikes: advert.CountLikes, Price: advert.Price, Phone: advert.Phone, IsLiked: advert.IsLiked,
		IsAgent: advert.IsAgent, Metro: advert.Metro, Address: advert.Address, AddressPoint: advert.AddressPoint,
		PriceHistory: priceHistory, Images: images, HouseProp: houseProperties, FlatProperties: flatProperties,
		YearCreation: int32(advert.YearCreation), Material: material, ComplexProperties: complexProperties,
		DateCreation: advert.DateCreation.String()}, nil

}

func (h *AdvertsServerHandler) UpdateAdvertById(ctx context.Context, reqAdv *genAdverts.UpdateAdvertByIdRequest) (*genAdverts.UpdateAdvertByIdResponse, error) {
	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())

	var materl models.MaterialBuilding

	switch reqAdv.Material {
	case genAdverts.MaterialBuilding_MATERIAL_BRICK:
		materl = models.MaterialBrick
	case genAdverts.MaterialBuilding_MATERIAL_MONOLITHIC:
		materl = models.MaterialMonolithic
	case genAdverts.MaterialBuilding_MATERIAL_WOOD:
		materl = models.MaterialWood
	case genAdverts.MaterialBuilding_MATERIAL_PANEL:
		materl = models.MaterialPanel
	case genAdverts.MaterialBuilding_MATERIAL_STALINSKY:
		materl = models.MaterialStalinsky
	case genAdverts.MaterialBuilding_MATERIAL_BLOCK:
		materl = models.MaterialBlock
	case genAdverts.MaterialBuilding_MATERIAL_MONOLITHIC_BLOCK:
		materl = models.MaterialMonolithicBlock
	case genAdverts.MaterialBuilding_MATERIAL_FRAME:
		materl = models.MaterialFrame
	case genAdverts.MaterialBuilding_MATERIAL_AERATED_CONCRETE_BLOCK:
		materl = models.MaterialAeratedConcreteBlock
	case genAdverts.MaterialBuilding_MATERIAL_GAS_SILICATE_BLOCK:
		materl = models.MaterialGasSilicateBlock
	case genAdverts.MaterialBuilding_MATERIAL_FOAM_CONCRETE_BLOCK:
		materl = models.MaterialFoamConcreteBlock
	}

	var statusArea models.StatusAreaHouse
	var statusHome models.StatusHomeHouse
	var housePropert *models.HouseProperties = nil
	var flatPropert *models.FlatProperties = nil

	if reqAdv.HouseProp != nil {
		switch reqAdv.HouseProp.StatusArea {
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

		switch reqAdv.HouseProp.StatusHome {
		case genAdverts.StatusHomeHouse_STATUS_HOME_LIVE:
			statusHome = "Live"
		case genAdverts.StatusHomeHouse_STATUS_HOME_REPAIR_NEED:
			statusHome = "RepairNeed"
		case genAdverts.StatusHomeHouse_STATUS_HOME_COMPLETE_NEED:
			statusHome = "CompleteNeed"
		case genAdverts.StatusHomeHouse_STATUS_HOME_RENOVATION:
			statusHome = "Renovation"
		}

		housePropert = &models.HouseProperties{CeilingHeight: reqAdv.HouseProp.CeilingHeight,
			SquareArea: reqAdv.HouseProp.SquareArea, SquareHouse: reqAdv.HouseProp.SquareHouse,
			BedroomCount: int(reqAdv.HouseProp.BedroomCount), StatusArea: statusArea,
			Cottage: reqAdv.HouseProp.Cottage, StatusHome: statusHome, Floor: int(reqAdv.HouseProp.Floor)}
	}

	updData := &models.AdvertUpdateData{
		ID:          reqAdv.Id,
		TypeAdvert:  reqAdv.AdvertType,
		TypeSale:    reqAdv.TypeSale,
		Title:       reqAdv.Title,
		Description: reqAdv.Description,
		Price:       reqAdv.Price,
		Phone:       reqAdv.Phone,
		IsAgent:     reqAdv.IsAgent,
		Address: models.AddressData{
			Province:     reqAdv.Address.Province,
			Town:         reqAdv.Address.Town,
			Street:       reqAdv.Address.Street,
			House:        reqAdv.Address.House,
			Metro:        reqAdv.Address.Metro,
			AddressPoint: reqAdv.Address.AddressPoint,
		},
		HouseProperties: housePropert,
		FlatProperties:  flatPropert,
		Material:        materl,
	}

	err := h.uc.UpdateAdvertById(ctx, updData)

	if err != nil {
		h.logger.Error(ctx.Value("requestId").(string) + " " + err.Error())
		return &genAdverts.UpdateAdvertByIdResponse{RespCode: StatusBadRequest}, err
	}

	return &genAdverts.UpdateAdvertByIdResponse{RespCode: StatusOk}, nil

}

func (h *AdvertsServerHandler) LikeAdvert(ctx context.Context, reqAdv *genAdverts.LikeAdvertRequest) (*genAdverts.LikeAdvertResponse, error) {
	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())

	err := h.uc.LikeAdvert(ctx, reqAdv.AdvId, reqAdv.UserId)

	if err != nil {
		h.logger.Error(ctx.Value("requestId").(string) + " " + err.Error())
		return &genAdverts.LikeAdvertResponse{RespCode: StatusBadRequest}, err
	}

	return &genAdverts.LikeAdvertResponse{RespCode: StatusOk}, nil

}

func (h *AdvertsServerHandler) DislikeAdvert(ctx context.Context, reqAdv *genAdverts.DislikeAdvertRequest) (*genAdverts.DislikeAdvertResponse, error) {
	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())

	err := h.uc.DislikeAdvert(ctx, reqAdv.AdvId, reqAdv.UserId)

	if err != nil {
		h.logger.Error(ctx.Value("requestId").(string) + " " + err.Error())
		return &genAdverts.DislikeAdvertResponse{RespCode: StatusBadRequest}, err
	}

	return &genAdverts.DislikeAdvertResponse{RespCode: StatusOk}, nil

}

func (h *AdvertsServerHandler) IncreasePriority(ctx context.Context, reqAdv *genAdverts.IncreasePriorityRequest) (*genAdverts.IncreasePriorityResponse, error) {
	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())

	newPriority, err := h.uc.UpdatePriority(ctx, reqAdv.AdvertId, reqAdv.Amount)
	if err != nil {
		h.logger.Error(ctx.Value("requestId").(string) + " " + err.Error())
		return nil, err
	}

	return &genAdverts.IncreasePriorityResponse{Amount: newPriority}, nil
}

func (h *AdvertsServerHandler) GetPriority(ctx context.Context, reqAdv *genAdverts.GetPriorityRequest) (*genAdverts.GetPriorityResponse, error) {
	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())

	priority, err := h.uc.GetPriority(ctx, reqAdv.AdvertId)
	if err != nil {
		h.logger.Error(ctx.Value("requestId").(string) + " " + err.Error())
		return nil, err
	}

	return &genAdverts.GetPriorityResponse{Amount: priority}, nil
}

func (h *AdvertsServerHandler) CreateHouseAdvert(ctx context.Context, reqAdv *genAdverts.CreateHouseAdvertRequest) (*genAdverts.CreateHouseAdvertResponse, error) {
	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())

	var materl models.MaterialBuilding

	switch reqAdv.Material {
	case genAdverts.MaterialBuilding_MATERIAL_BRICK:
		materl = models.MaterialBrick
	case genAdverts.MaterialBuilding_MATERIAL_MONOLITHIC:
		materl = models.MaterialMonolithic
	case genAdverts.MaterialBuilding_MATERIAL_WOOD:
		materl = models.MaterialWood
	case genAdverts.MaterialBuilding_MATERIAL_PANEL:
		materl = models.MaterialPanel
	case genAdverts.MaterialBuilding_MATERIAL_STALINSKY:
		materl = models.MaterialStalinsky
	case genAdverts.MaterialBuilding_MATERIAL_BLOCK:
		materl = models.MaterialBlock
	case genAdverts.MaterialBuilding_MATERIAL_MONOLITHIC_BLOCK:
		materl = models.MaterialMonolithicBlock
	case genAdverts.MaterialBuilding_MATERIAL_FRAME:
		materl = models.MaterialFrame
	case genAdverts.MaterialBuilding_MATERIAL_AERATED_CONCRETE_BLOCK:
		materl = models.MaterialAeratedConcreteBlock
	case genAdverts.MaterialBuilding_MATERIAL_GAS_SILICATE_BLOCK:
		materl = models.MaterialGasSilicateBlock
	case genAdverts.MaterialBuilding_MATERIAL_FOAM_CONCRETE_BLOCK:
		materl = models.MaterialFoamConcreteBlock
	}

	var statusArea models.StatusAreaHouse
	var statusHome models.StatusHomeHouse

	if reqAdv.CreateHouseProp != nil {
		switch reqAdv.CreateHouseProp.StatusArea {
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

		switch reqAdv.CreateHouseProp.StatusHome {
		case genAdverts.StatusHomeHouse_STATUS_HOME_LIVE:
			statusHome = "Live"
		case genAdverts.StatusHomeHouse_STATUS_HOME_REPAIR_NEED:
			statusHome = "RepairNeed"
		case genAdverts.StatusHomeHouse_STATUS_HOME_COMPLETE_NEED:
			statusHome = "CompleteNeed"
		case genAdverts.StatusHomeHouse_STATUS_HOME_RENOVATION:
			statusHome = "Renovation"
		}
	}

	updData := &models.AdvertHouseCreateData{
		UserID:         reqAdv.UserId,
		AdvertTypeSale: models.TypePlacementAdvert(reqAdv.TypeSale),
		Title:          reqAdv.Title,
		Description:    reqAdv.Description,
		Phone:          reqAdv.Phone,
		IsAgent:        reqAdv.IsAgent,
		CeilingHeight:  reqAdv.CreateHouseProp.CeilingHeight,
		SquareArea:     reqAdv.CreateHouseProp.SquareArea,
		SquareHouse:    reqAdv.CreateHouseProp.SquareHouse,
		BedroomCount:   int(reqAdv.CreateHouseProp.BedroomCount),
		StatusArea:     statusArea,
		Cottage:        reqAdv.CreateHouseProp.Cottage,
		StatusHome:     statusHome,
		Price:          reqAdv.Price,
		FloorGeneral:   int(reqAdv.CreateHouseProp.Floor),
		Material:       materl,
		Address: models.AddressData{
			Province:     reqAdv.Address.Province,
			Town:         reqAdv.Address.Town,
			Street:       reqAdv.Address.Street,
			House:        reqAdv.Address.House,
			Metro:        reqAdv.Address.Metro,
			AddressPoint: reqAdv.Address.AddressPoint,
		},
		YearCreation: int(reqAdv.YearCreation),
	}

	gotadv, err := h.uc.CreateHouseAdvert(ctx, updData)

	if err != nil {
		h.logger.Error(ctx.Value("requestId").(string) + " " + err.Error())
		return &genAdverts.CreateHouseAdvertResponse{RespCode: StatusBadRequest}, err
	}

	return &genAdverts.CreateHouseAdvertResponse{
		Id:           gotadv.ID,
		UserId:       gotadv.UserID,
		TypeSale:     string(gotadv.AdvertTypeSale),
		Title:        gotadv.Title,
		Description:  gotadv.Description,
		Phone:        gotadv.Phone,
		IsAgent:      gotadv.IsAgent,
		Priority:     int64(gotadv.Priority),
		DateCreation: gotadv.DateCreation.String(),
		RespCode:     StatusOk}, nil

}

func (h *AdvertsServerHandler) GetExistBuildingByAddress(ctx context.Context, reqAdv *genAdverts.GetExistBuildingByAddressRequest) (*genAdverts.GetExistBuildingByAddressResponse, error) {
	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())

	addrr := &models.AddressData{
		Province:     reqAdv.AdrData.Province,
		Town:         reqAdv.AdrData.Town,
		Street:       reqAdv.AdrData.Street,
		House:        reqAdv.AdrData.House,
		Metro:        reqAdv.AdrData.Metro,
		AddressPoint: reqAdv.AdrData.AddressPoint,
	}

	gotBuilding, err := h.uc.GetExistBuildingByAddress(ctx, addrr)

	if err != nil {
		h.logger.Error(ctx.Value("requestId").(string) + " " + err.Error())
		return &genAdverts.GetExistBuildingByAddressResponse{RespCode: StatusBadRequest}, err
	}

	var mater genAdverts.MaterialBuilding
	switch gotBuilding.Material {
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

	return &genAdverts.GetExistBuildingByAddressResponse{
		ComplexName:  gotBuilding.ComplexName,
		Floor:        int32(gotBuilding.Floor),
		Material:     mater,
		YearCreation: int32(gotBuilding.YearCreation),
		RespCode:     StatusOk}, nil

}

func (h *AdvertsServerHandler) CreateFlatAdvert(ctx context.Context, reqAdv *genAdverts.CreateFlatAdvertRequest) (*genAdverts.CreateFlatAdvertResponse, error) {
	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())

	var materl models.MaterialBuilding
	switch reqAdv.Material {
	case genAdverts.MaterialBuilding_MATERIAL_BRICK:
		materl = models.MaterialBrick
	case genAdverts.MaterialBuilding_MATERIAL_MONOLITHIC:
		materl = models.MaterialMonolithic
	case genAdverts.MaterialBuilding_MATERIAL_WOOD:
		materl = models.MaterialWood
	case genAdverts.MaterialBuilding_MATERIAL_PANEL:
		materl = models.MaterialPanel
	case genAdverts.MaterialBuilding_MATERIAL_STALINSKY:
		materl = models.MaterialStalinsky
	case genAdverts.MaterialBuilding_MATERIAL_BLOCK:
		materl = models.MaterialBlock
	case genAdverts.MaterialBuilding_MATERIAL_MONOLITHIC_BLOCK:
		materl = models.MaterialMonolithicBlock
	case genAdverts.MaterialBuilding_MATERIAL_FRAME:
		materl = models.MaterialFrame
	case genAdverts.MaterialBuilding_MATERIAL_AERATED_CONCRETE_BLOCK:
		materl = models.MaterialAeratedConcreteBlock
	case genAdverts.MaterialBuilding_MATERIAL_GAS_SILICATE_BLOCK:
		materl = models.MaterialGasSilicateBlock
	case genAdverts.MaterialBuilding_MATERIAL_FOAM_CONCRETE_BLOCK:
		materl = models.MaterialFoamConcreteBlock
	}

	gotFlat, err := h.uc.CreateFlatAdvert(ctx, &models.AdvertFlatCreateData{
		UserID:            reqAdv.UserId,
		AdvertTypeSale:    models.TypePlacementAdvert(reqAdv.TypeSale),
		Title:             reqAdv.Title,
		Description:       reqAdv.Description,
		Phone:             reqAdv.Phone,
		IsAgent:           reqAdv.IsAgent,
		Floor:             int(reqAdv.CreateFlatProp.Floor),
		CeilingHeight:     reqAdv.CreateFlatProp.CeilingHeight,
		SquareGeneral:     reqAdv.CreateFlatProp.SquareGeneral,
		RoomCount:         int(reqAdv.CreateFlatProp.RoomCount),
		SquareResidential: reqAdv.CreateFlatProp.SquareResidential,
		Apartment:         reqAdv.CreateFlatProp.Apartment,
		Price:             reqAdv.Price,
		FloorGeneral:      int(reqAdv.CreateFlatProp.FloorGeneral),
		Material:          materl,
		Address: models.AddressData{
			Province:     reqAdv.Address.Province,
			Town:         reqAdv.Address.Town,
			Street:       reqAdv.Address.Street,
			House:        reqAdv.Address.House,
			Metro:        reqAdv.Address.Metro,
			AddressPoint: reqAdv.Address.AddressPoint,
		},
		YearCreation: int(reqAdv.YearCreation),
	})

	if err != nil {
		h.logger.Error(ctx.Value("requestId").(string) + " " + err.Error())
		return &genAdverts.CreateFlatAdvertResponse{RespCode: StatusBadRequest}, err
	}

	return &genAdverts.CreateFlatAdvertResponse{
		Id:           gotFlat.ID,
		UserId:       gotFlat.UserID,
		TypeSale:     string(gotFlat.AdvertTypeSale),
		Title:        gotFlat.Title,
		Description:  gotFlat.Description,
		Phone:        gotFlat.Phone,
		IsAgent:      gotFlat.IsAgent,
		Priority:     int64(gotFlat.Priority),
		DateCreation: gotFlat.DateCreation.Format("2006-01-02 15:04:05"),
		RespCode:     StatusOk}, nil

}

func (h *AdvertsServerHandler) GetSquareAdvertsList(ctx context.Context, req *genAdverts.GetSquareAdvertsListRequest) (*genAdverts.GetSquareAdvertsListResponse, error) {

	utils.LogSuccesResponse(h.logger, "HERE GRPC ONE", utils.DeliveryLayer, "GetSquareAdvertsList")

	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())
	utils.LogSuccesResponse(h.logger, "HERE GRPC TWO", utils.DeliveryLayer, "GetSquareAdvertsList")

	gotList, err := h.uc.GetSquareAdvertsList(ctx, int(req.PageSize), int(req.Offset))

	if err != nil {
		h.logger.Error(ctx.Value("requestId").(string) + " " + err.Error())
		return &genAdverts.GetSquareAdvertsListResponse{RespCode: StatusBadRequest}, err
	}
	utils.LogSuccesResponse(h.logger, "HERE GRPC THREE", utils.DeliveryLayer, "GetSquareAdvertsList")

	foundAdverts := make([]*genAdverts.AdvertSquareData, 0)

	for _, adv := range gotList {
		newadv := &genAdverts.AdvertSquareData{
			Id:           adv.ID,
			TypeAdvert:   adv.TypeAdvert,
			Photo:        adv.Photo,
			TypeSale:     adv.TypeSale,
			Address:      adv.Address,
			Metro:        adv.Metro,
			DateCreation: adv.DateCreation.String(),
		}
		if adv.HouseProperties != nil {
			newadv.HouseProp = &gen.HouseSquareProperties{
				SquareArea:   adv.HouseProperties.SquareArea,
				SquareHouse:  adv.HouseProperties.SquareHouse,
				BedroomCount: int32(adv.HouseProperties.BedroomCount),
				Cottage:      adv.HouseProperties.Cottage,
				Floor:        int32(adv.HouseProperties.Floor)}
		}

		if adv.FlatProperties != nil {
			newadv.FlatProp = &gen.FlatSquareProperties{
				RoomCount:     int32(adv.FlatProperties.RoomCount),
				FloorGeneral:  int32(adv.FlatProperties.FloorGeneral),
				SquareGeneral: adv.FlatProperties.SquareGeneral,
				Floor:         int32(adv.FlatProperties.Floor)}
		}

		utils.LogSuccesResponse(h.logger, "HERE GRPC FOUR", utils.DeliveryLayer, "GetSquareAdvertsList")

		foundAdverts = append(foundAdverts, newadv)
	}
	utils.LogSuccesResponse(h.logger, "HERE GRPC FiVE", utils.DeliveryLayer, "GetSquareAdvertsList")

	return &genAdverts.GetSquareAdvertsListResponse{
		SquareData: foundAdverts,
		RespCode:   StatusOk}, nil

}

func (h *AdvertsServerHandler) DeleteAdvertById(ctx context.Context, reqAdv *genAdverts.DeleteAdvertByIdRequest) (*genAdverts.DeleteAdvertByIdResponse, error) {
	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())

	err := h.uc.DeleteAdvertById(ctx, reqAdv.Id)

	if err != nil {
		h.logger.Error(ctx.Value("requestId").(string) + " " + err.Error())
		return &genAdverts.DeleteAdvertByIdResponse{RespCode: StatusBadRequest}, err
	}

	return &genAdverts.DeleteAdvertByIdResponse{RespCode: StatusOk}, nil

}

// GetRectangleAdvertsList handles the request for retrieving a rectangle adverts with search.
func (h *AdvertsServerHandler) GetRectangleAdvertsList(ctx context.Context, req *genAdverts.GetRectangleAdvertsListRequest) (*genAdverts.GetRectangleAdvertsListResponse, error) {
	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())
	advList, err := h.uc.GetRectangleAdvertsList(ctx, models.AdvertFilter{
		MinPrice:   req.MinPrice,
		MaxPrice:   req.MaxPrice,
		Page:       int(req.Size),
		Offset:     int(req.Page),
		RoomCount:  int(req.RoomCount),
		Address:    req.Address,
		DealType:   req.DealType,
		AdvertType: req.AdvertType,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	protoAdverts := make([]*gen.AdvertRectangleData, 0, len(advList.Adverts))

	for _, advert := range advList.Adverts {
		var statusArea gen.StatusAreaHouse
		var statusHome gen.StatusHomeHouse
		var houseProperties *gen.HouseProperties = nil
		var flatProperties *gen.FlatProperties = nil
		if advert.HouseProperties != nil {
			houseProperties = &gen.HouseProperties{SquareArea: advert.HouseProperties.SquareArea, SquareHouse: advert.HouseProperties.SquareHouse, BedroomCount: int32(advert.HouseProperties.BedroomCount), StatusArea: statusArea, Cottage: advert.HouseProperties.Cottage, StatusHome: statusHome, Floor: int32(advert.HouseProperties.Floor)}
		}

		if advert.FlatProperties != nil {
			flatProperties = &gen.FlatProperties{RoomCount: int32(advert.FlatProperties.RoomCount), FloorGeneral: int32(advert.FlatProperties.FloorGeneral), SquareGeneral: advert.FlatProperties.SquareGeneral, Floor: int32(advert.FlatProperties.Floor)}
		}
		protoAdverts = append(protoAdverts, &genAdverts.AdvertRectangleData{Id: advert.ID, AdvertType: advert.TypeAdvert, TypeSale: advert.TypeSale, Title: advert.Title, Description: advert.Description, Price: int64(advert.Price), Phone: advert.Phone, IsLiked: advert.IsLiked, Address: advert.Address, AddressPoint: advert.AddressPoint, Photo: advert.Photo, HouseProp: houseProperties, FlatProperties: flatProperties, DateCreation: advert.DateCreation.String(), Rating: advert.Rating})
	}

	return &genAdverts.GetRectangleAdvertsListResponse{Adverts: protoAdverts, Info: &genAdverts.PageInfo{TotalElements: int64(advList.PageInfo.TotalElements), TotalPages: int64(advList.PageInfo.TotalPages), PageSize: int64(advList.PageInfo.PageSize), CurrentPage: int64(advList.PageInfo.CurrentPage)}}, nil
}

func (h *AdvertsServerHandler) GetRectangleAdvertsByUser(ctx context.Context, req *genAdverts.GetUserAdvertsRequest) (*genAdverts.GetUserAdvertsResponse, error) {
	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())
	advList, err := h.uc.GetRectangleAdvertsByUserId(ctx, int(req.Size), int(req.Page), req.UserId) // todo МБ тут ошибка и надо поменять местами int(req.Size), int(req.Page)

	if err != nil {
		log.Println(err)
		return &genAdverts.GetUserAdvertsResponse{
			RespCode: StatusBadRequest}, err
	}

	protoAdverts := make([]*gen.AdvertRectangleData, 0, len(advList))

	for _, advert := range advList {
		var statusArea gen.StatusAreaHouse
		var statusHome gen.StatusHomeHouse
		var houseProperties *gen.HouseProperties = nil
		var flatProperties *gen.FlatProperties = nil
		if advert.HouseProperties != nil {
			houseProperties = &gen.HouseProperties{SquareArea: advert.HouseProperties.SquareArea, SquareHouse: advert.HouseProperties.SquareHouse, BedroomCount: int32(advert.HouseProperties.BedroomCount), StatusArea: statusArea, Cottage: advert.HouseProperties.Cottage, StatusHome: statusHome, Floor: int32(advert.HouseProperties.Floor)}
		}

		if advert.FlatProperties != nil {
			flatProperties = &gen.FlatProperties{RoomCount: int32(advert.FlatProperties.RoomCount), FloorGeneral: int32(advert.FlatProperties.FloorGeneral), SquareGeneral: advert.FlatProperties.SquareGeneral, Floor: int32(advert.FlatProperties.Floor)}
		}
		protoAdverts = append(protoAdverts, &genAdverts.AdvertRectangleData{Id: advert.ID, AdvertType: advert.TypeAdvert, TypeSale: advert.TypeSale, Title: advert.Title, Description: advert.Description, Price: int64(advert.Price), Phone: advert.Phone, IsLiked: advert.IsLiked, Address: advert.Address, AddressPoint: advert.AddressPoint, Photo: advert.Photo, HouseProp: houseProperties, FlatProperties: flatProperties, DateCreation: advert.DateCreation.String(), Rating: advert.Rating})
	}

	return &genAdverts.GetUserAdvertsResponse{
		RectDataSlice: protoAdverts,
		RespCode:      StatusOk,
	}, nil
}

func (h *AdvertsServerHandler) GetLikedUserAdverts(ctx context.Context, req *genAdverts.GetLikedUserAdvertsRequest) (*genAdverts.GetLikedUserAdvertsResponse, error) {
	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())
	advList, err := h.uc.GetRectangleAdvertsLikedByUserId(ctx, int(req.PageSize), int(req.Offset), req.UserId) // todo МБ тут ошибка и надо поменять местами int(req.Size), int(req.Page)

	if err != nil {
		log.Println(err)
		return &genAdverts.GetLikedUserAdvertsResponse{
			RespCode: StatusBadRequest}, err
	}

	protoAdverts := make([]*gen.AdvertRectangleData, 0, len(advList))

	for _, advert := range advList {
		var statusArea gen.StatusAreaHouse
		var statusHome gen.StatusHomeHouse
		var houseProperties *gen.HouseProperties = nil
		var flatProperties *gen.FlatProperties = nil
		if advert.HouseProperties != nil {
			houseProperties = &gen.HouseProperties{SquareArea: advert.HouseProperties.SquareArea, SquareHouse: advert.HouseProperties.SquareHouse, BedroomCount: int32(advert.HouseProperties.BedroomCount), StatusArea: statusArea, Cottage: advert.HouseProperties.Cottage, StatusHome: statusHome, Floor: int32(advert.HouseProperties.Floor)}
		}

		if advert.FlatProperties != nil {
			flatProperties = &gen.FlatProperties{RoomCount: int32(advert.FlatProperties.RoomCount), FloorGeneral: int32(advert.FlatProperties.FloorGeneral), SquareGeneral: advert.FlatProperties.SquareGeneral, Floor: int32(advert.FlatProperties.Floor)}
		}
		protoAdverts = append(protoAdverts, &genAdverts.AdvertRectangleData{Id: advert.ID, AdvertType: advert.TypeAdvert, TypeSale: advert.TypeSale, Title: advert.Title, Description: advert.Description, Price: int64(advert.Price), Phone: advert.Phone, IsLiked: advert.IsLiked, Address: advert.Address, AddressPoint: advert.AddressPoint, Photo: advert.Photo, HouseProp: houseProperties, FlatProperties: flatProperties, DateCreation: advert.DateCreation.String()})
	}

	return &genAdverts.GetLikedUserAdvertsResponse{
		RectDataSlice: protoAdverts,
		RespCode:      StatusOk,
	}, nil
}

func (h *AdvertsServerHandler) GetRectangleAdvertsByComplex(ctx context.Context, req *genAdverts.GetComplexAdvertsRequest) (*genAdverts.GetComplexAdvertsResponse, error) {
	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())
	advList, err := h.uc.GetRectangleAdvertsByComplexId(ctx, int(req.PageSize), int(req.Offset), req.ComplexId) // todo МБ тут ошибка и надо поменять местами int(req.Size), int(req.Page)

	if err != nil {
		log.Println(err)
		return &genAdverts.GetComplexAdvertsResponse{
			RespCode: StatusBadRequest}, err
	}

	protoAdverts := make([]*gen.AdvertRectangleData, 0, len(advList))

	for _, advert := range advList {
		var statusArea gen.StatusAreaHouse
		var statusHome gen.StatusHomeHouse
		var houseProperties *gen.HouseProperties = nil
		var flatProperties *gen.FlatProperties = nil
		if advert.HouseProperties != nil {
			houseProperties = &gen.HouseProperties{SquareArea: advert.HouseProperties.SquareArea, SquareHouse: advert.HouseProperties.SquareHouse, BedroomCount: int32(advert.HouseProperties.BedroomCount), StatusArea: statusArea, Cottage: advert.HouseProperties.Cottage, StatusHome: statusHome, Floor: int32(advert.HouseProperties.Floor)}
		}

		if advert.FlatProperties != nil {
			flatProperties = &gen.FlatProperties{RoomCount: int32(advert.FlatProperties.RoomCount), FloorGeneral: int32(advert.FlatProperties.FloorGeneral), SquareGeneral: advert.FlatProperties.SquareGeneral, Floor: int32(advert.FlatProperties.Floor)}
		}
		protoAdverts = append(protoAdverts, &genAdverts.AdvertRectangleData{Id: advert.ID, AdvertType: advert.TypeAdvert, TypeSale: advert.TypeSale, Title: advert.Title, Description: advert.Description, Price: int64(advert.Price), Phone: advert.Phone, IsLiked: advert.IsLiked, Address: advert.Address, AddressPoint: advert.AddressPoint, Photo: advert.Photo, HouseProp: houseProperties, FlatProperties: flatProperties, DateCreation: advert.DateCreation.String()})
	}

	return &genAdverts.GetComplexAdvertsResponse{
		RectDataSlice: protoAdverts,
		RespCode:      StatusOk,
	}, nil
}
