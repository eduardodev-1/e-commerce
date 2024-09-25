package product_tests

import (
	"e-commerce/internal/core/adapters/repositories"
	"e-commerce/internal/core/domain/models"
	"e-commerce/internal/core/domain/services"
	"e-commerce/internal/core/handlers"
	"e-commerce/internal/infra/config"
	"e-commerce/internal/infra/database/postgres"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProduct(t *testing.T) {
	db := postgres.NewPsqlConn()
	allRepositories := repositories.NewRepositories(db)
	allServices := services.NewServices(allRepositories)
	allHandlers := handlers.NewHandlers(allServices)
	app := config.GetFiberConfig()

	app.Get("/v1/product/:id", allHandlers.ProductHandler.Get)

	// Teste para produto existente
	req := httptest.NewRequest(http.MethodGet, "/v1/product/1", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Erro ao fazer a requisição: %v", err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var product models.Product
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&product)
	if err != nil {
		t.Fatalf("Erro ao decodificar o corpo da resposta: %v", err)
	}
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.Name)
	assert.NotEmpty(t, product.Description)
	assert.NotEmpty(t, product.ImgURL)
	assert.NotEmpty(t, product.Price)
	assert.NotEmpty(t, product.SellerID)
	assert.NotEmpty(t, product.Quantity)

	// Teste para ID inválido
	req = httptest.NewRequest(http.MethodGet, "/v1/product/invalid", nil)
	resp, err = app.Test(req)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// Teste para produto não encontrado
	req = httptest.NewRequest(http.MethodGet, "/v1/product/9999", nil)
	resp, err = app.Test(req)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
