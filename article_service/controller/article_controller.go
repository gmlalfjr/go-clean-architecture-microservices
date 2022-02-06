package controller

import "github.com/gin-gonic/gin"

type ArticleController interface {
	CreateArticle(ctx *gin.Context)
}
