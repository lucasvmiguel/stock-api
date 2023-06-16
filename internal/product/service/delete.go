package service

import (
	"github.com/pkg/errors"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

// deletes a product by id
func (s *Service) DeleteByID(id uint) (*entity.Product, error) {
	p, err := s.repository.DeleteByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete product by id")
	}

	return p, nil
}
