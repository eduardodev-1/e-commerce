package handlers

import (
	"e-commerce/internal/core/adapters/repositories"
	"e-commerce/internal/core/domain/services"
	"e-commerce/internal/infra/config"
	"e-commerce/internal/infra/database/postgres"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProducts(t *testing.T) {
	db := postgres.NewPsqlConn()

	allRepositories := repositories.NewRepositories(db)

	allServices := services.NewServices(allRepositories)

	allHandlers := NewHandlers(allServices)

	app := config.GetFiberConfig()

	app.Get("/v1/product", allHandlers.ProductHandler.GetPaginatedList)

	// Crie uma requisição de teste
	req := httptest.NewRequest(http.MethodGet, "/v1/product", nil)
	resp, _ := app.Test(req)

	// Verifique o status da resposta
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
