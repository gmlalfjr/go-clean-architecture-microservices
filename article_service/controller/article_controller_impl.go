package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/exception"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/services"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/web"
	"net/http"
)

type ArticleControllerImpl struct {
	service services.ArticleService
}

func NewArticleController(service services.ArticleService) ArticleController {
	return &ArticleControllerImpl{service: service}
}

func (a ArticleControllerImpl) CreateArticle(ctx *gin.Context) {
	article := &web.ArticleCreateRequest{}
	userId, ok := ctx.Get("id")
	
	if ok == true {
		article.UserId = fmt.Sprint(userId)
	} else {
		ctx.JSON(http.StatusBadGateway, exception.ErrorResponse{
			Code:   http.StatusBadGateway,
			Status: "Something Error",
			Data:   "Failed get user",
		})
		return
	}
	if errBindJson := ctx.ShouldBindJSON(article); errBindJson != nil {
		ctx.JSON(http.StatusBadGateway, exception.ErrorResponse{
			Code:   400,
			Status: "Body Request Error",
			Data:   errBindJson.Error(),
		})
		return
	}

	data, err := a.service.CreateArticle(article)
	if err != nil {
		ctx.JSON(err.Code, exception.ErrorResponse{
			Code:   err.Code,
			Status: err.Status,
			Data:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, web.Response{
		Code:   http.StatusCreated,
		Status: data.Status,
		Data:   data,
	})
	return
}
