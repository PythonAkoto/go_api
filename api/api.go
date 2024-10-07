package api

import (
	"encoding/json"
	"net/http"
)

// create structs that will act as our parameters for our endpoint & responses

// create coin balance params
type CoinBalanceParams struct {
	Username string
}

// create coin response params
type CoinBalanceResponse struct {
	// success code, usually 200
	Code int

	// account balance
	Balance int64
}

// error response
type Error struct {
	// error code
	Code int

	// error message
	Message string
}

// http error message - for the person who called it
func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	// what the user gets back in the response
	json.NewEncoder(w).Encode((resp))
}

// create the ability for different type of error
var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Enexpected Error Occyred", http.StatusInternalServerError)
	}
)
