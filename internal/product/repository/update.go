package repository

import (
	"context"

	"github.com/pkg/errors"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

// updates a product by id from the database
func (r *Repository) UpdateByID(ctx context.Context, id int, product entity.Product) (*entity.Product, error) {
	product.ID = id

	result := r.run(ctx).Updates(&product)
	if result.RowsAffected == 0 {
		return nil, errors.Wrap(result.Error, "no product could be found")
	}

	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to update product by id")
	}

	result = r.run(ctx).First(&product, product.ID)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to get product that was updated")
	}

	return &product, nil
}
