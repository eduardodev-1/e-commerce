package product_tests

import (
	"bytes"
	"e-commerce/internal/core/adapters/repositories"
	"e-commerce/internal/core/domain/models"
	"e-commerce/internal/core/domain/services"
	"e-commerce/internal/core/handlers"
	"e-commerce/internal/infra/config"
	"e-commerce/internal/infra/database/postgres"
	"e-commerce/tests"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestUpdateProduct(t *testing.T) {
	db := postgres.NewPsqlConn()
	allRepositories := repositories.NewRepositories(db)
	allServices := services.NewServices(allRepositories)
	allHandlers := handlers.NewHandlers(allServices)
	app := config.GetFiberConfig()
	app.Use(tests.IsAuthenticatedMiddlewareMock)
	app.Put("/v1/product/:id", allHandlers.ProductHandler.Update)
	app.Get("/v1/product/:id", allHandlers.ProductHandler.Get)

	rng := rand.New(rand.NewSource(1))
	randomInt := rng.Intn(100) + 1
	// Teste para atualização bem-sucedida
	product := models.Product{
		Name:        "Updated Product " + fmt.Sprint(randomInt),
		Description: "Updated description " + fmt.Sprint(randomInt),
		ImgURL:      "http://example.com/updated.pn/" + fmt.Sprint(randomInt),
		Price:       float64(randomInt),
		SellerID:    1,
		Quantity:    randomInt,
	}
	body, _ := json.Marshal(product)
	req := httptest.NewRequest(http.MethodPut, "/v1/product/36", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Erro ao fazer a requisição: %v", err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)
	if err != nil {
		t.Fatalf("Erro ao decodificar o corpo da resposta: %v", err)
	}
	assert.Equal(t, "product updated successfully", response["message"])
	assert.NotEmpty(t, response["id"])
	productID, ok := response["id"].(float64)
	if !ok {
		t.Fatalf("ID do produto não encontrado na resposta inicial")
	}
	product.ID = int64(productID)

	//Teste para recuperar objeto atualizado
	req = getNewProductCreatedReq(strconv.FormatInt(product.ID, 10))
	resp, err = app.Test(req)
	if err != nil {
		t.Fatalf("Erro ao fazer a requisição: %v", err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var productUpdated models.Product
	decoder = json.NewDecoder(resp.Body)
	err = decoder.Decode(&productUpdated)
	if err != nil {
		t.Fatalf("Erro ao decodificar o corpo da resposta: %v", err)
	}
	assert.Equal(t, product.ID, productUpdated.ID)
	assert.Equal(t, product.Name, productUpdated.Name)
	assert.Equal(t, product.Description, productUpdated.Description)
	assert.Equal(t, product.ImgURL, productUpdated.ImgURL)
	assert.Equal(t, product.Price, productUpdated.Price)
	assert.Equal(t, product.SellerID, productUpdated.SellerID)
	assert.Equal(t, product.Quantity, productUpdated.Quantity)

	// Teste para ID inválido
	req = httptest.NewRequest(http.MethodPut, "/v1/product/invalid", nil)
	resp, err = app.Test(req)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// Teste para falha na atualização devido a solicitação inválida
	req = httptest.NewRequest(http.MethodPut, "/v1/product/1", nil) // Sem corpo
	resp, err = app.Test(req)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
