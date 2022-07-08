package handler

import "errors"

var (
	ErrNilDBClient = errors.New("db client cannot be nil")
)
