package services

import (
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/exception"
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/web"
)

type UserService interface {
	Register(request *web.UserRequestRegister) (*web.UserResponseRegister, *exception.ErrorResponse)
	Login(request *web.UserRequestLogin) (*web.UserResponseLogin, *exception.ErrorResponse)
}
