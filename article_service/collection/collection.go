package collection

import "go.mongodb.org/mongo-driver/mongo"

type Collection interface {
	ArticleCollection() *mongo.Collection
 }