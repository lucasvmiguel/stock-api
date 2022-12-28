package handler

import (
	"encoding/json"
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

var (
	nonexistentID = uint(0)
	fakeTime      = time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC)

	fakeProduct = &entity.Product{
		ID:            1,
		Code:          uuid.New(),
		Name:          "Product Foo",
		StockQuantity: 10,
		CreatedAt:     fakeTime,
		UpdatedAt:     fakeTime,
	}

	reqBody, _ = json.Marshal(map[string]interface{}{
		"name":           fakeProduct.Name,
		"stock_quantity": fakeProduct.StockQuantity,
	})
)

func TestNewHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := NewMockRepository(ctrl)

	_, err := NewHandler(repository)
	if err != nil {
		t.Error("should not return error when repository is not nil")
	}
}

func TestNewHandlerError(t *testing.T) {
	_, err := NewHandler(nil)
	if err == nil {
		t.Error("error should be returned when no repository is passed")
	}
}
