package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/jwt"
	"2024_1_TeaStealers/internal/pkg/middleware"
	"2024_1_TeaStealers/internal/pkg/users"
	"2024_1_TeaStealers/internal/pkg/utils"
	"net/http"
	"path/filepath"
	"slices"
	"strings"

	"github.com/satori/uuid"
)

// UserHandler handles HTTP requests for user.
type UserHandler struct {
	// uc represents the usecase interface for user.
	uc users.UserUsecase
}

// NewUserHandler creates a new instance of UserHandler.
func NewUserHandler(uc users.UserUsecase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) GetCurUser(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.CookieName)
	UUID, ok := id.(uuid.UUID)
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}
	userInfo, err := h.uc.GetUser(UUID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "user is not exists")
		return
	}
	if err := utils.WriteResponse(w, http.StatusOK, userInfo); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	}
}

func (h *UserHandler) UpdateUserPhoto(w http.ResponseWriter, r *http.Request) {
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

	fileName, err := h.uc.UpdateUserPhoto(file, fileType, UUID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "failed upload file")
		return
	}
	if err := utils.WriteResponse(w, http.StatusOK, fileName); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	}
}

func (h *UserHandler) DeleteUserPhoto(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.CookieName)
	UUID, ok := id.(uuid.UUID)
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}
	if err := h.uc.DeleteUserPhoto(UUID); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "error delete avatar")
		return
	}
	if err := utils.WriteResponse(w, http.StatusOK, "success delete avatar"); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	}
}

func (h *UserHandler) UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
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

	user, err := h.uc.UpdateUserInfo(id, data)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := utils.WriteResponse(w, http.StatusOK, user); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
	}
}

func (h *UserHandler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.CookieName)
	UUID, ok := id.(uuid.UUID)
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}
	data := &models.UserUpdatePassword{
		ID: UUID,
	}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}

	token, exp, err := h.uc.UpdateUserPassword(data)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	http.SetCookie(w, jwt.TokenCookie(middleware.CookieName, token, exp))

	if err := utils.WriteResponse(w, http.StatusOK, "success update password"); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
	}
}
