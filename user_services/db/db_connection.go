package db

import (
	"fmt"
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/helpers"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func Connection(conf helpers.Configuration) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Get("DB_USERNAME"),
		conf.Get("DB_PASSWORD"),
		conf.Get("PORT"),
		conf.Get("DATABASE"),
		)
	var err error
	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
	// Get generic database object sql.DB to use its functions
	sqlDB, err := client.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return client
}