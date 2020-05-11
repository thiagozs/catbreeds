package services

import (
	"hostgator-challenge/models"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "uuid"

func authMiddleware() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("my secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     identityKey,
		PayloadFunc:     payloadFunc,
		IdentityHandler: identityHandler,
		Authenticator:   authenticator,
		Authorizator:    authorizator,
		Unauthorized:    unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})
}

func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*models.Login); ok && v.UserName == "admin@gmail.com" {
		return true
	}
	return false
}

func authenticator(c *gin.Context) (interface{}, error) {
	var loginVals models.Login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	username := loginVals.UserName
	password := loginVals.Password

	if (username == "admin" && password == "123456") ||
		(username == "test" && password == "test") {
		return &models.Login{
			UserName: username,
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &models.Login{
		UserName: claims[identityKey].(string),
	}
}

func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*models.Login); ok {
		return jwt.MapClaims{
			identityKey: v.UserName,
		}
	}
	return jwt.MapClaims{}
}
