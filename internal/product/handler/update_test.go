package handler

import (
	"bytes"
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
)

func TestHandleUpdate(t *testing.T) {
	req, err := http.NewRequest(http.MethodPut, "/products/1", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	service := NewMockService(ctrl)
	service.
		EXPECT().
		UpdateByID(gomock.Eq(uint(1)), gomock.Eq(entity.Product{
			Name:          fakeProduct.Name,
			StockQuantity: fakeProduct.StockQuantity,
		})).
		Return(fakeProduct, nil)

	h, _ := NewHandler(NewHandlerArgs{Service: service})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleUpdate)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected, _ := json.Marshal(fakeProduct)
	expectedString := strings.TrimSpace(string(expected))
	result := strings.TrimSpace(rr.Body.String())
	if result != expectedString {
		t.Errorf("handler returned unexpected body: got %v want %v",
			result, expectedString)
	}
}

func TestHandleUpdateInvalidBody(t *testing.T) {
	reqBody, _ := json.Marshal(map[string]any{
		"name":           "",
		"stock_quantity": -10,
	})
	req, err := http.NewRequest(http.MethodPut, "/products/1", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	service := NewMockService(ctrl)
	service.
		EXPECT().
		UpdateByID(nil, nil).
		Return(nil, nil).
		Times(0)

	h, _ := NewHandler(NewHandlerArgs{Service: service})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleUpdate)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestHandleUpdateDBFailed(t *testing.T) {
	req, err := http.NewRequest(http.MethodPut, "/products/1", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	service := NewMockService(ctrl)
	service.
		EXPECT().
		UpdateByID(gomock.Eq(uint(1)), gomock.Eq(entity.Product{
			Name:          fakeProduct.Name,
			StockQuantity: fakeProduct.StockQuantity,
		})).
		Return(nil, errors.New(""))

	h, _ := NewHandler(NewHandlerArgs{Service: service})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleUpdate)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
