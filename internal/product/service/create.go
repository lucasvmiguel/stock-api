package service

import (
	"github.com/pkg/errors"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

// Creates a product
func (s *Service) Create(product entity.Product) (*entity.Product, error) {
	p, err := s.repository.Create(entity.Product{
		Name:          product.Name,
		StockQuantity: product.StockQuantity,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create product")
	}

	return p, nil
}
