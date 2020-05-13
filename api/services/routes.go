package services

import (
	_ "github.com/thiagozs/hostgator-challenge/api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) PublicRoutes() {

	s.Engine.NoRoute(s.Ctl.NoRoute)

	s.Engine.GET("/", s.Ctl.Welcome)
	s.Engine.GET("/ping", s.Ctl.Ping)
	s.Engine.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.Engine.POST("/login",
		s.Ctl.LoginHandler())

	s.Engine.GET("/login/refresh_token",
		s.Ctl.RefreshHandler())

	s.Engine.POST("/account", s.Ctl.CreateAccount)
	s.Engine.GET("/account/:id", s.Ctl.FindAccount)
}

func (s *Server) PrivateRoutes() {

	priv := s.Engine.Group("/breeds")
	priv.Use(s.Ctl.MiddlewareFunc())
	priv.GET("/:cat", s.Ctl.Breeds)

}
