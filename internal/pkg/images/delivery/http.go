package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/images"
	"2024_1_TeaStealers/internal/pkg/utils"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/satori/uuid"
)

// ImageHandler handles HTTP requests for manage image.
type ImageHandler struct {
	// uc represents the usecase interface for manage image.
	uc images.ImageUsecase
}

// NewImageHandler creates a new instance of ImageHandler.
func NewImageHandler(uc images.ImageUsecase) *ImageHandler {
	return &ImageHandler{uc: uc}
}

// ParseImageAndData parses the image and data from the request and saves the image to disk.
func ParseImageAndData(r *http.Request) (models.ImageCreateData, uuid.UUID, error) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		return models.ImageCreateData{}, uuid.UUID{}, err
	}

	advertIdStr := r.FormValue("advert_id")
	priorityStr := r.FormValue("priority")

	priority, err := strconv.Atoi(priorityStr)
	if err != nil {
		return models.ImageCreateData{}, uuid.UUID{}, err
	}

	advertId, err := uuid.FromString(advertIdStr)
	if err != nil {
		return models.ImageCreateData{}, uuid.UUID{}, err
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		return models.ImageCreateData{}, uuid.UUID{}, err
	}
	defer file.Close()

	idImage := uuid.NewV4()
	imagePath := os.Getenv("BASE_DIR") + advertId.String() + "|" + idImage.String() + ".jpg"
	dst, err := os.Create(imagePath)
	if err != nil {
		return models.ImageCreateData{}, uuid.UUID{}, err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return models.ImageCreateData{}, uuid.UUID{}, err
	}

	return models.ImageCreateData{Priority: priority, AdvertId: advertId}, idImage, nil
}

// CreateImage handles the request for creating a new image.
func (h *ImageHandler) CreateImage(w http.ResponseWriter, r *http.Request) {
	data, idImage, err := ParseImageAndData(r)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	newImage, err := h.uc.CreateImage(r.Context(), &data, idImage)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusCreated, newImage); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// GetImagesByAdvertId handles the request for retrieving a images by advert Id.
func (h *ImageHandler) GetImagesByAdvertId(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("advert_id")
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, "advert_id parameter is required")
		return
	}

	advertId, err := uuid.FromString(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid advert_id parameter")
		return
	}

	images, err := h.uc.GetImagesByAdvertId(r.Context(), advertId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, images); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// DeleteImageById handles the request for deleting a image by its Id.
func (h *ImageHandler) DeleteImageById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	imageId, err := uuid.FromString(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	err = h.uc.DeleteImageById(r.Context(), imageId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err = utils.WriteResponse(w, http.StatusOK, "DELETED image by id: "+id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}
