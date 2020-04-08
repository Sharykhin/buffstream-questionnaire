package response

import (
	"encoding/json"
	"net/http"

	"github.com/mattn/go-nulltype"
)

type (
	//Response is a general response json representation
	Response struct {
		StatusCode int                 `json:"StatusCode"`
		Data       Data                `json:"Data"`
		Error      nulltype.NullString `json:"Error"`
		Meta       Meta                `json:"Meta"`
	}
	// Data keeps json results
	Data map[string]interface{}
	// Meta keep such meta information like pagination, some specific keys and so on
	Meta map[string]interface{}
)

//OK sends 200 response
func OK(w http.ResponseWriter, data Data, meta Meta) {
	response := newResponse(http.StatusOK, data, meta, nil)
	sendJSON(w, &response)
}

// Created sends 201 response
func Created(w http.ResponseWriter, data Data, meta Meta) {
	response := newResponse(http.StatusCreated, data, meta, nil)
	sendJSON(w, &response)
}

// Created sends 204 response
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// BadRequest sends 400 response
func BadRequest(w http.ResponseWriter, err error) {
	response := newResponse(http.StatusBadRequest, nil, nil, err)
	sendJSON(w, &response)
}

// ServerError sends 500 response
func ServerError(w http.ResponseWriter, err error) {
	response := newResponse(http.StatusInternalServerError, nil, nil, err)
	sendJSON(w, &response)
}

// NotFound sends 404 response
func NotFound(w http.ResponseWriter, err error) {
	response := newResponse(http.StatusNotFound, nil, nil, err)
	sendJSON(w, &response)
}

func sendJSON(w http.ResponseWriter, res *Response) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	_ = json.NewEncoder(w).Encode(&res)
}

func newResponse(code int, data Data, meta Meta, err error) Response {
	var errResponse nulltype.NullString
	if err != nil {
		errResponse = nulltype.NullStringOf(err.Error())
	}
	return Response{
		StatusCode: code,
		Data:       data,
		Meta:       meta,
		Error:      errResponse,
	}
}
