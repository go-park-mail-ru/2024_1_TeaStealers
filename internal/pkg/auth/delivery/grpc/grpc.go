package grpc

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/auth"
	genAuth "2024_1_TeaStealers/internal/pkg/auth/delivery/grpc/gen"
	"context"
	"errors"
	"github.com/satori/uuid"
	"go.uber.org/zap"
)

const (
	SignUpMethod    = "SignUp"
	LoginMethod     = "Login"
	LogoutMethod    = "Logout"
	CheckAuthMethod = "CheckAuth"
)

// AuthHandler handles HTTP requests for user authentication.
type AuthServerHandler struct {
	genAuth.AuthServer
	// uc represents the usecase interface for authentication.
	uc     auth.AuthUsecase
	logger *zap.Logger
}

// NewServerAuthHandler creates a new instance of AuthHandler.
func NewServerAuthHandler(uc auth.AuthUsecase, logger *zap.Logger) *AuthServerHandler {
	return &AuthServerHandler{uc: uc, logger: logger}
}

// @Summary Register a new user
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.UserLoginData true "User data"
// @Success 201 {object} models.User
// @Failure 400 {string} string "Incorrect data format"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/signup [post]
func (h *AuthServerHandler) SignUp(ctx context.Context, req *genAuth.SignUpRequest) (*genAuth.SignUpInResponse, error) {
	data := models.UserSignUpData{Email: req.Email, Phone: req.Phone, Password: req.Password}
	data.Sanitize()
	newUser, token, exp, err := h.uc.SignUp(ctx, &data)
	if err != nil {
		h.logger.Error(err.Error())
		// utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, SignUpMethod, err, http.StatusBadRequest)
		return nil, errors.New("error signup")
	}
	newUser.Sanitize()

	layout := "2006-01-02 15:04:05"
	dateString := exp.Format(layout)

	h.logger.Info("success logIn")
	// utils.LogSuccesResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, SignUpMethod)
	return &genAuth.SignUpInResponse{Token: token, Exp: dateString}, nil

}

// @Summary User login
// @Description User login
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.UserLoginData true "User login data"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Incorrect password or login"
// @Failure 500 {string} string "Internal server error"
// @Router /auth/login [post]
func (h *AuthServerHandler) Login(ctx context.Context, req *genAuth.SignInRequest) (*genAuth.SignUpInResponse, error) {
	data := models.UserLoginData{Login: req.Email, Password: req.Password}
	data.Sanitize()
	_, token, exp, err := h.uc.Login(ctx, &data)

	if err != nil {
		h.logger.Error(err.Error())
		// utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, LoginMethod, err, http.StatusBadRequest)
		return nil, errors.New("error login")
	}

	layout := "2006-01-02 15:04:05"
	dateString := exp.Format(layout)

	h.logger.Info("success signUp")
	// utils.LogSuccesResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, LoginMethod)
	return &genAuth.SignUpInResponse{Token: token, Exp: dateString}, nil
}

func (h *AuthServerHandler) CheckAuth(ctx context.Context, req *genAuth.CheckAuthRequest) (*genAuth.CheckAuthResponse, error) {

	uuidUser, err := uuid.FromString(req.Id)
	if err != nil {
		h.logger.Error(err.Error())
		// utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, CheckAuthMethod, errors.New("user id is incorrect"), http.StatusUnauthorized)
		return nil, errors.New("incorrect user id")
	}
	err = h.uc.CheckAuth(ctx, uuidUser, int(req.Level))
	if err != nil {
		h.logger.Error(err.Error())
		// utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, CheckAuthMethod, err, http.StatusUnauthorized)
		return nil, errors.New("user not exists")
	}
	h.logger.Info("success checkAuth")
	// utils.LogSuccesResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, CheckAuthMethod)
	return &genAuth.CheckAuthResponse{Authorized: true}, nil
}
