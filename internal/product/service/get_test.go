package service

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

func TestGetByID_Successfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		GetByID(gomock.Eq(fakeProduct.ID)).
		Return(fakeProduct, nil)

	h, _ := NewService(repository)

	p, err := h.GetByID(fakeProduct.ID)
	if err != nil {
		t.Errorf("error should be nil, instead it got: %v", err)
	}

	if p == nil {
		t.Error("product should not be nil")
	}
}

func TestGetByID_RepositoryWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		GetByID(gomock.Eq(fakeProduct.ID)).
		Return(nil, errors.New(""))

	h, _ := NewService(repository)

	p, err := h.GetByID(fakeProduct.ID)
	if err == nil {
		t.Error("error should not be nil")
	}

	if p != nil {
		t.Error("product should be nil")
	}
}

func TestGetAll_Successfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		GetAll().
		Return([]*entity.Product{fakeProduct, fakeProduct}, nil)

	h, _ := NewService(repository)

	p, err := h.GetAll()
	if err != nil {
		t.Errorf("error should be nil, instead it got: %v", err)
	}

	if p == nil {
		t.Error("products should not be nil")
	}
}

func TestGetAll_RepositoryWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		GetAll().
		Return(nil, errors.New(""))

	h, _ := NewService(repository)

	p, err := h.GetAll()
	if err == nil {
		t.Error("error should not be nil")
	}

	if p != nil {
		t.Error("products should be nil")
	}
}
