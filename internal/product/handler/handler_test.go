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

	reqBody, _ = json.Marshal(map[string]any{
		"name":           fakeProduct.Name,
		"stock_quantity": fakeProduct.StockQuantity,
	})
)

func TestNewHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	service := NewMockService(ctrl)

	_, err := NewHandler(NewHandlerArgs{Service: service})
	if err != nil {
		t.Error("should not return error when service is not nil")
	}
}

func TestNewHandlerError(t *testing.T) {
	_, err := NewHandler(NewHandlerArgs{})
	if err == nil {
		t.Error("error should be returned when no service is passed")
	}
}
