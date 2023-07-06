package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

// Creates a product
func (s *Service) Create(ctx context.Context, product entity.Product) (*entity.Product, error) {
	ctx = s.transactor.Begin(ctx)

	p, err := s.repository.Create(ctx, entity.Product{
		Name:          product.Name,
		StockQuantity: product.StockQuantity,
	})
	if err != nil {
		s.transactor.Rollback(ctx)
		return nil, errors.Wrap(err, "failed to create product")
	}

	s.transactor.Commit(ctx)
	return p, nil
}
