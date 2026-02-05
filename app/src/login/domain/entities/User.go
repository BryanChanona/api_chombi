package entities

import "github.com/go-playground/validator/v10"

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// Modelo INTERNO para login (NO se expone)
type UserWithPassword struct {
	Id           int
	Name         string
	LastName     string
	Email        string
	PasswordHash string
}

// DTO de salida
type UserResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
}

func ValidateUser(user User) error {
	validate := validator.New()
	return validate.Struct(user)
}
