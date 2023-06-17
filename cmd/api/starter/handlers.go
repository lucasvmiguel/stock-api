package starter

import (
	productHandler "github.com/lucasvmiguel/stock-api/internal/product/handler"
)

// handlers is a struct that holds all handlers
type handlers struct {
	product *productHandler.Handler
}

// createHandlersArgs is the arguments struct for createHandlers function
type createHandlersArgs struct {
	services services
}

// createHandlers creates all handlers
func (s *Starter) createHandlers(args createHandlersArgs) (handlers, error) {
	productHandler, err := productHandler.NewHandler(productHandler.NewHandlerArgs{
		Service:                args.services.product,
		PaginationDefaultLimit: s.config.PaginationDefaultLimit,
	})
	if err != nil {
		return handlers{}, err
	}

	return handlers{
		product: productHandler,
	}, nil
}
