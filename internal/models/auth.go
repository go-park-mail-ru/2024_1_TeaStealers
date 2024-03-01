package models

import "github.com/satori/uuid"

type JwtPayload struct {
	ID    uuid.UUID
	Login string
}

type UserLoginData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
