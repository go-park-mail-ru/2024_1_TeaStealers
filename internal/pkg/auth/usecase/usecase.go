package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/auth"
	"2024_1_TeaStealers/internal/pkg/jwt"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"time"

	"github.com/satori/uuid"
)

// AuthUsecase represents the usecase for authentication.
type AuthUsecase struct {
	repo auth.AuthRepo
}

// NewAuthUsecase creates a new instance of AuthUsecase.
func NewAuthUsecase(repo auth.AuthRepo) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

// SignUp handles the user registration process.
func (u *AuthUsecase) SignUp(ctx context.Context, data *models.UserSignUpData) (*models.User, string, time.Time, error) {
	newUser := &models.User{
		ID:           uuid.NewV4(),
		Email:        data.Email,
		Phone:        data.Phone,
		PasswordHash: generateHashString(data.Password),
	}

	if err := u.repo.CreateUser(ctx, newUser); err != nil {
		return nil, "", time.Now(), err
	}

	token, exp, err := jwt.GenerateToken(newUser)
	if err != nil {
		return nil, "", time.Now(), err
	}

	return newUser, token, exp, nil
}

// Login handles the user login process.
func (u *AuthUsecase) Login(ctx context.Context, data *models.UserLoginData) (*models.User, string, time.Time, error) {
	user, err := u.repo.CheckUser(ctx, data.Login, generateHashString(data.Password))
	if err != nil {
		return nil, "", time.Now(), err
	}

	token, exp, err := jwt.GenerateToken(user)
	if err != nil {
		return nil, "", time.Now(), err
	}

	return user, token, exp, nil
}

// CheckAuth checking autorizing
func (u *AuthUsecase) CheckAuth(ctx context.Context, token string) (uuid.UUID, error) {
	claims, err := jwt.ParseToken(token)
	if err != nil {
		return uuid.Nil, err
	}
	id, err := jwt.ParseId(claims)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

// generateHashString returns a hash string for the given input string.
func generateHashString(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
