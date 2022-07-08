package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/lucasvmiguel/stock-api/internal/product/repository"
	"github.com/lucasvmiguel/stock-api/pkg/http/respond"
	"gorm.io/gorm"
)

var (
	ErrNameCannotBeBlank    = errors.New("name cannot be blank")
	ErrInvalidStockQuantity = errors.New("invalid stock quantity")
)

type HandlerPost struct {
	dbClient *gorm.DB
}

type Product struct {
	Name          string `json:"name"`
	StockQuantity int    `json:"stock_quantity"`
}

type PostRequestBody Product

func NewHandlerPost(dbClient *gorm.DB) (*HandlerPost, error) {
	if ErrNilDBClient == nil {
		return nil, ErrNilDBClient
	}

	return &HandlerPost{dbClient: dbClient}, nil
}

func (h *HandlerPost) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	postReqBody := PostRequestBody{}

	err := json.NewDecoder(req.Body).Decode(&postReqBody)
	if err != nil {
		respond.HTTPError(w, http.StatusBadRequest, errors.Wrap(err, "Invalid JSON body"))
		return
	}

	err = h.validatePostRequestBody(postReqBody)
	if err != nil {
		respond.HTTPError(w, http.StatusBadRequest, err)
		return
	}

	createdProduct, err := h.createProduct(postReqBody)
	if err != nil {
		respond.HTTPError(w, http.StatusInternalServerError, err)
		return
	}

	respond.HTTP(respond.Response{
		Body:       createdProduct,
		Err:        err,
		StatusCode: http.StatusCreated,
		Writer:     w,
	})
}

func (h *HandlerPost) validatePostRequestBody(postReqBody PostRequestBody) error {
	if postReqBody.Name == "" {
		return ErrNameCannotBeBlank
	}

	if postReqBody.StockQuantity < 0 {
		return ErrInvalidStockQuantity
	}

	return nil
}

func (h *HandlerPost) createProduct(postReqBody PostRequestBody) (*repository.Product, error) {
	product := &repository.Product{Name: postReqBody.Name, StockQuantity: postReqBody.StockQuantity}
	result := h.dbClient.Create(product)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "failed to create product")
	}

	return product, nil
}
