package main

import (
	"e-commerce/internal/config"
	"e-commerce/internal/controller"
	"e-commerce/internal/database"
	"e-commerce/internal/middleware"
	"e-commerce/internal/repositories"
	"e-commerce/internal/routes"
	"e-commerce/internal/services"
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
	routes.PublicRoutes(app, allControllers)
	app.Use(middleware.AuthMiddleware)
	routes.PrivateRoutes(app, allControllers)

	PORT := os.Getenv("FIBER_PORT")
	err = app.Listen(":" + PORT)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
