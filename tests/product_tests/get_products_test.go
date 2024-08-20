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

func TestGetProducts(t *testing.T) {
	db := postgres.NewPsqlConn()
	allRepositories := repositories.NewRepositories(db)
	allServices := services.NewServices(allRepositories)
	allHandlers := handlers.NewHandlers(allServices)
	app := config.GetFiberConfig()

	app.Get("/v1/product", allHandlers.ProductHandler.GetPaginatedList)

	// Crie uma requisição de teste
	req := httptest.NewRequest(http.MethodGet, "/v1/product", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Erro ao fazer a requisição: %v", err)
	}

	// Verifique o status da resposta
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Verifique o tipo de conteúdo
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	// Verifique o corpo da resposta
	var page models.Page
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&page)
	if err != nil {
		t.Fatalf("Erro ao decodificar o corpo da resposta: %v", err)
	}

	// Converta o conteúdo para o tipo esperado
	contentSlice, ok := page.Content.([]interface{})
	if !ok {
		t.Fatalf("Tipo de conteúdo inválido: %T", page.Content)
	}

	var products []*models.Product
	for _, item := range contentSlice {
		productMap, ok := item.(map[string]interface{})
		if !ok {
			t.Fatalf("Tipo de item inválido: %T", item)
		}
		product := &models.Product{
			ID:          int64(productMap["id"].(float64)),
			Name:        productMap["name"].(string),
			Description: productMap["description"].(string),
			ImgURL:      productMap["img_url"].(string),
			Price:       productMap["price"].(float64),
			SellerID:    int64(productMap["seller_id"].(float64)),
			Quantity:    int(productMap["quantity"].(float64)),
		}
		products = append(products, product)
	}
	assert.Greater(t, len(products), 0)

	for _, product := range products {
		assert.NotEmpty(t, product.ID)
		assert.NotEmpty(t, product.Name)
		assert.NotEmpty(t, product.Description)
		assert.NotEmpty(t, product.ImgURL)
		assert.NotEmpty(t, product.Price)
		assert.NotEmpty(t, product.SellerID)
		assert.NotEmpty(t, product.Quantity)
	}
	assert.IsType(t, 0, page.TotalElements, "TotalElements deve ser do tipo int, mas é %T", page.TotalElements)
	assert.IsType(t, 0, page.TotalPages, "TotalPages deve ser do tipo int, mas é %T", page.TotalPages)
	assert.IsType(t, 0, page.Size, "Size deve ser do tipo int, mas é %T", page.Size)
	assert.IsType(t, 0, page.Number, "Number deve ser do tipo int, mas é %T", page.Number)
	assert.IsType(t, "", page.Sort, "Sort deve ser do tipo string, mas é %T", page.Sort)
	assert.Equal(t, "id", page.Sort)
	assert.Equal(t, 0, page.Number)
}
