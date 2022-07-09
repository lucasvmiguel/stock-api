package handler

import (
	"encoding/json"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
	"github.com/lucasvmiguel/stock-api/pkg/http/respond"
	"github.com/lucasvmiguel/stock-api/pkg/parser"
	"github.com/lucasvmiguel/stock-api/pkg/validator"
)

type updateRequestBody struct {
	Name          string `validate:"omitempty"`
	StockQuantity int    `validate:"omitempty,numeric,min=0"`
}

// handles product update via http request
func (h *Handler) HandleUpdate(w http.ResponseWriter, req *http.Request) {
	id, err := parser.StringToUint(chi.URLParam(req, FieldID))
	if err != nil {
		respond.HTTPError(w, http.StatusBadRequest, err)
		return
	}

	reqBody := updateRequestBody{}
	err = json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		respond.HTTPError(w, http.StatusBadRequest, errors.Wrap(err, ErrInvalidJSONBody.Error()))
		return
	}

	errs := validator.Validate(reqBody)
	if errs != nil {
		respond.HTTP(respond.Response{Body: errs, StatusCode: http.StatusBadRequest, Writer: w})
		return
	}

	product, err := h.repository.UpdateByID(id, entity.Product{
		Name:          reqBody.Name,
		StockQuantity: reqBody.StockQuantity,
	})
	if err != nil {
		respond.HTTPError(w, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	if product == nil {
		respond.HTTPError(w, http.StatusNotFound, ErrNotFound)
		return
	}

	respond.HTTP(respond.Response{
		Body:       product,
		StatusCode: http.StatusOK,
		Writer:     w,
	})
}
