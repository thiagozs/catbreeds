package services

import (
	"github.com/gin-gonic/gin"

	"catbreeds/api/controllers"
	"catbreeds/api/database"
	mid "catbreeds/api/middlewares"
)

type Option func(sr *Server)

type Server struct {
	Engine *gin.Engine
	Models []interface{}
	Port   string
	Debug  bool
	DB     database.IGormRepo
	Ctl    controllers.ICtlRepo
}

// NewServer start a new service
func NewServer(opts ...Option) *Server {
	s := Server{}

	// get all options need
	for _, option := range opts {
		option(&s)
	}

	// Set are debug or not
	gin.SetMode(gin.ReleaseMode)
	if s.Debug {
		gin.SetMode(gin.DebugMode)
	}

	// load gin framework
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(mid.Cors())

	// Set Gin engine
	s.Engine = r

	return &s
}

// StartDB start all process for connection database and models
func (s *Server) MigrationDB() {
	// Running the migrations
	s.DB.GetDB().AutoMigrate(s.Models...)
}

func (s *Server) Run() error {
	// forward runner
	return s.Engine.Run(s.Port)
}

func (s *Server) Stop() {
	//TODO: improve this method
	s.DB.GetDB().Close()
}
