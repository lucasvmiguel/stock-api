package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
	"github.com/lucasvmiguel/stock-api/pkg/pagination"
)

// gets a product by id
func (s *Service) GetByID(ctx context.Context, id int) (*entity.Product, error) {
	product, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get product by id")
	}

	return product, nil
}

// gets all products
func (s *Service) GetAll(ctx context.Context) ([]*entity.Product, error) {
	products, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all products")
	}

	return products, nil
}

// gets products paginated
// cursor is the last id from the previous page
func (s *Service) GetPaginated(ctx context.Context, cursor int, limit int) (*pagination.Result[*entity.Product], error) {
	if limit < 1 || limit > 100 {
		return nil, errors.New("limit must be between 1 and 100")
	}

	products, err := s.repository.GetPaginated(ctx, cursor, limit)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get products paginated")
	}

	if len(products) == 0 {
		return &pagination.Result[*entity.Product]{
			Items:      products,
			NextCursor: nil,
		}, nil
	}

	lastProduct := products[len(products)-1]

	return &pagination.Result[*entity.Product]{
		Items:      products,
		NextCursor: &lastProduct.ID,
	}, nil
}
