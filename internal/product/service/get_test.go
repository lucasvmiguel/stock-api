package service

import (
	"context"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

const (
	LIMIT  = 10
	CURSOR = 100
)

func TestGetByID_Successfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	transactor := NewMockTransactor(ctrl)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		GetByID(gomock.Any(), gomock.Eq(fakeProduct.ID)).
		Return(fakeProduct, nil)

	s, _ := NewService(NewServiceArgs{
		Repository: repository,
		Transactor: transactor,
	})

	p, err := s.GetByID(context.Background(), fakeProduct.ID)
	if err != nil {
		t.Errorf("error should be nil, instead it got: %v", err)
	}

	if p == nil {
		t.Error("product should not be nil")
	}
}

func TestGetByID_RepositoryWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	transactor := NewMockTransactor(ctrl)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		GetByID(gomock.Any(), gomock.Eq(fakeProduct.ID)).
		Return(nil, errors.New(""))

	s, _ := NewService(NewServiceArgs{
		Repository: repository,
		Transactor: transactor,
	})

	p, err := s.GetByID(context.Background(), fakeProduct.ID)
	if err == nil {
		t.Error("error should not be nil")
	}

	if p != nil {
		t.Error("product should be nil")
	}
}

func TestGetAll_Successfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	transactor := NewMockTransactor(ctrl)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		GetAll(gomock.Any()).
		Return([]*entity.Product{fakeProduct, fakeProduct}, nil)

	s, _ := NewService(NewServiceArgs{
		Repository: repository,
		Transactor: transactor,
	})

	p, err := s.GetAll(context.Background())
	if err != nil {
		t.Errorf("error should be nil, instead it got: %v", err)
	}

	if p == nil {
		t.Error("products should not be nil")
	}
}

func TestGetAll_RepositoryWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	transactor := NewMockTransactor(ctrl)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		GetAll(gomock.Any()).
		Return(nil, errors.New(""))

	s, _ := NewService(NewServiceArgs{
		Repository: repository,
		Transactor: transactor,
	})

	p, err := s.GetAll(context.Background())
	if err == nil {
		t.Error("error should not be nil")
	}

	if p != nil {
		t.Error("products should be nil")
	}
}

func TestGetPaginated_Successfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	transactor := NewMockTransactor(ctrl)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		GetPaginated(gomock.Any(), gomock.Eq(CURSOR), gomock.Eq(LIMIT)).
		Return([]*entity.Product{fakeProduct, fakeProduct}, nil)

	s, _ := NewService(NewServiceArgs{
		Repository: repository,
		Transactor: transactor,
	})

	result, err := s.GetPaginated(context.Background(), CURSOR, LIMIT)
	if err != nil {
		t.Errorf("error should be nil, instead it got: %v", err)
	}

	if len(result.Items) != 2 {
		t.Error("2 products should be returned")
	}

	if *result.NextCursor != fakeProduct.ID {
		t.Error("next cursor should be last product id")
	}
}

func TestGetPaginated_SuccessfullyButNoMoreProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	transactor := NewMockTransactor(ctrl)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		GetPaginated(gomock.Any(), gomock.Eq(CURSOR), gomock.Eq(LIMIT)).
		Return([]*entity.Product{}, nil)

	s, _ := NewService(NewServiceArgs{
		Repository: repository,
		Transactor: transactor,
	})

	result, err := s.GetPaginated(context.Background(), CURSOR, LIMIT)
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
	limit := 10000
	ctrl := gomock.NewController(t)
	transactor := NewMockTransactor(ctrl)
	repository := NewMockRepository(ctrl)

	s, _ := NewService(NewServiceArgs{
		Repository: repository,
		Transactor: transactor,
	})

	_, err := s.GetPaginated(context.Background(), CURSOR, limit)
	if err == nil {
		t.Error("error should not be nil")
	}
}

func TestGetPaginated_RepositoryWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	transactor := NewMockTransactor(ctrl)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		GetPaginated(gomock.Any(), gomock.Eq(CURSOR), gomock.Eq(LIMIT)).
		Return(nil, errors.New(""))

	s, _ := NewService(NewServiceArgs{
		Repository: repository,
		Transactor: transactor,
	})

	_, err := s.GetPaginated(context.Background(), CURSOR, LIMIT)
	if err == nil {
		t.Error("error should not be nil")
	}
}
