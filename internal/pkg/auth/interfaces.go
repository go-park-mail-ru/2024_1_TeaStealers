package auth

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"time"

	"github.com/satori/uuid"
)

// AuthUsecase represents the usecase interface for authentication.
type AuthUsecase interface {
	SignUp(context.Context, *models.UserSignUpData) (*models.User, string, time.Time, error)
	Login(context.Context, *models.UserLoginData) (*models.User, string, time.Time, error)
	CheckAuth(context.Context, string) (uuid.UUID, error)
}

// AuthRepo represents the repository interface for authentication.
type AuthRepo interface {
	CreateUser(ctx context.Context, newUser *models.User) (*models.User, error)
	CheckUser(ctx context.Context, login string, passwordHash string) (*models.User, error)
	GetUserByLogin(ctx context.Context, login string) (*models.User, error)
	GetUserLevelById(id uuid.UUID) (int, error)
}
