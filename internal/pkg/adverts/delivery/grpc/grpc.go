package grpc

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/adverts"
	"2024_1_TeaStealers/internal/pkg/adverts/delivery/grpc/gen"
	genAdverts "2024_1_TeaStealers/internal/pkg/adverts/delivery/grpc/gen"
	"context"
	"log"

	"github.com/satori/uuid"
	"go.uber.org/zap"
)

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

func (h *AdvertsServerHandler) GetAdvertById(ctx context.Context, req *genAdverts.GetAdvertByIdRequest) (*genAdverts.GetAdvertByIdResponse, error) {
	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())

	advert, err := h.uc.GetAdvertById(ctx, req.Id)

	if err != nil {

		h.logger.Error(err.Error())
		// utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, SignUpMethod, err, http.StatusBadRequest)
		return nil, err
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
	return &genAdverts.GetAdvertByIdResponse{Id: advert.ID, AdvertType: advert.AdvertType, TypeSale: advert.TypeSale, Title: advert.Title, Description: advert.Description, CountViews: advert.CountViews, CountLikes: advert.CountLikes, Price: advert.Price, Phone: advert.Phone, IsLiked: advert.IsLiked, IsAgent: advert.IsAgent, Metro: advert.Metro, Address: advert.Address, AddressPoint: advert.AddressPoint, PriceHistory: priceHistory, Images: images, HouseProperties: houseProperties, FlatProperties: flatProperties, YearCreation: int32(advert.YearCreation), Material: material, ComplexProperties: complexProperties, DateCreation: advert.DateCreation.String()}, nil

}

// GetRectangeAdvertsList handles the request for retrieving a rectangle adverts with search.
func (h *AdvertsServerHandler) GetRectangleAdvertsList(ctx context.Context, req *genAdverts.GetRectangleAdvertsListRequest) (*genAdverts.GetRectangleAdvertsListResponse, error) {
	adverts, err := h.uc.GetRectangleAdvertsList(ctx, models.AdvertFilter{
		MinPrice:   req.MinPrice,
		MaxPrice:   req.MaxPrice,
		Page:       int(req.Page),
		Offset:     int(req.Size),
		RoomCount:  int(req.RoomCount),
		Address:    req.Address,
		DealType:   req.DealType,
		AdvertType: req.AdvertType,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	protoAdverts := make([]*gen.AdvertRectangleData, 0, len(adverts.Adverts))

	for _, advert := range adverts.Adverts {
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
		protoAdverts = append(protoAdverts, &genAdverts.AdvertRectangleData{Id: advert.ID, AdvertType: advert.TypeAdvert, TypeSale: advert.TypeSale, Title: advert.Title, Description: advert.Description, Price: int64(advert.Price), Phone: advert.Phone, IsLiked: advert.IsLiked, Metro: advert.Metro, Address: advert.Address, AddressPoint: advert.AddressPoint, Photo: advert.Photo, HouseProperties: houseProperties, FlatProperties: flatProperties, DateCreation: advert.DateCreation.String()})
	}

	return &genAdverts.GetRectangleAdvertsListResponse{Adverts: protoAdverts, Info: &genAdverts.PageInfo{TotalElements: int64(adverts.PageInfo.TotalElements), TotalPages: int64(adverts.PageInfo.TotalPages), PageSize: int64(adverts.PageInfo.PageSize), CurrentPage: int64(adverts.PageInfo.CurrentPage)}}, nil
}
