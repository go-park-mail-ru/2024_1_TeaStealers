package middleware

import (
	genAuth "2024_1_TeaStealers/internal/pkg/auth/delivery/grpc/gen"
	"2024_1_TeaStealers/internal/pkg/jwt"
	"2024_1_TeaStealers/internal/pkg/metrics"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/satori/uuid"
	"google.golang.org/grpc"

	"go.uber.org/zap"
)

// CookieName represents the name of the JWT cookie.
const CookieName = "jwt-tean"

type AuthMiddleware struct {
	logger *zap.Logger
	client genAuth.AuthClient
}

func NewAuthMiddleware(grpcConn *grpc.ClientConn, logger *zap.Logger) *AuthMiddleware {
	return &AuthMiddleware{client: genAuth.NewAuthClient(grpcConn), logger: logger}
}

// JwtMiddleware is a middleware function that handles JWT authentication.
func (md *AuthMiddleware) JwtMiddleware(next http.Handler, metrics metrics.MetricsHTTP) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		cookie, err := r.Cookie(CookieName)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := cookie.Value

		claims, err := jwt.ParseToken(token)
		if err != nil {
			log.Println(err)
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

		ctx := context.WithValue(r.Context(), "requestId", uuid.NewV4().String())
		resp, err := md.client.CheckAuth(ctx, &genAuth.CheckAuthRequest{Id: id, Level: int64(level)})
		if err != nil {
			metrics.IncreaseHits("401", "GET", "/api/check_auth (middleware)")
			w.WriteHeader(http.StatusUnauthorized)
			dur := time.Since(start)
			metrics.AddDurationToHandlerTimings("/api/check_auth (middleware)", "GET", dur)
			return
		}
		metrics.IncreaseHits("200", "GET", "/api/check_auth (middleware)")

		if !resp.Authorized {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), CookieName, id))
		dur := time.Since(start)
		metrics.AddDurationToHandlerTimings("/api/check_auth (middleware)", "GET", dur)
		next.ServeHTTP(w, r)
	})
}

// StatMiddleware is a middleware function that handles urls for likes and stat view.
func (md *AuthMiddleware) StatMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(CookieName)

		if err == nil {
			token := cookie.Value
			claims, err := jwt.ParseToken(token)
			if err == nil {
				id, _, _ := jwt.ParseClaims(claims)
				r = r.WithContext(context.WithValue(r.Context(), CookieName, id))
			}
		}
		next.ServeHTTP(w, r)
	})
}
