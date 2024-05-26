package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/users"
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"

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
func (u *UserUsecase) GetUser(ctx context.Context, id int64) (*models.User, error) {
	user, err := u.repo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUsecase) UpdateUserPhoto(ctx context.Context, file io.Reader, fileType string, id int64) (string, error) {
	newId := uuid.NewV4()
	newFileName := newId.String() + fileType
	subDirectory := "avatars"
	directory := filepath.Join(os.Getenv("DOCKER_DIR"), subDirectory)
	if err := os.MkdirAll(directory, 0755); err != nil {
		return "", err
	}
	destination, err := os.Create(directory + "/" + newFileName)
	if err != nil {
		return "", err
	}
	defer destination.Close()
	_, err = io.Copy(destination, file)
	if err != nil {
		return "", err
	}
	fileName, err := u.repo.UpdateUserPhoto(ctx, id, subDirectory+"/"+newFileName)
	if err != nil {
		return "", nil
	}
	return fileName, nil
}

func (u *UserUsecase) DeleteUserPhoto(ctx context.Context, id int64) error {
	if err := u.repo.DeleteUserPhoto(ctx, id); err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) UpdateUserInfo(ctx context.Context, id int64, data *models.UserUpdateData) (*models.User, error) {
	if data.Phone == "" {
		return nil, errors.New("phone cannot be empty")
	}
	if data.Email == "" {
		return nil, errors.New("email cannot be empty")
	}
	user, err := u.repo.UpdateUserInfo(ctx, id, data)
	if err != nil {
		return nil, errors.New("this email or phone already in use")
	}
	return user, nil
}
