package http

import (
	"2024_1_TeaStealers/internal/models"
	genAuth "2024_1_TeaStealers/internal/pkg/auth/delivery/grpc/gen"
	"2024_1_TeaStealers/internal/pkg/jwt"
	"2024_1_TeaStealers/internal/pkg/middleware"
	"2024_1_TeaStealers/internal/pkg/utils"
	"context"
	"errors"
	"net/http"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	SignUpMethod    = "SignUp"
	LoginMethod     = "Login"
	LogoutMethod    = "Logout"
	CheckAuthMethod = "CheckAuth"
)

// AuthClientHandler handles HTTP requests for user authentication.
type AuthClientHandler struct {
	client genAuth.AuthClient
	logger *zap.Logger
}

// NewClientAuthHandler creates a new instance of AuthHandler.
func NewClientAuthHandler(grpcConn *grpc.ClientConn, logger *zap.Logger) *AuthClientHandler {
	return &AuthClientHandler{client: genAuth.NewAuthClient(grpcConn), logger: logger}
}

// SignUp
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
func (h *AuthClientHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	data := models.UserSignUpData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, SignUpMethod, errors.New("error parse data"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}
	data.Sanitize()

	TokenExp, err := h.client.SignUp(r.Context(), &genAuth.SignUpRequest{Email: data.Email, Phone: data.Phone, Password: data.Password})
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, SignUpMethod, err, http.StatusConflict)
		utils.WriteError(w, http.StatusConflict, err.Error())
		return
	}

	token, exp := TokenExp.Token, TokenExp.Exp
	expTime, err := utils.StringToTime("2006-01-02 15:04:05", exp)
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, SignUpMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, jwt.TokenCookie(middleware.CookieName, token, expTime))

	if err = utils.WriteResponse(w, http.StatusCreated, "newUser"); err != nil { // todo а здесь нужно возвращать newUser???
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, SignUpMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, SignUpMethod)
		return
	}
}

// Login
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
func (h *AuthClientHandler) Login(w http.ResponseWriter, r *http.Request) {
	data := models.UserLoginData{}
	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, LoginMethod, errors.New("error parse data"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	data.Sanitize()

	md := metadata.New(map[string]string{"requestId": r.Context().Value("requestId").(string)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	loginResp, err := h.client.Login(ctx, &genAuth.SignInRequest{Email: data.Login, Password: data.Password})
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, SignUpMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusBadRequest, "incorrect password or login")
		return
	}

	token, exp := loginResp.Token, loginResp.Exp
	expTime, err := utils.StringToTime("2006-01-02 15:04:05", exp)
	if err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, LoginMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, jwt.TokenCookie(middleware.CookieName, token, expTime))

	if err = utils.WriteResponse(w, http.StatusOK, "user"); err != nil { // todo нужен user?
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, LoginMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, LoginMethod)
	}
}

// Logout
// @Summary User logout
// @Description User logout
// @Tags auth
// @Success 200 {string} string "Logged out"
// @Router /auth/logout [get]
func (h *AuthClientHandler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  middleware.CookieName,
		Value: "",
		Path:  "/",
	})
	if err := utils.WriteResponse(w, http.StatusOK, "success logout"); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, LogoutMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, LogoutMethod)
		return
	}
}

func (h *AuthClientHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	idUser := r.Context().Value(middleware.CookieName)

	if idUser == nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CheckAuthMethod, errors.New("user id is nil"), http.StatusUnauthorized)
		utils.WriteError(w, http.StatusUnauthorized, "token not found")
		return
	}
	uuidUser, ok := idUser.(int64)
	if !ok {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CheckAuthMethod, errors.New("user id is not a string"), http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, "user id is not a string")
		return
	}

	if err := utils.WriteResponse(w, http.StatusOK, uuidUser); err != nil {
		utils.LogErrorResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CheckAuthMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		utils.LogSuccesResponse(h.logger, r.Context().Value("requestId").(string), utils.DeliveryLayer, CheckAuthMethod)
		return
	}

}

func (h *AuthClientHandler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	uId, ok := r.Context().Value(middleware.CookieName).(int64)

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, "incorrect id")
		return
	}

	data := &models.UserUpdatePassword{
		ID: uId,
	}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}
	data.Sanitize()

	resp, err := h.client.UpdateUserPassword(r.Context(), &genAuth.UpdatePasswordRequest{Id: uId,
		OldPassword: data.OldPassword, NewPassword: data.NewPassword})
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	exp := resp.Exp[:19]
	expTime, err := utils.StringToTime("2006-01-02 15:04:05", exp)
	if err != nil {
		// utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, LoginMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.SetCookie(w, jwt.TokenCookie(middleware.CookieName, resp.Token, expTime))
	if err = utils.WriteResponse(w, http.StatusOK, "success update password"); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "error write response")
	}
}
