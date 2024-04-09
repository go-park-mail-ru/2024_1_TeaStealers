package models

import "html"

// UserSignUpData represents user information for signup
type UserSignUpData struct {
	// Email stands for users email
	Email string `json:"email"`
	// Phone stands for users phone
	Phone string `json:"phone"`
	// Password stands for users password
	Password string `json:"password"`
}

func (usSignUp *UserSignUpData) Sanitize() {
	usSignUp.Email = html.EscapeString(usSignUp.Email)
	usSignUp.Phone = html.EscapeString(usSignUp.Phone)
	usSignUp.Password = html.EscapeString(usSignUp.Password)
}

// UserLoginData represents user information for login
type UserLoginData struct {
	// Login stands for users nickname
	Login string `json:"login"`
	// Password stands for users password
	Password string `json:"password"`
}

func (usLogData *UserLoginData) Sanitize() {
	usLogData.Login = html.EscapeString(usLogData.Login)
	usLogData.Password = html.EscapeString(usLogData.Password)
}
