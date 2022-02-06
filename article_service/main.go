package main

import "github.com/gmlalfjr/go-clean-architecture-microservices/article-service/app"


func main() {
	r := app.RunServer()
	err := r.Run()
	if err != nil {
		panic(err)
	}
}