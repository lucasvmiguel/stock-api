package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/lucasvmiguel/stock-api/pkg/http/respond"
	"github.com/lucasvmiguel/stock-api/pkg/logger"
)

// handles delete product by id via http request
func (h *Handler) HandleDeleteByID(w http.ResponseWriter, req *http.Request) {
	logger := logger.HTTPLogEntry(req)

	id, err := strconv.Atoi(chi.URLParam(req, FieldID))
	if err != nil {
		respond.HTTPError(w, http.StatusBadRequest, err)
		return
	}

	product, err := h.service.DeleteByID(context.Background(), id)
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
		StatusCode: http.StatusNoContent,
		Writer:     w,
	})
}
