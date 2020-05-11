package services

import (
	"fmt"
	"net/http"
	"os"

	"hostgator-challenge/controllers"
	_ "hostgator-challenge/docs"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (sr *Server) PublicRoutes() {

	auth, err := authMiddleware()
	if err != nil {
		os.Exit(1)
	}

	sr.Engine.GET("/", controllers.Welcome)
	sr.Engine.GET("/ping", controllers.Ping)
	sr.Engine.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler))

	sr.Engine.POST("/login", auth.LoginHandler)
	sr.Engine.GET("/login/:id", controllers.FindLogin)

	sr.Engine.GET("/refresh_token", auth.RefreshHandler)

	sr.Engine.POST("/account", controllers.CreateLogin)
	sr.Engine.GET("/account/:id", controllers.FindLogin)
}

func (sr *Server) PrivateRoutes() {

	auth, err := authMiddleware()
	if err != nil {
		os.Exit(1)
	}

	priv := sr.Engine.Group("/breeds")
	priv.Use(auth.MiddlewareFunc())
	priv.GET("/:cat", func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		fmt.Printf("%#v\n", claims)
		c.JSON(http.StatusOK, gin.H{"message": "I have jwt!"})
	})

}
