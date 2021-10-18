package http_utils

import (
	"encoding/json"
  "github.com/katsun0921/bookstore_utils-go/rest_errors"
  "net/http"
)

func ResponseJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func ResponseError(w http.ResponseWriter, err rest_errors.RestErr) {
	ResponseJson(w, err.Status(), err)
}
