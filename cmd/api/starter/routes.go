package starter

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/lucasvmiguel/stock-api/internal/product/handler"
)

// createRoutes creates all HTTP routes
func (s *Starter) createRoutes() {
	// http middlewares
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.Timeout(60 * time.Second))

	// api http routes group (v1)
	s.router.Route("/api/v1", func(r chi.Router) {
		// product http routes
		r.Get("/products", s.handlers.product.HandleGetPaginated)
		r.Get("/products/all", s.handlers.product.HandleGetAll)
		r.Post("/products", s.handlers.product.HandleCreate)
		r.Get(fmt.Sprintf("/products/{%s}", handler.FieldID), s.handlers.product.HandleGetByID)
		r.Delete(fmt.Sprintf("/products/{%s}", handler.FieldID), s.handlers.product.HandleDeleteByID)
		r.Put(fmt.Sprintf("/products/{%s}", handler.FieldID), s.handlers.product.HandleUpdate)
		r.Patch(fmt.Sprintf("/products/{%s}", handler.FieldID), s.handlers.product.HandleUpdate)
	})

	// health http route
	s.router.Get("/health", func(w http.ResponseWriter, req *http.Request) { w.Write([]byte("Up and running")) })
}
