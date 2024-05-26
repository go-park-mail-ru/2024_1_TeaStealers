//go:generate mockgen -destination=mock/${GOFILE} -package=${GOPACKAGE}_mock -source=${GOFILE}
package users

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"io"
)

// UserUsecase represents the usecase interface for users.
type UserUsecase interface {
	GetUser(context.Context, int64) (*models.User, error)
	UpdateUserPhoto(context.Context, io.Reader, string, int64) (string, error)
	UpdateUserInfo(context.Context, int64, *models.UserUpdateData) (*models.User, error)
	DeleteUserPhoto(context.Context, int64) error
}

// UserRepo represents the repository interface for users.
type UserRepo interface {
	GetUserById(context.Context, int64) (*models.User, error)
	UpdateUserPhoto(context.Context, int64, string) (string, error)
	DeleteUserPhoto(context.Context, int64) error
	UpdateUserInfo(context.Context, int64, *models.UserUpdateData) (*models.User, error)
}
