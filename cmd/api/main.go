package routes

import (
	"e-commerce/internal/core/adapters/handlers"
	"e-commerce/internal/core/adapters/repositories"
	"e-commerce/internal/core/domain/services"
	"e-commerce/internal/infra/config"
	"e-commerce/internal/infra/database/postgres"
	"log"
	"os"
)

func main() {
	db := postgres.NewPsqlConn()

	allRepositories := repositories.NewRepositories(db)

	allServices := services.NewServices(allRepositories)

	allHandlers := handlers.NewHandlers(allServices)

	app := config.GetFiberConfig()

	Public(app, allHandlers)

	Private(app, allHandlers)

	PORT := os.Getenv("FIBER_PORT")

	if err := app.Listen(":" + PORT); err != nil {
		log.Fatal(err.Error())
	}
}
