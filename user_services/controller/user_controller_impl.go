package controller

import (
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/exception"
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/services"
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/web"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type UserControllerImpl struct {
	service services.UserService
}

func NewController(service services.UserService) UserController{
	return &UserControllerImpl{service: service}
}


func (u UserControllerImpl) Register(ctx *fiber.Ctx) error {
	var request web.UserRequestRegister
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.JSON(exception.BadRequestError(err.Error()))
	}
	response, errRegister := u.service.Register(&request)
	if errRegister != nil {
		return ctx.Status(errRegister.Code).JSON(exception.ErrorResponse{
			Code:   errRegister.Code,
			Status: errRegister.Status,
			Data:   errRegister.Data,
		})
	}
	return ctx.Status(http.StatusCreated).JSON(web.SuccessCreate(response))
}

func (u UserControllerImpl) Login(ctx *fiber.Ctx) error {
	var request web.UserRequestLogin
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.JSON(exception.BadRequestError(err.Error()))
	}
	response, errLogin := u.service.Login(&request)

	if errLogin != nil {
		return ctx.Status(errLogin.Code).JSON(exception.ErrorResponse{
			Code:   errLogin.Code,
			Status: errLogin.Status,
			Data:   errLogin.Data,
		})
	}
	return ctx.JSON(web.Response{
		Code:   200,
		Status: "Success Login",
		Data:   response,
	})
}
