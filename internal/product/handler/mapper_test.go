package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
	"github.com/lucasvmiguel/stock-api/pkg/pagination"
)

func TestMapProductToResponseBody(t *testing.T) {
	responseBody := mapProductToResponseBody(fakeProduct)

	assert.Equal(t, responseBody, fakeProductResponseBody)
}

func TestMapProductsToResponseBody(t *testing.T) {
	responseBody := mapProductsToResponseBody([]*entity.Product{fakeProduct})

	assert.Equal(t, responseBody, productsResponseBody{fakeProductResponseBody})
}

func TestMapProductsToPaginatedResponseBody(t *testing.T) {
	fakeNextCursor := 10

	result := &pagination.Result[*entity.Product]{
		NextCursor: &fakeNextCursor,
		Items: []*entity.Product{
			fakeProduct,
		},
	}

	responseBody := mapProductsToPaginatedResponseBody(result)

	assert.Equal(t, responseBody, paginatedProductsResponseBody{
		NextCursor: &fakeNextCursor,
		Items:      []productResponseBody{fakeProductResponseBody},
	})
}
