//go:generate mockgen -destination=mock/${GOFILE} -package=${GOPACKAGE}_mock -source=${GOFILE}
package users

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"io"
	"time"

	"github.com/satori/uuid"
)

// UserUsecase represents the usecase interface for users.
type UserUsecase interface {
	GetUser(context.Context, uuid.UUID) (*models.User, error)
	UpdateUserPhoto(context.Context, io.Reader, string, uuid.UUID) (string, error)
	UpdateUserInfo(context.Context, uuid.UUID, *models.UserUpdateData) (*models.User, error)
	DeleteUserPhoto(context.Context, uuid.UUID) error
	UpdateUserPassword(context.Context, *models.UserUpdatePassword) (string, time.Time, error) // тут менять левел юзера + генерировать новый жвт
}

// UserRepo represents the repository interface for users.
type UserRepo interface {
	GetUserById(context.Context, uuid.UUID) (*models.User, error)
	UpdateUserPhoto(context.Context, uuid.UUID, string) (string, error)
	DeleteUserPhoto(context.Context, uuid.UUID) error
	UpdateUserInfo(context.Context, uuid.UUID, *models.UserUpdateData) (*models.User, error)
	UpdateUserPassword(context.Context, uuid.UUID, string) (int, error)
	CheckUserPassword(context.Context, uuid.UUID, string) error
}
