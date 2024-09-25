package product_tests

import (
	"e-commerce/internal/core/adapters/repositories"
	"e-commerce/internal/core/domain/models"
	"e-commerce/internal/core/domain/services"
	"e-commerce/internal/core/handlers"
	"e-commerce/internal/infra/config"
	"e-commerce/internal/infra/database/postgres"
	"e-commerce/tests"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDeleteProduct(t *testing.T) {
	db := postgres.NewPsqlConn()
	allRepositories := repositories.NewRepositories(db)
	allServices := services.NewServices(allRepositories)
	allHandlers := handlers.NewHandlers(allServices)
	app := config.GetFiberConfig()
	app.Use(tests.IsAuthenticatedMiddlewareMock)
	app.Delete("/v1/product/:id", allHandlers.ProductHandler.Delete)
	app.Post("/v1/product", allHandlers.ProductHandler.Post)
	product := models.Product{
		Name:        "New Product Test",
		Description: "Product description Test",
		ImgURL:      "http://example.com/image.png/test",
		Price:       999.0,
		SellerID:    1,
		Quantity:    10,
	}
	// Teste para criação bem-sucedida
	successReq := getSuccessReq(product)
	resp, err := app.Test(successReq)
	if err != nil {
		t.Fatalf("Erro ao fazer a requisição: %v", err)
	}
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var response map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)
	if err != nil {
		t.Fatalf("Erro ao decodificar o corpo da resposta: %v", err)
	}
	assert.Equal(t, "product created successfully", response["message"])
	assert.NotEmpty(t, response["id"])
	product.ID = int64(response["id"].(float64))

	// Teste para exclusão bem-sucedida
	req := httptest.NewRequest(http.MethodDelete, "/v1/product/"+strconv.FormatInt(product.ID, 10), nil)
	resp, err = app.Test(req)
	if err != nil {
		t.Fatalf("Erro ao fazer a requisição: %v", err)
	}
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	if resp.StatusCode != http.StatusNoContent {
		decoder = json.NewDecoder(resp.Body)
		err = decoder.Decode(&response)
		if err != nil {
			t.Fatalf("Erro ao decodificar o corpo da resposta: %v", err)
		}
		log.Println(response["error"])
	}
	// Teste para ID inválido
	req = httptest.NewRequest(http.MethodDelete, "/v1/product/invalid", nil)
	resp, err = app.Test(req)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// Teste para produto não encontrado
	req = httptest.NewRequest(http.MethodDelete, "/v1/product/"+strconv.FormatInt(product.ID, 10), nil)
	resp, err = app.Test(req)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	if resp.StatusCode != http.StatusNotFound {
		decoder = json.NewDecoder(resp.Body)
		err = decoder.Decode(&response)
		if err != nil {
			t.Fatalf("Erro ao decodificar o corpo da resposta: %v", err)
		}
		log.Println(response["error"])
	}
}
