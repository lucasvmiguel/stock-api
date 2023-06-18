// handler package is responsible to handle http requests of the product domain
package handler

import (
	"errors"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
	"github.com/lucasvmiguel/stock-api/pkg/pagination"
)

var (
	// field id is used as url param for different handlers (eg: get-by-id handler)
	FieldID = "id"
	// error when service is nil
	ErrNilSercice = errors.New("service cannot be nil")
	// error when json body is not valid
	ErrInvalidJSONBody = errors.New("invalid JSON body")
	// error when product was not found
	ErrNotFound = errors.New("product not found")
	// error internal server error
	ErrInternalServerError = errors.New("internal server error")
	// error when invalid limit query parameter is passed to the get paginated handler
	ErrInvalidLimitQueryParam = errors.New("invalid limit query parameter, it must be a integer")
	// error when invalid cursor query parameter is passed to the get paginated handler
	ErrInvalidCursorQueryParam = errors.New("invalid cursor query parameter, it must be a integer")
)

// generic pagination result type
// https://github.com/golang/mock/issues/621#issuecomment-1094351718
type GenericTypePaginationResult = *pagination.Result[*entity.Product]

// service interface to run different features
type Service interface {
	Create(product entity.Product) (*entity.Product, error)
	GetAll() ([]*entity.Product, error)
	GetByID(id uint) (*entity.Product, error)
	GetPaginated(cursor uint, limit uint) (GenericTypePaginationResult, error)
	UpdateByID(id uint, product entity.Product) (*entity.Product, error)
	DeleteByID(id uint) (*entity.Product, error)
}

// product handler that has methods to handle different types of http requests
type Handler struct {
	service                Service
	paginationDefaultLimit int
}

// new product handler arguments
type NewHandlerArgs struct {
	Service                Service
	PaginationDefaultLimit int
}

// creates a new product handler
func NewHandler(args NewHandlerArgs) (*Handler, error) {
	if args.Service == nil {
		return nil, ErrNilSercice
	}

	paginationDefaultLimit := args.PaginationDefaultLimit
	if paginationDefaultLimit == 0 {
		paginationDefaultLimit = 10
	}

	return &Handler{
		service:                args.Service,
		paginationDefaultLimit: args.PaginationDefaultLimit,
	}, nil
}
