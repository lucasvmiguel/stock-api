package service

import (
	"github.com/pkg/errors"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

// gets a product by id
func (s *Service) GetByID(id uint) (*entity.Product, error) {
	p, err := s.repository.GetByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get product by id")
	}

	return p, nil
}

// gets all products
func (s *Service) GetAll() ([]*entity.Product, error) {
	p, err := s.repository.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all products")
	}

	return p, nil
}
