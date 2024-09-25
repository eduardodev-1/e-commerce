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
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestPostProduct(t *testing.T) {
	db := postgres.NewPsqlConn()
	allRepositories := repositories.NewRepositories(db)
	allServices := services.NewServices(allRepositories)
	allHandlers := handlers.NewHandlers(allServices)
	app := config.GetFiberConfig()
	app.Use(tests.IsAuthenticatedMiddlewareMock)
	app.Post("/v1/product", allHandlers.ProductHandler.Post)
	app.Get("/v1/product/:id", allHandlers.ProductHandler.Get)
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

	//Teste para recuperar objeto criado
	req := getNewProductCreatedReq(strconv.FormatInt(product.ID, 10))
	resp, err = app.Test(req)
	if err != nil {
		t.Fatalf("Erro ao fazer a requisição: %v", err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var productCreated models.Product
	decoder = json.NewDecoder(resp.Body)
	err = decoder.Decode(&productCreated)
	if err != nil {
		t.Fatalf("Erro ao decodificar o corpo da resposta: %v", err)
	}
	assert.Equal(t, product.ID, productCreated.ID)
	assert.Equal(t, product.Name, productCreated.Name)
	assert.Equal(t, product.Description, productCreated.Description)
	assert.Equal(t, product.ImgURL, productCreated.ImgURL)
	assert.Equal(t, product.Price, productCreated.Price)
	assert.Equal(t, product.SellerID, productCreated.SellerID)
	assert.Equal(t, product.Quantity, productCreated.Quantity)

	// Teste para falha na criação devido a solicitação inválida
	withoutBodyReq := getWithoutBodyReq()
	resp, err = app.Test(withoutBodyReq)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func getNewProductCreatedReq(s string) *http.Request {
	return httptest.NewRequest(http.MethodGet, "/v1/product/"+s, nil)
}

func getWithoutBodyReq() *http.Request {
	return httptest.NewRequest(http.MethodPost, "/v1/product", nil) // Sem corpo
}

func getSuccessReq(product models.Product) *http.Request {
	body, _ := json.Marshal(product)
	req := httptest.NewRequest(http.MethodPost, "/v1/product", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	return req
}
