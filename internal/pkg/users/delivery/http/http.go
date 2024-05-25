package http

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/middleware"
	genUsers "2024_1_TeaStealers/internal/pkg/users/delivery/grpc/gen"
	"2024_1_TeaStealers/internal/pkg/utils"
	"net/http"

	"google.golang.org/grpc"
)

// UserClientHandler handles HTTP requests for user.
type UserClientHandler struct {
	client genUsers.UsersClient
}

// NewClientUserHandler creates a new instance of UserHandler.
func NewClientUserHandler(grpcConn *grpc.ClientConn) *UserClientHandler {
	return &UserClientHandler{client: genUsers.NewUsersClient(grpcConn)}
}

func (h *UserClientHandler) GetCurUser(w http.ResponseWriter, r *http.Request) {
	uId, ok := r.Context().Value(middleware.CookieName).(int64)

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "error parse id")
		return
	}

	resp, err := h.client.GetCurUser(r.Context(), &genUsers.GetUserRequest{Id: uId})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "user is not exists")
		return
	}

	userInfo := &models.User{ID: uId, FirstName: resp.FirstName, SecondName: resp.Surname,
		Phone: resp.Phone, Email: resp.Email}
	userInfo.Sanitize()

	if err := utils.WriteResponse(w, http.StatusOK, userInfo); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	}
}

func (h *UserClientHandler) UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	uId, ok := r.Context().Value(middleware.CookieName).(int64)
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

	resp, err := h.client.UpdateUserInfo(r.Context(), &genUsers.UpdateUserInfoRequest{Id: uId,
		FirstName: data.FirstName, Surname: data.SecondName,
		Phone: data.Phone, Email: data.Email})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	if resp.Updated {
		if err := utils.WriteResponse(w, http.StatusOK, data); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, "error write response")
		}
	}
}
