package handler

import (
	"net/http"

	"github.com/lucasvmiguel/stock-api/pkg/http/respond"
	"github.com/lucasvmiguel/stock-api/pkg/logger"
)

// handles get all products via http request
func (h *Handler) HandleGetAll(w http.ResponseWriter, req *http.Request) {
	logger := logger.HTTPLogEntry(req)

	products, err := h.service.GetAll()
	if err != nil {
		logger.Err(err).Msg(ErrInternalServerError.Error())
		respond.HTTPError(w, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	respond.HTTP(respond.Response{
		Body:       mapProductsToResponseBody(products),
		StatusCode: http.StatusOK,
		Writer:     w,
	})
}
