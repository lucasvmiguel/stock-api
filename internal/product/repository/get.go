package repository

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
)

// gets all products from the database
func (r *Repository) GetAll() ([]*entity.Product, error) {
	products := []*entity.Product{}
	result := r.dbClient.Find(&products)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to get all products")
	}

	return products, nil
}

// gets a product by id from the database
func (r *Repository) GetByID(id uint) (*entity.Product, error) {
	product := &entity.Product{}
	result := r.dbClient.First(&product, id)
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
func (r *Repository) GetPaginated(cursor uint, limit uint) ([]*entity.Product, error) {
	products := []*entity.Product{}
	result := r.dbClient.Where("id > ?", int(cursor)).Limit(int(limit)).Find(&products)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to get products paginated")
	}

	return products, nil
}
