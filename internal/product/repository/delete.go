package repository

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

// deletes a product by id from the database
func (r *Repository) DeleteByID(ctx context.Context, id int) (*entity.Product, error) {
	product := &entity.Product{}
	result := r.run(ctx).First(product, id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}

	result = r.run(ctx).Delete(product, id)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to delete product by id")
	}

	return product, nil
}
