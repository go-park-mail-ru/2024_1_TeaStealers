package models

// UserLoginData represents user information for login and signup
type UserLoginData struct {
	// Login stands for users nickname
	Login string `json:"login"`
	// Password stands for users password
	Password string `json:"password"`
}
