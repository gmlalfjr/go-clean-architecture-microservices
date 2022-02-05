package main

import (
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/app"
)

func main()  {
	applicationServer := app.RunServer()
	err := applicationServer.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
