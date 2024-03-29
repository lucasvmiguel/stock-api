package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/lucasvmiguel/stock-api/pkg/http/respond"
	"github.com/lucasvmiguel/stock-api/pkg/logger"
	"github.com/lucasvmiguel/stock-api/pkg/validator"
)

// getPaginatedQueryParams is a struct that holds all query params for get paginated
type getPaginatedQueryParams struct {
	Limit  int `validate:"numeric,min=1,max=100"`
	Cursor int `validate:"numeric,min=0"`
}

// handles get all products via http request
func (h *Handler) HandleGetPaginated(w http.ResponseWriter, req *http.Request) {
	logger := logger.HTTPLogEntry(req)

	paginatedQueryParams, err := h.buildGetPaginatedQueryParams(req)
	if err != nil {
		respond.HTTPError(w, http.StatusBadRequest, err)
		return
	}

	errs := validator.Validate(paginatedQueryParams)
	if errs != nil {
		respond.HTTP(respond.Response{Body: errs, StatusCode: http.StatusBadRequest, Writer: w})
		return
	}

	result, err := h.service.GetPaginated(context.Background(), paginatedQueryParams.Cursor, paginatedQueryParams.Limit)
	if err != nil {
		logger.Err(err).Msg(ErrInternalServerError.Error())
		respond.HTTPError(w, http.StatusInternalServerError, ErrInternalServerError)
		return
	}

	respond.HTTP(respond.Response{
		Body:       mapProductsToPaginatedResponseBody(result),
		StatusCode: http.StatusOK,
		Writer:     w,
	})
}

func (h *Handler) buildGetPaginatedQueryParams(req *http.Request) (getPaginatedQueryParams, error) {
	paginatedQueryParams := getPaginatedQueryParams{}
	limitQueryParam := req.URL.Query().Get("limit")
	cursorQueryParam := req.URL.Query().Get("cursor")

	if limitQueryParam != "" {
		limit, err := strconv.Atoi(limitQueryParam)
		if err != nil {
			return paginatedQueryParams, ErrInvalidLimitQueryParam
		}
		paginatedQueryParams.Limit = limit
	} else {
		paginatedQueryParams.Limit = h.paginationDefaultLimit
	}

	if cursorQueryParam != "" {
		cursor, err := strconv.Atoi(cursorQueryParam)
		if err != nil {
			return paginatedQueryParams, ErrInvalidCursorQueryParam
		}
		paginatedQueryParams.Cursor = cursor
	}

	return paginatedQueryParams, nil
}
