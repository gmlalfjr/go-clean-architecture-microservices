package web

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gmlalfjr/go-clean-architecture-microservices/users_service/exception"
	"time"
)

type UserResponseLogin struct {
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (u *UserResponseLogin) GenerateToken(data map[string] string) (*UserResponseLogin, *exception.ErrorResponse) {

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  data["id"],
		"email": data["email"],
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})
	generateRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  data["id"],
		"email": data["email"],
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})
	tokenString, err := generateToken.SignedString([]byte("accessToken"))
	refreshTokenString, errRefresh := generateRefreshToken.SignedString([]byte("refreshToken"))
	if err != nil || errRefresh != nil {
		return nil, exception.NewInternalServerError("Failed Generated Token")
	}
	u.AccessToken = tokenString
	u.RefreshToken = refreshTokenString
	return u, nil
}
