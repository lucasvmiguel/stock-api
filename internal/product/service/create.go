package service

import (
	"github.com/lucasvmiguel/stock-api/internal/product/entity"
	"github.com/pkg/errors"
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
