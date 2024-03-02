package middleware

import (
	"2024_1_TeaStealers/internal/pkg/jwt"
	"context"
	"net/http"
	"time"
)

// CookieName represents the name of the JWT cookie.
const CookieName = "jwt-tean"

// JwtMiddleware is a middleware function that handles JWT authentication.
func JwtMiddleware(next http.Handler) http.Handler {
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

		id, err := jwt.ParseId(claims)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), CookieName, id))

		next.ServeHTTP(w, r)
	})
}
