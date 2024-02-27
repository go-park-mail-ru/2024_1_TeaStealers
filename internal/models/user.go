package models

import "time"

type User struct {
	ID           int64     `json:"ID"`
	CreatedAt    time.Time `json:"created_at"`
	FirstName    string    `json:"First_Name"`
	LastName     string    `json:"Last_name"`
	Email        string    `json:"email"`
	PasswordHash []byte    `json:"-"`
}
