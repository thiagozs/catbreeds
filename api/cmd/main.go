package main

import (
	"fmt"
	"os"
	"path/filepath"

	"hostgator-challenge/api/controllers"
	"hostgator-challenge/api/database"
	_ "hostgator-challenge/api/docs"
	"hostgator-challenge/api/models"
	"hostgator-challenge/api/services"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

	version := "1.0.0"
	fmt.Printf("Start Server %s on port :8080...\n", version)

	d, err := gorm.Open("sqlite3", filepath.Base("database.db"))
	if err != nil {
		fmt.Printf("Error on start data base, got: %s\n ", err.Error())
		os.Exit(1)
	}
	db := database.NewGormRepo(d)
	ctl := controllers.NewCtlRepo(db)

	// options server...
	opts := func(s *services.Server) {
		s.Debug = false
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
