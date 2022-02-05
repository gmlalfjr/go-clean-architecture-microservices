package repository

import (
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/entity"
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/exception"
	"gorm.io/gorm"
)



type UserRepositoryImpl struct {

}
func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u UserRepositoryImpl) Register(db *gorm.DB, user *entity.User) (*entity.User, *exception.ErrorResponse) {
	if err := db.Table("users").Create(&user).Error; err !=nil {

		return nil, &exception.ErrorResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}
	}
	return user, nil
}

func (u UserRepositoryImpl) Login(db *gorm.DB, user *entity.User) (*entity.User, *exception.ErrorResponse){
	if err := db.Table("users").Where("email = ?", user.Email).First(&user).Error; err != nil {
		return  nil, &exception.ErrorResponse{
			Code:   500,
			Status: "Not Found user",
			Data:   err.Error(),
		}
	}
	return user, nil
}
