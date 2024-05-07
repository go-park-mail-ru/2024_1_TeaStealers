package http

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/jwt"
	"2024_1_TeaStealers/internal/pkg/middleware"
	genUsers "2024_1_TeaStealers/internal/pkg/users/delivery/grpc/gen"
	"2024_1_TeaStealers/internal/pkg/utils"
	"github.com/satori/uuid"
	"google.golang.org/grpc"
	"net/http"
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
	id := r.Context().Value(middleware.CookieName).(string)
	uId, err := uuid.FromString(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "error parse id")
		return
	}

	resp, err := h.client.GetCurUser(r.Context(), &genUsers.GetUserRequest{Id: id})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "user is not exists")
		return
	}
	dateBirth, err := utils.StringToTime(resp.DateBirthday, resp.DateBirthday)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "error parse time")
		return
	}

	userInfo := &models.User{ID: uId, FirstName: resp.FirstName, SecondName: resp.Surname,
		DateBirthday: dateBirth, Phone: resp.Phone, Email: resp.Email}
	userInfo.Sanitize()

	if err := utils.WriteResponse(w, http.StatusOK, userInfo); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
		return
	}
}

func (h *UserClientHandler) UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
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

	resp, err := h.client.UpdateUserInfo(r.Context(), &genUsers.UpdateUserInfoRequest{Id: id.String(),
		FirstName: data.FirstName, Surname: data.SecondName, DateBirthday: data.DateBirthday.String(),
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

func (h *UserClientHandler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
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
	data.Sanitize()

	resp, err := h.client.UpdateUserPassword(r.Context(), &genUsers.UpdatePasswordRequest{Id: UUID.String(),
		OldPassword: data.OldPassword, NewPassword: data.NewPassword})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	exp, err := utils.StringToTime(resp.Exp, resp.Exp)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "error parse time")
		return
	}

	http.SetCookie(w, jwt.TokenCookie(middleware.CookieName, resp.Token, exp))
	if err = utils.WriteResponse(w, http.StatusOK, "success update password"); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
	}
}
