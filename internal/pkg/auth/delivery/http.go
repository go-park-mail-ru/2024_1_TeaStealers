package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/auth"
	"2024_1_TeaStealers/internal/pkg/jwt"
	"2024_1_TeaStealers/internal/pkg/middleware"
	"2024_1_TeaStealers/internal/pkg/utils"
	"context"
	"errors"
	"net/http"

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
type AuthHandler struct {
	// uc represents the usecase interface for authentication.
	uc     auth.AuthUsecase
	logger *zap.Logger
}

// NewAuthHandler creates a new instance of AuthHandler.
func NewAuthHandler(uc auth.AuthUsecase, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{uc: uc, logger: logger}
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
func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "requestId", uuid.NewV4().String())

	data := models.UserSignUpData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, SignUpMethod, errors.New("error parse data"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}

	newUser, token, exp, err := h.uc.SignUp(ctx, &data)
	if err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, SignUpMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "data already is used")
		return
	}
	newUser.Sanitize()

	http.SetCookie(w, jwt.TokenCookie(middleware.CookieName, token, exp))

	if err = utils.WriteResponse(w, http.StatusCreated, newUser); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, SignUpMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, SignUpMethod)
	}
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
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "requestId", uuid.NewV4().String())

	data := models.UserLoginData{}
	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, LoginMethod, errors.New("error parse data"), http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	data.Sanitize()
	user, token, exp, err := h.uc.Login(ctx, &data)

	if err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, LoginMethod, err, http.StatusBadRequest)
		utils.WriteError(w, http.StatusBadRequest, "incorrect password or login")
		return
	}
	user.Sanitize()

	http.SetCookie(w, jwt.TokenCookie(middleware.CookieName, token, exp))

	if err := utils.WriteResponse(w, http.StatusOK, user); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, LoginMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, LoginMethod)
	}
}

// @Summary User logout
// @Description User logout
// @Tags auth
// @Success 200 {string} string "Logged out"
// @Router /auth/logout [get]
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "requestId", uuid.NewV4().String())

	http.SetCookie(w, &http.Cookie{
		Name:  middleware.CookieName,
		Value: "",
		Path:  "/",
	})
	if err := utils.WriteResponse(w, http.StatusOK, "success logout"); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, LogoutMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, LogoutMethod)
	}
}

func (h *AuthHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "requestId", uuid.NewV4().String())

	idUser := ctx.Value(middleware.CookieName)
	if idUser == nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, CheckAuthMethod, errors.New("user id is nill"), http.StatusUnauthorized)
		utils.WriteError(w, http.StatusUnauthorized, "token not found")
		return
	}
	uuidUser, ok := idUser.(uuid.UUID)
	if !ok {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, CheckAuthMethod, errors.New("user id is incorrect"), http.StatusUnauthorized)
		utils.WriteError(w, http.StatusUnauthorized, "incorrect user id")
		return
	}
	err := h.uc.CheckAuth(ctx, uuidUser)
	if err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, CheckAuthMethod, err, http.StatusUnauthorized)
		utils.WriteError(w, http.StatusUnauthorized, "user not exists")
		return
	}
	if err = utils.WriteResponse(w, http.StatusOK, uuidUser); err != nil {
		utils.LogErrorResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, CheckAuthMethod, err, http.StatusInternalServerError)
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	} else {
		utils.LogSuccesResponse(h.logger, ctx.Value("requestId").(string), utils.DeliveryLayer, CheckAuthMethod)
	}
}
