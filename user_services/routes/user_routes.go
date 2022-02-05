package routes

import (
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/controller"
	"github.com/gofiber/fiber/v2"
)

func Route(f *fiber.App, controller controller.UserController) {
	f.Post("/api/register", controller.Register)
	f.Post("/api/login", controller.Login)
}
