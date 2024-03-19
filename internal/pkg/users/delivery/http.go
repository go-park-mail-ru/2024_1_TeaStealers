package delivery

import (
	"2024_1_TeaStealers/internal/pkg/middleware"
	"2024_1_TeaStealers/internal/pkg/users"
	"2024_1_TeaStealers/internal/pkg/utils"
	"fmt"
	"github.com/satori/uuid"
	"net/http"
	"path/filepath"
	"slices"
	"strings"
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
	}
	fmt.Println(UUID)
	userInfo, err := h.uc.GetUser(UUID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "user is not exists")
	}
	if err := utils.WriteResponse(w, http.StatusOK, userInfo); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
	}
}

func (h *UserHandler) UpdateUserPhoto(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.CookieName)
	UUID, ok := id.(uuid.UUID)
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
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
	fmt.Println(fileType)
	if !slices.Contains(allowedExtensions, fileType) {
		utils.WriteError(w, http.StatusBadRequest, "jpg, jpeg, png only")
	}

	fileName, err := h.uc.UpdateUserPhoto(file, fileType, UUID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "failed upload file")
	}
	if err := utils.WriteResponse(w, http.StatusOK, fileName); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
	}
}
