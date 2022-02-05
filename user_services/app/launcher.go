package app

import (
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/controller"
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/db"
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/entity"
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/helpers"
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/repository"
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/routes"
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/services"
	"github.com/gofiber/fiber/v2"
)

func RunServer() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName: "Test App v1.0.1",
	})
	conf := helpers.NewConfiguration()
	dbConn := db.Connection(conf)
	err := dbConn.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}
	userRepository := repository.NewUserRepository()
	userService := services.NewUserService(userRepository, dbConn)
	userController := controller.NewController(userService)
	routes.Route(app, userController)

	return app
}
