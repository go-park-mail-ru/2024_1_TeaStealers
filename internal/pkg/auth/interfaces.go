package auth

import (
	"time"

	_ "github.com/go-park-mail-ru/2024_1_TeaStealers/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int64     `json:"ID"`
	CreatedAt    time.Time `json:"created_at"`
	FirstName    string    `json:"First_Name"`
	LastName     string    `json:"Last_name"`
	Email        string    `json:"email"`
	PasswordHash []byte    `json:"-"`
	JWTSession   []byte    `json:"JWT"`
}

type Usecase interface {
	Register(email, password string) (*User, error)

	Login(email, password string) error

	ResetPassword(email string) error
}

func (r *User) Register(email, password string) (*User, error) {

	var err error
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return &User{}, err
	}

	user := &User{
		ID:           generateID(),
		CreatedAt:    time.Now(),
		Email:        email,
		PasswordHash: passwordHash,
	}

	user.JWTSession, _ = generateJWT(user.ID)
	_, err = r.Create(r)

	return user, err
}

type Repository interface {
	Create(user *User) error
	UpdateInfo(user *User) error
	Delete(user *User) error
	ReadByJWT(session []byte) error
	GetByEmail(email string) (*User, error)
}

func (r *User) GetByEmail (e,ail string) (*User, error) {

}

func (r *User) Create(user *User) error {
	_, err := 
}

func generateID() int64 {
	return 10
}

func generateJWT(id int64) ([]byte, error) {
	var err error
	return []byte("123"), err
}
