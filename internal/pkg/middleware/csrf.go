package middleware

import (
	"context"
	"github.com/satori/uuid"
	"net/http"
)

// CookieNameCsrf represents the name of the csrf cookie.
const CookieNameCsrf = "csrf-tean"

type CsrfMiddleware struct{}

func NewCsrfMiddleware() *CsrfMiddleware {
	return &CsrfMiddleware{}
}

// SetCSRFToken is a middleware function that handles csrf token cookie.
func (h *CsrfMiddleware) SetCSRFToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := uuid.NewV4().String()
		http.SetCookie(w, &http.Cookie{
			Name:     CookieNameCsrf,
			Value:    token,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		})
		ctx := context.WithValue(r.Context(), "csrftoken", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
