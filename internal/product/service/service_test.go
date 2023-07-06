package service

import (
	"testing"
	"time"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/uuid"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

var (
	fakeTime = time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC)

	fakeProduct = &entity.Product{
		ID:            1,
		Code:          uuid.New(),
		Name:          "Product Foo",
		StockQuantity: 10,
		CreatedAt:     fakeTime,
		UpdatedAt:     fakeTime,
	}
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	transactor := NewMockTransactor(ctrl)
	repository := NewMockRepository(ctrl)

	_, err := NewService(NewServiceArgs{
		Repository: repository,
		Transactor: transactor,
	})
	if err != nil {
		t.Error("should not return error when repository is not nil")
	}
}

func TestNewServiceError(t *testing.T) {
	_, err := NewService(NewServiceArgs{})
	if err == nil {
		t.Error("error should be returned when no repository is passed")
	}
}
