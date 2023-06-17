package starter

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/lucasvmiguel/stock-api/internal/product/handler"
	"github.com/lucasvmiguel/stock-api/pkg/logger"
)

// createRoutesArgs is the arguments struct for createRoutes function
type createRoutesArgs struct {
	router   *chi.Mux
	handlers handlers
}

// createRoutes creates all HTTP routes
func (s *Starter) createRoutes(args createRoutesArgs) {
	// http middlewares
	args.router.Use(middleware.RequestID)
	args.router.Use(middleware.RealIP)
	args.router.Use(logger.HTTPMiddleware(s.config.ServiceName))
	args.router.Use(middleware.Recoverer)
	args.router.Use(middleware.Timeout(60 * time.Second))

	// api http routes group (v1)
	args.router.Route("/api/v1", func(r chi.Router) {
		// product http routes
		r.Get("/products", args.handlers.product.HandleGetPaginated)
		r.Get("/products/all", args.handlers.product.HandleGetAll)
		r.Post("/products", args.handlers.product.HandleCreate)
		r.Get(fmt.Sprintf("/products/{%s}", handler.FieldID), args.handlers.product.HandleGetByID)
		r.Delete(fmt.Sprintf("/products/{%s}", handler.FieldID), args.handlers.product.HandleDeleteByID)
		r.Put(fmt.Sprintf("/products/{%s}", handler.FieldID), args.handlers.product.HandleUpdate)
		r.Patch(fmt.Sprintf("/products/{%s}", handler.FieldID), args.handlers.product.HandleUpdate)
	})

	// health http route
	args.router.Get("/health", func(w http.ResponseWriter, req *http.Request) { w.Write([]byte("Up and running")) })
}
