package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestHandleUpdate(t *testing.T) {
	req, err := http.NewRequest("PUT", "/products/1", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	h, _ := NewHandler(&mockRepo{})

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
	reqBody, _ := json.Marshal(map[string]interface{}{
		"Name":          "",
		"StockQuantity": -10,
	})
	req, err := http.NewRequest("PUT", "/products/1", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	h, _ := NewHandler(&mockRepo{})

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
	req, err := http.NewRequest("PUT", "/products/1", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	h, _ := NewHandler(&mockBrokeRepo{})

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
