package jwt

import (
	"2024_1_TeaStealers/internal/models"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/satori/uuid"
	"os"
	"time"
)

// GenerateToken returns a new JWT token for the given user.
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

// ParseToken parses the provided JWT token string and returns the parsed token.
func ParseToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

// ParseId parses the user ID from the JWT token claims.
func ParseId(claims *jwt.Token) (uuid.UUID, error) {
	payloadMap, ok := claims.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, errors.New("invalid claims")
	}
	idStr, ok := payloadMap["id"].(string)
	if !ok {
		return uuid.Nil, errors.New("incorrect id")
	}
	id, err := uuid.FromString(idStr)
	if err != nil {
		return uuid.Nil, errors.New("incorrect id")
	}

	return id, nil
}
