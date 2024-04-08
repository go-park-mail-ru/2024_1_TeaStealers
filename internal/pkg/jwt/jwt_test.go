package jwt

import (
	"2024_1_TeaStealers/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTokenValid(t *testing.T) {
	user := &models.User{
		ID: uuid.NewV4(),
	}
	token, _, _ := GenerateToken(user)
	parsedToken, err := ParseToken(token)

	assert.NoError(t, err)
	assert.NotNil(t, parsedToken)
}

func TestParseTokenInvalid(t *testing.T) {
	token := "invalid_jwt_token"
	parsedToken, err := ParseToken(token)

	assert.Error(t, err)
	assert.Nil(t, parsedToken)
}

/*
	func TestParseClaimsValid(t *testing.T) {
		claims := &jwt.Token{
			Claims: jwt.MapClaims{
				"id":    uuid.NewV4(),
				"level": 1.0,
			},
		}
		id, level, err := ParseClaims(claims)

		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, id)
		assert.Equal(t, 1, level)
	}
*/
func TestParseClaimsInvalid(t *testing.T) {
	claims := &jwt.Token{}

	id, level, err := ParseClaims(claims)

	assert.Error(t, err)
	assert.Equal(t, uuid.Nil, id)
	assert.Equal(t, 0, level)
}
