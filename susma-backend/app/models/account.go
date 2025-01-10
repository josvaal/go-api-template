package models

type RegisterRequest struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	ProfilePicture string `json:"profile_picture"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Account struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}
