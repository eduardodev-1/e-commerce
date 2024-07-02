package main

import (
	"github.com/eduardodev-1/e-commerce/internal/config"
	"github.com/eduardodev-1/e-commerce/internal/controller"
	"github.com/eduardodev-1/e-commerce/internal/database"
	"github.com/eduardodev-1/e-commerce/internal/middleware"
	"github.com/eduardodev-1/e-commerce/internal/repositories"
	"github.com/eduardodev-1/e-commerce/internal/routes"
	"github.com/eduardodev-1/e-commerce/internal/services"
	"log"
	"os"
)

func main() {
	//InitialConfigs
	app := config.GetFiberConfig()
	// DB Connection
	db, err := database.NewPsqlConn()
	if err != nil {
		log.Fatal(err.Error())
	}
	// Repositories
	allRepositories := repositories.NewRepositories(db)
	// Services
	allServices := services.NewServices(allRepositories)
	// Controllers
	allControllers := controller.NewControllers(allServices)
	// Routes
	routes.SetupPublicRoutes(app, allControllers)
	app.Use(middleware.AuthMiddleware)
	routes.SetupPrivateRoutes(app, allControllers)

	PORT := os.Getenv("FIBER_PORT")
	err = app.Listen(":" + PORT)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
