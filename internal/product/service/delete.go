package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

// deletes a product by id
func (s *Service) DeleteByID(ctx context.Context, id int) (*entity.Product, error) {
	ctx = s.transactor.Begin(ctx)

	p, err := s.repository.DeleteByID(ctx, id)
	if err != nil {
		s.transactor.Rollback(ctx)
		return nil, errors.Wrap(err, "failed to delete product by id")
	}

	s.transactor.Commit(ctx)
	return p, nil
}
