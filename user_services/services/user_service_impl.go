package services

import (
	"fmt"
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/entity"
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/exception"
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/repository"
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/web"
	"gorm.io/gorm"
	"time"
)

type UserServiceImpl struct {
	repo repository.UserRepository
	DB *gorm.DB
}
func NewUserService(repo repository.UserRepository, db *gorm.DB) UserService {
	return &UserServiceImpl{
		repo: repo,
		DB:   db,
	}
}
func (u UserServiceImpl) Register( request *web.UserRequestRegister) (*web.UserResponseRegister, *exception.ErrorResponse) {
	err := request.Validate()
	if err != nil {
		return nil, exception.BadRequestError(err.Error())
	}
	pass, errHash := request.HashPassword()
	if errHash != nil {
		return nil, errHash
	}
	user := &entity.User{
		FullName:   request.FullName,
		Email:      request.Email,
		Password:   pass,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}
	data, errRegister := u.repo.Register(u.DB, user)
	if errRegister != nil {
		return nil, errRegister
	}

	return &web.UserResponseRegister{
		FullName:   data.FullName,
		Email:      data.Email,
		CreatedAt:  &data.CreatedAt,
		ModifiedAt: &data.ModifiedAt,
	}, nil
}

func (u UserServiceImpl) Login( request *web.UserRequestLogin) (*web.UserResponseLogin, *exception.ErrorResponse) {
	err := request.Validate()
	if err != nil {
		return nil, exception.BadRequestError(err.Error())
	}
	user := &entity.User{
		Email:      request.Email,
		Password:   request.Password,
	}
	data, errLogin := u.repo.Login(u.DB, user)
	if errLogin != nil {
		return nil, errLogin
	}
	errCheckPassword := request.CheckPasswordHash(data.Password)
	if errCheckPassword != nil {
		return nil, errCheckPassword
	}
	var userResponse web.UserResponseLogin
	generateToken, errGenerate := userResponse.GenerateToken(map[string]string{
		"id": fmt.Sprint(data.Id),
		"email": data.Email,
	})
	if errGenerate != nil {
		return nil, errGenerate
	}

	return generateToken, nil
}
