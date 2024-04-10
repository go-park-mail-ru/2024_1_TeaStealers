package usecase

import (
	"2024_1_TeaStealers/internal/models"
	"2024_1_TeaStealers/internal/pkg/jwt"
	"2024_1_TeaStealers/internal/pkg/users"
	"2024_1_TeaStealers/internal/pkg/utils"
	"errors"
	"io"
	"os"
	"path/filepath"
	"time"

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

func (u *UserUsecase) UpdateUserPassword(data *models.UserUpdatePassword) (string, time.Time, error) {
	oldPasswordHash := utils.GenerateHashString(data.OldPassword)
	newPasswordHash := utils.GenerateHashString(data.NewPassword)
	if oldPasswordHash == newPasswordHash {
		return "", time.Now(), errors.New("passwords must not match")
	}
	if err := u.repo.CheckUserPassword(data.ID, oldPasswordHash); err != nil {
		return "", time.Now(), errors.New("invalid old password")
	}
	level, err := u.repo.UpdateUserPassword(data.ID, newPasswordHash)
	if err != nil {
		return "", time.Now(), errors.New("incorrect id or passwordhash")
	}
	user := &models.User{
		ID:          data.ID,
		LevelUpdate: level,
	}
	token, exp, err := jwt.GenerateToken(user)
	if err != nil {
		return "", time.Now(), errors.New("unable to generate token")
	}
	return token, exp, nil
}
