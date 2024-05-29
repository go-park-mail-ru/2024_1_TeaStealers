package grpc

import (
	"2024_1_TeaStealers/internal/models"
	complex "2024_1_TeaStealers/internal/pkg/complexes"
	genComplex "2024_1_TeaStealers/internal/pkg/complexes/delivery/grpc/gen"
	"context"
	"log"

	"go.uber.org/zap"
)

// ComplexServerHandler handles HTTP requests for complexes.
type ComplexServerHandler struct {
	genComplex.ComplexServer
	// uc represents the usecase interface for authentication.
	uc     complex.ComplexUsecase
	logger *zap.Logger
}

// NewComplexServerHandler creates a new instance of AuthHandler.
func NewComplexServerHandler(uc complex.ComplexUsecase, logger *zap.Logger) *ComplexServerHandler {
	return &ComplexServerHandler{uc: uc, logger: logger}
}

func (h *ComplexServerHandler) CreateComplex(ctx context.Context, req *genComplex.CreateComplexRequest) (*genComplex.CreateComplexResponse, error) {

	var classHousing models.ClassHouse
	switch req.ClassHousing {
	case genComplex.ClassHouse_CLASS_HOUSE_BUSINESS:
		classHousing = models.ClassHouseBusiness
	case genComplex.ClassHouse_CLASS_HOUSE_PREMIUM:
		classHousing = models.ClassHousePremium
	case genComplex.ClassHouse_CLASS_HOUSE_ELITE:
		classHousing = models.ClassHouseElite
	case genComplex.ClassHouse_CLASS_HOUSE_ECONOM:
		classHousing = models.ClassHouseEconom
	case genComplex.ClassHouse_CLASS_HOUSE_COMFORT:
		classHousing = models.ClassHouseComfort
	}

	data := models.ComplexCreateData{CompanyId: req.CompanyId, Name: req.Name, Address: req.Address, Description: req.Description, WithoutFinishingOption: req.WithoutFinishingOption, FinishingOption: req.FinishingOption, PreFinishingOption: req.PreFinishingOption, ClassHousing: classHousing, Parking: req.Parking, Security: req.Security}
	data.Sanitize()

	newComplex, err := h.uc.CreateComplex(ctx, &data)
	if err != nil {
		log.Println(err)
		return &genComplex.CreateComplexResponse{RespCode: 400}, err
	}
	newComplex.Sanitize()

	return &genComplex.CreateComplexResponse{RespCode: 200}, nil
}

func (h *ComplexServerHandler) UpdateComplexPhoto(ctx context.Context, req *genComplex.UpdateComplexPhotoRequest) (*genComplex.UpdateComplexPhotoResponse, error) {
	fileName, err := h.uc.UpdateComplexPhoto(ctx, req.ComplexId, req.FileName)
	if err != nil {
		log.Println(err)
		return &genComplex.UpdateComplexPhotoResponse{RespCode: 400}, err
	}

	return &genComplex.UpdateComplexPhotoResponse{Filename: fileName, RespCode: 200}, nil
}

// GetComplexById handles the request for getting complex by id
func (h *ComplexServerHandler) GetComplexById(ctx context.Context, req *genComplex.GetComplexByIdRequest) (*genComplex.GetComplexByIdResponse, error) {
	complexData, err := h.uc.GetComplexById(ctx, req.Id)
	if err != nil {
		log.Println(err)
		return &genComplex.GetComplexByIdResponse{RespCode: 400}, err
	}

	var classHousing genComplex.ClassHouse

	switch complexData.ClassHousing {
	case models.ClassHouseBusiness:
		classHousing = genComplex.ClassHouse_CLASS_HOUSE_BUSINESS
	case models.ClassHousePremium:
		classHousing = genComplex.ClassHouse_CLASS_HOUSE_PREMIUM
	case models.ClassHouseElite:
		classHousing = genComplex.ClassHouse_CLASS_HOUSE_ELITE
	case models.ClassHouseEconom:
		classHousing = genComplex.ClassHouse_CLASS_HOUSE_ECONOM
	case models.ClassHouseComfort:
		classHousing = genComplex.ClassHouse_CLASS_HOUSE_COMFORT
	}

	return &genComplex.GetComplexByIdResponse{Id: complexData.ID, CompanyId: complexData.CompanyId, Name: complexData.Name, Address: complexData.Address, Photo: complexData.Photo, Description: complexData.Description, DateBeginBuild: complexData.DateBeginBuild.String(), DateEndBuild: complexData.DateEndBuild.String(), WithoutFinishingOption: complexData.WithoutFinishingOption, FinishingOption: complexData.FinishingOption, PreFinishingOption: complexData.PreFinishingOption, ClassHousing: classHousing, Parking: complexData.Parking, Security: complexData.Security, RespCode: 200}, nil
}

func (h *ComplexServerHandler) CreateCompany(ctx context.Context, req *genComplex.CreateCompanyRequest) (*genComplex.CreateCompanyResponse, error) {
	data := models.CompanyCreateData{Name: req.Name, YearFounded: int(req.YearFounded), Phone: req.Phone, Description: req.Description}
	data.Sanitize()

	newCompany, err := h.uc.CreateCompany(ctx, &data)
	if err != nil {
		log.Println(err)
		return &genComplex.CreateCompanyResponse{RespCode: 400}, err
	}
	newCompany.Sanitize()

	return &genComplex.CreateCompanyResponse{RespCode: 200}, nil
}

func (h *ComplexServerHandler) UpdateCompanyPhoto(ctx context.Context, req *genComplex.UpdateCompanyPhotoRequest) (*genComplex.UpdateCompanyPhotoResponse, error) {
	fileName, err := h.uc.UpdateCompanyPhoto(ctx, req.CompanyId, req.FileName)
	if err != nil {
		log.Println(err)
		return &genComplex.UpdateCompanyPhotoResponse{RespCode: 400}, err
	}

	return &genComplex.UpdateCompanyPhotoResponse{Filename: fileName, RespCode: 200}, nil
}

// GetCompanyById handles the request for getting company by id
func (h *ComplexServerHandler) GetCompanyById(ctx context.Context, req *genComplex.GetCompanyByIdRequest) (*genComplex.GetCompanyByIdResponse, error) {

	companyData, err := h.uc.GetCompanyById(ctx, req.Id)
	if err != nil {
		log.Println(err)
		return &genComplex.GetCompanyByIdResponse{RespCode: 400}, err
	}

	return &genComplex.GetCompanyByIdResponse{Id: companyData.ID, Name: companyData.Name, Photo: companyData.Photo, YearFounded: int32(companyData.YearFounded), Phone: companyData.Phone, Description: companyData.Description, RespCode: 200}, nil
}
