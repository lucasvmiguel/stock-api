package service

import (
	"github.com/lucasvmiguel/stock-api/internal/product/entity"
	"github.com/pkg/errors"
)

// Updates a product by id
func (s *Service) UpdateByID(id uint, product entity.Product) (*entity.Product, error) {
	p, err := s.repository.UpdateByID(id, entity.Product{
		Name:          product.Name,
		StockQuantity: product.StockQuantity,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to update product")
	}

	if p == nil {
		return nil, nil
	}

	return p, nil
}
