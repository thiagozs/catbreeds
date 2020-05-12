package services

import (
	"fmt"
	"hostgator-challenge/database"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Option func(sr *Server)

type Server struct {
	Engine     *gin.Engine
	Models     []interface{}
	Port       string
	Debug      bool
	DialectDB  string
	FileNameDB string
	DB         *gorm.DB
	DBI        database.IGormRepo
}

// NewServer start a new service
func NewServer(opts ...Option) *Server {
	srv := Server{}

	// get all options need
	for _, option := range opts {
		option(&srv)
	}

	// Set are debug or not
	gin.SetMode(gin.ReleaseMode)
	if srv.Debug {
		gin.SetMode(gin.DebugMode)
	}

	// load gin framework
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(corsMiddleware())

	// set gin to engine
	srv.Engine = r

	return &srv
}

// StartDB start all process for connection database and models
func (sr *Server) StartDB() {
	// load database instance
	dialect, err := gorm.Open(sr.DialectDB, filepath.Base(sr.FileNameDB))
	if err != nil {
		fmt.Printf("Error on start data base, got: %s\n ", err.Error())
		os.Exit(1)
	}
	sr.DB = dialect
	//defer dialect.Close()

	// Running the migrations
	dialect.AutoMigrate(sr.Models...)

	// Start a new dialect to Repository
	sr.DBI = database.NewGormRepo(dialect)

	// use a context of gin to forward databse settings and conn.
	sr.Engine.Use(func(c *gin.Context) {
		c.Set("db", sr.DBI)
		c.Next()
	})

}

func (sr *Server) Run() error {
	// forward runner
	return sr.Engine.Run()
}

func (sr *Server) Stop() {
	//TODO: improve this method
	sr.DB.Close()
}
