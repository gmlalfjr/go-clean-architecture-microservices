package repository

import (
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/entity"
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/exception"
	"gorm.io/gorm"
)
type UserRepository interface {
	Register(db *gorm.DB, user *entity.User) (*entity.User, *exception.ErrorResponse)
	Login(db *gorm.DB, user *entity.User) (*entity.User, *exception.ErrorResponse)
}