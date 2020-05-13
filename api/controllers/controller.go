package controllers

import (
	"net/http"
	"os"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/thiagozs/hostgator-challenge/api/database"
	"github.com/thiagozs/hostgator-challenge/libs"
)

// ICtlRepo base repository
type ICtlRepo interface {
	GetJwt() *jwt.GinJWTMiddleware
	FindAccount(c *gin.Context)
	CreateAccount(c *gin.Context)
	Breeds(c *gin.Context)
	Welcome(c *gin.Context)
	Ping(c *gin.Context)
	RefreshHandler() func(c *gin.Context)
	LoginHandler() func(c *gin.Context)
	MiddlewareFunc() gin.HandlerFunc
	NoRoute(c *gin.Context)
}

// CtlRepo struct
type CtlRepo struct {
	DB database.IGormRepo
}

// NewCtlRepo new repository
func NewCtlRepo(db database.IGormRepo) ICtlRepo {
	return &CtlRepo{db}
}

// GetJwt auth middleware
func (ctl *CtlRepo) GetJwt() *jwt.GinJWTMiddleware {
	ajwt, err := libs.AuthMiddleware(ctl.DB)
	if err != nil {
		os.Exit(1)
	}
	return ajwt
}

// NoRoute show 404 when not have a endpoint
func (ctl *CtlRepo) NoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Page not found"})
}
