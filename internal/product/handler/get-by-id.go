package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/lucasvmiguel/stock-api/pkg/http/respond"
	"github.com/lucasvmiguel/stock-api/pkg/logger"
	"github.com/lucasvmiguel/stock-api/pkg/parser"
)

// handles get product by id via http request
func (h *Handler) HandleGetByID(w http.ResponseWriter, req *http.Request) {
	logger := logger.HTTPLogEntry(req)

	id, err := parser.StringToUint(chi.URLParam(req, FieldID))
	if err != nil {
		respond.HTTPError(w, http.StatusBadRequest, err)
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		logger.Err(err).Msg(ErrInternalServerError.Error())
		respond.HTTPError(w, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	if product == nil {
		respond.HTTPError(w, http.StatusNotFound, ErrNotFound)
		return
	}

	respond.HTTP(respond.Response{
		Body:       mapProductToResponseBody(product),
		StatusCode: http.StatusOK,
		Writer:     w,
	})
}
