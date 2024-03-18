package users

import (
	"2024_1_TeaStealers/internal/models"
	"github.com/satori/uuid"
)

// UserUsecase represents the usecase interface for users.
type UserUsecase interface {
	GetUser(uuid.UUID) (*models.User, error)
}

// UserRepo represents the repository interface for users.
type UserRepo interface {
	GetUserById(uuid.UUID) (*models.User, error)
}
