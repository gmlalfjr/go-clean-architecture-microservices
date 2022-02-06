package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/controller"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/middleware"
)

func ArticleRoutes(ctx *gin.Engine, controller controller.ArticleController) {
	ctx.POST("/create", middleware.VerifyAuthorization , controller.CreateArticle)
}