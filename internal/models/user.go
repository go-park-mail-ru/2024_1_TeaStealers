package models

import (
	"github.com/satori/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Login        string    `json:"login"`
	Phone        string    `json:"phone"`
	PasswordHash string    `json:"-"`
}
