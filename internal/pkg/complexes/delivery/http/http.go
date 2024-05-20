package delivery

import (
	"2024_1_TeaStealers/internal/models"
	genComplex "2024_1_TeaStealers/internal/pkg/complexes/delivery/grpc/gen"
	"2024_1_TeaStealers/internal/pkg/utils"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/satori/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// UserClientHandler handles HTTP requests for user.
type ComplexClientHandler struct {
	client genComplex.ComplexClient
	logger *zap.Logger
}

// NewClientUserHandler creates a new instance of UserHandler.
func NewClientComplexHandler(grpcConn *grpc.ClientConn, logger *zap.Logger) *ComplexClientHandler {
	return &ComplexClientHandler{client: genComplex.NewComplexClient(grpcConn), logger: logger}
}

func (h *ComplexClientHandler) CreateComplex(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("csrftoken")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "csrf cookie not found")
		return
	}
	data := models.ComplexCreateData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}
	data.Sanitize()

	var classHousing genComplex.ClassHouse

	switch data.ClassHousing {
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

	newComplex, err := h.client.CreateComplex(r.Context(), &genComplex.CreateComplexRequest{CompanyId: data.CompanyId, Name: data.Name, Address: data.Address, Description: data.Description, DateBeginBuild: data.DateBeginBuild.String(), DateEndBuild: data.DateEndBuild.String(), WithoutFinishingOption: data.WithoutFinishingOption, FinishingOption: data.FinishingOption, PreFinishingOption: data.PreFinishingOption, ClassHousing: classHousing, Parking: data.Parking, Security: data.Security})
	//newComplex.Sanitize()
	if err != nil {
		log.Println(err)
		utils.WriteError(w, http.StatusBadRequest, "data already is used")
		return
	}

	if err = utils.WriteResponse(w, http.StatusCreated, newComplex); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *ComplexClientHandler) UpdateComplexPhoto(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("csrftoken")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "csrf cookie not found")
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	complexId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}
	if err := r.ParseMultipartForm(5 << 20); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "max size file 5 mb")
		return
	}

	file, head, err := r.FormFile("file")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "bad data request")
		return
	}
	defer file.Close()

	allowedExtensions := []string{".jpg", ".jpeg", ".png"}
	fileType := strings.ToLower(filepath.Ext(head.Filename))
	if !slices.Contains(allowedExtensions, fileType) {
		utils.WriteError(w, http.StatusBadRequest, "jpg, jpeg, png only")
		return
	}

	newId := uuid.NewV4()
	newFileName := newId.String() + fileType
	subDirectory := "complexes"
	directory := filepath.Join(os.Getenv("DOCKER_DIR"), subDirectory)
	if err := os.MkdirAll(directory, 0755); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "File system error")
		return
	}
	destination, err := os.Create(directory + "/" + newFileName)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "File system error")
	}
	defer destination.Close()
	_, err = io.Copy(destination, file)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "File system error")
	}

	fileName, err := h.client.UpdateComplexPhoto(r.Context(), &genComplex.UpdateComplexPhotoRequest{ComplexId: complexId, FileName: subDirectory + "/" + newFileName})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "failed upload file")
		return
	}
	if err := utils.WriteResponse(w, http.StatusOK, fileName); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	}
}

// GetComplexById handles the request for getting complex by id
func (h *ComplexClientHandler) GetComplexById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	complexId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	complexData, err := h.client.GetComplexById(r.Context(), &genComplex.GetComplexByIdRequest{Id: complexId})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	//complexData.Sanitize()

	if err = utils.WriteResponse(w, http.StatusOK, complexData); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *ComplexClientHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("csrftoken")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "csrf cookie not found")
		return
	}
	data := models.CompanyCreateData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}
	data.Sanitize()

	newCompany, err := h.client.CreateCompany(r.Context(), &genComplex.CreateCompanyRequest{Name: data.Name, YearFounded: int32(data.YearFounded), Phone: data.Phone, Description: data.Description})
	if err != nil {
		log.Println(err)
		utils.WriteError(w, http.StatusBadRequest, "data already is used")
		return
	}
	//newCompany.Sanitize()

	if err = utils.WriteResponse(w, http.StatusCreated, newCompany); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *ComplexClientHandler) UpdateCompanyPhoto(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("csrftoken")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "csrf cookie not found")
		return
	}
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	companyId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}
	if err := r.ParseMultipartForm(5 << 20); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "max size file 5 mb")
		return
	}

	file, head, err := r.FormFile("file")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "bad data request")
		return
	}
	defer file.Close()

	allowedExtensions := []string{".jpg", ".jpeg", ".png"}
	fileType := strings.ToLower(filepath.Ext(head.Filename))
	if !slices.Contains(allowedExtensions, fileType) {
		utils.WriteError(w, http.StatusBadRequest, "jpg, jpeg, png only")
		return
	}

	newId := uuid.NewV4()
	newFileName := newId.String() + fileType
	subDirectory := "companies"
	directory := filepath.Join(os.Getenv("DOCKER_DIR"), subDirectory)
	if err := os.MkdirAll(directory, 0755); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "File system error")
		return
	}
	destination, err := os.Create(directory + "/" + newFileName)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "File system error")
		return
	}
	defer destination.Close()
	_, err = io.Copy(destination, file)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "File system error")
		return
	}

	fileName, err := h.client.UpdateCompanyPhoto(r.Context(), &genComplex.UpdateCompanyPhotoRequest{CompanyId: companyId, FileName: subDirectory + "/" + newFileName})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "failed upload file")
		return
	}
	if err := utils.WriteResponse(w, http.StatusOK, fileName); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	}
}

// GetCompanyById handles the request for getting company by id
func (h *ComplexClientHandler) GetCompanyById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	companyId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	companyData, err := h.client.GetCompanyById(r.Context(), &genComplex.GetCompanyByIdRequest{Id: companyId})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	//companyData.Sanitize()

	if err = utils.WriteResponse(w, http.StatusOK, companyData); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}
