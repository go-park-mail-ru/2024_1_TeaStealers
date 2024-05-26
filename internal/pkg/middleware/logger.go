package middleware

import (
	"context"
	"net/http"

	"github.com/satori/uuid"
)

func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "requestId", uuid.NewV4().String())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
