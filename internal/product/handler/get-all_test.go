package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

func TestHandleGetAll(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	h, _ := NewHandler(&mockRepo{})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleGetAll)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected, _ := json.Marshal([]entity.Product{*fakeProduct, *fakeProduct})
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

	h, _ := NewHandler(&mockBrokeRepo{})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleGetAll)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
