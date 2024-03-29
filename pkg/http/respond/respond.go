package respond

import (
	"encoding/json"
	"net/http"
)

// Response is a param struct for the function HTTP
type Response struct {
	Body       interface{}
	StatusCode int
	Writer     http.ResponseWriter
}

type errorBody struct {
	Message string `json:"message"`
}

var (
	contentTypeKey   = "Content-Type"
	contentTypeValue = "application/json"
)

// helper function that respond a HTTP in json format
func HTTP(resp Response) {
	resp.Writer.Header().Set(contentTypeKey, contentTypeValue)
	resp.Writer.WriteHeader(resp.StatusCode)

	if resp.Body != nil {
		json.NewEncoder(resp.Writer).Encode(resp.Body)
	}
}

// helper function that respond an HTTP error in json format
func HTTPError(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set(contentTypeKey, contentTypeValue)

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorBody{Message: err.Error()})
}
