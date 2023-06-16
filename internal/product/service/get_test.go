package service

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

const (
	LIMIT  = uint(10)
	CURSOR = uint(100)
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

func TestGetPaginated_Successfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		GetPaginated(gomock.Eq(CURSOR), gomock.Eq(LIMIT)).
		Return([]*entity.Product{fakeProduct, fakeProduct}, nil)

	h, _ := NewService(repository)

	result, err := h.GetPaginated(CURSOR, LIMIT)
	if err != nil {
		t.Errorf("error should be nil, instead it got: %v", err)
	}

	if len(result.Items) != 2 {
		t.Error("2 products should be returned")
	}

	if result.NextCursor != &fakeProduct.ID {
		t.Error("next cursor should be last product id")
	}
}

func TestGetPaginated_SuccessfullyButNoMoreProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		GetPaginated(gomock.Eq(CURSOR), gomock.Eq(LIMIT)).
		Return([]*entity.Product{}, nil)

	h, _ := NewService(repository)

	result, err := h.GetPaginated(CURSOR, LIMIT)
	if err != nil {
		t.Errorf("error should be nil, instead it got: %v", err)
	}

	if len(result.Items) != 0 {
		t.Error("no products should be returned")
	}

	if result.NextCursor != nil {
		t.Error("next cursor should be nil")
	}
}

func TestGetPaginated_InvalidLimit(t *testing.T) {
	limit := uint(10000)
	ctrl := gomock.NewController(t)
	repository := NewMockRepository(ctrl)

	h, _ := NewService(repository)

	_, err := h.GetPaginated(CURSOR, limit)
	if err == nil {
		t.Error("error should not be nil")
	}
}

func TestGetPaginated_RepositoryWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		GetPaginated(gomock.Eq(CURSOR), gomock.Eq(LIMIT)).
		Return(nil, errors.New(""))

	h, _ := NewService(repository)

	_, err := h.GetPaginated(CURSOR, LIMIT)
	if err == nil {
		t.Error("error should not be nil")
	}
}
