package delivery

import (
	"2024_1_TeaStealers/internal/pkg/images"
	"2024_1_TeaStealers/internal/pkg/utils"
	"fmt"
	"net/http"
	"path/filepath"
	"slices"
	"strings"

	"github.com/gorilla/mux"
	"github.com/satori/uuid"
)

// ImagesHandler handles HTTP requests for images.
type ImagesHandler struct {
	// uc represents the usecase interface for images.
	uc images.ImageUsecase
}

// NewImageHandler creates a new instance of ImagesHandler.
func NewImageHandler(uc images.ImageUsecase) *ImagesHandler {
	return &ImagesHandler{uc: uc}
}

// UploadImage upload image for advert
// @Summary Upload new images for advert
// @Description Upload new images for advert
// @Tags images
// @Accept json
// @Produce json
// @Param id formData string true "id advert"
// @Param file formData file true "image (.jpg, .jpeg, .png)"
// @Success 201 {object} models.Image
// @Failure 400 {string} string "Incorrect data format"
// @Failure 400 {string} string "max size 5 mb"
// @Failure 500 {string} string "failed to upload images"
// @Router /adverts/image [post]
func (h *ImagesHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(5 << 20); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "max size 5 mb")
		return
	}
	advertIdStr := r.FormValue("id")
	advertUUID, err := uuid.FromString(advertIdStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect advert id")
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
	fmt.Println(fileType)
	if !slices.Contains(allowedExtensions, fileType) {
		utils.WriteError(w, http.StatusBadRequest, "jpg, jpeg, png only")
	}

	image, err := h.uc.UploadImage(file, fileType, advertUUID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "failed to upload image")
		return
	}

	if err = utils.WriteResponse(w, http.StatusCreated, image); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// GetAdvertImages get list of image for advert
// @Summary get list of image for advert
// @Description get list of image for advert
// @Tags images
// @Accept json
// @Produce json
// @Param id path string true "id advert"
// @Success 200 {array} []models.ImageResp
// @Failure 400 {string} string "incorrect advert id"
// @Failure 500 {string} string "Internal Server Error"
// @Router /advert/{id}/image [get]
func (h *ImagesHandler) GetAdvertImages(w http.ResponseWriter, r *http.Request) {
	queryParam := mux.Vars(r)["id"]
	advrtUUID, err := uuid.FromString(queryParam)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect advert id")
		return
	}
	resp, err := h.uc.GetAdvertImages(advrtUUID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data")
		return
	}
	if err = utils.WriteResponse(w, http.StatusOK, resp); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// DeleteImage delete image.
// @Summary get delete image.
// @Description delete image.
// @Tags images
// @Accept json
// @Produce json
// @Param id path string true "id advert"
// @Success 200 {array} []models.ImageResp
// @Failure 400 {string} string "incorrect image id"
// @Failure 500 {string} string "Internal Server Error"
// @Router /advert/{id}/image [delete]
func (h *ImagesHandler) DeleteImage(w http.ResponseWriter, r *http.Request) {
	queryParam := mux.Vars(r)["id"]
	imageUUID, err := uuid.FromString(queryParam)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect image id")
		return
	}
	resp, err := h.uc.DeleteImage(imageUUID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data")
		return
	}
	if err = utils.WriteResponse(w, http.StatusOK, resp); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}
