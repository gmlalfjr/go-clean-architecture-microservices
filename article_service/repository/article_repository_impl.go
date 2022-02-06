package repository

import (
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/collection"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/db"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/entity"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/exception"
	"go.mongodb.org/mongo-driver/bson"
)

type ArticleRepositoryImplementation struct {
	Collection collection.Collection
}

func NewArticleRepository(coll collection.Collection) ArticleRepository {
	return &ArticleRepositoryImplementation{
		Collection: coll,
	}
}

func (a ArticleRepositoryImplementation) CreateArticle(article *entity.Article) (*entity.Article, *exception.ErrorResponse) {

	ctx, cancel := db.NewMongoContext()
	defer cancel()

	_, err :=a.Collection.ArticleCollection().InsertOne(ctx, bson.M{
		"userId": article.UserId,
		"text":     article.Text,
		"title":    article.Title,
		"status": article.Status,
		"createdAt": article.CreatedAt,
		"modifiedAt": article.ModifiedAt,
	})
	if err != nil {
		return nil, &exception.ErrorResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}
	}

	return article, nil
}
