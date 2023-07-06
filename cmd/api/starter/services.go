package starter

import (
	productService "github.com/lucasvmiguel/stock-api/internal/product/service"
	"github.com/lucasvmiguel/stock-api/pkg/transactor"
)

// services is a struct that holds all services
type services struct {
	product *productService.Service
}

// createServicesArgs is the arguments struct for createServices function
type createServicesArgs struct {
	Repositories repositories
	Transactor   *transactor.Transactor
}

// createServices creates all services
func (s *Starter) createServices(args createServicesArgs) (services, error) {
	productSvc, err := productService.NewService(productService.NewServiceArgs{
		Repository: args.Repositories.product,
		Transactor: args.Transactor,
	})
	if err != nil {
		return services{}, err
	}

	return services{
		product: productSvc,
	}, nil
}
