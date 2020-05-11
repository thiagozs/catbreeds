package services

import (
	"hostgator-challenge/controllers"

	_ "hostgator-challenge/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (sr *Server) PublicRoutes() {
	sr.Engine.GET("/", controllers.Welcome)
	sr.Engine.GET("/ping", controllers.Ping)
	sr.Engine.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler))

	login := sr.Engine.Group("/login")
	login.POST("", controllers.CreateLogin)
	login.GET("/:id", controllers.FindLogin)
}

func (sr *Server) PrivateRoutes() {

}
