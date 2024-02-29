package auth

import (
	"2024_1_TeaStealers/internal/models"
	"context"
	"time"
)

type AuthUsecase interface {
	SignUp(context.Context, *models.UserLoginData) (*models.User, string, time.Time, error)
	Login(context.Context, *models.UserLoginData) (*models.User, string, time.Time, error)
}

type AuthRepo interface {
	CreateUser(ctx context.Context, newUser *models.User) error
	CheckUser(ctx context.Context, login string, passwordHash string) (*models.User, error)
	GetUserByLogin(cts context.Context, login string) (*models.User, error)
}

// func (r *User) Register(email, password string) (*User, error) {

// 	var err error
// 	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)

// 	if err != nil {
// 		return &User{}, err
// 	}

// 	user := &User{
// 		ID:           generateID(),
// 		CreatedAt:    time.Now(),
// 		Email:        email,
// 		PasswordHash: passwordHash,
// 	}

// 	user.JWTSession, _ = generateJWT(user.ID)
// 	_, err = r.Create(r)

// 	return user, err
// }

// func (r *User) GetByEmail (e,ail string) (*User, error) {

// }

// func (r *User) Create(user *User) error {
// 	_, err :=
// }

// func generateID() int64 {
// 	return 10
// }

// func generateJWT(id int64) ([]byte, error) {
// 	var err error
// 	return []byte("123"), err
// }
