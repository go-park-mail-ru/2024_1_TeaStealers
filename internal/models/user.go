package models

import (
	"time"

	"github.com/satori/uuid"
)

// User represents user information
type User struct {
	// ID uniquely identifies the user.
	ID uuid.UUID `json:"id"`
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
