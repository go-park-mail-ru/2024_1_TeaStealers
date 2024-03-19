package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/users"
	"errors"
	"github.com/satori/uuid"
	"io"
	"os"
	"path/filepath"
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

func (u *UserUsecase) UpdateUserPhoto(file io.Reader, fileType string, id uuid.UUID) (string, error) {
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
	fileName, err := u.repo.UpdateUserPhoto(id, subDirectory+"/"+newFileName)
	if err != nil {
		return "", nil
	}
	return fileName, nil
}

func (u *UserUsecase) DeleteUserPhoto(id uuid.UUID) error {
	if err := u.repo.DeleteUserPhoto(id); err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) UpdateUserInfo(id uuid.UUID, data *models.UserUpdateData) (*models.User, error) {
	if data.Phone == "" {
		return nil, errors.New("phone cannot be empty")
	}
	if data.Email == "" {
		return nil, errors.New("email cannot be empty")
	}
	user, err := u.repo.UpdateUserInfo(id, data)
	if err != nil {
		return nil, errors.New("this email or phone already in use")
	}
	return user, nil
}
