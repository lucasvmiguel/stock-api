package service

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

func TestCreate_Successfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		Create(gomock.Eq(entity.Product{
			Name:          fakeProduct.Name,
			StockQuantity: fakeProduct.StockQuantity,
		})).
		Return(fakeProduct, nil)

	h, _ := NewService(repository)

	p, err := h.Create(entity.Product{
		Name:          fakeProduct.Name,
		StockQuantity: fakeProduct.StockQuantity,
	})
	if err != nil {
		t.Errorf("error should be nil, instead it got: %v", err)
	}

	if p == nil {
		t.Error("product should not be nil")
	}
}

func TestCreate_RepositoryWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		Create(gomock.Eq(entity.Product{
			Name:          fakeProduct.Name,
			StockQuantity: fakeProduct.StockQuantity,
		})).
		Return(nil, errors.New(""))

	h, _ := NewService(repository)

	p, err := h.Create(entity.Product{
		Name:          fakeProduct.Name,
		StockQuantity: fakeProduct.StockQuantity,
	})
	if err == nil {
		t.Error("error should not be nil")
	}

	if p != nil {
		t.Error("product should be nil")
	}
}
