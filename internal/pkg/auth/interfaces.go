package auth

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"time"
)

// AuthUsecase represents the usecase interface for authentication.
type AuthUsecase interface {
	SignUp(context.Context, *models.UserLoginData) (*models.User, string, time.Time, error)
	Login(context.Context, *models.UserLoginData) (*models.User, string, time.Time, error)
}

// AuthRepo represents the repository interface for authentication.
type AuthRepo interface {
	CreateUser(ctx context.Context, newUser *models.User) error
	CheckUser(ctx context.Context, login string, passwordHash string) (*models.User, error)
	GetUserByLogin(cts context.Context, login string) (*models.User, error)
}
