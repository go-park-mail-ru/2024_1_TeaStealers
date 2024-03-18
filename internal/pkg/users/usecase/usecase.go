package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/users"
	"github.com/satori/uuid"
)

// UserUsecase represents the usecase for user.
type UserUsecase struct {
	repo users.UserRepo
}

// NewUserUsecase creates a new instance of UserUsecase.
func NewUserUsecase(repo users.UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

// GetUser ...
func (u *UserUsecase) GetUser(id uuid.UUID) (*models.User, error) {
	user, err := u.repo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
