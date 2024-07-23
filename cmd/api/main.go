package main

import (
	"e-commerce/internal/config"
	"e-commerce/internal/core/adapters/repositories"
	"e-commerce/internal/core/adapters/services"
	"e-commerce/internal/database/postgres"
	"e-commerce/internal/network/handlers"
	"e-commerce/internal/network/middleware"
	"e-commerce/internal/network/routes"
	"log"
	"os"
)

func main() {
	db := postgres.NewPsqlConn()

	allRepositories := repositories.NewRepositories(db)

	allServices := services.NewServices(allRepositories)

	allHandlers := handlers.NewHandlers(allServices)

	app := config.GetFiberConfig()

	routes.Public(app, allHandlers)
	app.Use(middleware.AuthMiddleware)
	routes.Private(app, allHandlers)

	PORT := os.Getenv("FIBER_PORT")
	err := app.Listen(":" + PORT)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
