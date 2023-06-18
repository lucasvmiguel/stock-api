package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	gomock "github.com/golang/mock/gomock"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

func TestHandleCreate(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	service := NewMockService(ctrl)
	service.
		EXPECT().
		Create(gomock.Eq(entity.Product{
			Name:          fakeProduct.Name,
			StockQuantity: fakeProduct.StockQuantity,
		})).
		Return(fakeProduct, nil)

	h, _ := NewHandler(NewHandlerArgs{Service: service})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleCreate)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	expected, _ := json.Marshal(fakeProductResponseBody)
	expectedString := strings.TrimSpace(string(expected))
	result := strings.TrimSpace(rr.Body.String())
	if result != expectedString {
		t.Errorf("handler returned unexpected body: got %v want %v",
			result, expectedString)
	}
}

func TestHandleCreateInvalidBody(t *testing.T) {
	reqBody, _ := json.Marshal(map[string]any{
		"name":           "",
		"stock_quantity": -10,
	})
	req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	service := NewMockService(ctrl)
	service.
		EXPECT().
		Create(nil).
		Return(nil, nil).
		Times(0)

	h, _ := NewHandler(NewHandlerArgs{Service: service})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleCreate)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestHandleCreateDBFailed(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	service := NewMockService(ctrl)
	service.
		EXPECT().
		Create(gomock.Eq(entity.Product{
			Name:          fakeProduct.Name,
			StockQuantity: fakeProduct.StockQuantity,
		})).
		Return(nil, errors.New(""))

	h, _ := NewHandler(NewHandlerArgs{Service: service})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleCreate)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
