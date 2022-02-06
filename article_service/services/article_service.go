package services

import (
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/exception"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/web"
)

type ArticleService interface {
	CreateArticle(request *web.ArticleCreateRequest)(*web.ArticleCreateResponse, *exception.ErrorResponse)
}