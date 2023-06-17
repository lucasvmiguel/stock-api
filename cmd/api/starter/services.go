package starter

import (
	productService "github.com/lucasvmiguel/stock-api/internal/product/service"
)

// services is a struct that holds all services
type services struct {
	product *productService.Service
}

// createServicesArgs is the arguments struct for createServices function
type createServicesArgs struct {
	repositories repositories
}

// createServices creates all services
func (s *Starter) createServices(args createServicesArgs) (services, error) {
	productSvc, err := productService.NewService(args.repositories.product)
	if err != nil {
		return services{}, err
	}

	return services{
		product: productSvc,
	}, nil
}
