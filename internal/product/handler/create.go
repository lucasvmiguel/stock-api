package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
	"github.com/lucasvmiguel/stock-api/pkg/http/respond"
	"github.com/lucasvmiguel/stock-api/pkg/validator"
)

type createRequestBody struct {
	Name          string `validate:"required"`
	StockQuantity int    `validate:"numeric,min=0"`
}

// handles create product via http request
func (h *Handler) HandleCreate(w http.ResponseWriter, req *http.Request) {
	reqBody := createRequestBody{}

	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		respond.HTTPError(w, http.StatusBadRequest, errors.Wrap(err, ErrInvalidJSONBody.Error()))
		return
	}

	errs := validator.Validate(reqBody)
	if errs != nil {
		respond.HTTP(respond.Response{Body: errs, StatusCode: http.StatusBadRequest, Writer: w})
		return
	}

	product, err := h.repository.Create(entity.Product{
		Name:          reqBody.Name,
		StockQuantity: reqBody.StockQuantity,
	})
	if err != nil {
		respond.HTTPError(w, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	respond.HTTP(respond.Response{
		Body:       product,
		StatusCode: http.StatusCreated,
		Writer:     w,
	})
}
