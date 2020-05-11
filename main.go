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

// @host thiagozs.com

// @BasePath /
func main() {

	var version string = "1.0.0"

	fmt.Printf("Start Server %s on port :8080...\n", version)
	// options...
	debug := func(srv *services.Server) {
		srv.Debug = true
	}

	models := func(srv *services.Server) {
		srv.Models = append(srv.Models, &models.Login{})
	}

	dbset := func(srv *services.Server) {
		srv.DialectDB = "sqlite3"
		srv.FileNameDB = "database.db"
	}

	srv := services.NewServer(debug, models, dbset)
	srv.StartDB()
	srv.PublicRoutes()
	srv.PrivateRoutes()

	if err := srv.Run(); err != nil {
		fmt.Printf("Error on run server, got : %s\n", err.Error())
		os.Exit(1)
	}

}
