package main

import (
	"e-commerce/internal/config"
	"e-commerce/internal/core/services"
	"e-commerce/internal/database/postgres"
	"e-commerce/internal/handler"
	"e-commerce/internal/middleware"
	"e-commerce/internal/repositories"
	"e-commerce/internal/routes"
	"log"
	"os"
)

func main() {
	db := postgres.NewPsqlConn()

	allRepositories := repositories.NewRepositories(db)

	allServices := services.NewServices(allRepositories)

	allHandlers := handler.NewHandlers(allServices)

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
