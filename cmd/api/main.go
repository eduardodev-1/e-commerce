package main

import (
	"e-commerce/internal/core/adapters/repositories"
	"e-commerce/internal/core/domain/services"
	"e-commerce/internal/core/handlers"
	"e-commerce/internal/infra/config"
	"e-commerce/internal/infra/database/postgres"
	"e-commerce/internal/routes"
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

	routes.Private(app, allHandlers)

	PORT := os.Getenv("FIBER_PORT")

	log.Fatal(app.Listen(":" + PORT))
}
