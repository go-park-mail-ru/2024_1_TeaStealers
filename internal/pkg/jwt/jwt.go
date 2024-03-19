package jwt

import (
	"2024_1_TeaStealers/internal/models"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/satori/uuid"
)

// GenerateToken returns a new JWT token for the given user.
func GenerateToken(user *models.User) (string, time.Time, error) {
	exp := time.Now().Add(time.Hour * 24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"level": user.LevelUpdate,
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

// ParseClaims parses the user ID from the JWT token claims.
func ParseClaims(claims *jwt.Token) (uuid.UUID, int, error) {
	payloadMap, ok := claims.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, 0, errors.New("invalid claims")
	}
	idStr, ok := payloadMap["id"].(string)
	if !ok {
		return uuid.Nil, 0, errors.New("incorrect id")
	}
	id, err := uuid.FromString(idStr)
	if err != nil {
		return uuid.Nil, 0, errors.New("incorrect id")
	}
	levelStr, ok := payloadMap["level"].(float64)
	if !ok {
		return uuid.Nil, 0, errors.New("incorrect level")
	}

	return id, int(levelStr), nil
}
