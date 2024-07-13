package main

import (
	"e-commerce/internal/config"
	"e-commerce/internal/controller"
	"e-commerce/internal/core/services"
	"e-commerce/internal/database/postgres"
	"e-commerce/internal/middleware"
	"e-commerce/internal/repositories"
	"e-commerce/internal/routes"
	"log"
	"os"
)

func main() {
	db := postgres.NewPsqlConn()

	app := config.GetFiberConfig()

	allRepositories, _ := repositories.NewRepositories(db, "postgres")

	allServices := services.NewServices(allRepositories)

	allControllers := controller.NewControllers(allServices)

	routes.PublicRoutes(app, allControllers)
	app.Use(middleware.AuthMiddleware)
	routes.PrivateRoutes(app, allControllers)

	PORT := os.Getenv("FIBER_PORT")
	err := app.Listen(":" + PORT)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
