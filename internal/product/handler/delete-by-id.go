package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lucasvmiguel/stock-api/pkg/http/respond"
	"github.com/lucasvmiguel/stock-api/pkg/parser"
)

// handles delete product by id via http request
func (h *Handler) HandleDeleteByID(w http.ResponseWriter, req *http.Request) {
	id, err := parser.StringToUint(chi.URLParam(req, FieldID))
	if err != nil {
		respond.HTTPError(w, http.StatusBadRequest, err)
		return
	}

	product, err := h.repository.DeleteByID(id)
	if err != nil {
		respond.HTTPError(w, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	if product == nil {
		respond.HTTPError(w, http.StatusNotFound, ErrNotFound)
		return
	}

	respond.HTTP(respond.Response{
		StatusCode: http.StatusNoContent,
		Writer:     w,
	})
}
