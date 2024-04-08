package models

// UserSignUpData represents user information for signup
type UserSignUpData struct {
	// Email stands for users email
	Email string `json:"email"`
	// Phone stands for users phone
	Phone string `json:"phone"`
	// Password stands for users password
	Password string `json:"password"`
}

// UserLoginData represents user information for login
type UserLoginData struct {
	// Login stands for users nickname
	Login string `json:"login"`
	// Password stands for users password
	Password string `json:"password"`
}
