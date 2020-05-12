package services

import (
	"fmt"
	"hostgator-challenge/database"
	"hostgator-challenge/libs"
	"hostgator-challenge/models"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "uuid"

type Authorizator func(data interface{}, c *gin.Context) bool
type Authenticator func(c *gin.Context) (interface{}, error)

func authMiddleware(db database.IGormRepo) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("my secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     identityKey,
		PayloadFunc:     payloadFunc,
		IdentityHandler: identityHandler,
		Authenticator:   authenticator(db),
		Authorizator:    authorizator(db),
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

func authorizator(db database.IGormRepo) Authorizator {
	return func(data interface{}, c *gin.Context) bool {
		if v, ok := data.(*models.Login); ok {
			var login models.Login
			if err := db.FindOne(models.Login{ID: v.ID}, &login); err != nil {
				return false
			}
			if v.UserName == login.UserName {
				return true
			}
		}
		return false
	}
}

func authenticator(db database.IGormRepo) Authenticator {
	return func(c *gin.Context) (interface{}, error) {
		var loginVals models.Login
		if err := c.ShouldBind(&loginVals); err != nil {
			return "", jwt.ErrMissingLoginValues
		}

		var login models.Login
		if err := db.FindOne(models.Login{ID: loginVals.ID}, &login); err != nil {
			return models.Login{}, err
		}

		pwd := libs.NewPasswordGen()
		match, err := pwd.Compare(loginVals.Password, login.Password)
		if !match || err != nil {
			fmt.Println("Password Invalid")
		}

		if loginVals.UserName == login.UserName && match {
			return &models.Login{UserName: loginVals.UserName}, nil
		}

		return nil, jwt.ErrFailedAuthentication
	}
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
