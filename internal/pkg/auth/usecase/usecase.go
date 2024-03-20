package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/auth"
	"2024_1_TeaStealers/internal/pkg/jwt"
	"2024_1_TeaStealers/internal/pkg/utils"
	"context"
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
		PasswordHash: utils.GenerateHashString(data.Password),
	}

	userResponse, err := u.repo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, "", time.Now(), err
	}

	token, exp, err := jwt.GenerateToken(newUser)
	if err != nil {
		return nil, "", time.Now(), err
	}

	return userResponse, token, exp, nil
}

// Login handles the user login process.
func (u *AuthUsecase) Login(ctx context.Context, data *models.UserLoginData) (*models.User, string, time.Time, error) {
	user, err := u.repo.CheckUser(ctx, data.Login, utils.GenerateHashString(data.Password))
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
	id, _, err := jwt.ParseClaims(claims)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
