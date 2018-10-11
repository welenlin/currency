package api

import (
	"encoding/json"
	"net/http"
)

// STANDARDIZE API HANDLER

type JsonResponse struct {
	// Reserved field to add some meta information to the API response
	Meta interface{} `json:"meta,omitempty"`
	Data interface{} `json:"data"`
}

type JsonErrorResponse struct {
	Error *ApiError `json:"error"`
}

type ApiError struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc HandlerFunc
}

func ResponseJSON(w http.ResponseWriter, data JsonResponse) {
	jsonData, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func ErrorResponse(w http.ResponseWriter, err JsonErrorResponse) {
	jsonData, _ := json.Marshal(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Error.Status)
	w.Write(jsonData)
}

/*
HandlerFunc return function for API call
*/
type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

/*
ServeHTTP listen and serve request
*/
func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Close = true

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	if err := fn(w, r); err != nil {
		sendError(w, r, err)
	}
}

/*
sendError print response to browser
*/
func sendError(w http.ResponseWriter, r *http.Request, err error) {

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
	}
}
