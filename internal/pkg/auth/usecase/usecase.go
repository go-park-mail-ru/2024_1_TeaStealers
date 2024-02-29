package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/auth"
	"2024_1_TeaStealers/internal/pkg/middleware"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"github.com/satori/uuid"
	"time"
)

type AuthUsecase struct {
	repo auth.AuthRepo
}

func NewAuthUsecase(repo auth.AuthRepo) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

func (u *AuthUsecase) SignUp(ctx context.Context, data *models.UserLoginData) (*models.User, string, time.Time, error) {
	newUser := &models.User{
		ID:           uuid.NewV4(),
		Login:        data.Login,
		Phone:        "",
		PasswordHash: generateHashString(data.Password),
	}

	if err := u.repo.CreateUser(ctx, newUser); err != nil {
		return nil, "", time.Now(), err
	}

	token, exp, err := middleware.GenerateToken(newUser)
	if err != nil {
		return nil, "", time.Now(), err
	}

	return newUser, token, exp, nil
}

func (u *AuthUsecase) Login(ctx context.Context, data *models.UserLoginData) (*models.User, string, time.Time, error) {
	user, err := u.repo.CheckUser(ctx, data.Login, generateHashString(data.Password))
	if err != nil {
		return nil, "", time.Now(), err
	}

	token, exp, err := middleware.GenerateToken(user)
	if err != nil {
		return nil, "", time.Now(), err
	}

	return user, token, exp, nil
}

func generateHashString(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
