package delivery

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/auth"
	"2024_1_TeaStealers/internal/pkg/middleware"
	"2024_1_TeaStealers/internal/pkg/utils"
	"errors"
	"net/http"
	"time"
)

// AuthHandler handles HTTP requests for user authentication.
type AuthHandler struct {
	// uc represents the usecase interface for authentication.
	uc auth.AuthUsecase
}

// NewAuthHandler creates a new instance of AuthHandler.
func NewAuthHandler(uc auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{uc: uc}
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
	data := models.UserSignUpData{}

	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect data format")
		return
	}

	newUser, token, exp, err := h.uc.SignUp(r.Context(), &data)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "data already is used")
		return
	}

	http.SetCookie(w, tokenCookie(middleware.CookieName, token, exp))

	if err = utils.WriteResponse(w, http.StatusCreated, newUser); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
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
	data := models.UserLoginData{}
	if err := utils.ReadRequestData(r, &data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, token, exp, err := h.uc.Login(r.Context(), &data)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "incorrect password or login")
	}

	http.SetCookie(w, tokenCookie(middleware.CookieName, token, exp))

	if err := utils.WriteResponse(w, http.StatusOK, user); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// @Summary User logout
// @Description User logout
// @Tags auth
// @Success 200 {string} string "Logged out"
// @Router /auth/logout [get]
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  middleware.CookieName,
		Value: "",
		Path:  "/",
	})
	if err := utils.WriteResponse(w, http.StatusOK, "success logout"); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *AuthHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	tokenCookie, err := r.Cookie(middleware.CookieName)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			utils.WriteError(w, http.StatusUnauthorized, "token cookie not found")
			return
		}
		utils.WriteError(w, http.StatusUnauthorized, "fail to get token cookie")
		return
	}

	id, err := h.uc.CheckAuth(r.Context(), tokenCookie.Value)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "jws token is invalid")
		return
	}
	if err = utils.WriteResponse(w, http.StatusOK, id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
	}
}

// tokenCookie creates a new cookie for storing the authentication token.
func tokenCookie(name, token string, exp time.Time) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    token,
		Expires:  exp,
		Path:     "/",
		HttpOnly: true,
	}
}
