package collection

import (
	"go.mongodb.org/mongo-driver/mongo"
)


type CollectionImpl struct {
	DB *mongo.Database
}

func NewCollection(db *mongo.Database) Collection {
	return &CollectionImpl{DB: db }
}
func (c CollectionImpl) ArticleCollection() *mongo.Collection {
	return c.DB.Collection("article")
}