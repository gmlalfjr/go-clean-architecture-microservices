package db

import (
	"context"
	"fmt"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"time"
)

func ConnectionMongo(configuration helpers.Configuration) *mongo.Database {
	ctx, cancel := NewMongoContext()
	defer cancel()

	option := options.Client().
		ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@clusterdb.4ynlz.mongodb.net/%s?retryWrites=true&w=majority",
			configuration.Get("DATABASE_USERNAME"),
			configuration.Get("DATABASE_PASSWORD"),
			configuration.Get("DATABASE"))).
		SetMinPoolSize(uint64(100)).
		SetMaxPoolSize(uint64(100)).
		SetMaxConnIdleTime(time.Duration(100) * time.Second)

	client, err := mongo.NewClient(option)
	if err != nil {
		panic(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	database := client.Database(configuration.Get("DATABASE"))
	return database
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}