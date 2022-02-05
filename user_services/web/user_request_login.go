package web

import (
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/exception"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

)

type UserRequestLogin struct {
	Email string `json:"email" validate:"required"`
	Password string`json:"password" validate:"required"`
}

func (u *UserRequestLogin) Validate() error {
	validate := validator.New()

	return validate.Struct(u)
}


func (u *UserRequestLogin) CheckPasswordHash(hash string) *exception.ErrorResponse {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(u.Password))
	if err != nil {
		return &exception.ErrorResponse{
			Code:   400,
			Status: "Wrong Credentials",
			Data:   err.Error(),
		}
	}
	return nil
}
