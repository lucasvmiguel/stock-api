package repository

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

// gets all products from the database
func (r *Repository) GetAll(ctx context.Context) ([]*entity.Product, error) {
	products := []*entity.Product{}
	result := r.run(ctx).Order(idColumnName).Find(&products)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to get all products")
	}

	return products, nil
}

// gets a product by id from the database
func (r *Repository) GetByID(ctx context.Context, id int) (*entity.Product, error) {
	product := &entity.Product{}
	result := r.run(ctx).First(&product, id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to get product by id")
	}

	return product, nil
}

// gets products paginated from the database
// cursor is the last id from the previous page
func (r *Repository) GetPaginated(ctx context.Context, cursor int, limit int) ([]*entity.Product, error) {
	products := []*entity.Product{}
	result := r.run(ctx).Where("id > ?", int(cursor)).Limit(int(limit)).Order(idColumnName).Find(&products)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to get products paginated")
	}

	return products, nil
}
