package service

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

func TestUpdateByID_Successfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		UpdateByID(gomock.Eq(uint(1)), gomock.Eq(entity.Product{
			Name:          fakeProduct.Name,
			StockQuantity: fakeProduct.StockQuantity,
		})).
		Return(fakeProduct, nil)

	h, _ := NewService(repository)

	p, err := h.UpdateByID(uint(1), entity.Product{
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

func TestUpdateByID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		UpdateByID(gomock.Eq(uint(1)), gomock.Eq(entity.Product{
			Name:          fakeProduct.Name,
			StockQuantity: fakeProduct.StockQuantity,
		})).
		Return(nil, nil)

	h, _ := NewService(repository)

	p, err := h.UpdateByID(uint(1), entity.Product{
		Name:          fakeProduct.Name,
		StockQuantity: fakeProduct.StockQuantity,
	})
	if err != nil {
		t.Errorf("error should be nil, instead it got: %v", err)
	}

	if p != nil {
		t.Error("product should be nil")
	}
}

func TestUpdateByID_RepositoryWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		UpdateByID(gomock.Eq(uint(0)), gomock.Eq(entity.Product{
			Name:          fakeProduct.Name,
			StockQuantity: fakeProduct.StockQuantity,
		})).
		Return(nil, errors.New(""))

	h, _ := NewService(repository)

	p, err := h.UpdateByID(uint(0), entity.Product{
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
