package starter

import (
	productService "github.com/lucasvmiguel/stock-api/internal/product/service"
)

// services is a struct that holds all services
type services struct {
	product *productService.Service
}

// createServices creates all services
func (s *Starter) createServices() (services, error) {
	productSvc, err := productService.NewService(s.repositories.product)
	if err != nil {
		return services{}, err
	}

	return services{
		product: productSvc,
	}, nil
}
