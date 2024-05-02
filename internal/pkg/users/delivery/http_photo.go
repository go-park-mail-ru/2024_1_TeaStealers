package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/middleware"
	"2024_1_TeaStealers/internal/pkg/users"
	"2024_1_TeaStealers/internal/pkg/utils"
	"net/http"
	"path/filepath"
	"slices"
	"strings"

	"github.com/satori/uuid"
)

// UserHandlerPhoto handles HTTP requests for user.
type UserHandlerPhoto struct {
	// uc represents the usecase interface for user.
	uc users.UserUsecase
}

// NewUserHandlerPhoto creates a new instance of UserHandler.
func NewUserHandlerPhoto(uc users.UserUsecase) *UserHandlerPhoto {
	return &UserHandlerPhoto{uc: uc}
}

func (h *UserHandlerPhoto) UpdateUserPhoto(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("csrftoken")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "csrf cookie not found")
		return
	}
	id := r.Context().Value(middleware.CookieName)
	UUID, ok := id.(uuid.UUID)
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
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

	fileName, err := h.uc.UpdateUserPhoto(r.Context(), file, fileType, UUID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "failed upload file")
		return
	}
	if err := utils.WriteResponse(w, http.StatusOK, fileName); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	}
}

func (h *UserHandlerPhoto) DeleteUserPhoto(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("csrftoken")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "csrf cookie not found")
		return
	}
	id := r.Context().Value(middleware.CookieName)
	UUID, ok := id.(uuid.UUID)
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}
	if err := h.uc.DeleteUserPhoto(r.Context(), UUID); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "error delete avatar")
		return
	}
	if err := utils.WriteResponse(w, http.StatusOK, "success delete avatar"); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	}
}

func (h *UserHandlerPhoto) UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("csrftoken")
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "csrf cookie not found")
		return
	}
	id, ok := r.Context().Value(middleware.CookieName).(uuid.UUID)
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}
	data := &models.UserUpdateData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}
	data.Sanitize()

	user, err := h.uc.UpdateUserInfo(r.Context(), id, data)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	user.Sanitize()

	if err := utils.WriteResponse(w, http.StatusOK, user); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
	}
}
