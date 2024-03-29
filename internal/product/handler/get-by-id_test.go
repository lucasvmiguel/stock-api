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
)

func TestHandleGetByID(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/products/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	service := NewMockService(ctrl)
	service.
		EXPECT().
		GetByID(gomock.Any(), gomock.Eq(1)).
		Return(fakeProduct, nil)

	h, _ := NewHandler(NewHandlerArgs{Service: service})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleGetByID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected, _ := json.Marshal(fakeProductResponseBody)
	expectedString := strings.TrimSpace(string(expected))
	result := strings.TrimSpace(rr.Body.String())
	if result != expectedString {
		t.Errorf("handler returned unexpected body: got %v want %v",
			result, expectedString)
	}
}

func TestHandleGetByIDNotFound(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/products/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	service := NewMockService(ctrl)
	service.
		EXPECT().
		GetByID(gomock.Any(), gomock.Eq(nonexistentID)).
		Return(nil, nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "0")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	h, _ := NewHandler(NewHandlerArgs{Service: service})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleGetByID)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestHandleGetByIDDBFailed(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/products/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	service := NewMockService(ctrl)
	service.
		EXPECT().
		GetByID(gomock.Any(), gomock.Eq(1)).
		Return(nil, errors.New(""))

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	h, _ := NewHandler(NewHandlerArgs{Service: service})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleGetByID)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
