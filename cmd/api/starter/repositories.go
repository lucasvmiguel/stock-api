package starter

import (
	"gorm.io/gorm"

	productRepository "github.com/lucasvmiguel/stock-api/internal/product/repository"
)

// repositories is a struct that holds all repositories
type repositories struct {
	product *productRepository.Repository
}

// createRepositoriesArgs is the arguments struct for createRepositories function
type createRepositoriesArgs struct {
	gormDB *gorm.DB
}

// createRepositories creates all repositories
func (s *Starter) createRepositories(args createRepositoriesArgs) (repositories, error) {
	productRepo, err := productRepository.NewRepository(args.gormDB)
	if err != nil {
		return repositories{}, err
	}

	return repositories{
		product: productRepo,
	}, nil
}
