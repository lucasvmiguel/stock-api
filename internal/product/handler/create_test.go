package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleCreate(t *testing.T) {
	req, err := http.NewRequest("POST", "/products", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	h, _ := NewHandler(&mockRepo{})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleCreate)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	expected, _ := json.Marshal(fakeProduct)
	expectedString := strings.TrimSpace(string(expected))
	result := strings.TrimSpace(rr.Body.String())
	if result != expectedString {
		t.Errorf("handler returned unexpected body: got %v want %v",
			result, expectedString)
	}
}

func TestHandleCreateInvalidBody(t *testing.T) {
	reqBody, _ := json.Marshal(map[string]interface{}{
		"Name":          "",
		"StockQuantity": -10,
	})
	req, err := http.NewRequest("POST", "/products", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	h, _ := NewHandler(&mockRepo{})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleCreate)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestHandleCreateDBFailed(t *testing.T) {
	req, err := http.NewRequest("POST", "/products", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	h, _ := NewHandler(&mockBrokeRepo{})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleCreate)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
