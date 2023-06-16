package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	gomock "github.com/golang/mock/gomock"
	"github.com/lucasvmiguel/stock-api/internal/product/entity"
	"github.com/lucasvmiguel/stock-api/pkg/pagination"
)

const (
	LIMIT  = uint(10)
	CURSOR = uint(100)
)

func TestHandleGetPaginated(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/products?limit=10&cursor=100", nil)
	if err != nil {
		t.Fatal(err)
	}

	nextCursor := uint(110)
	productsPaginated := &pagination.Result[*entity.Product]{
		Items:      []*entity.Product{fakeProduct},
		NextCursor: &nextCursor,
	}

	ctrl := gomock.NewController(t)
	service := NewMockService(ctrl)
	service.
		EXPECT().
		GetPaginated(gomock.Eq(CURSOR), gomock.Eq(LIMIT)).
		Return(productsPaginated, nil)

	h, _ := NewHandler(NewHandlerArgs{Service: service})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleGetPaginated)

	rctx := chi.NewRouteContext()
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected, _ := json.Marshal(productsPaginated)
	expectedString := strings.TrimSpace(string(expected))
	result := strings.TrimSpace(rr.Body.String())
	if result != expectedString {
		t.Errorf("handler returned unexpected body: got %v want %v",
			result, expectedString)
	}
}

func TestHandleGetPaginatedEnd(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/products?limit=10&cursor=200", nil)
	if err != nil {
		t.Fatal(err)
	}

	productsPaginated := &pagination.Result[*entity.Product]{
		Items:      []*entity.Product{},
		NextCursor: nil,
	}

	ctrl := gomock.NewController(t)
	service := NewMockService(ctrl)
	service.
		EXPECT().
		GetPaginated(gomock.Eq(uint(200)), gomock.Eq(LIMIT)).
		Return(productsPaginated, nil)

	h, _ := NewHandler(NewHandlerArgs{Service: service})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleGetPaginated)

	rctx := chi.NewRouteContext()
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected, _ := json.Marshal(productsPaginated)
	expectedString := strings.TrimSpace(string(expected))
	result := strings.TrimSpace(rr.Body.String())
	if result != expectedString {
		t.Errorf("handler returned unexpected body: got %v want %v",
			result, expectedString)
	}
}

func TestHandleGetPaginated_InvalidLimit(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/products?limit=100000", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	service := NewMockService(ctrl)

	h, _ := NewHandler(NewHandlerArgs{Service: service})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleGetPaginated)

	rctx := chi.NewRouteContext()
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestHandleGetPaginated_InvalidCursor(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/products?cursor=abc", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	service := NewMockService(ctrl)

	h, _ := NewHandler(NewHandlerArgs{Service: service})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleGetPaginated)

	rctx := chi.NewRouteContext()
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestHandleGetPaginated_FailedService(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/products?limit=10&cursor=100", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	service := NewMockService(ctrl)
	service.
		EXPECT().
		GetPaginated(gomock.Eq(CURSOR), gomock.Eq(LIMIT)).
		Return(nil, errors.New("error"))

	h, _ := NewHandler(NewHandlerArgs{Service: service})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleGetPaginated)

	rctx := chi.NewRouteContext()
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
