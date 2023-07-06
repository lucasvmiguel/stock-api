package service

import (
	"context"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestDeleteByID_Successfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	transactor := NewMockTransactor(ctrl)
	transactor.EXPECT().Begin(gomock.Any()).Return(context.Background())
	transactor.EXPECT().Commit(gomock.Any())
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		DeleteByID(gomock.Any(), gomock.Eq(fakeProduct.ID)).
		Return(fakeProduct, nil)

	s, _ := NewService(NewServiceArgs{
		Repository: repository,
		Transactor: transactor,
	})

	p, err := s.DeleteByID(context.Background(), fakeProduct.ID)
	if err != nil {
		t.Errorf("error should be nil, instead it got: %v", err)
	}

	if p == nil {
		t.Error("product should not be nil")
	}
}

func TestDeleteByID_RepositoryWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	transactor := NewMockTransactor(ctrl)
	transactor.EXPECT().Begin(gomock.Any()).Return(context.Background())
	transactor.EXPECT().Rollback(gomock.Any())
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		DeleteByID(gomock.Any(), gomock.Eq(fakeProduct.ID)).
		Return(nil, errors.New(""))

	s, _ := NewService(NewServiceArgs{
		Repository: repository,
		Transactor: transactor,
	})

	p, err := s.DeleteByID(context.Background(), fakeProduct.ID)
	if err == nil {
		t.Error("error should not be nil")
	}

	if p != nil {
		t.Error("product should be nil")
	}
}
