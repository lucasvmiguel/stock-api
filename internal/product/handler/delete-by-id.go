package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/lucasvmiguel/stock-api/pkg/http/respond"
)

func (h *Handler) HandleDeleteByID(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(req, "id"), 10, 64)
	if err != nil {
		respond.HTTPError(w, http.StatusBadRequest, err)
		return
	}

	product, err := h.repository.DeleteByID(uint(id))
	if err != nil {
		respond.HTTPError(w, http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		respond.HTTPError(w, http.StatusNotFound, ErrNotFound)
	}

	respond.HTTP(respond.Response{
		StatusCode: http.StatusNoContent,
		Writer:     w,
	})
}
