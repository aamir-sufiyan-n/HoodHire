package dto

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type LoginDto struct {
	Email string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required,min=5"`
}
type SignupDto struct{
	Username string `json:"username" validate:"required"`
	Email string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required,min=5"`
}


