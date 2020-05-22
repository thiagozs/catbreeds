package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"catbreeds/api/controllers"
	"catbreeds/api/database"
	_ "catbreeds/api/docs"
	"catbreeds/api/models"
	"catbreeds/api/services"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	port  = flag.String("p", ":8080", "default port for api")
	debug = flag.Bool("d", false, "log debug is on or off")
)

// @title Codding challenge
// @version 2.0
// @description Simple documentation of API.
// @termsOfService https://thiagozs.com/terms/

// @contact.name API Support
// @contact.url https://thiagozs.com
// @contact.email thiago.zilli@gmail.com

// @license.name Reserved Commons
// @license.url https://thiagozs.com/license

// @host localhost:8080
// @schemes http
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Bearer
func main() {

	flag.Parse()

	version := "1.0.0"
	fmt.Printf("Start Server v%s on port %s...\n", version, *port)

	d, err := gorm.Open("sqlite3", filepath.Base("database.db"))
	if err != nil {
		fmt.Printf("Error on start data base, got: %s\n ", err.Error())
		os.Exit(1)
	}
	db := database.NewGormRepo(d)
	ctl := controllers.NewCtlRepo(db)

	// options server...
	opts := func(s *services.Server) {
		s.Port = *port
		s.Debug = *debug
		s.Models = append(s.Models, &models.CatAPI{}, &models.Login{})
		s.DB = db
		s.Ctl = ctl
	}

	s := services.NewServer(opts)
	s.MigrationDB()
	s.PublicRoutes()
	s.PrivateRoutes()

	if err := s.Run(); err != nil {
		fmt.Printf("Error on run server, got : %s\n", err.Error())
		os.Exit(1)
	}

}
