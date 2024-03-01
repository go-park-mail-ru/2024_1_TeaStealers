package middleware

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/satori/uuid"
	"net/http"
	"os"
	"time"
)

const CookieName = "jwt-tean"

func GenerateToken(user *models.User) (string, time.Time, error) {
	exp := time.Now().Add(time.Hour * 24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"login": user.Login,
		"exp":   exp.Unix(),
	})
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", time.Now(), err
	}
	return tokenStr, exp, nil
}

func ParseToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(CookieName)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token := cookie.Value

		claims, err := ParseToken(token)
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

		payload, err := ParsePayload(claims)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "payload", payload))

		next.ServeHTTP(w, r)
	})
}

func ParsePayload(claims *jwt.Token) (*models.JwtPayload, error) {
	payloadMap, ok := claims.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid format (claims)")
	}
	idStr, ok := payloadMap["id"].(string)
	if !ok {
		return nil, errors.New("incorrect id")
	}
	login, ok := payloadMap["login"].(string)
	if !ok {
		return nil, errors.New("incorrect login")
	}
	id, err := uuid.FromString(idStr)
	if err != nil {
		return nil, errors.New("incorrect id")
	}

	return &models.JwtPayload{
		ID:    id,
		Login: login,
	}, nil
}
