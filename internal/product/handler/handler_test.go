package handler

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

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
		"Name":          fakeProduct.Name,
		"StockQuantity": fakeProduct.StockQuantity,
	})
)

type mockRepo struct{}

func (r *mockRepo) Create(product entity.Product) (*entity.Product, error) {
	return fakeProduct, nil
}

func (r *mockRepo) GetAll() ([]*entity.Product, error) {
	return []*entity.Product{fakeProduct, fakeProduct}, nil
}

func (r *mockRepo) GetByID(id uint) (*entity.Product, error) {
	if id == nonexistentID {
		return nil, nil
	}
	return fakeProduct, nil
}

func (r *mockRepo) UpdateByID(id uint, product entity.Product) (*entity.Product, error) {
	if id == nonexistentID {
		return nil, nil
	}
	return fakeProduct, nil
}

func (r *mockRepo) DeleteByID(id uint) (*entity.Product, error) {
	if id == nonexistentID {
		return nil, nil
	}
	return fakeProduct, nil
}

type mockBrokeRepo struct{}

func (r *mockBrokeRepo) Create(product entity.Product) (*entity.Product, error) {
	return nil, errors.New("")
}

func (r *mockBrokeRepo) GetAll() ([]*entity.Product, error) {
	return nil, errors.New("")
}

func (r *mockBrokeRepo) GetByID(id uint) (*entity.Product, error) {
	return nil, errors.New("")
}

func (r *mockBrokeRepo) UpdateByID(id uint, product entity.Product) (*entity.Product, error) {
	return nil, errors.New("")
}

func (r *mockBrokeRepo) DeleteByID(id uint) (*entity.Product, error) {
	return nil, errors.New("")
}

func TestNewHandler(t *testing.T) {
	_, err := NewHandler(&mockRepo{})
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
