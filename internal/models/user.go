package models

import (
	"html"
	"time"
)

// User represents user information
type User struct {
	// ID uniquely identifies the user.
	ID int64 `json:"id"`
	// PasswordHash is the hashed password of the user.
	PasswordHash string `json:"-"`
	// LevelUpdate is the level of changes password of the user.
	LevelUpdate int `json:"-"`
	// FirstName is the first name of user.
	FirstName string `json:"firstName"`
	// SecondName is the second name of user.
	SecondName string `json:"secondName"`
	// DateBirthday is the date birthday of user.
	DateBirthday time.Time `json:"dateBirthday"`
	// Phone is the phone of user.
	Phone string `json:"phone"`
	// Email is the email of user.
	Email string `json:"email"`
	// Photo is the filename of photo for user.
	Photo string `json:"photo"`
	// DateCreation is the date creation of user.
	DateCreation time.Time `json:"-"`
	// IsDeleted is deleted user.
	IsDeleted bool `json:"-"`
}

func (user *User) Sanitize() {
	user.PasswordHash = html.EscapeString(user.PasswordHash)
	user.FirstName = html.EscapeString(user.FirstName)
	user.SecondName = html.EscapeString(user.SecondName)
	user.Phone = html.EscapeString(user.Phone)
	user.Email = html.EscapeString(user.Email)
	user.Photo = html.EscapeString(user.Photo)
}

// UserUpdateData represents user update information
type UserUpdateData struct {
	// FirstName is the first name of user.
	FirstName string `json:"firstName"`
	// SecondName is the second name of user.
	SecondName string `json:"secondName"`
	// Phone is the phone of user.
	Phone string `json:"phone"`
	// Email is the email of user.
	Email string `json:"email"`
	// Photo is the filename of photo for user.
}

func (user *UserUpdateData) Sanitize() {
	user.FirstName = html.EscapeString(user.FirstName)
	user.SecondName = html.EscapeString(user.SecondName)
	user.Phone = html.EscapeString(user.Phone)
	user.Email = html.EscapeString(user.Email)
}

type UserUpdatePassword struct {
	// ID uniquely identifies the user.
	ID int64 `json:"id"`
	// OldPassword ...
	OldPassword string `json:"oldPassword"`
	// NewPassword ...
	NewPassword string `json:"newPassword"`
}

func (user *UserUpdatePassword) Sanitize() {
	user.OldPassword = html.EscapeString(user.OldPassword)
	user.NewPassword = html.EscapeString(user.NewPassword)
}
