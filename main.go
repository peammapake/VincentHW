package main

import (
	"fmt"

	"github.com/peammapake/vincent-inventory-api/initializer"
	"github.com/peammapake/vincent-inventory-api/internal/app"
	"github.com/peammapake/vincent-inventory-api/internal/server"
	"github.com/peammapake/vincent-inventory-api/pkg/routes"
	"github.com/peammapake/vincent-inventory-api/repository"
	"github.com/peammapake/vincent-inventory-api/router"
)

func main() {
	// Read config file using Viper
	config, err := initializer.ReadConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// Initialize Fiber framework router
	r := router.New()

	// Initialize DB connection
	err = initializer.ConnectToDB(config)
	if err != nil {
		panic(fmt.Errorf("fatal error connecting to database: %s", err))
	}
	// Unnescessary for GORM 1.2+ but just in case
	defer initializer.CloseDB()

	// Init Server
	s := initServer()

	// Routes
	routes.PublicRoutes(r, s)
	routes.NotFoundRoute(r, s)

	// Start port listening
	port := config.Server.ServerPort
	err = r.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(fmt.Errorf("fatal error starting server: %s", err))
	}

}

func initServer() *server.Server {
	repo := repository.NewRepository()
	app := app.NewApp(repo)
	return server.NewServer(app)
}
