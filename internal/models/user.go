package models

import (
	"github.com/satori/uuid"
)

// User represents user information
type User struct {
	// ID uniquely identifies the user.
	ID uuid.UUID `json:"id"`
	// Login is the username of the user.
	Login string `json:"login"`
	// PasswordHash is the hashed password of the user.
	PasswordHash string `json:"-"`
}
