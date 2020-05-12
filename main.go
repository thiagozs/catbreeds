package main

import (
	"fmt"
	"os"

	_ "hostgator-challenge/docs"
	"hostgator-challenge/models"
	"hostgator-challenge/services"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// @title CATAPI
// @version 1.0
// @description A little server about cats and yours highlights.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://thiagozs.com/
// @contact.email thiago.zilli@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080

// @BasePath /
func main() {

	var version string = "1.0.0"

	fmt.Printf("Start Server %s on port :8080...\n", version)
	// options...
	opts := func(srv *services.Server) {
		srv.Debug = true
		srv.DialectDB = "sqlite3"
		srv.FileNameDB = "database.db"
		srv.Models = append(srv.Models, &models.Login{})
		srv.Models = append(srv.Models, &models.CatAPI{})
	}

	srv := services.NewServer(opts)
	srv.StartDB()
	srv.PublicRoutes()
	srv.PrivateRoutes()

	if err := srv.Run(); err != nil {
		fmt.Printf("Error on run server, got : %s\n", err.Error())
		os.Exit(1)
	}

}
