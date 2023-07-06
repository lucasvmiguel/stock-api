package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

// Updates a product by id
func (s *Service) UpdateByID(ctx context.Context, id int, product entity.Product) (*entity.Product, error) {
	ctx = s.transactor.Begin(ctx)

	p, err := s.repository.UpdateByID(ctx, id, entity.Product{
		Name:          product.Name,
		StockQuantity: product.StockQuantity,
	})
	if err != nil {
		s.transactor.Rollback(ctx)
		return nil, errors.Wrap(err, "failed to update product")
	}

	defer s.transactor.Commit(ctx)

	if p == nil {
		return nil, nil
	}

	return p, nil
}
