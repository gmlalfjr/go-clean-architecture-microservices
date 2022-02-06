package middleware

import (
	"errors"
	"github.com/gmlalfjr/go-clean-architecture-microservices/article-service/exception"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func VerifyAuthorization(c *gin.Context) {

	getHeader := c.GetHeader("Authorization")
	if len(getHeader) <= 0 {
		c.JSON(http.StatusBadRequest, exception.NewError(errors.New("not authorize"), 400, "Bad Request Error"))
		c.Abort()
		return
	}
	if !strings.Contains(getHeader, "Bearer") {
		c.JSON(http.StatusBadRequest,exception.NewError(errors.New("not accepted"), 400, "Bad Request Error"))
		c.Abort()
		return
	}
	tokenString := strings.Replace(getHeader, "Bearer ", "", -1)
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("accessToken"), nil
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, exception.NewError(err, 400, "Something Error"))
		c.Abort()
		return
	}

	for key, val := range claims {
		c.Set(key, val)
	}
	c.Next()
}
