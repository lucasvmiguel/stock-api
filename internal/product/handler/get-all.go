package handler

import (
	"net/http"

	"github.com/lucasvmiguel/stock-api/internal/product/entity"
	"github.com/lucasvmiguel/stock-api/pkg/http/respond"
	"github.com/lucasvmiguel/stock-api/pkg/logger"
)

// GetAllResponseBody is the response body for get all products
type getAllResponseBody []productResponseBody

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
		Body:       h.buildGetAllResponseBody(products),
		StatusCode: http.StatusOK,
		Writer:     w,
	})
}

func (h *Handler) buildGetAllResponseBody(products []*entity.Product) getAllResponseBody {
	getAllResponseBody := getAllResponseBody{}

	for _, product := range products {
		getAllResponseBody = append(getAllResponseBody, h.buildProductResponseBody(product))
	}

	return getAllResponseBody
}
