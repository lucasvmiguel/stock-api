package service

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestDeleteByID_Successfully(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		DeleteByID(gomock.Eq(fakeProduct.ID)).
		Return(fakeProduct, nil)

	h, _ := NewService(repository)

	p, err := h.DeleteByID(fakeProduct.ID)
	if err != nil {
		t.Errorf("error should be nil, instead it got: %v", err)
	}

	if p == nil {
		t.Error("product should not be nil")
	}
}

func TestDeleteByID_RepositoryWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		DeleteByID(gomock.Eq(fakeProduct.ID)).
		Return(nil, errors.New(""))

	h, _ := NewService(repository)

	p, err := h.DeleteByID(fakeProduct.ID)
	if err == nil {
		t.Error("error should not be nil")
	}

	if p != nil {
		t.Error("product should be nil")
	}
}
