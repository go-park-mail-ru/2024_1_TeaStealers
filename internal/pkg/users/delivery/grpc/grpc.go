package grpc

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/users"
	genUsers "2024_1_TeaStealers/internal/pkg/users/delivery/grpc/gen"
	"context"
	"errors"

	"github.com/satori/uuid"
)

// UserServerHandler handles HTTP requests for user.
type UserServerHandler struct {
	// uc represents the usecase interface for user.
	uc users.UserUsecase
	genUsers.UsersServer
}

// NewUserServerHandler creates a new instance of UserHandler.
func NewUserServerHandler(uc users.UserUsecase) *UserServerHandler {
	return &UserServerHandler{uc: uc}
}

func (h *UserServerHandler) GetCurUser(ctx context.Context, req *genUsers.GetUserRequest) (*genUsers.GetUserResponse, error) {
	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())

	userId := req.Id

	userInfo, err := h.uc.GetUser(ctx, userId)
	if err != nil {
		return nil, errors.New("user is not exists")
	}
	userInfo.Sanitize()

	return &genUsers.GetUserResponse{FirstName: userInfo.FirstName, Surname: userInfo.SecondName,
		DateBirthday: userInfo.DateBirthday.String(), Phone: userInfo.Phone, Email: userInfo.Email,
		Photo: userInfo.Photo}, nil
}

func (h *UserServerHandler) UpdateUserInfo(ctx context.Context, req *genUsers.UpdateUserInfoRequest) (*genUsers.UpdateUserInfoResponse, error) {
	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())

	userId := req.Id

	data := &models.UserUpdateData{FirstName: req.FirstName, SecondName: req.Surname, Phone: req.Phone, Email: req.Email}
	data.Sanitize()

	user, err := h.uc.UpdateUserInfo(ctx, userId, data)
	if err != nil {
		return nil, err
	}
	user.Sanitize()

	return &genUsers.UpdateUserInfoResponse{Updated: true}, nil
}

func (h *UserServerHandler) UpdateUserPassword(ctx context.Context, req *genUsers.UpdatePasswordRequest) (*genUsers.UpdatePasswordResponse, error) {
	ctx = context.WithValue(ctx, "requestId", uuid.NewV4().String())

	userId := req.Id

	data := &models.UserUpdatePassword{
		ID:          userId,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	}
	data.Sanitize()

	token, exp, err := h.uc.UpdateUserPassword(ctx, data)
	if err != nil {
		return nil, err
	}
	return &genUsers.UpdatePasswordResponse{Updated: true, Token: token, Exp: exp.String()}, nil
}
