package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/collection"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/controller"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/db"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/helpers"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/repository"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/routes"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/services"
)

func RunServer() *gin.Engine {
	r := gin.Default()
	conf := helpers.NewConfiguration()
	conn := db.ConnectionMongo(conf)
	newCollection := collection.NewCollection(conn)
	newRepo := repository.NewArticleRepository(newCollection)
	newService := services.NewArticleService(newRepo)
	newController := controller.NewArticleController(newService)

	routes.ArticleRoutes(r, newController)

	return r
}