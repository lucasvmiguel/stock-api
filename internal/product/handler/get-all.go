package handler

import (
	"net/http"

	"github.com/lucasvmiguel/stock-api/pkg/http/respond"
)

func (h *Handler) HandleGetAll(w http.ResponseWriter, req *http.Request) {
	products, err := h.repository.GetAll()
	if err != nil {
		respond.HTTPError(w, http.StatusInternalServerError, err)
		return
	}

	respond.HTTP(respond.Response{
		Body:       products,
		StatusCode: http.StatusOK,
		Writer:     w,
	})
}
