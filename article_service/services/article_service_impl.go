package services

import (
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/entity"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/exception"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/repository"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/web"
	"time"
)

type ArticleServiceImpl struct {
	repo repository.ArticleRepository
}
func NewArticleService( repo repository.ArticleRepository) ArticleService {
	return &ArticleServiceImpl{repo: repo}
}

func (a ArticleServiceImpl) CreateArticle(request *web.ArticleCreateRequest) (*web.ArticleCreateResponse, *exception.ErrorResponse) {
	article := &entity.Article{
		UserId: request.UserId,
		Title:      request.Title,
		Text:       request.Text,
		Status:     request.Status,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}
	data, err := a.repo.CreateArticle(article)
	if err != nil {
		return nil, exception.NewError(err, err.Code, err.Status)
	}

	return &web.ArticleCreateResponse{
		Title:  data.Title,
		Text:   data.Text,
		Status: data.Status,
	}, nil
}
