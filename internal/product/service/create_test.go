package service

import (
	"context"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

func TestCreate_Successfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	transactor := NewMockTransactor(ctrl)
	transactor.EXPECT().Begin(gomock.Any()).Return(context.Background())
	transactor.EXPECT().Commit(gomock.Any())
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		Create(gomock.Any(), gomock.Eq(entity.Product{
			Name:          fakeProduct.Name,
			StockQuantity: fakeProduct.StockQuantity,
		})).
		Return(fakeProduct, nil)

	s, _ := NewService(NewServiceArgs{
		Repository: repository,
		Transactor: transactor,
	})

	p, err := s.Create(context.Background(), entity.Product{
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
	transactor := NewMockTransactor(ctrl)
	transactor.EXPECT().Begin(gomock.Any()).Return(context.Background())
	transactor.EXPECT().Rollback(gomock.Any())
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		Create(gomock.Any(), gomock.Eq(entity.Product{
			Name:          fakeProduct.Name,
			StockQuantity: fakeProduct.StockQuantity,
		})).
		Return(nil, errors.New(""))

	s, _ := NewService(NewServiceArgs{
		Repository: repository,
		Transactor: transactor,
	})

	p, err := s.Create(context.Background(), entity.Product{
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
