package handler

import (
	"time"

	"github.com/google/uuid"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
	"github.com/lucasvmiguel/stock-api/pkg/pagination"
)

// productResponseBody is the response body for a product
type productResponseBody struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Code          uuid.UUID `json:"code"`
	StockQuantity int       `json:"stock_quantity"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// productsResponseBody is the response body for a list of products
type productsResponseBody []productResponseBody

// paginatedProductsResponseBody is the response body for a paginated list of products
type paginatedProductsResponseBody struct {
	Items      []productResponseBody `json:"items"`
	NextCursor *uint                 `json:"next_cursor"`
}

// maps a product entity to a product response body
func mapProductToResponseBody(product *entity.Product) productResponseBody {
	return productResponseBody{
		ID:            product.ID,
		Name:          product.Name,
		StockQuantity: product.StockQuantity,
		Code:          product.Code,
		CreatedAt:     product.CreatedAt,
		UpdatedAt:     product.UpdatedAt,
	}
}

// maps a list of product entities to products response body
func mapProductsToResponseBody(products []*entity.Product) productsResponseBody {
	response := productsResponseBody{}

	for _, product := range products {
		response = append(response, mapProductToResponseBody(product))
	}

	return response
}

// maps a list of product entities to paginated products response body
func mapProductsToPaginatedResponseBody(result *pagination.Result[*entity.Product]) paginatedProductsResponseBody {
	responseBody := paginatedProductsResponseBody{
		NextCursor: result.NextCursor,
		Items:      []productResponseBody{},
	}

	for _, product := range result.Items {
		responseBody.Items = append(responseBody.Items, mapProductToResponseBody(product))
	}

	return responseBody
}
