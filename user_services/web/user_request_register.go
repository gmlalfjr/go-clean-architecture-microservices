package web

import (
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/exception"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserRequestRegister struct {
	FullName string `json:"fullName" validate:"required"`
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	PassportRepeat string `json:"passwordRepeat" validate:"required"`
	CreatedAt *time.Time
	ModifiedAt *time.Time
}

func (u *UserRequestRegister) Validate() error {
	validate := validator.New()

	return validate.Struct(u)
}

func (u *UserRequestRegister) HashPassword() (string, *exception.ErrorResponse) {
	if u.Password != u.PassportRepeat {
		return "", exception.BadRequestError("Password Did Not Match")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return "", exception.NewInternalServerError(err.Error())
	}
	return string(hash), nil
}

