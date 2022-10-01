package models

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"Email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
