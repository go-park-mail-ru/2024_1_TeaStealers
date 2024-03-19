package users

import (
	"2024_1_TeaStealers/internal/models"
	"github.com/satori/uuid"
	"io"
)

// UserUsecase represents the usecase interface for users.
type UserUsecase interface {
	GetUser(uuid.UUID) (*models.User, error)
	UpdateUserPhoto(io.Reader, string, uuid.UUID) (string, error)
	UpdateUserInfo(uuid.UUID, *models.UserUpdateData) (*models.User, error)
	DeleteUserPhoto(uuid.UUID) error
	//UpdateUserPasswrd(password *models.UserUpdatePassword) error //тут менять левел юзера + генерировать новый жвт
}

// UserRepo represents the repository interface for users.
type UserRepo interface {
	GetUserById(uuid.UUID) (*models.User, error)
	UpdateUserPhoto(uuid.UUID, string) (string, error)
	DeleteUserPhoto(uuid.UUID) error
	UpdateUserInfo(uuid.UUID, *models.UserUpdateData) (*models.User, error)
}
