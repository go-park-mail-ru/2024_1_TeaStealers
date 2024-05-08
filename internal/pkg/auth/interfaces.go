//go:generate mockgen -destination=mock/${GOFILE} -package=${GOPACKAGE}_mock -source=${GOFILE}
package auth

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"time"
)

const (
	SignUpMethod           = "SignUp"
	LoginMethod            = "Login"
	CheckAuthMethod        = "CheckAuth"
	GetUserLevelByIdMethod = "GetUserLevelById"
	CreateUserMethod       = "CreateUser"
	CheckUserMethod        = "CheckUser"
	GetUserByLoginMethod   = "GetUserByLogin"
	BeginTxMethod          = "BeginTx"
)

// AuthUsecase represents the usecase interface for authentication.
type AuthUsecase interface {
	SignUp(context.Context, *models.UserSignUpData) (*models.User, string, time.Time, error)
	Login(context.Context, *models.UserLoginData) (*models.User, string, time.Time, error)
	CheckAuth(ctx context.Context, id int64, jwtLevel int) error
}

// AuthRepo represents the repository interface for authentication.
type AuthRepo interface {
	CreateUser(ctx context.Context, newUser *models.User) (*models.User, error)
	CheckUser(ctx context.Context, login string, passwordHash string) (*models.User, error)
	GetUserByLogin(ctx context.Context, login string) (*models.User, error)
	GetUserLevelById(ctx context.Context, id int64) (int, error)
}
