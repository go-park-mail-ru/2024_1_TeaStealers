package delivery

import (
	"2024_1_TeaStealers/internal/pkg/middleware"
	"2024_1_TeaStealers/internal/pkg/users"
	"2024_1_TeaStealers/internal/pkg/utils"
	"fmt"
	"github.com/satori/uuid"
	"net/http"
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
