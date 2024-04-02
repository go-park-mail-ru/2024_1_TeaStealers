package middleware

import (
	"2024_1_TeaStealers/internal/pkg/auth"
	"2024_1_TeaStealers/internal/pkg/jwt"
	"context"
	"net/http"
	"time"
)

// CookieName represents the name of the JWT cookie.
const CookieName = "jwt-tean"

type AuthMiddleware struct {
	uc auth.AuthUsecase
}

func NewAuthMiddleware(uc auth.AuthUsecase) *AuthMiddleware {
	return &AuthMiddleware{uc: uc}
}

// JwtMiddleware is a middleware function that handles JWT authentication.
func (md *AuthMiddleware) JwtTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(CookieName)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token := cookie.Value
		claims, err := jwt.ParseToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		timeExp, err := claims.Claims.GetExpirationTime()
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if timeExp.Before(time.Now()) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		id, level, err := jwt.ParseClaims(claims)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if err := md.uc.GetUserLevelById(id, level); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), CookieName, id))

		next.ServeHTTP(w, r)
	})
}
