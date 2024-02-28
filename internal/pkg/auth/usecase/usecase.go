package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/auth"
)

type UsecaseImpl struct {
	repo auth.Repository
}

func NewUsecase(repo auth.Repository) *UsecaseImpl {
	return &UsecaseImpl{repo: repo}
}

func (u *UsecaseImpl) Register(email, password string) (*models.User, error) {
	user := &models.User{
		Email:        email,
		PasswordHash: []byte(password),
	}
	err := u.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UsecaseImpl) Login(email, password string) error {
	//
}

func (u *UsecaseImpl) ResetPassword(email string) error {
	//
}
