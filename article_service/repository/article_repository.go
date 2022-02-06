package repository

import (
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/entity"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/exception"
)

type ArticleRepository interface {
	CreateArticle(article *entity.Article) (*entity.Article, *exception.ErrorResponse)
}