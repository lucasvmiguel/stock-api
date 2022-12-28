package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi/v5"
	gomock "github.com/golang/mock/gomock"
)

func TestHandleDeleteByID(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, "/products/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	service := NewMockService(ctrl)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		DeleteByID(gomock.Eq(uint(1))).
		Return(fakeProduct, nil)

	h, _ := NewHandler(repository, service)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleDeleteByID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}

func TestHandleDeleteByIDNotFound(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, "/products/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	service := NewMockService(ctrl)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		DeleteByID(gomock.Eq(nonexistentID)).
		Return(nil, nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", strconv.FormatUint(uint64(nonexistentID), 10))
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	h, _ := NewHandler(repository, service)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleDeleteByID)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestHandleDeleteByIDDBFailed(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, "/products/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	service := NewMockService(ctrl)
	repository := NewMockRepository(ctrl)
	repository.
		EXPECT().
		DeleteByID(gomock.Eq(uint(1))).
		Return(nil, errors.New(""))

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	h, _ := NewHandler(repository, service)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleDeleteByID)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
