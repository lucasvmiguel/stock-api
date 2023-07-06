package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	gomock "github.com/golang/mock/gomock"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

func TestHandleGetAll(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	products := []*entity.Product{fakeProduct, fakeProduct}
	ctrl := gomock.NewController(t)
	service := NewMockService(ctrl)
	service.
		EXPECT().
		GetAll(gomock.Any()).
		Return(products, nil)

	h, _ := NewHandler(NewHandlerArgs{Service: service})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleGetAll)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected, _ := json.Marshal([]productResponseBody{fakeProductResponseBody, fakeProductResponseBody})
	expectedString := strings.TrimSpace(string(expected))
	result := strings.TrimSpace(rr.Body.String())
	if result != expectedString {
		t.Errorf("handler returned unexpected body: got %v want %v",
			result, expectedString)
	}
}

func TestHandleGetAllDBFailed(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	service := NewMockService(ctrl)
	service.
		EXPECT().
		GetAll(gomock.Any()).
		Return(nil, errors.New(""))

	h, _ := NewHandler(NewHandlerArgs{Service: service})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleGetAll)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
