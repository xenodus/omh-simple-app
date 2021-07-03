package models

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
}

func (r *JSONResponse) PrintJSONResponse(res http.ResponseWriter, statusCode int, message string, data interface{}) {
	r.StatusCode = statusCode
	r.Message = message
	r.Data = data

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(r.StatusCode)
	json.NewEncoder(res).Encode(r)
}

func (r *JSONResponse) PrintNotFoundResponse(res http.ResponseWriter) {
	r.PrintJSONResponse(res, http.StatusNotFound, "not found", nil)
}

func (r *JSONResponse) PrintUnexpectedErrorResponse(res http.ResponseWriter) {
	r.PrintJSONResponse(res, http.StatusInternalServerError, "unexpected error", nil)
}
